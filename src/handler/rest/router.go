package rest

import (
	"fmt"
	"golang-heroku/docs/swagger"
	"golang-heroku/src/business/usecases"
	logCase "golang-heroku/src/business/usecases/log"
	"golang-heroku/src/utils/config"
	"golang-heroku/src/utils/configreader"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type REST interface {
	Run()
}

type rest struct {
	http         *gin.Engine
	conf         config.Application
	configreader configreader.Interface
	log          logCase.Logger
	uc           *usecases.Usecases
}

func Init(conf config.Application, configreader configreader.Interface, log logCase.Logger, uc *usecases.Usecases) REST {
	gin.SetMode("")
	httpServer := gin.New()
	r := &rest{
		http:         httpServer,
		conf:         conf,
		configreader: configreader,
		log:          log,
		uc:           uc,
	}
	r.Register()

	return r
}

func (r *rest) Register() {
	r.http.GET("/ping", r.Ping)
	r.http.GET("/event", r.GetList)
	r.http.POST("/booking", r.PostBook)
	r.http.PUT("/payment", r.Payment)
	r.http.GET("/ticket/:gmail", r.MyTicket)
	r.registerSwaggerRoutes()
}

func (r *rest) Run() {
	if err := r.http.Run(":8080"); err != nil {
		r.log.LogError(err.Error())
	}
}

func (r *rest) registerSwaggerRoutes() {
	if r.conf.Gin.Swagger.Enabled {
		swagger.SwaggerInfo.Title = r.conf.Meta.Title
		swagger.SwaggerInfo.Description = r.conf.Meta.Description
		swagger.SwaggerInfo.Version = r.conf.Meta.Version
		swagger.SwaggerInfo.Host = r.conf.Meta.Host
		swagger.SwaggerInfo.BasePath = r.conf.Meta.BasePath

		r.http.GET(fmt.Sprintf("%s/*any", r.conf.Gin.Swagger.Path),
			ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
