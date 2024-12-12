// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: com/digitalasset/canton/mediator/admin/v30/mediator_administration_service.proto

package v30

import (
	context "context"
	v30 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/admin/pruning/v30"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MediatorAdministrationService_Prune_FullMethodName                  = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/Prune"
	MediatorAdministrationService_SetSchedule_FullMethodName            = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/SetSchedule"
	MediatorAdministrationService_SetCron_FullMethodName                = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/SetCron"
	MediatorAdministrationService_SetMaxDuration_FullMethodName         = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/SetMaxDuration"
	MediatorAdministrationService_SetRetention_FullMethodName           = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/SetRetention"
	MediatorAdministrationService_ClearSchedule_FullMethodName          = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/ClearSchedule"
	MediatorAdministrationService_GetSchedule_FullMethodName            = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/GetSchedule"
	MediatorAdministrationService_LocatePruningTimestamp_FullMethodName = "/com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService/LocatePruningTimestamp"
)

// MediatorAdministrationServiceClient is the client API for MediatorAdministrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediatorAdministrationServiceClient interface {
	Prune(ctx context.Context, in *MediatorPruning_PruneRequest, opts ...grpc.CallOption) (*MediatorPruning_PruneResponse, error)
	SetSchedule(ctx context.Context, in *v30.SetSchedule_Request, opts ...grpc.CallOption) (*v30.SetSchedule_Response, error)
	SetCron(ctx context.Context, in *v30.SetCron_Request, opts ...grpc.CallOption) (*v30.SetCron_Response, error)
	SetMaxDuration(ctx context.Context, in *v30.SetMaxDuration_Request, opts ...grpc.CallOption) (*v30.SetMaxDuration_Response, error)
	SetRetention(ctx context.Context, in *v30.SetRetention_Request, opts ...grpc.CallOption) (*v30.SetRetention_Response, error)
	ClearSchedule(ctx context.Context, in *v30.ClearSchedule_Request, opts ...grpc.CallOption) (*v30.ClearSchedule_Response, error)
	GetSchedule(ctx context.Context, in *v30.GetSchedule_Request, opts ...grpc.CallOption) (*v30.GetSchedule_Response, error)
	LocatePruningTimestamp(ctx context.Context, in *v30.LocatePruningTimestamp_Request, opts ...grpc.CallOption) (*v30.LocatePruningTimestamp_Response, error)
}

type mediatorAdministrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediatorAdministrationServiceClient(cc grpc.ClientConnInterface) MediatorAdministrationServiceClient {
	return &mediatorAdministrationServiceClient{cc}
}

