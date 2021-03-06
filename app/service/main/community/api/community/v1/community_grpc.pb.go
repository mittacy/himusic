// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package community

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

// CommunityServiceClient is the client API for CommunityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunityServiceClient interface {
	CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityReply, error)
	UpdateCommunity(ctx context.Context, in *UpdateCommunityRequest, opts ...grpc.CallOption) (*UpdateCommunityReply, error)
	DeleteCommunity(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityReply, error)
	GetCommunity(ctx context.Context, in *GetCommunityRequest, opts ...grpc.CallOption) (*GetCommunityReply, error)
	ListCommunity(ctx context.Context, in *ListCommunityRequest, opts ...grpc.CallOption) (*ListCommunityReply, error)
}

type communityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunityServiceClient(cc grpc.ClientConnInterface) CommunityServiceClient {
	return &communityServiceClient{cc}
}

func (c *communityServiceClient) CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityReply, error) {
	out := new(CreateCommunityReply)
	err := c.cc.Invoke(ctx, "/api.community.CommunityService/CreateCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) UpdateCommunity(ctx context.Context, in *UpdateCommunityRequest, opts ...grpc.CallOption) (*UpdateCommunityReply, error) {
	out := new(UpdateCommunityReply)
	err := c.cc.Invoke(ctx, "/api.community.CommunityService/UpdateCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) DeleteCommunity(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityReply, error) {
	out := new(DeleteCommunityReply)
	err := c.cc.Invoke(ctx, "/api.community.CommunityService/DeleteCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetCommunity(ctx context.Context, in *GetCommunityRequest, opts ...grpc.CallOption) (*GetCommunityReply, error) {
	out := new(GetCommunityReply)
	err := c.cc.Invoke(ctx, "/api.community.CommunityService/GetCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) ListCommunity(ctx context.Context, in *ListCommunityRequest, opts ...grpc.CallOption) (*ListCommunityReply, error) {
	out := new(ListCommunityReply)
	err := c.cc.Invoke(ctx, "/api.community.CommunityService/ListCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunityServiceServer is the server API for CommunityService service.
// All implementations must embed UnimplementedCommunityServiceServer
// for forward compatibility
type CommunityServiceServer interface {
	CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityReply, error)
	UpdateCommunity(context.Context, *UpdateCommunityRequest) (*UpdateCommunityReply, error)
	DeleteCommunity(context.Context, *DeleteCommunityRequest) (*DeleteCommunityReply, error)
	GetCommunity(context.Context, *GetCommunityRequest) (*GetCommunityReply, error)
	ListCommunity(context.Context, *ListCommunityRequest) (*ListCommunityReply, error)
	mustEmbedUnimplementedCommunityServiceServer()
}

// UnimplementedCommunityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunityServiceServer struct {
}

func (UnimplementedCommunityServiceServer) CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunity not implemented")
}
func (UnimplementedCommunityServiceServer) UpdateCommunity(context.Context, *UpdateCommunityRequest) (*UpdateCommunityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommunity not implemented")
}
func (UnimplementedCommunityServiceServer) DeleteCommunity(context.Context, *DeleteCommunityRequest) (*DeleteCommunityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommunity not implemented")
}
func (UnimplementedCommunityServiceServer) GetCommunity(context.Context, *GetCommunityRequest) (*GetCommunityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunity not implemented")
}
func (UnimplementedCommunityServiceServer) ListCommunity(context.Context, *ListCommunityRequest) (*ListCommunityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommunity not implemented")
}
func (UnimplementedCommunityServiceServer) mustEmbedUnimplementedCommunityServiceServer() {}

// UnsafeCommunityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunityServiceServer will
// result in compilation errors.
type UnsafeCommunityServiceServer interface {
	mustEmbedUnimplementedCommunityServiceServer()
}

func RegisterCommunityServiceServer(s grpc.ServiceRegistrar, srv CommunityServiceServer) {
	s.RegisterService(&CommunityService_ServiceDesc, srv)
}

func _CommunityService_CreateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.community.CommunityService/CreateCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreateCommunity(ctx, req.(*CreateCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_UpdateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).UpdateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.community.CommunityService/UpdateCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).UpdateCommunity(ctx, req.(*UpdateCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_DeleteCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).DeleteCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.community.CommunityService/DeleteCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).DeleteCommunity(ctx, req.(*DeleteCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.community.CommunityService/GetCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetCommunity(ctx, req.(*GetCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_ListCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).ListCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.community.CommunityService/ListCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).ListCommunity(ctx, req.(*ListCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunityService_ServiceDesc is the grpc.ServiceDesc for CommunityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.community.CommunityService",
	HandlerType: (*CommunityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunity",
			Handler:    _CommunityService_CreateCommunity_Handler,
		},
		{
			MethodName: "UpdateCommunity",
			Handler:    _CommunityService_UpdateCommunity_Handler,
		},
		{
			MethodName: "DeleteCommunity",
			Handler:    _CommunityService_DeleteCommunity_Handler,
		},
		{
			MethodName: "GetCommunity",
			Handler:    _CommunityService_GetCommunity_Handler,
		},
		{
			MethodName: "ListCommunity",
			Handler:    _CommunityService_ListCommunity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/community/v1/community.proto",
}
