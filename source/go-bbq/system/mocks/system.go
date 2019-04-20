// Code generated by MockGen. DO NOT EDIT.
// Source: system/system.go

// Package mock_system is a generated GoMock package.
package mock_system

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	system "github.com/ssargent/go-bbq/system"
)

// MockTenantService is a mock of TenantService interface
type MockTenantService struct {
	ctrl     *gomock.Controller
	recorder *MockTenantServiceMockRecorder
}

// MockTenantServiceMockRecorder is the mock recorder for MockTenantService
type MockTenantServiceMockRecorder struct {
	mock *MockTenantService
}

// NewMockTenantService creates a new mock instance
func NewMockTenantService(ctrl *gomock.Controller) *MockTenantService {
	mock := &MockTenantService{ctrl: ctrl}
	mock.recorder = &MockTenantServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTenantService) EXPECT() *MockTenantServiceMockRecorder {
	return m.recorder
}

// GetByKey mocks base method
func (m *MockTenantService) GetByKey(key string) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKey", key)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByKey indicates an expected call of GetByKey
func (mr *MockTenantServiceMockRecorder) GetByKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKey", reflect.TypeOf((*MockTenantService)(nil).GetByKey), key)
}

// CreateTenant mocks base method
func (m *MockTenantService) CreateTenant(tenant system.Tenant) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", tenant)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant
func (mr *MockTenantServiceMockRecorder) CreateTenant(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockTenantService)(nil).CreateTenant), tenant)
}

// DeleteTenant mocks base method
func (m *MockTenantService) DeleteTenant(tenant system.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenant", tenant)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTenant indicates an expected call of DeleteTenant
func (mr *MockTenantServiceMockRecorder) DeleteTenant(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenant", reflect.TypeOf((*MockTenantService)(nil).DeleteTenant), tenant)
}

// MockAccountService is a mock of AccountService interface
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// GetAccount mocks base method
func (m *MockAccountService) GetAccount(accountName string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", accountName)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockAccountServiceMockRecorder) GetAccount(accountName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountService)(nil).GetAccount), accountName)
}

// Login mocks base method
func (m *MockAccountService) Login(login, password string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", login, password)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAccountServiceMockRecorder) Login(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAccountService)(nil).Login), login, password)
}

// GetAccounts mocks base method
func (m *MockAccountService) GetAccounts() ([]system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts")
	ret0, _ := ret[0].([]system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts
func (mr *MockAccountServiceMockRecorder) GetAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAccountService)(nil).GetAccounts))
}

// CreateAccount mocks base method
func (m *MockAccountService) CreateAccount(account system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", account)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountServiceMockRecorder) CreateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountService)(nil).CreateAccount), account)
}

// UpdateAccount mocks base method
func (m *MockAccountService) UpdateAccount(account system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", account)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockAccountServiceMockRecorder) UpdateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountService)(nil).UpdateAccount), account)
}

// DeleteAccount mocks base method
func (m *MockAccountService) DeleteAccount(account system.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockAccountServiceMockRecorder) DeleteAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccountService)(nil).DeleteAccount), account)
}

// MockTenantRepository is a mock of TenantRepository interface
type MockTenantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTenantRepositoryMockRecorder
}

// MockTenantRepositoryMockRecorder is the mock recorder for MockTenantRepository
type MockTenantRepositoryMockRecorder struct {
	mock *MockTenantRepository
}

// NewMockTenantRepository creates a new mock instance
func NewMockTenantRepository(ctrl *gomock.Controller) *MockTenantRepository {
	mock := &MockTenantRepository{ctrl: ctrl}
	mock.recorder = &MockTenantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTenantRepository) EXPECT() *MockTenantRepositoryMockRecorder {
	return m.recorder
}

// GetByKey mocks base method
func (m *MockTenantRepository) GetByKey(key string) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKey", key)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByKey indicates an expected call of GetByKey
func (mr *MockTenantRepositoryMockRecorder) GetByKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKey", reflect.TypeOf((*MockTenantRepository)(nil).GetByKey), key)
}

// Create mocks base method
func (m *MockTenantRepository) Create(tenant system.Tenant) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", tenant)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTenantRepositoryMockRecorder) Create(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTenantRepository)(nil).Create), tenant)
}

// Delete mocks base method
func (m *MockTenantRepository) Delete(tenant system.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", tenant)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTenantRepositoryMockRecorder) Delete(tenant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTenantRepository)(nil).Delete), tenant)
}

// MockAccountRepository is a mock of AccountRepository interface
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// GetByEmail mocks base method
func (m *MockAccountRepository) GetByEmail(email string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail
func (mr *MockAccountRepositoryMockRecorder) GetByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockAccountRepository)(nil).GetByEmail), email)
}

// GetByID mocks base method
func (m *MockAccountRepository) GetByID(id uuid.UUID) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockAccountRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccountRepository)(nil).GetByID), id)
}

// GetByLogin mocks base method
func (m *MockAccountRepository) GetByLogin(accountName string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByLogin", accountName)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByLogin indicates an expected call of GetByLogin
func (mr *MockAccountRepositoryMockRecorder) GetByLogin(accountName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLogin", reflect.TypeOf((*MockAccountRepository)(nil).GetByLogin), accountName)
}

// GetAll mocks base method
func (m *MockAccountRepository) GetAll() ([]system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockAccountRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAccountRepository)(nil).GetAll))
}

// Create mocks base method
func (m *MockAccountRepository) Create(account system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", account)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAccountRepositoryMockRecorder) Create(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepository)(nil).Create), account)
}

// Update mocks base method
func (m *MockAccountRepository) Update(account system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", account)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockAccountRepositoryMockRecorder) Update(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountRepository)(nil).Update), account)
}

// Delete mocks base method
func (m *MockAccountRepository) Delete(account system.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAccountRepositoryMockRecorder) Delete(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccountRepository)(nil).Delete), account)
}
