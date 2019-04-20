package account

import (
	"encoding/json"
	"net/http"
	"time"

	//	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	//"github.com/google/uuid"

	"github.com/ssargent/go-bbq/config"
	"github.com/ssargent/go-bbq/internal/infrastructure"
	"github.com/ssargent/go-bbq/system"
)

// LoginModel is a simple model to capture logins.
type LoginModel struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

// LoginResult captues whether the login was successful
type LoginResult struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Error   string `json:"error,omitempty"`
}

type accountHandler struct {
	service system.AccountService
	config  *config.Config
}

// NewAccountHandler will create an api Handler for a new account.
func NewAccountHandler(config *config.Config, service system.AccountService) infrastructure.ApiHandler {
	return &accountHandler{service: service, config: config}
}

func (handler *accountHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	// This api is special in that it requires some public apis.  but it also
	// has private apis.  Protect those here.
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(handler.config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/{login}", handler.getAccountByLogin)

	})

	router.Post("/", handler.createAccount)
	router.Post("/login", handler.login)
	router.Post("/signin", handler.signin)
	return router
}

func (handler *accountHandler) createToken(account system.Account) string {
	claims := jwt.MapClaims{
		"sub":   account.ID,
		"iss":   "https://bbq.k8s.ssargent.net/",
		"aud":   "https://bbq.k8s.ssargent.net/",
		"exp":   time.Now().Add(time.Second * time.Duration(100000)).Unix(),
		"iat":   time.Now().Unix(),
		"login": account.LoginName,
		"fn":    account.FullName,
	}
	_, tokenString, _ := handler.config.TokenAuth.Encode(claims)
	return tokenString
}

// signin will retrun a jwt cookie.  This is used for signing in locally via the web app.
// if you need api style access, use login.
func (handler *accountHandler) signin(w http.ResponseWriter, r *http.Request) {
	newLogin := LoginModel{}

	if err := json.NewDecoder(r.Body).Decode(&newLogin); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	authenticated, err := handler.service.Login(newLogin.LoginName, newLogin.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, LoginResult{Success: false, Error: err.Error()})
		return
	}

	jwtToken := handler.createToken(authenticated)
	result := LoginResult{Success: true, Token: jwtToken}

	c := http.Cookie{
		Name:     "jwt",
		Value:    jwtToken,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &c)

	result.Token = ""
	render.JSON(w, r, result)
}

func (handler *accountHandler) login(w http.ResponseWriter, r *http.Request) {
	newLogin := LoginModel{}

	if err := json.NewDecoder(r.Body).Decode(&newLogin); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	authenticated, err := handler.service.Login(newLogin.LoginName, newLogin.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, LoginResult{Success: false, Error: err.Error()})
		return
	}

	jwtToken := handler.createToken(authenticated)

	result := LoginResult{Success: true, Token: jwtToken}

	render.JSON(w, r, result)
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
