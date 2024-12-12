// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: com/daml/ledger/api/v1/admin/command_inspection_service.proto

package admin

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
	CommandInspectionService_GetCommandStatus_FullMethodName = "/com.daml.ledger.api.v1.admin.CommandInspectionService/GetCommandStatus"
)

// CommandInspectionServiceClient is the client API for CommandInspectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandInspectionServiceClient interface {
	GetCommandStatus(ctx context.Context, in *GetCommandStatusRequest, opts ...grpc.CallOption) (*GetCommandStatusResponse, error)
}

type commandInspectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandInspectionServiceClient(cc grpc.ClientConnInterface) CommandInspectionServiceClient {
	return &commandInspectionServiceClient{cc}
}

func (c *commandInspectionServiceClient) GetCommandStatus(ctx context.Context, in *GetCommandStatusRequest, opts ...grpc.CallOption) (*GetCommandStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommandStatusResponse)
	err := c.cc.Invoke(ctx, CommandInspectionService_GetCommandStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandInspectionServiceServer is the server API for CommandInspectionService service.
// All implementations must embed UnimplementedCommandInspectionServiceServer
// for forward compatibility.
type CommandInspectionServiceServer interface {
	GetCommandStatus(context.Context, *GetCommandStatusRequest) (*GetCommandStatusResponse, error)
	mustEmbedUnimplementedCommandInspectionServiceServer()
}

// UnimplementedCommandInspectionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommandInspectionServiceServer struct{}

func (UnimplementedCommandInspectionServiceServer) GetCommandStatus(context.Context, *GetCommandStatusRequest) (*GetCommandStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommandStatus not implemented")
}
func (UnimplementedCommandInspectionServiceServer) mustEmbedUnimplementedCommandInspectionServiceServer() {
}
func (UnimplementedCommandInspectionServiceServer) testEmbeddedByValue() {}

// UnsafeCommandInspectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandInspectionServiceServer will
// result in compilation errors.
type UnsafeCommandInspectionServiceServer interface {
	mustEmbedUnimplementedCommandInspectionServiceServer()
}

func RegisterCommandInspectionServiceServer(s grpc.ServiceRegistrar, srv CommandInspectionServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommandInspectionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommandInspectionService_ServiceDesc, srv)
}

func _CommandInspectionService_GetCommandStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommandStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandInspectionServiceServer).GetCommandStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandInspectionService_GetCommandStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandInspectionServiceServer).GetCommandStatus(ctx, req.(*GetCommandStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandInspectionService_ServiceDesc is the grpc.ServiceDesc for CommandInspectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandInspectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.admin.CommandInspectionService",
	HandlerType: (*CommandInspectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommandStatus",
			Handler:    _CommandInspectionService_GetCommandStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v1/admin/command_inspection_service.proto",
}
