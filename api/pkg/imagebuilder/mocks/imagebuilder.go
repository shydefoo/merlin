// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mlp "github.com/caraml-dev/merlin/mlp"

	mock "github.com/stretchr/testify/mock"

	models "github.com/caraml-dev/merlin/models"
)

// ImageBuilder is an autogenerated mock type for the ImageBuilder type
type ImageBuilder struct {
	mock.Mock
}

// BuildImage provides a mock function with given fields: ctx, project, model, version
func (_m *ImageBuilder) BuildImage(ctx context.Context, project mlp.Project, model *models.Model, version *models.Version) (string, error) {
	ret := _m.Called(ctx, project, model, version)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, mlp.Project, *models.Model, *models.Version) string); ok {
		r0 = rf(ctx, project, model, version)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mlp.Project, *models.Model, *models.Version) error); ok {
		r1 = rf(ctx, project, model, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainers provides a mock function with given fields: ctx, project, model, version
func (_m *ImageBuilder) GetContainers(ctx context.Context, project mlp.Project, model *models.Model, version *models.Version) ([]*models.Container, error) {
	ret := _m.Called(ctx, project, model, version)

	var r0 []*models.Container
	if rf, ok := ret.Get(0).(func(context.Context, mlp.Project, *models.Model, *models.Version) []*models.Container); ok {
		r0 = rf(ctx, project, model, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mlp.Project, *models.Model, *models.Version) error); ok {
		r1 = rf(ctx, project, model, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMainAppPath provides a mock function with given fields: version
func (_m *ImageBuilder) GetMainAppPath(version *models.Version) (string, error) {
	ret := _m.Called(version)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.Version) string); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Version) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewImageBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewImageBuilder creates a new instance of ImageBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImageBuilder(t mockConstructorTestingTNewImageBuilder) *ImageBuilder {
	mock := &ImageBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
