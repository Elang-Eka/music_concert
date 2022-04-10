package rest

import (
	"azura-test/src/entity"

	"github.com/gin-gonic/gin"
)

// @Summary Get Event
// @Description Get list of event
// @Tags Event
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.Event{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /event [GET]
func (r *rest) GetList(ctx *gin.Context) {
	data, err := r.uc.Event.FindAll(ctx)
	if err != nil {
		r.log.LogError(err.Error())
		r.httpRespError(ctx, 500, err)
	}

	r.httpRespSuccess(ctx, 200, data)
}

func (r *rest) GetEvent(ctx *gin.Context, code int) (entity.Event, error) {
	event, err := r.uc.Event.GetEvent(ctx, code)
	if err != nil {
		r.log.LogError(err.Error())
		return entity.Event{}, err
	}

	return event, nil
}
