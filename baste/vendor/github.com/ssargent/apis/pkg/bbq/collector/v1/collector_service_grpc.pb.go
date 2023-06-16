// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: bbq/collector/v1/collector_service.proto

package collectorv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CollectorService_Record_FullMethodName  = "/bbq.collector.v1.CollectorService/Record"
	CollectorService_Session_FullMethodName = "/bbq.collector.v1.CollectorService/Session"
)

// CollectorServiceClient is the client API for CollectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectorServiceClient interface {
	Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error)
	Session(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionResponse, error)
}

type collectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorServiceClient(cc grpc.ClientConnInterface) CollectorServiceClient {
	return &collectorServiceClient{cc}
}

func (c *collectorServiceClient) Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, CollectorService_Record_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorServiceClient) Session(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, CollectorService_Session_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServiceServer is the server API for CollectorService service.
// All implementations must embed UnimplementedCollectorServiceServer
// for forward compatibility
type CollectorServiceServer interface {
	Record(context.Context, *RecordRequest) (*RecordResponse, error)
	Session(context.Context, *SessionRequest) (*SessionResponse, error)
	mustEmbedUnimplementedCollectorServiceServer()
}

// UnimplementedCollectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollectorServiceServer struct {
}

func (UnimplementedCollectorServiceServer) Record(context.Context, *RecordRequest) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (UnimplementedCollectorServiceServer) Session(context.Context, *SessionRequest) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Session not implemented")
}
func (UnimplementedCollectorServiceServer) mustEmbedUnimplementedCollectorServiceServer() {}

// UnsafeCollectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectorServiceServer will
// result in compilation errors.
type UnsafeCollectorServiceServer interface {
	mustEmbedUnimplementedCollectorServiceServer()
}

func RegisterCollectorServiceServer(s grpc.ServiceRegistrar, srv CollectorServiceServer) {
	s.RegisterService(&CollectorService_ServiceDesc, srv)
}

func _CollectorService_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorService_Record_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).Record(ctx, req.(*RecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorService_Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorService_Session_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).Session(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectorService_ServiceDesc is the grpc.ServiceDesc for CollectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bbq.collector.v1.CollectorService",
	HandlerType: (*CollectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Record",
			Handler:    _CollectorService_Record_Handler,
		},
		{
			MethodName: "Session",
			Handler:    _CollectorService_Session_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bbq/collector/v1/collector_service.proto",
}