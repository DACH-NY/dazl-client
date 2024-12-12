# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import sequencer_authentication_service_pb2 as com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2

GRPC_GENERATED_VERSION = '1.67.1'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in com/digitalasset/canton/domain/api/v30/sequencer_authentication_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class SequencerAuthenticationServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Challenge = channel.unary_unary(
                '/com.digitalasset.canton.domain.api.v30.SequencerAuthenticationService/Challenge',
                request_serializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.ChallengeRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.ChallengeResponse.FromString,
                _registered_method=True)
        self.Authenticate = channel.unary_unary(
                '/com.digitalasset.canton.domain.api.v30.SequencerAuthenticationService/Authenticate',
                request_serializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.AuthenticateRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.AuthenticateResponse.FromString,
                _registered_method=True)


class SequencerAuthenticationServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Challenge(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Authenticate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SequencerAuthenticationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Challenge': grpc.unary_unary_rpc_method_handler(
                    servicer.Challenge,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.ChallengeRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.ChallengeResponse.SerializeToString,
            ),
            'Authenticate': grpc.unary_unary_rpc_method_handler(
                    servicer.Authenticate,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.AuthenticateRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.AuthenticateResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.digitalasset.canton.domain.api.v30.SequencerAuthenticationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.digitalasset.canton.domain.api.v30.SequencerAuthenticationService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SequencerAuthenticationService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Challenge(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.domain.api.v30.SequencerAuthenticationService/Challenge',
            com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.ChallengeRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.ChallengeResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Authenticate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.domain.api.v30.SequencerAuthenticationService/Authenticate',
            com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.AuthenticateRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v30_dot_sequencer__authentication__service__pb2.SequencerAuthentication.AuthenticateResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
