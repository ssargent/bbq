package health

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-redis/cache"
	"github.com/ssargent/bbq/bbq-apiserver/config"
)

// DeviceApi - duck typing?
type HealthApi struct {
	config *config.Config
}

// Device is the representation of a bbq device
type HealthStatus struct {
	ApiServer string `json:"apiServer"`
	Database  string `json:"database"`
	Cache     string `json:"cache"`
}

// New returns a new instance of a config
func New(configuration *config.Config) *HealthApi {
	return &HealthApi{configuration}
}

// HealthRoutes are the routes scoped to a specific tenant
func (api *HealthApi) HealthRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/livecheck", api.LiveCheck)

	return router
}

// LiveCheck is the GET for /devices
func (api *HealthApi) LiveCheck(w http.ResponseWriter, r *http.Request) {

	health := &HealthStatus{}
	database := "ready"
	cacheStatus := "ready"
	apiServer := "ready"

	rows, err := api.config.Database.Query(
		"SELECT 1")

	defer rows.Close()
	if err != nil {
		database = "not-ready"
	}

	err = api.config.Cache.Set(&cache.Item{
		Key:        "HEALTH$CHECK",
		Object:     "health-check",
		Expiration: time.Minute * 10,
	})

	if err != nil {
		cacheStatus = "not-ready"
	}

	health.Cache = cacheStatus
	health.Database = database

	if health.Cache != "ready" || health.Database != "ready" {
		apiServer = "not-ready"
		w.WriteHeader(500)
	}

	health.ApiServer = apiServer

	render.JSON(w, r, health)
}
