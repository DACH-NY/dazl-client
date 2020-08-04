# Copyright (c) 2020 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import package_service_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2


class PackageServiceStub(object):
  """Allows clients to query the DAML-LF packages that are supported by the server.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListPackages = channel.unary_unary(
        '/com.daml.ledger.api.v1.PackageService/ListPackages',
        request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.ListPackagesRequest.SerializeToString,
        response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.ListPackagesResponse.FromString,
        )
    self.GetPackage = channel.unary_unary(
        '/com.daml.ledger.api.v1.PackageService/GetPackage',
        request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageRequest.SerializeToString,
        response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageResponse.FromString,
        )
    self.GetPackageStatus = channel.unary_unary(
        '/com.daml.ledger.api.v1.PackageService/GetPackageStatus',
        request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageStatusRequest.SerializeToString,
        response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageStatusResponse.FromString,
        )


class PackageServiceServicer(object):
  """Allows clients to query the DAML-LF packages that are supported by the server.
  """

  def ListPackages(self, request, context):
    """Returns the identifiers of all supported packages.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetPackage(self, request, context):
    """Returns the contents of a single package, or a ``NOT_FOUND`` error if the requested package is unknown.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetPackageStatus(self, request, context):
    """Returns the status of a single package.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_PackageServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListPackages': grpc.unary_unary_rpc_method_handler(
          servicer.ListPackages,
          request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.ListPackagesRequest.FromString,
          response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.ListPackagesResponse.SerializeToString,
      ),
      'GetPackage': grpc.unary_unary_rpc_method_handler(
          servicer.GetPackage,
          request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageRequest.FromString,
          response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageResponse.SerializeToString,
      ),
      'GetPackageStatus': grpc.unary_unary_rpc_method_handler(
          servicer.GetPackageStatus,
          request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageStatusRequest.FromString,
          response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_package__service__pb2.GetPackageStatusResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'com.daml.ledger.api.v1.PackageService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
