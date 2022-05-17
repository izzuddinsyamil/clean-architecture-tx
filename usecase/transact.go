package usecase

import (
	"context"
	repository "repo-pattern-w-trx-management/repo"
)

func (uc *usecase) Transact(ctx context.Context, senderId, receiverId, amount int) (err error) {
	err = uc.ar.Atomic(ctx, func(ar repository.AtomicRepository) error {
		err := ar.DeductBalance(ctx, senderId, amount)
		if err != nil {
			return err
		}

		err = ar.AddBalance(ctx, receiverId, amount)
		if err != nil {
			return err
		}

		return nil
	})
	return
}
