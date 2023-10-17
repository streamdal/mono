// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sp_common.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Common status codes used in gRPC method responses
type ResponseCode int32

const (
	ResponseCode_RESPONSE_CODE_UNSET                 ResponseCode = 0
	ResponseCode_RESPONSE_CODE_OK                    ResponseCode = 1
	ResponseCode_RESPONSE_CODE_BAD_REQUEST           ResponseCode = 2
	ResponseCode_RESPONSE_CODE_NOT_FOUND             ResponseCode = 3
	ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR ResponseCode = 4
	ResponseCode_RESPONSE_CODE_GENERIC_ERROR         ResponseCode = 5
)

// Enum value maps for ResponseCode.
var (
	ResponseCode_name = map[int32]string{
		0: "RESPONSE_CODE_UNSET",
		1: "RESPONSE_CODE_OK",
		2: "RESPONSE_CODE_BAD_REQUEST",
		3: "RESPONSE_CODE_NOT_FOUND",
		4: "RESPONSE_CODE_INTERNAL_SERVER_ERROR",
		5: "RESPONSE_CODE_GENERIC_ERROR",
	}
	ResponseCode_value = map[string]int32{
		"RESPONSE_CODE_UNSET":                 0,
		"RESPONSE_CODE_OK":                    1,
		"RESPONSE_CODE_BAD_REQUEST":           2,
		"RESPONSE_CODE_NOT_FOUND":             3,
		"RESPONSE_CODE_INTERNAL_SERVER_ERROR": 4,
		"RESPONSE_CODE_GENERIC_ERROR":         5,
	}
)

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}

func (x ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_common_proto_enumTypes[0].Descriptor()
}

func (ResponseCode) Type() protoreflect.EnumType {
	return &file_sp_common_proto_enumTypes[0]
}

func (x ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCode.Descriptor instead.
func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{0}
}

// Each SDK client is a $service + $component + $operation_type
type OperationType int32

const (
	OperationType_OPERATION_TYPE_UNSET    OperationType = 0
	OperationType_OPERATION_TYPE_CONSUMER OperationType = 1
	OperationType_OPERATION_TYPE_PRODUCER OperationType = 2
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_TYPE_UNSET",
		1: "OPERATION_TYPE_CONSUMER",
		2: "OPERATION_TYPE_PRODUCER",
	}
	OperationType_value = map[string]int32{
		"OPERATION_TYPE_UNSET":    0,
		"OPERATION_TYPE_CONSUMER": 1,
		"OPERATION_TYPE_PRODUCER": 2,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_common_proto_enumTypes[1].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_sp_common_proto_enumTypes[1]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{1}
}

type TailResponseType int32

const (
	TailResponseType_TAIL_RESPONSE_TYPE_UNSET   TailResponseType = 0
	TailResponseType_TAIL_RESPONSE_TYPE_PAYLOAD TailResponseType = 1
	TailResponseType_TAIL_RESPONSE_TYPE_ERROR   TailResponseType = 2
)

// Enum value maps for TailResponseType.
var (
	TailResponseType_name = map[int32]string{
		0: "TAIL_RESPONSE_TYPE_UNSET",
		1: "TAIL_RESPONSE_TYPE_PAYLOAD",
		2: "TAIL_RESPONSE_TYPE_ERROR",
	}
	TailResponseType_value = map[string]int32{
		"TAIL_RESPONSE_TYPE_UNSET":   0,
		"TAIL_RESPONSE_TYPE_PAYLOAD": 1,
		"TAIL_RESPONSE_TYPE_ERROR":   2,
	}
)

func (x TailResponseType) Enum() *TailResponseType {
	p := new(TailResponseType)
	*p = x
	return p
}

func (x TailResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TailResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_common_proto_enumTypes[2].Descriptor()
}

func (TailResponseType) Type() protoreflect.EnumType {
	return &file_sp_common_proto_enumTypes[2]
}

func (x TailResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TailResponseType.Descriptor instead.
func (TailResponseType) EnumDescriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{2}
}

type TailRequestType int32

const (
	TailRequestType_TAIL_REQUEST_TYPE_UNSET TailRequestType = 0
	TailRequestType_TAIL_REQUEST_TYPE_START TailRequestType = 1
	TailRequestType_TAIL_REQUEST_TYPE_STOP  TailRequestType = 2
)

