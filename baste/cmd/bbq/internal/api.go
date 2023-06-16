package internal

import (
	"fmt"
	"net/http"
	"strings"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/ssargent/bbq/cmd/bbq/internal/collector"
	"github.com/ssargent/bbq/cmd/bbq/internal/config"
)

type API struct {
	cfg   *config.Config
	cache *cache.Cache
	DB    *sqlx.DB
}

func NewApi(cfg *config.Config, cache *cache.Cache, db *sqlx.DB) *API {
	// setup any services here...  Add those to the API Struct...
	return &API{
		cfg:   cfg,
		cache: cache,
		DB:    db,
	}
}

func (a *API) ListenAndServe() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	grpcPaths := make([]string, 0)

	if a.cfg.Services.CollectorEnabled {
		collectorPath, collectorHandler, err := collector.NewCollectorServiceHandler()
		if err != nil {
			return fmt.Errorf("collector.NewCollectorServiceHandler: %w", err)
		}

		fmt.Printf("collectorPath: %s\n", collectorPath)
		r.Mount(collectorPath, collectorHandler)

		// Register the service for grpc reflection.  If reflection is enabled it will use this
		// and wire it up below.
		grpcPaths = append(grpcPaths, strings.Replace(collectorPath, "/", "", -1))
	}

	reflector := grpcreflect.NewStaticReflector(grpcPaths...)

	r1path, r1handler := grpcreflect.NewHandlerV1(reflector)
	_, r2handler := grpcreflect.NewHandlerV1Alpha(reflector)

	r.Handle(r1path, r1handler)
	r.Handle("/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo", r2handler)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	h2s := &http2.Server{}
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", a.cfg.Port),
		Handler: h2c.NewHandler(r, h2s),
	}
	return srv.ListenAndServe()
}
