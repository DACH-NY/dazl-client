// Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: com/daml/ledger/api/v1/event_query_service.proto

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

// EventQueryServiceClient is the client API for EventQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventQueryServiceClient interface {
	// Get the create and the consuming exercise event for the contract with the provided ID.
	// No events will be returned for contracts that have been pruned because they
	// have already been archived before the latest pruning offset.
	GetEventsByContractId(ctx context.Context, in *GetEventsByContractIdRequest, opts ...grpc.CallOption) (*GetEventsByContractIdResponse, error)
	// Get all create and consuming exercise events for the contracts with the provided contract key.
	// Only events for unpruned contracts will be returned.
	// Matching events are delivered in reverse chronological order, i.e.,
	// the most recent events are delivered first.
	GetEventsByContractKey(ctx context.Context, in *GetEventsByContractKeyRequest, opts ...grpc.CallOption) (*GetEventsByContractKeyResponse, error)
}

type eventQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventQueryServiceClient(cc grpc.ClientConnInterface) EventQueryServiceClient {
	return &eventQueryServiceClient{cc}
}

func (c *eventQueryServiceClient) GetEventsByContractId(ctx context.Context, in *GetEventsByContractIdRequest, opts ...grpc.CallOption) (*GetEventsByContractIdResponse, error) {
	out := new(GetEventsByContractIdResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.EventQueryService/GetEventsByContractId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventQueryServiceClient) GetEventsByContractKey(ctx context.Context, in *GetEventsByContractKeyRequest, opts ...grpc.CallOption) (*GetEventsByContractKeyResponse, error) {
	out := new(GetEventsByContractKeyResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.EventQueryService/GetEventsByContractKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventQueryServiceServer is the server API for EventQueryService service.
// All implementations must embed UnimplementedEventQueryServiceServer
// for forward compatibility
type EventQueryServiceServer interface {
	// Get the create and the consuming exercise event for the contract with the provided ID.
	// No events will be returned for contracts that have been pruned because they
	// have already been archived before the latest pruning offset.
	GetEventsByContractId(context.Context, *GetEventsByContractIdRequest) (*GetEventsByContractIdResponse, error)
	// Get all create and consuming exercise events for the contracts with the provided contract key.
	// Only events for unpruned contracts will be returned.
	// Matching events are delivered in reverse chronological order, i.e.,
	// the most recent events are delivered first.
	GetEventsByContractKey(context.Context, *GetEventsByContractKeyRequest) (*GetEventsByContractKeyResponse, error)
	mustEmbedUnimplementedEventQueryServiceServer()
}

// UnimplementedEventQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventQueryServiceServer struct {
}

func (UnimplementedEventQueryServiceServer) GetEventsByContractId(context.Context, *GetEventsByContractIdRequest) (*GetEventsByContractIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsByContractId not implemented")
}
func (UnimplementedEventQueryServiceServer) GetEventsByContractKey(context.Context, *GetEventsByContractKeyRequest) (*GetEventsByContractKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsByContractKey not implemented")
}
func (UnimplementedEventQueryServiceServer) mustEmbedUnimplementedEventQueryServiceServer() {}

// UnsafeEventQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventQueryServiceServer will
// result in compilation errors.
type UnsafeEventQueryServiceServer interface {
	mustEmbedUnimplementedEventQueryServiceServer()
}

func RegisterEventQueryServiceServer(s grpc.ServiceRegistrar, srv EventQueryServiceServer) {
	s.RegisterService(&EventQueryService_ServiceDesc, srv)
}

func _EventQueryService_GetEventsByContractId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsByContractIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventQueryServiceServer).GetEventsByContractId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.EventQueryService/GetEventsByContractId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventQueryServiceServer).GetEventsByContractId(ctx, req.(*GetEventsByContractIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventQueryService_GetEventsByContractKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsByContractKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventQueryServiceServer).GetEventsByContractKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.EventQueryService/GetEventsByContractKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventQueryServiceServer).GetEventsByContractKey(ctx, req.(*GetEventsByContractKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventQueryService_ServiceDesc is the grpc.ServiceDesc for EventQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.EventQueryService",
	HandlerType: (*EventQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventsByContractId",
			Handler:    _EventQueryService_GetEventsByContractId_Handler,
		},
		{
			MethodName: "GetEventsByContractKey",
			Handler:    _EventQueryService_GetEventsByContractKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v1/event_query_service.proto",
}