// Enum value maps for TailRequestType.
var (
	TailRequestType_name = map[int32]string{
		0: "TAIL_REQUEST_TYPE_UNSET",
		1: "TAIL_REQUEST_TYPE_START",
		2: "TAIL_REQUEST_TYPE_STOP",
	}
	TailRequestType_value = map[string]int32{
		"TAIL_REQUEST_TYPE_UNSET": 0,
		"TAIL_REQUEST_TYPE_START": 1,
		"TAIL_REQUEST_TYPE_STOP":  2,
	}
)

func (x TailRequestType) Enum() *TailRequestType {
	p := new(TailRequestType)
	*p = x
	return p
}

func (x TailRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TailRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_common_proto_enumTypes[3].Descriptor()
}

func (TailRequestType) Type() protoreflect.EnumType {
	return &file_sp_common_proto_enumTypes[3]
}

func (x TailRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TailRequestType.Descriptor instead.
func (TailRequestType) EnumDescriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{3}
}

// Common response message for many gRPC methods
type StandardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Co-relation ID for the request / response
	Id      string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code    ResponseCode `protobuf:"varint,2,opt,name=code,proto3,enum=protos.ResponseCode" json:"code,omitempty"`
	Message string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StandardResponse) Reset() {
	*x = StandardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardResponse) ProtoMessage() {}

func (x *StandardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sp_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardResponse.ProtoReflect.Descriptor instead.
func (*StandardResponse) Descriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{0}
}

func (x *StandardResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StandardResponse) GetCode() ResponseCode {
	if x != nil {
		return x.Code
	}
	return ResponseCode_RESPONSE_CODE_UNSET
}

func (x *StandardResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Used to indicate who a command is intended for
type Audience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the service -- let's include the service name on all calls, we can
	// optimize later ~DS
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Name of the component the SDK is interacting with (ie. kafka-$topic-name)
	ComponentName string `protobuf:"bytes,2,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	// Consumer or Producer
	OperationType OperationType `protobuf:"varint,3,opt,name=operation_type,json=operationType,proto3,enum=protos.OperationType" json:"operation_type,omitempty"`
	// Name for the consumer or producer
	OperationName string `protobuf:"bytes,4,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
}

func (x *Audience) Reset() {
	*x = Audience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audience) ProtoMessage() {}

func (x *Audience) ProtoReflect() protoreflect.Message {
	mi := &file_sp_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audience.ProtoReflect.Descriptor instead.
func (*Audience) Descriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{1}
}

func (x *Audience) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Audience) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

func (x *Audience) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_OPERATION_TYPE_UNSET
}

func (x *Audience) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels   map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value    float64           `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Audience *Audience         `protobuf:"bytes,4,opt,name=audience,proto3" json:"audience,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_sp_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{2}
}

func (x *Metric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metric) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Metric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Metric) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

type TailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       TailRequestType   `protobuf:"varint,1,opt,name=type,proto3,enum=protos.TailRequestType" json:"type,omitempty"`
	XId        *string           `protobuf:"bytes,2,opt,name=_id,json=Id,proto3,oneof" json:"_id,omitempty"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
	Audience   *Audience         `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	PipelineId *string           `protobuf:"bytes,4,opt,name=pipeline_id,json=pipelineId,proto3,oneof" json:"pipeline_id,omitempty"`
	XMetadata  map[string]string `protobuf:"bytes,1000,rep,name=_metadata,json=Metadata,proto3" json:"_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}

func (x *TailRequest) Reset() {
	*x = TailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TailRequest) ProtoMessage() {}

func (x *TailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TailRequest.ProtoReflect.Descriptor instead.
func (*TailRequest) Descriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{3}
}

func (x *TailRequest) GetType() TailRequestType {
	if x != nil {
		return x.Type
	}
	return TailRequestType_TAIL_REQUEST_TYPE_UNSET
}

func (x *TailRequest) GetXId() string {
	if x != nil && x.XId != nil {
		return *x.XId
	}
	return ""
}

func (x *TailRequest) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *TailRequest) GetPipelineId() string {
	if x != nil && x.PipelineId != nil {
		return *x.PipelineId
	}
	return ""
}

func (x *TailRequest) GetXMetadata() map[string]string {
	if x != nil {
		return x.XMetadata
	}
	return nil
}

// TailResponse originates in the SDK and then is sent to snitch servers where
// it is forwarded to the correct frontend streaming gRPC connection
type TailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          TailResponseType `protobuf:"varint,1,opt,name=type,proto3,enum=protos.TailResponseType" json:"type,omitempty"`
	TailRequestId string           `protobuf:"bytes,2,opt,name=tail_request_id,json=tailRequestId,proto3" json:"tail_request_id,omitempty"`
	Audience      *Audience        `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	PipelineId    string           `protobuf:"bytes,4,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	SessionId     string           `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Timestamp in nanoseconds
	TimestampNs int64 `protobuf:"varint,6,opt,name=timestamp_ns,json=timestampNs,proto3" json:"timestamp_ns,omitempty"`
	// Payload data. For errors, this will be the error message
	// For payloads, this will be JSON of the payload data, post processing
	OriginalData []byte `protobuf:"bytes,7,opt,name=original_data,json=originalData,proto3" json:"original_data,omitempty"`
	// For payloads, this will be the new data, post processing
	NewData   []byte            `protobuf:"bytes,8,opt,name=new_data,json=newData,proto3" json:"new_data,omitempty"`
	XMetadata map[string]string `protobuf:"bytes,1000,rep,name=_metadata,json=Metadata,proto3" json:"_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
	// Set by server to indicate that the response is a keepalive message
	XKeepalive *bool `protobuf:"varint,1001,opt,name=_keepalive,json=Keepalive,proto3,oneof" json:"_keepalive,omitempty"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}

func (x *TailResponse) Reset() {
	*x = TailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TailResponse) ProtoMessage() {}

func (x *TailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sp_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TailResponse.ProtoReflect.Descriptor instead.
func (*TailResponse) Descriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{4}
}

func (x *TailResponse) GetType() TailResponseType {
	if x != nil {
		return x.Type
	}
	return TailResponseType_TAIL_RESPONSE_TYPE_UNSET
}

func (x *TailResponse) GetTailRequestId() string {
	if x != nil {
		return x.TailRequestId
	}
	return ""
}

func (x *TailResponse) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *TailResponse) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *TailResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *TailResponse) GetTimestampNs() int64 {
	if x != nil {
		return x.TimestampNs
	}
	return 0
}

func (x *TailResponse) GetOriginalData() []byte {
	if x != nil {
		return x.OriginalData
	}
	return nil
}

func (x *TailResponse) GetNewData() []byte {
	if x != nil {
		return x.NewData
	}
	return nil
}

func (x *TailResponse) GetXMetadata() map[string]string {
	if x != nil {
		return x.XMetadata
	}
	return nil
}

func (x *TailResponse) GetXKeepalive() bool {
	if x != nil && x.XKeepalive != nil {
		return *x.XKeepalive
	}
	return false
}

type AudienceRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes     float64 `protobuf:"fixed64,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Processed float64 `protobuf:"fixed64,2,opt,name=processed,proto3" json:"processed,omitempty"`
}

func (x *AudienceRate) Reset() {
	*x = AudienceRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudienceRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudienceRate) ProtoMessage() {}

func (x *AudienceRate) ProtoReflect() protoreflect.Message {
	mi := &file_sp_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudienceRate.ProtoReflect.Descriptor instead.
func (*AudienceRate) Descriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{5}
}

func (x *AudienceRate) GetBytes() float64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *AudienceRate) GetProcessed() float64 {
	if x != nil {
		return x.Processed
	}
	return 0
}

type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonSchema []byte            `protobuf:"bytes,1,opt,name=json_schema,json=jsonSchema,proto3" json:"json_schema,omitempty"`
	XVersion   int32             `protobuf:"varint,100,opt,name=_version,json=Version,proto3" json:"_version,omitempty"`                                                                                            // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
	XMetadata  map[string]string `protobuf:"bytes,1000,rep,name=_metadata,json=Metadata,proto3" json:"_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_sp_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_sp_common_proto_rawDescGZIP(), []int{6}
}

func (x *Schema) GetJsonSchema() []byte {
	if x != nil {
		return x.JsonSchema
	}
	return nil
}

