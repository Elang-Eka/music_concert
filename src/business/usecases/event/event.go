package event

import (
	eventDom "azura-test/src/business/domain/event"
	log "azura-test/src/business/usecases/log"
	"azura-test/src/entity"
	"context"
)

type Interface interface {
	FindAll(ctx context.Context) ([]entity.Event, error)
	GetEvent(ctx context.Context, code int) (entity.Event, error)
	// CreateTransaction(ctx context.Context, trx entity.Transaction) (entity.Transaction, error)
	// CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error)
}

type event struct {
	log log.Logger
	dom eventDom.Interface
}

func Init(log log.Logger, dom eventDom.Interface) Interface {
	return &event{
		log: log,
		dom: dom,
	}
}

func (e *event) FindAll(ctx context.Context) ([]entity.Event, error) {
	e.log.LogAccess("Get list of event")

	return e.dom.FindAll(ctx)
}

func (e *event) GetEvent(ctx context.Context, code int) (entity.Event, error) {
	e.log.LogAccess("Get trasaction_id")

	return e.dom.GetEvent(ctx, code)
}

// func (e *event) CreateTransaction(ctx context.Context, trx entity.Transaction) (entity.Transaction, error) {
// 	e.log.LogAccess("Create Transaction")

// 	return e.dom.CreateTransaction(ctx, trx)
// }
// func (e *event) CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error) {
// 	e.log.LogAccess("Create User")

// 	return e.dom.CreateUser(ctx, user, trx_id)
// }
