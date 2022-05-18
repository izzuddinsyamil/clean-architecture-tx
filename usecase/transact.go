package usecase

import (
	"context"
	repository "repo-pattern-w-trx-management/repo"
)

func (uc *usecase) Transact(ctx context.Context, senderId, receiverId, amount int) (err error) {
	err = uc.r.Atomic(ctx, func(repo repository.Repository) error {
		err := repo.CreateTransaction(ctx, senderId, receiverId, amount)
		if err != nil {
			return err
		}

		err = repo.DeductBalance(ctx, senderId, amount)
		if err != nil {
			return err
		}

		err = repo.AddBalance(ctx, receiverId, amount)
		if err != nil {
			return err
		}

		return nil
	})
	return
}
