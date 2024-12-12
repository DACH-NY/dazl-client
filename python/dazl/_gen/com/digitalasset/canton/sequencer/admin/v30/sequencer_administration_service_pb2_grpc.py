# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import sequencer_administration_service_pb2 as com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2

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
        + f' but the generated code in com/digitalasset/canton/sequencer/admin/v30/sequencer_administration_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class SequencerAdministrationServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.PruningStatus = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/PruningStatus',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.PruningStatusRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.PruningStatusResponse.FromString,
                _registered_method=True)
        self.TrafficControlState = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/TrafficControlState',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.TrafficControlStateRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.TrafficControlStateResponse.FromString,
                _registered_method=True)
        self.SetTrafficPurchased = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/SetTrafficPurchased',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SetTrafficPurchasedRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SetTrafficPurchasedResponse.FromString,
                _registered_method=True)
        self.Snapshot = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/Snapshot',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SnapshotRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SnapshotResponse.FromString,
                _registered_method=True)
        self.OnboardingState = channel.unary_stream(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/OnboardingState',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.OnboardingStateRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.OnboardingStateResponse.FromString,
                _registered_method=True)
        self.DisableMember = channel.unary_unary(
                '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/DisableMember',
                request_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.DisableMemberRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.DisableMemberResponse.FromString,
                _registered_method=True)


class SequencerAdministrationServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def PruningStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TrafficControlState(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetTrafficPurchased(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Snapshot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def OnboardingState(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DisableMember(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SequencerAdministrationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'PruningStatus': grpc.unary_unary_rpc_method_handler(
                    servicer.PruningStatus,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.PruningStatusRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.PruningStatusResponse.SerializeToString,
            ),
            'TrafficControlState': grpc.unary_unary_rpc_method_handler(
                    servicer.TrafficControlState,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.TrafficControlStateRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.TrafficControlStateResponse.SerializeToString,
            ),
            'SetTrafficPurchased': grpc.unary_unary_rpc_method_handler(
                    servicer.SetTrafficPurchased,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SetTrafficPurchasedRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SetTrafficPurchasedResponse.SerializeToString,
            ),
            'Snapshot': grpc.unary_unary_rpc_method_handler(
                    servicer.Snapshot,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SnapshotRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SnapshotResponse.SerializeToString,
            ),
            'OnboardingState': grpc.unary_stream_rpc_method_handler(
                    servicer.OnboardingState,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.OnboardingStateRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.OnboardingStateResponse.SerializeToString,
            ),
            'DisableMember': grpc.unary_unary_rpc_method_handler(
                    servicer.DisableMember,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.DisableMemberRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.DisableMemberResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SequencerAdministrationService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def PruningStatus(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/PruningStatus',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.PruningStatusRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.PruningStatusResponse.FromString,
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
    def TrafficControlState(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/TrafficControlState',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.TrafficControlStateRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.TrafficControlStateResponse.FromString,
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
    def SetTrafficPurchased(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/SetTrafficPurchased',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SetTrafficPurchasedRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SetTrafficPurchasedResponse.FromString,
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
    def Snapshot(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/Snapshot',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SnapshotRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.SnapshotResponse.FromString,
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
    def OnboardingState(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/OnboardingState',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.OnboardingStateRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.OnboardingStateResponse.FromString,
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
    def DisableMember(request,
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
            '/com.digitalasset.canton.sequencer.admin.v30.SequencerAdministrationService/DisableMember',
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.DisableMemberRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_sequencer_dot_admin_dot_v30_dot_sequencer__administration__service__pb2.DisableMemberResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)