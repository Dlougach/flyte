// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"

	mock "github.com/stretchr/testify/mock"

	models "github.com/flyteorg/flyte/flyteadmin/pkg/repositories/models"
)

// NodeExecutionEventRepoInterface is an autogenerated mock type for the NodeExecutionEventRepoInterface type
type NodeExecutionEventRepoInterface struct {
	mock.Mock
}

type NodeExecutionEventRepoInterface_Create struct {
	*mock.Call
}

func (_m NodeExecutionEventRepoInterface_Create) Return(_a0 error) *NodeExecutionEventRepoInterface_Create {
	return &NodeExecutionEventRepoInterface_Create{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionEventRepoInterface) OnCreate(ctx context.Context, id *core.NodeExecutionIdentifier, input models.NodeExecutionEvent) *NodeExecutionEventRepoInterface_Create {
	c_call := _m.On("Create", ctx, id, input)
	return &NodeExecutionEventRepoInterface_Create{Call: c_call}
}

func (_m *NodeExecutionEventRepoInterface) OnCreateMatch(matchers ...interface{}) *NodeExecutionEventRepoInterface_Create {
	c_call := _m.On("Create", matchers...)
	return &NodeExecutionEventRepoInterface_Create{Call: c_call}
}

// Create provides a mock function with given fields: ctx, id, input
func (_m *NodeExecutionEventRepoInterface) Create(ctx context.Context, id *core.NodeExecutionIdentifier, input models.NodeExecutionEvent) error {
	ret := _m.Called(ctx, id, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.NodeExecutionIdentifier, models.NodeExecutionEvent) error); ok {
		r0 = rf(ctx, id, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
