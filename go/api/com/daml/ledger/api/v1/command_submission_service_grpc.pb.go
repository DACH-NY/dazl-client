// Copyright (c) 2017-2022 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.5
// source: com/daml/ledger/api/v1/command_submission_service.proto

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

// CommandSubmissionServiceClient is the client API for CommandSubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandSubmissionServiceClient interface {
	// Submit a single composite command.
	// Errors:
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “NOT_FOUND“: if the request does not include a valid ledger id or if a resource is missing (e.g. contract key)
	// due to for example contention on resources
	// - “ALREADY_EXISTS“ if a resource is duplicated (e.g. contract key)
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	// - “ABORTED“: if the number of in-flight commands reached the maximum (if a limit is configured)
	// - “FAILED_PRECONDITION“: on consistency errors (e.g. the contract key has changed since the submission)
	// or if an interpretation error occurred
	// - “UNAVAILABLE“: if the participant is not yet ready to submit commands or if the service has been shut down.
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type commandSubmissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandSubmissionServiceClient(cc grpc.ClientConnInterface) CommandSubmissionServiceClient {
	return &commandSubmissionServiceClient{cc}
}

func (c *commandSubmissionServiceClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.CommandSubmissionService/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandSubmissionServiceServer is the server API for CommandSubmissionService service.
// All implementations must embed UnimplementedCommandSubmissionServiceServer
// for forward compatibility
type CommandSubmissionServiceServer interface {
	// Submit a single composite command.
	// Errors:
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “NOT_FOUND“: if the request does not include a valid ledger id or if a resource is missing (e.g. contract key)
	// due to for example contention on resources
	// - “ALREADY_EXISTS“ if a resource is duplicated (e.g. contract key)
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	// - “ABORTED“: if the number of in-flight commands reached the maximum (if a limit is configured)
	// - “FAILED_PRECONDITION“: on consistency errors (e.g. the contract key has changed since the submission)
	// or if an interpretation error occurred
	// - “UNAVAILABLE“: if the participant is not yet ready to submit commands or if the service has been shut down.
	Submit(context.Context, *SubmitRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCommandSubmissionServiceServer()
}

// UnimplementedCommandSubmissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandSubmissionServiceServer struct {
}

func (UnimplementedCommandSubmissionServiceServer) Submit(context.Context, *SubmitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
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
		FullMethod: "/com.daml.ledger.api.v1.CommandSubmissionService/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandSubmissionServiceServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandSubmissionService_ServiceDesc is the grpc.ServiceDesc for CommandSubmissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandSubmissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.CommandSubmissionService",
	HandlerType: (*CommandSubmissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _CommandSubmissionService_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v1/command_submission_service.proto",
}
