import { WireType } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { ValidJSONStep } from "./steps/sp_steps_valid_json.js";
import { InferSchemaStep } from "./steps/sp_steps_inferschema.js";
import { KVStep } from "./steps/sp_steps_kv.js";
import { HttpRequestStep } from "./steps/sp_steps_httprequest.js";
import { CustomStep } from "./steps/sp_steps_custom.js";
import { DecodeStep } from "./steps/sp_steps_decode.js";
import { EncodeStep } from "./steps/sp_steps_encode.js";
import { TransformStep } from "./steps/sp_steps_transform.js";
import { DetectiveStep } from "./steps/sp_steps_detective.js";
import { NotificationConfig } from "./sp_notify.js";
/**
 * A condition defines how the SDK should handle a step response -- should it
 * continue executing the pipeline, should it abort, should it notify the server?
 * Each step can have multiple conditions.
 *
 * @generated from protobuf enum protos.PipelineStepCondition
 */
export var PipelineStepCondition;
(function (PipelineStepCondition) {
    /**
     * @generated from protobuf enum value: PIPELINE_STEP_CONDITION_UNSET = 0;
     */
    PipelineStepCondition[PipelineStepCondition["UNSET"] = 0] = "UNSET";
    /**
     * Abort executing the current pipeline AND continue executing any other pipelines
     *
     * @generated from protobuf enum value: PIPELINE_STEP_CONDITION_ABORT_CURRENT = 1;
     */
    PipelineStepCondition[PipelineStepCondition["ABORT_CURRENT"] = 1] = "ABORT_CURRENT";
    /**
     * Notify the server about the step condition
     *
     * @generated from protobuf enum value: PIPELINE_STEP_CONDITION_NOTIFY = 2;
     */
    PipelineStepCondition[PipelineStepCondition["NOTIFY"] = 2] = "NOTIFY";
    /**
     * Abort executing ALL pipelines
     *
     * @generated from protobuf enum value: PIPELINE_STEP_CONDITION_ABORT_ALL = 3;
     */
    PipelineStepCondition[PipelineStepCondition["ABORT_ALL"] = 3] = "ABORT_ALL";
})(PipelineStepCondition || (PipelineStepCondition = {}));
// @generated message type with reflection information, may provide speed optimized methods
class Pipeline$Type extends MessageType {
    constructor() {
        super("protos.Pipeline", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "steps", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => PipelineStep },
            { no: 4, name: "_notification_configs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => NotificationConfig }
        ]);
    }
    create(value) {
        const message = { id: "", name: "", steps: [], NotificationConfigs: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                case /* repeated protos.PipelineStep steps */ 3:
                    message.steps.push(PipelineStep.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated protos.NotificationConfig _notification_configs */ 4:
                    message.NotificationConfigs.push(NotificationConfig.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* repeated protos.PipelineStep steps = 3; */
        for (let i = 0; i < message.steps.length; i++)
            PipelineStep.internalBinaryWrite(message.steps[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* repeated protos.NotificationConfig _notification_configs = 4; */
        for (let i = 0; i < message.NotificationConfigs.length; i++)
            NotificationConfig.internalBinaryWrite(message.NotificationConfigs[i], writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.Pipeline
 */
export const Pipeline = new Pipeline$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PipelineStep$Type extends MessageType {
    constructor() {
        super("protos.PipelineStep", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "on_success", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["protos.PipelineStepCondition", PipelineStepCondition, "PIPELINE_STEP_CONDITION_"] },
            { no: 3, name: "on_failure", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["protos.PipelineStepCondition", PipelineStepCondition, "PIPELINE_STEP_CONDITION_"] },
            { no: 1000, name: "detective", kind: "message", oneof: "step", T: () => DetectiveStep },
            { no: 1001, name: "transform", kind: "message", oneof: "step", T: () => TransformStep },
            { no: 1002, name: "encode", kind: "message", oneof: "step", T: () => EncodeStep },
            { no: 1003, name: "decode", kind: "message", oneof: "step", T: () => DecodeStep },
            { no: 1004, name: "custom", kind: "message", oneof: "step", T: () => CustomStep },
            { no: 1005, name: "http_request", kind: "message", oneof: "step", T: () => HttpRequestStep },
            { no: 1006, name: "kv", kind: "message", oneof: "step", T: () => KVStep },
            { no: 1007, name: "infer_schema", kind: "message", oneof: "step", T: () => InferSchemaStep },
            { no: 1008, name: "valid_json", kind: "message", oneof: "step", T: () => ValidJSONStep },
            { no: 10000, name: "_wasm_id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 10001, name: "_wasm_bytes", kind: "scalar", opt: true, T: 12 /*ScalarType.BYTES*/ },
            { no: 10002, name: "_wasm_function", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { name: "", onSuccess: [], onFailure: [], step: { oneofKind: undefined } };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* repeated protos.PipelineStepCondition on_success */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.onSuccess.push(reader.int32());
                    else
                        message.onSuccess.push(reader.int32());
                    break;
                case /* repeated protos.PipelineStepCondition on_failure */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.onFailure.push(reader.int32());
                    else
                        message.onFailure.push(reader.int32());
                    break;
                case /* protos.steps.DetectiveStep detective */ 1000:
                    message.step = {
                        oneofKind: "detective",
                        detective: DetectiveStep.internalBinaryRead(reader, reader.uint32(), options, message.step.detective)
                    };
                    break;
                case /* protos.steps.TransformStep transform */ 1001:
                    message.step = {
                        oneofKind: "transform",
                        transform: TransformStep.internalBinaryRead(reader, reader.uint32(), options, message.step.transform)
                    };
                    break;
                case /* protos.steps.EncodeStep encode */ 1002:
                    message.step = {
                        oneofKind: "encode",
                        encode: EncodeStep.internalBinaryRead(reader, reader.uint32(), options, message.step.encode)
                    };
                    break;
                case /* protos.steps.DecodeStep decode */ 1003:
                    message.step = {
                        oneofKind: "decode",
                        decode: DecodeStep.internalBinaryRead(reader, reader.uint32(), options, message.step.decode)
                    };
                    break;
                case /* protos.steps.CustomStep custom */ 1004:
                    message.step = {
                        oneofKind: "custom",
                        custom: CustomStep.internalBinaryRead(reader, reader.uint32(), options, message.step.custom)
                    };
                    break;
                case /* protos.steps.HttpRequestStep http_request */ 1005:
                    message.step = {
                        oneofKind: "httpRequest",
                        httpRequest: HttpRequestStep.internalBinaryRead(reader, reader.uint32(), options, message.step.httpRequest)
                    };
                    break;
                case /* protos.steps.KVStep kv */ 1006:
                    message.step = {
                        oneofKind: "kv",
                        kv: KVStep.internalBinaryRead(reader, reader.uint32(), options, message.step.kv)
                    };
                    break;
                case /* protos.steps.InferSchemaStep infer_schema */ 1007:
                    message.step = {
                        oneofKind: "inferSchema",
                        inferSchema: InferSchemaStep.internalBinaryRead(reader, reader.uint32(), options, message.step.inferSchema)
                    };
                    break;
                case /* protos.steps.ValidJSONStep valid_json */ 1008:
                    message.step = {
                        oneofKind: "validJson",
                        validJson: ValidJSONStep.internalBinaryRead(reader, reader.uint32(), options, message.step.validJson)
                    };
                    break;
                case /* optional string _wasm_id */ 10000:
                    message.WasmId = reader.string();
                    break;
                case /* optional bytes _wasm_bytes */ 10001:
                    message.WasmBytes = reader.bytes();
                    break;
                case /* optional string _wasm_function */ 10002:
                    message.WasmFunction = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* repeated protos.PipelineStepCondition on_success = 2; */
        if (message.onSuccess.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.onSuccess.length; i++)
                writer.int32(message.onSuccess[i]);
            writer.join();
        }
        /* repeated protos.PipelineStepCondition on_failure = 3; */
        if (message.onFailure.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.onFailure.length; i++)
                writer.int32(message.onFailure[i]);
            writer.join();
        }
        /* protos.steps.DetectiveStep detective = 1000; */
        if (message.step.oneofKind === "detective")
            DetectiveStep.internalBinaryWrite(message.step.detective, writer.tag(1000, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.TransformStep transform = 1001; */
        if (message.step.oneofKind === "transform")
            TransformStep.internalBinaryWrite(message.step.transform, writer.tag(1001, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.EncodeStep encode = 1002; */
        if (message.step.oneofKind === "encode")
            EncodeStep.internalBinaryWrite(message.step.encode, writer.tag(1002, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.DecodeStep decode = 1003; */
        if (message.step.oneofKind === "decode")
            DecodeStep.internalBinaryWrite(message.step.decode, writer.tag(1003, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.CustomStep custom = 1004; */
        if (message.step.oneofKind === "custom")
            CustomStep.internalBinaryWrite(message.step.custom, writer.tag(1004, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.HttpRequestStep http_request = 1005; */
        if (message.step.oneofKind === "httpRequest")
            HttpRequestStep.internalBinaryWrite(message.step.httpRequest, writer.tag(1005, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.KVStep kv = 1006; */
        if (message.step.oneofKind === "kv")
            KVStep.internalBinaryWrite(message.step.kv, writer.tag(1006, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.InferSchemaStep infer_schema = 1007; */
        if (message.step.oneofKind === "inferSchema")
            InferSchemaStep.internalBinaryWrite(message.step.inferSchema, writer.tag(1007, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.ValidJSONStep valid_json = 1008; */
        if (message.step.oneofKind === "validJson")
            ValidJSONStep.internalBinaryWrite(message.step.validJson, writer.tag(1008, WireType.LengthDelimited).fork(), options).join();
        /* optional string _wasm_id = 10000; */
        if (message.WasmId !== undefined)
            writer.tag(10000, WireType.LengthDelimited).string(message.WasmId);
        /* optional bytes _wasm_bytes = 10001; */
        if (message.WasmBytes !== undefined)
            writer.tag(10001, WireType.LengthDelimited).bytes(message.WasmBytes);
        /* optional string _wasm_function = 10002; */
        if (message.WasmFunction !== undefined)
            writer.tag(10002, WireType.LengthDelimited).string(message.WasmFunction);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.PipelineStep
 */
export const PipelineStep = new PipelineStep$Type();
