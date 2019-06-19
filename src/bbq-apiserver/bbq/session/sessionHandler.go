package session

import (
	"github.com/go-chi/chi"

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
	/*
		router.Get("/", handler.getMonitors)
		router.Get("/address/{address}", handler.getMonitorByAddress)
		router.Get("/{monitorName}", handler.getMonitorByName)
		router.Post("/", handler.createMonitor)
		router.Delete("/{monitorName}", handler.deleteMonitor)
	*/
	return router
}
