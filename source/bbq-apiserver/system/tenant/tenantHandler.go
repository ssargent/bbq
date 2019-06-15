package tenant

import (
	"encoding/json"
	"net/http"

	//	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	//"github.com/google/uuid"

	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
	"github.com/ssargent/bbq/bbq-apiserver/system"
)

type newTenantModel struct {
	TenantName string `json:"tenantName"`
	URLKey     string `json:"urlKey"`
	LoginName  string `json:"loginName"`
	Password   string `json:"password"`
	FullName   string `json:"fullName"`
	Email      string `json:"email"`
}

type newTenantResult struct {
	Tenant  system.Tenant  `json:"tenant"`
	Account system.Account `json:"account"`
}

type tenantHandler struct {
	accountService system.AccountService
	tenantService  system.TenantService
	config         *config.Config
}

// NewAccountHandler will create an api Handler for a new account.
func NewTenantHandler(config *config.Config, tenantService system.TenantService, accountService system.AccountService) infrastructure.ApiHandler {
	return &tenantHandler{tenantService: tenantService, accountService: accountService, config: config}
}

func (handler *tenantHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	// This api is special in that it requires some public apis.  but it also
	// has private apis.  Protect those here.
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(handler.config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/{key}", handler.getByKey)

	})

	router.Post("/", handler.createTenant)
	return router
}

func (handler *tenantHandler) createTenant(w http.ResponseWriter, r *http.Request) {
	newTenant := newTenantModel{}

	// need a TenantModel here to encapsulate the account creation process as well.
	if err := json.NewDecoder(r.Body).Decode(&newTenant); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	tenant := system.Tenant{
		Name:      newTenant.TenantName,
		URLKey:    newTenant.URLKey,
		IsEnabled: true,
	}

	createdTenant, err := handler.tenantService.CreateTenant(tenant)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	account := system.Account{
		LoginName:     newTenant.LoginName,
		LoginPassword: newTenant.Password,
		FullName:      newTenant.FullName,
		Email:         newTenant.Email,
		IsEnabled:     true,
		TenantID:      createdTenant.ID,
	}

	createdAccount, err := handler.accountService.CreateAccount(account)

	if err != nil {
		//Possibly Mark the tenant as eligible for cleanup/delete.
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))

		err = handler.tenantService.DeleteTenant(createdTenant)

		return
	}

	result := newTenantResult{
		Tenant:  createdTenant,
		Account: createdAccount,
	}

	render.JSON(w, r, result)
}

func (handler *tenantHandler) getByKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	if len(key) == 0 {
		http.Error(w, http.StatusText(400), 400)
	}

	tenant, err := handler.tenantService.GetByKey(key)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	render.JSON(w, r, tenant)
}
