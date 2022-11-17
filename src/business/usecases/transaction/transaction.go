package transaction

import (
	"context"
	transactionDom "golang-heroku/src/business/domain/transaction"
	log "golang-heroku/src/business/usecases/log"
	"golang-heroku/src/entity"
)

type Interface interface {
	CreateTransaction(ctx context.Context, trx entity.Transaction, event entity.Event) (entity.Transaction, error)
	UpdateTransaction(ctx context.Context, trx entity.Transaction, action string) (entity.Transaction, error)
	GetTrx(ctx context.Context, code int) (entity.Transaction, error)
}

type transaction struct {
	log log.Logger
	dom transactionDom.Interface
}

func Init(log log.Logger, dom transactionDom.Interface) Interface {
	return &transaction{
		log: log,
		dom: dom,
	}
}

func (t *transaction) CreateTransaction(ctx context.Context, trx entity.Transaction, event entity.Event) (entity.Transaction, error) {
	t.log.LogAccess("Create Transaction")

	return t.dom.CreateTransaction(ctx, trx, event)
}

func (t *transaction) UpdateTransaction(ctx context.Context, trx entity.Transaction, action string) (entity.Transaction, error) {
	t.log.LogAccess("Update Transaction")

	return t.dom.UpdateTransaction(ctx, trx, action)
}

func (t *transaction) GetTrx(ctx context.Context, code int) (entity.Transaction, error) {
	t.log.LogAccess("Get Data Transaction")

	return t.dom.GetTrx(ctx, code)

}
