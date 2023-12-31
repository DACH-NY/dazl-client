// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: com/daml/ledger/api/v1/admin/metering_report_service.proto

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MeteringReportServiceClient is the client API for MeteringReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeteringReportServiceClient interface {
	// Retrieve a metering report.
	GetMeteringReport(ctx context.Context, in *GetMeteringReportRequest, opts ...grpc.CallOption) (*GetMeteringReportResponse, error)
}

type meteringReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeteringReportServiceClient(cc grpc.ClientConnInterface) MeteringReportServiceClient {
	return &meteringReportServiceClient{cc}
}

func (c *meteringReportServiceClient) GetMeteringReport(ctx context.Context, in *GetMeteringReportRequest, opts ...grpc.CallOption) (*GetMeteringReportResponse, error) {
	out := new(GetMeteringReportResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.MeteringReportService/GetMeteringReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeteringReportServiceServer is the server API for MeteringReportService service.
// All implementations must embed UnimplementedMeteringReportServiceServer
// for forward compatibility
type MeteringReportServiceServer interface {
	// Retrieve a metering report.
	GetMeteringReport(context.Context, *GetMeteringReportRequest) (*GetMeteringReportResponse, error)
	mustEmbedUnimplementedMeteringReportServiceServer()
}

// UnimplementedMeteringReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeteringReportServiceServer struct {
}

func (UnimplementedMeteringReportServiceServer) GetMeteringReport(context.Context, *GetMeteringReportRequest) (*GetMeteringReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeteringReport not implemented")
}
func (UnimplementedMeteringReportServiceServer) mustEmbedUnimplementedMeteringReportServiceServer() {}

// UnsafeMeteringReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeteringReportServiceServer will
// result in compilation errors.
type UnsafeMeteringReportServiceServer interface {
	mustEmbedUnimplementedMeteringReportServiceServer()
}

func RegisterMeteringReportServiceServer(s grpc.ServiceRegistrar, srv MeteringReportServiceServer) {
	s.RegisterService(&MeteringReportService_ServiceDesc, srv)
}

func _MeteringReportService_GetMeteringReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeteringReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeteringReportServiceServer).GetMeteringReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.MeteringReportService/GetMeteringReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeteringReportServiceServer).GetMeteringReport(ctx, req.(*GetMeteringReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeteringReportService_ServiceDesc is the grpc.ServiceDesc for MeteringReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeteringReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.admin.MeteringReportService",
	HandlerType: (*MeteringReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMeteringReport",
			Handler:    _MeteringReportService_GetMeteringReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v1/admin/metering_report_service.proto",
}
