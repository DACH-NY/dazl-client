# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file

import builtins as _builtins, sys, typing as _typing

from google.protobuf.message import Message as _Message

from .experimental_features_pb2 import ExperimentalFeatures

if sys.version_info >= (3, 8):
    from typing import Literal as _L
else:
    from typing_extensions import Literal as _L

__all__ = [
    "GetLedgerApiVersionRequest",
    "GetLedgerApiVersionResponse",
    "FeaturesDescriptor",
    "UserManagementFeature",
]


class GetLedgerApiVersionRequest(_Message):
    ledger_id: _builtins.str
    def __init__(self, *, ledger_id: _typing.Optional[_builtins.str] = ...): ...
    def HasField(self, field_name: _L["ledger_id"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["ledger_id"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetLedgerApiVersionResponse(_Message):
    version: _builtins.str
    @property
    def features(self) -> FeaturesDescriptor: ...
    def __init__(self, *, version: _typing.Optional[_builtins.str] = ..., features: _typing.Optional[FeaturesDescriptor] = ...): ...
    def HasField(self, field_name: _L["version", "features"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["version", "features"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class FeaturesDescriptor(_Message):
    @property
    def user_management(self) -> UserManagementFeature: ...
    @property
    def experimental(self) -> ExperimentalFeatures: ...
    def __init__(self, *, user_management: _typing.Optional[UserManagementFeature] = ..., experimental: _typing.Optional[ExperimentalFeatures] = ...): ...
    def HasField(self, field_name: _L["user_management", "experimental"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["user_management", "experimental"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class UserManagementFeature(_Message):
    supported: _builtins.bool
    max_rights_per_user: _builtins.int
    max_users_page_size: _builtins.int
    def __init__(self, *, supported: _typing.Optional[_builtins.bool] = ..., max_rights_per_user: _typing.Optional[_builtins.int] = ..., max_users_page_size: _typing.Optional[_builtins.int] = ...): ...
    def HasField(self, field_name: _L["supported", "max_rights_per_user", "max_users_page_size"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["supported", "max_rights_per_user", "max_users_page_size"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...
