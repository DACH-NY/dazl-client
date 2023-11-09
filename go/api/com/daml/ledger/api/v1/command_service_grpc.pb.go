// Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: com/daml/ledger/api/v1/command_service.proto

package v1

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

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	// Submits a single composite command and waits for its result.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWait(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Submits a single composite command, waits for its result, and returns the transaction id.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWaitForTransactionId(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*SubmitAndWaitForTransactionIdResponse, error)
	// Submits a single composite command, waits for its result, and returns the transaction.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWaitForTransaction(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*SubmitAndWaitForTransactionResponse, error)
	// Submits a single composite command, waits for its result, and returns the transaction tree.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWaitForTransactionTree(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*SubmitAndWaitForTransactionTreeResponse, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) SubmitAndWait(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.CommandService/SubmitAndWait", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) SubmitAndWaitForTransactionId(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*SubmitAndWaitForTransactionIdResponse, error) {
	out := new(SubmitAndWaitForTransactionIdResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.CommandService/SubmitAndWaitForTransactionId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) SubmitAndWaitForTransaction(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*SubmitAndWaitForTransactionResponse, error) {
	out := new(SubmitAndWaitForTransactionResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.CommandService/SubmitAndWaitForTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) SubmitAndWaitForTransactionTree(ctx context.Context, in *SubmitAndWaitRequest, opts ...grpc.CallOption) (*SubmitAndWaitForTransactionTreeResponse, error) {
	out := new(SubmitAndWaitForTransactionTreeResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.CommandService/SubmitAndWaitForTransactionTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility
type CommandServiceServer interface {
	// Submits a single composite command and waits for its result.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWait(context.Context, *SubmitAndWaitRequest) (*emptypb.Empty, error)
	// Submits a single composite command, waits for its result, and returns the transaction id.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWaitForTransactionId(context.Context, *SubmitAndWaitRequest) (*SubmitAndWaitForTransactionIdResponse, error)
	// Submits a single composite command, waits for its result, and returns the transaction.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWaitForTransaction(context.Context, *SubmitAndWaitRequest) (*SubmitAndWaitForTransactionResponse, error)
	// Submits a single composite command, waits for its result, and returns the transaction tree.
	// Propagates the gRPC error of failed submissions including Daml interpretation errors.
	SubmitAndWaitForTransactionTree(context.Context, *SubmitAndWaitRequest) (*SubmitAndWaitForTransactionTreeResponse, error)
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (UnimplementedCommandServiceServer) SubmitAndWait(context.Context, *SubmitAndWaitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAndWait not implemented")
}
func (UnimplementedCommandServiceServer) SubmitAndWaitForTransactionId(context.Context, *SubmitAndWaitRequest) (*SubmitAndWaitForTransactionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAndWaitForTransactionId not implemented")
}
func (UnimplementedCommandServiceServer) SubmitAndWaitForTransaction(context.Context, *SubmitAndWaitRequest) (*SubmitAndWaitForTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAndWaitForTransaction not implemented")
}
func (UnimplementedCommandServiceServer) SubmitAndWaitForTransactionTree(context.Context, *SubmitAndWaitRequest) (*SubmitAndWaitForTransactionTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAndWaitForTransactionTree not implemented")
}
func (UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}

// UnsafeCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceServer will
// result in compilation errors.
type UnsafeCommandServiceServer interface {
	mustEmbedUnimplementedCommandServiceServer()
}

func RegisterCommandServiceServer(s grpc.ServiceRegistrar, srv CommandServiceServer) {
	s.RegisterService(&CommandService_ServiceDesc, srv)
}

func _CommandService_SubmitAndWait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAndWaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).SubmitAndWait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.CommandService/SubmitAndWait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).SubmitAndWait(ctx, req.(*SubmitAndWaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_SubmitAndWaitForTransactionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAndWaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).SubmitAndWaitForTransactionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.CommandService/SubmitAndWaitForTransactionId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).SubmitAndWaitForTransactionId(ctx, req.(*SubmitAndWaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_SubmitAndWaitForTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAndWaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).SubmitAndWaitForTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.CommandService/SubmitAndWaitForTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).SubmitAndWaitForTransaction(ctx, req.(*SubmitAndWaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_SubmitAndWaitForTransactionTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAndWaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).SubmitAndWaitForTransactionTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.CommandService/SubmitAndWaitForTransactionTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).SubmitAndWaitForTransactionTree(ctx, req.(*SubmitAndWaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandService_ServiceDesc is the grpc.ServiceDesc for CommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitAndWait",
			Handler:    _CommandService_SubmitAndWait_Handler,
		},
		{
			MethodName: "SubmitAndWaitForTransactionId",
			Handler:    _CommandService_SubmitAndWaitForTransactionId_Handler,
		},
		{
			MethodName: "SubmitAndWaitForTransaction",
			Handler:    _CommandService_SubmitAndWaitForTransaction_Handler,
		},
		{
			MethodName: "SubmitAndWaitForTransactionTree",
			Handler:    _CommandService_SubmitAndWaitForTransactionTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v1/command_service.proto",
}
