// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: com/digitalasset/canton/admin/participant/v30/traffic_control_service.proto

package v30

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
	TrafficControlService_TrafficControlState_FullMethodName = "/com.digitalasset.canton.admin.participant.v30.TrafficControlService/TrafficControlState"
)

// TrafficControlServiceClient is the client API for TrafficControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrafficControlServiceClient interface {
	TrafficControlState(ctx context.Context, in *TrafficControlStateRequest, opts ...grpc.CallOption) (*TrafficControlStateResponse, error)
}

type trafficControlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrafficControlServiceClient(cc grpc.ClientConnInterface) TrafficControlServiceClient {
	return &trafficControlServiceClient{cc}
}

func (c *trafficControlServiceClient) TrafficControlState(ctx context.Context, in *TrafficControlStateRequest, opts ...grpc.CallOption) (*TrafficControlStateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrafficControlStateResponse)
	err := c.cc.Invoke(ctx, TrafficControlService_TrafficControlState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrafficControlServiceServer is the server API for TrafficControlService service.
// All implementations must embed UnimplementedTrafficControlServiceServer
// for forward compatibility.
type TrafficControlServiceServer interface {
	TrafficControlState(context.Context, *TrafficControlStateRequest) (*TrafficControlStateResponse, error)
	mustEmbedUnimplementedTrafficControlServiceServer()
}

// UnimplementedTrafficControlServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrafficControlServiceServer struct{}

func (UnimplementedTrafficControlServiceServer) TrafficControlState(context.Context, *TrafficControlStateRequest) (*TrafficControlStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrafficControlState not implemented")
}
func (UnimplementedTrafficControlServiceServer) mustEmbedUnimplementedTrafficControlServiceServer() {}
func (UnimplementedTrafficControlServiceServer) testEmbeddedByValue()                               {}

// UnsafeTrafficControlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrafficControlServiceServer will
// result in compilation errors.
type UnsafeTrafficControlServiceServer interface {
	mustEmbedUnimplementedTrafficControlServiceServer()
}

func RegisterTrafficControlServiceServer(s grpc.ServiceRegistrar, srv TrafficControlServiceServer) {
	// If the following call pancis, it indicates UnimplementedTrafficControlServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrafficControlService_ServiceDesc, srv)
}

func _TrafficControlService_TrafficControlState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficControlStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficControlServiceServer).TrafficControlState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrafficControlService_TrafficControlState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficControlServiceServer).TrafficControlState(ctx, req.(*TrafficControlStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrafficControlService_ServiceDesc is the grpc.ServiceDesc for TrafficControlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrafficControlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.canton.admin.participant.v30.TrafficControlService",
	HandlerType: (*TrafficControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrafficControlState",
			Handler:    _TrafficControlService_TrafficControlState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/canton/admin/participant/v30/traffic_control_service.proto",
}