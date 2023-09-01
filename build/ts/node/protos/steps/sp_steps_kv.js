// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "steps/sp_steps_kv.proto" (package "protos.steps", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
// !!!!!!!!!!!!!!!!!!!!!!!!!!! IMPORTANT !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// 
// KV consists of two parts:
// 
// - KVStep
// - HostFuncKVExists
// 
// KVStep is used in PipelineSteps that will execute a specific KV request;
// the actual KV lookup is performed by the KV WASM func that calls out to
// HostFuncKVExists() that is a function exported by the SDK.
// 
// The HostFuncKVExists() function is needed because as of 08.30.2023, WASM does
// not have socket support, so we need to call out to the SDK to perform the
// actual KV API call.
// 
/**
 * Used by frontend when constructing a pipeline that contains a KV step that
 * performs a KVExists request.
 * protolint:disable:next ENUM_FIELD_NAMES_PREFIX
 *
 * @generated from protobuf enum protos.steps.KVExistsMode
 */
export var KVExistsMode;
(function (KVExistsMode) {
    /**
     * @generated from protobuf enum value: KV_EXISTS_MODE_UNSET = 0;
     */
    KVExistsMode[KVExistsMode["KV_EXISTS_MODE_UNSET"] = 0] = "KV_EXISTS_MODE_UNSET";
    /**
     * Will cause the KV lookup to use the key string as-is for the lookup
     *
     * @generated from protobuf enum value: KV_EXISTS_MODE_STATIC = 1;
     */
    KVExistsMode[KVExistsMode["KV_EXISTS_MODE_STATIC"] = 1] = "KV_EXISTS_MODE_STATIC";
    /**
     * DYNAMIC mode will cause the KV lookup WASM to use the key to lookup the
     * associated value and use the result for the key existence check.
     *
     * For example, if "key" in KVExistsRequest is set to "foo", KV WASM will do
     * the following:
     *
     * 1. Lookup the value of "foo" in the payload (which is "bar")
     * 2. Use "bar" as the "key" for the KV lookup
     *
     * @generated from protobuf enum value: KV_EXISTS_MODE_DYNAMIC = 2;
     */
    KVExistsMode[KVExistsMode["KV_EXISTS_MODE_DYNAMIC"] = 2] = "KV_EXISTS_MODE_DYNAMIC";
})(KVExistsMode || (KVExistsMode = {}));
// @generated message type with reflection information, may provide speed optimized methods
class KVExistsRequest$Type extends MessageType {
    constructor() {
        super("protos.steps.KVExistsRequest", [
            { no: 1, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "mode", kind: "enum", T: () => ["protos.steps.KVExistsMode", KVExistsMode] }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.KVExistsRequest
 */
export const KVExistsRequest = new KVExistsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVExistsResponse$Type extends MessageType {
    constructor() {
        super("protos.steps.KVExistsResponse", [
            { no: 1, name: "exists", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "is_error", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.KVExistsResponse
 */
export const KVExistsResponse = new KVExistsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVStep$Type extends MessageType {
    constructor() {
        super("protos.steps.KVStep", [
            { no: 1, name: "kv_exists_request", kind: "message", oneof: "request", T: () => KVExistsRequest }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.KVStep
 */
export const KVStep = new KVStep$Type();
//# sourceMappingURL=sp_steps_kv.js.map