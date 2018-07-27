// Code generated by MockGen. DO NOT EDIT.
// Source: ./uc/INTERACTOR.go

// Package uc is a generated GoMock package.
package uc

import (
	domain "github.com/err0r500/go-realworld-clean/domain"
	uc "github.com/err0r500/go-realworld-clean/uc"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Log mocks base method
func (m *MockLogger) Log(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log
func (mr *MockLoggerMockRecorder) Log(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogger)(nil).Log), arg0...)
}

// MockAuthHandler is a mock of AuthHandler interface
type MockAuthHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthHandlerMockRecorder
}

// MockAuthHandlerMockRecorder is the mock recorder for MockAuthHandler
type MockAuthHandlerMockRecorder struct {
	mock *MockAuthHandler
}

// NewMockAuthHandler creates a new mock instance
func NewMockAuthHandler(ctrl *gomock.Controller) *MockAuthHandler {
	mock := &MockAuthHandler{ctrl: ctrl}
	mock.recorder = &MockAuthHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthHandler) EXPECT() *MockAuthHandlerMockRecorder {
	return m.recorder
}

// GenUserToken mocks base method
func (m *MockAuthHandler) GenUserToken(userName string) (string, error) {
	ret := m.ctrl.Call(m, "GenUserToken", userName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenUserToken indicates an expected call of GenUserToken
func (mr *MockAuthHandlerMockRecorder) GenUserToken(userName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenUserToken", reflect.TypeOf((*MockAuthHandler)(nil).GenUserToken), userName)
}

// GetUserName mocks base method
func (m *MockAuthHandler) GetUserName(token string) (string, error) {
	ret := m.ctrl.Call(m, "GetUserName", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserName indicates an expected call of GetUserName
func (mr *MockAuthHandlerMockRecorder) GetUserName(token interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserName", reflect.TypeOf((*MockAuthHandler)(nil).GetUserName), token)
}

// MockUserRW is a mock of UserRW interface
type MockUserRW struct {
	ctrl     *gomock.Controller
	recorder *MockUserRWMockRecorder
}

// MockUserRWMockRecorder is the mock recorder for MockUserRW
type MockUserRWMockRecorder struct {
	mock *MockUserRW
}

// NewMockUserRW creates a new mock instance
func NewMockUserRW(ctrl *gomock.Controller) *MockUserRW {
	mock := &MockUserRW{ctrl: ctrl}
	mock.recorder = &MockUserRWMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRW) EXPECT() *MockUserRWMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserRW) Create(username, email, password string) (*domain.User, error) {
	ret := m.ctrl.Call(m, "Create", username, email, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserRWMockRecorder) Create(username, email, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRW)(nil).Create), username, email, password)
}

// GetByName mocks base method
func (m *MockUserRW) GetByName(userName string) (*domain.User, error) {
	ret := m.ctrl.Call(m, "GetByName", userName)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockUserRWMockRecorder) GetByName(userName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockUserRW)(nil).GetByName), userName)
}

// GetByEmailAndPassword mocks base method
func (m *MockUserRW) GetByEmailAndPassword(email, password string) (*domain.User, error) {
	ret := m.ctrl.Call(m, "GetByEmailAndPassword", email, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmailAndPassword indicates an expected call of GetByEmailAndPassword
func (mr *MockUserRWMockRecorder) GetByEmailAndPassword(email, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmailAndPassword", reflect.TypeOf((*MockUserRW)(nil).GetByEmailAndPassword), email, password)
}

// Save mocks base method
func (m *MockUserRW) Save(user domain.User) error {
	ret := m.ctrl.Call(m, "Save", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockUserRWMockRecorder) Save(user interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRW)(nil).Save), user)
}

// MockArticleRW is a mock of ArticleRW interface
type MockArticleRW struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRWMockRecorder
}

// MockArticleRWMockRecorder is the mock recorder for MockArticleRW
type MockArticleRWMockRecorder struct {
	mock *MockArticleRW
}

// NewMockArticleRW creates a new mock instance
func NewMockArticleRW(ctrl *gomock.Controller) *MockArticleRW {
	mock := &MockArticleRW{ctrl: ctrl}
	mock.recorder = &MockArticleRWMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticleRW) EXPECT() *MockArticleRWMockRecorder {
	return m.recorder
}

// GetByAuthorsNameOrderedByMostRecentAsc mocks base method
func (m *MockArticleRW) GetByAuthorsNameOrderedByMostRecentAsc(usernames []string) ([]domain.Article, error) {
	ret := m.ctrl.Call(m, "GetByAuthorsNameOrderedByMostRecentAsc", usernames)
	ret0, _ := ret[0].([]domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAuthorsNameOrderedByMostRecentAsc indicates an expected call of GetByAuthorsNameOrderedByMostRecentAsc
func (mr *MockArticleRWMockRecorder) GetByAuthorsNameOrderedByMostRecentAsc(usernames interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAuthorsNameOrderedByMostRecentAsc", reflect.TypeOf((*MockArticleRW)(nil).GetByAuthorsNameOrderedByMostRecentAsc), usernames)
}

// GetRecentFiltered mocks base method
func (m *MockArticleRW) GetRecentFiltered(filters uc.Filters) ([]domain.Article, error) {
	ret := m.ctrl.Call(m, "GetRecentFiltered", filters)
	ret0, _ := ret[0].([]domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecentFiltered indicates an expected call of GetRecentFiltered
func (mr *MockArticleRWMockRecorder) GetRecentFiltered(filters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentFiltered", reflect.TypeOf((*MockArticleRW)(nil).GetRecentFiltered), filters)
}

// MockUserValidator is a mock of UserValidator interface
type MockUserValidator struct {
	ctrl     *gomock.Controller
	recorder *MockUserValidatorMockRecorder
}

// MockUserValidatorMockRecorder is the mock recorder for MockUserValidator
type MockUserValidatorMockRecorder struct {
	mock *MockUserValidator
}

// NewMockUserValidator creates a new mock instance
func NewMockUserValidator(ctrl *gomock.Controller) *MockUserValidator {
	mock := &MockUserValidator{ctrl: ctrl}
	mock.recorder = &MockUserValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserValidator) EXPECT() *MockUserValidatorMockRecorder {
	return m.recorder
}

// CheckUser mocks base method
func (m *MockUserValidator) CheckUser(user domain.User) error {
	ret := m.ctrl.Call(m, "CheckUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUser indicates an expected call of CheckUser
func (mr *MockUserValidatorMockRecorder) CheckUser(user interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockUserValidator)(nil).CheckUser), user)
}
