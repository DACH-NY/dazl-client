# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import update_service_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2

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
        + f' but the generated code in com/daml/ledger/api/v2/update_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class UpdateServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetUpdates = channel.unary_stream(
                '/com.daml.ledger.api.v2.UpdateService/GetUpdates',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesResponse.FromString,
                _registered_method=True)
        self.GetUpdateTrees = channel.unary_stream(
                '/com.daml.ledger.api.v2.UpdateService/GetUpdateTrees',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdateTreesResponse.FromString,
                _registered_method=True)
        self.GetTransactionTreeByEventId = channel.unary_unary(
                '/com.daml.ledger.api.v2.UpdateService/GetTransactionTreeByEventId',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByEventIdRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionTreeResponse.FromString,
                _registered_method=True)
        self.GetTransactionTreeById = channel.unary_unary(
                '/com.daml.ledger.api.v2.UpdateService/GetTransactionTreeById',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByIdRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionTreeResponse.FromString,
                _registered_method=True)
        self.GetTransactionByEventId = channel.unary_unary(
                '/com.daml.ledger.api.v2.UpdateService/GetTransactionByEventId',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByEventIdRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionResponse.FromString,
                _registered_method=True)
        self.GetTransactionById = channel.unary_unary(
                '/com.daml.ledger.api.v2.UpdateService/GetTransactionById',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByIdRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionResponse.FromString,
                _registered_method=True)


class UpdateServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetUpdates(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUpdateTrees(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTransactionTreeByEventId(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTransactionTreeById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTransactionByEventId(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTransactionById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UpdateServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetUpdates': grpc.unary_stream_rpc_method_handler(
                    servicer.GetUpdates,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesResponse.SerializeToString,
            ),
            'GetUpdateTrees': grpc.unary_stream_rpc_method_handler(
                    servicer.GetUpdateTrees,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdateTreesResponse.SerializeToString,
            ),
            'GetTransactionTreeByEventId': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTransactionTreeByEventId,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByEventIdRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionTreeResponse.SerializeToString,
            ),
            'GetTransactionTreeById': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTransactionTreeById,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByIdRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionTreeResponse.SerializeToString,
            ),
            'GetTransactionByEventId': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTransactionByEventId,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByEventIdRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionResponse.SerializeToString,
            ),
            'GetTransactionById': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTransactionById,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByIdRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.daml.ledger.api.v2.UpdateService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.daml.ledger.api.v2.UpdateService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class UpdateService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetUpdates(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/com.daml.ledger.api.v2.UpdateService/GetUpdates',
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesResponse.FromString,
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
    def GetUpdateTrees(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/com.daml.ledger.api.v2.UpdateService/GetUpdateTrees',
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdatesRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetUpdateTreesResponse.FromString,
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
    def GetTransactionTreeByEventId(request,
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
            '/com.daml.ledger.api.v2.UpdateService/GetTransactionTreeByEventId',
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByEventIdRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionTreeResponse.FromString,
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
    def GetTransactionTreeById(request,
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
            '/com.daml.ledger.api.v2.UpdateService/GetTransactionTreeById',
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByIdRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionTreeResponse.FromString,
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
    def GetTransactionByEventId(request,
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
            '/com.daml.ledger.api.v2.UpdateService/GetTransactionByEventId',
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByEventIdRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionResponse.FromString,
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
    def GetTransactionById(request,
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
            '/com.daml.ledger.api.v2.UpdateService/GetTransactionById',
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionByIdRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v2_dot_update__service__pb2.GetTransactionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
