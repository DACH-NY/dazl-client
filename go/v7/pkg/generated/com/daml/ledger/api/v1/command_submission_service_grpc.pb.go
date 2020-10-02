// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// CommandSubmissionServiceClient is the client API for CommandSubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandSubmissionServiceClient interface {
	// Submit a single composite command.
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type commandSubmissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandSubmissionServiceClient(cc grpc.ClientConnInterface) CommandSubmissionServiceClient {
	return &commandSubmissionServiceClient{cc}
}

var commandSubmissionServiceSubmitStreamDesc = &grpc.StreamDesc{
	StreamName: "Submit",
}

func (c *commandSubmissionServiceClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.CommandSubmissionService/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandSubmissionServiceService is the service API for CommandSubmissionService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterCommandSubmissionServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type CommandSubmissionServiceService struct {
	// Submit a single composite command.
	Submit func(context.Context, *SubmitRequest) (*empty.Empty, error)
}

func (s *CommandSubmissionServiceService) submit(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.Submit == nil {
		return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
	}
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.daml.ledger.api.v1.CommandSubmissionService/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterCommandSubmissionServiceService registers a service implementation with a gRPC server.
func RegisterCommandSubmissionServiceService(s grpc.ServiceRegistrar, srv *CommandSubmissionServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "com.daml.ledger.api.v1.CommandSubmissionService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Submit",
				Handler:    srv.submit,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "com/daml/ledger/api/v1/command_submission_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewCommandSubmissionServiceService creates a new CommandSubmissionServiceService containing the
// implemented methods of the CommandSubmissionService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewCommandSubmissionServiceService(s interface{}) *CommandSubmissionServiceService {
	ns := &CommandSubmissionServiceService{}
	if h, ok := s.(interface {
		Submit(context.Context, *SubmitRequest) (*empty.Empty, error)
	}); ok {
		ns.Submit = h.Submit
	}
	return ns
}

// UnstableCommandSubmissionServiceService is the service API for CommandSubmissionService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableCommandSubmissionServiceService interface {
	// Submit a single composite command.
	Submit(context.Context, *SubmitRequest) (*empty.Empty, error)
}