func (x *Schema) GetXVersion() int32 {
	if x != nil {
		return x.XVersion
	}
	return 0
}

func (x *Schema) GetXMetadata() map[string]string {
	if x != nil {
		return x.XMetadata
	}
	return nil
}

var File_sp_common_proto protoreflect.FileDescriptor

var file_sp_common_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x66, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xb9, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcf, 0x01,
	0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xba, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x24, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x58, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x22, 0xe8, 0x03, 0x0a,
	0x0c, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74,
	0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x4e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0a, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x4b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x58, 0x5f, 0x6b, 0x65,
	0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x06,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6a, 0x73, 0x6f,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0xe8, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0xc3, 0x01, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x13,
	0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x41, 0x44,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04,
	0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x05, 0x2a, 0x63, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x45, 0x52, 0x10, 0x02, 0x2a, 0x6e, 0x0a, 0x10, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41,
	0x49, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x41, 0x49, 0x4c,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41, 0x49, 0x4c,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x2a, 0x67, 0x0a, 0x0f, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x41, 0x49,
	0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x41, 0x49, 0x4c, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x41, 0x49, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x6e, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_common_proto_rawDescOnce sync.Once
	file_sp_common_proto_rawDescData = file_sp_common_proto_rawDesc
)

func file_sp_common_proto_rawDescGZIP() []byte {
	file_sp_common_proto_rawDescOnce.Do(func() {
		file_sp_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_common_proto_rawDescData)
	})
	return file_sp_common_proto_rawDescData
}

var file_sp_common_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_sp_common_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_sp_common_proto_goTypes = []interface{}{
	(ResponseCode)(0),        // 0: protos.ResponseCode
	(OperationType)(0),       // 1: protos.OperationType
	(TailResponseType)(0),    // 2: protos.TailResponseType
	(TailRequestType)(0),     // 3: protos.TailRequestType
	(*StandardResponse)(nil), // 4: protos.StandardResponse
	(*Audience)(nil),         // 5: protos.Audience
	(*Metric)(nil),           // 6: protos.Metric
	(*TailRequest)(nil),      // 7: protos.TailRequest
	(*TailResponse)(nil),     // 8: protos.TailResponse
	(*AudienceRate)(nil),     // 9: protos.AudienceRate
	(*Schema)(nil),           // 10: protos.Schema
	nil,                      // 11: protos.Metric.LabelsEntry
	nil,                      // 12: protos.TailRequest.MetadataEntry
	nil,                      // 13: protos.TailResponse.MetadataEntry
	nil,                      // 14: protos.Schema.MetadataEntry
}
var file_sp_common_proto_depIdxs = []int32{
	0,  // 0: protos.StandardResponse.code:type_name -> protos.ResponseCode
	1,  // 1: protos.Audience.operation_type:type_name -> protos.OperationType
	11, // 2: protos.Metric.labels:type_name -> protos.Metric.LabelsEntry
	5,  // 3: protos.Metric.audience:type_name -> protos.Audience
	3,  // 4: protos.TailRequest.type:type_name -> protos.TailRequestType
	5,  // 5: protos.TailRequest.audience:type_name -> protos.Audience
	12, // 6: protos.TailRequest._metadata:type_name -> protos.TailRequest.MetadataEntry
	2,  // 7: protos.TailResponse.type:type_name -> protos.TailResponseType
	5,  // 8: protos.TailResponse.audience:type_name -> protos.Audience
	13, // 9: protos.TailResponse._metadata:type_name -> protos.TailResponse.MetadataEntry
	14, // 10: protos.Schema._metadata:type_name -> protos.Schema.MetadataEntry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_sp_common_proto_init() }
func file_sp_common_proto_init() {
	if File_sp_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sp_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandardResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audience); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudienceRate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_sp_common_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_sp_common_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_common_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_common_proto_goTypes,
		DependencyIndexes: file_sp_common_proto_depIdxs,
		EnumInfos:         file_sp_common_proto_enumTypes,
		MessageInfos:      file_sp_common_proto_msgTypes,
	}.Build()
	File_sp_common_proto = out.File
	file_sp_common_proto_rawDesc = nil
	file_sp_common_proto_goTypes = nil
	file_sp_common_proto_depIdxs = nil
}
