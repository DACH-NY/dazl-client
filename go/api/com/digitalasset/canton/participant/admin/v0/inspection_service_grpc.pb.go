// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: com/digitalasset/canton/participant/admin/v0/inspection_service.proto

package v0

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
	InspectionService_LookupContractDomain_FullMethodName    = "/com.digitalasset.canton.participant.admin.v0.InspectionService/LookupContractDomain"
	InspectionService_LookupTransactionDomain_FullMethodName = "/com.digitalasset.canton.participant.admin.v0.InspectionService/LookupTransactionDomain"
	InspectionService_LookupOffsetByTime_FullMethodName      = "/com.digitalasset.canton.participant.admin.v0.InspectionService/LookupOffsetByTime"
	InspectionService_LookupOffsetByIndex_FullMethodName     = "/com.digitalasset.canton.participant.admin.v0.InspectionService/LookupOffsetByIndex"
	InspectionService_CountInFlight_FullMethodName           = "/com.digitalasset.canton.participant.admin.v0.InspectionService/CountInFlight"
)

// InspectionServiceClient is the client API for InspectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InspectionServiceClient interface {
	LookupContractDomain(ctx context.Context, in *LookupContractDomain_Request, opts ...grpc.CallOption) (*LookupContractDomain_Response, error)
	LookupTransactionDomain(ctx context.Context, in *LookupTransactionDomain_Request, opts ...grpc.CallOption) (*LookupTransactionDomain_Response, error)
	LookupOffsetByTime(ctx context.Context, in *LookupOffsetByTime_Request, opts ...grpc.CallOption) (*LookupOffsetByTime_Response, error)
	LookupOffsetByIndex(ctx context.Context, in *LookupOffsetByIndex_Request, opts ...grpc.CallOption) (*LookupOffsetByIndex_Response, error)
	CountInFlight(ctx context.Context, in *CountInFlight_Request, opts ...grpc.CallOption) (*CountInFlight_Response, error)
}

type inspectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInspectionServiceClient(cc grpc.ClientConnInterface) InspectionServiceClient {
	return &inspectionServiceClient{cc}
}

