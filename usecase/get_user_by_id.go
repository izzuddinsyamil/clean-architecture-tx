package usecase

import (
	"context"
	"repo-pattern-w-trx-management/model"
)

func (uc *usecase) GetUserById(ctx context.Context, id int) (model.User, error) {
	return uc.r.GetUserById(ctx, id)
}
