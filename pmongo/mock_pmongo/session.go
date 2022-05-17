// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/easyops-cn/mongo-driver-helper/pmongo (interfaces: SessionInterface)

// Package mock_pmongo is a generated GoMock package.
package mock_pmongo

import (
	context "context"
	reflect "reflect"

	pmongo "github.com/easyops-cn/mongo-driver-helper/pmongo"
	gomock "github.com/golang/mock/gomock"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockSessionInterface is a mock of SessionInterface interface.
type MockSessionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSessionInterfaceMockRecorder
}

// MockSessionInterfaceMockRecorder is the mock recorder for MockSessionInterface.
type MockSessionInterfaceMockRecorder struct {
	mock *MockSessionInterface
}

// NewMockSessionInterface creates a new mock instance.
func NewMockSessionInterface(ctrl *gomock.Controller) *MockSessionInterface {
	mock := &MockSessionInterface{ctrl: ctrl}
	mock.recorder = &MockSessionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionInterface) EXPECT() *MockSessionInterfaceMockRecorder {
	return m.recorder
}

// AbortTransaction mocks base method.
func (m *MockSessionInterface) AbortTransaction(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortTransaction indicates an expected call of AbortTransaction.
func (mr *MockSessionInterfaceMockRecorder) AbortTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortTransaction", reflect.TypeOf((*MockSessionInterface)(nil).AbortTransaction), arg0)
}

// AdvanceClusterTime mocks base method.
func (m *MockSessionInterface) AdvanceClusterTime(arg0 bson.Raw) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvanceClusterTime", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdvanceClusterTime indicates an expected call of AdvanceClusterTime.
func (mr *MockSessionInterfaceMockRecorder) AdvanceClusterTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvanceClusterTime", reflect.TypeOf((*MockSessionInterface)(nil).AdvanceClusterTime), arg0)
}

// AdvanceOperationTime mocks base method.
func (m *MockSessionInterface) AdvanceOperationTime(arg0 *primitive.Timestamp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvanceOperationTime", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdvanceOperationTime indicates an expected call of AdvanceOperationTime.
func (mr *MockSessionInterfaceMockRecorder) AdvanceOperationTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvanceOperationTime", reflect.TypeOf((*MockSessionInterface)(nil).AdvanceOperationTime), arg0)
}

// Client mocks base method.
func (m *MockSessionInterface) Client() pmongo.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(pmongo.ClientInterface)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockSessionInterfaceMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockSessionInterface)(nil).Client))
}

// ClusterTime mocks base method.
func (m *MockSessionInterface) ClusterTime() bson.Raw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterTime")
	ret0, _ := ret[0].(bson.Raw)
	return ret0
}

// ClusterTime indicates an expected call of ClusterTime.
func (mr *MockSessionInterfaceMockRecorder) ClusterTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterTime", reflect.TypeOf((*MockSessionInterface)(nil).ClusterTime))
}

// CommitTransaction mocks base method.
func (m *MockSessionInterface) CommitTransaction(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTransaction indicates an expected call of CommitTransaction.
func (mr *MockSessionInterfaceMockRecorder) CommitTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTransaction", reflect.TypeOf((*MockSessionInterface)(nil).CommitTransaction), arg0)
}

// EndSession mocks base method.
func (m *MockSessionInterface) EndSession(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndSession", arg0)
}

// EndSession indicates an expected call of EndSession.
func (mr *MockSessionInterfaceMockRecorder) EndSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndSession", reflect.TypeOf((*MockSessionInterface)(nil).EndSession), arg0)
}

// OperationTime mocks base method.
func (m *MockSessionInterface) OperationTime() *primitive.Timestamp {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperationTime")
	ret0, _ := ret[0].(*primitive.Timestamp)
	return ret0
}

// OperationTime indicates an expected call of OperationTime.
func (mr *MockSessionInterfaceMockRecorder) OperationTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationTime", reflect.TypeOf((*MockSessionInterface)(nil).OperationTime))
}

// StartTransaction mocks base method.
func (m *MockSessionInterface) StartTransaction(arg0 ...*options.TransactionOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartTransaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartTransaction indicates an expected call of StartTransaction.
func (mr *MockSessionInterfaceMockRecorder) StartTransaction(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTransaction", reflect.TypeOf((*MockSessionInterface)(nil).StartTransaction), arg0...)
}

// WithTransaction mocks base method.
func (m *MockSessionInterface) WithTransaction(arg0 context.Context, arg1 func(mongo.SessionContext) (interface{}, error), arg2 ...*options.TransactionOptions) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithTransaction", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockSessionInterfaceMockRecorder) WithTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockSessionInterface)(nil).WithTransaction), varargs...)
}