func (c *inspectionServiceClient) LookupContractDomain(ctx context.Context, in *LookupContractDomain_Request, opts ...grpc.CallOption) (*LookupContractDomain_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupContractDomain_Response)
	err := c.cc.Invoke(ctx, InspectionService_LookupContractDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectionServiceClient) LookupTransactionDomain(ctx context.Context, in *LookupTransactionDomain_Request, opts ...grpc.CallOption) (*LookupTransactionDomain_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupTransactionDomain_Response)
	err := c.cc.Invoke(ctx, InspectionService_LookupTransactionDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectionServiceClient) LookupOffsetByTime(ctx context.Context, in *LookupOffsetByTime_Request, opts ...grpc.CallOption) (*LookupOffsetByTime_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupOffsetByTime_Response)
	err := c.cc.Invoke(ctx, InspectionService_LookupOffsetByTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectionServiceClient) LookupOffsetByIndex(ctx context.Context, in *LookupOffsetByIndex_Request, opts ...grpc.CallOption) (*LookupOffsetByIndex_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupOffsetByIndex_Response)
	err := c.cc.Invoke(ctx, InspectionService_LookupOffsetByIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectionServiceClient) CountInFlight(ctx context.Context, in *CountInFlight_Request, opts ...grpc.CallOption) (*CountInFlight_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountInFlight_Response)
	err := c.cc.Invoke(ctx, InspectionService_CountInFlight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InspectionServiceServer is the server API for InspectionService service.
// All implementations must embed UnimplementedInspectionServiceServer
// for forward compatibility.
type InspectionServiceServer interface {
	LookupContractDomain(context.Context, *LookupContractDomain_Request) (*LookupContractDomain_Response, error)
	LookupTransactionDomain(context.Context, *LookupTransactionDomain_Request) (*LookupTransactionDomain_Response, error)
	LookupOffsetByTime(context.Context, *LookupOffsetByTime_Request) (*LookupOffsetByTime_Response, error)
	LookupOffsetByIndex(context.Context, *LookupOffsetByIndex_Request) (*LookupOffsetByIndex_Response, error)
	CountInFlight(context.Context, *CountInFlight_Request) (*CountInFlight_Response, error)
	mustEmbedUnimplementedInspectionServiceServer()
}

// UnimplementedInspectionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInspectionServiceServer struct{}

func (UnimplementedInspectionServiceServer) LookupContractDomain(context.Context, *LookupContractDomain_Request) (*LookupContractDomain_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupContractDomain not implemented")
}
func (UnimplementedInspectionServiceServer) LookupTransactionDomain(context.Context, *LookupTransactionDomain_Request) (*LookupTransactionDomain_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupTransactionDomain not implemented")
}
func (UnimplementedInspectionServiceServer) LookupOffsetByTime(context.Context, *LookupOffsetByTime_Request) (*LookupOffsetByTime_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupOffsetByTime not implemented")
}
func (UnimplementedInspectionServiceServer) LookupOffsetByIndex(context.Context, *LookupOffsetByIndex_Request) (*LookupOffsetByIndex_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupOffsetByIndex not implemented")
}
func (UnimplementedInspectionServiceServer) CountInFlight(context.Context, *CountInFlight_Request) (*CountInFlight_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountInFlight not implemented")
}
func (UnimplementedInspectionServiceServer) mustEmbedUnimplementedInspectionServiceServer() {}
func (UnimplementedInspectionServiceServer) testEmbeddedByValue()                           {}

// UnsafeInspectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InspectionServiceServer will
// result in compilation errors.
type UnsafeInspectionServiceServer interface {
	mustEmbedUnimplementedInspectionServiceServer()
}

func RegisterInspectionServiceServer(s grpc.ServiceRegistrar, srv InspectionServiceServer) {
	// If the following call pancis, it indicates UnimplementedInspectionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InspectionService_ServiceDesc, srv)
}

func _InspectionService_LookupContractDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupContractDomain_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectionServiceServer).LookupContractDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InspectionService_LookupContractDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectionServiceServer).LookupContractDomain(ctx, req.(*LookupContractDomain_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _InspectionService_LookupTransactionDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupTransactionDomain_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectionServiceServer).LookupTransactionDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InspectionService_LookupTransactionDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectionServiceServer).LookupTransactionDomain(ctx, req.(*LookupTransactionDomain_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _InspectionService_LookupOffsetByTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupOffsetByTime_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectionServiceServer).LookupOffsetByTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InspectionService_LookupOffsetByTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectionServiceServer).LookupOffsetByTime(ctx, req.(*LookupOffsetByTime_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _InspectionService_LookupOffsetByIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupOffsetByIndex_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectionServiceServer).LookupOffsetByIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InspectionService_LookupOffsetByIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectionServiceServer).LookupOffsetByIndex(ctx, req.(*LookupOffsetByIndex_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _InspectionService_CountInFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountInFlight_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectionServiceServer).CountInFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InspectionService_CountInFlight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectionServiceServer).CountInFlight(ctx, req.(*CountInFlight_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// InspectionService_ServiceDesc is the grpc.ServiceDesc for InspectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InspectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.canton.participant.admin.v0.InspectionService",
	HandlerType: (*InspectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupContractDomain",
			Handler:    _InspectionService_LookupContractDomain_Handler,
		},
		{
			MethodName: "LookupTransactionDomain",
			Handler:    _InspectionService_LookupTransactionDomain_Handler,
		},
		{
			MethodName: "LookupOffsetByTime",
			Handler:    _InspectionService_LookupOffsetByTime_Handler,
		},
		{
			MethodName: "LookupOffsetByIndex",
			Handler:    _InspectionService_LookupOffsetByIndex_Handler,
		},
		{
			MethodName: "CountInFlight",
			Handler:    _InspectionService_CountInFlight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/canton/participant/admin/v0/inspection_service.proto",
}
