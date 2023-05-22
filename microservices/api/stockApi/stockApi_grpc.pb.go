// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: microservices/api/stockApi/stockApi.proto

package stockApi

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
	StockService_AddProducts_FullMethodName   = "/stockApi.StockService/AddProducts"
	StockService_GetProducts_FullMethodName   = "/stockApi.StockService/GetProducts"
	StockService_RemoveProduct_FullMethodName = "/stockApi.StockService/RemoveProduct"
)

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	AddProducts(ctx context.Context, in *AddProductsRequest, opts ...grpc.CallOption) (*AddProductsReply, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsReply, error)
	RemoveProduct(ctx context.Context, in *RemoveProductRequest, opts ...grpc.CallOption) (*RemoveProductReply, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) AddProducts(ctx context.Context, in *AddProductsRequest, opts ...grpc.CallOption) (*AddProductsReply, error) {
	out := new(AddProductsReply)
	err := c.cc.Invoke(ctx, StockService_AddProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsReply, error) {
	out := new(GetProductsReply)
	err := c.cc.Invoke(ctx, StockService_GetProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) RemoveProduct(ctx context.Context, in *RemoveProductRequest, opts ...grpc.CallOption) (*RemoveProductReply, error) {
	out := new(RemoveProductReply)
	err := c.cc.Invoke(ctx, StockService_RemoveProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility
type StockServiceServer interface {
	AddProducts(context.Context, *AddProductsRequest) (*AddProductsReply, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsReply, error)
	RemoveProduct(context.Context, *RemoveProductRequest) (*RemoveProductReply, error)
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (UnimplementedStockServiceServer) AddProducts(context.Context, *AddProductsRequest) (*AddProductsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProducts not implemented")
}
func (UnimplementedStockServiceServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedStockServiceServer) RemoveProduct(context.Context, *RemoveProductRequest) (*RemoveProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProduct not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_AddProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).AddProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_AddProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).AddProducts(ctx, req.(*AddProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_RemoveProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).RemoveProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_RemoveProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).RemoveProduct(ctx, req.(*RemoveProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stockApi.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProducts",
			Handler:    _StockService_AddProducts_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _StockService_GetProducts_Handler,
		},
		{
			MethodName: "RemoveProduct",
			Handler:    _StockService_RemoveProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservices/api/stockApi/stockApi.proto",
}