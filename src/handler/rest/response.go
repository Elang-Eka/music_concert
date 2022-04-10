package rest

import (
	"azura-test/src/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (r *rest) httpRespError(ctx *gin.Context, StatusCode int, err error) {
	statusStr := http.StatusText(StatusCode)
	errResp := &entity.HTTPResp{
		Message: entity.HTTPMessage{
			Title: "Error",
			Body:  "Request failure",
		},
		Meta: entity.Meta{
			Path:       r.conf.Meta.Host + ctx.Request.URL.String(),
			StatusCode: StatusCode,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", ctx.Request.Method, ctx.Request.URL.RequestURI(), StatusCode, statusStr),
			Error: &entity.MetaError{
				Code:    StatusCode,
				Message: err.Error(),
			},
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}
	r.log.LogError(err.Error())
	ctx.Header("x-request-id", "RequestId")
	ctx.AbortWithStatusJSON(StatusCode, errResp)
}

func (r *rest) httpRespSuccess(ctx *gin.Context, StatusCode int, data interface{}) {
	meta := entity.Meta{
		Path:       r.conf.Meta.Host + ctx.Request.URL.String(),
		StatusCode: StatusCode,
		Status:     http.StatusText(StatusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", ctx.Request.Method, ctx.Request.URL.RequestURI(), StatusCode, http.StatusText(StatusCode)),
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	resp := &entity.HTTPResp{
		Message: entity.HTTPMessage{
			Title: "OK",
			Body:  "Request successful",
		},
		Meta: meta,
		Data: data,
	}

	raw, err := json.Marshal(&resp)

	if err != nil {
		r.httpRespError(ctx, 404, err)
		return
	}

	ctx.Header("x-request-id", "RequestId")
	ctx.Data(StatusCode, "application/json", raw)
}

func (r *rest) Bind(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindWith(obj, binding.Default(ctx.Request.Method, ctx.ContentType()))
}

// @Summary Health Check
// @Description This endpoint will hit the server
// @Tags Server
// @Produce json
// @Success 200 string example="PONG!"
// @Router /ping [GET]
func (r *rest) Ping(ctx *gin.Context) {
	r.httpRespSuccess(ctx, 200, "PONG!")
}
