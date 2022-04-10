package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// @Summary Get Ticket user
// @Description For users to know the number of tickets
// @Tags Ticket
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.UserTicket{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /ticket/{email} [GET]
func (r *rest) MyTicket(ctx *gin.Context) {
	gmail := ctx.Param("gmail")
	user, err := r.uc.User.FindUser(ctx, gmail)
	if err != nil {
		r.httpRespError(ctx, 200, errors.New("User not found"))
	} else {
		data, err := r.uc.User.GetMyTicket(ctx, user)
		if len(data) != 0 {
			user.Ticket = append(user.Ticket, data...)
			if err != nil {
				r.log.LogError(err.Error())
				r.httpRespError(ctx, 200, err)
			}
			r.httpRespSuccess(ctx, 200, user)
		} else {
			r.httpRespError(ctx, 500, errors.New("You don't have ticket / maybe you haven't completed the payment"))
		}
	}
}
