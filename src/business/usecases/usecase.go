package usecases

import (
	"golang-heroku/src/business/domain"
	"golang-heroku/src/business/usecases/event"
	log "golang-heroku/src/business/usecases/log"
	"golang-heroku/src/business/usecases/transaction"
	"golang-heroku/src/business/usecases/user"
)

type Usecases struct {
	Event       event.Interface
	Transaction transaction.Interface
	User        user.Interface
}

func Init(log log.Logger, d *domain.Domains) *Usecases {
	return &Usecases{
		Event:       event.Init(log, d.Event),
		Transaction: transaction.Init(log, d.Transaction),
		User:        user.Init(log, d.User),
	}
}
