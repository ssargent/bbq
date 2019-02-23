package devices 

import (
	"net/http"
	"strconv"
	"encoding/json"
 

	"github.com/google/uuid"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/ssargent/go-bbq/infrastructure"
	"github.com/ssargent/go-bbq/config"
)

// Config - duck typing?
type Config struct {
	*config.Config
}

// New returns a new instance of a config
func New(configuration *config.Config) *Config {
	return &Config{configuration}
}


// Device is the representation of a bbq device
type Device struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	TenantID uuid.UUID `json:"tenantid"`
}

// Routes is the total list of Routes tenants aside.
func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", config.GetAllDevices)
	router.Get("/{deviceid}", config.GetDevice)
	router.Post("/", config.CreateDevice)
	router.Delete("/{deviceid}", config.DeleteDevice)
  
	return router;
}

// TenantRoutes are the routes scoped to a specific tenant
func (config *Config)  TenantRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", config.getTenantDevices)
	router.Get("/{deviceid}", config.GetTenantDevice)
	router.Post("/", config.CreateTenantDevice)
	router.Delete("/{deviceid}", config.DeleteTenantDevice)

	return router;
}

// GetAllDevices is the GET for /devices
func (config *Config) GetAllDevices(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
        count = 10
    }
    if start < 0 {
        start = 0
	}
	
	devices, err := GetAllDevices(config.Database, count, start)
    if err != nil { 
		render.JSON(w, r, err.Error())
        return
    }

    render.JSON(w, r, devices)
}

// GetDevice is the get for /device/{deviceid}
func (config *Config) GetDevice(w http.ResponseWriter, r *http.Request) {
	deviceid,err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	device, err := GetDevice(config.Database, deviceid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	render.JSON(w, r, device)
}

// CreateDevice creates a device
func (config *Config) CreateDevice(w http.ResponseWriter, r *http.Request) {
	data := Device{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	device, err := CreateDevice(config.Database, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, device)
}

// DeleteDevice deletes the device
func (config *Config) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	deviceid,err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	if err = DeleteDevice(config.Database, deviceid); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
    w.Write([]byte("204 - No Content"))
}
 
func (config *Config) getTenantDevices(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	devices, err := GetTenantDevices(config.Database, tenantKey)
    if err != nil { 
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
    }

    render.JSON(w, r, devices)
}

// GetTenantDevice gets a tenant device
func (config *Config) GetTenantDevice(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	deviceid,err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	device, err := GetTenantDevice(config.Database, tenantKey, deviceid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	render.JSON(w, r, device)
}

// CreateTenantDevice creates a tenant device
func (config *Config) CreateTenantDevice(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")

	data := Device{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	device, err := CreateTenantDevice(config.Database, tenantKey, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, device)
}

// DeleteTenantDevice deletes a tenant device
func (config *Config) DeleteTenantDevice(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	deviceid,err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	err = DeleteTenantDevice(config.Database, tenantKey, deviceid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	w.WriteHeader(http.StatusNoContent)
    w.Write([]byte("204 - No Content"))
}
