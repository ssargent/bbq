package temperature

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"

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
type Reading struct {
	ID         int       `json:"id"`
	Probe0     float32   `json:"probe0"`
	Probe1     float32   `json:"probe1"`
	Probe2     float32   `json:"probe2"`
	Probe3     float32   `json:"probe3"`
	RecordedAt time.Time `json:"recordedat"`
	SessionID  uuid.UUID `json:"sessionid"`
}

// TenantRoutes are the routes scoped to a specific tenant
func (config *Config) TenantRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{sessionid}", config.getTenantSessionReadings)
	//	router.Post("/{sessionid}", config.createTenantSessionReadings)

	return router
}

func (config *Config) getTenantSessionReadings(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	sessionid, err := uuid.Parse(chi.URLParam(r, "sessionid"))

	if err != nil {
		fmt.Println("Error Parsing UUID", err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	readings, err := getTenantSessionReadings(config.Database, tenantKey, sessionid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	render.JSON(w, r, readings)

}

func (config *Config) createTenantSessionReadings(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	sessionid, err := uuid.Parse(chi.URLParam(r, "sessionid"))

	if err != nil {
		fmt.Println("Error Parsing UUID", err)
		http.Error(w, http.StatusText(400), 400)
		return
	} else {

		session, err := getTenantSessionReadings(config.Database, tenantKey, sessionid)

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		} else {

			render.JSON(w, r, session)
		}
	}
}
