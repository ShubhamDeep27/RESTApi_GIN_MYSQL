// Code generated by MockGen. DO NOT EDIT.
// Source: rest/gin/dao (interfaces: EmployeeDao)

// Package mocks is a generated GoMock package.
package service

import (
	reflect "reflect"
	models "rest/gin/models"

	gomock "github.com/golang/mock/gomock"
)

// MockEmployeeDao is a mock of EmployeeDao interface.
type MockEmployeeDao struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeDaoMockRecorder
}

// MockEmployeeDaoMockRecorder is the mock recorder for MockEmployeeDao.
type MockEmployeeDaoMockRecorder struct {
	mock *MockEmployeeDao
}

// NewMockEmployeeDao creates a new mock instance.
func NewMockEmployeeDao(ctrl *gomock.Controller) *MockEmployeeDao {
	mock := &MockEmployeeDao{ctrl: ctrl}
	mock.recorder = &MockEmployeeDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeDao) EXPECT() *MockEmployeeDaoMockRecorder {
	return m.recorder
}

// CreateEmployees mocks base method.
func (m *MockEmployeeDao) CreateEmployees(arg0 *models.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployees", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEmployees indicates an expected call of CreateEmployees.
func (mr *MockEmployeeDaoMockRecorder) CreateEmployees(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployees", reflect.TypeOf((*MockEmployeeDao)(nil).CreateEmployees), arg0)
}

// GetAllEmployees mocks base method.
func (m *MockEmployeeDao) GetAllEmployees() ([]*models.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEmployees")
	ret0, _ := ret[0].([]*models.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEmployees indicates an expected call of GetAllEmployees.
func (mr *MockEmployeeDaoMockRecorder) GetAllEmployees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEmployees", reflect.TypeOf((*MockEmployeeDao)(nil).GetAllEmployees))
}

// GetEmployeeById mocks base method.
func (m *MockEmployeeDao) GetEmployeeById(arg0 string) (*models.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeById", arg0)
	ret0, _ := ret[0].(*models.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeById indicates an expected call of GetEmployeeById.
func (mr *MockEmployeeDaoMockRecorder) GetEmployeeById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeById", reflect.TypeOf((*MockEmployeeDao)(nil).GetEmployeeById), arg0)
}

// UpdateEmployee mocks base method.
func (m *MockEmployeeDao) UpdateEmployee(arg0 *models.Employee, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockEmployeeDaoMockRecorder) UpdateEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockEmployeeDao)(nil).UpdateEmployee), arg0, arg1)
}
