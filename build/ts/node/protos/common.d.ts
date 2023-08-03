import { MessageType } from "@protobuf-ts/runtime";
/**
 * Common response message for many gRPC methods
 *
 * @generated from protobuf message protos.StandardResponse
 */
export interface StandardResponse {
    /**
     * Co-relation ID for the request / response
     *
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: protos.ResponseCode code = 2;
     */
    code: ResponseCode;
    /**
     * @generated from protobuf field: string message = 3;
     */
    message: string;
}
/**
 * Used to indicate who a command is intended for
 *
 * @generated from protobuf message protos.Audience
 */
export interface Audience {
    /**
     * Name of the service -- let's include the service name on all calls, we can
     * optimize later ~DS
     *
     * @generated from protobuf field: string service_name = 1;
     */
    serviceName: string;
    /**
     * Name of the component the SDK is interacting with (ie. kafka-$topic-name)
     *
     * @generated from protobuf field: string component_name = 2;
     */
    componentName: string;
    /**
     * Consumer or Producer
     *
     * @generated from protobuf field: protos.OperationType operation_type = 3;
     */
    operationType: OperationType;
    /**
     * Name for the consumer or producer
     *
     * @generated from protobuf field: string operation_name = 4;
     */
    operationName: string;
}
/**
 * Common status codes used in gRPC method responses
 *
 * @generated from protobuf enum protos.ResponseCode
 */
export declare enum ResponseCode {
    /**
     * @generated from protobuf enum value: RESPONSE_CODE_UNSET = 0;
     */
    UNSET = 0,
    /**
     * @generated from protobuf enum value: RESPONSE_CODE_OK = 1;
     */
    OK = 1,
    /**
     * @generated from protobuf enum value: RESPONSE_CODE_BAD_REQUEST = 2;
     */
    BAD_REQUEST = 2,
    /**
     * @generated from protobuf enum value: RESPONSE_CODE_NOT_FOUND = 3;
     */
    NOT_FOUND = 3,
    /**
     * @generated from protobuf enum value: RESPONSE_CODE_INTERNAL_SERVER_ERROR = 4;
     */
    INTERNAL_SERVER_ERROR = 4,
    /**
     * @generated from protobuf enum value: RESPONSE_CODE_GENERIC_ERROR = 5;
     */
    GENERIC_ERROR = 5
}
/**
 * Each SDK client is a $service + $component + $operation_type
 *
 * @generated from protobuf enum protos.OperationType
 */
export declare enum OperationType {
    /**
     * @generated from protobuf enum value: OPERATION_TYPE_UNSET = 0;
     */
    UNSET = 0,
    /**
     * @generated from protobuf enum value: OPERATION_TYPE_CONSUMER = 1;
     */
    CONSUMER = 1,
    /**
     * @generated from protobuf enum value: OPERATION_TYPE_PRODUCER = 2;
     */
    PRODUCER = 2
}
declare class StandardResponse$Type extends MessageType<StandardResponse> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.StandardResponse
 */
export declare const StandardResponse: StandardResponse$Type;
declare class Audience$Type extends MessageType<Audience> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.Audience
 */
export declare const Audience: Audience$Type;
export {};
