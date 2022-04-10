package usecases

import (
	"azura-test/src/business/domain"
	"azura-test/src/business/usecases/event"
	log "azura-test/src/business/usecases/log"
	"azura-test/src/business/usecases/transaction"
	"azura-test/src/business/usecases/user"
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
