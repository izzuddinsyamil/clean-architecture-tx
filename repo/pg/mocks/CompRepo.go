// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"
	model "repo-pattern-w-trx-management/model"

	mock "github.com/stretchr/testify/mock"

	pg_repo "repo-pattern-w-trx-management/repo/pg"
)

// CompRepo is an autogenerated mock type for the CompRepo type
type CompRepo struct {
	mock.Mock
}

// AddBalance provides a mock function with given fields: ctx, userId, amount
func (_m *CompRepo) AddBalance(ctx context.Context, userId int, amount int) error {
	ret := _m.Called(ctx, userId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, userId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Atomic provides a mock function with given fields: ctx, fn
func (_m *CompRepo) Atomic(ctx context.Context, fn func(pg_repo.CompRepo) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(pg_repo.CompRepo) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTransaction provides a mock function with given fields: ctx, senderId, receiverId, amount
func (_m *CompRepo) CreateTransaction(ctx context.Context, senderId int, receiverId int, amount int) error {
	ret := _m.Called(ctx, senderId, receiverId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) error); ok {
		r0 = rf(ctx, senderId, receiverId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, name, balance
func (_m *CompRepo) CreateUser(ctx context.Context, name string, balance int) error {
	ret := _m.Called(ctx, name, balance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, name, balance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeductBalance provides a mock function with given fields: ctx, userId, amount
func (_m *CompRepo) DeductBalance(ctx context.Context, userId int, amount int) error {
	ret := _m.Called(ctx, userId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, userId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx
func (_m *CompRepo) GetUser(ctx context.Context) ([]model.User, error) {
	ret := _m.Called(ctx)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(context.Context) []model.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewCompRepoT interface {
	mock.TestingT
	Cleanup(func())
}

// NewCompRepo creates a new instance of CompRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCompRepo(t NewCompRepoT) *CompRepo {
	mock := &CompRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
