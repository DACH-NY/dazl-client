# Copyright (c) 2017-2022 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file

import builtins as _builtins, sys, typing as _typing

from google.protobuf.any_pb2 import Any
from google.protobuf.internal.containers import RepeatedCompositeFieldContainer, RepeatedScalarFieldContainer
from google.protobuf.message import Message as _Message
from google.protobuf.wrappers_pb2 import StringValue
from google.rpc.status_pb2 import Status

from .contract_metadata_pb2 import ContractMetadata
from .value_pb2 import Identifier, Record, Value

if sys.version_info >= (3, 8):
    from typing import Literal as _L
else:
    from typing_extensions import Literal as _L

__all__ = [
    "Event",
    "CreatedEvent",
    "InterfaceView",
    "ArchivedEvent",
    "ExercisedEvent",
]


class Event(_Message):
    @property
    def created(self) -> CreatedEvent: ...
    @property
    def archived(self) -> ArchivedEvent: ...
    @_typing.overload
    def __init__(self): ...
    @_typing.overload
    def __init__(self, *, created: CreatedEvent = ...): ...
    @_typing.overload
    def __init__(self, *, archived: ArchivedEvent = ...): ...
    def HasField(self, field_name: _L["event", "created", "archived"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["event", "created", "archived"]) -> None: ...
    def WhichOneof(self, oneof_group: _L["event"]) -> _L[None, "created", "archived"]: ...

class CreatedEvent(_Message):
    event_id: _builtins.str
    contract_id: _builtins.str
    @property
    def template_id(self) -> Identifier: ...
    @property
    def contract_key(self) -> Value: ...
    @property
    def create_arguments(self) -> Record: ...
    @property
    def create_arguments_blob(self) -> Any: ...
    @property
    def interface_views(self) -> RepeatedCompositeFieldContainer[InterfaceView]: ...
    @property
    def witness_parties(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    @property
    def signatories(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    @property
    def observers(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    @property
    def agreement_text(self) -> StringValue: ...
    @property
    def metadata(self) -> ContractMetadata: ...
    def __init__(self, *, event_id: _typing.Optional[_builtins.str] = ..., contract_id: _typing.Optional[_builtins.str] = ..., template_id: _typing.Optional[Identifier] = ..., contract_key: _typing.Optional[Value] = ..., create_arguments: _typing.Optional[Record] = ..., create_arguments_blob: _typing.Optional[Any] = ..., interface_views: _typing.Optional[_typing.Iterable[InterfaceView]] = ..., witness_parties: _typing.Optional[_typing.Iterable[_builtins.str]] = ..., signatories: _typing.Optional[_typing.Iterable[_builtins.str]] = ..., observers: _typing.Optional[_typing.Iterable[_builtins.str]] = ..., agreement_text: _typing.Optional[StringValue] = ..., metadata: _typing.Optional[ContractMetadata] = ...): ...
    def HasField(self, field_name: _L["event_id", "contract_id", "template_id", "contract_key", "create_arguments", "create_arguments_blob", "interface_views", "witness_parties", "signatories", "observers", "agreement_text", "metadata"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["event_id", "contract_id", "template_id", "contract_key", "create_arguments", "create_arguments_blob", "interface_views", "witness_parties", "signatories", "observers", "agreement_text", "metadata"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class InterfaceView(_Message):
    @property
    def interface_id(self) -> Identifier: ...
    @property
    def view_status(self) -> Status: ...
    @property
    def view_value(self) -> Record: ...
    def __init__(self, *, interface_id: _typing.Optional[Identifier] = ..., view_status: _typing.Optional[Status] = ..., view_value: _typing.Optional[Record] = ...): ...
    def HasField(self, field_name: _L["interface_id", "view_status", "view_value"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["interface_id", "view_status", "view_value"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class ArchivedEvent(_Message):
    event_id: _builtins.str
    contract_id: _builtins.str
    @property
    def template_id(self) -> Identifier: ...
    @property
    def witness_parties(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    def __init__(self, *, event_id: _typing.Optional[_builtins.str] = ..., contract_id: _typing.Optional[_builtins.str] = ..., template_id: _typing.Optional[Identifier] = ..., witness_parties: _typing.Optional[_typing.Iterable[_builtins.str]] = ...): ...
    def HasField(self, field_name: _L["event_id", "contract_id", "template_id", "witness_parties"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["event_id", "contract_id", "template_id", "witness_parties"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class ExercisedEvent(_Message):
    event_id: _builtins.str
    contract_id: _builtins.str
    @property
    def template_id(self) -> Identifier: ...
    @property
    def interface_id(self) -> Identifier: ...
    choice: _builtins.str
    @property
    def choice_argument(self) -> Value: ...
    @property
    def acting_parties(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    consuming: _builtins.bool
    @property
    def witness_parties(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    @property
    def child_event_ids(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    @property
    def exercise_result(self) -> Value: ...
    def __init__(self, *, event_id: _typing.Optional[_builtins.str] = ..., contract_id: _typing.Optional[_builtins.str] = ..., template_id: _typing.Optional[Identifier] = ..., interface_id: _typing.Optional[Identifier] = ..., choice: _typing.Optional[_builtins.str] = ..., choice_argument: _typing.Optional[Value] = ..., acting_parties: _typing.Optional[_typing.Iterable[_builtins.str]] = ..., consuming: _typing.Optional[_builtins.bool] = ..., witness_parties: _typing.Optional[_typing.Iterable[_builtins.str]] = ..., child_event_ids: _typing.Optional[_typing.Iterable[_builtins.str]] = ..., exercise_result: _typing.Optional[Value] = ...): ...
    def HasField(self, field_name: _L["event_id", "contract_id", "template_id", "interface_id", "choice", "choice_argument", "acting_parties", "consuming", "witness_parties", "child_event_ids", "exercise_result"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["event_id", "contract_id", "template_id", "interface_id", "choice", "choice_argument", "acting_parties", "consuming", "witness_parties", "child_event_ids", "exercise_result"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...
