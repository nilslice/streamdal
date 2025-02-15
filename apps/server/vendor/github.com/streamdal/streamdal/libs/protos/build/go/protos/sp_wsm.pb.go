// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sp_wsm.proto

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

// Included in Wasm response; the SDK should use the WASMExitCode to determine
// what to do next - should it execute next step, should it notify or should it
// stop execution/abort the rest of the steps in current or all pipelines.
//
// Example:
//
// a. Wasm func returns WASM_EXIT_CODE_FALSE - read PipelineStep.on_false
// conditions to determine what to do next.
//
// b. Wasm func returns WASM_EXIT_CODE_TRUE - read PipelineStep.on_true
// conditions to determine what to do next.
//
// .. and so on.
// TODO: This might be a dupe - should Wasm use ExecStatus instead of this?
// protolint:disable:next ENUM_FIELD_NAMES_PREFIX
type WASMExitCode int32

const (
	WASMExitCode_WASM_EXIT_CODE_UNSET WASMExitCode = 0
	WASMExitCode_WASM_EXIT_CODE_TRUE  WASMExitCode = 1
	WASMExitCode_WASM_EXIT_CODE_FALSE WASMExitCode = 2
	WASMExitCode_WASM_EXIT_CODE_ERROR WASMExitCode = 3
)

// Enum value maps for WASMExitCode.
var (
	WASMExitCode_name = map[int32]string{
		0: "WASM_EXIT_CODE_UNSET",
		1: "WASM_EXIT_CODE_TRUE",
		2: "WASM_EXIT_CODE_FALSE",
		3: "WASM_EXIT_CODE_ERROR",
	}
	WASMExitCode_value = map[string]int32{
		"WASM_EXIT_CODE_UNSET": 0,
		"WASM_EXIT_CODE_TRUE":  1,
		"WASM_EXIT_CODE_FALSE": 2,
		"WASM_EXIT_CODE_ERROR": 3,
	}
)

func (x WASMExitCode) Enum() *WASMExitCode {
	p := new(WASMExitCode)
	*p = x
	return p
}

func (x WASMExitCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WASMExitCode) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_wsm_proto_enumTypes[0].Descriptor()
}

func (WASMExitCode) Type() protoreflect.EnumType {
	return &file_sp_wsm_proto_enumTypes[0]
}

