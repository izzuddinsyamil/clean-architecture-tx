package usecase

import (
	"context"
	"errors"
	repository "repo-pattern-w-trx-management/repo"
	"repo-pattern-w-trx-management/repo/mocks"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestTransact(t *testing.T) {
	ctx := context.Background()
	tests := map[string]struct {
		senderId, receiverId, amount int
		errWant                      error
		repo                         repository.Repository
	}{
		"Happy case": {
			senderId:   1,
			receiverId: 2,
			amount:     1000,
			errWant:    nil,
			repo: func() repository.Repository {
				repoMock := &mocks.Repository{}
				repoMock.On("Atomic", ctx, mock.AnythingOfType("func(repo.Repository) error")).
					Return(nil).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(func(repository.Repository) error)

						rMock := &mocks.Repository{}
						rMock.On("CreateTransaction", ctx, 1, 2, 1000).
							Return(nil).Once()
						rMock.On("DeductBalance", ctx, 1, 1000).
							Return(nil).Once()
						rMock.On("AddBalance", ctx, 2, 1000).
							Return(nil).Once()

						arg(rMock)
					})

				return repoMock
			}(),
		},
		"Error case: got error when creating transaction": {
			senderId:   1,
			receiverId: 2,
			amount:     1000,
			errWant:    nil,
			repo: func() repository.Repository {
				repoMock := &mocks.Repository{}
				repoMock.On("Atomic", ctx, mock.AnythingOfType("func(repo.Repository) error")).
					Return(nil).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(func(repository.Repository) error)

						rMock := &mocks.Repository{}
						rMock.On("CreateTransaction", ctx, 1, 2, 1000).
							Return(errors.New("some db error")).Once()
						arg(rMock)
					})

				return repoMock
			}(),
		},
		"Error case: got error when deducting balance": {
			senderId:   1,
			receiverId: 2,
			amount:     1000,
			errWant:    nil,
			repo: func() repository.Repository {
				repoMock := &mocks.Repository{}
				repoMock.On("Atomic", ctx, mock.AnythingOfType("func(repo.Repository) error")).
					Return(nil).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(func(repository.Repository) error)

						rMock := &mocks.Repository{}
						rMock.On("CreateTransaction", ctx, 1, 2, 1000).
							Return(nil).Once()
						rMock.On("DeductBalance", ctx, 1, 1000).
							Return(errors.New("some db error")).Once()
						arg(rMock)
					})

				return repoMock
			}(),
		},
	}

	for tName, tCase := range tests {
		t.Run(tName, func(t *testing.T) {
			uc := usecase{
				r: tCase.repo,
			}

			errGot := uc.Transact(ctx, tCase.senderId, tCase.receiverId, tCase.amount)
			assert.Equal(t, tCase.errWant, errGot)
		})
	}
}
