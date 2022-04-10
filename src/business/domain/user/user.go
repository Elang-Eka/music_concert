package user

import (
	log "azura-test/src/business/usecases/log"
	"azura-test/src/entity"
	"azura-test/src/utils/sql"
	"context"
)

type Interface interface {
	FindUser(ctx context.Context, email string) (entity.UserTicket, error)
	CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error)
	GetMyTicket(ctx context.Context, user entity.UserTicket) ([]entity.Ticket, error)
}

type user struct {
	log log.Logger
	db  sql.Interface
}

func Init(log log.Logger, db sql.Interface) Interface {
	return &user{
		log: log,
		db:  db,
	}
}

func (u *user) CreateUser(ctx context.Context, user entity.User, trx_id int) (entity.User, error) {
	tx, err := u.db.Leader().Begin(ctx)
	if err != nil {
		return entity.User{}, err
	}

	res, err := tx.Exec(addUser, user.Name, user.Age, user.Gender, user.Email, trx_id)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return entity.User{}, err
		}
		return entity.User{}, err
	}
	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		if err := tx.Rollback(); err != nil {
			return entity.User{}, err
		}
		return entity.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return entity.User{}, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return entity.User{}, err
	}

	user.Id = int(lid)
	user.Id_Trx = trx_id

	return user, nil
}

func (u *user) GetMyTicket(ctx context.Context, user entity.UserTicket) ([]entity.Ticket, error) {
	rows, err := u.db.Leader().Query(ctx, getUserTicket, user.Email)
	if err != nil {
		return nil, err
	}

	Ticket := []entity.Ticket{}
	for rows.Next() {
		ticket := entity.Ticket{}
		if err = rows.StructScan(&ticket); err != nil {
			u.log.LogError(err.Error())
			return nil, err
		}
		Ticket = append(Ticket, ticket)
	}

	return Ticket, nil
}

func (u *user) FindUser(ctx context.Context, email string) (entity.UserTicket, error) {
	rows, err := u.db.Leader().Query(ctx, getUser, email)
	if err != nil {
		return entity.UserTicket{}, err
	}

	user := entity.UserTicket{}
	rows.Next()
	if err = rows.StructScan(&user); err != nil {
		u.log.LogError(err.Error())
	}

	return user, nil
}
