// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: com/digitalasset/canton/participant/admin/v0/transfer_service.proto

package v0

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

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferServiceClient interface {
	TransferOut(ctx context.Context, in *AdminTransferOutRequest, opts ...grpc.CallOption) (*AdminTransferOutResponse, error)
	TransferIn(ctx context.Context, in *AdminTransferInRequest, opts ...grpc.CallOption) (*AdminTransferInResponse, error)
	TransferSearch(ctx context.Context, in *AdminTransferSearchQuery, opts ...grpc.CallOption) (*AdminTransferSearchResponse, error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) TransferOut(ctx context.Context, in *AdminTransferOutRequest, opts ...grpc.CallOption) (*AdminTransferOutResponse, error) {
	out := new(AdminTransferOutResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.TransferService/TransferOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) TransferIn(ctx context.Context, in *AdminTransferInRequest, opts ...grpc.CallOption) (*AdminTransferInResponse, error) {
	out := new(AdminTransferInResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.TransferService/TransferIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) TransferSearch(ctx context.Context, in *AdminTransferSearchQuery, opts ...grpc.CallOption) (*AdminTransferSearchResponse, error) {
	out := new(AdminTransferSearchResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.TransferService/TransferSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServiceServer is the server API for TransferService service.
// All implementations must embed UnimplementedTransferServiceServer
// for forward compatibility
type TransferServiceServer interface {
	TransferOut(context.Context, *AdminTransferOutRequest) (*AdminTransferOutResponse, error)
	TransferIn(context.Context, *AdminTransferInRequest) (*AdminTransferInResponse, error)
	TransferSearch(context.Context, *AdminTransferSearchQuery) (*AdminTransferSearchResponse, error)
	mustEmbedUnimplementedTransferServiceServer()
}

// UnimplementedTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransferServiceServer struct {
}

func (UnimplementedTransferServiceServer) TransferOut(context.Context, *AdminTransferOutRequest) (*AdminTransferOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOut not implemented")
}
func (UnimplementedTransferServiceServer) TransferIn(context.Context, *AdminTransferInRequest) (*AdminTransferInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferIn not implemented")
}
func (UnimplementedTransferServiceServer) TransferSearch(context.Context, *AdminTransferSearchQuery) (*AdminTransferSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferSearch not implemented")
}
func (UnimplementedTransferServiceServer) mustEmbedUnimplementedTransferServiceServer() {}

// UnsafeTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServiceServer will
// result in compilation errors.
type UnsafeTransferServiceServer interface {
	mustEmbedUnimplementedTransferServiceServer()
}

func RegisterTransferServiceServer(s grpc.ServiceRegistrar, srv TransferServiceServer) {
	s.RegisterService(&TransferService_ServiceDesc, srv)
}

func _TransferService_TransferOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminTransferOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).TransferOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.TransferService/TransferOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).TransferOut(ctx, req.(*AdminTransferOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_TransferIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminTransferInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).TransferIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.TransferService/TransferIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).TransferIn(ctx, req.(*AdminTransferInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_TransferSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminTransferSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).TransferSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.TransferService/TransferSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).TransferSearch(ctx, req.(*AdminTransferSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// TransferService_ServiceDesc is the grpc.ServiceDesc for TransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.canton.participant.admin.v0.TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferOut",
			Handler:    _TransferService_TransferOut_Handler,
		},
		{
			MethodName: "TransferIn",
			Handler:    _TransferService_TransferIn_Handler,
		},
		{
			MethodName: "TransferSearch",
			Handler:    _TransferService_TransferSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/canton/participant/admin/v0/transfer_service.proto",
}
