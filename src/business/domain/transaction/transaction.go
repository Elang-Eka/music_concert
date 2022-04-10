package transaction

import (
	log "azura-test/src/business/usecases/log"
	"azura-test/src/entity"
	"azura-test/src/utils/sql"
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Interface interface {
	CreateTransaction(ctx context.Context, trx entity.Transaction, event entity.Event) (entity.Transaction, error)
	UpdateTransaction(ctx context.Context, trx entity.Transaction, action string) (entity.Transaction, error)

	GetTrx(ctx context.Context, code int) (entity.Transaction, error)
}

type transaction struct {
	log log.Logger
	db  sql.Interface
}

func Init(log log.Logger, db sql.Interface) Interface {
	return &transaction{
		log: log,
		db:  db,
	}
}

func (t *transaction) CreateTransaction(ctx context.Context, trx entity.Transaction, event entity.Event) (entity.Transaction, error) {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(100000)
	unixCode := rand.Intn(1000)
	price := (event.Price * trx.Quantity) + unixCode
	tx, err := t.db.Leader().Begin(ctx)
	if err != nil {
		return entity.Transaction{}, err
	}

	res, err := tx.Exec(addTransaction, trx.Event, trx.Quantity, price, trx.PayMethod, code)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return entity.Transaction{}, err
		}
		return entity.Transaction{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		if err := tx.Rollback(); err != nil {
			return entity.Transaction{}, err
		}
		return entity.Transaction{}, err
	}

	if err := tx.Commit(); err != nil {
		return entity.Transaction{}, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return entity.Transaction{}, err
	}

	trx.TPrice = price
	trx.Id = int(lid)
	trx.Code = code
	trx.Action = "waiting"

	return trx, nil
}

func (t *transaction) UpdateTransaction(ctx context.Context, trx entity.Transaction, action string) (entity.Transaction, error) {
	tx, err := t.db.Leader().Begin(ctx)
	if err != nil {
		return entity.Transaction{}, err
	}

	res, err := tx.Exec(updateTransaction, action, trx.Event, trx.TPrice, trx.PayMethod, trx.Code)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return entity.Transaction{}, err
		}
		return entity.Transaction{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		if err := tx.Rollback(); err != nil {
			return entity.Transaction{}, err
		}
		return entity.Transaction{}, err
	}

	if err := tx.Commit(); err != nil {
		return entity.Transaction{}, err
	}

	trx.Action = action
	return trx, nil
}

// func (t *transaction) AcceptTransaction(ctx context.Context, trx entity.Transaction, event entity.Event) (entity.Transaction, error) {
// 	rand.Seed(time.Now().UnixNano())
// 	code := rand.Intn(100000)
// 	unixCode := rand.Intn(1000)
// 	price := (event.Price * trx.Quantity) + unixCode
// 	tx, err := t.db.Leader().Begin(ctx)
// 	if err != nil {
// 		fmt.Println("Ada error bosq : ", err)
// 		return entity.Transaction{}, err
// 	}

// 	res, err := tx.Exec(accept, trx.Event, trx.Quantity, price, trx.PayMethod, code)
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

// func (t *transaction) RejectTransaction(ctx context.Context, trx entity.Transaction, event entity.Event) (entity.Transaction, error) {
// 	rand.Seed(time.Now().UnixNano())
// 	code := rand.Intn(100000)
// 	unixCode := rand.Intn(1000)
// 	price := (event.Price * trx.Quantity) + unixCode
// 	tx, err := t.db.Leader().Begin(ctx)
// 	if err != nil {
// 		fmt.Println("Ada error bosq : ", err)
// 		return entity.Transaction{}, err
// 	}

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

func (t *transaction) GetTrx(ctx context.Context, code int) (entity.Transaction, error) {
	rows, err := t.db.Leader().Query(ctx, getTransaction, code)
	if err != nil {
		fmt.Println("Ada error bosq : ", err)
		return entity.Transaction{}, err
	}
	Transaction := entity.Transaction{}
	rows.Next()
	transaction := entity.Transaction{}
	if err := rows.StructScan(&transaction); err != nil {
		t.log.LogError(err.Error())
		fmt.Println("Ada error bosq : ", err)
	}
	Transaction = transaction

	return Transaction, nil
}
