// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bo "ic-indexer-service/app/model/bo"
import context "context"
import mock "github.com/stretchr/testify/mock"

import request "ic-indexer-service/app/model/request"

// UpdateHandler is an autogenerated mock type for the UpdateHandler type
type UpdateHandler struct {
	mock.Mock
}

// UpdateIcecreamDetails provides a mock function with given fields: ctx, icecream, params
func (_m *UpdateHandler) UpdateIcecreamDetails(ctx context.Context, icecream *bo.ESIcecream, params request.IcecreamIndexRequest) error {
	ret := _m.Called(ctx, icecream, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *bo.ESIcecream, request.IcecreamIndexRequest) error); ok {
		r0 = rf(ctx, icecream, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
