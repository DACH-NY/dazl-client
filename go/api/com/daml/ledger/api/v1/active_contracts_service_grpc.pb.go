// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.5
// source: com/daml/ledger/api/v1/active_contracts_service.proto

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

// ActiveContractsServiceClient is the client API for ActiveContractsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActiveContractsServiceClient interface {
	// Returns a stream of the latest snapshot of active contracts.
	// If there are no active contracts, the stream returns a single GetActiveContractsResponse message with the offset at which the snapshot has been taken.
	// Clients SHOULD use the offset in the last GetActiveContractsResponse message to continue streaming transactions with the transaction service.
	// Clients SHOULD NOT assume that the set of active contracts they receive reflects the state at the ledger end.
	GetActiveContracts(ctx context.Context, in *GetActiveContractsRequest, opts ...grpc.CallOption) (ActiveContractsService_GetActiveContractsClient, error)
}

type activeContractsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActiveContractsServiceClient(cc grpc.ClientConnInterface) ActiveContractsServiceClient {
	return &activeContractsServiceClient{cc}
}

func (c *activeContractsServiceClient) GetActiveContracts(ctx context.Context, in *GetActiveContractsRequest, opts ...grpc.CallOption) (ActiveContractsService_GetActiveContractsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ActiveContractsService_ServiceDesc.Streams[0], "/com.daml.ledger.api.v1.ActiveContractsService/GetActiveContracts", opts...)
	if err != nil {
		return nil, err
	}
	x := &activeContractsServiceGetActiveContractsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActiveContractsService_GetActiveContractsClient interface {
	Recv() (*GetActiveContractsResponse, error)
	grpc.ClientStream
}

type activeContractsServiceGetActiveContractsClient struct {
	grpc.ClientStream
}

func (x *activeContractsServiceGetActiveContractsClient) Recv() (*GetActiveContractsResponse, error) {
	m := new(GetActiveContractsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActiveContractsServiceServer is the server API for ActiveContractsService service.
// All implementations must embed UnimplementedActiveContractsServiceServer
// for forward compatibility
type ActiveContractsServiceServer interface {
	// Returns a stream of the latest snapshot of active contracts.
	// If there are no active contracts, the stream returns a single GetActiveContractsResponse message with the offset at which the snapshot has been taken.
	// Clients SHOULD use the offset in the last GetActiveContractsResponse message to continue streaming transactions with the transaction service.
	// Clients SHOULD NOT assume that the set of active contracts they receive reflects the state at the ledger end.
	GetActiveContracts(*GetActiveContractsRequest, ActiveContractsService_GetActiveContractsServer) error
	mustEmbedUnimplementedActiveContractsServiceServer()
}

// UnimplementedActiveContractsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActiveContractsServiceServer struct {
}

func (UnimplementedActiveContractsServiceServer) GetActiveContracts(*GetActiveContractsRequest, ActiveContractsService_GetActiveContractsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActiveContracts not implemented")
}
func (UnimplementedActiveContractsServiceServer) mustEmbedUnimplementedActiveContractsServiceServer() {
}

// UnsafeActiveContractsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActiveContractsServiceServer will
// result in compilation errors.
type UnsafeActiveContractsServiceServer interface {
	mustEmbedUnimplementedActiveContractsServiceServer()
}

func RegisterActiveContractsServiceServer(s grpc.ServiceRegistrar, srv ActiveContractsServiceServer) {
	s.RegisterService(&ActiveContractsService_ServiceDesc, srv)
}

func _ActiveContractsService_GetActiveContracts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetActiveContractsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActiveContractsServiceServer).GetActiveContracts(m, &activeContractsServiceGetActiveContractsServer{stream})
}

type ActiveContractsService_GetActiveContractsServer interface {
	Send(*GetActiveContractsResponse) error
	grpc.ServerStream
}

type activeContractsServiceGetActiveContractsServer struct {
	grpc.ServerStream
}

func (x *activeContractsServiceGetActiveContractsServer) Send(m *GetActiveContractsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ActiveContractsService_ServiceDesc is the grpc.ServiceDesc for ActiveContractsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActiveContractsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.ActiveContractsService",
	HandlerType: (*ActiveContractsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActiveContracts",
			Handler:       _ActiveContractsService_GetActiveContracts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/daml/ledger/api/v1/active_contracts_service.proto",
}
