from shared import sp_shared_pb2 as _sp_shared_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class KVCreateHTTPRequest(_message.Message):
    __slots__ = ["kvs", "overwrite"]
    KVS_FIELD_NUMBER: _ClassVar[int]
    OVERWRITE_FIELD_NUMBER: _ClassVar[int]
    kvs: _containers.RepeatedCompositeFieldContainer[KVObject]
    overwrite: bool
    def __init__(self, kvs: _Optional[_Iterable[_Union[KVObject, _Mapping]]] = ..., overwrite: bool = ...) -> None: ...

class KVInstruction(_message.Message):
    __slots__ = ["action", "id", "object", "requested_at_unix_ts_nano_utc"]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    OBJECT_FIELD_NUMBER: _ClassVar[int]
    REQUESTED_AT_UNIX_TS_NANO_UTC_FIELD_NUMBER: _ClassVar[int]
    action: _sp_shared_pb2.KVAction
    id: str
    object: KVObject
    requested_at_unix_ts_nano_utc: int
    def __init__(self, id: _Optional[str] = ..., action: _Optional[_Union[_sp_shared_pb2.KVAction, str]] = ..., object: _Optional[_Union[KVObject, _Mapping]] = ..., requested_at_unix_ts_nano_utc: _Optional[int] = ...) -> None: ...

class KVObject(_message.Message):
    __slots__ = ["created_at_unix_ts_nano_utc", "key", "updated_at_unix_ts_nano_utc", "value"]
    CREATED_AT_UNIX_TS_NANO_UTC_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_UNIX_TS_NANO_UTC_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    created_at_unix_ts_nano_utc: int
    key: str
    updated_at_unix_ts_nano_utc: int
    value: bytes
    def __init__(self, key: _Optional[str] = ..., value: _Optional[bytes] = ..., created_at_unix_ts_nano_utc: _Optional[int] = ..., updated_at_unix_ts_nano_utc: _Optional[int] = ...) -> None: ...

class KVRequest(_message.Message):
    __slots__ = ["instructions", "overwrite"]
    INSTRUCTIONS_FIELD_NUMBER: _ClassVar[int]
    OVERWRITE_FIELD_NUMBER: _ClassVar[int]
    instructions: _containers.RepeatedCompositeFieldContainer[KVInstruction]
    overwrite: bool
    def __init__(self, instructions: _Optional[_Iterable[_Union[KVInstruction, _Mapping]]] = ..., overwrite: bool = ...) -> None: ...

class KVUpdateHTTPRequest(_message.Message):
    __slots__ = ["kvs"]
    KVS_FIELD_NUMBER: _ClassVar[int]
    kvs: _containers.RepeatedCompositeFieldContainer[KVObject]
    def __init__(self, kvs: _Optional[_Iterable[_Union[KVObject, _Mapping]]] = ...) -> None: ...
