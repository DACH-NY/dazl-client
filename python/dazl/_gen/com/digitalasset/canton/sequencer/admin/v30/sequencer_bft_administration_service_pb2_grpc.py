# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import sequencer_bft_administration_service_pb2 as com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2

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
        + f' but the generated code in com/digitalasset/canton/sequencer/admin/v30/sequencer_bft_administration_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class SequencerBftAdministrationServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AddPeerEndpoint = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/AddPeerEndpoint',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.AddPeerEndpointRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.AddPeerEndpointResponse.FromString,
                _registered_method=True)
        self.RemovePeerEndpoint = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/RemovePeerEndpoint',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.RemovePeerEndpointRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.RemovePeerEndpointResponse.FromString,
                _registered_method=True)
        self.GetPeerNetworkStatus = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/GetPeerNetworkStatus',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetPeerNetworkStatusRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetPeerNetworkStatusResponse.FromString,
                _registered_method=True)
        self.GetOrderingTopology = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/GetOrderingTopology',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetOrderingTopologyRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetOrderingTopologyResponse.FromString,
                _registered_method=True)


class SequencerBftAdministrationServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def AddPeerEndpoint(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemovePeerEndpoint(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPeerNetworkStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOrderingTopology(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SequencerBftAdministrationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'AddPeerEndpoint': grpc.unary_unary_rpc_method_handler(
                    servicer.AddPeerEndpoint,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.AddPeerEndpointRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.AddPeerEndpointResponse.SerializeToString,
            ),
            'RemovePeerEndpoint': grpc.unary_unary_rpc_method_handler(
                    servicer.RemovePeerEndpoint,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.RemovePeerEndpointRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.RemovePeerEndpointResponse.SerializeToString,
            ),
            'GetPeerNetworkStatus': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPeerNetworkStatus,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetPeerNetworkStatusRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetPeerNetworkStatusResponse.SerializeToString,
            ),
            'GetOrderingTopology': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOrderingTopology,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetOrderingTopologyRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetOrderingTopologyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SequencerBftAdministrationService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def AddPeerEndpoint(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/AddPeerEndpoint',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.AddPeerEndpointRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.AddPeerEndpointResponse.FromString,
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
    def RemovePeerEndpoint(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/RemovePeerEndpoint',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.RemovePeerEndpointRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.RemovePeerEndpointResponse.FromString,
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
    def GetPeerNetworkStatus(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/GetPeerNetworkStatus',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetPeerNetworkStatusRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetPeerNetworkStatusResponse.FromString,
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
    def GetOrderingTopology(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerBftAdministrationService/GetOrderingTopology',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetOrderingTopologyRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__bft__administration__service__pb2.GetOrderingTopologyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
