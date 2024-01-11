// This file is generated by rust-protobuf 3.3.0. Do not edit
// .proto file is parsed by protoc --rust-out=...
// @generated

// https://github.com/rust-lang/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![allow(unused_attributes)]
#![cfg_attr(rustfmt, rustfmt::skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unused_results)]
#![allow(unused_mut)]

//! Generated file from `sp_wsm.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_3_3_0;

///  SDK generates a WASM request and passes this to the WASM func
// @@protoc_insertion_point(message:protos.WASMRequest)
#[derive(PartialEq,Clone,Default,Debug)]
pub struct WASMRequest {
    // message fields
    ///  The actual step that the WASM func will operate on. This is the same step
    ///  that is declared in protos.Pipeline.
    // @@protoc_insertion_point(field:protos.WASMRequest.step)
    pub step: ::protobuf::MessageField<super::sp_pipeline::PipelineStep>,
    ///  Payload data that WASM func will operate on
    // @@protoc_insertion_point(field:protos.WASMRequest.input_payload)
    pub input_payload: ::std::vec::Vec<u8>,
    ///  Potentially filled out result from previous step. If this is first step in
    ///  the pipeline, it will be empty.
    // @@protoc_insertion_point(field:protos.WASMRequest.input_step)
    pub input_step: ::std::option::Option<::std::vec::Vec<u8>>,
    // @@protoc_insertion_point(field:protos.WASMRequest.inter_step_result)
    pub inter_step_result: ::protobuf::MessageField<InterStepResult>,
    // special fields
    // @@protoc_insertion_point(special_field:protos.WASMRequest.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a WASMRequest {
    fn default() -> &'a WASMRequest {
        <WASMRequest as ::protobuf::Message>::default_instance()
    }
}

