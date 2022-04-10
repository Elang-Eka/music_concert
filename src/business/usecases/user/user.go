package user

import (
	userDom "azura-test/src/business/domain/user"
	log "azura-test/src/business/usecases/log"
	"azura-test/src/entity"
	"context"
)

type Interface interface {
	FindUser(ctx context.Context, email string) (entity.UserTicket, error)
	CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error)
	GetMyTicket(ctx context.Context, user entity.UserTicket) ([]entity.Ticket, error)
}

type user struct {
	log log.Logger
	dom userDom.Interface
}

func Init(log log.Logger, dom userDom.Interface) Interface {
	return &user{
		log: log,
		dom: dom,
	}
}

func (u *user) CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error) {
	u.log.LogAccess("Create User")

	return u.dom.CreateUser(ctx, user, trx_id)
}

func (u *user) FindUser(ctx context.Context, gmail string) (entity.UserTicket, error) {
	u.log.LogAccess("GetUserID")

	return u.dom.FindUser(ctx, gmail)
}

func (u *user) GetMyTicket(ctx context.Context, user entity.UserTicket) ([]entity.Ticket, error) {
	u.log.LogAccess("Get User Ticket")

	return u.dom.GetMyTicket(ctx, user)
}
