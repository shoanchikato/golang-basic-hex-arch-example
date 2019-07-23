// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shoanchikato/demo/hex/toy (interfaces: Repository,Service,Handler)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	toy "github.com/shoanchikato/demo/hex/toy"
	http "net/http"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRepository) Create(arg0 *toy.Toy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0)
}

// FindAll mocks base method
func (m *MockRepository) FindAll() ([]*toy.Toy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*toy.Toy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll))
}

// FindByID mocks base method
func (m *MockRepository) FindByID(arg0 string) (*toy.Toy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*toy.Toy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockRepositoryMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), arg0)
}

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateToy mocks base method
func (m *MockService) CreateToy(arg0 *toy.Toy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateToy indicates an expected call of CreateToy
func (mr *MockServiceMockRecorder) CreateToy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToy", reflect.TypeOf((*MockService)(nil).CreateToy), arg0)
}

// FindAllToys mocks base method
func (m *MockService) FindAllToys() ([]*toy.Toy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllToys")
	ret0, _ := ret[0].([]*toy.Toy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllToys indicates an expected call of FindAllToys
func (mr *MockServiceMockRecorder) FindAllToys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllToys", reflect.TypeOf((*MockService)(nil).FindAllToys))
}

// FindToyByID mocks base method
func (m *MockService) FindToyByID(arg0 string) (*toy.Toy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindToyByID", arg0)
	ret0, _ := ret[0].(*toy.Toy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindToyByID indicates an expected call of FindToyByID
func (mr *MockServiceMockRecorder) FindToyByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindToyByID", reflect.TypeOf((*MockService)(nil).FindToyByID), arg0)
}

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockHandler) Create(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", arg0, arg1)
}

// Create indicates an expected call of Create
func (mr *MockHandlerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockHandler)(nil).Create), arg0, arg1)
}

// Get mocks base method
func (m *MockHandler) Get(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", arg0, arg1)
}

// Get indicates an expected call of Get
func (mr *MockHandlerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHandler)(nil).Get), arg0, arg1)
}

// GetByID mocks base method
func (m *MockHandler) GetByID(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetByID", arg0, arg1)
}

// GetByID indicates an expected call of GetByID
func (mr *MockHandlerMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockHandler)(nil).GetByID), arg0, arg1)
}