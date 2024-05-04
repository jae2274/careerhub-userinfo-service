// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: careerhub/userinfo_service/restapi/restapi_grpc/scrapJob.proto

package restapi_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ScrapJobGrpcClient is the client API for ScrapJobGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScrapJobGrpcClient interface {
	GetScrapJobs(ctx context.Context, in *GetScrapJobsRequest, opts ...grpc.CallOption) (*GetScrapJobsResponse, error)
	AddScrapJob(ctx context.Context, in *AddScrapJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveScrapJob(ctx context.Context, in *RemoveScrapJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type scrapJobGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewScrapJobGrpcClient(cc grpc.ClientConnInterface) ScrapJobGrpcClient {
	return &scrapJobGrpcClient{cc}
}

func (c *scrapJobGrpcClient) GetScrapJobs(ctx context.Context, in *GetScrapJobsRequest, opts ...grpc.CallOption) (*GetScrapJobsResponse, error) {
	out := new(GetScrapJobsResponse)
	err := c.cc.Invoke(ctx, "/careerhub.userinfo_service.restapi_grpc.ScrapJobGrpc/GetScrapJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapJobGrpcClient) AddScrapJob(ctx context.Context, in *AddScrapJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/careerhub.userinfo_service.restapi_grpc.ScrapJobGrpc/AddScrapJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapJobGrpcClient) RemoveScrapJob(ctx context.Context, in *RemoveScrapJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/careerhub.userinfo_service.restapi_grpc.ScrapJobGrpc/RemoveScrapJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScrapJobGrpcServer is the server API for ScrapJobGrpc service.
// All implementations must embed UnimplementedScrapJobGrpcServer
// for forward compatibility
type ScrapJobGrpcServer interface {
	GetScrapJobs(context.Context, *GetScrapJobsRequest) (*GetScrapJobsResponse, error)
	AddScrapJob(context.Context, *AddScrapJobRequest) (*emptypb.Empty, error)
	RemoveScrapJob(context.Context, *RemoveScrapJobRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedScrapJobGrpcServer()
}

// UnimplementedScrapJobGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedScrapJobGrpcServer struct {
}

func (UnimplementedScrapJobGrpcServer) GetScrapJobs(context.Context, *GetScrapJobsRequest) (*GetScrapJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScrapJobs not implemented")
}
func (UnimplementedScrapJobGrpcServer) AddScrapJob(context.Context, *AddScrapJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddScrapJob not implemented")
}
func (UnimplementedScrapJobGrpcServer) RemoveScrapJob(context.Context, *RemoveScrapJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveScrapJob not implemented")
}
func (UnimplementedScrapJobGrpcServer) mustEmbedUnimplementedScrapJobGrpcServer() {}

// UnsafeScrapJobGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScrapJobGrpcServer will
// result in compilation errors.
type UnsafeScrapJobGrpcServer interface {
	mustEmbedUnimplementedScrapJobGrpcServer()
}

func RegisterScrapJobGrpcServer(s grpc.ServiceRegistrar, srv ScrapJobGrpcServer) {
	s.RegisterService(&ScrapJobGrpc_ServiceDesc, srv)
}

func _ScrapJobGrpc_GetScrapJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScrapJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapJobGrpcServer).GetScrapJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.userinfo_service.restapi_grpc.ScrapJobGrpc/GetScrapJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapJobGrpcServer).GetScrapJobs(ctx, req.(*GetScrapJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapJobGrpc_AddScrapJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddScrapJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapJobGrpcServer).AddScrapJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.userinfo_service.restapi_grpc.ScrapJobGrpc/AddScrapJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapJobGrpcServer).AddScrapJob(ctx, req.(*AddScrapJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapJobGrpc_RemoveScrapJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveScrapJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapJobGrpcServer).RemoveScrapJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.userinfo_service.restapi_grpc.ScrapJobGrpc/RemoveScrapJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapJobGrpcServer).RemoveScrapJob(ctx, req.(*RemoveScrapJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScrapJobGrpc_ServiceDesc is the grpc.ServiceDesc for ScrapJobGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScrapJobGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "careerhub.userinfo_service.restapi_grpc.ScrapJobGrpc",
	HandlerType: (*ScrapJobGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScrapJobs",
			Handler:    _ScrapJobGrpc_GetScrapJobs_Handler,
		},
		{
			MethodName: "AddScrapJob",
			Handler:    _ScrapJobGrpc_AddScrapJob_Handler,
		},
		{
			MethodName: "RemoveScrapJob",
			Handler:    _ScrapJobGrpc_RemoveScrapJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "careerhub/userinfo_service/restapi/restapi_grpc/scrapJob.proto",
}
