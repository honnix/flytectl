// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"

	filters "github.com/flyteorg/flytectl/pkg/filters"

	mock "github.com/stretchr/testify/mock"

	service "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"
)

// AdminFetcherExtInterface is an autogenerated mock type for the AdminFetcherExtInterface type
type AdminFetcherExtInterface struct {
	mock.Mock
}

type AdminFetcherExtInterface_AdminServiceClient struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_AdminServiceClient) Return(_a0 service.AdminServiceClient) *AdminFetcherExtInterface_AdminServiceClient {
	return &AdminFetcherExtInterface_AdminServiceClient{Call: _m.Call.Return(_a0)}
}

func (_m *AdminFetcherExtInterface) OnAdminServiceClient() *AdminFetcherExtInterface_AdminServiceClient {
	c := _m.On("AdminServiceClient")
	return &AdminFetcherExtInterface_AdminServiceClient{Call: c}
}

func (_m *AdminFetcherExtInterface) OnAdminServiceClientMatch(matchers ...interface{}) *AdminFetcherExtInterface_AdminServiceClient {
	c := _m.On("AdminServiceClient", matchers...)
	return &AdminFetcherExtInterface_AdminServiceClient{Call: c}
}

// AdminServiceClient provides a mock function with given fields:
func (_m *AdminFetcherExtInterface) AdminServiceClient() service.AdminServiceClient {
	ret := _m.Called()

	var r0 service.AdminServiceClient
	if rf, ok := ret.Get(0).(func() service.AdminServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(service.AdminServiceClient)
		}
	}

	return r0
}

