// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: com/daml/ledger/api/v1/ledger_configuration_service.proto

package v1

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

// LedgerConfigurationServiceClient is the client API for LedgerConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LedgerConfigurationServiceClient interface {
	GetLedgerConfiguration(ctx context.Context, in *GetLedgerConfigurationRequest, opts ...grpc.CallOption) (LedgerConfigurationService_GetLedgerConfigurationClient, error)
}

type ledgerConfigurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerConfigurationServiceClient(cc grpc.ClientConnInterface) LedgerConfigurationServiceClient {
	return &ledgerConfigurationServiceClient{cc}
}

func (c *ledgerConfigurationServiceClient) GetLedgerConfiguration(ctx context.Context, in *GetLedgerConfigurationRequest, opts ...grpc.CallOption) (LedgerConfigurationService_GetLedgerConfigurationClient, error) {
	stream, err := c.cc.NewStream(ctx, &LedgerConfigurationService_ServiceDesc.Streams[0], "/com.daml.ledger.api.v1.LedgerConfigurationService/GetLedgerConfiguration", opts...)
	if err != nil {
		return nil, err
	}
	x := &ledgerConfigurationServiceGetLedgerConfigurationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LedgerConfigurationService_GetLedgerConfigurationClient interface {
	Recv() (*GetLedgerConfigurationResponse, error)
	grpc.ClientStream
}

type ledgerConfigurationServiceGetLedgerConfigurationClient struct {
	grpc.ClientStream
}

func (x *ledgerConfigurationServiceGetLedgerConfigurationClient) Recv() (*GetLedgerConfigurationResponse, error) {
	m := new(GetLedgerConfigurationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LedgerConfigurationServiceServer is the server API for LedgerConfigurationService service.
// All implementations must embed UnimplementedLedgerConfigurationServiceServer
// for forward compatibility
type LedgerConfigurationServiceServer interface {
	GetLedgerConfiguration(*GetLedgerConfigurationRequest, LedgerConfigurationService_GetLedgerConfigurationServer) error
	mustEmbedUnimplementedLedgerConfigurationServiceServer()
}

// UnimplementedLedgerConfigurationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLedgerConfigurationServiceServer struct {
}

func (UnimplementedLedgerConfigurationServiceServer) GetLedgerConfiguration(*GetLedgerConfigurationRequest, LedgerConfigurationService_GetLedgerConfigurationServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLedgerConfiguration not implemented")
}
func (UnimplementedLedgerConfigurationServiceServer) mustEmbedUnimplementedLedgerConfigurationServiceServer() {
}

// UnsafeLedgerConfigurationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LedgerConfigurationServiceServer will
// result in compilation errors.
type UnsafeLedgerConfigurationServiceServer interface {
	mustEmbedUnimplementedLedgerConfigurationServiceServer()
}

func RegisterLedgerConfigurationServiceServer(s grpc.ServiceRegistrar, srv LedgerConfigurationServiceServer) {
	s.RegisterService(&LedgerConfigurationService_ServiceDesc, srv)
}

func _LedgerConfigurationService_GetLedgerConfiguration_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLedgerConfigurationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LedgerConfigurationServiceServer).GetLedgerConfiguration(m, &ledgerConfigurationServiceGetLedgerConfigurationServer{stream})
}

type LedgerConfigurationService_GetLedgerConfigurationServer interface {
	Send(*GetLedgerConfigurationResponse) error
	grpc.ServerStream
}

type ledgerConfigurationServiceGetLedgerConfigurationServer struct {
	grpc.ServerStream
}

func (x *ledgerConfigurationServiceGetLedgerConfigurationServer) Send(m *GetLedgerConfigurationResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LedgerConfigurationService_ServiceDesc is the grpc.ServiceDesc for LedgerConfigurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LedgerConfigurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.LedgerConfigurationService",
	HandlerType: (*LedgerConfigurationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLedgerConfiguration",
			Handler:       _LedgerConfigurationService_GetLedgerConfiguration_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/daml/ledger/api/v1/ledger_configuration_service.proto",
}
