package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/internal/apis/data/temperature"
	"github.com/ssargent/bbq/bbq-apiserver/internal/apis/health"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure/redis"
	"github.com/ssargent/bbq/bbq-apiserver/security/claims"

	"github.com/ssargent/bbq/bbq-apiserver/bbq/device"
	"github.com/ssargent/bbq/bbq-apiserver/bbq/monitor"
	"github.com/ssargent/bbq/bbq-apiserver/bbq/session"
	"github.com/ssargent/bbq/bbq-apiserver/bbq/subject"

	"github.com/ssargent/bbq/bbq-apiserver/data/sensors"

	//"github.com/ssargent/bbq/bbq-apiserver/system"
	"github.com/ssargent/bbq/bbq-apiserver/system/account"
	"github.com/ssargent/bbq/bbq-apiserver/system/tenant"
)

// Routes wtse-1
func Routes(c *config.Config) *chi.Mux {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300})

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.Recoverer,
		middleware.RequestID,
		cors.Handler)

	healthAPI := health.New(c)
	temperatureAPI := temperature.New(c)

	caching := redis.NewRedisCacheService(c)

	authentication := claims.NewClaimsAuthenticationService()

	sensorReadingRepository := sensors.NewSensorReadingRepository(c.Database)
	sensorReadingService := sensors.NewSensorReadingService(caching, sensorReadingRepository)
	sensorReadingHandler := sensors.NewSensorReadingHandler(c, authentication, sensorReadingService)

	deviceRepository := device.NewDeviceRepository(c.Database)
	deviceService := device.NewDeviceService(caching, deviceRepository)
	deviceHandler := device.NewDeviceHandler(c, authentication, deviceService)

	monitorRepository := monitor.NewMonitorRepository(c.Database)
	monitorService := monitor.NewMonitorService(caching, monitorRepository)
	monitorHandler := monitor.NewMonitorHandler(c, authentication, monitorService)

	accountRepository := account.NewAccountRepository(c)
	accountService := account.NewAccountService(caching, accountRepository)
	accountHandler := account.NewAccountHandler(c, accountService)

	tenantRepository := tenant.NewTenantRepository(c)
	tenantService := tenant.NewTenantService(c, tenantRepository)
	tenantHandler := tenant.NewTenantHandler(c, tenantService, accountService)

	subjectRepository := subject.NewSubjectRepository(c.Database)
	subjectService := subject.NewSubjectService(caching, subjectRepository)
	subjectHandler := subject.NewSubjectHandler(c, authentication, subjectService)

	sessionRepository := session.NewSessionRepository(c.Database)

	unitOfWork := bbq.BBQUnitOfWork{
		Monitor: monitorRepository,
		Device:  deviceRepository,
		Session: sessionRepository,
		Subject: subjectRepository,
	}

	sessionService := session.NewSessionService(caching, unitOfWork, deviceService, monitorService, subjectService)
	sessionHandler := session.NewSessionHandler(c, authentication, sessionService)

	router.Route("/v1", func(r chi.Router) {
		//	r.Mount("/bbq/devices", devicesAPI.Routes())
		r.Mount("/health", healthAPI.HealthRoutes())

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)

			r.Mount("/bbq/devices", deviceHandler.Routes())
			r.Mount("/bbq/monitors", monitorHandler.Routes())
			r.Mount("/bbq/sessions", sessionHandler.Routes())
			r.Mount("/bbq/subjects", subjectHandler.Routes())
			r.Mount("/{tenantkey}/data/temperature", temperatureAPI.TenantRoutes())

			r.Mount("/data/sensors", sensorReadingHandler.Routes())
		})

		r.Mount("/system/accounts", accountHandler.Routes())
		r.Mount("/system/tenants", tenantHandler.Routes())
	})

	return router
}

func main() {

	fmt.Println("Starting BBQ Server")
	configuration := &config.Config{}
	configuration.Initialize(
		os.Getenv("BBQ_DB_USER"),
		os.Getenv("BBQ_DB_PASSWORD"),
		os.Getenv("BBQ_DB_NAME"),
		os.Getenv("BBQ_DB_HOST"),
		os.Getenv("BBQ_REDIS_MASTER"),
		os.Getenv("BBQ_REDIS_PASSWORD"),
	)

	router := Routes(configuration)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Println("Serving application at PORT :" + configuration.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router)) // Note, the port is usually gotten from the environment.

}
