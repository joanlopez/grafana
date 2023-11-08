// Code generated by mockery v2.27.1. DO NOT EDIT.

package ssosettingstests

import (
	context "context"

	models "github.com/grafana/grafana/pkg/services/ssosettings/models"
	mock "github.com/stretchr/testify/mock"
)

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, provider
func (_m *MockStore) Delete(ctx context.Context, provider string) error {
	ret := _m.Called(ctx, provider)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, provider)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, provider
func (_m *MockStore) Get(ctx context.Context, provider string) (*models.SSOSetting, error) {
	ret := _m.Called(ctx, provider)

	var r0 *models.SSOSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.SSOSetting, error)); ok {
		return rf(ctx, provider)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.SSOSetting); ok {
		r0 = rf(ctx, provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SSOSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *MockStore) GetAll(ctx context.Context) ([]*models.SSOSetting, error) {
	ret := _m.Called(ctx)

	var r0 []*models.SSOSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*models.SSOSetting, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*models.SSOSetting); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.SSOSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, provider, data
func (_m *MockStore) Patch(ctx context.Context, provider string, data map[string]interface{}) error {
	ret := _m.Called(ctx, provider, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) error); ok {
		r0 = rf(ctx, provider, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: ctx, provider, data
func (_m *MockStore) Upsert(ctx context.Context, provider string, data map[string]interface{}) error {
	ret := _m.Called(ctx, provider, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) error); ok {
		r0 = rf(ctx, provider, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockStore creates a new instance of MockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStore(t mockConstructorTestingTNewMockStore) *MockStore {
	mock := &MockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}