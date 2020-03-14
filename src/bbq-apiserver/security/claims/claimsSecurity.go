package claims

import (
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/security"
)

type claimsAuthenticationService struct {
}

func NewClaimsAuthenticationService() security.AuthenticationService {
	return &claimsAuthenticationService{}
}

func (c *claimsAuthenticationService) GetLoginSession(request *http.Request) (security.LoginSession, error) {
	_, claims, _ := jwtauth.FromContext(request.Context())

	tenantString := claims["tenant"].(string)
	accountString := claims["sub"].(string)
	loginString := claims["login"].(string)
	fullName := claims["fn"].(string)

	account, err := uuid.Parse(accountString)

	if err != nil {
		return security.LoginSession{}, errors.New("Cannot parse account from claims")
	}

	tenant, err := uuid.Parse(tenantString)

	if err != nil {
		return security.LoginSession{}, errors.New("Cannot parse tenant from claims")
	}

	loginSession := security.LoginSession{
		TenantId:  tenant,
		AccountId: account,
		LoginName: loginString,
		FullName:  fullName,
		Claims:    claims,
	}

	return loginSession, nil
}
