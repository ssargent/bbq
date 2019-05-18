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

	"github.com/ssargent/go-bbq/config"
	"github.com/ssargent/go-bbq/internal/apis/bbq/devices"
	"github.com/ssargent/go-bbq/internal/apis/bbq/monitors"
	"github.com/ssargent/go-bbq/internal/apis/bbq/sessions"
	"github.com/ssargent/go-bbq/internal/apis/data/temperature"
	"github.com/ssargent/go-bbq/internal/apis/health"
	"github.com/ssargent/go-bbq/internal/infrastructure/redis"

	"github.com/ssargent/go-bbq/bbq/device"
	//"github.com/ssargent/go-bbq/system"
	"github.com/ssargent/go-bbq/system/account"
	"github.com/ssargent/go-bbq/system/tenant"
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
		cors.Handler)

	healthAPI := health.New(c)
	devicesAPI := devices.New(c)
	monitorsAPI := monitors.New(c)
	sessionsAPI := sessions.New(c)
	temperatureAPI := temperature.New(c)

	caching := redis.NewRedisCacheService(c)
	accountRepository := account.NewAccountRepository(c)
	accountService := account.NewAccountService(caching, accountRepository)
	accountHandler := account.NewAccountHandler(c, accountService)

	tenantRepository := tenant.NewTenantRepository(c)
	tenantService := tenant.NewTenantService(c, tenantRepository)
	tenantHandler := tenant.NewTenantHandler(c, tenantService, accountService)

	deviceRepository := device.NewDeviceRepository(c.Database)
	deviceService := device.NewDeviceService(caching, deviceRepository)
	deviceHandler := device.NewDeviceHandler(c, deviceService)

	router.Route("/v1", func(r chi.Router) {
		//	r.Mount("/bbq/devices", devicesAPI.Routes())
		r.Mount("/health", healthAPI.HealthRoutes())

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)

			r.Mount("/{tenantkey}/bbq/devices", devicesAPI.TenantRoutes())
			r.Mount("/{tenantkey}/bbq/monitors", monitorsAPI.TenantRoutes())
			r.Mount("/{tenantkey}/bbq/sessions", sessionsAPI.TenantRoutes())
			r.Mount("/{tenantkey}/data/temperature", temperatureAPI.TenantRoutes())
			r.Mount("/bbq/devices", deviceHandler.Routes())
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
