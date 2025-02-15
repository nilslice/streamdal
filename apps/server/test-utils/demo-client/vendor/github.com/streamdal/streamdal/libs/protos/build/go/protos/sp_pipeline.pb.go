// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sp_pipeline.proto

package protos

import (
	steps "github.com/streamdal/streamdal/libs/protos/build/go/protos/steps"
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

// Defines the ways in which a pipeline can be aborted
type AbortCondition int32

const (
	AbortCondition_ABORT_CONDITION_UNSET         AbortCondition = 0
	AbortCondition_ABORT_CONDITION_ABORT_CURRENT AbortCondition = 1
	AbortCondition_ABORT_CONDITION_ABORT_ALL     AbortCondition = 2
)

// Enum value maps for AbortCondition.
var (
	AbortCondition_name = map[int32]string{
		0: "ABORT_CONDITION_UNSET",
		1: "ABORT_CONDITION_ABORT_CURRENT",
		2: "ABORT_CONDITION_ABORT_ALL",
	}
	AbortCondition_value = map[string]int32{
		"ABORT_CONDITION_UNSET":         0,
		"ABORT_CONDITION_ABORT_CURRENT": 1,
		"ABORT_CONDITION_ABORT_ALL":     2,
	}
)

func (x AbortCondition) Enum() *AbortCondition {
	p := new(AbortCondition)
	*p = x
	return p
}

func (x AbortCondition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AbortCondition) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_pipeline_proto_enumTypes[0].Descriptor()
}

func (AbortCondition) Type() protoreflect.EnumType {
	return &file_sp_pipeline_proto_enumTypes[0]
}

