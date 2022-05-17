package usecase

import "context"

func (uc *usecase) CreateUser(ctx context.Context, name string, balance int) (err error) {
	return uc.r.CreateUser(ctx, name, balance)
}
