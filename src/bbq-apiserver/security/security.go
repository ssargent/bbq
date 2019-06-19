package security

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//go:generate mockgen  -destination=./mocks/security.go -package=mock_security github.com/ssargent/bbq/bbq-apiserver/security AuthenticationService

type LoginSession struct {
	AccountId uuid.UUID
	TenantId  uuid.UUID
	LoginName string
	FullName  string
	Claims    jwt.MapClaims
}

type AuthenticationService interface {
	GetLoginSession(request *http.Request) (LoginSession, error)
}
