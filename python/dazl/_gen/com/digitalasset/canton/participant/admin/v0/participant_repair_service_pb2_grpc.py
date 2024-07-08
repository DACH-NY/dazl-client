# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import participant_repair_service_pb2 as com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2


class ParticipantRepairServiceStub(object):
    """Moving ACS from one participant to another
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Download = channel.unary_stream(
                '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/Download',
                request_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.DownloadRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.AcsSnapshotChunk.FromString,
                )
        self.Upload = channel.stream_unary(
                '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/Upload',
                request_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.UploadRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.UploadResponse.FromString,
                )
        self.PurgeContracts = channel.unary_unary(
                '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/PurgeContracts',
                request_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.PurgeContractsRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.PurgeContractsResponse.FromString,
                )
        self.MigrateDomain = channel.unary_unary(
                '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/MigrateDomain',
                request_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.MigrateDomainRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.MigrateDomainResponse.FromString,
                )


class ParticipantRepairServiceServicer(object):
    """Moving ACS from one participant to another
    """

    def Download(self, request, context):
        """get contracts for a party
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Upload(self, request_iterator, context):
        """upload contracts for a party
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PurgeContracts(self, request, context):
        """purge contracts
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def MigrateDomain(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ParticipantRepairServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Download': grpc.unary_stream_rpc_method_handler(
                    servicer.Download,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.DownloadRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.AcsSnapshotChunk.SerializeToString,
            ),
            'Upload': grpc.stream_unary_rpc_method_handler(
                    servicer.Upload,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.UploadRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.UploadResponse.SerializeToString,
            ),
            'PurgeContracts': grpc.unary_unary_rpc_method_handler(
                    servicer.PurgeContracts,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.PurgeContractsRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.PurgeContractsResponse.SerializeToString,
            ),
            'MigrateDomain': grpc.unary_unary_rpc_method_handler(
                    servicer.MigrateDomain,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.MigrateDomainRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.MigrateDomainResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.digitalasset.canton.participant.admin.v0.ParticipantRepairService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ParticipantRepairService(object):
    """Moving ACS from one participant to another
    """

    @staticmethod
    def Download(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/Download',
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.DownloadRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.AcsSnapshotChunk.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Upload(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/Upload',
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.UploadRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.UploadResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PurgeContracts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/PurgeContracts',
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.PurgeContractsRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.PurgeContractsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def MigrateDomain(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/com.digitalasset.canton.participant.admin.v0.ParticipantRepairService/MigrateDomain',
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.MigrateDomainRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_participant__repair__service__pb2.MigrateDomainResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
