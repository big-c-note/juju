// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state/migrations (interfaces: MigrationRelationNetworks,RelationNetworksSource,RelationNetworksModel)

// Package migrations is a generated GoMock package.
package migrations

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	description "github.com/juju/description/v3"
)

// MockMigrationRelationNetworks is a mock of MigrationRelationNetworks interface
type MockMigrationRelationNetworks struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationRelationNetworksMockRecorder
}

// MockMigrationRelationNetworksMockRecorder is the mock recorder for MockMigrationRelationNetworks
type MockMigrationRelationNetworksMockRecorder struct {
	mock *MockMigrationRelationNetworks
}

// NewMockMigrationRelationNetworks creates a new mock instance
func NewMockMigrationRelationNetworks(ctrl *gomock.Controller) *MockMigrationRelationNetworks {
	mock := &MockMigrationRelationNetworks{ctrl: ctrl}
	mock.recorder = &MockMigrationRelationNetworksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMigrationRelationNetworks) EXPECT() *MockMigrationRelationNetworksMockRecorder {
	return m.recorder
}

// CIDRS mocks base method
func (m *MockMigrationRelationNetworks) CIDRS() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CIDRS")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CIDRS indicates an expected call of CIDRS
func (mr *MockMigrationRelationNetworksMockRecorder) CIDRS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CIDRS", reflect.TypeOf((*MockMigrationRelationNetworks)(nil).CIDRS))
}

// Id mocks base method
func (m *MockMigrationRelationNetworks) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockMigrationRelationNetworksMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockMigrationRelationNetworks)(nil).Id))
}

// RelationKey mocks base method
func (m *MockMigrationRelationNetworks) RelationKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// RelationKey indicates an expected call of RelationKey
func (mr *MockMigrationRelationNetworksMockRecorder) RelationKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationKey", reflect.TypeOf((*MockMigrationRelationNetworks)(nil).RelationKey))
}

// MockRelationNetworksSource is a mock of RelationNetworksSource interface
type MockRelationNetworksSource struct {
	ctrl     *gomock.Controller
	recorder *MockRelationNetworksSourceMockRecorder
}

// MockRelationNetworksSourceMockRecorder is the mock recorder for MockRelationNetworksSource
type MockRelationNetworksSourceMockRecorder struct {
	mock *MockRelationNetworksSource
}

// NewMockRelationNetworksSource creates a new mock instance
func NewMockRelationNetworksSource(ctrl *gomock.Controller) *MockRelationNetworksSource {
	mock := &MockRelationNetworksSource{ctrl: ctrl}
	mock.recorder = &MockRelationNetworksSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelationNetworksSource) EXPECT() *MockRelationNetworksSourceMockRecorder {
	return m.recorder
}

// AllRelationNetworks mocks base method
func (m *MockRelationNetworksSource) AllRelationNetworks() ([]MigrationRelationNetworks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRelationNetworks")
	ret0, _ := ret[0].([]MigrationRelationNetworks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRelationNetworks indicates an expected call of AllRelationNetworks
func (mr *MockRelationNetworksSourceMockRecorder) AllRelationNetworks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRelationNetworks", reflect.TypeOf((*MockRelationNetworksSource)(nil).AllRelationNetworks))
}

// MockRelationNetworksModel is a mock of RelationNetworksModel interface
type MockRelationNetworksModel struct {
	ctrl     *gomock.Controller
	recorder *MockRelationNetworksModelMockRecorder
}

// MockRelationNetworksModelMockRecorder is the mock recorder for MockRelationNetworksModel
type MockRelationNetworksModelMockRecorder struct {
	mock *MockRelationNetworksModel
}

// NewMockRelationNetworksModel creates a new mock instance
func NewMockRelationNetworksModel(ctrl *gomock.Controller) *MockRelationNetworksModel {
	mock := &MockRelationNetworksModel{ctrl: ctrl}
	mock.recorder = &MockRelationNetworksModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelationNetworksModel) EXPECT() *MockRelationNetworksModelMockRecorder {
	return m.recorder
}

// AddRelationNetwork mocks base method
func (m *MockRelationNetworksModel) AddRelationNetwork(arg0 description.RelationNetworkArgs) description.RelationNetwork {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRelationNetwork", arg0)
	ret0, _ := ret[0].(description.RelationNetwork)
	return ret0
}

// AddRelationNetwork indicates an expected call of AddRelationNetwork
func (mr *MockRelationNetworksModelMockRecorder) AddRelationNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRelationNetwork", reflect.TypeOf((*MockRelationNetworksModel)(nil).AddRelationNetwork), arg0)
}
