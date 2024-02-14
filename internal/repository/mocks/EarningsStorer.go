// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// EarningsStorer is an autogenerated mock type for the EarningsStorer type
type EarningsStorer struct {
	mock.Mock
}

// GetEarningsByEmpoyeeID provides a mock function with given fields: ID
func (_m *EarningsStorer) GetEarningsByEmpoyeeID(ID string) (repository.Earnings, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for GetEarningsByEmpoyeeID")
	}

	var r0 repository.Earnings
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (repository.Earnings, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(string) repository.Earnings); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(repository.Earnings)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertEarnings provides a mock function with given fields: earnings
func (_m *EarningsStorer) InsertEarnings(earnings repository.Earnings) (repository.Earnings, error) {
	ret := _m.Called(earnings)

	if len(ret) == 0 {
		panic("no return value specified for InsertEarnings")
	}

	var r0 repository.Earnings
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.Earnings) (repository.Earnings, error)); ok {
		return rf(earnings)
	}
	if rf, ok := ret.Get(0).(func(repository.Earnings) repository.Earnings); ok {
		r0 = rf(earnings)
	} else {
		r0 = ret.Get(0).(repository.Earnings)
	}

	if rf, ok := ret.Get(1).(func(repository.Earnings) error); ok {
		r1 = rf(earnings)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEarningsStorer creates a new instance of EarningsStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEarningsStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *EarningsStorer {
	mock := &EarningsStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
