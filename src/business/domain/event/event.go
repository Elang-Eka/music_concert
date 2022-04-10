package event

import (
	log "azura-test/src/business/usecases/log"
	"azura-test/src/entity"
	"azura-test/src/utils/sql"
	"context"
)

type Interface interface {
	FindAll(ctx context.Context) ([]entity.Event, error)
	GetEvent(ctx context.Context, code int) (entity.Event, error)
}

type event struct {
	log log.Logger
	db  sql.Interface
}

func Init(log log.Logger, db sql.Interface) Interface {
	return &event{
		log: log,
		db:  db,
	}
}

func (e *event) FindAll(ctx context.Context) ([]entity.Event, error) {
	rows, err := e.db.Leader().Query(ctx, findAll)
	if err != nil {
		return nil, err
	}
	Event := []entity.Event{}
	for rows.Next() {
		event := entity.Event{}
		if err := rows.StructScan(&event); err != nil {
			e.log.LogError(err.Error())
			continue
		}
		Event = append(Event, event)
	}

	return Event, nil
}

func (e *event) GetEvent(ctx context.Context, code int) (event entity.Event, err error) {
	rows, err := e.db.Leader().Query(ctx, getEvent, code)
	if err != nil {
		return
	}

	trx := entity.Event{}
	rows.Next()
	if err = rows.StructScan(&trx); err != nil {
		e.log.LogError(err.Error())
		return
	}

	return trx, nil
}
