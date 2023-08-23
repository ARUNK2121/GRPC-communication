// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// A_ServiceClient is the client API for A_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type A_ServiceClient interface {
	FunctionA(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type a_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewA_ServiceClient(cc grpc.ClientConnInterface) A_ServiceClient {
	return &a_ServiceClient{cc}
}

func (c *a_ServiceClient) FunctionA(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/service_a.A_Service/FunctionA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// A_ServiceServer is the server API for A_Service service.
// All implementations must embed UnimplementedA_ServiceServer
// for forward compatibility
type A_ServiceServer interface {
	FunctionA(context.Context, *SampleRequest) (*SampleResponse, error)
	mustEmbedUnimplementedA_ServiceServer()
}

// UnimplementedA_ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedA_ServiceServer struct {
}

func (UnimplementedA_ServiceServer) FunctionA(context.Context, *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FunctionA not implemented")
}
func (UnimplementedA_ServiceServer) mustEmbedUnimplementedA_ServiceServer() {}

// UnsafeA_ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to A_ServiceServer will
// result in compilation errors.
type UnsafeA_ServiceServer interface {
	mustEmbedUnimplementedA_ServiceServer()
}

func RegisterA_ServiceServer(s grpc.ServiceRegistrar, srv A_ServiceServer) {
	s.RegisterService(&A_Service_ServiceDesc, srv)
}

func _A_Service_FunctionA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(A_ServiceServer).FunctionA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_a.A_Service/FunctionA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(A_ServiceServer).FunctionA(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// A_Service_ServiceDesc is the grpc.ServiceDesc for A_Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var A_Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_a.A_Service",
	HandlerType: (*A_ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FunctionA",
			Handler:    _A_Service_FunctionA_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/service-A.proto",
}
