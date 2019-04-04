package account 

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/ssargent/go-bbq/pkg"
)

type accountHandler struct {
	service *system.AccountService
}

func NewAccountHandler(service *system.AccountService) *ApiHandler (
	return &accountHandler{ service: service}
)

func (handler *accountHandler) Register() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{login}", api.getAccountByLogin)

	return router
}

func (handler *accountHandler) getAccountByLogin(w http.ResponseWriter, r *http.Request) {
	login, err := strconv.Atoi(chi.URLParam(r, "login"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	account, err := service.GetAccount(login)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	render.JSON(w, r, account)
}