impl WASMRequest {
    pub fn new() -> WASMRequest {
        ::std::default::Default::default()
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(4);
        let mut oneofs = ::std::vec::Vec::with_capacity(0);
        fields.push(::protobuf::reflect::rt::v2::make_message_field_accessor::<_, super::sp_pipeline::PipelineStep>(
            "step",
            |m: &WASMRequest| { &m.step },
            |m: &mut WASMRequest| { &mut m.step },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "input_payload",
            |m: &WASMRequest| { &m.input_payload },
            |m: &mut WASMRequest| { &mut m.input_payload },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_option_accessor::<_, _>(
            "input_step",
            |m: &WASMRequest| { &m.input_step },
            |m: &mut WASMRequest| { &mut m.input_step },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_message_field_accessor::<_, InterStepResult>(
            "inter_step_result",
            |m: &WASMRequest| { &m.inter_step_result },
            |m: &mut WASMRequest| { &mut m.inter_step_result },
        ));
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<WASMRequest>(
            "WASMRequest",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for WASMRequest {
    const NAME: &'static str = "WASMRequest";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                10 => {
                    ::protobuf::rt::read_singular_message_into_field(is, &mut self.step)?;
                },
                18 => {
                    self.input_payload = is.read_bytes()?;
                },
                26 => {
                    self.input_step = ::std::option::Option::Some(is.read_bytes()?);
                },
                34 => {
                    ::protobuf::rt::read_singular_message_into_field(is, &mut self.inter_step_result)?;
                },
                tag => {
                    ::protobuf::rt::read_unknown_or_skip_group(tag, is, self.special_fields.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u64 {
        let mut my_size = 0;
        if let Some(v) = self.step.as_ref() {
            let len = v.compute_size();
            my_size += 1 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
        }
        if !self.input_payload.is_empty() {
            my_size += ::protobuf::rt::bytes_size(2, &self.input_payload);
        }
        if let Some(v) = self.input_step.as_ref() {
            my_size += ::protobuf::rt::bytes_size(3, &v);
        }
        if let Some(v) = self.inter_step_result.as_ref() {
            let len = v.compute_size();
            my_size += 1 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if let Some(v) = self.step.as_ref() {
            ::protobuf::rt::write_message_field_with_cached_size(1, v, os)?;
        }
        if !self.input_payload.is_empty() {
            os.write_bytes(2, &self.input_payload)?;
        }
        if let Some(v) = self.input_step.as_ref() {
            os.write_bytes(3, v)?;
        }
        if let Some(v) = self.inter_step_result.as_ref() {
            ::protobuf::rt::write_message_field_with_cached_size(4, v, os)?;
        }
        os.write_unknown_fields(self.special_fields.unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn special_fields(&self) -> &::protobuf::SpecialFields {
        &self.special_fields
    }

    fn mut_special_fields(&mut self) -> &mut ::protobuf::SpecialFields {
        &mut self.special_fields
    }

    fn new() -> WASMRequest {
        WASMRequest::new()
    }

    fn clear(&mut self) {
        self.step.clear();
        self.input_payload.clear();
        self.input_step = ::std::option::Option::None;
        self.inter_step_result.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static WASMRequest {
        static instance: WASMRequest = WASMRequest {
            step: ::protobuf::MessageField::none(),
            input_payload: ::std::vec::Vec::new(),
            input_step: ::std::option::Option::None,
            inter_step_result: ::protobuf::MessageField::none(),
            special_fields: ::protobuf::SpecialFields::new(),
        };
        &instance
    }
}

impl ::protobuf::MessageFull for WASMRequest {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("WASMRequest").unwrap()).clone()
    }
}

impl ::std::fmt::Display for WASMRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for WASMRequest {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

///  Returned by all WASM functions
// @@protoc_insertion_point(message:protos.WASMResponse)
#[derive(PartialEq,Clone,Default,Debug)]
pub struct WASMResponse {
    // message fields
    ///  Potentially modified input payload. Concept: All WASM funcs accept an
    ///  input_payload in WASMRequest, WASM func reads input payload, modifies it
    ///  and writes the modified output to output_payload.
    // @@protoc_insertion_point(field:protos.WASMResponse.output_payload)
    pub output_payload: ::std::vec::Vec<u8>,
    ///  Exit code that the WASM func exited with; more info in WASMExitCode's comment
    // @@protoc_insertion_point(field:protos.WASMResponse.exit_code)
    pub exit_code: ::protobuf::EnumOrUnknown<WASMExitCode>,
    ///  Additional info about the reason a specific exit code was returned
    // @@protoc_insertion_point(field:protos.WASMResponse.exit_msg)
    pub exit_msg: ::std::string::String,
    ///  Potential additional step output - ie. if a WASM func is an HTTPGet,
    ///  output_step would contain the HTTP response body; if the WASM func is a
    ///  KVGet, the output_step would be the value of the fetched key.
    // @@protoc_insertion_point(field:protos.WASMResponse.output_step)
    pub output_step: ::std::option::Option<::std::vec::Vec<u8>>,
    // @@protoc_insertion_point(field:protos.WASMResponse.inter_step_result)
    pub inter_step_result: ::protobuf::MessageField<InterStepResult>,
    // special fields
    // @@protoc_insertion_point(special_field:protos.WASMResponse.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a WASMResponse {
    fn default() -> &'a WASMResponse {
        <WASMResponse as ::protobuf::Message>::default_instance()
    }
}

impl WASMResponse {
    pub fn new() -> WASMResponse {
        ::std::default::Default::default()
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(5);
        let mut oneofs = ::std::vec::Vec::with_capacity(0);
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "output_payload",
            |m: &WASMResponse| { &m.output_payload },
            |m: &mut WASMResponse| { &mut m.output_payload },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "exit_code",
            |m: &WASMResponse| { &m.exit_code },
            |m: &mut WASMResponse| { &mut m.exit_code },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "exit_msg",
            |m: &WASMResponse| { &m.exit_msg },
            |m: &mut WASMResponse| { &mut m.exit_msg },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_option_accessor::<_, _>(
            "output_step",
            |m: &WASMResponse| { &m.output_step },
            |m: &mut WASMResponse| { &mut m.output_step },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_message_field_accessor::<_, InterStepResult>(
            "inter_step_result",
            |m: &WASMResponse| { &m.inter_step_result },
            |m: &mut WASMResponse| { &mut m.inter_step_result },
        ));
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<WASMResponse>(
            "WASMResponse",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for WASMResponse {
    const NAME: &'static str = "WASMResponse";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                10 => {
                    self.output_payload = is.read_bytes()?;
                },
                16 => {
                    self.exit_code = is.read_enum_or_unknown()?;
                },
                26 => {
                    self.exit_msg = is.read_string()?;
                },
                34 => {
                    self.output_step = ::std::option::Option::Some(is.read_bytes()?);
                },
                42 => {
                    ::protobuf::rt::read_singular_message_into_field(is, &mut self.inter_step_result)?;
                },
                tag => {
                    ::protobuf::rt::read_unknown_or_skip_group(tag, is, self.special_fields.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u64 {
        let mut my_size = 0;
        if !self.output_payload.is_empty() {
            my_size += ::protobuf::rt::bytes_size(1, &self.output_payload);
        }
        if self.exit_code != ::protobuf::EnumOrUnknown::new(WASMExitCode::WASM_EXIT_CODE_UNSET) {
            my_size += ::protobuf::rt::int32_size(2, self.exit_code.value());
        }
        if !self.exit_msg.is_empty() {
            my_size += ::protobuf::rt::string_size(3, &self.exit_msg);
        }
        if let Some(v) = self.output_step.as_ref() {
            my_size += ::protobuf::rt::bytes_size(4, &v);
        }
        if let Some(v) = self.inter_step_result.as_ref() {
            let len = v.compute_size();
            my_size += 1 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if !self.output_payload.is_empty() {
            os.write_bytes(1, &self.output_payload)?;
        }
        if self.exit_code != ::protobuf::EnumOrUnknown::new(WASMExitCode::WASM_EXIT_CODE_UNSET) {
            os.write_enum(2, ::protobuf::EnumOrUnknown::value(&self.exit_code))?;
        }
        if !self.exit_msg.is_empty() {
            os.write_string(3, &self.exit_msg)?;
        }
        if let Some(v) = self.output_step.as_ref() {
            os.write_bytes(4, v)?;
        }
        if let Some(v) = self.inter_step_result.as_ref() {
            ::protobuf::rt::write_message_field_with_cached_size(5, v, os)?;
        }
        os.write_unknown_fields(self.special_fields.unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn special_fields(&self) -> &::protobuf::SpecialFields {
        &self.special_fields
    }

    fn mut_special_fields(&mut self) -> &mut ::protobuf::SpecialFields {
        &mut self.special_fields
    }

    fn new() -> WASMResponse {
        WASMResponse::new()
    }

    fn clear(&mut self) {
        self.output_payload.clear();
        self.exit_code = ::protobuf::EnumOrUnknown::new(WASMExitCode::WASM_EXIT_CODE_UNSET);
        self.exit_msg.clear();
        self.output_step = ::std::option::Option::None;
        self.inter_step_result.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static WASMResponse {
        static instance: WASMResponse = WASMResponse {
            output_payload: ::std::vec::Vec::new(),
            exit_code: ::protobuf::EnumOrUnknown::from_i32(0),
            exit_msg: ::std::string::String::new(),
            output_step: ::std::option::Option::None,
            inter_step_result: ::protobuf::MessageField::none(),
            special_fields: ::protobuf::SpecialFields::new(),
        };
        &instance
    }
}

impl ::protobuf::MessageFull for WASMResponse {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("WASMResponse").unwrap()).clone()
    }
}

impl ::std::fmt::Display for WASMResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for WASMResponse {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

///  Intended for communicating wasm results between steps.
///  Currently only used for passing results from a Detective Step to a Transform step
// @@protoc_insertion_point(message:protos.InterStepResult)
#[derive(PartialEq,Clone,Default,Debug)]
pub struct InterStepResult {
    // message oneof groups
    pub input_from: ::std::option::Option<inter_step_result::Input_from>,
    // special fields
    // @@protoc_insertion_point(special_field:protos.InterStepResult.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a InterStepResult {
    fn default() -> &'a InterStepResult {
        <InterStepResult as ::protobuf::Message>::default_instance()
    }
}

impl InterStepResult {
    pub fn new() -> InterStepResult {
        ::std::default::Default::default()
    }

    // .protos.steps.DetectiveStepResult detective_result = 1;

    pub fn detective_result(&self) -> &super::sp_steps_detective::DetectiveStepResult {
        match self.input_from {
            ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(ref v)) => v,
            _ => <super::sp_steps_detective::DetectiveStepResult as ::protobuf::Message>::default_instance(),
        }
    }

    pub fn clear_detective_result(&mut self) {
        self.input_from = ::std::option::Option::None;
    }

    pub fn has_detective_result(&self) -> bool {
        match self.input_from {
            ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(..)) => true,
            _ => false,
        }
    }

    // Param is passed by value, moved
    pub fn set_detective_result(&mut self, v: super::sp_steps_detective::DetectiveStepResult) {
        self.input_from = ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(v))
    }

    // Mutable pointer to the field.
    pub fn mut_detective_result(&mut self) -> &mut super::sp_steps_detective::DetectiveStepResult {
        if let ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(_)) = self.input_from {
        } else {
            self.input_from = ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(super::sp_steps_detective::DetectiveStepResult::new()));
        }
        match self.input_from {
            ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(ref mut v)) => v,
            _ => panic!(),
        }
    }

    // Take field
    pub fn take_detective_result(&mut self) -> super::sp_steps_detective::DetectiveStepResult {
        if self.has_detective_result() {
            match self.input_from.take() {
                ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(v)) => v,
                _ => panic!(),
            }
        } else {
            super::sp_steps_detective::DetectiveStepResult::new()
        }
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(1);
        let mut oneofs = ::std::vec::Vec::with_capacity(1);
        fields.push(::protobuf::reflect::rt::v2::make_oneof_message_has_get_mut_set_accessor::<_, super::sp_steps_detective::DetectiveStepResult>(
            "detective_result",
            InterStepResult::has_detective_result,
            InterStepResult::detective_result,
            InterStepResult::mut_detective_result,
            InterStepResult::set_detective_result,
        ));
        oneofs.push(inter_step_result::Input_from::generated_oneof_descriptor_data());
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<InterStepResult>(
            "InterStepResult",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for InterStepResult {
    const NAME: &'static str = "InterStepResult";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                10 => {
                    self.input_from = ::std::option::Option::Some(inter_step_result::Input_from::DetectiveResult(is.read_message()?));
                },
                tag => {
                    ::protobuf::rt::read_unknown_or_skip_group(tag, is, self.special_fields.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u64 {
        let mut my_size = 0;
        if let ::std::option::Option::Some(ref v) = self.input_from {
            match v {
                &inter_step_result::Input_from::DetectiveResult(ref v) => {
                    let len = v.compute_size();
                    my_size += 1 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
                },
            };
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if let ::std::option::Option::Some(ref v) = self.input_from {
            match v {
                &inter_step_result::Input_from::DetectiveResult(ref v) => {
                    ::protobuf::rt::write_message_field_with_cached_size(1, v, os)?;
                },
            };
        }
        os.write_unknown_fields(self.special_fields.unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn special_fields(&self) -> &::protobuf::SpecialFields {
        &self.special_fields
    }

    fn mut_special_fields(&mut self) -> &mut ::protobuf::SpecialFields {
        &mut self.special_fields
    }

    fn new() -> InterStepResult {
        InterStepResult::new()
    }

    fn clear(&mut self) {
        self.input_from = ::std::option::Option::None;
        self.special_fields.clear();
    }

    fn default_instance() -> &'static InterStepResult {
        static instance: InterStepResult = InterStepResult {
            input_from: ::std::option::Option::None,
            special_fields: ::protobuf::SpecialFields::new(),
        };
        &instance
    }
}

impl ::protobuf::MessageFull for InterStepResult {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("InterStepResult").unwrap()).clone()
    }
}

impl ::std::fmt::Display for InterStepResult {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for InterStepResult {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

/// Nested message and enums of message `InterStepResult`
pub mod inter_step_result {

    #[derive(Clone,PartialEq,Debug)]
    #[non_exhaustive]
    // @@protoc_insertion_point(oneof:protos.InterStepResult.input_from)
    pub enum Input_from {
        // @@protoc_insertion_point(oneof_field:protos.InterStepResult.detective_result)
        DetectiveResult(super::super::sp_steps_detective::DetectiveStepResult),
    }

    impl ::protobuf::Oneof for Input_from {
    }

    impl ::protobuf::OneofFull for Input_from {
        fn descriptor() -> ::protobuf::reflect::OneofDescriptor {
            static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::OneofDescriptor> = ::protobuf::rt::Lazy::new();
            descriptor.get(|| <super::InterStepResult as ::protobuf::MessageFull>::descriptor().oneof_by_name("input_from").unwrap()).clone()
        }
    }

    impl Input_from {
        pub(in super) fn generated_oneof_descriptor_data() -> ::protobuf::reflect::GeneratedOneofDescriptorData {
            ::protobuf::reflect::GeneratedOneofDescriptorData::new::<Input_from>("input_from")
        }
    }
}

///  Included in WASM response; the SDK should use the WASMExitCode to determine
///  what to do next - should it execute next step, should it notify or should it
///  stop executing/abort the rest of the steps in the pipeline.
///
///  Example:
///
///  a. WASM func returns WASM_EXIT_CODE_FAILURE - read PipelineStep.on_failure
///  conditions to determine what to do next.
///
///  b. WASM func returns WASM_EXIT_CODE_SUCCESS - read PipelineStep.on_success
///  conditions to determine what to do next.
///
///  .. and so on.
///  protolint:disable:next ENUM_FIELD_NAMES_PREFIX
#[derive(Clone,Copy,PartialEq,Eq,Debug,Hash)]
// @@protoc_insertion_point(enum:protos.WASMExitCode)
pub enum WASMExitCode {
    // @@protoc_insertion_point(enum_value:protos.WASMExitCode.WASM_EXIT_CODE_UNSET)
    WASM_EXIT_CODE_UNSET = 0,
    // @@protoc_insertion_point(enum_value:protos.WASMExitCode.WASM_EXIT_CODE_SUCCESS)
    WASM_EXIT_CODE_SUCCESS = 1,
    // @@protoc_insertion_point(enum_value:protos.WASMExitCode.WASM_EXIT_CODE_FAILURE)
    WASM_EXIT_CODE_FAILURE = 2,
    // @@protoc_insertion_point(enum_value:protos.WASMExitCode.WASM_EXIT_CODE_INTERNAL_ERROR)
    WASM_EXIT_CODE_INTERNAL_ERROR = 3,
}

impl ::protobuf::Enum for WASMExitCode {
    const NAME: &'static str = "WASMExitCode";

    fn value(&self) -> i32 {
        *self as i32
    }

    fn from_i32(value: i32) -> ::std::option::Option<WASMExitCode> {
        match value {
            0 => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_UNSET),
            1 => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_SUCCESS),
            2 => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_FAILURE),
            3 => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_INTERNAL_ERROR),
            _ => ::std::option::Option::None
        }
    }

    fn from_str(str: &str) -> ::std::option::Option<WASMExitCode> {
        match str {
            "WASM_EXIT_CODE_UNSET" => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_UNSET),
            "WASM_EXIT_CODE_SUCCESS" => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_SUCCESS),
            "WASM_EXIT_CODE_FAILURE" => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_FAILURE),
            "WASM_EXIT_CODE_INTERNAL_ERROR" => ::std::option::Option::Some(WASMExitCode::WASM_EXIT_CODE_INTERNAL_ERROR),
            _ => ::std::option::Option::None
        }
    }

    const VALUES: &'static [WASMExitCode] = &[
        WASMExitCode::WASM_EXIT_CODE_UNSET,
        WASMExitCode::WASM_EXIT_CODE_SUCCESS,
        WASMExitCode::WASM_EXIT_CODE_FAILURE,
        WASMExitCode::WASM_EXIT_CODE_INTERNAL_ERROR,
    ];
}

impl ::protobuf::EnumFull for WASMExitCode {
    fn enum_descriptor() -> ::protobuf::reflect::EnumDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::EnumDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().enum_by_package_relative_name("WASMExitCode").unwrap()).clone()
    }

    fn descriptor(&self) -> ::protobuf::reflect::EnumValueDescriptor {
        let index = *self as usize;
        Self::enum_descriptor().value_by_index(index)
    }
}

impl ::std::default::Default for WASMExitCode {
    fn default() -> Self {
        WASMExitCode::WASM_EXIT_CODE_UNSET
    }
}

impl WASMExitCode {
    fn generated_enum_descriptor_data() -> ::protobuf::reflect::GeneratedEnumDescriptorData {
        ::protobuf::reflect::GeneratedEnumDescriptorData::new::<WASMExitCode>("WASMExitCode")
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n\x0csp_wsm.proto\x12\x06protos\x1a\x11sp_pipeline.proto\x1a\x1esteps/s\
    p_steps_detective.proto\"\xef\x01\n\x0bWASMRequest\x12(\n\x04step\x18\
    \x01\x20\x01(\x0b2\x14.protos.PipelineStepR\x04step\x12#\n\rinput_payloa\
    d\x18\x02\x20\x01(\x0cR\x0cinputPayload\x12\"\n\ninput_step\x18\x03\x20\
    \x01(\x0cH\0R\tinputStep\x88\x01\x01\x12H\n\x11inter_step_result\x18\x04\
    \x20\x01(\x0b2\x17.protos.InterStepResultH\x01R\x0finterStepResult\x88\
    \x01\x01B\r\n\x0b_input_stepB\x14\n\x12_inter_step_result\"\x99\x02\n\
    \x0cWASMResponse\x12%\n\x0eoutput_payload\x18\x01\x20\x01(\x0cR\routputP\
    ayload\x121\n\texit_code\x18\x02\x20\x01(\x0e2\x14.protos.WASMExitCodeR\
    \x08exitCode\x12\x19\n\x08exit_msg\x18\x03\x20\x01(\tR\x07exitMsg\x12$\n\
    \x0boutput_step\x18\x04\x20\x01(\x0cH\0R\noutputStep\x88\x01\x01\x12H\n\
    \x11inter_step_result\x18\x05\x20\x01(\x0b2\x17.protos.InterStepResultH\
    \x01R\x0finterStepResult\x88\x01\x01B\x0e\n\x0c_output_stepB\x14\n\x12_i\
    nter_step_result\"o\n\x0fInterStepResult\x12N\n\x10detective_result\x18\
    \x01\x20\x01(\x0b2!.protos.steps.DetectiveStepResultH\0R\x0fdetectiveRes\
    ultB\x0c\n\ninput_from*\x83\x01\n\x0cWASMExitCode\x12\x18\n\x14WASM_EXIT\
    _CODE_UNSET\x10\0\x12\x1a\n\x16WASM_EXIT_CODE_SUCCESS\x10\x01\x12\x1a\n\
    \x16WASM_EXIT_CODE_FAILURE\x10\x02\x12!\n\x1dWASM_EXIT_CODE_INTERNAL_ERR\
    OR\x10\x03B<Z:github.com/streamdal/streamdal/libs/protos/build/go/protos\
    J\xce\x14\n\x06\x12\x04\0\0I\x01\n\x08\n\x01\x0c\x12\x03\0\0\x12\n\x08\n\
    \x01\x02\x12\x03\x02\0\x0f\n\t\n\x02\x03\0\x12\x03\x04\0\x1b\n\t\n\x02\
    \x03\x01\x12\x03\x05\0(\n\x08\n\x01\x08\x12\x03\x07\0Q\n\t\n\x02\x08\x0b\
    \x12\x03\x07\0Q\n\x9e\x04\n\x02\x05\0\x12\x04\x17\0\x1c\x01\x1a\x91\x04\
    \x20Included\x20in\x20WASM\x20response;\x20the\x20SDK\x20should\x20use\
    \x20the\x20WASMExitCode\x20to\x20determine\n\x20what\x20to\x20do\x20next\
    \x20-\x20should\x20it\x20execute\x20next\x20step,\x20should\x20it\x20not\
    ify\x20or\x20should\x20it\n\x20stop\x20executing/abort\x20the\x20rest\
    \x20of\x20the\x20steps\x20in\x20the\x20pipeline.\n\n\x20Example:\n\n\x20\
    a.\x20WASM\x20func\x20returns\x20WASM_EXIT_CODE_FAILURE\x20-\x20read\x20\
    PipelineStep.on_failure\n\x20conditions\x20to\x20determine\x20what\x20to\
    \x20do\x20next.\n\n\x20b.\x20WASM\x20func\x20returns\x20WASM_EXIT_CODE_S\
    UCCESS\x20-\x20read\x20PipelineStep.on_success\n\x20conditions\x20to\x20\
    determine\x20what\x20to\x20do\x20next.\n\n\x20..\x20and\x20so\x20on.\n\
    \x20protolint:disable:next\x20ENUM_FIELD_NAMES_PREFIX\n\n\n\n\x03\x05\0\
    \x01\x12\x03\x17\x05\x11\n\x0b\n\x04\x05\0\x02\0\x12\x03\x18\x02\x1b\n\
    \x0c\n\x05\x05\0\x02\0\x01\x12\x03\x18\x02\x16\n\x0c\n\x05\x05\0\x02\0\
    \x02\x12\x03\x18\x19\x1a\n\x0b\n\x04\x05\0\x02\x01\x12\x03\x19\x02\x1d\n\
    \x0c\n\x05\x05\0\x02\x01\x01\x12\x03\x19\x02\x18\n\x0c\n\x05\x05\0\x02\
    \x01\x02\x12\x03\x19\x1b\x1c\nK\n\x04\x05\0\x02\x02\x12\x03\x1a\x02\x1d\
    \">\x20Probably\x20need\x20better\x20names\x20for\x20these\x20as\x20FAIL\
    URE\x20is\x20too\x20harsh\n\n\x0c\n\x05\x05\0\x02\x02\x01\x12\x03\x1a\
    \x02\x18\n\x0c\n\x05\x05\0\x02\x02\x02\x12\x03\x1a\x1b\x1c\n\x0b\n\x04\
    \x05\0\x02\x03\x12\x03\x1b\x02$\n\x0c\n\x05\x05\0\x02\x03\x01\x12\x03\
    \x1b\x02\x1f\n\x0c\n\x05\x05\0\x02\x03\x02\x12\x03\x1b\"#\nK\n\x02\x04\0\
    \x12\x04\x1f\0,\x01\x1a?\x20SDK\x20generates\x20a\x20WASM\x20request\x20\
    and\x20passes\x20this\x20to\x20the\x20WASM\x20func\n\n\n\n\x03\x04\0\x01\
    \x12\x03\x1f\x08\x13\n~\n\x04\x04\0\x02\0\x12\x03\"\x02\x1f\x1aq\x20The\
    \x20actual\x20step\x20that\x20the\x20WASM\x20func\x20will\x20operate\x20\
    on.\x20This\x20is\x20the\x20same\x20step\n\x20that\x20is\x20declared\x20\
    in\x20protos.Pipeline.\n\n\x0c\n\x05\x04\0\x02\0\x06\x12\x03\"\x02\x15\n\
    \x0c\n\x05\x04\0\x02\0\x01\x12\x03\"\x16\x1a\n\x0c\n\x05\x04\0\x02\0\x03\
    \x12\x03\"\x1d\x1e\n:\n\x04\x04\0\x02\x01\x12\x03%\x02\x1a\x1a-\x20Paylo\
    ad\x20data\x20that\x20WASM\x20func\x20will\x20operate\x20on\n\n\x0c\n\
    \x05\x04\0\x02\x01\x05\x12\x03%\x02\x07\n\x0c\n\x05\x04\0\x02\x01\x01\
    \x12\x03%\x08\x15\n\x0c\n\x05\x04\0\x02\x01\x03\x12\x03%\x18\x19\nz\n\
    \x04\x04\0\x02\x02\x12\x03)\x02\x20\x1am\x20Potentially\x20filled\x20out\
    \x20result\x20from\x20previous\x20step.\x20If\x20this\x20is\x20first\x20\
    step\x20in\n\x20the\x20pipeline,\x20it\x20will\x20be\x20empty.\n\n\x0c\n\
    \x05\x04\0\x02\x02\x04\x12\x03)\x02\n\n\x0c\n\x05\x04\0\x02\x02\x05\x12\
    \x03)\x0b\x10\n\x0c\n\x05\x04\0\x02\x02\x01\x12\x03)\x11\x1b\n\x0c\n\x05\
    \x04\0\x02\x02\x03\x12\x03)\x1e\x1f\n\x0b\n\x04\x04\0\x02\x03\x12\x03+\
    \x021\n\x0c\n\x05\x04\0\x02\x03\x04\x12\x03+\x02\n\n\x0c\n\x05\x04\0\x02\
    \x03\x06\x12\x03+\x0b\x1a\n\x0c\n\x05\x04\0\x02\x03\x01\x12\x03+\x1b,\n\
    \x0c\n\x05\x04\0\x02\x03\x03\x12\x03+/0\n,\n\x02\x04\x01\x12\x04/\0A\x01\
    \x1a\x20\x20Returned\x20by\x20all\x20WASM\x20functions\n\n\n\n\x03\x04\
    \x01\x01\x12\x03/\x08\x14\n\xd2\x01\n\x04\x04\x01\x02\0\x12\x033\x02\x1b\
    \x1a\xc4\x01\x20Potentially\x20modified\x20input\x20payload.\x20Concept:\
    \x20All\x20WASM\x20funcs\x20accept\x20an\n\x20input_payload\x20in\x20WAS\
    MRequest,\x20WASM\x20func\x20reads\x20input\x20payload,\x20modifies\x20i\
    t\n\x20and\x20writes\x20the\x20modified\x20output\x20to\x20output_payloa\
    d.\n\n\x0c\n\x05\x04\x01\x02\0\x05\x12\x033\x02\x07\n\x0c\n\x05\x04\x01\
    \x02\0\x01\x12\x033\x08\x16\n\x0c\n\x05\x04\x01\x02\0\x03\x12\x033\x19\
    \x1a\n\\\n\x04\x04\x01\x02\x01\x12\x036\x02\x1d\x1aO\x20Exit\x20code\x20\
    that\x20the\x20WASM\x20func\x20exited\x20with;\x20more\x20info\x20in\x20\
    WASMExitCode's\x20comment\n\n\x0c\n\x05\x04\x01\x02\x01\x06\x12\x036\x02\
    \x0e\n\x0c\n\x05\x04\x01\x02\x01\x01\x12\x036\x0f\x18\n\x0c\n\x05\x04\
    \x01\x02\x01\x03\x12\x036\x1b\x1c\nQ\n\x04\x04\x01\x02\x02\x12\x039\x02\
    \x16\x1aD\x20Additional\x20info\x20about\x20the\x20reason\x20a\x20specif\
    ic\x20exit\x20code\x20was\x20returned\n\n\x0c\n\x05\x04\x01\x02\x02\x05\
    \x12\x039\x02\x08\n\x0c\n\x05\x04\x01\x02\x02\x01\x12\x039\t\x11\n\x0c\n\
    \x05\x04\x01\x02\x02\x03\x12\x039\x14\x15\n\xdc\x01\n\x04\x04\x01\x02\
    \x03\x12\x03>\x02!\x1a\xce\x01\x20Potential\x20additional\x20step\x20out\
    put\x20-\x20ie.\x20if\x20a\x20WASM\x20func\x20is\x20an\x20HTTPGet,\n\x20\
    output_step\x20would\x20contain\x20the\x20HTTP\x20response\x20body;\x20i\
    f\x20the\x20WASM\x20func\x20is\x20a\n\x20KVGet,\x20the\x20output_step\
    \x20would\x20be\x20the\x20value\x20of\x20the\x20fetched\x20key.\n\n\x0c\
    \n\x05\x04\x01\x02\x03\x04\x12\x03>\x02\n\n\x0c\n\x05\x04\x01\x02\x03\
    \x05\x12\x03>\x0b\x10\n\x0c\n\x05\x04\x01\x02\x03\x01\x12\x03>\x11\x1c\n\
    \x0c\n\x05\x04\x01\x02\x03\x03\x12\x03>\x1f\x20\n\x0b\n\x04\x04\x01\x02\
    \x04\x12\x03@\x021\n\x0c\n\x05\x04\x01\x02\x04\x04\x12\x03@\x02\n\n\x0c\
    \n\x05\x04\x01\x02\x04\x06\x12\x03@\x0b\x1a\n\x0c\n\x05\x04\x01\x02\x04\
    \x01\x12\x03@\x1b,\n\x0c\n\x05\x04\x01\x02\x04\x03\x12\x03@/0\n\x98\x01\
    \n\x02\x04\x02\x12\x04E\0I\x01\x1a\x8b\x01\x20Intended\x20for\x20communi\
    cating\x20wasm\x20results\x20between\x20steps.\n\x20Currently\x20only\
    \x20used\x20for\x20passing\x20results\x20from\x20a\x20Detective\x20Step\
    \x20to\x20a\x20Transform\x20step\n\n\n\n\x03\x04\x02\x01\x12\x03E\x08\
    \x17\n\x0c\n\x04\x04\x02\x08\0\x12\x04F\x02H\x03\n\x0c\n\x05\x04\x02\x08\
    \0\x01\x12\x03F\x08\x12\n\x0b\n\x04\x04\x02\x02\0\x12\x03G\x043\n\x0c\n\
    \x05\x04\x02\x02\0\x06\x12\x03G\x04\x1d\n\x0c\n\x05\x04\x02\x02\0\x01\
    \x12\x03G\x1e.\n\x0c\n\x05\x04\x02\x02\0\x03\x12\x03G12b\x06proto3\
";

/// `FileDescriptorProto` object which was a source for this generated file
fn file_descriptor_proto() -> &'static ::protobuf::descriptor::FileDescriptorProto {
    static file_descriptor_proto_lazy: ::protobuf::rt::Lazy<::protobuf::descriptor::FileDescriptorProto> = ::protobuf::rt::Lazy::new();
    file_descriptor_proto_lazy.get(|| {
        ::protobuf::Message::parse_from_bytes(file_descriptor_proto_data).unwrap()
    })
}

/// `FileDescriptor` object which allows dynamic access to files
pub fn file_descriptor() -> &'static ::protobuf::reflect::FileDescriptor {
    static generated_file_descriptor_lazy: ::protobuf::rt::Lazy<::protobuf::reflect::GeneratedFileDescriptor> = ::protobuf::rt::Lazy::new();
    static file_descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::FileDescriptor> = ::protobuf::rt::Lazy::new();
    file_descriptor.get(|| {
        let generated_file_descriptor = generated_file_descriptor_lazy.get(|| {
            let mut deps = ::std::vec::Vec::with_capacity(2);
            deps.push(super::sp_pipeline::file_descriptor().clone());
            deps.push(super::sp_steps_detective::file_descriptor().clone());
            let mut messages = ::std::vec::Vec::with_capacity(3);
            messages.push(WASMRequest::generated_message_descriptor_data());
            messages.push(WASMResponse::generated_message_descriptor_data());
            messages.push(InterStepResult::generated_message_descriptor_data());
            let mut enums = ::std::vec::Vec::with_capacity(1);
            enums.push(WASMExitCode::generated_enum_descriptor_data());
            ::protobuf::reflect::GeneratedFileDescriptor::new_generated(
                file_descriptor_proto(),
                deps,
                messages,
                enums,
            )
        });
        ::protobuf::reflect::FileDescriptor::new_generated_2(generated_file_descriptor)
    })
}
