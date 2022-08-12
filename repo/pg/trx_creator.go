package pg_repo

import (
	"context"
	"database/sql"
	errCode "repo-pattern-w-trx-management/error"

	"github.com/palantir/stacktrace"
)

type TxCreator interface {
	CreateTransaction(ctx context.Context, senderId, receiverId, amount int) (err error)
}

type txCreatorRepo struct {
	conn *sql.DB
	db   DBTX
}

func NewTxCreatorRepo(db *sql.DB) TxCreator {
	return &txCreatorRepo{
		conn: db,
		db:   db,
	}
}

func (tc *txCreatorRepo) CreateTransaction(ctx context.Context, senderId, receiverId, amount int) (err error) {

	var (
		tracer = "repo.AddBalance"
		q      = `
		insert into transactions (amount, receiver_id, sender_id) values ($1, $2, $3)`
	)

	_, err = tc.db.ExecContext(ctx, q, amount, receiverId, senderId)
	if err != nil {
		err = stacktrace.PropagateWithCode(err, errCode.EcodeInternal, tracer)
		return
	}

	return
}
