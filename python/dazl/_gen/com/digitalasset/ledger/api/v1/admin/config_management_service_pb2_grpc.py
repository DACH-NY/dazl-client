# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import config_management_service_pb2 as com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2


class ConfigManagementServiceStub(object):
  """Status: experimental interface, will change before it is deemed production
  ready

  Ledger configuration management service provides methods for the ledger administrator
  to change the current ledger configuration. The services provides methods to modify
  different aspects of the configuration.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetTimeModel = channel.unary_unary(
        '/com.digitalasset.ledger.api.v1.admin.ConfigManagementService/GetTimeModel',
        request_serializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelRequest.SerializeToString,
        response_deserializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelResponse.FromString,
        )
    self.SetTimeModel = channel.unary_unary(
        '/com.digitalasset.ledger.api.v1.admin.ConfigManagementService/SetTimeModel',
        request_serializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelRequest.SerializeToString,
        response_deserializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelResponse.FromString,
        )


class ConfigManagementServiceServicer(object):
  """Status: experimental interface, will change before it is deemed production
  ready

  Ledger configuration management service provides methods for the ledger administrator
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
    In case of failure this method responds with:
    - INVALID_ARGUMENT if arguments are invalid, or the provided configuration generation
    does not match the current active configuration generation. The caller is expected
    to retry by again fetching current time model using 'GetTimeModel', applying changes
    and resubmitting.
    - UNIMPLEMENTED if this method is not supported by the backing ledger.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ConfigManagementServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetTimeModel': grpc.unary_unary_rpc_method_handler(
          servicer.GetTimeModel,
          request_deserializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelRequest.FromString,
          response_serializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.GetTimeModelResponse.SerializeToString,
      ),
      'SetTimeModel': grpc.unary_unary_rpc_method_handler(
          servicer.SetTimeModel,
          request_deserializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelRequest.FromString,
          response_serializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_admin_dot_config__management__service__pb2.SetTimeModelResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'com.digitalasset.ledger.api.v1.admin.ConfigManagementService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
