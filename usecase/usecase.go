package usecase

import (
	repository "repo-pattern-w-trx-management/repo/pg"
)

type usecase struct {
	r repository.Repository
}

func NewUsecase(r repository.Repository) *usecase {
	return &usecase{
		r: r,
	}
}
