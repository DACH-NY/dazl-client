# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import config_management_service_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2


class ConfigManagementServiceStub(object):
    """Status: experimental interface, will change before it is deemed production
    ready

    The ledger configuration management service provides methods for the ledger administrator
    to change the current ledger configuration. The services provides methods to modify
    different aspects of the configuration.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetTimeModel = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.ConfigManagementService/GetTimeModel',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelResponse.FromString,
                )
        self.SetTimeModel = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.ConfigManagementService/SetTimeModel',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelResponse.FromString,
                )


class ConfigManagementServiceServicer(object):
    """Status: experimental interface, will change before it is deemed production
    ready

    The ledger configuration management service provides methods for the ledger administrator
    to change the current ledger configuration. The services provides methods to modify
    different aspects of the configuration.
    """

    def GetTimeModel(self, request, context):
        """Return the currently active time model and the current configuration generation.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetTimeModel(self, request, context):
        """Set the ledger time model.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ConfigManagementServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetTimeModel': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTimeModel,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelResponse.SerializeToString,
            ),
            'SetTimeModel': grpc.unary_unary_rpc_method_handler(
                    servicer.SetTimeModel,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.daml.ledger.api.v1.admin.ConfigManagementService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ConfigManagementService(object):
    """Status: experimental interface, will change before it is deemed production
    ready

    The ledger configuration management service provides methods for the ledger administrator
    to change the current ledger configuration. The services provides methods to modify
    different aspects of the configuration.
    """

    @staticmethod
    def GetTimeModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/com.daml.ledger.api.v1.admin.ConfigManagementService/GetTimeModel',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetTimeModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/com.daml.ledger.api.v1.admin.ConfigManagementService/SetTimeModel',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
