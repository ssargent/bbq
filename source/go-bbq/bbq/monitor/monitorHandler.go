package monitor

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ssargent/go-bbq/bbq"
	"github.com/ssargent/go-bbq/config"
	"github.com/ssargent/go-bbq/internal/infrastructure"
)

type monitorHandler struct {
	service bbq.MonitorService
	config  *config.Config
}

// NewDeviceHandler will create an api Handler for a devices.
func NewMonitorHandler(config *config.Config, service bbq.MonitorService) infrastructure.ApiHandler {
	return &monitorHandler{service: service, config: config}
}

func (handler *monitorHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.getMonitors)
	router.Get("/address/{address}", handler.getMonitorByAddress)
	router.Get("/{deviceName}", handler.getMonitorByName)
	router.Post("/", handler.createMonitor)
	router.Delete("/{deviceid}", handler.deleteMonitor)

	return router
}

func (handler *monitorHandler) getMonitors(w http.ResponseWriter, r *http.Request) {

}

func (handler *monitorHandler) getMonitorByAddress(w http.ResponseWriter, r *http.Request) {

}

func (handler *monitorHandler) getMonitorByName(w http.ResponseWriter, r *http.Request) {

}

func (handler *monitorHandler) createMonitor(w http.ResponseWriter, r *http.Request) {

}

func (handler *monitorHandler) deleteMonitor(w http.ResponseWriter, r *http.Request) {

}
