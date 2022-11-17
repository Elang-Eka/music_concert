package domain

import (
	"golang-heroku/src/business/domain/event"
	"golang-heroku/src/business/domain/transaction"
	"golang-heroku/src/business/domain/user"
	log "golang-heroku/src/business/usecases/log"
	"golang-heroku/src/utils/sql"
)

type Domains struct {
	Event       event.Interface
	Transaction transaction.Interface
	User        user.Interface
}

func Init(log log.Logger, db sql.Interface) *Domains {
	return &Domains{
		Event:       event.Init(log, db),
		Transaction: transaction.Init(log, db),
		User:        user.Init(log, db),
	}
}
