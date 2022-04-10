package rest

import (
	"azura-test/src/entity"

	"github.com/gin-gonic/gin"
)

func (r *rest) CreateUser(ctx *gin.Context, trx entity.Transaction) entity.Transaction {
	len := len(trx.User)
	for i := 0; i < len; i++ {
		user := trx.User[i]
		Usr, err := r.uc.User.CreateUser(ctx, user, trx.Id)
		if err != nil {
			r.log.LogError(err.Error())
			r.httpRespError(ctx, 500, err)
		}
		trx.User[i] = Usr
	}

	return trx
}
