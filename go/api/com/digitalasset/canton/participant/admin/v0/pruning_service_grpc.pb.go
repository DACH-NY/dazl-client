// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: com/digitalasset/canton/participant/admin/v0/pruning_service.proto

package v0

import (
	context "context"
	v0 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/pruning/admin/v0"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PruningServiceClient is the client API for PruningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PruningServiceClient interface {
	// Prune the participant specifying the offset before and at which ledger transactions
	// should be removed. Only returns when the potentially long-running prune request ends
	// successfully or with one of the following errors:
	// - ``INVALID_ARGUMENT``: if the payload, particularly the offset is malformed or missing
	// - ``INTERNAL``: if the participant has encountered a failure and has potentially
	//   applied pruning partially. Such cases warrant verifying the participant health before
	//   retrying the prune with the same (or a larger, valid) offset. Successful retries
	//   after such errors ensure that different components reach a consistent pruning state.
	// - ``FAILED_PRECONDITION``: if the participant is not yet able to prune at the specified
	//   offset or if pruning is invoked on a participant running the Community Edition.
	Prune(ctx context.Context, in *PruneRequest, opts ...grpc.CallOption) (*PruneResponse, error)
	// Retrieve the safe pruning offset
	GetSafePruningOffset(ctx context.Context, in *GetSafePruningOffsetRequest, opts ...grpc.CallOption) (*GetSafePruningOffsetResponse, error)
	// Enable automatic pruning using the specified schedule parameters
	// The following errors may occur on the SetSchedule or Update commands:
	// - ``INVALID_ARGUMENT``: if a parameter is missing or an invalid cron expression
	//   or duration.
	// - ``FAILED_PRECONDITION``: if automatic background pruning has not been enabled
	//   or if invoked on a participant running the Community Edition.
	SetSchedule(ctx context.Context, in *v0.SetSchedule_Request, opts ...grpc.CallOption) (*v0.SetSchedule_Response, error)
	// Modify individual pruning schedule parameters.
	// - ``INVALID_ARGUMENT``: if the payload is malformed or no schedule is configured
	SetCron(ctx context.Context, in *v0.SetCron_Request, opts ...grpc.CallOption) (*v0.SetCron_Response, error)
	SetMaxDuration(ctx context.Context, in *v0.SetMaxDuration_Request, opts ...grpc.CallOption) (*v0.SetMaxDuration_Response, error)
	SetRetention(ctx context.Context, in *v0.SetRetention_Request, opts ...grpc.CallOption) (*v0.SetRetention_Response, error)
	// Disable automatic pruning and remove the persisted schedule configuration.
	ClearSchedule(ctx context.Context, in *v0.ClearSchedule_Request, opts ...grpc.CallOption) (*v0.ClearSchedule_Response, error)
	// Retrieve the automatic pruning configuration.
	GetSchedule(ctx context.Context, in *v0.GetSchedule_Request, opts ...grpc.CallOption) (*v0.GetSchedule_Response, error)
}

type pruningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPruningServiceClient(cc grpc.ClientConnInterface) PruningServiceClient {
	return &pruningServiceClient{cc}
}

