package usecase

import (
	"context"
	"repo-pattern-w-trx-management/model"
)

func (uc *usecase) GetUserList(ctx context.Context) (u []model.User, err error) {
	return uc.r.GetUser(ctx)
}
