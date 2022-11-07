// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: todo.proto

package model

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

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodosClient interface {
	GetById(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoList, error)
	Register(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error)
	Remove(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error)
	Edit(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) GetById(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error) {
	out := new(GetByIDResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := c.cc.Invoke(ctx, "/model.Todos/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Register(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Remove(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Edit(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServer is the server API for Todos service.
// All implementations must embed UnimplementedTodosServer
// for forward compatibility
type TodosServer interface {
	GetById(context.Context, *GetByIDRequest) (*GetByIDResponse, error)
	List(context.Context, *emptypb.Empty) (*TodoList, error)
	Register(context.Context, *Todo) (*MutationResponse, error)
	Remove(context.Context, *Todo) (*MutationResponse, error)
	Edit(context.Context, *Todo) (*MutationResponse, error)
	mustEmbedUnimplementedTodosServer()
}

// UnimplementedTodosServer must be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (UnimplementedTodosServer) GetById(context.Context, *GetByIDRequest) (*GetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedTodosServer) List(context.Context, *emptypb.Empty) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTodosServer) Register(context.Context, *Todo) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedTodosServer) Remove(context.Context, *Todo) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedTodosServer) Edit(context.Context, *Todo) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedTodosServer) mustEmbedUnimplementedTodosServer() {}

// UnsafeTodosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodosServer will
// result in compilation errors.
type UnsafeTodosServer interface {
	mustEmbedUnimplementedTodosServer()
}

func RegisterTodosServer(s grpc.ServiceRegistrar, srv TodosServer) {
	s.RegisterService(&Todos_ServiceDesc, srv)
}

func _Todos_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).GetById(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Register(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Remove(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Edit(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

// Todos_ServiceDesc is the grpc.ServiceDesc for Todos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _Todos_GetById_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todos_List_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Todos_Register_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Todos_Remove_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _Todos_Edit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
