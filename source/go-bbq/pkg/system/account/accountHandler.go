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

type loginModel struct {
	loginname string
	password  string
}

type loginResult struct {
	success bool
	token   string
}

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
	router.Post("/login", handler.login)

	return router
}

func (handler *accountHandler) login(w http.ResponseWriter, r *http.Request) {
	newLogin := loginModel{}

	if err := json.NewDecoder(r.Body).Decode(&newLogin); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	authenticated, err := handler.service.Login(newLogin.loginname, newLogin.password)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	jwtToken := handler.service.CreateToken(authenticated)

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
