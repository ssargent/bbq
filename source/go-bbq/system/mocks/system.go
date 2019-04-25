// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ssargent/go-bbq/system (interfaces: TenantService,TenantRepository,AccountService,AccountRepository)

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

// CreateTenant mocks base method
func (m *MockTenantService) CreateTenant(arg0 system.Tenant) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", arg0)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant
func (mr *MockTenantServiceMockRecorder) CreateTenant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockTenantService)(nil).CreateTenant), arg0)
}

// DeleteTenant mocks base method
func (m *MockTenantService) DeleteTenant(arg0 system.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenant", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTenant indicates an expected call of DeleteTenant
func (mr *MockTenantServiceMockRecorder) DeleteTenant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenant", reflect.TypeOf((*MockTenantService)(nil).DeleteTenant), arg0)
}

// GetByKey mocks base method
func (m *MockTenantService) GetByKey(arg0 string) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKey", arg0)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByKey indicates an expected call of GetByKey
func (mr *MockTenantServiceMockRecorder) GetByKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKey", reflect.TypeOf((*MockTenantService)(nil).GetByKey), arg0)
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

// Create mocks base method
func (m *MockTenantRepository) Create(arg0 system.Tenant) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTenantRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTenantRepository)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockTenantRepository) Delete(arg0 system.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTenantRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTenantRepository)(nil).Delete), arg0)
}

// GetByKey mocks base method
func (m *MockTenantRepository) GetByKey(arg0 string) (system.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKey", arg0)
	ret0, _ := ret[0].(system.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByKey indicates an expected call of GetByKey
func (mr *MockTenantRepositoryMockRecorder) GetByKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKey", reflect.TypeOf((*MockTenantRepository)(nil).GetByKey), arg0)
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

// CreateAccount mocks base method
func (m *MockAccountService) CreateAccount(arg0 system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountServiceMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountService)(nil).CreateAccount), arg0)
}

// DeleteAccount mocks base method
func (m *MockAccountService) DeleteAccount(arg0 system.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockAccountServiceMockRecorder) DeleteAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccountService)(nil).DeleteAccount), arg0)
}

// GetAccount mocks base method
func (m *MockAccountService) GetAccount(arg0 string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockAccountServiceMockRecorder) GetAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountService)(nil).GetAccount), arg0)
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

// Login mocks base method
func (m *MockAccountService) Login(arg0, arg1 string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAccountServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAccountService)(nil).Login), arg0, arg1)
}

// UpdateAccount mocks base method
func (m *MockAccountService) UpdateAccount(arg0 system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockAccountServiceMockRecorder) UpdateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountService)(nil).UpdateAccount), arg0)
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

// Create mocks base method
func (m *MockAccountRepository) Create(arg0 system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAccountRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepository)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockAccountRepository) Delete(arg0 system.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAccountRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccountRepository)(nil).Delete), arg0)
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

// GetByEmail mocks base method
func (m *MockAccountRepository) GetByEmail(arg0 string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail
func (mr *MockAccountRepositoryMockRecorder) GetByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockAccountRepository)(nil).GetByEmail), arg0)
}

// GetByID mocks base method
func (m *MockAccountRepository) GetByID(arg0 uuid.UUID) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockAccountRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccountRepository)(nil).GetByID), arg0)
}

// GetByLogin mocks base method
func (m *MockAccountRepository) GetByLogin(arg0 string) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByLogin", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByLogin indicates an expected call of GetByLogin
func (mr *MockAccountRepositoryMockRecorder) GetByLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLogin", reflect.TypeOf((*MockAccountRepository)(nil).GetByLogin), arg0)
}

// Update mocks base method
func (m *MockAccountRepository) Update(arg0 system.Account) (system.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(system.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockAccountRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountRepository)(nil).Update), arg0)
}
