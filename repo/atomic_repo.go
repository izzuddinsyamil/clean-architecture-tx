package repo

import (
	"context"
	"database/sql"
	"fmt"
	errCode "repo-pattern-w-trx-management/error"
	"repo-pattern-w-trx-management/model"

	"github.com/palantir/stacktrace"
)

type AtomicRepository interface {
	Atomic(ctx context.Context, fn func(r AtomicRepository) error) error
	GetUser(ctx context.Context) (u []model.User, err error)
	GetUserById(ctx context.Context, id int) (u model.User, err error)
	CreateUser(ctx context.Context, name string, balance int) (err error)
	DeductBalance(ctx context.Context, userId, amount int) (err error)
	AddBalance(ctx context.Context, userId, amount int) (err error)
}

// this interface is automatically generated by 'sqlc' and it can be found
// inside the 'db.go' file. It's implemented by both '*sql.DB' and '*sql.TX'
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// this func is automatically generated by 'sqlc' and it can be found
// inside the 'db.go' file. It accepts both '*sql.DB' and '*sql.TX' as
// input.
func New(db DBTX) *Queries {
	return &Queries{db: db}
}

// this struct is automatically generated by 'sqlc' and it can be found
// inside the 'db.go' file.
type Queries struct {
	db DBTX
}

type atomicRepo struct {
	conn *sql.DB
	db   *Queries
}

func NewAtomicRepo(db *sql.DB) AtomicRepository {
	return &atomicRepo{
		conn: db,
		db:   New(db),
	}
}

func (r *atomicRepo) Atomic(ctx context.Context, fn func(r AtomicRepository) error) (err error) {
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

	newRepo := &atomicRepo{
		conn: r.conn,
		db:   New(tx),
	}

	err = fn(newRepo)
	return
}

func (r *atomicRepo) GetUser(ctx context.Context) (u []model.User, err error) {
	var (
		tracer = "repo.GetUser"
		q      = `
		select id, name, balance
		from users`
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

func (r *atomicRepo) GetUserById(ctx context.Context, id int) (u model.User, err error) {
	var (
		tracer = "repo.GetUserById"
		q      = `
		select id, name, balance
		from users
		where id = $1`
	)

	err = r.conn.QueryRowContext(ctx, q, id).Scan(&u.Id, &u.Name, &u.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			err = stacktrace.PropagateWithCode(err, errCode.EcodeNotFound, tracer)
			return
		}

		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}
	return
}

func (r *atomicRepo) CreateUser(ctx context.Context, name string, balance int) (err error) {
	var (
		tracer = "repo.CreateUser"
		q      = `
		INSERT INTO users (name, balance) values ($1, $2)`
	)

	_, err = r.conn.ExecContext(ctx, q, name, balance)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}

func (r *atomicRepo) DeductBalance(ctx context.Context, userId, amount int) (err error) {
	var (
		tracer = "repo.DeductBalance"
		q      = `
		update users
		set balance = balance - $1
		where id = $2`
	)

	_, err = r.db.db.ExecContext(ctx, q, amount, userId)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}

func (r *atomicRepo) AddBalance(ctx context.Context, userId, amount int) (err error) {
	var (
		tracer = "repo.AddBalance"
		q      = `
		update users
		set balance = balance + $1
		where id = $2`
	)

	_, err = r.db.db.ExecContext(ctx, q, amount, userId)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}
