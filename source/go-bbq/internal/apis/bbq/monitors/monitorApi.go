package monitors
 
import (
	"net/http"
	"strconv"
	"encoding/json"
 
	"github.com/google/uuid"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/ssargent/go-bbq/internal/infrastructure"
	"github.com/ssargent/go-bbq/internal/config"
)


// Config - duck typing?
type Config struct {
	*config.Config
}

// New returns a new instance of a config
func New(configuration *config.Config) *Config {
	return &Config{configuration}
}


// Monitor represents a bbq monitor
type Monitor struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Address string `json:"address"`
	TenantID uuid.UUID `json:"tenantid"`
}

// TenantRoutes are the routes scoped to a specific tenant
func (config *Config)  TenantRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", config.getTenantMonitors)
	router.Get("/{monitorid}", config.getTenantMonitor)
	router.Post("/", config.createTenantMonitor)
	router.Delete("/{monitorid}", config.deleteTenantMonitor)

	return router;
}

func (config *Config) getTenantMonitors(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	monitors, err := getTenantMonitors(config.Database, tenantKey)

    if err != nil { 
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
    }

    render.JSON(w, r, monitors)
}

func (config *Config) getTenantMonitor(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	monitorid,err := strconv.Atoi(chi.URLParam(r, "monitorid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	monitor, err := getTenantMonitor(config.Database, tenantKey, monitorid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	render.JSON(w, r, monitor)
}

func (config *Config) createTenantMonitor(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")

	data := Monitor{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	monitor, err := createTenantMonitor(config.Database, tenantKey, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, monitor)
}

func (config *Config) deleteTenantMonitor(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	monitorid,err := strconv.Atoi(chi.URLParam(r, "monitorid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	err = deleteTenantMonitor(config.Database, tenantKey, monitorid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	w.WriteHeader(http.StatusNoContent)
    w.Write([]byte("204 - No Content"))
}
