package repo

import (
	"context"
	"database/sql"
	errCode "repo-pattern-w-trx-management/error"
	"repo-pattern-w-trx-management/model"

	"github.com/palantir/stacktrace"
)

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *repo {
	return &repo{db: db}
}

func (r *repo) GetUser(ctx context.Context) (u []model.User, err error) {
	var (
		tracer = "repo.GetUser"
		q      = `
		select id, name, balance
		from users
		order by id asc`
	)

	rows, err := r.db.QueryContext(ctx, q)
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

func (r *repo) GetUserById(ctx context.Context, id int) (u model.User, err error) {
	var (
		tracer = "repo.GetUserById"
		q      = `
		select id, name, balance
		from users
		where id = $1`
	)

	err = r.db.QueryRowContext(ctx, q, id).Scan(&u.Id, &u.Name, &u.Balance)
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

func (r *repo) CreateUser(ctx context.Context, name string, balance int) (err error) {
	var (
		tracer = "repo.CreateUser"
		q      = `
		INSERT INTO users (name, balance) values ($1, $2)`
	)

	_, err = r.db.ExecContext(ctx, q, name, balance)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}
