// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeflow/katib/manager/modelstore (interfaces: ModelStore)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/kubeflow/katib/api"
	reflect "reflect"
)

// MockModelStore is a mock of ModelStore interface
type MockModelStore struct {
	ctrl     *gomock.Controller
	recorder *MockModelStoreMockRecorder
}

// MockModelStoreMockRecorder is the mock recorder for MockModelStore
type MockModelStoreMockRecorder struct {
	mock *MockModelStore
}

// NewMockModelStore creates a new mock instance
func NewMockModelStore(ctrl *gomock.Controller) *MockModelStore {
	mock := &MockModelStore{ctrl: ctrl}
	mock.recorder = &MockModelStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelStore) EXPECT() *MockModelStoreMockRecorder {
	return m.recorder
}

// GetSavedModel mocks base method
func (m *MockModelStore) GetSavedModel(arg0 *api.GetSavedModelRequest) (*api.ModelInfo, error) {
	ret := m.ctrl.Call(m, "GetSavedModel", arg0)
	ret0, _ := ret[0].(*api.ModelInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavedModel indicates an expected call of GetSavedModel
func (mr *MockModelStoreMockRecorder) GetSavedModel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavedModel", reflect.TypeOf((*MockModelStore)(nil).GetSavedModel), arg0)
}

// GetSavedModels mocks base method
func (m *MockModelStore) GetSavedModels(arg0 *api.GetSavedModelsRequest) ([]*api.ModelInfo, error) {
	ret := m.ctrl.Call(m, "GetSavedModels", arg0)
	ret0, _ := ret[0].([]*api.ModelInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavedModels indicates an expected call of GetSavedModels
func (mr *MockModelStoreMockRecorder) GetSavedModels(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavedModels", reflect.TypeOf((*MockModelStore)(nil).GetSavedModels), arg0)
}

// GetSavedStudies mocks base method
func (m *MockModelStore) GetSavedStudies() ([]*api.StudyOverview, error) {
	ret := m.ctrl.Call(m, "GetSavedStudies")
	ret0, _ := ret[0].([]*api.StudyOverview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavedStudies indicates an expected call of GetSavedStudies
func (mr *MockModelStoreMockRecorder) GetSavedStudies() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavedStudies", reflect.TypeOf((*MockModelStore)(nil).GetSavedStudies))
}

// SaveModel mocks base method
func (m *MockModelStore) SaveModel(arg0 *api.SaveModelRequest) error {
	ret := m.ctrl.Call(m, "SaveModel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveModel indicates an expected call of SaveModel
func (mr *MockModelStoreMockRecorder) SaveModel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveModel", reflect.TypeOf((*MockModelStore)(nil).SaveModel), arg0)
}

// SaveStudy mocks base method
func (m *MockModelStore) SaveStudy(arg0 *api.SaveStudyRequest) error {
	ret := m.ctrl.Call(m, "SaveStudy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveStudy indicates an expected call of SaveStudy
func (mr *MockModelStoreMockRecorder) SaveStudy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveStudy", reflect.TypeOf((*MockModelStore)(nil).SaveStudy), arg0)
}
