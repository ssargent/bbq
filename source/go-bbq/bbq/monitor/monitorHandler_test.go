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

func TestGetMonitorsEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountId, _ := uuid.NewUUID()
	tenant, _ := uuid.NewUUID()
	mon := bbq.Monitor{
		ID:          1,
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}

	auth := jwtauth.New("HS256", []byte("password"), nil)
	claims := jwt.MapClaims{
		"sub":    accountId,
		"tenant": tenant,
		"iss":    "https://bbq.k8s.ssargent.net/",
		"aud":    "https://bbq.k8s.ssargent.net/",
		"exp":    time.Now().Add(time.Second * time.Duration(100000)).Unix(),
		"iat":    time.Now().Unix(),
		"login":  "chef",
		"fn":     "Chef Hetfield",
	}

	testConfig := config.Config{
		TokenAuth: auth,
	}

	loginSession := security.LoginSession{
		AccountId: accountId,
		TenantId:  tenant,
		LoginName: "chef",
		FullName:  "Chef Hetfield",
		Claims:    claims,
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
