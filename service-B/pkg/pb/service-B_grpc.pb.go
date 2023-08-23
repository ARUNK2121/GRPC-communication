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

// B_ServiceClient is the client API for B_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type B_ServiceClient interface {
	FunctionB(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
	FunctionB_TO_A(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type b_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewB_ServiceClient(cc grpc.ClientConnInterface) B_ServiceClient {
	return &b_ServiceClient{cc}
}

func (c *b_ServiceClient) FunctionB(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/service_b.B_Service/FunctionB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b_ServiceClient) FunctionB_TO_A(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/service_b.B_Service/FunctionB_TO_A", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// B_ServiceServer is the server API for B_Service service.
// All implementations must embed UnimplementedB_ServiceServer
// for forward compatibility
type B_ServiceServer interface {
	FunctionB(context.Context, *SampleRequest) (*SampleResponse, error)
	FunctionB_TO_A(context.Context, *SampleRequest) (*SampleResponse, error)
	mustEmbedUnimplementedB_ServiceServer()
}

// UnimplementedB_ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedB_ServiceServer struct {
}

func (UnimplementedB_ServiceServer) FunctionB(context.Context, *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FunctionB not implemented")
}
func (UnimplementedB_ServiceServer) FunctionB_TO_A(context.Context, *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FunctionB_TO_A not implemented")
}
func (UnimplementedB_ServiceServer) mustEmbedUnimplementedB_ServiceServer() {}

// UnsafeB_ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to B_ServiceServer will
// result in compilation errors.
type UnsafeB_ServiceServer interface {
	mustEmbedUnimplementedB_ServiceServer()
}

func RegisterB_ServiceServer(s grpc.ServiceRegistrar, srv B_ServiceServer) {
	s.RegisterService(&B_Service_ServiceDesc, srv)
}

func _B_Service_FunctionB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B_ServiceServer).FunctionB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_b.B_Service/FunctionB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B_ServiceServer).FunctionB(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B_Service_FunctionB_TO_A_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B_ServiceServer).FunctionB_TO_A(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_b.B_Service/FunctionB_TO_A",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B_ServiceServer).FunctionB_TO_A(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// B_Service_ServiceDesc is the grpc.ServiceDesc for B_Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var B_Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_b.B_Service",
	HandlerType: (*B_ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FunctionB",
			Handler:    _B_Service_FunctionB_Handler,
		},
		{
			MethodName: "FunctionB_TO_A",
			Handler:    _B_Service_FunctionB_TO_A_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/service-B/pb/service-B.proto",
}