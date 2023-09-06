// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_kv.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
/**
 * KVAction is a shared type that is used for protos.KVCommand and protos.KVStep.
 * Note that only a subset of actions are used for protos.KVCommand (CREATE,
 * UPDATE, DELETE, DELETE_ALL) while protos.KVStep uses most of them.
 *
 * protolint:disable:next ENUM_FIELD_NAMES_PREFIX
 *
 * @generated from protobuf enum protos.KVAction
 */
export var KVAction;
(function (KVAction) {
    /**
     * @generated from protobuf enum value: KV_ACTION_UNSET = 0;
     */
    KVAction[KVAction["KV_ACTION_UNSET"] = 0] = "KV_ACTION_UNSET";
    /**
     * @generated from protobuf enum value: KV_ACTION_GET = 1;
     */
    KVAction[KVAction["KV_ACTION_GET"] = 1] = "KV_ACTION_GET";
    /**
     * @generated from protobuf enum value: KV_ACTION_CREATE = 2;
     */
    KVAction[KVAction["KV_ACTION_CREATE"] = 2] = "KV_ACTION_CREATE";
    /**
     * @generated from protobuf enum value: KV_ACTION_UPDATE = 3;
     */
    KVAction[KVAction["KV_ACTION_UPDATE"] = 3] = "KV_ACTION_UPDATE";
    /**
     * @generated from protobuf enum value: KV_ACTION_EXISTS = 4;
     */
    KVAction[KVAction["KV_ACTION_EXISTS"] = 4] = "KV_ACTION_EXISTS";
    /**
     * @generated from protobuf enum value: KV_ACTION_DELETE = 5;
     */
    KVAction[KVAction["KV_ACTION_DELETE"] = 5] = "KV_ACTION_DELETE";
    /**
     * @generated from protobuf enum value: KV_ACTION_DELETE_ALL = 6;
     */
    KVAction[KVAction["KV_ACTION_DELETE_ALL"] = 6] = "KV_ACTION_DELETE_ALL";
})(KVAction || (KVAction = {}));
// @generated message type with reflection information, may provide speed optimized methods
class KVObject$Type extends MessageType {
    constructor() {
        super("protos.KVObject", [
            { no: 1, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 3, name: "created_at_unix_ts_nano_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "updated_at_unix_ts_nano_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVObject
 */
export const KVObject = new KVObject$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVInstruction$Type extends MessageType {
    constructor() {
        super("protos.KVInstruction", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "action", kind: "enum", T: () => ["protos.KVAction", KVAction] },
            { no: 3, name: "object", kind: "message", T: () => KVObject },
            { no: 4, name: "requested_at_unix_ts_nano_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVInstruction
 */
export const KVInstruction = new KVInstruction$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVRequest$Type extends MessageType {
    constructor() {
        super("protos.KVRequest", [
            { no: 1, name: "instructions", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KVInstruction },
            { no: 2, name: "overwrite", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVRequest
 */
export const KVRequest = new KVRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVCreateHTTPRequest$Type extends MessageType {
    constructor() {
        super("protos.KVCreateHTTPRequest", [
            { no: 1, name: "kvs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KVObject },
            { no: 2, name: "overwrite", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVCreateHTTPRequest
 */
export const KVCreateHTTPRequest = new KVCreateHTTPRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVUpdateHTTPRequest$Type extends MessageType {
    constructor() {
        super("protos.KVUpdateHTTPRequest", [
            { no: 1, name: "kvs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KVObject }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVUpdateHTTPRequest
 */
export const KVUpdateHTTPRequest = new KVUpdateHTTPRequest$Type();
//# sourceMappingURL=sp_kv.js.map