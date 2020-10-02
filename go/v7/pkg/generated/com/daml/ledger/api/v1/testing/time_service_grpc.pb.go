// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testing

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeServiceClient interface {
	// Returns a stream of time updates.
	// Always returns at least one response, where the first one is the current time.
	// Subsequent responses are emitted whenever the ledger server's time is updated.
	GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (TimeService_GetTimeClient, error)
	// Allows clients to change the ledger's clock in an atomic get-and-set operation.
	SetTime(ctx context.Context, in *SetTimeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type timeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeServiceClient(cc grpc.ClientConnInterface) TimeServiceClient {
	return &timeServiceClient{cc}
}

var timeServiceGetTimeStreamDesc = &grpc.StreamDesc{
	StreamName:    "GetTime",
	ServerStreams: true,
}

func (c *timeServiceClient) GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (TimeService_GetTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, timeServiceGetTimeStreamDesc, "/com.daml.ledger.api.v1.testing.TimeService/GetTime", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeServiceGetTimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TimeService_GetTimeClient interface {
	Recv() (*GetTimeResponse, error)
	grpc.ClientStream
}

type timeServiceGetTimeClient struct {
	grpc.ClientStream
}

func (x *timeServiceGetTimeClient) Recv() (*GetTimeResponse, error) {
	m := new(GetTimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var timeServiceSetTimeStreamDesc = &grpc.StreamDesc{
	StreamName: "SetTime",
}

func (c *timeServiceClient) SetTime(ctx context.Context, in *SetTimeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.testing.TimeService/SetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServiceService is the service API for TimeService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterTimeServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type TimeServiceService struct {
	// Returns a stream of time updates.
	// Always returns at least one response, where the first one is the current time.
	// Subsequent responses are emitted whenever the ledger server's time is updated.
	GetTime func(*GetTimeRequest, TimeService_GetTimeServer) error
	// Allows clients to change the ledger's clock in an atomic get-and-set operation.
	SetTime func(context.Context, *SetTimeRequest) (*empty.Empty, error)
}

func (s *TimeServiceService) getTime(_ interface{}, stream grpc.ServerStream) error {
	if s.GetTime == nil {
		return status.Errorf(codes.Unimplemented, "method GetTime not implemented")
	}
	m := new(GetTimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.GetTime(m, &timeServiceGetTimeServer{stream})
}
func (s *TimeServiceService) setTime(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.SetTime == nil {
		return nil, status.Errorf(codes.Unimplemented, "method SetTime not implemented")
	}
	in := new(SetTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.SetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.daml.ledger.api.v1.testing.TimeService/SetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SetTime(ctx, req.(*SetTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

type TimeService_GetTimeServer interface {
	Send(*GetTimeResponse) error
	grpc.ServerStream
}

type timeServiceGetTimeServer struct {
	grpc.ServerStream
}

func (x *timeServiceGetTimeServer) Send(m *GetTimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterTimeServiceService registers a service implementation with a gRPC server.
func RegisterTimeServiceService(s grpc.ServiceRegistrar, srv *TimeServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "com.daml.ledger.api.v1.testing.TimeService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SetTime",
				Handler:    srv.setTime,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "GetTime",
				Handler:       srv.getTime,
				ServerStreams: true,
			},
		},
		Metadata: "com/daml/ledger/api/v1/testing/time_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewTimeServiceService creates a new TimeServiceService containing the
// implemented methods of the TimeService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewTimeServiceService(s interface{}) *TimeServiceService {
	ns := &TimeServiceService{}
	if h, ok := s.(interface {
		GetTime(*GetTimeRequest, TimeService_GetTimeServer) error
	}); ok {
		ns.GetTime = h.GetTime
	}
	if h, ok := s.(interface {
		SetTime(context.Context, *SetTimeRequest) (*empty.Empty, error)
	}); ok {
		ns.SetTime = h.SetTime
	}
	return ns
}

// UnstableTimeServiceService is the service API for TimeService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableTimeServiceService interface {
	// Returns a stream of time updates.
	// Always returns at least one response, where the first one is the current time.
	// Subsequent responses are emitted whenever the ledger server's time is updated.
	GetTime(*GetTimeRequest, TimeService_GetTimeServer) error
	// Allows clients to change the ledger's clock in an atomic get-and-set operation.
	SetTime(context.Context, *SetTimeRequest) (*empty.Empty, error)
}