func (c *mediatorAdministrationServiceClient) Prune(ctx context.Context, in *MediatorPruning_PruneRequest, opts ...grpc.CallOption) (*MediatorPruning_PruneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MediatorPruning_PruneResponse)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_Prune_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediatorAdministrationServiceClient) SetSchedule(ctx context.Context, in *v30.SetSchedule_Request, opts ...grpc.CallOption) (*v30.SetSchedule_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v30.SetSchedule_Response)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_SetSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediatorAdministrationServiceClient) SetCron(ctx context.Context, in *v30.SetCron_Request, opts ...grpc.CallOption) (*v30.SetCron_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v30.SetCron_Response)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_SetCron_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediatorAdministrationServiceClient) SetMaxDuration(ctx context.Context, in *v30.SetMaxDuration_Request, opts ...grpc.CallOption) (*v30.SetMaxDuration_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v30.SetMaxDuration_Response)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_SetMaxDuration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediatorAdministrationServiceClient) SetRetention(ctx context.Context, in *v30.SetRetention_Request, opts ...grpc.CallOption) (*v30.SetRetention_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v30.SetRetention_Response)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_SetRetention_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediatorAdministrationServiceClient) ClearSchedule(ctx context.Context, in *v30.ClearSchedule_Request, opts ...grpc.CallOption) (*v30.ClearSchedule_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v30.ClearSchedule_Response)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_ClearSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediatorAdministrationServiceClient) GetSchedule(ctx context.Context, in *v30.GetSchedule_Request, opts ...grpc.CallOption) (*v30.GetSchedule_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v30.GetSchedule_Response)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_GetSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediatorAdministrationServiceClient) LocatePruningTimestamp(ctx context.Context, in *v30.LocatePruningTimestamp_Request, opts ...grpc.CallOption) (*v30.LocatePruningTimestamp_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v30.LocatePruningTimestamp_Response)
	err := c.cc.Invoke(ctx, MediatorAdministrationService_LocatePruningTimestamp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediatorAdministrationServiceServer is the server API for MediatorAdministrationService service.
// All implementations must embed UnimplementedMediatorAdministrationServiceServer
// for forward compatibility.
type MediatorAdministrationServiceServer interface {
	Prune(context.Context, *MediatorPruning_PruneRequest) (*MediatorPruning_PruneResponse, error)
	SetSchedule(context.Context, *v30.SetSchedule_Request) (*v30.SetSchedule_Response, error)
	SetCron(context.Context, *v30.SetCron_Request) (*v30.SetCron_Response, error)
	SetMaxDuration(context.Context, *v30.SetMaxDuration_Request) (*v30.SetMaxDuration_Response, error)
	SetRetention(context.Context, *v30.SetRetention_Request) (*v30.SetRetention_Response, error)
	ClearSchedule(context.Context, *v30.ClearSchedule_Request) (*v30.ClearSchedule_Response, error)
	GetSchedule(context.Context, *v30.GetSchedule_Request) (*v30.GetSchedule_Response, error)
	LocatePruningTimestamp(context.Context, *v30.LocatePruningTimestamp_Request) (*v30.LocatePruningTimestamp_Response, error)
	mustEmbedUnimplementedMediatorAdministrationServiceServer()
}

// UnimplementedMediatorAdministrationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMediatorAdministrationServiceServer struct{}

func (UnimplementedMediatorAdministrationServiceServer) Prune(context.Context, *MediatorPruning_PruneRequest) (*MediatorPruning_PruneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prune not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) SetSchedule(context.Context, *v30.SetSchedule_Request) (*v30.SetSchedule_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSchedule not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) SetCron(context.Context, *v30.SetCron_Request) (*v30.SetCron_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCron not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) SetMaxDuration(context.Context, *v30.SetMaxDuration_Request) (*v30.SetMaxDuration_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMaxDuration not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) SetRetention(context.Context, *v30.SetRetention_Request) (*v30.SetRetention_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRetention not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) ClearSchedule(context.Context, *v30.ClearSchedule_Request) (*v30.ClearSchedule_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearSchedule not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) GetSchedule(context.Context, *v30.GetSchedule_Request) (*v30.GetSchedule_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) LocatePruningTimestamp(context.Context, *v30.LocatePruningTimestamp_Request) (*v30.LocatePruningTimestamp_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocatePruningTimestamp not implemented")
}
func (UnimplementedMediatorAdministrationServiceServer) mustEmbedUnimplementedMediatorAdministrationServiceServer() {
}
func (UnimplementedMediatorAdministrationServiceServer) testEmbeddedByValue() {}

// UnsafeMediatorAdministrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediatorAdministrationServiceServer will
// result in compilation errors.
type UnsafeMediatorAdministrationServiceServer interface {
	mustEmbedUnimplementedMediatorAdministrationServiceServer()
}

func RegisterMediatorAdministrationServiceServer(s grpc.ServiceRegistrar, srv MediatorAdministrationServiceServer) {
	// If the following call pancis, it indicates UnimplementedMediatorAdministrationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MediatorAdministrationService_ServiceDesc, srv)
}

func _MediatorAdministrationService_Prune_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediatorPruning_PruneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).Prune(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_Prune_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).Prune(ctx, req.(*MediatorPruning_PruneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediatorAdministrationService_SetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v30.SetSchedule_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).SetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_SetSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).SetSchedule(ctx, req.(*v30.SetSchedule_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediatorAdministrationService_SetCron_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v30.SetCron_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).SetCron(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_SetCron_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).SetCron(ctx, req.(*v30.SetCron_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediatorAdministrationService_SetMaxDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v30.SetMaxDuration_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).SetMaxDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_SetMaxDuration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).SetMaxDuration(ctx, req.(*v30.SetMaxDuration_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediatorAdministrationService_SetRetention_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v30.SetRetention_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).SetRetention(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_SetRetention_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).SetRetention(ctx, req.(*v30.SetRetention_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediatorAdministrationService_ClearSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v30.ClearSchedule_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).ClearSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_ClearSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).ClearSchedule(ctx, req.(*v30.ClearSchedule_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediatorAdministrationService_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v30.GetSchedule_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_GetSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).GetSchedule(ctx, req.(*v30.GetSchedule_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediatorAdministrationService_LocatePruningTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v30.LocatePruningTimestamp_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediatorAdministrationServiceServer).LocatePruningTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediatorAdministrationService_LocatePruningTimestamp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediatorAdministrationServiceServer).LocatePruningTimestamp(ctx, req.(*v30.LocatePruningTimestamp_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MediatorAdministrationService_ServiceDesc is the grpc.ServiceDesc for MediatorAdministrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediatorAdministrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService",
	HandlerType: (*MediatorAdministrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prune",
			Handler:    _MediatorAdministrationService_Prune_Handler,
		},
		{
			MethodName: "SetSchedule",
			Handler:    _MediatorAdministrationService_SetSchedule_Handler,
		},
		{
			MethodName: "SetCron",
			Handler:    _MediatorAdministrationService_SetCron_Handler,
		},
		{
			MethodName: "SetMaxDuration",
			Handler:    _MediatorAdministrationService_SetMaxDuration_Handler,
		},
		{
			MethodName: "SetRetention",
			Handler:    _MediatorAdministrationService_SetRetention_Handler,
		},
		{
			MethodName: "ClearSchedule",
			Handler:    _MediatorAdministrationService_ClearSchedule_Handler,
		},
		{
			MethodName: "GetSchedule",
			Handler:    _MediatorAdministrationService_GetSchedule_Handler,
		},
		{
			MethodName: "LocatePruningTimestamp",
			Handler:    _MediatorAdministrationService_LocatePruningTimestamp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/canton/mediator/admin/v30/mediator_administration_service.proto",
}
