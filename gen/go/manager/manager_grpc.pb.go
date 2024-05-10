// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: manager/manager.proto

package manager

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateLoginItem(ctx context.Context, in *CreateLoginItemRequest, opts ...grpc.CallOption) (*CreateLoginItemResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error)
	GetLoginItem(ctx context.Context, in *GetLoginItemRequest, opts ...grpc.CallOption) (*GetLoginItemResponse, error)
	GetLoginItems(ctx context.Context, in *GetLoginItemsRequest, opts ...grpc.CallOption) (*GetLoginItemsResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateLoginItem(ctx context.Context, in *CreateLoginItemRequest, opts ...grpc.CallOption) (*CreateLoginItemResponse, error) {
	out := new(CreateLoginItemResponse)
	err := c.cc.Invoke(ctx, "/manager.Manager/CreateLoginItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, "/manager.Manager/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, "/manager.Manager/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetLoginItem(ctx context.Context, in *GetLoginItemRequest, opts ...grpc.CallOption) (*GetLoginItemResponse, error) {
	out := new(GetLoginItemResponse)
	err := c.cc.Invoke(ctx, "/manager.Manager/GetLoginItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetLoginItems(ctx context.Context, in *GetLoginItemsRequest, opts ...grpc.CallOption) (*GetLoginItemsResponse, error) {
	out := new(GetLoginItemsResponse)
	err := c.cc.Invoke(ctx, "/manager.Manager/GetLoginItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateLoginItem(context.Context, *CreateLoginItemRequest) (*CreateLoginItemResponse, error)
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error)
	GetLoginItem(context.Context, *GetLoginItemRequest) (*GetLoginItemResponse, error)
	GetLoginItems(context.Context, *GetLoginItemsRequest) (*GetLoginItemsResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateLoginItem(context.Context, *CreateLoginItemRequest) (*CreateLoginItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLoginItem not implemented")
}
func (UnimplementedManagerServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedManagerServer) GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedManagerServer) GetLoginItem(context.Context, *GetLoginItemRequest) (*GetLoginItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginItem not implemented")
}
func (UnimplementedManagerServer) GetLoginItems(context.Context, *GetLoginItemsRequest) (*GetLoginItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginItems not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_CreateLoginItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLoginItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateLoginItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/CreateLoginItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateLoginItem(ctx, req.(*CreateLoginItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetItems(ctx, req.(*GetItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetLoginItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetLoginItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/GetLoginItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetLoginItem(ctx, req.(*GetLoginItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetLoginItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetLoginItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/GetLoginItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetLoginItems(ctx, req.(*GetLoginItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLoginItem",
			Handler:    _Manager_CreateLoginItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _Manager_GetItem_Handler,
		},
		{
			MethodName: "GetItems",
			Handler:    _Manager_GetItems_Handler,
		},
		{
			MethodName: "GetLoginItem",
			Handler:    _Manager_GetLoginItem_Handler,
		},
		{
			MethodName: "GetLoginItems",
			Handler:    _Manager_GetLoginItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager/manager.proto",
}
