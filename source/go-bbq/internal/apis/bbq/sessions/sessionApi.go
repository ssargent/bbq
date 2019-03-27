package sessions

import (
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	//"encoding/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/lib/pq"
	"github.com/ssargent/go-bbq/internal/config"
	"github.com/ssargent/go-bbq/internal/infrastructure"
)

// Session represents a bbq session.
type Session struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Subject     string      `json:"subject"`
	Type        string      `json:"type"`
	Weight      float64     `json:"weight"`
	Device      string      `json:"device"`
	Monitor     string      `json:"monitor"`
	StartTime   time.Time   `json:"starttime"`
	EndTime     pq.NullTime `json:"endtime"`
	TenantID    uuid.UUID   `json:"tenantid"`
	UID         uuid.UUID   `json:"uid"`
}

type sessionRecord struct {
	ID          int         `json:"id"`
	DeviceID    int         `json:"deviceid"`
	MonitorID   int         `json:"monitorid"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	StartTime   time.Time   `json:"starttime"`
	SubjectID   int         `json:"subjectid"`
	Weight      float64     `json:"weight"`
	TenantID    uuid.UUID   `json:"tenantid"`
	UID         uuid.UUID   `json:"uid"`
	EndTime     pq.NullTime `json:"endtime"`
}

// Config - duck typing?
type Config struct {
	*config.Config
}

// New returns a new instance of a config
func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

// TenantRoutes are the routes scoped to a specific tenant
func (config *Config) TenantRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", config.getTenantSessions)
	router.Get("/{sessionid}", config.getTenantSession)
	router.Post("/", config.createTenantSession)
	//router.Post("/{sessionid}/probe-data", config.recordProbeData)
	//router.Delete("/{sessionid}", config.deleteTenantSession)

	return router
}

/*
func (config *Config) recordProbeData(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	sessionid, err := uuid.Parse(chi.URLParam(r, "sessionid"))

	result := make([]ProbeData, 0)


}*/

func (config *Config) getTenantSessions(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	sessions, err := getTenantSessions(config.Database, tenantKey)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, sessions)
}

func (config *Config) getTenantSession(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	fmt.Println("Trying to load sessionid from request")
	//sessionid, err := strconv.Atoi(chi.URLParam(r, "sessionid"))
	sessionid, err := uuid.Parse(chi.URLParam(r, "sessionid"))

	if err != nil {
		fmt.Println("Error Parsing UUID", err)
		http.Error(w, http.StatusText(400), 400)

		render.Status(r, 500)
	} else {

		session, err := GetTenantSession(config.Database, tenantKey, sessionid)

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			render.Status(r, 404)
		} else {

			render.JSON(w, r, session)
		}
	}
}

func (config *Config) createTenantSession(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")

	data := Session{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	session, err := createTenantSession(config.Database, tenantKey, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, session)
}

func (config *Config) deleteTenantSession(w http.ResponseWriter, r *http.Request) {

}
