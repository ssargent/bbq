package monitor

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
	"github.com/ssargent/bbq/bbq-apiserver/security"
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
	router.Delete("/{monitorName}", handler.deleteMonitor)

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
	loginSession, err := handler.authentication.GetLoginSession(r)

	data := bbq.Monitor{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	monitor, err := handler.service.CreateMonitor(loginSession.TenantId, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, monitor)
}

func (handler *monitorHandler) deleteMonitor(w http.ResponseWriter, r *http.Request) {
	loginSession, err := handler.authentication.GetLoginSession(r)

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	monitorName := strings.Replace(chi.URLParam(r, "monitorName"), "-", " ", -1)
	monitorName = strings.ToLower(monitorName)

	monitor, err := handler.service.GetMonitorByName(loginSession.TenantId, monitorName)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	err = handler.service.DeleteMonitor(loginSession.TenantId, monitor)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("204 - No Content"))
}
