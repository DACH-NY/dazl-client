// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: com/digitalasset/canton/participant/admin/v0/package_service.proto

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PackageServiceClient is the client API for PackageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageServiceClient interface {
	// List all Daml-LF archives on the participant node - return their hashes
	ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error)
	// Lists all the modules in package on the participant node
	ListPackageContents(ctx context.Context, in *ListPackageContentsRequest, opts ...grpc.CallOption) (*ListPackageContentsResponse, error)
	// List all DARs on the participant node - return their hashes and filenames
	ListDars(ctx context.Context, in *ListDarsRequest, opts ...grpc.CallOption) (*ListDarsResponse, error)
	// List content of a Dar
	ListDarContents(ctx context.Context, in *ListDarContentsRequest, opts ...grpc.CallOption) (*ListDarContentsResponse, error)
	// Upload a DAR file and all packages inside to the participant node
	UploadDar(ctx context.Context, in *UploadDarRequest, opts ...grpc.CallOption) (*UploadDarResponse, error)
	// Remove a package that is not vetted
	RemovePackage(ctx context.Context, in *RemovePackageRequest, opts ...grpc.CallOption) (*RemovePackageResponse, error)
	// Remove a DAR that is not needed
	RemoveDar(ctx context.Context, in *RemoveDarRequest, opts ...grpc.CallOption) (*RemoveDarResponse, error)
	// Obtain a DAR file by hash -- for inspection & download
	GetDar(ctx context.Context, in *GetDarRequest, opts ...grpc.CallOption) (*GetDarResponse, error)
	// Share a DAR with another participant
	Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List requests this participant has made to share DARs with another participant
	ListShareRequests(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShareRequestsResponse, error)
	// List offers to share a DAR that this participant has received
	ListShareOffers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShareOffersResponse, error)
	// Accept a DAR sharing offer (this will install the DAR into the participant)
	AcceptShareOffer(ctx context.Context, in *AcceptShareOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Reject a DAR sharing offer
	RejectShareOffer(ctx context.Context, in *RejectShareOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Add party to our DAR distribution whitelist
	WhitelistAdd(ctx context.Context, in *WhitelistChangeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Remove party from our DAR distribution whitelist
	WhitelistRemove(ctx context.Context, in *WhitelistChangeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List all parties currently on the whitelist
	WhitelistList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WhitelistListResponse, error)
}

type packageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageServiceClient(cc grpc.ClientConnInterface) PackageServiceClient {
	return &packageServiceClient{cc}
}

func (c *packageServiceClient) ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error) {
	out := new(ListPackagesResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/ListPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) ListPackageContents(ctx context.Context, in *ListPackageContentsRequest, opts ...grpc.CallOption) (*ListPackageContentsResponse, error) {
	out := new(ListPackageContentsResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/ListPackageContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) ListDars(ctx context.Context, in *ListDarsRequest, opts ...grpc.CallOption) (*ListDarsResponse, error) {
	out := new(ListDarsResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/ListDars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) ListDarContents(ctx context.Context, in *ListDarContentsRequest, opts ...grpc.CallOption) (*ListDarContentsResponse, error) {
	out := new(ListDarContentsResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/ListDarContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) UploadDar(ctx context.Context, in *UploadDarRequest, opts ...grpc.CallOption) (*UploadDarResponse, error) {
	out := new(UploadDarResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/UploadDar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) RemovePackage(ctx context.Context, in *RemovePackageRequest, opts ...grpc.CallOption) (*RemovePackageResponse, error) {
	out := new(RemovePackageResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/RemovePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) RemoveDar(ctx context.Context, in *RemoveDarRequest, opts ...grpc.CallOption) (*RemoveDarResponse, error) {
	out := new(RemoveDarResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/RemoveDar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) GetDar(ctx context.Context, in *GetDarRequest, opts ...grpc.CallOption) (*GetDarResponse, error) {
	out := new(GetDarResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/GetDar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) ListShareRequests(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShareRequestsResponse, error) {
	out := new(ListShareRequestsResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/ListShareRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) ListShareOffers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShareOffersResponse, error) {
	out := new(ListShareOffersResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/ListShareOffers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) AcceptShareOffer(ctx context.Context, in *AcceptShareOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/AcceptShareOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) RejectShareOffer(ctx context.Context, in *RejectShareOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/RejectShareOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) WhitelistAdd(ctx context.Context, in *WhitelistChangeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/WhitelistAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) WhitelistRemove(ctx context.Context, in *WhitelistChangeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/WhitelistRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) WhitelistList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WhitelistListResponse, error) {
	out := new(WhitelistListResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.canton.participant.admin.v0.PackageService/WhitelistList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageServiceServer is the server API for PackageService service.
// All implementations must embed UnimplementedPackageServiceServer
// for forward compatibility
type PackageServiceServer interface {
	// List all Daml-LF archives on the participant node - return their hashes
	ListPackages(context.Context, *ListPackagesRequest) (*ListPackagesResponse, error)
	// Lists all the modules in package on the participant node
	ListPackageContents(context.Context, *ListPackageContentsRequest) (*ListPackageContentsResponse, error)
	// List all DARs on the participant node - return their hashes and filenames
	ListDars(context.Context, *ListDarsRequest) (*ListDarsResponse, error)
	// List content of a Dar
	ListDarContents(context.Context, *ListDarContentsRequest) (*ListDarContentsResponse, error)
	// Upload a DAR file and all packages inside to the participant node
	UploadDar(context.Context, *UploadDarRequest) (*UploadDarResponse, error)
	// Remove a package that is not vetted
	RemovePackage(context.Context, *RemovePackageRequest) (*RemovePackageResponse, error)
	// Remove a DAR that is not needed
	RemoveDar(context.Context, *RemoveDarRequest) (*RemoveDarResponse, error)
	// Obtain a DAR file by hash -- for inspection & download
	GetDar(context.Context, *GetDarRequest) (*GetDarResponse, error)
	// Share a DAR with another participant
	Share(context.Context, *ShareRequest) (*emptypb.Empty, error)
	// List requests this participant has made to share DARs with another participant
	ListShareRequests(context.Context, *emptypb.Empty) (*ListShareRequestsResponse, error)
	// List offers to share a DAR that this participant has received
	ListShareOffers(context.Context, *emptypb.Empty) (*ListShareOffersResponse, error)
	// Accept a DAR sharing offer (this will install the DAR into the participant)
	AcceptShareOffer(context.Context, *AcceptShareOfferRequest) (*emptypb.Empty, error)
	// Reject a DAR sharing offer
	RejectShareOffer(context.Context, *RejectShareOfferRequest) (*emptypb.Empty, error)
	// Add party to our DAR distribution whitelist
	WhitelistAdd(context.Context, *WhitelistChangeRequest) (*emptypb.Empty, error)
	// Remove party from our DAR distribution whitelist
	WhitelistRemove(context.Context, *WhitelistChangeRequest) (*emptypb.Empty, error)
	// List all parties currently on the whitelist
	WhitelistList(context.Context, *emptypb.Empty) (*WhitelistListResponse, error)
	mustEmbedUnimplementedPackageServiceServer()
}

// UnimplementedPackageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPackageServiceServer struct {
}

func (UnimplementedPackageServiceServer) ListPackages(context.Context, *ListPackagesRequest) (*ListPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPackages not implemented")
}
func (UnimplementedPackageServiceServer) ListPackageContents(context.Context, *ListPackageContentsRequest) (*ListPackageContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPackageContents not implemented")
}
func (UnimplementedPackageServiceServer) ListDars(context.Context, *ListDarsRequest) (*ListDarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDars not implemented")
}
func (UnimplementedPackageServiceServer) ListDarContents(context.Context, *ListDarContentsRequest) (*ListDarContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDarContents not implemented")
}
func (UnimplementedPackageServiceServer) UploadDar(context.Context, *UploadDarRequest) (*UploadDarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDar not implemented")
}
func (UnimplementedPackageServiceServer) RemovePackage(context.Context, *RemovePackageRequest) (*RemovePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePackage not implemented")
}
func (UnimplementedPackageServiceServer) RemoveDar(context.Context, *RemoveDarRequest) (*RemoveDarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDar not implemented")
}
func (UnimplementedPackageServiceServer) GetDar(context.Context, *GetDarRequest) (*GetDarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDar not implemented")
}
func (UnimplementedPackageServiceServer) Share(context.Context, *ShareRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedPackageServiceServer) ListShareRequests(context.Context, *emptypb.Empty) (*ListShareRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShareRequests not implemented")
}
func (UnimplementedPackageServiceServer) ListShareOffers(context.Context, *emptypb.Empty) (*ListShareOffersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShareOffers not implemented")
}
func (UnimplementedPackageServiceServer) AcceptShareOffer(context.Context, *AcceptShareOfferRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptShareOffer not implemented")
}
func (UnimplementedPackageServiceServer) RejectShareOffer(context.Context, *RejectShareOfferRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectShareOffer not implemented")
}
func (UnimplementedPackageServiceServer) WhitelistAdd(context.Context, *WhitelistChangeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistAdd not implemented")
}
func (UnimplementedPackageServiceServer) WhitelistRemove(context.Context, *WhitelistChangeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistRemove not implemented")
}
func (UnimplementedPackageServiceServer) WhitelistList(context.Context, *emptypb.Empty) (*WhitelistListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistList not implemented")
}
func (UnimplementedPackageServiceServer) mustEmbedUnimplementedPackageServiceServer() {}

// UnsafePackageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageServiceServer will
// result in compilation errors.
type UnsafePackageServiceServer interface {
	mustEmbedUnimplementedPackageServiceServer()
}

func RegisterPackageServiceServer(s grpc.ServiceRegistrar, srv PackageServiceServer) {
	s.RegisterService(&PackageService_ServiceDesc, srv)
}

func _PackageService_ListPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/ListPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListPackages(ctx, req.(*ListPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_ListPackageContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPackageContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListPackageContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/ListPackageContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListPackageContents(ctx, req.(*ListPackageContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_ListDars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListDars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/ListDars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListDars(ctx, req.(*ListDarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_ListDarContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDarContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListDarContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/ListDarContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListDarContents(ctx, req.(*ListDarContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_UploadDar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).UploadDar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/UploadDar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).UploadDar(ctx, req.(*UploadDarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_RemovePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).RemovePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/RemovePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).RemovePackage(ctx, req.(*RemovePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_RemoveDar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).RemoveDar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/RemoveDar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).RemoveDar(ctx, req.(*RemoveDarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_GetDar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).GetDar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/GetDar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).GetDar(ctx, req.(*GetDarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).Share(ctx, req.(*ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_ListShareRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListShareRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/ListShareRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListShareRequests(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_ListShareOffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListShareOffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/ListShareOffers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListShareOffers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_AcceptShareOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptShareOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).AcceptShareOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/AcceptShareOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).AcceptShareOffer(ctx, req.(*AcceptShareOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_RejectShareOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectShareOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).RejectShareOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/RejectShareOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).RejectShareOffer(ctx, req.(*RejectShareOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_WhitelistAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).WhitelistAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/WhitelistAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).WhitelistAdd(ctx, req.(*WhitelistChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_WhitelistRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).WhitelistRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/WhitelistRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).WhitelistRemove(ctx, req.(*WhitelistChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_WhitelistList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).WhitelistList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.canton.participant.admin.v0.PackageService/WhitelistList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).WhitelistList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PackageService_ServiceDesc is the grpc.ServiceDesc for PackageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.canton.participant.admin.v0.PackageService",
	HandlerType: (*PackageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPackages",
			Handler:    _PackageService_ListPackages_Handler,
		},
		{
			MethodName: "ListPackageContents",
			Handler:    _PackageService_ListPackageContents_Handler,
		},
		{
			MethodName: "ListDars",
			Handler:    _PackageService_ListDars_Handler,
		},
		{
			MethodName: "ListDarContents",
			Handler:    _PackageService_ListDarContents_Handler,
		},
		{
			MethodName: "UploadDar",
			Handler:    _PackageService_UploadDar_Handler,
		},
		{
			MethodName: "RemovePackage",
			Handler:    _PackageService_RemovePackage_Handler,
		},
		{
			MethodName: "RemoveDar",
			Handler:    _PackageService_RemoveDar_Handler,
		},
		{
			MethodName: "GetDar",
			Handler:    _PackageService_GetDar_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _PackageService_Share_Handler,
		},
		{
			MethodName: "ListShareRequests",
			Handler:    _PackageService_ListShareRequests_Handler,
		},
		{
			MethodName: "ListShareOffers",
			Handler:    _PackageService_ListShareOffers_Handler,
		},
		{
			MethodName: "AcceptShareOffer",
			Handler:    _PackageService_AcceptShareOffer_Handler,
		},
		{
			MethodName: "RejectShareOffer",
			Handler:    _PackageService_RejectShareOffer_Handler,
		},
		{
			MethodName: "WhitelistAdd",
			Handler:    _PackageService_WhitelistAdd_Handler,
		},
		{
			MethodName: "WhitelistRemove",
			Handler:    _PackageService_WhitelistRemove_Handler,
		},
		{
			MethodName: "WhitelistList",
			Handler:    _PackageService_WhitelistList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/canton/participant/admin/v0/package_service.proto",
}