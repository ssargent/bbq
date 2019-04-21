package devices

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/ssargent/go-bbq/config"
	"github.com/ssargent/go-bbq/internal/infrastructure"
)

// DeviceApi - duck typing?
type DeviceApi struct {
	config *config.Config
}

// New returns a new instance of a config
func New(configuration *config.Config) *DeviceApi {
	return &DeviceApi{configuration}
}

// Device is the representation of a bbq device
type Device struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenantid"`
}

// TenantRoutes are the routes scoped to a specific tenant
func (api *DeviceApi) TenantRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", api.getTenantDevices)
	router.Get("/{deviceid}", api.GetTenantDevice)
	router.Post("/", api.CreateTenantDevice)
	router.Delete("/{deviceid}", api.DeleteTenantDevice)

	return router
}

// GetAllDevices is the GET for /devices
func (api *DeviceApi) GetAllDevices(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	devices, err := GetAllDevices(api.config, count, start)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, devices)
}

// GetDevice is the get for /device/{deviceid}
func (api *DeviceApi) GetDevice(w http.ResponseWriter, r *http.Request) {
	deviceid, err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	device, err := GetDevice(api.config, deviceid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	render.JSON(w, r, device)
}

// CreateDevice creates a device
func (api *DeviceApi) CreateDevice(w http.ResponseWriter, r *http.Request) {
	data := Device{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	device, err := CreateDevice(api.config, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, device)
}

// DeleteDevice deletes the device
func (api *DeviceApi) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	deviceid, err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	if err = DeleteDevice(api.config, deviceid); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("204 - No Content"))
}

func (api *DeviceApi) getTenantDevices(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	devices, err := GetTenantDevices(api.config, tenantKey)
	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, devices)
}

// GetTenantDevice gets a tenant device
func (api *DeviceApi) GetTenantDevice(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	deviceid, err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	device, err := GetTenantDevice(api.config, tenantKey, deviceid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	render.JSON(w, r, device)
}

// CreateTenantDevice creates a tenant device
func (api *DeviceApi) CreateTenantDevice(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")

	data := Device{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	device, err := CreateTenantDevice(api.config, tenantKey, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, device)
}

// DeleteTenantDevice deletes a tenant device
func (api *DeviceApi) DeleteTenantDevice(w http.ResponseWriter, r *http.Request) {
	tenantKey := chi.URLParam(r, "tenantkey")
	deviceid, err := strconv.Atoi(chi.URLParam(r, "deviceid"))

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	err = DeleteTenantDevice(api.config, tenantKey, deviceid)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("204 - No Content"))
}
