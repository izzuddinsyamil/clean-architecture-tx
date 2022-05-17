package usecase

import (
	"context"
	"repo-pattern-w-trx-management/model"
)

type repo interface {
	GetUser(ctx context.Context) (u []model.User, err error)
	GetUserById(ctx context.Context, id int) (u model.User, err error)
	CreateUser(ctx context.Context, name string, balance int) (err error)
}

type usecase struct {
	r repo
}

func NewUsecase(r repo) *usecase {
	return &usecase{
		r: r,
	}
}
