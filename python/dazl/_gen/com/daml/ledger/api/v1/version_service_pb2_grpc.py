# Copyright (c) 2017-2021 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import (
    version_service_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_version__service__pb2,
)


class VersionServiceStub(object):
    """Allows clients to retrieve information about the ledger API version"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetLedgerApiVersion = channel.unary_unary(
            "/com.daml.ledger.api.v1.VersionService/GetLedgerApiVersion",
            request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_version__service__pb2.GetLedgerApiVersionRequest.SerializeToString,
            response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_version__service__pb2.GetLedgerApiVersionResponse.FromString,
        )


class VersionServiceServicer(object):
    """Allows clients to retrieve information about the ledger API version"""

    def GetLedgerApiVersion(self, request, context):
        """Read the Ledger API version"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_VersionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "GetLedgerApiVersion": grpc.unary_unary_rpc_method_handler(
            servicer.GetLedgerApiVersion,
            request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_version__service__pb2.GetLedgerApiVersionRequest.FromString,
            response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_version__service__pb2.GetLedgerApiVersionResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "com.daml.ledger.api.v1.VersionService", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))


# This class is part of an EXPERIMENTAL API.
class VersionService(object):
    """Allows clients to retrieve information about the ledger API version"""

    @staticmethod
    def GetLedgerApiVersion(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/com.daml.ledger.api.v1.VersionService/GetLedgerApiVersion",
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_version__service__pb2.GetLedgerApiVersionRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_version__service__pb2.GetLedgerApiVersionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )
