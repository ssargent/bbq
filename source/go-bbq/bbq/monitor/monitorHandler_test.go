package monitor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/ssargent/go-bbq/bbq"
	mock_bbq "github.com/ssargent/go-bbq/bbq/mocks"
	"github.com/ssargent/go-bbq/config"
	"github.com/ssargent/go-bbq/security"
	mock_security "github.com/ssargent/go-bbq/security/mocks"
	"github.com/stretchr/testify/assert"
)

func getLoginSessionHelper(accountId uuid.UUID, tenantId uuid.UUID, login string, name string) security.LoginSession {
	claims := jwt.MapClaims{
		"sub":    accountId,
		"tenant": tenantId,
		"iss":    "https://bbq.k8s.ssargent.net/",
		"aud":    "https://bbq.k8s.ssargent.net/",
		"exp":    time.Now().Add(time.Second * time.Duration(100000)).Unix(),
		"iat":    time.Now().Unix(),
		"login":  "chef",
		"fn":     "Chef Hetfield",
	}

	loginSession := security.LoginSession{
		AccountId: accountId,
		TenantId:  tenantId,
		LoginName: "chef",
		FullName:  "Chef Hetfield",
		Claims:    claims,
	}

	return loginSession
}

func getMonitorHelper(id int, name string, address string, tenant uuid.UUID) bbq.Monitor {
	return bbq.Monitor{
		ID:          id,
		Name:        name,
		Description: name,
		Address:     address,
		TenantID:    tenant,
	}
}

func TestGetMonitorsEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountId, _ := uuid.NewUUID()
	tenant, _ := uuid.NewUUID()

	mon := getMonitorHelper(1, "My Monitor", "deadbeef", tenant)
	loginSession := getLoginSessionHelper(accountId, tenant, "chef", "Chef Hetfield")

	auth := jwtauth.New("HS256", []byte("password"), nil)

	testConfig := config.Config{
		TokenAuth: auth,
	}

	authenticationService := mock_security.NewMockAuthenticationService(mockCtrl)
	monitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	monitorHandler := NewMonitorHandler(&testConfig, authenticationService, monitorService)

	authenticationService.EXPECT().GetLoginSession(gomock.Any()).Return(loginSession, nil).Times(1)
	monitorService.EXPECT().GetMonitors(tenant).Return([]bbq.Monitor{mon}, nil).Times(1)

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	var monitorsResult []bbq.Monitor

	monitorHandler.Routes().ServeHTTP(response, request)

	if err := json.NewDecoder(response.Body).Decode(&monitorsResult); err != nil {
		t.Error("Cannot convert json to monitor collection")
	}

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
	assert.NotEqual(t, "[]", response.Body.String())
	assert.NotEmpty(t, monitorsResult)
	assert.ElementsMatch(t, []bbq.Monitor{mon}, monitorsResult)
}

func TestGetMonitorByAddressEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountId, _ := uuid.NewUUID()
	tenant, _ := uuid.NewUUID()

	mon := getMonitorHelper(1, "My Monitor", "deadbeef", tenant)
	loginSession := getLoginSessionHelper(accountId, tenant, "chef", "Chef Hetfield")

	auth := jwtauth.New("HS256", []byte("password"), nil)

	testConfig := config.Config{
		TokenAuth: auth,
	}

	authenticationService := mock_security.NewMockAuthenticationService(mockCtrl)
	monitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	monitorHandler := NewMonitorHandler(&testConfig, authenticationService, monitorService)

	authenticationService.EXPECT().GetLoginSession(gomock.Any()).Return(loginSession, nil).Times(1)
	monitorService.EXPECT().GetMonitorByAddress(tenant, "deadbeef").Return(mon, nil).Times(1)

	request, _ := http.NewRequest("GET", "/address/deadbeef", nil)
	response := httptest.NewRecorder()
	var monitorResult bbq.Monitor

	monitorHandler.Routes().ServeHTTP(response, request)

	if err := json.NewDecoder(response.Body).Decode(&monitorResult); err != nil {
		t.Error("Cannot convert json to monitor")
	}

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
	assert.NotNil(t, monitorResult)
	assert.Equal(t, mon, monitorResult)

}
