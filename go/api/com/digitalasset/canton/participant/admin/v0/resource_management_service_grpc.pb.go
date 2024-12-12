// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: com/digitalasset/canton/participant/admin/v0/resource_management_service.proto

package v0

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ResourceManagementService_UpdateResourceLimits_FullMethodName = "/com.digitalasset.canton.participant.admin.v0.ResourceManagementService/UpdateResourceLimits"
	ResourceManagementService_GetResourceLimits_FullMethodName    = "/com.digitalasset.canton.participant.admin.v0.ResourceManagementService/GetResourceLimits"
)

// ResourceManagementServiceClient is the client API for ResourceManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceManagementServiceClient interface {
	UpdateResourceLimits(ctx context.Context, in *ResourceLimits, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetResourceLimits(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResourceLimits, error)
}

type resourceManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceManagementServiceClient(cc grpc.ClientConnInterface) ResourceManagementServiceClient {
	return &resourceManagementServiceClient{cc}
}

func (c *resourceManagementServiceClient) UpdateResourceLimits(ctx context.Context, in *ResourceLimits, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ResourceManagementService_UpdateResourceLimits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceManagementServiceClient) GetResourceLimits(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResourceLimits, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourceLimits)
	err := c.cc.Invoke(ctx, ResourceManagementService_GetResourceLimits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceManagementServiceServer is the server API for ResourceManagementService service.
// All implementations must embed UnimplementedResourceManagementServiceServer
// for forward compatibility.
type ResourceManagementServiceServer interface {
	UpdateResourceLimits(context.Context, *ResourceLimits) (*emptypb.Empty, error)
	GetResourceLimits(context.Context, *emptypb.Empty) (*ResourceLimits, error)
	mustEmbedUnimplementedResourceManagementServiceServer()
}

// UnimplementedResourceManagementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResourceManagementServiceServer struct{}

func (UnimplementedResourceManagementServiceServer) UpdateResourceLimits(context.Context, *ResourceLimits) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResourceLimits not implemented")
}
func (UnimplementedResourceManagementServiceServer) GetResourceLimits(context.Context, *emptypb.Empty) (*ResourceLimits, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceLimits not implemented")
}
func (UnimplementedResourceManagementServiceServer) mustEmbedUnimplementedResourceManagementServiceServer() {
}
func (UnimplementedResourceManagementServiceServer) testEmbeddedByValue() {}

// UnsafeResourceManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceManagementServiceServer will
// result in compilation errors.
type UnsafeResourceManagementServiceServer interface {
	mustEmbedUnimplementedResourceManagementServiceServer()
}

func RegisterResourceManagementServiceServer(s grpc.ServiceRegistrar, srv ResourceManagementServiceServer) {
	// If the following call pancis, it indicates UnimplementedResourceManagementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResourceManagementService_ServiceDesc, srv)
}

func _ResourceManagementService_UpdateResourceLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceLimits)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceManagementServiceServer).UpdateResourceLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceManagementService_UpdateResourceLimits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceManagementServiceServer).UpdateResourceLimits(ctx, req.(*ResourceLimits))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceManagementService_GetResourceLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceManagementServiceServer).GetResourceLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceManagementService_GetResourceLimits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceManagementServiceServer).GetResourceLimits(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceManagementService_ServiceDesc is the grpc.ServiceDesc for ResourceManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.canton.participant.admin.v0.ResourceManagementService",
	HandlerType: (*ResourceManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateResourceLimits",
			Handler:    _ResourceManagementService_UpdateResourceLimits_Handler,
		},
		{
			MethodName: "GetResourceLimits",
			Handler:    _ResourceManagementService_GetResourceLimits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/canton/participant/admin/v0/resource_management_service.proto",
}
