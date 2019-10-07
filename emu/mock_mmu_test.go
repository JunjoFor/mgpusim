// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/mem/vm/mmu (interfaces: MMU)

package emu

import (
	gomock "github.com/golang/mock/gomock"
	vm "gitlab.com/akita/mem/vm"
	ca "gitlab.com/akita/util/ca"
	reflect "reflect"
)

// MockMMU is a mock of MMU interface
type MockMMU struct {
	ctrl     *gomock.Controller
	recorder *MockMMUMockRecorder
}

// MockMMUMockRecorder is the mock recorder for MockMMU
type MockMMUMockRecorder struct {
	mock *MockMMU
}

// NewMockMMU creates a new mock instance
func NewMockMMU(ctrl *gomock.Controller) *MockMMU {
	mock := &MockMMU{ctrl: ctrl}
	mock.recorder = &MockMMUMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMMU) EXPECT() *MockMMUMockRecorder {
	return m.recorder
}

// CreatePage mocks base method
func (m *MockMMU) CreatePage(arg0 *vm.Page) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreatePage", arg0)
}

// CreatePage indicates an expected call of CreatePage
func (mr *MockMMUMockRecorder) CreatePage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePage", reflect.TypeOf((*MockMMU)(nil).CreatePage), arg0)
}

// GetNumPages mocks base method
func (m *MockMMU) GetNumPages(arg0 ca.PID) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumPages", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNumPages indicates an expected call of GetNumPages
func (mr *MockMMUMockRecorder) GetNumPages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumPages", reflect.TypeOf((*MockMMU)(nil).GetNumPages), arg0)
}

// GetOrCreatePageTable mocks base method
func (m *MockMMU) GetOrCreatePageTable(arg0 ca.PID) vm.PageTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreatePageTable", arg0)
	ret0, _ := ret[0].(vm.PageTable)
	return ret0
}

// GetOrCreatePageTable indicates an expected call of GetOrCreatePageTable
func (mr *MockMMUMockRecorder) GetOrCreatePageTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreatePageTable", reflect.TypeOf((*MockMMU)(nil).GetOrCreatePageTable), arg0)
}

// GetPageWithGivenVAddr mocks base method
func (m *MockMMU) GetPageWithGivenVAddr(arg0 uint64, arg1 ca.PID) *vm.Page {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPageWithGivenVAddr", arg0, arg1)
	ret0, _ := ret[0].(*vm.Page)
	return ret0
}

// GetPageWithGivenVAddr indicates an expected call of GetPageWithGivenVAddr
func (mr *MockMMUMockRecorder) GetPageWithGivenVAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPageWithGivenVAddr", reflect.TypeOf((*MockMMU)(nil).GetPageWithGivenVAddr), arg0, arg1)
}

// InvalidatePage mocks base method
func (m *MockMMU) InvalidatePage(arg0 uint64, arg1 ca.PID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InvalidatePage", arg0, arg1)
}

// InvalidatePage indicates an expected call of InvalidatePage
func (mr *MockMMUMockRecorder) InvalidatePage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidatePage", reflect.TypeOf((*MockMMU)(nil).InvalidatePage), arg0, arg1)
}

// MarkPageAsMigrating mocks base method
func (m *MockMMU) MarkPageAsMigrating(arg0 uint64, arg1 ca.PID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkPageAsMigrating", arg0, arg1)
}

// MarkPageAsMigrating indicates an expected call of MarkPageAsMigrating
func (mr *MockMMUMockRecorder) MarkPageAsMigrating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkPageAsMigrating", reflect.TypeOf((*MockMMU)(nil).MarkPageAsMigrating), arg0, arg1)
}

// RemovePage mocks base method
func (m *MockMMU) RemovePage(arg0 ca.PID, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemovePage", arg0, arg1)
}

// RemovePage indicates an expected call of RemovePage
func (mr *MockMMUMockRecorder) RemovePage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePage", reflect.TypeOf((*MockMMU)(nil).RemovePage), arg0, arg1)
}

// Translate mocks base method
func (m *MockMMU) Translate(arg0 ca.PID, arg1 uint64) (uint64, *vm.Page) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(*vm.Page)
	return ret0, ret1
}

// Translate indicates an expected call of Translate
func (mr *MockMMUMockRecorder) Translate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockMMU)(nil).Translate), arg0, arg1)
}

// ValidatePage mocks base method
func (m *MockMMU) ValidatePage(arg0 uint64, arg1 ca.PID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ValidatePage", arg0, arg1)
}

// ValidatePage indicates an expected call of ValidatePage
func (mr *MockMMUMockRecorder) ValidatePage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePage", reflect.TypeOf((*MockMMU)(nil).ValidatePage), arg0, arg1)
}
