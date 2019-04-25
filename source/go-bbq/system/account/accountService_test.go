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

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockRepo.EXPECT().GetByLogin("chef").Return(login, nil).Times(1)
	mockCacheService.EXPECT().GetItem("system$accounts$chef", &returnedLogin).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem("system$accounts$chef", login, time.Minute*10).Return(nil).Times(1)

	accountService.GetAccount("chef")
}
