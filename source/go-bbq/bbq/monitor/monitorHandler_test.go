package monitor

import (
	"context"
	"fmt"
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
	_, tokenString, _ := auth.Encode(claims)

	testConfig := config.Config{
		TokenAuth: auth,
	}

	monitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	monitorHandler := NewMonitorHandler(&testConfig, monitorService)

	monitorService.EXPECT().GetMonitors(tenant).Return([]bbq.Monitor{mon}, nil).Times(1)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "jwt", claims)

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	fmt.Print(tokenString)

	monitorHandler.Routes().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
