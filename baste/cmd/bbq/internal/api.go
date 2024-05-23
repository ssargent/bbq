package internal

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
	"github.com/ssargent/bbq/cmd/bbq/internal/collector"
	"github.com/ssargent/bbq/cmd/bbq/internal/intake"
	"github.com/ssargent/bbq/internal/config"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	customFunc grpc_zap.CodeToLevel
)

type API struct {
	cfg    *config.Config
	cache  *cache.Cache
	DB     *pgxpool.Pool
	logger *zap.Logger
}

func NewApi(logger *zap.Logger, cfg *config.Config, cache *cache.Cache, db *pgxpool.Pool) *API {
	// setup any services here...  Add those to the API Struct...
	return &API{
		logger: logger,
		cfg:    cfg,
		cache:  cache,
		DB:     db,
	}
}

func (a *API) ListenAndServe() error {
	errorChannel := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go a.rest(errorChannel, &wg)
	go a.grpc(errorChannel, &wg)

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-errorChannel:
		close(errorChannel)
		log.Fatal(err)
	}

	return nil
}

func (a *API) grpc(errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.cfg.Grpc.Port))
	if err != nil {
		errors <- err
	}

	opts := []grpc_zap.Option{}
	/*opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}*/

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(a.logger)
	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(a.logger, opts...),
	),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(a.logger, opts...),
		))
	intake.RegisterIntake(grpcServer, a.cfg, a.cache, a.DB, a.logger)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		errors <- err
	}
}

func (a *API) rest(errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	grpcPaths := make([]string, 0)

	if a.cfg.Services.CollectorEnabled {
		collectorPath, collectorHandler, err := collector.NewCollectorServiceHandler()
		if err != nil {
			errors <- fmt.Errorf("collector.NewCollectorServiceHandler: %w", err)
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
	errors <- srv.ListenAndServe()
}