func (x WASMExitCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WASMExitCode.Descriptor instead.
func (WASMExitCode) EnumDescriptor() ([]byte, []int) {
	return file_sp_wsm_proto_rawDescGZIP(), []int{0}
}

// SDK generates a WASM request and passes this to the WASM func
type WASMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual step that the WASM func will operate on. This is the same step
	// that is declared in protos.Pipeline.
	Step *PipelineStep `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
	// Payload data that WASM func will operate on
	InputPayload []byte `protobuf:"bytes,2,opt,name=input_payload,json=inputPayload,proto3" json:"input_payload,omitempty"`
	// Potentially filled out result from previous step. If this is first step in
	// the pipeline, it will be empty.
	InputStep []byte `protobuf:"bytes,3,opt,name=input_step,json=inputStep,proto3,oneof" json:"input_step,omitempty"`
	// Potential input from a previous step if `Step.Dynamic == true`
	// This is used for communicating data between steps.
	// For example, when trying to find email addresses in a payload and
	// then passing on the results to a transform step to obfuscate them
	InterStepResult *InterStepResult `protobuf:"bytes,4,opt,name=inter_step_result,json=interStepResult,proto3,oneof" json:"inter_step_result,omitempty"`
}

func (x *WASMRequest) Reset() {
	*x = WASMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_wsm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WASMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WASMRequest) ProtoMessage() {}

func (x *WASMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_wsm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WASMRequest.ProtoReflect.Descriptor instead.
func (*WASMRequest) Descriptor() ([]byte, []int) {
	return file_sp_wsm_proto_rawDescGZIP(), []int{0}
}

func (x *WASMRequest) GetStep() *PipelineStep {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *WASMRequest) GetInputPayload() []byte {
	if x != nil {
		return x.InputPayload
	}
	return nil
}

func (x *WASMRequest) GetInputStep() []byte {
	if x != nil {
		return x.InputStep
	}
	return nil
}

func (x *WASMRequest) GetInterStepResult() *InterStepResult {
	if x != nil {
		return x.InterStepResult
	}
	return nil
}

// Returned by all WASM functions
type WASMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Potentially modified input payload. Concept: All WASM funcs accept an
	// input_payload in WASMRequest, WASM func reads input payload, modifies it
	// and writes the modified output to output_payload.
	OutputPayload []byte `protobuf:"bytes,1,opt,name=output_payload,json=outputPayload,proto3" json:"output_payload,omitempty"`
	// Exit code that the WASM func exited with; more info in WASMExitCode's comment
	ExitCode WASMExitCode `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3,enum=protos.WASMExitCode" json:"exit_code,omitempty"`
	// Additional info about the reason a specific exit code was returned
	ExitMsg string `protobuf:"bytes,3,opt,name=exit_msg,json=exitMsg,proto3" json:"exit_msg,omitempty"`
	// Potential additional step output - ie. if a WASM func is an HTTPGet,
	// output_step would contain the HTTP response body; if the WASM func is a
	// KVGet, the output_step would be the value of the fetched key.
	OutputStep []byte `protobuf:"bytes,4,opt,name=output_step,json=outputStep,proto3,oneof" json:"output_step,omitempty"`
	// If `Step.Dynamic == true`, this field should be filled out by the WASM module
	// This is used for communicating data between steps.
	// For example, when trying to find email addresses in a payload and
	// then passing on the results to a transform step to obfuscate them
	InterStepResult *InterStepResult `protobuf:"bytes,5,opt,name=inter_step_result,json=interStepResult,proto3,oneof" json:"inter_step_result,omitempty"`
}

func (x *WASMResponse) Reset() {
	*x = WASMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_wsm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WASMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WASMResponse) ProtoMessage() {}

func (x *WASMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sp_wsm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WASMResponse.ProtoReflect.Descriptor instead.
func (*WASMResponse) Descriptor() ([]byte, []int) {
	return file_sp_wsm_proto_rawDescGZIP(), []int{1}
}

func (x *WASMResponse) GetOutputPayload() []byte {
	if x != nil {
		return x.OutputPayload
	}
	return nil
}

func (x *WASMResponse) GetExitCode() WASMExitCode {
	if x != nil {
		return x.ExitCode
	}
	return WASMExitCode_WASM_EXIT_CODE_UNSET
}

func (x *WASMResponse) GetExitMsg() string {
	if x != nil {
		return x.ExitMsg
	}
	return ""
}

func (x *WASMResponse) GetOutputStep() []byte {
	if x != nil {
		return x.OutputStep
	}
	return nil
}

func (x *WASMResponse) GetInterStepResult() *InterStepResult {
	if x != nil {
		return x.InterStepResult
	}
	return nil
}

// Intended for communicating wasm results between steps.
// Currently only used for passing results from a Detective Step to a Transform step
type InterStepResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to InputFrom:
	//
	//	*InterStepResult_DetectiveResult
	InputFrom isInterStepResult_InputFrom `protobuf_oneof:"input_from"`
}

func (x *InterStepResult) Reset() {
	*x = InterStepResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_wsm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterStepResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterStepResult) ProtoMessage() {}

