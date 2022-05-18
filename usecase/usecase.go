package usecase

import (
	repository "repo-pattern-w-trx-management/repo"
)

type usecase struct {
	r repository.Repository
}

func NewUsecase(r repository.Repository) *usecase {
	return &usecase{
		r: r,
	}
}