type AdminFetcherExtInterface_FetchAllVerOfLP struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchAllVerOfLP) Return(_a0 []*admin.LaunchPlan, _a1 error) *AdminFetcherExtInterface_FetchAllVerOfLP {
	return &AdminFetcherExtInterface_FetchAllVerOfLP{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfLP(ctx context.Context, lpName string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchAllVerOfLP {
	c := _m.On("FetchAllVerOfLP", ctx, lpName, project, domain, filter)
	return &AdminFetcherExtInterface_FetchAllVerOfLP{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfLPMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchAllVerOfLP {
	c := _m.On("FetchAllVerOfLP", matchers...)
	return &AdminFetcherExtInterface_FetchAllVerOfLP{Call: c}
}

// FetchAllVerOfLP provides a mock function with given fields: ctx, lpName, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchAllVerOfLP(ctx context.Context, lpName string, project string, domain string, filter filters.Filters) ([]*admin.LaunchPlan, error) {
	ret := _m.Called(ctx, lpName, project, domain, filter)

	var r0 []*admin.LaunchPlan
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) []*admin.LaunchPlan); ok {
		r0 = rf(ctx, lpName, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.LaunchPlan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, lpName, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchAllVerOfTask struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchAllVerOfTask) Return(_a0 []*admin.Task, _a1 error) *AdminFetcherExtInterface_FetchAllVerOfTask {
	return &AdminFetcherExtInterface_FetchAllVerOfTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfTask(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchAllVerOfTask {
	c := _m.On("FetchAllVerOfTask", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchAllVerOfTask{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfTaskMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchAllVerOfTask {
	c := _m.On("FetchAllVerOfTask", matchers...)
	return &AdminFetcherExtInterface_FetchAllVerOfTask{Call: c}
}

// FetchAllVerOfTask provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchAllVerOfTask(ctx context.Context, name string, project string, domain string, filter filters.Filters) ([]*admin.Task, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 []*admin.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) []*admin.Task); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchAllVerOfWorkflow struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchAllVerOfWorkflow) Return(_a0 []*admin.Workflow, _a1 error) *AdminFetcherExtInterface_FetchAllVerOfWorkflow {
	return &AdminFetcherExtInterface_FetchAllVerOfWorkflow{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfWorkflow(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchAllVerOfWorkflow {
	c := _m.On("FetchAllVerOfWorkflow", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchAllVerOfWorkflow{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfWorkflowMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchAllVerOfWorkflow {
	c := _m.On("FetchAllVerOfWorkflow", matchers...)
	return &AdminFetcherExtInterface_FetchAllVerOfWorkflow{Call: c}
}

// FetchAllVerOfWorkflow provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchAllVerOfWorkflow(ctx context.Context, name string, project string, domain string, filter filters.Filters) ([]*admin.Workflow, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 []*admin.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) []*admin.Workflow); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchExecution struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchExecution) Return(_a0 *admin.Execution, _a1 error) *AdminFetcherExtInterface_FetchExecution {
	return &AdminFetcherExtInterface_FetchExecution{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchExecution(ctx context.Context, name string, project string, domain string) *AdminFetcherExtInterface_FetchExecution {
	c := _m.On("FetchExecution", ctx, name, project, domain)
	return &AdminFetcherExtInterface_FetchExecution{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchExecutionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchExecution {
	c := _m.On("FetchExecution", matchers...)
	return &AdminFetcherExtInterface_FetchExecution{Call: c}
}

// FetchExecution provides a mock function with given fields: ctx, name, project, domain
func (_m *AdminFetcherExtInterface) FetchExecution(ctx context.Context, name string, project string, domain string) (*admin.Execution, error) {
	ret := _m.Called(ctx, name, project, domain)

	var r0 *admin.Execution
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *admin.Execution); ok {
		r0 = rf(ctx, name, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchLPLatestVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchLPLatestVersion) Return(_a0 *admin.LaunchPlan, _a1 error) *AdminFetcherExtInterface_FetchLPLatestVersion {
	return &AdminFetcherExtInterface_FetchLPLatestVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchLPLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchLPLatestVersion {
	c := _m.On("FetchLPLatestVersion", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchLPLatestVersion{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchLPLatestVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchLPLatestVersion {
	c := _m.On("FetchLPLatestVersion", matchers...)
	return &AdminFetcherExtInterface_FetchLPLatestVersion{Call: c}
}

// FetchLPLatestVersion provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchLPLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) (*admin.LaunchPlan, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 *admin.LaunchPlan
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) *admin.LaunchPlan); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.LaunchPlan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchLPVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchLPVersion) Return(_a0 *admin.LaunchPlan, _a1 error) *AdminFetcherExtInterface_FetchLPVersion {
	return &AdminFetcherExtInterface_FetchLPVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchLPVersion(ctx context.Context, name string, version string, project string, domain string) *AdminFetcherExtInterface_FetchLPVersion {
	c := _m.On("FetchLPVersion", ctx, name, version, project, domain)
	return &AdminFetcherExtInterface_FetchLPVersion{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchLPVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchLPVersion {
	c := _m.On("FetchLPVersion", matchers...)
	return &AdminFetcherExtInterface_FetchLPVersion{Call: c}
}

// FetchLPVersion provides a mock function with given fields: ctx, name, version, project, domain
func (_m *AdminFetcherExtInterface) FetchLPVersion(ctx context.Context, name string, version string, project string, domain string) (*admin.LaunchPlan, error) {
	ret := _m.Called(ctx, name, version, project, domain)

	var r0 *admin.LaunchPlan
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.LaunchPlan); ok {
		r0 = rf(ctx, name, version, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.LaunchPlan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, version, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchProjectDomainAttributes struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchProjectDomainAttributes) Return(_a0 *admin.ProjectDomainAttributesGetResponse, _a1 error) *AdminFetcherExtInterface_FetchProjectDomainAttributes {
	return &AdminFetcherExtInterface_FetchProjectDomainAttributes{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchProjectDomainAttributes(ctx context.Context, project string, domain string, rsType admin.MatchableResource) *AdminFetcherExtInterface_FetchProjectDomainAttributes {
	c := _m.On("FetchProjectDomainAttributes", ctx, project, domain, rsType)
	return &AdminFetcherExtInterface_FetchProjectDomainAttributes{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchProjectDomainAttributesMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchProjectDomainAttributes {
	c := _m.On("FetchProjectDomainAttributes", matchers...)
	return &AdminFetcherExtInterface_FetchProjectDomainAttributes{Call: c}
}

// FetchProjectDomainAttributes provides a mock function with given fields: ctx, project, domain, rsType
func (_m *AdminFetcherExtInterface) FetchProjectDomainAttributes(ctx context.Context, project string, domain string, rsType admin.MatchableResource) (*admin.ProjectDomainAttributesGetResponse, error) {
	ret := _m.Called(ctx, project, domain, rsType)

	var r0 *admin.ProjectDomainAttributesGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, admin.MatchableResource) *admin.ProjectDomainAttributesGetResponse); ok {
		r0 = rf(ctx, project, domain, rsType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ProjectDomainAttributesGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, admin.MatchableResource) error); ok {
		r1 = rf(ctx, project, domain, rsType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchTaskLatestVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchTaskLatestVersion) Return(_a0 *admin.Task, _a1 error) *AdminFetcherExtInterface_FetchTaskLatestVersion {
	return &AdminFetcherExtInterface_FetchTaskLatestVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchTaskLatestVersion {
	c := _m.On("FetchTaskLatestVersion", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchTaskLatestVersion{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskLatestVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchTaskLatestVersion {
	c := _m.On("FetchTaskLatestVersion", matchers...)
	return &AdminFetcherExtInterface_FetchTaskLatestVersion{Call: c}
}

// FetchTaskLatestVersion provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchTaskLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) (*admin.Task, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 *admin.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) *admin.Task); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchTaskVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchTaskVersion) Return(_a0 *admin.Task, _a1 error) *AdminFetcherExtInterface_FetchTaskVersion {
	return &AdminFetcherExtInterface_FetchTaskVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskVersion(ctx context.Context, name string, version string, project string, domain string) *AdminFetcherExtInterface_FetchTaskVersion {
	c := _m.On("FetchTaskVersion", ctx, name, version, project, domain)
	return &AdminFetcherExtInterface_FetchTaskVersion{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchTaskVersion {
	c := _m.On("FetchTaskVersion", matchers...)
	return &AdminFetcherExtInterface_FetchTaskVersion{Call: c}
}

// FetchTaskVersion provides a mock function with given fields: ctx, name, version, project, domain
func (_m *AdminFetcherExtInterface) FetchTaskVersion(ctx context.Context, name string, version string, project string, domain string) (*admin.Task, error) {
	ret := _m.Called(ctx, name, version, project, domain)

	var r0 *admin.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.Task); ok {
		r0 = rf(ctx, name, version, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, version, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchWorkflowAttributes struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchWorkflowAttributes) Return(_a0 *admin.WorkflowAttributesGetResponse, _a1 error) *AdminFetcherExtInterface_FetchWorkflowAttributes {
	return &AdminFetcherExtInterface_FetchWorkflowAttributes{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowAttributes(ctx context.Context, project string, domain string, name string, rsType admin.MatchableResource) *AdminFetcherExtInterface_FetchWorkflowAttributes {
	c := _m.On("FetchWorkflowAttributes", ctx, project, domain, name, rsType)
	return &AdminFetcherExtInterface_FetchWorkflowAttributes{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowAttributesMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchWorkflowAttributes {
	c := _m.On("FetchWorkflowAttributes", matchers...)
	return &AdminFetcherExtInterface_FetchWorkflowAttributes{Call: c}
}

// FetchWorkflowAttributes provides a mock function with given fields: ctx, project, domain, name, rsType
func (_m *AdminFetcherExtInterface) FetchWorkflowAttributes(ctx context.Context, project string, domain string, name string, rsType admin.MatchableResource) (*admin.WorkflowAttributesGetResponse, error) {
	ret := _m.Called(ctx, project, domain, name, rsType)

	var r0 *admin.WorkflowAttributesGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, admin.MatchableResource) *admin.WorkflowAttributesGetResponse); ok {
		r0 = rf(ctx, project, domain, name, rsType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.WorkflowAttributesGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, admin.MatchableResource) error); ok {
		r1 = rf(ctx, project, domain, name, rsType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchWorkflowLatestVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchWorkflowLatestVersion) Return(_a0 *admin.Workflow, _a1 error) *AdminFetcherExtInterface_FetchWorkflowLatestVersion {
	return &AdminFetcherExtInterface_FetchWorkflowLatestVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchWorkflowLatestVersion {
	c := _m.On("FetchWorkflowLatestVersion", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchWorkflowLatestVersion{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowLatestVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchWorkflowLatestVersion {
	c := _m.On("FetchWorkflowLatestVersion", matchers...)
	return &AdminFetcherExtInterface_FetchWorkflowLatestVersion{Call: c}
}

// FetchWorkflowLatestVersion provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchWorkflowLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) (*admin.Workflow, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 *admin.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) *admin.Workflow); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchWorkflowVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchWorkflowVersion) Return(_a0 *admin.Workflow, _a1 error) *AdminFetcherExtInterface_FetchWorkflowVersion {
	return &AdminFetcherExtInterface_FetchWorkflowVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowVersion(ctx context.Context, name string, version string, project string, domain string) *AdminFetcherExtInterface_FetchWorkflowVersion {
	c := _m.On("FetchWorkflowVersion", ctx, name, version, project, domain)
	return &AdminFetcherExtInterface_FetchWorkflowVersion{Call: c}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchWorkflowVersion {
	c := _m.On("FetchWorkflowVersion", matchers...)
	return &AdminFetcherExtInterface_FetchWorkflowVersion{Call: c}
}

// FetchWorkflowVersion provides a mock function with given fields: ctx, name, version, project, domain
func (_m *AdminFetcherExtInterface) FetchWorkflowVersion(ctx context.Context, name string, version string, project string, domain string) (*admin.Workflow, error) {
	ret := _m.Called(ctx, name, version, project, domain)

	var r0 *admin.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.Workflow); ok {
		r0 = rf(ctx, name, version, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, version, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_ListExecution struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_ListExecution) Return(_a0 *admin.ExecutionList, _a1 error) *AdminFetcherExtInterface_ListExecution {
	return &AdminFetcherExtInterface_ListExecution{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnListExecution(ctx context.Context, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_ListExecution {
	c := _m.On("ListExecution", ctx, project, domain, filter)
	return &AdminFetcherExtInterface_ListExecution{Call: c}
}

func (_m *AdminFetcherExtInterface) OnListExecutionMatch(matchers ...interface{}) *AdminFetcherExtInterface_ListExecution {
	c := _m.On("ListExecution", matchers...)
	return &AdminFetcherExtInterface_ListExecution{Call: c}
}

// ListExecution provides a mock function with given fields: ctx, project, domain, filter
func (_m *AdminFetcherExtInterface) ListExecution(ctx context.Context, project string, domain string, filter filters.Filters) (*admin.ExecutionList, error) {
	ret := _m.Called(ctx, project, domain, filter)

	var r0 *admin.ExecutionList
	if rf, ok := ret.Get(0).(func(context.Context, string, string, filters.Filters) *admin.ExecutionList); ok {
		r0 = rf(ctx, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ExecutionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_ListProjects struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_ListProjects) Return(_a0 *admin.Projects, _a1 error) *AdminFetcherExtInterface_ListProjects {
	return &AdminFetcherExtInterface_ListProjects{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnListProjects(ctx context.Context, filter filters.Filters) *AdminFetcherExtInterface_ListProjects {
	c := _m.On("ListProjects", ctx, filter)
	return &AdminFetcherExtInterface_ListProjects{Call: c}
}

func (_m *AdminFetcherExtInterface) OnListProjectsMatch(matchers ...interface{}) *AdminFetcherExtInterface_ListProjects {
	c := _m.On("ListProjects", matchers...)
	return &AdminFetcherExtInterface_ListProjects{Call: c}
}

// ListProjects provides a mock function with given fields: ctx, filter
func (_m *AdminFetcherExtInterface) ListProjects(ctx context.Context, filter filters.Filters) (*admin.Projects, error) {
	ret := _m.Called(ctx, filter)

	var r0 *admin.Projects
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filters) *admin.Projects); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Projects)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, filters.Filters) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
