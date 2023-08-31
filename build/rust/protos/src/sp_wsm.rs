// This file is generated by rust-protobuf 3.2.0. Do not edit
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
const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_3_2_0;

///  SDK generates a WASM request and passes this to the WASM func
#[derive(PartialEq,Clone,Default,Debug)]
// @@protoc_insertion_point(message:protos.WASMRequest)
pub struct WASMRequest {
    // message fields
    // @@protoc_insertion_point(field:protos.WASMRequest.step)
    pub step: ::protobuf::MessageField<super::sp_pipeline::PipelineStep>,
    // @@protoc_insertion_point(field:protos.WASMRequest.input)
    pub input: ::std::vec::Vec<u8>,
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
        let mut fields = ::std::vec::Vec::with_capacity(2);
        let mut oneofs = ::std::vec::Vec::with_capacity(0);
        fields.push(::protobuf::reflect::rt::v2::make_message_field_accessor::<_, super::sp_pipeline::PipelineStep>(
            "step",
            |m: &WASMRequest| { &m.step },
            |m: &mut WASMRequest| { &mut m.step },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "input",
            |m: &WASMRequest| { &m.input },
            |m: &mut WASMRequest| { &mut m.input },
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
                    self.input = is.read_bytes()?;
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
        if !self.input.is_empty() {
            my_size += ::protobuf::rt::bytes_size(2, &self.input);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if let Some(v) = self.step.as_ref() {
            ::protobuf::rt::write_message_field_with_cached_size(1, v, os)?;
        }
        if !self.input.is_empty() {
            os.write_bytes(2, &self.input)?;
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
        self.input.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static WASMRequest {
        static instance: WASMRequest = WASMRequest {
            step: ::protobuf::MessageField::none(),
            input: ::std::vec::Vec::new(),
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
#[derive(PartialEq,Clone,Default,Debug)]
// @@protoc_insertion_point(message:protos.WASMResponse)
pub struct WASMResponse {
    // message fields
    // @@protoc_insertion_point(field:protos.WASMResponse.output)
    pub output: ::std::vec::Vec<u8>,
    // @@protoc_insertion_point(field:protos.WASMResponse.exit_code)
    pub exit_code: ::protobuf::EnumOrUnknown<WASMExitCode>,
    // @@protoc_insertion_point(field:protos.WASMResponse.exit_msg)
    pub exit_msg: ::std::string::String,
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
        let mut fields = ::std::vec::Vec::with_capacity(3);
        let mut oneofs = ::std::vec::Vec::with_capacity(0);
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "output",
            |m: &WASMResponse| { &m.output },
            |m: &mut WASMResponse| { &mut m.output },
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
                    self.output = is.read_bytes()?;
                },
                16 => {
                    self.exit_code = is.read_enum_or_unknown()?;
                },
                26 => {
                    self.exit_msg = is.read_string()?;
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
        if !self.output.is_empty() {
            my_size += ::protobuf::rt::bytes_size(1, &self.output);
        }
        if self.exit_code != ::protobuf::EnumOrUnknown::new(WASMExitCode::WASM_EXIT_CODE_UNSET) {
            my_size += ::protobuf::rt::int32_size(2, self.exit_code.value());
        }
        if !self.exit_msg.is_empty() {
            my_size += ::protobuf::rt::string_size(3, &self.exit_msg);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if !self.output.is_empty() {
            os.write_bytes(1, &self.output)?;
        }
        if self.exit_code != ::protobuf::EnumOrUnknown::new(WASMExitCode::WASM_EXIT_CODE_UNSET) {
            os.write_enum(2, ::protobuf::EnumOrUnknown::value(&self.exit_code))?;
        }
        if !self.exit_msg.is_empty() {
            os.write_string(3, &self.exit_msg)?;
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
        self.output.clear();
        self.exit_code = ::protobuf::EnumOrUnknown::new(WASMExitCode::WASM_EXIT_CODE_UNSET);
        self.exit_msg.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static WASMResponse {
        static instance: WASMResponse = WASMResponse {
            output: ::std::vec::Vec::new(),
            exit_code: ::protobuf::EnumOrUnknown::from_i32(0),
            exit_msg: ::std::string::String::new(),
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
    \n\x0csp_wsm.proto\x12\x06protos\x1a\x11sp_pipeline.proto\"M\n\x0bWASMRe\
    quest\x12(\n\x04step\x18\x01\x20\x01(\x0b2\x14.protos.PipelineStepR\x04s\
    tep\x12\x14\n\x05input\x18\x02\x20\x01(\x0cR\x05input\"t\n\x0cWASMRespon\
    se\x12\x16\n\x06output\x18\x01\x20\x01(\x0cR\x06output\x121\n\texit_code\
    \x18\x02\x20\x01(\x0e2\x14.protos.WASMExitCodeR\x08exitCode\x12\x19\n\
    \x08exit_msg\x18\x03\x20\x01(\tR\x07exitMsg*\x83\x01\n\x0cWASMExitCode\
    \x12\x18\n\x14WASM_EXIT_CODE_UNSET\x10\0\x12\x1a\n\x16WASM_EXIT_CODE_SUC\
    CESS\x10\x01\x12\x1a\n\x16WASM_EXIT_CODE_FAILURE\x10\x02\x12!\n\x1dWASM_\
    EXIT_CODE_INTERNAL_ERROR\x10\x03B4Z2github.com/streamdal/snitch-protos/b\
    uild/go/protosJ\xf3\t\n\x06\x12\x04\0\0(\x01\n\x08\n\x01\x0c\x12\x03\0\0\
    \x12\n\x08\n\x01\x02\x12\x03\x02\0\x0f\n\t\n\x02\x03\0\x12\x03\x04\0\x1b\
    \n\x08\n\x01\x08\x12\x03\x06\0I\n\t\n\x02\x08\x0b\x12\x03\x06\0I\n\x9e\
    \x04\n\x02\x05\0\x12\x04\x16\0\x1b\x01\x1a\x91\x04\x20Included\x20in\x20\
    WASM\x20response;\x20the\x20SDK\x20should\x20use\x20the\x20WASMExitCode\
    \x20to\x20determine\n\x20what\x20to\x20do\x20next\x20-\x20should\x20it\
    \x20execute\x20next\x20step,\x20should\x20it\x20notify\x20or\x20should\
    \x20it\n\x20stop\x20executing/abort\x20the\x20rest\x20of\x20the\x20steps\
    \x20in\x20the\x20pipeline.\n\n\x20Example:\n\n\x20a.\x20WASM\x20func\x20\
    returns\x20WASM_EXIT_CODE_FAILURE\x20-\x20read\x20PipelineStep.on_failur\
    e\n\x20conditions\x20to\x20determine\x20what\x20to\x20do\x20next.\n\n\
    \x20b.\x20WASM\x20func\x20returns\x20WASM_EXIT_CODE_SUCCESS\x20-\x20read\
    \x20PipelineStep.on_success\n\x20conditions\x20to\x20determine\x20what\
    \x20to\x20do\x20next.\n\n\x20..\x20and\x20so\x20on.\n\x20protolint:disab\
    le:next\x20ENUM_FIELD_NAMES_PREFIX\n\n\n\n\x03\x05\0\x01\x12\x03\x16\x05\
    \x11\n\x0b\n\x04\x05\0\x02\0\x12\x03\x17\x02\x1b\n\x0c\n\x05\x05\0\x02\0\
    \x01\x12\x03\x17\x02\x16\n\x0c\n\x05\x05\0\x02\0\x02\x12\x03\x17\x19\x1a\
    \n\x0b\n\x04\x05\0\x02\x01\x12\x03\x18\x02\x1d\n\x0c\n\x05\x05\0\x02\x01\
    \x01\x12\x03\x18\x02\x18\n\x0c\n\x05\x05\0\x02\x01\x02\x12\x03\x18\x1b\
    \x1c\nK\n\x04\x05\0\x02\x02\x12\x03\x19\x02\x1d\">\x20Probably\x20need\
    \x20better\x20names\x20for\x20these\x20as\x20FAILURE\x20is\x20too\x20har\
    sh\n\n\x0c\n\x05\x05\0\x02\x02\x01\x12\x03\x19\x02\x18\n\x0c\n\x05\x05\0\
    \x02\x02\x02\x12\x03\x19\x1b\x1c\n\x0b\n\x04\x05\0\x02\x03\x12\x03\x1a\
    \x02$\n\x0c\n\x05\x05\0\x02\x03\x01\x12\x03\x1a\x02\x1f\n\x0c\n\x05\x05\
    \0\x02\x03\x02\x12\x03\x1a\"#\nK\n\x02\x04\0\x12\x04\x1e\0!\x01\x1a?\x20\
    SDK\x20generates\x20a\x20WASM\x20request\x20and\x20passes\x20this\x20to\
    \x20the\x20WASM\x20func\n\n\n\n\x03\x04\0\x01\x12\x03\x1e\x08\x13\n\x0b\
    \n\x04\x04\0\x02\0\x12\x03\x1f\x02\x1f\n\x0c\n\x05\x04\0\x02\0\x06\x12\
    \x03\x1f\x02\x15\n\x0c\n\x05\x04\0\x02\0\x01\x12\x03\x1f\x16\x1a\n\x0c\n\
    \x05\x04\0\x02\0\x03\x12\x03\x1f\x1d\x1e\n\x0b\n\x04\x04\0\x02\x01\x12\
    \x03\x20\x02\x12\n\x0c\n\x05\x04\0\x02\x01\x05\x12\x03\x20\x02\x07\n\x0c\
    \n\x05\x04\0\x02\x01\x01\x12\x03\x20\x08\r\n\x0c\n\x05\x04\0\x02\x01\x03\
    \x12\x03\x20\x10\x11\n,\n\x02\x04\x01\x12\x04$\0(\x01\x1a\x20\x20Returne\
    d\x20by\x20all\x20WASM\x20functions\n\n\n\n\x03\x04\x01\x01\x12\x03$\x08\
    \x14\n\x0b\n\x04\x04\x01\x02\0\x12\x03%\x02\x13\n\x0c\n\x05\x04\x01\x02\
    \0\x05\x12\x03%\x02\x07\n\x0c\n\x05\x04\x01\x02\0\x01\x12\x03%\x08\x0e\n\
    \x0c\n\x05\x04\x01\x02\0\x03\x12\x03%\x11\x12\n\x0b\n\x04\x04\x01\x02\
    \x01\x12\x03&\x02\x1d\n\x0c\n\x05\x04\x01\x02\x01\x06\x12\x03&\x02\x0e\n\
    \x0c\n\x05\x04\x01\x02\x01\x01\x12\x03&\x0f\x18\n\x0c\n\x05\x04\x01\x02\
    \x01\x03\x12\x03&\x1b\x1c\n\x0b\n\x04\x04\x01\x02\x02\x12\x03'\x02\x16\n\
    \x0c\n\x05\x04\x01\x02\x02\x05\x12\x03'\x02\x08\n\x0c\n\x05\x04\x01\x02\
    \x02\x01\x12\x03'\t\x11\n\x0c\n\x05\x04\x01\x02\x02\x03\x12\x03'\x14\x15\
    b\x06proto3\
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
            let mut deps = ::std::vec::Vec::with_capacity(1);
            deps.push(super::sp_pipeline::file_descriptor().clone());
            let mut messages = ::std::vec::Vec::with_capacity(2);
            messages.push(WASMRequest::generated_message_descriptor_data());
            messages.push(WASMResponse::generated_message_descriptor_data());
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
