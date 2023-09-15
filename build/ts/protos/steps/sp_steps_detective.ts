// @generated by protobuf-ts 2.9.0 with parameter long_type_string
// @generated from protobuf file "steps/sp_steps_detective.proto" (package "protos.steps", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message protos.steps.DetectiveStep
 */
export interface DetectiveStep {
    /**
     * @generated from protobuf field: optional string path = 1;
     */
    path?: string;
    /**
     * @generated from protobuf field: repeated string args = 2;
     */
    args: string[]; // args determined by match_type
    /**
     * @generated from protobuf field: optional bool negate = 3;
     */
    negate?: boolean;
    /**
     * @generated from protobuf field: protos.steps.DetectiveType type = 4;
     */
    type: DetectiveType;
}
/**
 * 1000-1999 reserved for core match types
 *
 * @generated from protobuf enum protos.steps.DetectiveType
 */
export enum DetectiveType {
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IS_EMPTY = 1000;
     */
    IS_EMPTY = 1000,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_HAS_FIELD = 1001;
     */
    HAS_FIELD = 1001,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IS_TYPE = 1002;
     */
    IS_TYPE = 1002,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_CONTAINS_ANY = 1003;
     */
    STRING_CONTAINS_ANY = 1003,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_CONTAINS_ALL = 1004;
     */
    STRING_CONTAINS_ALL = 1004,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_EQUAL = 1005;
     */
    STRING_EQUAL = 1005,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IPV4_ADDRESS = 1006;
     */
    IPV4_ADDRESS = 1006,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IPV6_ADDRESS = 1007;
     */
    IPV6_ADDRESS = 1007,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_MAC_ADDRESS = 1008;
     */
    MAC_ADDRESS = 1008,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_REGEX = 1009;
     */
    REGEX = 1009,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_TIMESTAMP_RFC3339 = 1010;
     */
    TIMESTAMP_RFC3339 = 1010,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_TIMESTAMP_UNIX_NANO = 1011;
     */
    TIMESTAMP_UNIX_NANO = 1011,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_TIMESTAMP_UNIX = 1012;
     */
    TIMESTAMP_UNIX = 1012,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_BOOLEAN_TRUE = 1013;
     */
    BOOLEAN_TRUE = 1013,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_BOOLEAN_FALSE = 1014;
     */
    BOOLEAN_FALSE = 1014,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_UUID = 1015;
     */
    UUID = 1015,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_URL = 1016;
     */
    URL = 1016,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_HOSTNAME = 1017;
     */
    HOSTNAME = 1017,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_LENGTH_MIN = 1018;
     */
    STRING_LENGTH_MIN = 1018,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_LENGTH_MAX = 1019;
     */
    STRING_LENGTH_MAX = 1019,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_LENGTH_RANGE = 1020;
     */
    STRING_LENGTH_RANGE = 1020,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_SEMVER = 2021;
     */
    SEMVER = 2021,
    /**
     * / Payloads containing values with any PII - runs all PII matchers
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_ANY = 2000;
     */
    PII_ANY = 2000,
    /**
     * Payloads containing values with a credit card number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_CREDIT_CARD = 2001;
     */
    PII_CREDIT_CARD = 2001,
    /**
     * Payloads containing values with a social security number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_SSN = 2002;
     */
    PII_SSN = 2002,
    /**
     * Payloads containing values with an email address
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_EMAIL = 2003;
     */
    PII_EMAIL = 2003,
    /**
     * Payloads containing values with a phone number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_PHONE = 2004;
     */
    PII_PHONE = 2004,
    /**
     * Payloads containing values with a driver's license
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_DRIVER_LICENSE = 2005;
     */
    PII_DRIVER_LICENSE = 2005,
    /**
     * Payloads containing values with a passport ID
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_PASSPORT_ID = 2006;
     */
    PII_PASSPORT_ID = 2006,
    /**
     * Payloads containing values with a VIN number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_VIN_NUMBER = 2007;
     */
    PII_VIN_NUMBER = 2007,
    /**
     * Payloads containing values with various serial number formats
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_SERIAL_NUMBER = 2008;
     */
    PII_SERIAL_NUMBER = 2008,
    /**
     * Payloads containing fields named "login", "username", "user", "userid", "user_id", "user", "password", "pass", "passwd", "pwd"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_LOGIN = 2009;
     */
    PII_LOGIN = 2009,
    /**
     * Payloads containing fields named "taxpayer_id", "tax_id", "taxpayerid", "taxid"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_TAXPAYER_ID = 2010;
     */
    PII_TAXPAYER_ID = 2010,
    /**
     * Payloads containing fields named "address", "street", "city", "state", "zip", "zipcode", "zip_code", "country"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_ADDRESS = 2011;
     */
    PII_ADDRESS = 2011,
    /**
     * Payloads containing fields named "signature", "signature_image", "signature_image_url", "signature_image_uri"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_SIGNATURE = 2012;
     */
    PII_SIGNATURE = 2012,
    /**
     * Payloads containing values that contain GPS data or coordinates like "lat", "lon", "latitude", "longitude"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_GEOLOCATION = 2013;
     */
    PII_GEOLOCATION = 2013,
    /**
     * Payloads containing fields like "school", "university", "college", "education"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_EDUCATION = 2014;
     */
    PII_EDUCATION = 2014,
    /**
     * Payloads containing fields like "account", "bank", "credit", "debit", "financial", "finance"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_FINANCIAL = 2015;
     */
    PII_FINANCIAL = 2015,
    /**
     * Payloads containing fields like "patient", "health", "healthcare", "health care", "medical"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_HEALTH = 2016;
     */
    PII_HEALTH = 2016,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_EQUAL_TO = 3000;
     */
    NUMERIC_EQUAL_TO = 3000,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_GREATER_THAN = 3001;
     */
    NUMERIC_GREATER_THAN = 3001,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_GREATER_EQUAL = 3002;
     */
    NUMERIC_GREATER_EQUAL = 3002,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_LESS_THAN = 3003;
     */
    NUMERIC_LESS_THAN = 3003,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_LESS_EQUAL = 3004;
     */
    NUMERIC_LESS_EQUAL = 3004,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_RANGE = 3005;
     */
    NUMERIC_RANGE = 3005,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_MIN = 3006;
     */
    NUMERIC_MIN = 3006,
    /**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_MAX = 3007;
     */
    NUMERIC_MAX = 3007
}
// @generated message type with reflection information, may provide speed optimized methods
class DetectiveStep$Type extends MessageType<DetectiveStep> {
    constructor() {
        super("protos.steps.DetectiveStep", [
            { no: 1, name: "path", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "args", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "negate", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "type", kind: "enum", T: () => ["protos.steps.DetectiveType", DetectiveType, "DETECTIVE_TYPE_"] }
        ]);
    }
    create(value?: PartialMessage<DetectiveStep>): DetectiveStep {
        const message = { args: [], type: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<DetectiveStep>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DetectiveStep): DetectiveStep {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional string path */ 1:
                    message.path = reader.string();
                    break;
                case /* repeated string args */ 2:
                    message.args.push(reader.string());
                    break;
                case /* optional bool negate */ 3:
                    message.negate = reader.bool();
                    break;
                case /* protos.steps.DetectiveType type */ 4:
                    message.type = reader.int32();
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
    internalBinaryWrite(message: DetectiveStep, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* optional string path = 1; */
        if (message.path !== undefined)
            writer.tag(1, WireType.LengthDelimited).string(message.path);
        /* repeated string args = 2; */
        for (let i = 0; i < message.args.length; i++)
            writer.tag(2, WireType.LengthDelimited).string(message.args[i]);
        /* optional bool negate = 3; */
        if (message.negate !== undefined)
            writer.tag(3, WireType.Varint).bool(message.negate);
        /* protos.steps.DetectiveType type = 4; */
        if (message.type !== 0)
            writer.tag(4, WireType.Varint).int32(message.type);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.DetectiveStep
 */
export const DetectiveStep = new DetectiveStep$Type();