func (x AbortCondition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AbortCondition.Descriptor instead.
func (AbortCondition) EnumDescriptor() ([]byte, []int) {
	return file_sp_pipeline_proto_rawDescGZIP(), []int{0}
}

// Pipeline is a structure that holds one or more pipeline steps. This structure
// is intended to be immutable; clients are expected to generate WASMRequest's
// that contain a pipeline step.
type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID should NOT be set by external gRPC client on CreatePipelineRequest - it
	// will be ignored; it _does_ need to be set on UpdatePipelineRequest.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Friendly name for the pipeline
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// One or more steps to execute
	Steps []*PipelineStep `protobuf:"bytes,3,rep,name=steps,proto3" json:"steps,omitempty"`
	// Notification configs for this pipeline. Only filled out
	// in external API responses
	XNotificationConfigs []*NotificationConfig `protobuf:"bytes,4,rep,name=_notification_configs,json=NotificationConfigs,proto3" json:"_notification_configs,omitempty"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_sp_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_sp_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *Pipeline) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pipeline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pipeline) GetSteps() []*PipelineStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Pipeline) GetXNotificationConfigs() []*NotificationConfig {
	if x != nil {
		return x.XNotificationConfigs
	}
	return nil
}

// Conditions define how the SDK should handle a Wasm response in a step.
// Should it continue executing the pipeline, should it abort, should it notify
// and on_error.
type PipelineStepConditions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Should we abort execution?
	Abort AbortCondition `protobuf:"varint,1,opt,name=abort,proto3,enum=protos.AbortCondition" json:"abort,omitempty"`
	// Should we trigger a notification?
	Notify bool `protobuf:"varint,2,opt,name=notify,proto3" json:"notify,omitempty"`
	// Should we include additional metadata that SDK should pass back to user?
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PipelineStepConditions) Reset() {
	*x = PipelineStepConditions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineStepConditions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineStepConditions) ProtoMessage() {}

func (x *PipelineStepConditions) ProtoReflect() protoreflect.Message {
	mi := &file_sp_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineStepConditions.ProtoReflect.Descriptor instead.
func (*PipelineStepConditions) Descriptor() ([]byte, []int) {
	return file_sp_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineStepConditions) GetAbort() AbortCondition {
	if x != nil {
		return x.Abort
	}
	return AbortCondition_ABORT_CONDITION_UNSET
}

func (x *PipelineStepConditions) GetNotify() bool {
	if x != nil {
		return x.Notify
	}
	return false
}

func (x *PipelineStepConditions) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// A pipeline step is a single step in a pipeline.
type PipelineStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Friendly name for the step
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SDKs should read this when Wasm returns 'true' to determine what to do next.
	OnTrue *PipelineStepConditions `protobuf:"bytes,2,opt,name=on_true,json=onTrue,proto3" json:"on_true,omitempty"`
	// SDKs should read this when Wasm returns 'false' to determine what to do next.
	OnFalse *PipelineStepConditions `protobuf:"bytes,3,opt,name=on_false,json=onFalse,proto3" json:"on_false,omitempty"`
	// Indicates whether to use the results from a previous step as input to this step
	Dynamic bool `protobuf:"varint,4,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	// SDKs should read this when Wasm returns 'error' to determine what to do next.
	OnError *PipelineStepConditions `protobuf:"bytes,5,opt,name=on_error,json=onError,proto3" json:"on_error,omitempty"`
	// Types that are assignable to Step:
	//
	//	*PipelineStep_Detective
	//	*PipelineStep_Transform
	//	*PipelineStep_Encode
	//	*PipelineStep_Decode
	//	*PipelineStep_Custom
	//	*PipelineStep_HttpRequest
	//	*PipelineStep_Kv
	//	*PipelineStep_InferSchema
	//	*PipelineStep_ValidJson
	//	*PipelineStep_SchemaValidation
	Step isPipelineStep_Step `protobuf_oneof:"step"`
	// ID is a uuid(sha256(_wasm_bytes)) that is set by server
	XWasmId *string `protobuf:"bytes,10000,opt,name=_wasm_id,json=WasmId,proto3,oneof" json:"_wasm_id,omitempty"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
	// WASM module bytes (set by server)
	XWasmBytes []byte `protobuf:"bytes,10001,opt,name=_wasm_bytes,json=WasmBytes,proto3,oneof" json:"_wasm_bytes,omitempty"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
	// WASM function name to execute (set by server)
	XWasmFunction *string `protobuf:"bytes,10002,opt,name=_wasm_function,json=WasmFunction,proto3,oneof" json:"_wasm_function,omitempty"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}

func (x *PipelineStep) Reset() {
	*x = PipelineStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineStep) ProtoMessage() {}

func (x *PipelineStep) ProtoReflect() protoreflect.Message {
	mi := &file_sp_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineStep.ProtoReflect.Descriptor instead.
func (*PipelineStep) Descriptor() ([]byte, []int) {
	return file_sp_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *PipelineStep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PipelineStep) GetOnTrue() *PipelineStepConditions {
	if x != nil {
		return x.OnTrue
	}
	return nil
}

func (x *PipelineStep) GetOnFalse() *PipelineStepConditions {
	if x != nil {
		return x.OnFalse
	}
	return nil
}

func (x *PipelineStep) GetDynamic() bool {
	if x != nil {
		return x.Dynamic
	}
	return false
}

func (x *PipelineStep) GetOnError() *PipelineStepConditions {
	if x != nil {
		return x.OnError
	}
	return nil
}

func (m *PipelineStep) GetStep() isPipelineStep_Step {
	if m != nil {
		return m.Step
	}
	return nil
}

func (x *PipelineStep) GetDetective() *steps.DetectiveStep {
	if x, ok := x.GetStep().(*PipelineStep_Detective); ok {
		return x.Detective
	}
	return nil
}

func (x *PipelineStep) GetTransform() *steps.TransformStep {
	if x, ok := x.GetStep().(*PipelineStep_Transform); ok {
		return x.Transform
	}
	return nil
}

func (x *PipelineStep) GetEncode() *steps.EncodeStep {
	if x, ok := x.GetStep().(*PipelineStep_Encode); ok {
		return x.Encode
	}
	return nil
}

func (x *PipelineStep) GetDecode() *steps.DecodeStep {
	if x, ok := x.GetStep().(*PipelineStep_Decode); ok {
		return x.Decode
	}
	return nil
}

func (x *PipelineStep) GetCustom() *steps.CustomStep {
	if x, ok := x.GetStep().(*PipelineStep_Custom); ok {
		return x.Custom
	}
	return nil
}

func (x *PipelineStep) GetHttpRequest() *steps.HttpRequestStep {
	if x, ok := x.GetStep().(*PipelineStep_HttpRequest); ok {
		return x.HttpRequest
	}
	return nil
}

func (x *PipelineStep) GetKv() *steps.KVStep {
	if x, ok := x.GetStep().(*PipelineStep_Kv); ok {
		return x.Kv
	}
	return nil
}

func (x *PipelineStep) GetInferSchema() *steps.InferSchemaStep {
	if x, ok := x.GetStep().(*PipelineStep_InferSchema); ok {
		return x.InferSchema
	}
	return nil
}

func (x *PipelineStep) GetValidJson() *steps.ValidJSONStep {
	if x, ok := x.GetStep().(*PipelineStep_ValidJson); ok {
		return x.ValidJson
	}
	return nil
}

func (x *PipelineStep) GetSchemaValidation() *steps.SchemaValidationStep {
	if x, ok := x.GetStep().(*PipelineStep_SchemaValidation); ok {
		return x.SchemaValidation
	}
	return nil
}

func (x *PipelineStep) GetXWasmId() string {
	if x != nil && x.XWasmId != nil {
		return *x.XWasmId
	}
	return ""
}

func (x *PipelineStep) GetXWasmBytes() []byte {
	if x != nil {
		return x.XWasmBytes
	}
	return nil
}

func (x *PipelineStep) GetXWasmFunction() string {
	if x != nil && x.XWasmFunction != nil {
		return *x.XWasmFunction
	}
	return ""
}

type isPipelineStep_Step interface {
	isPipelineStep_Step()
}

type PipelineStep_Detective struct {
	Detective *steps.DetectiveStep `protobuf:"bytes,1000,opt,name=detective,proto3,oneof"`
}

type PipelineStep_Transform struct {
	Transform *steps.TransformStep `protobuf:"bytes,1001,opt,name=transform,proto3,oneof"`
}

type PipelineStep_Encode struct {
	Encode *steps.EncodeStep `protobuf:"bytes,1002,opt,name=encode,proto3,oneof"`
}

type PipelineStep_Decode struct {
	Decode *steps.DecodeStep `protobuf:"bytes,1003,opt,name=decode,proto3,oneof"`
}

type PipelineStep_Custom struct {
	Custom *steps.CustomStep `protobuf:"bytes,1004,opt,name=custom,proto3,oneof"`
}

type PipelineStep_HttpRequest struct {
	HttpRequest *steps.HttpRequestStep `protobuf:"bytes,1005,opt,name=http_request,json=httpRequest,proto3,oneof"`
}

type PipelineStep_Kv struct {
	Kv *steps.KVStep `protobuf:"bytes,1006,opt,name=kv,proto3,oneof"`
}

type PipelineStep_InferSchema struct {
	InferSchema *steps.InferSchemaStep `protobuf:"bytes,1007,opt,name=infer_schema,json=inferSchema,proto3,oneof"`
}

type PipelineStep_ValidJson struct {
	ValidJson *steps.ValidJSONStep `protobuf:"bytes,1008,opt,name=valid_json,json=validJson,proto3,oneof"`
}

type PipelineStep_SchemaValidation struct {
	SchemaValidation *steps.SchemaValidationStep `protobuf:"bytes,1009,opt,name=schema_validation,json=schemaValidation,proto3,oneof"`
}

func (*PipelineStep_Detective) isPipelineStep_Step() {}

func (*PipelineStep_Transform) isPipelineStep_Step() {}

func (*PipelineStep_Encode) isPipelineStep_Step() {}

func (*PipelineStep_Decode) isPipelineStep_Step() {}

func (*PipelineStep_Custom) isPipelineStep_Step() {}

func (*PipelineStep_HttpRequest) isPipelineStep_Step() {}

func (*PipelineStep_Kv) isPipelineStep_Step() {}

func (*PipelineStep_InferSchema) isPipelineStep_Step() {}

func (*PipelineStep_ValidJson) isPipelineStep_Step() {}

func (*PipelineStep_SchemaValidation) isPipelineStep_Step() {}

var File_sp_pipeline_proto protoreflect.FileDescriptor

var file_sp_pipeline_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x70, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0f, 0x73, 0x70, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x2f, 0x73, 0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x2f, 0x73, 0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f, 0x73, 0x70,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f, 0x73, 0x70,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f, 0x73, 0x70, 0x5f, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f, 0x73, 0x70, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f, 0x73,
	0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f, 0x73, 0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f,
	0x73, 0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f,
	0x73, 0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x08, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x4e, 0x0a, 0x15, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x16, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x2c, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x48, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf6,
	0x07, 0x0a, 0x0c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x6f, 0x6e, 0x54, 0x72, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x08,
	0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07,
	0x6f, 0x6e, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x12, 0x39, 0x0a, 0x08, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x09,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52,
	0x09, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x65, 0x70, 0x48, 0x00, 0x52, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a,
	0x06, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x18, 0xeb, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0xec, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65,
	0x70, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52,
	0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x43, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0xed, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52,
	0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x02,
	0x6b, 0x76, 0x18, 0xee, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x4b, 0x56, 0x53, 0x74, 0x65, 0x70, 0x48,
	0x00, 0x52, 0x02, 0x6b, 0x76, 0x12, 0x43, 0x0a, 0x0c, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0xef, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x65,
	0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x69,
	0x6e, 0x66, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3d, 0x0a, 0x0a, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0xf0, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52, 0x09,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x11, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xf1,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x65, 0x70, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52, 0x10, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x08, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x90, 0x4e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x06, 0x57, 0x61, 0x73, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0b, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x91, 0x4e, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x09, 0x57, 0x61, 0x73, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x92, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0c,
	0x57, 0x61, 0x73, 0x6d, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x06, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x42, 0x0b, 0x0a, 0x09, 0x58, 0x5f, 0x77, 0x61, 0x73,
	0x6d, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x58, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x58, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x6d, 0x0a, 0x0e, 0x41, 0x62, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x42, 0x4f,
	0x52, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x45, 0x54, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x4f,
	0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x55,
	0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x42, 0x4f, 0x52, 0x54,
	0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_pipeline_proto_rawDescOnce sync.Once
	file_sp_pipeline_proto_rawDescData = file_sp_pipeline_proto_rawDesc
)

func file_sp_pipeline_proto_rawDescGZIP() []byte {
	file_sp_pipeline_proto_rawDescOnce.Do(func() {
		file_sp_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_pipeline_proto_rawDescData)
	})
	return file_sp_pipeline_proto_rawDescData
}

var file_sp_pipeline_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sp_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sp_pipeline_proto_goTypes = []interface{}{
	(AbortCondition)(0),                // 0: protos.AbortCondition
	(*Pipeline)(nil),                   // 1: protos.Pipeline
	(*PipelineStepConditions)(nil),     // 2: protos.PipelineStepConditions
	(*PipelineStep)(nil),               // 3: protos.PipelineStep
	nil,                                // 4: protos.PipelineStepConditions.MetadataEntry
	(*NotificationConfig)(nil),         // 5: protos.NotificationConfig
	(*steps.DetectiveStep)(nil),        // 6: protos.steps.DetectiveStep
	(*steps.TransformStep)(nil),        // 7: protos.steps.TransformStep
	(*steps.EncodeStep)(nil),           // 8: protos.steps.EncodeStep
	(*steps.DecodeStep)(nil),           // 9: protos.steps.DecodeStep
	(*steps.CustomStep)(nil),           // 10: protos.steps.CustomStep
	(*steps.HttpRequestStep)(nil),      // 11: protos.steps.HttpRequestStep
	(*steps.KVStep)(nil),               // 12: protos.steps.KVStep
	(*steps.InferSchemaStep)(nil),      // 13: protos.steps.InferSchemaStep
	(*steps.ValidJSONStep)(nil),        // 14: protos.steps.ValidJSONStep
	(*steps.SchemaValidationStep)(nil), // 15: protos.steps.SchemaValidationStep
}
var file_sp_pipeline_proto_depIdxs = []int32{
	3,  // 0: protos.Pipeline.steps:type_name -> protos.PipelineStep
	5,  // 1: protos.Pipeline._notification_configs:type_name -> protos.NotificationConfig
	0,  // 2: protos.PipelineStepConditions.abort:type_name -> protos.AbortCondition
	4,  // 3: protos.PipelineStepConditions.metadata:type_name -> protos.PipelineStepConditions.MetadataEntry
	2,  // 4: protos.PipelineStep.on_true:type_name -> protos.PipelineStepConditions
	2,  // 5: protos.PipelineStep.on_false:type_name -> protos.PipelineStepConditions
	2,  // 6: protos.PipelineStep.on_error:type_name -> protos.PipelineStepConditions
	6,  // 7: protos.PipelineStep.detective:type_name -> protos.steps.DetectiveStep
	7,  // 8: protos.PipelineStep.transform:type_name -> protos.steps.TransformStep
	8,  // 9: protos.PipelineStep.encode:type_name -> protos.steps.EncodeStep
	9,  // 10: protos.PipelineStep.decode:type_name -> protos.steps.DecodeStep
	10, // 11: protos.PipelineStep.custom:type_name -> protos.steps.CustomStep
	11, // 12: protos.PipelineStep.http_request:type_name -> protos.steps.HttpRequestStep
	12, // 13: protos.PipelineStep.kv:type_name -> protos.steps.KVStep
	13, // 14: protos.PipelineStep.infer_schema:type_name -> protos.steps.InferSchemaStep
	14, // 15: protos.PipelineStep.valid_json:type_name -> protos.steps.ValidJSONStep
	15, // 16: protos.PipelineStep.schema_validation:type_name -> protos.steps.SchemaValidationStep
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_sp_pipeline_proto_init() }
func file_sp_pipeline_proto_init() {
	if File_sp_pipeline_proto != nil {
		return
	}
	file_sp_notify_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sp_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
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
		file_sp_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineStepConditions); i {
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
		file_sp_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineStep); i {
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
	file_sp_pipeline_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PipelineStep_Detective)(nil),
		(*PipelineStep_Transform)(nil),
		(*PipelineStep_Encode)(nil),
		(*PipelineStep_Decode)(nil),
		(*PipelineStep_Custom)(nil),
		(*PipelineStep_HttpRequest)(nil),
		(*PipelineStep_Kv)(nil),
		(*PipelineStep_InferSchema)(nil),
		(*PipelineStep_ValidJson)(nil),
		(*PipelineStep_SchemaValidation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_pipeline_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_pipeline_proto_goTypes,
		DependencyIndexes: file_sp_pipeline_proto_depIdxs,
		EnumInfos:         file_sp_pipeline_proto_enumTypes,
		MessageInfos:      file_sp_pipeline_proto_msgTypes,
	}.Build()
	File_sp_pipeline_proto = out.File
	file_sp_pipeline_proto_rawDesc = nil
	file_sp_pipeline_proto_goTypes = nil
	file_sp_pipeline_proto_depIdxs = nil
}
