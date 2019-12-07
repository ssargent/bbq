package session

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
	"github.com/ssargent/bbq/bbq-apiserver/security"
)

type sessionHandler struct {
	service        bbq.SessionService
	authentication security.AuthenticationService
	config         *config.Config
}

// NewDeviceHandler will create an api Handler for a devices.
func NewSessionHandler(config *config.Config, authentication security.AuthenticationService, service bbq.SessionService) infrastructure.ApiHandler {
	return &sessionHandler{service: service, authentication: authentication, config: config}
}

func (handler *sessionHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.getSessions)
	router.Get("/address/{address}", handler.getSessionsByMonitorAddress)
	/*	router.Post("/", handler.createMonitor)
		router.Delete("/{monitorName}", handler.deleteMonitor)
	*/
	return router
}

func (handler *sessionHandler) getSessions(w http.ResponseWriter, r *http.Request) {
	loginSession, err := handler.authentication.GetLoginSession(r)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	sessions, err := handler.service.GetSessions(loginSession.TenantId)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, sessions)
}

func (handler *sessionHandler) getSessionsByMonitorAddress(w http.ResponseWriter, r *http.Request) {
	loginSession, err := handler.authentication.GetLoginSession(r)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	monitor, err := handler.service.GetSessionByMonitorAddress(loginSession.TenantId, chi.URLParam(r, "address"))

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, monitor)
}
