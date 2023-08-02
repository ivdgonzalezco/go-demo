// Code generated by MockGen. DO NOT EDIT.
// Source: ./uc/INTERACTOR.go

// Package uc is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/err0r500/go-realworld-clean/domain"
	gomock "github.com/golang/mock/gomock"
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

// Create mocks base method
func (m *MockArticleRW) Create(arg0 domain.Article) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockArticleRWMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleRW)(nil).Create), arg0)
}

// Save mocks base method
func (m *MockArticleRW) Save(arg0 domain.Article) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockArticleRWMockRecorder) Save(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockArticleRW)(nil).Save), arg0)
}

// GetBySlug mocks base method
func (m *MockArticleRW) GetBySlug(slug string) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "GetBySlug", slug)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug
func (mr *MockArticleRWMockRecorder) GetBySlug(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockArticleRW)(nil).GetBySlug), slug)
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
func (m *MockArticleRW) GetRecentFiltered(filters []domain.ArticleFilter) ([]domain.Article, error) {
	ret := m.ctrl.Call(m, "GetRecentFiltered", filters)
	ret0, _ := ret[0].([]domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecentFiltered indicates an expected call of GetRecentFiltered
func (mr *MockArticleRWMockRecorder) GetRecentFiltered(filters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentFiltered", reflect.TypeOf((*MockArticleRW)(nil).GetRecentFiltered), filters)
}

// Delete mocks base method
func (m *MockArticleRW) Delete(slug string) error {
	ret := m.ctrl.Call(m, "Delete", slug)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockArticleRWMockRecorder) Delete(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleRW)(nil).Delete), slug)
}

// MockCommentRW is a mock of CommentRW interface
type MockCommentRW struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRWMockRecorder
}

// MockCommentRWMockRecorder is the mock recorder for MockCommentRW
type MockCommentRWMockRecorder struct {
	mock *MockCommentRW
}

// NewMockCommentRW creates a new mock instance
func NewMockCommentRW(ctrl *gomock.Controller) *MockCommentRW {
	mock := &MockCommentRW{ctrl: ctrl}
	mock.recorder = &MockCommentRWMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentRW) EXPECT() *MockCommentRWMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCommentRW) Create(comment domain.Comment) (*domain.Comment, error) {
	ret := m.ctrl.Call(m, "Create", comment)
	ret0, _ := ret[0].(*domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCommentRWMockRecorder) Create(comment interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommentRW)(nil).Create), comment)
}

// GetByID mocks base method
func (m *MockCommentRW) GetByID(id int) (*domain.Comment, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockCommentRWMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCommentRW)(nil).GetByID), id)
}

// Delete mocks base method
func (m *MockCommentRW) Delete(id int) error {
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCommentRWMockRecorder) Delete(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentRW)(nil).Delete), id)
}

// MockTagsRW is a mock of TagsRW interface
type MockTagsRW struct {
	ctrl     *gomock.Controller
	recorder *MockTagsRWMockRecorder
}

// MockTagsRWMockRecorder is the mock recorder for MockTagsRW
type MockTagsRWMockRecorder struct {
	mock *MockTagsRW
}

// NewMockTagsRW creates a new mock instance
func NewMockTagsRW(ctrl *gomock.Controller) *MockTagsRW {
	mock := &MockTagsRW{ctrl: ctrl}
	mock.recorder = &MockTagsRWMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagsRW) EXPECT() *MockTagsRWMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockTagsRW) GetAll() ([]string, error) {
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockTagsRWMockRecorder) GetAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTagsRW)(nil).GetAll))
}

// Add mocks base method
func (m *MockTagsRW) Add(newTags []string) error {
	ret := m.ctrl.Call(m, "Add", newTags)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockTagsRWMockRecorder) Add(newTags interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTagsRW)(nil).Add), newTags)
}

// MockSlugger is a mock of Slugger interface
type MockSlugger struct {
	ctrl     *gomock.Controller
	recorder *MockSluggerMockRecorder
}

// MockSluggerMockRecorder is the mock recorder for MockSlugger
type MockSluggerMockRecorder struct {
	mock *MockSlugger
}

// NewMockSlugger creates a new mock instance
func NewMockSlugger(ctrl *gomock.Controller) *MockSlugger {
	mock := &MockSlugger{ctrl: ctrl}
	mock.recorder = &MockSluggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSlugger) EXPECT() *MockSluggerMockRecorder {
	return m.recorder
}

// NewSlug mocks base method
func (m *MockSlugger) NewSlug(arg0 string) string {
	ret := m.ctrl.Call(m, "NewSlug", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// NewSlug indicates an expected call of NewSlug
func (mr *MockSluggerMockRecorder) NewSlug(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSlug", reflect.TypeOf((*MockSlugger)(nil).NewSlug), arg0)
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

// MockArticleValidator is a mock of ArticleValidator interface
type MockArticleValidator struct {
	ctrl     *gomock.Controller
	recorder *MockArticleValidatorMockRecorder
}

// MockArticleValidatorMockRecorder is the mock recorder for MockArticleValidator
type MockArticleValidatorMockRecorder struct {
	mock *MockArticleValidator
}

// NewMockArticleValidator creates a new mock instance
func NewMockArticleValidator(ctrl *gomock.Controller) *MockArticleValidator {
	mock := &MockArticleValidator{ctrl: ctrl}
	mock.recorder = &MockArticleValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticleValidator) EXPECT() *MockArticleValidatorMockRecorder {
	return m.recorder
}

// BeforeCreationCheck mocks base method
func (m *MockArticleValidator) BeforeCreationCheck(article *domain.Article) error {
	ret := m.ctrl.Call(m, "BeforeCreationCheck", article)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeCreationCheck indicates an expected call of BeforeCreationCheck
func (mr *MockArticleValidatorMockRecorder) BeforeCreationCheck(article interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeCreationCheck", reflect.TypeOf((*MockArticleValidator)(nil).BeforeCreationCheck), article)
}

// BeforeUpdateCheck mocks base method
func (m *MockArticleValidator) BeforeUpdateCheck(article *domain.Article) error {
	ret := m.ctrl.Call(m, "BeforeUpdateCheck", article)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeUpdateCheck indicates an expected call of BeforeUpdateCheck
func (mr *MockArticleValidatorMockRecorder) BeforeUpdateCheck(article interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeUpdateCheck", reflect.TypeOf((*MockArticleValidator)(nil).BeforeUpdateCheck), article)
}
