// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: bbq/collector/v1/collector_service.proto

package collectorv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ssargent/apis/pkg/bbq/collector/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CollectorServiceName is the fully-qualified name of the CollectorService service.
	CollectorServiceName = "bbq.collector.v1.CollectorService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CollectorServiceRecordProcedure is the fully-qualified name of the CollectorService's Record RPC.
	CollectorServiceRecordProcedure = "/bbq.collector.v1.CollectorService/Record"
	// CollectorServiceSessionProcedure is the fully-qualified name of the CollectorService's Session
	// RPC.
	CollectorServiceSessionProcedure = "/bbq.collector.v1.CollectorService/Session"
)

// CollectorServiceClient is a client for the bbq.collector.v1.CollectorService service.
type CollectorServiceClient interface {
	Record(context.Context, *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error)
	Session(context.Context, *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error)
}

// NewCollectorServiceClient constructs a client for the bbq.collector.v1.CollectorService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCollectorServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CollectorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &collectorServiceClient{
		record: connect_go.NewClient[v1.RecordRequest, v1.RecordResponse](
			httpClient,
			baseURL+CollectorServiceRecordProcedure,
			opts...,
		),
		session: connect_go.NewClient[v1.SessionRequest, v1.SessionResponse](
			httpClient,
			baseURL+CollectorServiceSessionProcedure,
			opts...,
		),
	}
}

// collectorServiceClient implements CollectorServiceClient.
type collectorServiceClient struct {
	record  *connect_go.Client[v1.RecordRequest, v1.RecordResponse]
	session *connect_go.Client[v1.SessionRequest, v1.SessionResponse]
}

// Record calls bbq.collector.v1.CollectorService.Record.
func (c *collectorServiceClient) Record(ctx context.Context, req *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error) {
	return c.record.CallUnary(ctx, req)
}

// Session calls bbq.collector.v1.CollectorService.Session.
func (c *collectorServiceClient) Session(ctx context.Context, req *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error) {
	return c.session.CallUnary(ctx, req)
}

// CollectorServiceHandler is an implementation of the bbq.collector.v1.CollectorService service.
type CollectorServiceHandler interface {
	Record(context.Context, *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error)
	Session(context.Context, *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error)
}

// NewCollectorServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCollectorServiceHandler(svc CollectorServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	collectorServiceRecordHandler := connect_go.NewUnaryHandler(
		CollectorServiceRecordProcedure,
		svc.Record,
		opts...,
	)
	collectorServiceSessionHandler := connect_go.NewUnaryHandler(
		CollectorServiceSessionProcedure,
		svc.Session,
		opts...,
	)
	return "/bbq.collector.v1.CollectorService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CollectorServiceRecordProcedure:
			collectorServiceRecordHandler.ServeHTTP(w, r)
		case CollectorServiceSessionProcedure:
			collectorServiceSessionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCollectorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCollectorServiceHandler struct{}

func (UnimplementedCollectorServiceHandler) Record(context.Context, *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bbq.collector.v1.CollectorService.Record is not implemented"))
}

func (UnimplementedCollectorServiceHandler) Session(context.Context, *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bbq.collector.v1.CollectorService.Session is not implemented"))
}
