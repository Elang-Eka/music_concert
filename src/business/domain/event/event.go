package event

import (
	log "azura-test/src/business/usecases/log"
	"azura-test/src/entity"
	"azura-test/src/utils/sql"
	"context"
	"fmt"
)

type Interface interface {
	FindAll(ctx context.Context) ([]entity.Event, error)
	GetEvent(ctx context.Context, code int) (entity.Event, error)
	// CreateTransaction(ctx context.Context, trx entity.Transaction) (entity.Transaction, error)
	// CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error)
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
		fmt.Println("Ada error bosq : ", err)
		return nil, err
	}
	Event := []entity.Event{}
	for rows.Next() {
		event := entity.Event{}
		if err := rows.StructScan(&event); err != nil {
			e.log.LogError(err.Error())
			fmt.Println("Ada error bosq hmm : ", err)
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
		fmt.Println("Akhirnya ketemu error")
		e.log.LogError(err.Error())
		return
	}

	return trx, nil
}

// func (e *event) CreateTransaction(ctx context.Context, trx entity.Transaction) (entity.Transaction, error) {
// 	event, err := e.GetEvent(ctx, trx.Event)
// 	if err != nil {
// 		return entity.Transaction{}, err
// 	}

// 	price := event.Price * trx.Quantity
// 	tx, err := e.db.Leader().Begin(ctx)
// 	if err != nil {
// 		fmt.Println("Ada error bosq : ", err)
// 		return entity.Transaction{}, err
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	code := rand.Intn(100000)
// 	res, err := tx.Exec(addTransaction, trx.Event, trx.Quantity, price, trx.PayMethod, code)
// 	if err != nil {
// 		fmt.Println("Ada error bosq : ", err)
// 		if err := tx.Rollback(); err != nil {
// 			return entity.Transaction{}, err
// 		}
// 		return entity.Transaction{}, err
// 	}

// 	num, err := res.RowsAffected()
// 	if err != nil || num < 1 {
// 		if err := tx.Rollback(); err != nil {
// 			return entity.Transaction{}, err
// 		}
// 		return entity.Transaction{}, err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return entity.Transaction{}, err
// 	}

// 	lid, err := res.LastInsertId()
// 	if err != nil {
// 		return entity.Transaction{}, err
// 	}

// 	trx.TPrice = price
// 	trx.Id = int(lid)
// 	trx.Code = code

// 	return trx, nil
// }

// func (e *event) CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error) {
// 	tx, err := e.db.Leader().Begin(ctx)
// 	if err != nil {
// 		return entity.User{}, err
// 	}
// 	// fmt.Println(user)
// 	// fmt.Println("Ini user yang tersedia :", user)
// 	res, err := tx.Exec(addUser, user.Name, user.Age, user.Gender, user.Email, trx_id)
// 	if err != nil {
// 		// fmt.Println("Ada error bosq : ", err)
// 		if err := tx.Rollback(); err != nil {
// 			// fmt.Println("Ada error bosq : ", err)
// 			return entity.User{}, err
// 		}
// 		return entity.User{}, err
// 	}
// 	num, err := res.RowsAffected()
// 	if err != nil || num < 1 {
// 		// fmt.Println("Ada error bosq : ", err)
// 		if err := tx.Rollback(); err != nil {
// 			// fmt.Println("Ada error bosq : ", err)
// 			return entity.User{}, err
// 		}
// 		return entity.User{}, err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		// fmt.Println("Ada error bosq : ", err)
// 		return entity.User{}, err
// 	}

// 	lid, err := res.LastInsertId()
// 	if err != nil {
// 		return entity.User{}, err
// 	}

// 	user.Id = int(lid)
// 	user.Id_Trx = trx_id

// 	return user, nil
// }
