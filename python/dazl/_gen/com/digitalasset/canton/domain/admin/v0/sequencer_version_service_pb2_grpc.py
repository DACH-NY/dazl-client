# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from ....protocol.v0 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2

GRPC_GENERATED_VERSION = '1.65.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.66.0'
SCHEDULED_RELEASE_DATE = 'August 6, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in com/digitalasset/canton/domain/admin/v0/sequencer_version_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class SequencerVersionServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Handshake = channel.unary_unary(
                '/com.digitalasset.canton.domain.admin.v0.SequencerVersionService/Handshake',
                request_serializer=com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2.Handshake.Request.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2.Handshake.Response.FromString,
                _registered_method=True)


class SequencerVersionServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Handshake(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SequencerVersionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Handshake': grpc.unary_unary_rpc_method_handler(
                    servicer.Handshake,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2.Handshake.Request.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2.Handshake.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.digitalasset.canton.domain.admin.v0.SequencerVersionService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.digitalasset.canton.domain.admin.v0.SequencerVersionService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SequencerVersionService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Handshake(request,
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
            '/com.digitalasset.canton.domain.admin.v0.SequencerVersionService/Handshake',
            com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2.Handshake.Request.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2.Handshake.Response.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
