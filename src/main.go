package main

import (
	"fmt"
	"golang-heroku/src/business/domain"
	"golang-heroku/src/business/usecases"
	"golang-heroku/src/handler/rest"
	"golang-heroku/src/utils/config"
	"golang-heroku/src/utils/configreader"
	logger "golang-heroku/src/utils/logger"
	"golang-heroku/src/utils/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("HAIIII")
	cfg := config.Init()
	confreader := configreader.Init(configreader.Options{
		Name: "conf",
		Type: "yaml",
		Path: "../etc/cfg",
	})
	confreader.ReadConfig(&cfg)

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// if !files.IsExist(godotenv.Load("../.env")) {

	// }

	cfg.SQL.Driver = os.Getenv("DB_DRIVER")
	cfg.SQL.User = os.Getenv("DB_USER")
	cfg.SQL.Password = os.Getenv("DB_PASS")
	cfg.SQL.Host = os.Getenv("DB_HOST")
	cfg.SQL.DB = os.Getenv("DB_NAME")
	cfg.SQL.Port = os.Getenv("DB_PORT")

	fmt.Printf("%+v\n", cfg)

	log := logger.Init()

	db := sql.Init(cfg.SQL, log)

	d := domain.Init(log, db)

	uc := usecases.Init(log, d)

	r := rest.Init(cfg, confreader, log, uc)
	r.Run()
}
