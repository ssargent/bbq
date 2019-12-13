package session

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
	"github.com/lib/pq"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	mock_bbq "github.com/ssargent/bbq/bbq-apiserver/bbq/mocks"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/security"
	mock_security "github.com/ssargent/bbq/bbq-apiserver/security/mocks"
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

func getSessionHelper(id int, name string, description string, sessionType string, weight float64, device string, monitor string, start time.Time, endTime pq.NullTime, uid uuid.UUID, tenant uuid.UUID) bbq.Session {
	return bbq.Session{
		ID:          id,
		Name:        name,
		Description: description,
		Type:        sessionType,
		Weight:      weight,
		Device:      device,
		Monitor:     monitor,
		StartTime:   start,
		EndTime:     endTime,
		TenantID:    tenant,
		UID:         uid,
	}
}

func getSimpleSessionHelper(tenant uuid.UUID) bbq.Session {
	food := "Smoked Goat"
	notFinished := pq.NullTime{
		Valid: false,
	}

	return getSessionHelper(4, food, food, "", 10, "LBGE", "Inkbird 4 Probe", time.Now(), notFinished, uuid.New(), tenant)
}

func TestGetSessionsEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountId, _ := uuid.NewUUID()
	tenant, _ := uuid.NewUUID()

	mon := getSimpleSessionHelper(tenant)
	loginSession := getLoginSessionHelper(accountId, tenant, "chef", "Chef Hetfield")

	auth := jwtauth.New("HS256", []byte("password"), nil)

	testConfig := config.Config{
		TokenAuth: auth,
	}

	authenticationService := mock_security.NewMockAuthenticationService(mockCtrl)
	monitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	monitorHandler := NewSessionHandler(&testConfig, authenticationService, monitorService)

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
