// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: observer.proto

package observerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Observer_UpdateObserver_FullMethodName = "/observerpb.Observer/UpdateObserver"
)

// ObserverClient is the client API for Observer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObserverClient interface {
	UpdateObserver(ctx context.Context, in *UpdateObserverRequest, opts ...grpc.CallOption) (*UpdateObserverResponse, error)
}

type observerClient struct {
	cc grpc.ClientConnInterface
}

func NewObserverClient(cc grpc.ClientConnInterface) ObserverClient {
	return &observerClient{cc}
}

func (c *observerClient) UpdateObserver(ctx context.Context, in *UpdateObserverRequest, opts ...grpc.CallOption) (*UpdateObserverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateObserverResponse)
	err := c.cc.Invoke(ctx, Observer_UpdateObserver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObserverServer is the server API for Observer service.
// All implementations must embed UnimplementedObserverServer
// for forward compatibility.
type ObserverServer interface {
	UpdateObserver(context.Context, *UpdateObserverRequest) (*UpdateObserverResponse, error)
	mustEmbedUnimplementedObserverServer()
}

// UnimplementedObserverServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedObserverServer struct{}

func (UnimplementedObserverServer) UpdateObserver(context.Context, *UpdateObserverRequest) (*UpdateObserverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObserver not implemented")
}
func (UnimplementedObserverServer) mustEmbedUnimplementedObserverServer() {}
func (UnimplementedObserverServer) testEmbeddedByValue()                  {}

// UnsafeObserverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObserverServer will
// result in compilation errors.
type UnsafeObserverServer interface {
	mustEmbedUnimplementedObserverServer()
}

func RegisterObserverServer(s grpc.ServiceRegistrar, srv ObserverServer) {
	// If the following call pancis, it indicates UnimplementedObserverServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Observer_ServiceDesc, srv)
}

func _Observer_UpdateObserver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObserverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObserverServer).UpdateObserver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Observer_UpdateObserver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObserverServer).UpdateObserver(ctx, req.(*UpdateObserverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Observer_ServiceDesc is the grpc.ServiceDesc for Observer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Observer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "observerpb.Observer",
	HandlerType: (*ObserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateObserver",
			Handler:    _Observer_UpdateObserver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "observer.proto",
}
