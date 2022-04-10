package rest

import (
	"azura-test/src/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Booking Ticket
// @Description This api to book ticket concert
// @Tags Ticket
// @Accept json
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Transaction{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /booking [POST]
func (r *rest) PostBook(ctx *gin.Context) {
	var trx entity.Transaction
	if err := ctx.ShouldBindJSON(&trx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := r.GetEvent(ctx, trx.Event)
	if err != nil {
		r.log.LogError(err.Error())
		r.httpRespError(ctx, 500, err)
	}

	data, err := r.uc.Transaction.CreateTransaction(ctx, trx, event)

	if err != nil {
		r.log.LogError(err.Error())
		r.httpRespError(ctx, 500, err)
	}
	data = r.CreateUser(ctx, data)

	r.httpRespSuccess(ctx, 200, data)
}

func (r *rest) GetTransaction(ctx *gin.Context, trx entity.Transaction) (entity.Transaction, error) {
	data, err := r.uc.Transaction.GetTrx(ctx, trx.Code)
	if err != nil {
		r.log.LogError(err.Error())
		return entity.Transaction{}, err
	}

	return data, nil
}

func (r *rest) UpdateTransaction(ctx *gin.Context, trx entity.Transaction, action string) (entity.Transaction, error) {
	data, err := r.uc.Transaction.UpdateTransaction(ctx, trx, action)
	if err != nil {
		r.log.LogError(err.Error())
		r.httpRespError(ctx, 500, err)
	}

	return data, nil
}
