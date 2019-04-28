package account

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/golang/mock/gomock"
	mock_infrastructure "github.com/ssargent/go-bbq/internal/infrastructure/mocks"
	"github.com/ssargent/go-bbq/system"
	mock_system "github.com/ssargent/go-bbq/system/mocks"
)

/*
type Account struct {
	ID            uuid.UUID `json:"id"`
	LoginName     string    `json:"loginName"`
	LoginPassword string    `json:"loginPassword,omitempty"`
	FullName      string    `json:"fullName"`
	Email         string    `json:"email"`
	IsEnabled     bool      `json:"isenabled"`
	TenantID      uuid.UUID `json:"tenantid"`
}
*/
func TestGetAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_system.NewMockAccountRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	accountService := NewAccountService(mockCacheService, mockRepo)

	id, err := uuid.NewUUID()

	if err != nil {
		return
	}

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	login := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}
	var returnedLogin system.Account

	mockRepo.EXPECT().GetByLogin("chef").Return(login, nil).Times(1)
	mockCacheService.EXPECT().GetItem("system$accounts$chef", &returnedLogin).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem("system$accounts$chef", login, time.Minute*10).Return(nil).Times(1)

	accountService.GetAccount("chef")
}

func TestGetAccountWhenCached(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_system.NewMockAccountRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	accountService := NewAccountService(mockCacheService, mockRepo)
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	tenant, err := uuid.NewUUID()
	if err != nil {
		return
	}

	login := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}
	var returnedLogin system.Account

	mockRepo.EXPECT().GetByLogin("chef").Return(login, nil).Times(0)
	mockCacheService.EXPECT().GetItem("system$accounts$chef", &returnedLogin).Return(nil).Times(1)
	mockCacheService.EXPECT().SetItem("system$accounts$chef", login, time.Minute*10).Return(nil).Times(0)

	accountService.GetAccount("chef")
}

func TestCreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_system.NewMockAccountRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	accountService := NewAccountService(mockCacheService, mockRepo)
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	tenant, err := uuid.NewUUID()
	if err != nil {
		return
	}

	emptyAccount := system.Account{}
	login := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}

	createdLogin := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}
	mockRepo.EXPECT().GetByLogin(login.LoginName).Return(emptyAccount, nil).Times(1)
	mockRepo.EXPECT().GetByEmail(login.Email).Return(emptyAccount, nil).Times(1)
	mockRepo.EXPECT().Create(gomock.Any()).Return(createdLogin, nil).Times(1)

	account, err := accountService.CreateAccount(login)

	if err != nil {
		t.Fatalf("Expected err to be nil got %s", err)
	}

	if account.Empty() {
		t.Fatal("Expected non empty account, but got empty account")
	}

	if account.LoginPassword == login.LoginPassword {
		t.Fatalf("Expected LoginPassword to be cleared, got %s", account.LoginPassword)
	}
}

func TestCreateAccountWhenLoginExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_system.NewMockAccountRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	accountService := NewAccountService(mockCacheService, mockRepo)
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	tenant, err := uuid.NewUUID()
	if err != nil {
		return
	}

	emptyAccount := system.Account{}
	login := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}

	mockRepo.EXPECT().GetByLogin(login.LoginName).Return(login, nil).Times(1)
	mockRepo.EXPECT().GetByEmail(login.Email).Return(emptyAccount, nil).Times(0)
	mockRepo.EXPECT().Create(gomock.Any()).Times(0)

	account, err := accountService.CreateAccount(login)

	if err == nil {
		t.Fatal("Expected err to be loginname already exists, got nil")
	}

	if !account.Empty() {
		t.Fatal("Expected empty account to be empty, got non empty account")
	}
}

func TestCreateAccountWhenEmailExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_system.NewMockAccountRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	accountService := NewAccountService(mockCacheService, mockRepo)
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	tenant, err := uuid.NewUUID()
	if err != nil {
		return
	}

	emptyAccount := system.Account{}
	login := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}

	mockRepo.EXPECT().GetByLogin(login.LoginName).Return(emptyAccount, nil).Times(1)
	mockRepo.EXPECT().GetByEmail(login.Email).Return(login, nil).MaxTimes(1)
	mockRepo.EXPECT().Create(gomock.Any()).MaxTimes(0)

	account, err := accountService.CreateAccount(login)

	if err == nil {
		t.Fatal("Expected err to be email already exists, got nil")
	}

	if !account.Empty() {
		t.Fatal("Expected empty account to be empty, got non empty account")
	}
}

func TestUpdateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_system.NewMockAccountRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	accountService := NewAccountService(mockCacheService, mockRepo)
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	tenant, err := uuid.NewUUID()
	if err != nil {
		return
	}

	emptyAccount := system.Account{}
	login := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}

	updatedLogin := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}
	mockRepo.EXPECT().GetByEmail(login.Email).Return(emptyAccount, nil).Times(1)
	mockRepo.EXPECT().Update(gomock.Any()).Return(updatedLogin, nil).Times(1)

	account, err := accountService.UpdateAccount(login)

	if err != nil {
		t.Fatalf("Expected err to be nil got %s", err)
	}

	if account.Empty() {
		t.Fatal("Expected non empty account, but got empty account")
	}

	if account.LoginPassword == login.LoginPassword {
		t.Fatalf("Expected LoginPassword to be cleared, got %s", account.LoginPassword)
	}
}

func TestUpdateAccountWhenEmailExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_system.NewMockAccountRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	accountService := NewAccountService(mockCacheService, mockRepo)
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	tenant, err := uuid.NewUUID()
	if err != nil {
		return
	}

	emptyAccount := system.Account{}

	login := system.Account{
		ID:            id,
		LoginName:     "chef",
		LoginPassword: "i+mU$tB3EnCry+eD",
		FullName:      "Chef Hetfield",
		Email:         "chef@myfamilycooks.com",
		IsEnabled:     true,
		TenantID:      tenant,
	}

	mockRepo.EXPECT().GetByEmail(login.Email).Return(emptyAccount, errors.New("a login with that email already exists.  please choose another")).MaxTimes(1)
	mockRepo.EXPECT().Update(gomock.Any()).MaxTimes(0)

	account, err := accountService.UpdateAccount(login)

	if err == nil {
		t.Fatal("Expected err to be email already exists, got nil")
	}

	if !account.Empty() {
		t.Fatal("Expected empty account to be empty, got non empty account")
	}
}
