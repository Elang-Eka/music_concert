package rest

import (
	"azura-test/src/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// @Summary Update Transaction
// @Description This API to input payment from user
// @Tags Ticket
// @Accept json
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Transaction{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /payment [PUT]
func (r *rest) Payment(ctx *gin.Context) {
	var pay entity.Transaction
	if err := ctx.ShouldBindJSON(&pay); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := r.Validation(ctx, pay)
	if err != nil {
		r.httpRespError(ctx, 500, err)
	} else {
		r.httpRespSuccess(ctx, 200, "successful payment")
	}
}

func (r *rest) Validation(ctx *gin.Context, pay entity.Transaction) (entity.Transaction, error) {
	data, err := r.GetTransaction(ctx, pay)
	if err != nil {
		return entity.Transaction{}, err
	}

	t_now := time.Now().In(time.UTC).Unix()
	trx_time := data.Date.Unix() - 25200
	diff := t_now - trx_time
	if diff > 1800 {
		reject, err := r.UpdateTransaction(ctx, data, "rejected")
		if err != nil {
			return entity.Transaction{}, err
		}
		return reject, errors.New("Event Ticket Expired! Please order again")
	} else {
		accept, err := r.UpdateTransaction(ctx, data, "accepted")
		if err != nil {
			return entity.Transaction{}, err
		}
		return accept, nil
	}
}
