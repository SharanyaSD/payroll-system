// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// PayrollStorer is an autogenerated mock type for the PayrollStorer type
type PayrollStorer struct {
	mock.Mock
}

// CreatePayroll provides a mock function with given fields: payroll
func (_m *PayrollStorer) CreatePayroll(payroll repository.Payroll) (repository.Payroll, error) {
	ret := _m.Called(payroll)

	if len(ret) == 0 {
		panic("no return value specified for CreatePayroll")
	}

	var r0 repository.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.Payroll) (repository.Payroll, error)); ok {
		return rf(payroll)
	}
	if rf, ok := ret.Get(0).(func(repository.Payroll) repository.Payroll); ok {
		r0 = rf(payroll)
	} else {
		r0 = ret.Get(0).(repository.Payroll)
	}

	if rf, ok := ret.Get(1).(func(repository.Payroll) error); ok {
		r1 = rf(payroll)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPayroll provides a mock function with given fields:
func (_m *PayrollStorer) GetPayroll() ([]repository.Payroll, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPayroll")
	}

	var r0 []repository.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]repository.Payroll, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []repository.Payroll); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPayrollStorer creates a new instance of PayrollStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPayrollStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *PayrollStorer {
	mock := &PayrollStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