func (c *pruningServiceClient) Prune(ctx context.Context, in *PruneRequest, opts ...grpc.CallOption) (*PruneResponse, error) {
	out := new(PruneResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/Prune", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pruningServiceClient) GetSafePruningOffset(ctx context.Context, in *GetSafePruningOffsetRequest, opts ...grpc.CallOption) (*GetSafePruningOffsetResponse, error) {
	out := new(GetSafePruningOffsetResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/GetSafePruningOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pruningServiceClient) SetSchedule(ctx context.Context, in *v0.SetSchedule_Request, opts ...grpc.CallOption) (*v0.SetSchedule_Response, error) {
	out := new(v0.SetSchedule_Response)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/SetSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pruningServiceClient) SetCron(ctx context.Context, in *v0.SetCron_Request, opts ...grpc.CallOption) (*v0.SetCron_Response, error) {
	out := new(v0.SetCron_Response)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/SetCron", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pruningServiceClient) SetMaxDuration(ctx context.Context, in *v0.SetMaxDuration_Request, opts ...grpc.CallOption) (*v0.SetMaxDuration_Response, error) {
	out := new(v0.SetMaxDuration_Response)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/SetMaxDuration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pruningServiceClient) SetRetention(ctx context.Context, in *v0.SetRetention_Request, opts ...grpc.CallOption) (*v0.SetRetention_Response, error) {
	out := new(v0.SetRetention_Response)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/SetRetention", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pruningServiceClient) ClearSchedule(ctx context.Context, in *v0.ClearSchedule_Request, opts ...grpc.CallOption) (*v0.ClearSchedule_Response, error) {
	out := new(v0.ClearSchedule_Response)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/ClearSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pruningServiceClient) GetSchedule(ctx context.Context, in *v0.GetSchedule_Request, opts ...grpc.CallOption) (*v0.GetSchedule_Response, error) {
	out := new(v0.GetSchedule_Response)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PruningService/GetSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PruningServiceServer is the server API for PruningService service.
// All implementations must embed UnimplementedPruningServiceServer
// for forward compatibility
type PruningServiceServer interface {
	// Prune the participant specifying the offset before and at which ledger transactions
	// should be removed. Only returns when the potentially long-running prune request ends
	// successfully or with one of the following errors:
	// - ``INVALID_ARGUMENT``: if the payload, particularly the offset is malformed or missing
	// - ``INTERNAL``: if the participant has encountered a failure and has potentially
	//   applied pruning partially. Such cases warrant verifying the participant health before
	//   retrying the prune with the same (or a larger, valid) offset. Successful retries
	//   after such errors ensure that different components reach a consistent pruning state.
	// - ``FAILED_PRECONDITION``: if the participant is not yet able to prune at the specified
	//   offset or if pruning is invoked on a participant running the Community Edition.
	Prune(context.Context, *PruneRequest) (*PruneResponse, error)
	// Retrieve the safe pruning offset
	GetSafePruningOffset(context.Context, *GetSafePruningOffsetRequest) (*GetSafePruningOffsetResponse, error)
	// Enable automatic pruning using the specified schedule parameters
	// The following errors may occur on the SetSchedule or Update commands:
	// - ``INVALID_ARGUMENT``: if a parameter is missing or an invalid cron expression
	//   or duration.
	// - ``FAILED_PRECONDITION``: if automatic background pruning has not been enabled
	//   or if invoked on a participant running the Community Edition.
	SetSchedule(context.Context, *v0.SetSchedule_Request) (*v0.SetSchedule_Response, error)
	// Modify individual pruning schedule parameters.
	// - ``INVALID_ARGUMENT``: if the payload is malformed or no schedule is configured
	SetCron(context.Context, *v0.SetCron_Request) (*v0.SetCron_Response, error)
	SetMaxDuration(context.Context, *v0.SetMaxDuration_Request) (*v0.SetMaxDuration_Response, error)
	SetRetention(context.Context, *v0.SetRetention_Request) (*v0.SetRetention_Response, error)
	// Disable automatic pruning and remove the persisted schedule configuration.
	ClearSchedule(context.Context, *v0.ClearSchedule_Request) (*v0.ClearSchedule_Response, error)
	// Retrieve the automatic pruning configuration.
	GetSchedule(context.Context, *v0.GetSchedule_Request) (*v0.GetSchedule_Response, error)
	mustEmbedUnimplementedPruningServiceServer()
}

// UnimplementedPruningServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPruningServiceServer struct {
}

func (UnimplementedPruningServiceServer) Prune(context.Context, *PruneRequest) (*PruneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prune not implemented")
}
func (UnimplementedPruningServiceServer) GetSafePruningOffset(context.Context, *GetSafePruningOffsetRequest) (*GetSafePruningOffsetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSafePruningOffset not implemented")
}
func (UnimplementedPruningServiceServer) SetSchedule(context.Context, *v0.SetSchedule_Request) (*v0.SetSchedule_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSchedule not implemented")
}
func (UnimplementedPruningServiceServer) SetCron(context.Context, *v0.SetCron_Request) (*v0.SetCron_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCron not implemented")
}
func (UnimplementedPruningServiceServer) SetMaxDuration(context.Context, *v0.SetMaxDuration_Request) (*v0.SetMaxDuration_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMaxDuration not implemented")
}
func (UnimplementedPruningServiceServer) SetRetention(context.Context, *v0.SetRetention_Request) (*v0.SetRetention_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRetention not implemented")
}
func (UnimplementedPruningServiceServer) ClearSchedule(context.Context, *v0.ClearSchedule_Request) (*v0.ClearSchedule_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearSchedule not implemented")
}
func (UnimplementedPruningServiceServer) GetSchedule(context.Context, *v0.GetSchedule_Request) (*v0.GetSchedule_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (UnimplementedPruningServiceServer) mustEmbedUnimplementedPruningServiceServer() {}

// UnsafePruningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PruningServiceServer will
// result in compilation errors.
type UnsafePruningServiceServer interface {
	mustEmbedUnimplementedPruningServiceServer()
}

func RegisterPruningServiceServer(s grpc.ServiceRegistrar, srv PruningServiceServer) {
	s.RegisterService(&PruningService_ServiceDesc, srv)
}

func _PruningService_Prune_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PruneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).Prune(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/Prune",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).Prune(ctx, req.(*PruneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PruningService_GetSafePruningOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSafePruningOffsetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).GetSafePruningOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/GetSafePruningOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).GetSafePruningOffset(ctx, req.(*GetSafePruningOffsetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PruningService_SetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v0.SetSchedule_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).SetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/SetSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).SetSchedule(ctx, req.(*v0.SetSchedule_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PruningService_SetCron_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v0.SetCron_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).SetCron(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/SetCron",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).SetCron(ctx, req.(*v0.SetCron_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PruningService_SetMaxDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v0.SetMaxDuration_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).SetMaxDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/SetMaxDuration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).SetMaxDuration(ctx, req.(*v0.SetMaxDuration_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PruningService_SetRetention_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v0.SetRetention_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).SetRetention(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/SetRetention",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).SetRetention(ctx, req.(*v0.SetRetention_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PruningService_ClearSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v0.ClearSchedule_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).ClearSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/ClearSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).ClearSchedule(ctx, req.(*v0.ClearSchedule_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PruningService_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v0.GetSchedule_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PruningServiceServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PruningService/GetSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PruningServiceServer).GetSchedule(ctx, req.(*v0.GetSchedule_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PruningService_ServiceDesc is the grpc.ServiceDesc for PruningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PruningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.canton.participant.admin.v0.PruningService",
	HandlerType: (*PruningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prune",
			Handler:    _PruningService_Prune_Handler,
		},
		{
			MethodName: "GetSafePruningOffset",
			Handler:    _PruningService_GetSafePruningOffset_Handler,
		},
		{
			MethodName: "SetSchedule",
			Handler:    _PruningService_SetSchedule_Handler,
		},
		{
			MethodName: "SetCron",
			Handler:    _PruningService_SetCron_Handler,
		},
		{
			MethodName: "SetMaxDuration",
			Handler:    _PruningService_SetMaxDuration_Handler,
		},
		{
			MethodName: "SetRetention",
			Handler:    _PruningService_SetRetention_Handler,
		},
		{
			MethodName: "ClearSchedule",
			Handler:    _PruningService_ClearSchedule_Handler,
		},
		{
			MethodName: "GetSchedule",
			Handler:    _PruningService_GetSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/canton/participant/admin/v0/pruning_service.proto",
}