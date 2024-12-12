// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: com/daml/ledger/api/v2/version_service.proto

package v2

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
	VersionService_GetLedgerApiVersion_FullMethodName = "/com.daml.ledger.api.v2.VersionService/GetLedgerApiVersion"
)

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionServiceClient interface {
	GetLedgerApiVersion(ctx context.Context, in *GetLedgerApiVersionRequest, opts ...grpc.CallOption) (*GetLedgerApiVersionResponse, error)
}

type versionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionServiceClient(cc grpc.ClientConnInterface) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) GetLedgerApiVersion(ctx context.Context, in *GetLedgerApiVersionRequest, opts ...grpc.CallOption) (*GetLedgerApiVersionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLedgerApiVersionResponse)
	err := c.cc.Invoke(ctx, VersionService_GetLedgerApiVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
// All implementations must embed UnimplementedVersionServiceServer
// for forward compatibility.
type VersionServiceServer interface {
	GetLedgerApiVersion(context.Context, *GetLedgerApiVersionRequest) (*GetLedgerApiVersionResponse, error)
	mustEmbedUnimplementedVersionServiceServer()
}

// UnimplementedVersionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVersionServiceServer struct{}

func (UnimplementedVersionServiceServer) GetLedgerApiVersion(context.Context, *GetLedgerApiVersionRequest) (*GetLedgerApiVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLedgerApiVersion not implemented")
}
func (UnimplementedVersionServiceServer) mustEmbedUnimplementedVersionServiceServer() {}
func (UnimplementedVersionServiceServer) testEmbeddedByValue()                        {}

// UnsafeVersionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionServiceServer will
// result in compilation errors.
type UnsafeVersionServiceServer interface {
	mustEmbedUnimplementedVersionServiceServer()
}

func RegisterVersionServiceServer(s grpc.ServiceRegistrar, srv VersionServiceServer) {
	// If the following call pancis, it indicates UnimplementedVersionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VersionService_ServiceDesc, srv)
}

func _VersionService_GetLedgerApiVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLedgerApiVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).GetLedgerApiVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VersionService_GetLedgerApiVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).GetLedgerApiVersion(ctx, req.(*GetLedgerApiVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionService_ServiceDesc is the grpc.ServiceDesc for VersionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v2.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLedgerApiVersion",
			Handler:    _VersionService_GetLedgerApiVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v2/version_service.proto",
}