package device

import (
	"encoding/json"
	"net/http"

	//	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	//"github.com/google/uuid"

	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
	"github.com/ssargent/bbq/bbq-apiserver/security"
)

type deviceHandler struct {
	service        bbq.DeviceService
	authentication security.AuthenticationService
	config         *config.Config
}

// NewDeviceHandler will create an api Handler for a devices.
func NewDeviceHandler(config *config.Config, authentication security.AuthenticationService, service bbq.DeviceService) infrastructure.ApiHandler {
	return &deviceHandler{service: service, authentication: authentication, config: config}
}

func (handler *deviceHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.getDevices)
	router.Get("/{deviceName}", handler.getDevice)
	router.Post("/", handler.createDevice)
	//router.Delete("/{deviceid}", handler.deleteDevice)

	return router
}

func (handler *deviceHandler) getDevices(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	tenantString := claims["tenant"].(string)
	tenant, err := uuid.Parse(tenantString)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	devices, err := handler.service.GetDevices(tenant)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, devices)
}

func (handler *deviceHandler) getDevice(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	tenantString := claims["tenant"].(string)
	deviceName := chi.URLParam(r, "deviceName")

	tenant, err := uuid.Parse(tenantString)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	device, err := handler.service.GetDeviceByName(tenant, deviceName)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, device)
}

func (handler *deviceHandler) createDevice(w http.ResponseWriter, r *http.Request) {
	loginSession, err := handler.authentication.GetLoginSession(r)

	data := bbq.Device{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	device, err := handler.service.CreateDevice(loginSession.TenantId, data)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, device)
}
