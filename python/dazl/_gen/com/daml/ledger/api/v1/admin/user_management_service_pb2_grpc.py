# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import user_management_service_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2

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
        + f' but the generated code in com/daml/ledger/api/v1/admin/user_management_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class UserManagementServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateUser = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/CreateUser',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.CreateUserRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.CreateUserResponse.FromString,
                _registered_method=True)
        self.GetUser = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/GetUser',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GetUserRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GetUserResponse.FromString,
                _registered_method=True)
        self.UpdateUser = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/UpdateUser',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserResponse.FromString,
                _registered_method=True)
        self.DeleteUser = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/DeleteUser',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.DeleteUserRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.DeleteUserResponse.FromString,
                _registered_method=True)
        self.ListUsers = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/ListUsers',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUsersRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUsersResponse.FromString,
                _registered_method=True)
        self.GrantUserRights = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/GrantUserRights',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GrantUserRightsRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GrantUserRightsResponse.FromString,
                _registered_method=True)
        self.RevokeUserRights = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/RevokeUserRights',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.RevokeUserRightsRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.RevokeUserRightsResponse.FromString,
                _registered_method=True)
        self.ListUserRights = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/ListUserRights',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUserRightsRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUserRightsResponse.FromString,
                _registered_method=True)
        self.UpdateUserIdentityProviderId = channel.unary_unary(
                '/com.daml.ledger.api.v1.admin.UserManagementService/UpdateUserIdentityProviderId',
                request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserIdentityProviderRequest.SerializeToString,
                response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserIdentityProviderResponse.FromString,
                _registered_method=True)


class UserManagementServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListUsers(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GrantUserRights(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RevokeUserRights(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListUserRights(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateUserIdentityProviderId(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserManagementServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateUser': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateUser,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.CreateUserRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.CreateUserResponse.SerializeToString,
            ),
            'GetUser': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUser,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GetUserRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GetUserResponse.SerializeToString,
            ),
            'UpdateUser': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateUser,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserResponse.SerializeToString,
            ),
            'DeleteUser': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteUser,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.DeleteUserRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.DeleteUserResponse.SerializeToString,
            ),
            'ListUsers': grpc.unary_unary_rpc_method_handler(
                    servicer.ListUsers,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUsersRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUsersResponse.SerializeToString,
            ),
            'GrantUserRights': grpc.unary_unary_rpc_method_handler(
                    servicer.GrantUserRights,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GrantUserRightsRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GrantUserRightsResponse.SerializeToString,
            ),
            'RevokeUserRights': grpc.unary_unary_rpc_method_handler(
                    servicer.RevokeUserRights,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.RevokeUserRightsRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.RevokeUserRightsResponse.SerializeToString,
            ),
            'ListUserRights': grpc.unary_unary_rpc_method_handler(
                    servicer.ListUserRights,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUserRightsRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUserRightsResponse.SerializeToString,
            ),
            'UpdateUserIdentityProviderId': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateUserIdentityProviderId,
                    request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserIdentityProviderRequest.FromString,
                    response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserIdentityProviderResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.daml.ledger.api.v1.admin.UserManagementService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.daml.ledger.api.v1.admin.UserManagementService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class UserManagementService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateUser(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/CreateUser',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.CreateUserRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.CreateUserResponse.FromString,
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
    def GetUser(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/GetUser',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GetUserRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GetUserResponse.FromString,
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
    def UpdateUser(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/UpdateUser',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserResponse.FromString,
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
    def DeleteUser(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/DeleteUser',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.DeleteUserRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.DeleteUserResponse.FromString,
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
    def ListUsers(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/ListUsers',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUsersRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUsersResponse.FromString,
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
    def GrantUserRights(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/GrantUserRights',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GrantUserRightsRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.GrantUserRightsResponse.FromString,
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
    def RevokeUserRights(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/RevokeUserRights',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.RevokeUserRightsRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.RevokeUserRightsResponse.FromString,
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
    def ListUserRights(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/ListUserRights',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUserRightsRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.ListUserRightsResponse.FromString,
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
    def UpdateUserIdentityProviderId(request,
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
            '/com.daml.ledger.api.v1.admin.UserManagementService/UpdateUserIdentityProviderId',
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserIdentityProviderRequest.SerializeToString,
            com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_user__management__service__pb2.UpdateUserIdentityProviderResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
