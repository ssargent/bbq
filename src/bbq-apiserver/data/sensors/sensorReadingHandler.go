package sensors

import (
	"net/http"

	//	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	//"github.com/google/uuid"

	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/data"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
	"github.com/ssargent/bbq/bbq-apiserver/security"
)

type sensorReadingHandler struct {
	service        data.SensorReadingService
	authentication security.AuthenticationService
	config         *config.Config
}

// NewSensorReadingHandler will create an api Handler for a Sensor Readings.
func NewSensorReadingHandler(config *config.Config, authentication security.AuthenticationService, service data.SensorReadingService) infrastructure.ApiHandler {
	return &sensorReadingHandler{service: service, authentication: authentication, config: config}
}

func (handler *sensorReadingHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{sessionid}", handler.getSensorReadings)
	router.Get("/{sessionid}/raw", handler.getRawSensorReadings)

	return router
}

func (handler *sensorReadingHandler) getSensorReadings(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	tenantString := claims["tenant"].(string)
	tenant, err := uuid.Parse(tenantString)

	sessionidString := chi.URLParam(r, "sessionid")
	sessionid, err := uuid.Parse(sessionidString)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	sessionData, err := handler.service.GetReadings(tenant, sessionid)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, sessionData)
}

func (handler *sensorReadingHandler) getRawSensorReadings(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	tenantString := claims["tenant"].(string)
	tenant, err := uuid.Parse(tenantString)

	sessionidString := chi.URLParam(r, "sessionid")
	sessionid, err := uuid.Parse(sessionidString)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	sessionData, err := handler.service.GetRawReadings(tenant, sessionid)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, sessionData)
}
