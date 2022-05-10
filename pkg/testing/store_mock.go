// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation and Dapr Contributors.
// Licensed under the MIT License.
// ------------------------------------------------------------

// Code generated by mockery v2.9.4. DO NOT EDIT.

package testing

import (
	context "context"

	configuration "github.com/dapr/components-contrib/configuration"

	mock "github.com/stretchr/testify/mock"
)

// MockConfigurationStore is an autogenerated mock type for the Store type
type MockConfigurationStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, req
func (_m *MockConfigurationStore) Get(ctx context.Context, req *configuration.GetRequest) (*configuration.GetResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *configuration.GetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *configuration.GetRequest) *configuration.GetResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*configuration.GetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *configuration.GetRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: metadata
func (_m *MockConfigurationStore) Init(metadata configuration.Metadata) error {
	ret := _m.Called(metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(configuration.Metadata) error); ok {
		r0 = rf(metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: ctx, req, handler
func (_m *MockConfigurationStore) Subscribe(ctx context.Context, req *configuration.SubscribeRequest, handler configuration.UpdateHandler) (string, error) {
	ret := _m.Called(ctx, req, handler)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *configuration.SubscribeRequest, configuration.UpdateHandler) string); ok {
		r0 = rf(ctx, req, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *configuration.SubscribeRequest, configuration.UpdateHandler) error); ok {
		r1 = rf(ctx, req, handler)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unsubscribe provides a mock function with given fields: ctx, req
func (_m *MockConfigurationStore) Unsubscribe(ctx context.Context, req *configuration.UnsubscribeRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *configuration.UnsubscribeRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type FailingConfigurationStore struct {
	Failure Failure
}

func (f *FailingConfigurationStore) Get(ctx context.Context, req *configuration.GetRequest) (*configuration.GetResponse, error) {
	if err := f.Failure.PerformFailure(req.Metadata["key"]); err != nil {
		return nil, err
	}
	return &configuration.GetResponse{}, nil
}

func (f *FailingConfigurationStore) Init(metadata configuration.Metadata) error {
	return nil
}

func (f *FailingConfigurationStore) Subscribe(ctx context.Context, req *configuration.SubscribeRequest, handler configuration.UpdateHandler) (string, error) {
	handler(ctx, &configuration.UpdateEvent{
		Items: []*configuration.Item{
			{
				Key:   req.Metadata["key"],
				Value: "testConfig",
			},
		},
	})
	if err := f.Failure.PerformFailure(req.Metadata["key"]); err != nil {
		return "", err
	}
	return "subscribeID", nil
}

func (f *FailingConfigurationStore) Unsubscribe(ctx context.Context, req *configuration.UnsubscribeRequest) error {
	return f.Failure.PerformFailure(req.ID)
}