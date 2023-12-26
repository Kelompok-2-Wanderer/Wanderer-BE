// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	reports "wanderer/features/reports"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetBookingCurrentYear provides a mock function with given fields: ctx
func (_m *Repository) GetBookingCurrentYear(ctx context.Context) ([]reports.GraphBooking, error) {
	ret := _m.Called(ctx)

	var r0 []reports.GraphBooking
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]reports.GraphBooking, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []reports.GraphBooking); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reports.GraphBooking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecentBooking provides a mock function with given fields: ctx
func (_m *Repository) GetRecentBooking(ctx context.Context) ([]reports.Booking, error) {
	ret := _m.Called(ctx)

	var r0 []reports.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]reports.Booking, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []reports.Booking); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reports.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopTour provides a mock function with given fields: ctx
func (_m *Repository) GetTopTour(ctx context.Context) ([]reports.Tour, error) {
	ret := _m.Called(ctx)

	var r0 []reports.Tour
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]reports.Tour, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []reports.Tour); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reports.Tour)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalBooking provides a mock function with given fields: ctx
func (_m *Repository) GetTotalBooking(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalLocation provides a mock function with given fields: ctx
func (_m *Repository) GetTotalLocation(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalTour provides a mock function with given fields: ctx
func (_m *Repository) GetTotalTour(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUser provides a mock function with given fields: ctx
func (_m *Repository) GetTotalUser(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}