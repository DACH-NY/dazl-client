// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: com/daml/ledger/api/v2/command_submission_service.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CommandSubmissionService_Submit_FullMethodName             = "/com.daml.ledger.api.v2.CommandSubmissionService/Submit"
	CommandSubmissionService_SubmitReassignment_FullMethodName = "/com.daml.ledger.api.v2.CommandSubmissionService/SubmitReassignment"
)

// CommandSubmissionServiceClient is the client API for CommandSubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandSubmissionServiceClient interface {
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	SubmitReassignment(ctx context.Context, in *SubmitReassignmentRequest, opts ...grpc.CallOption) (*SubmitReassignmentResponse, error)
}

type commandSubmissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandSubmissionServiceClient(cc grpc.ClientConnInterface) CommandSubmissionServiceClient {
	return &commandSubmissionServiceClient{cc}
}

func (c *commandSubmissionServiceClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, CommandSubmissionService_Submit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandSubmissionServiceClient) SubmitReassignment(ctx context.Context, in *SubmitReassignmentRequest, opts ...grpc.CallOption) (*SubmitReassignmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitReassignmentResponse)
	err := c.cc.Invoke(ctx, CommandSubmissionService_SubmitReassignment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandSubmissionServiceServer is the server API for CommandSubmissionService service.
// All implementations must embed UnimplementedCommandSubmissionServiceServer
// for forward compatibility
type CommandSubmissionServiceServer interface {
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	SubmitReassignment(context.Context, *SubmitReassignmentRequest) (*SubmitReassignmentResponse, error)
	mustEmbedUnimplementedCommandSubmissionServiceServer()
}

// UnimplementedCommandSubmissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandSubmissionServiceServer struct {
}

func (UnimplementedCommandSubmissionServiceServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedCommandSubmissionServiceServer) SubmitReassignment(context.Context, *SubmitReassignmentRequest) (*SubmitReassignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitReassignment not implemented")
}
func (UnimplementedCommandSubmissionServiceServer) mustEmbedUnimplementedCommandSubmissionServiceServer() {
}

// UnsafeCommandSubmissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandSubmissionServiceServer will
// result in compilation errors.
type UnsafeCommandSubmissionServiceServer interface {
	mustEmbedUnimplementedCommandSubmissionServiceServer()
}

func RegisterCommandSubmissionServiceServer(s grpc.ServiceRegistrar, srv CommandSubmissionServiceServer) {
	s.RegisterService(&CommandSubmissionService_ServiceDesc, srv)
}

func _CommandSubmissionService_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandSubmissionServiceServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandSubmissionService_Submit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandSubmissionServiceServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandSubmissionService_SubmitReassignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitReassignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandSubmissionServiceServer).SubmitReassignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandSubmissionService_SubmitReassignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandSubmissionServiceServer).SubmitReassignment(ctx, req.(*SubmitReassignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandSubmissionService_ServiceDesc is the grpc.ServiceDesc for CommandSubmissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandSubmissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v2.CommandSubmissionService",
	HandlerType: (*CommandSubmissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _CommandSubmissionService_Submit_Handler,
		},
		{
			MethodName: "SubmitReassignment",
			Handler:    _CommandSubmissionService_SubmitReassignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v2/command_submission_service.proto",
}