func (x *InterStepResult) ProtoReflect() protoreflect.Message {
	mi := &file_sp_wsm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterStepResult.ProtoReflect.Descriptor instead.
func (*InterStepResult) Descriptor() ([]byte, []int) {
	return file_sp_wsm_proto_rawDescGZIP(), []int{2}
}

func (m *InterStepResult) GetInputFrom() isInterStepResult_InputFrom {
	if m != nil {
		return m.InputFrom
	}
	return nil
}

func (x *InterStepResult) GetDetectiveResult() *steps.DetectiveStepResult {
	if x, ok := x.GetInputFrom().(*InterStepResult_DetectiveResult); ok {
		return x.DetectiveResult
	}
	return nil
}

type isInterStepResult_InputFrom interface {
	isInterStepResult_InputFrom()
}

type InterStepResult_DetectiveResult struct {
	DetectiveResult *steps.DetectiveStepResult `protobuf:"bytes,1,opt,name=detective_result,json=detectiveResult,proto3,oneof"`
}

func (*InterStepResult_DetectiveResult) isInterStepResult_InputFrom() {}

var File_sp_wsm_proto protoreflect.FileDescriptor

var file_sp_wsm_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x5f, 0x77, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x11, 0x73, 0x70, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x2f, 0x73, 0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0b, 0x57, 0x41,
	0x53, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x65, 0x70, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x11,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x48, 0x01, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x99, 0x02, 0x0a, 0x0c,
	0x57, 0x41, 0x53, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x57, 0x41, 0x53, 0x4d, 0x45, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x65, 0x78,
	0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x69, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x24, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x74, 0x65, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x53, 0x74, 0x65, 0x70, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x01, 0x52, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x74, 0x65,
	0x70, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x65, 0x70,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6f, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4e, 0x0a, 0x10, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x65,
	0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x2a, 0x75, 0x0a, 0x0c, 0x57, 0x41, 0x53, 0x4d,
	0x45, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x41, 0x53, 0x4d,
	0x5f, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x41, 0x53, 0x4d, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x57,
	0x41, 0x53, 0x4d, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x41,
	0x4c, 0x53, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x41, 0x53, 0x4d, 0x5f, 0x45, 0x58,
	0x49, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42,
	0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61,
	0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_wsm_proto_rawDescOnce sync.Once
	file_sp_wsm_proto_rawDescData = file_sp_wsm_proto_rawDesc
)

func file_sp_wsm_proto_rawDescGZIP() []byte {
	file_sp_wsm_proto_rawDescOnce.Do(func() {
		file_sp_wsm_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_wsm_proto_rawDescData)
	})
	return file_sp_wsm_proto_rawDescData
}

var file_sp_wsm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sp_wsm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sp_wsm_proto_goTypes = []interface{}{
	(WASMExitCode)(0),                 // 0: protos.WASMExitCode
	(*WASMRequest)(nil),               // 1: protos.WASMRequest
	(*WASMResponse)(nil),              // 2: protos.WASMResponse
	(*InterStepResult)(nil),           // 3: protos.InterStepResult
	(*PipelineStep)(nil),              // 4: protos.PipelineStep
	(*steps.DetectiveStepResult)(nil), // 5: protos.steps.DetectiveStepResult
}
var file_sp_wsm_proto_depIdxs = []int32{
	4, // 0: protos.WASMRequest.step:type_name -> protos.PipelineStep
	3, // 1: protos.WASMRequest.inter_step_result:type_name -> protos.InterStepResult
	0, // 2: protos.WASMResponse.exit_code:type_name -> protos.WASMExitCode
	3, // 3: protos.WASMResponse.inter_step_result:type_name -> protos.InterStepResult
	5, // 4: protos.InterStepResult.detective_result:type_name -> protos.steps.DetectiveStepResult
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sp_wsm_proto_init() }
func file_sp_wsm_proto_init() {
	if File_sp_wsm_proto != nil {
		return
	}
	file_sp_pipeline_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sp_wsm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WASMRequest); i {
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
		file_sp_wsm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WASMResponse); i {
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
		file_sp_wsm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterStepResult); i {
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
	file_sp_wsm_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_sp_wsm_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_sp_wsm_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*InterStepResult_DetectiveResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_wsm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_wsm_proto_goTypes,
		DependencyIndexes: file_sp_wsm_proto_depIdxs,
		EnumInfos:         file_sp_wsm_proto_enumTypes,
		MessageInfos:      file_sp_wsm_proto_msgTypes,
	}.Build()
	File_sp_wsm_proto = out.File
	file_sp_wsm_proto_rawDesc = nil
	file_sp_wsm_proto_goTypes = nil
	file_sp_wsm_proto_depIdxs = nil
}
