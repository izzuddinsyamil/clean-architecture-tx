package repo

import (
	"context"
	"database/sql"
	"fmt"
	errCode "repo-pattern-w-trx-management/error"
	"repo-pattern-w-trx-management/model"

	"github.com/palantir/stacktrace"
)

type Repository interface {
	Atomic(ctx context.Context, fn func(r Repository) error) error
	GetUser(ctx context.Context) (u []model.User, err error)
	CreateUser(ctx context.Context, name string, balance int) (err error)
	DeductBalance(ctx context.Context, userId, amount int) (err error)
	AddBalance(ctx context.Context, userId, amount int) (err error)
	CreateTransaction(ctx context.Context, senderId, receiverId, amount int) (err error)
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repo struct {
	conn *sql.DB
	db   DBTX
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		conn: db,
		db:   db,
	}
}

func (r *repo) Atomic(ctx context.Context, fn func(r Repository) error) (err error) {
	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}

		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	newRepo := &repo{
		conn: r.conn,
		db:   tx,
	}

	err = fn(newRepo)
	return
}

func (r *repo) GetUser(ctx context.Context) (u []model.User, err error) {
	var (
		tracer = "repo.GetUser"
		q      = `
		select id, name, balance
		from users
		order by id asc`
	)

	rows, err := r.conn.QueryContext(ctx, q)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user model.User
		if errScan := rows.Scan(&user.Id, &user.Name, &user.Balance); err != nil {
			err = stacktrace.PropagateWithCode(errScan, errCode.EcodeInternal, tracer)
			return
		}

		u = append(u, user)
	}

	return
}

func (r *repo) CreateUser(ctx context.Context, name string, balance int) (err error) {
	var (
		tracer = "repo.CreateUser"
		q      = `
		insert into users (name, balance) values ($1, $2)`
	)

	_, err = r.conn.ExecContext(ctx, q, name, balance)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}

func (r *repo) DeductBalance(ctx context.Context, userId, amount int) (err error) {
	var (
		tracer = "repo.DeductBalance"
		q      = `
		update users
		set balance = balance - $1
		where id = $2`
	)

	_, err = r.db.ExecContext(ctx, q, amount, userId)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}

func (r *repo) AddBalance(ctx context.Context, userId, amount int) (err error) {
	var (
		tracer = "repo.AddBalance"
		q      = `
		update users
		set balance = balance + $1
		where id = $2`
	)

	_, err = r.db.ExecContext(ctx, q, amount, userId)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}

func (r *repo) CreateTransaction(ctx context.Context, senderId, receiverId, amount int) (err error) {
	var (
		tracer = "repo.AddBalance"
		q      = `
		insert into transactions (amount, receiver_id, sender_id) values ($1, $2, $3)`
	)

	_, err = r.db.ExecContext(ctx, q, amount, receiverId, senderId)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}
