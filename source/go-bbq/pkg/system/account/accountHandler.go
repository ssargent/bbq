package account

import (
	"encoding/json"
	"net/http"
	//	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	//"github.com/google/uuid"

	"github.com/ssargent/go-bbq/internal/infrastructure"
	"github.com/ssargent/go-bbq/pkg"
	"github.com/ssargent/go-bbq/pkg/system"
)

type accountHandler struct {
	service system.AccountService
}

// NewAccountHandler will create an api Handler for a new account.
func NewAccountHandler(service system.AccountService) pkg.ApiHandler {
	return &accountHandler{service: service}
}

func (handler *accountHandler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{login}", handler.getAccountByLogin)
	router.Post("/", handler.createAccount)

	return router
}

func (handler *accountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	account := system.Account{}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	createdAccount, err := handler.service.CreateAccount(account)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, createdAccount)
}

func (handler *accountHandler) getAccountByLogin(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")

	if len(login) == 0 {
		http.Error(w, http.StatusText(400), 400)
	}

	account, err := handler.service.GetAccount(login)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}
	account.LoginPassword = ""
	render.JSON(w, r, account)
}
