package monitor

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/ssargent/go-bbq/bbq"
	"github.com/ssargent/go-bbq/config"
	"github.com/ssargent/go-bbq/internal/infrastructure"
	"github.com/ssargent/go-bbq/security"
)

type monitorHandler struct {
	service        bbq.MonitorService
	authentication security.AuthenticationService
	config         *config.Config
}

// NewDeviceHandler will create an api Handler for a devices.
func NewMonitorHandler(config *config.Config, authentication security.AuthenticationService, service bbq.MonitorService) infrastructure.ApiHandler {
	return &monitorHandler{service: service, authentication: authentication, config: config}
}

func (handler *monitorHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.getMonitors)
	router.Get("/address/{address}", handler.getMonitorByAddress)
	router.Get("/{monitorName}", handler.getMonitorByName)
	router.Post("/", handler.createMonitor)
	router.Delete("/{monitorId}", handler.deleteMonitor)

	return router
}

func (handler *monitorHandler) getMonitors(w http.ResponseWriter, r *http.Request) {
	loginSession, err := handler.authentication.GetLoginSession(r)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	monitors, err := handler.service.GetMonitors(loginSession.TenantId)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, monitors)
}

func (handler *monitorHandler) getMonitorByAddress(w http.ResponseWriter, r *http.Request) {
	loginSession, err := handler.authentication.GetLoginSession(r)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	monitor, err := handler.service.GetMonitorByAddress(loginSession.TenantId, chi.URLParam(r, "address"))

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, monitor)
}

func (handler *monitorHandler) getMonitorByName(w http.ResponseWriter, r *http.Request) {
	loginSession, err := handler.authentication.GetLoginSession(r)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	monitorName := strings.Replace(chi.URLParam(r, "monitorName"), "-", " ", -1)
	monitorName = strings.ToLower(monitorName)
	monitor, err := handler.service.GetMonitorByName(loginSession.TenantId, monitorName)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, monitor)
}

func (handler *monitorHandler) createMonitor(w http.ResponseWriter, r *http.Request) {

}

func (handler *monitorHandler) deleteMonitor(w http.ResponseWriter, r *http.Request) {

}
