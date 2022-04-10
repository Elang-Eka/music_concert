package domain

import (
	"azura-test/src/business/domain/event"
	"azura-test/src/business/domain/transaction"
	"azura-test/src/business/domain/user"
	log "azura-test/src/business/usecases/log"
	"azura-test/src/utils/sql"
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
