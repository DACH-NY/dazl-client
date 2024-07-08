# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import party_name_management_pb2 as com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_party__name__management__pb2


class PartyNameManagementServiceStub(object):
    """*
    Local participant service allowing to set the display name for a party

    The display name is a local property to the participant. The participant is encouraged to perform
    a Daml based KYC process and add some automation which will update the display names based
    on the desired update rules.

    As such, this function here just offers the bare functionality to perform this.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.setPartyDisplayName = channel.unary_unary(
                '/com.digitalasset.canton.participant.admin.v0.PartyNameManagementService/setPartyDisplayName',
                request_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_party__name__management__pb2.SetPartyDisplayNameRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_party__name__management__pb2.SetPartyDisplayNameResponse.FromString,
                )


class PartyNameManagementServiceServicer(object):
    """*
    Local participant service allowing to set the display name for a party

    The display name is a local property to the participant. The participant is encouraged to perform
    a Daml based KYC process and add some automation which will update the display names based
    on the desired update rules.

    As such, this function here just offers the bare functionality to perform this.
    """

    def setPartyDisplayName(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PartyNameManagementServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'setPartyDisplayName': grpc.unary_unary_rpc_method_handler(
                    servicer.setPartyDisplayName,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_party__name__management__pb2.SetPartyDisplayNameRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_party__name__management__pb2.SetPartyDisplayNameResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.digitalasset.canton.participant.admin.v0.PartyNameManagementService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PartyNameManagementService(object):
    """*
    Local participant service allowing to set the display name for a party

    The display name is a local property to the participant. The participant is encouraged to perform
    a Daml based KYC process and add some automation which will update the display names based
    on the desired update rules.

    As such, this function here just offers the bare functionality to perform this.
    """

    @staticmethod
    def setPartyDisplayName(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/com.digitalasset.canton.participant.admin.v0.PartyNameManagementService/setPartyDisplayName',
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_party__name__management__pb2.SetPartyDisplayNameRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_participant_dot_admin_dot_v0_dot_party__name__management__pb2.SetPartyDisplayNameResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
