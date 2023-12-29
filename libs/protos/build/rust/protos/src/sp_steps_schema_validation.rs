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

//! Generated file from `steps/sp_steps_schema_validation.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_3_3_0;

// @@protoc_insertion_point(message:protos.steps.SchemaValidationStep)
#[derive(PartialEq,Clone,Default,Debug)]
pub struct SchemaValidationStep {
    // message fields
    // @@protoc_insertion_point(field:protos.steps.SchemaValidationStep.type)
    pub type_: ::protobuf::EnumOrUnknown<SchemaValidationType>,
    // @@protoc_insertion_point(field:protos.steps.SchemaValidationStep.condition)
    pub condition: ::protobuf::EnumOrUnknown<SchemaValidationCondition>,
    // message oneof groups
    pub options: ::std::option::Option<schema_validation_step::Options>,
    // special fields
    // @@protoc_insertion_point(special_field:protos.steps.SchemaValidationStep.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a SchemaValidationStep {
    fn default() -> &'a SchemaValidationStep {
        <SchemaValidationStep as ::protobuf::Message>::default_instance()
    }
}

impl SchemaValidationStep {
    pub fn new() -> SchemaValidationStep {
        ::std::default::Default::default()
    }

    // .protos.steps.SchemaValidationJSONSchema json_schema = 101;

    pub fn json_schema(&self) -> &SchemaValidationJSONSchema {
        match self.options {
            ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(ref v)) => v,
            _ => <SchemaValidationJSONSchema as ::protobuf::Message>::default_instance(),
        }
    }

    pub fn clear_json_schema(&mut self) {
        self.options = ::std::option::Option::None;
    }

    pub fn has_json_schema(&self) -> bool {
        match self.options {
            ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(..)) => true,
            _ => false,
        }
    }

    // Param is passed by value, moved
    pub fn set_json_schema(&mut self, v: SchemaValidationJSONSchema) {
        self.options = ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(v))
    }

    // Mutable pointer to the field.
    pub fn mut_json_schema(&mut self) -> &mut SchemaValidationJSONSchema {
        if let ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(_)) = self.options {
        } else {
            self.options = ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(SchemaValidationJSONSchema::new()));
        }
        match self.options {
            ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(ref mut v)) => v,
            _ => panic!(),
        }
    }

    // Take field
    pub fn take_json_schema(&mut self) -> SchemaValidationJSONSchema {
        if self.has_json_schema() {
            match self.options.take() {
                ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(v)) => v,
                _ => panic!(),
            }
        } else {
            SchemaValidationJSONSchema::new()
        }
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(3);
        let mut oneofs = ::std::vec::Vec::with_capacity(1);
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "type",
            |m: &SchemaValidationStep| { &m.type_ },
            |m: &mut SchemaValidationStep| { &mut m.type_ },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "condition",
            |m: &SchemaValidationStep| { &m.condition },
            |m: &mut SchemaValidationStep| { &mut m.condition },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_oneof_message_has_get_mut_set_accessor::<_, SchemaValidationJSONSchema>(
            "json_schema",
            SchemaValidationStep::has_json_schema,
            SchemaValidationStep::json_schema,
            SchemaValidationStep::mut_json_schema,
            SchemaValidationStep::set_json_schema,
        ));
        oneofs.push(schema_validation_step::Options::generated_oneof_descriptor_data());
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<SchemaValidationStep>(
            "SchemaValidationStep",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for SchemaValidationStep {
    const NAME: &'static str = "SchemaValidationStep";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                8 => {
                    self.type_ = is.read_enum_or_unknown()?;
                },
                16 => {
                    self.condition = is.read_enum_or_unknown()?;
                },
                810 => {
                    self.options = ::std::option::Option::Some(schema_validation_step::Options::JsonSchema(is.read_message()?));
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
        if self.type_ != ::protobuf::EnumOrUnknown::new(SchemaValidationType::SCHEMA_VALIDATION_TYPE_UNKNOWN) {
            my_size += ::protobuf::rt::int32_size(1, self.type_.value());
        }
        if self.condition != ::protobuf::EnumOrUnknown::new(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_UNKNOWN) {
            my_size += ::protobuf::rt::int32_size(2, self.condition.value());
        }
        if let ::std::option::Option::Some(ref v) = self.options {
            match v {
                &schema_validation_step::Options::JsonSchema(ref v) => {
                    let len = v.compute_size();
                    my_size += 2 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
                },
            };
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if self.type_ != ::protobuf::EnumOrUnknown::new(SchemaValidationType::SCHEMA_VALIDATION_TYPE_UNKNOWN) {
            os.write_enum(1, ::protobuf::EnumOrUnknown::value(&self.type_))?;
        }
        if self.condition != ::protobuf::EnumOrUnknown::new(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_UNKNOWN) {
            os.write_enum(2, ::protobuf::EnumOrUnknown::value(&self.condition))?;
        }
        if let ::std::option::Option::Some(ref v) = self.options {
            match v {
                &schema_validation_step::Options::JsonSchema(ref v) => {
                    ::protobuf::rt::write_message_field_with_cached_size(101, v, os)?;
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

    fn new() -> SchemaValidationStep {
        SchemaValidationStep::new()
    }

    fn clear(&mut self) {
        self.type_ = ::protobuf::EnumOrUnknown::new(SchemaValidationType::SCHEMA_VALIDATION_TYPE_UNKNOWN);
        self.condition = ::protobuf::EnumOrUnknown::new(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_UNKNOWN);
        self.options = ::std::option::Option::None;
        self.special_fields.clear();
    }

    fn default_instance() -> &'static SchemaValidationStep {
        static instance: SchemaValidationStep = SchemaValidationStep {
            type_: ::protobuf::EnumOrUnknown::from_i32(0),
            condition: ::protobuf::EnumOrUnknown::from_i32(0),
            options: ::std::option::Option::None,
            special_fields: ::protobuf::SpecialFields::new(),
        };
        &instance
    }
}

impl ::protobuf::MessageFull for SchemaValidationStep {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("SchemaValidationStep").unwrap()).clone()
    }
}

impl ::std::fmt::Display for SchemaValidationStep {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for SchemaValidationStep {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

/// Nested message and enums of message `SchemaValidationStep`
pub mod schema_validation_step {

    #[derive(Clone,PartialEq,Debug)]
    #[non_exhaustive]
    // @@protoc_insertion_point(oneof:protos.steps.SchemaValidationStep.options)
    pub enum Options {
        // @@protoc_insertion_point(oneof_field:protos.steps.SchemaValidationStep.json_schema)
        JsonSchema(super::SchemaValidationJSONSchema),
    }

    impl ::protobuf::Oneof for Options {
    }

    impl ::protobuf::OneofFull for Options {
        fn descriptor() -> ::protobuf::reflect::OneofDescriptor {
            static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::OneofDescriptor> = ::protobuf::rt::Lazy::new();
            descriptor.get(|| <super::SchemaValidationStep as ::protobuf::MessageFull>::descriptor().oneof_by_name("options").unwrap()).clone()
        }
    }

    impl Options {
        pub(in super) fn generated_oneof_descriptor_data() -> ::protobuf::reflect::GeneratedOneofDescriptorData {
            ::protobuf::reflect::GeneratedOneofDescriptorData::new::<Options>("options")
        }
    }
}

// @@protoc_insertion_point(message:protos.steps.SchemaValidationJSONSchema)
#[derive(PartialEq,Clone,Default,Debug)]
pub struct SchemaValidationJSONSchema {
    // message fields
    // @@protoc_insertion_point(field:protos.steps.SchemaValidationJSONSchema.json_schema)
    pub json_schema: ::std::vec::Vec<u8>,
    // special fields
    // @@protoc_insertion_point(special_field:protos.steps.SchemaValidationJSONSchema.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a SchemaValidationJSONSchema {
    fn default() -> &'a SchemaValidationJSONSchema {
        <SchemaValidationJSONSchema as ::protobuf::Message>::default_instance()
    }
}

impl SchemaValidationJSONSchema {
    pub fn new() -> SchemaValidationJSONSchema {
        ::std::default::Default::default()
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(1);
        let mut oneofs = ::std::vec::Vec::with_capacity(0);
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "json_schema",
            |m: &SchemaValidationJSONSchema| { &m.json_schema },
            |m: &mut SchemaValidationJSONSchema| { &mut m.json_schema },
        ));
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<SchemaValidationJSONSchema>(
            "SchemaValidationJSONSchema",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for SchemaValidationJSONSchema {
    const NAME: &'static str = "SchemaValidationJSONSchema";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                10 => {
                    self.json_schema = is.read_bytes()?;
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
        if !self.json_schema.is_empty() {
            my_size += ::protobuf::rt::bytes_size(1, &self.json_schema);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if !self.json_schema.is_empty() {
            os.write_bytes(1, &self.json_schema)?;
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

    fn new() -> SchemaValidationJSONSchema {
        SchemaValidationJSONSchema::new()
    }

    fn clear(&mut self) {
        self.json_schema.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static SchemaValidationJSONSchema {
        static instance: SchemaValidationJSONSchema = SchemaValidationJSONSchema {
            json_schema: ::std::vec::Vec::new(),
            special_fields: ::protobuf::SpecialFields::new(),
        };
        &instance
    }
}

impl ::protobuf::MessageFull for SchemaValidationJSONSchema {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("SchemaValidationJSONSchema").unwrap()).clone()
    }
}

impl ::std::fmt::Display for SchemaValidationJSONSchema {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for SchemaValidationJSONSchema {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

///  TODO: expand for protobuf, avro, etc.
#[derive(Clone,Copy,PartialEq,Eq,Debug,Hash)]
// @@protoc_insertion_point(enum:protos.steps.SchemaValidationType)
pub enum SchemaValidationType {
    // @@protoc_insertion_point(enum_value:protos.steps.SchemaValidationType.SCHEMA_VALIDATION_TYPE_UNKNOWN)
    SCHEMA_VALIDATION_TYPE_UNKNOWN = 0,
    // @@protoc_insertion_point(enum_value:protos.steps.SchemaValidationType.SCHEMA_VALIDATION_TYPE_JSONSCHEMA)
    SCHEMA_VALIDATION_TYPE_JSONSCHEMA = 1,
}

impl ::protobuf::Enum for SchemaValidationType {
    const NAME: &'static str = "SchemaValidationType";

    fn value(&self) -> i32 {
        *self as i32
    }

    fn from_i32(value: i32) -> ::std::option::Option<SchemaValidationType> {
        match value {
            0 => ::std::option::Option::Some(SchemaValidationType::SCHEMA_VALIDATION_TYPE_UNKNOWN),
            1 => ::std::option::Option::Some(SchemaValidationType::SCHEMA_VALIDATION_TYPE_JSONSCHEMA),
            _ => ::std::option::Option::None
        }
    }

    fn from_str(str: &str) -> ::std::option::Option<SchemaValidationType> {
        match str {
            "SCHEMA_VALIDATION_TYPE_UNKNOWN" => ::std::option::Option::Some(SchemaValidationType::SCHEMA_VALIDATION_TYPE_UNKNOWN),
            "SCHEMA_VALIDATION_TYPE_JSONSCHEMA" => ::std::option::Option::Some(SchemaValidationType::SCHEMA_VALIDATION_TYPE_JSONSCHEMA),
            _ => ::std::option::Option::None
        }
    }

    const VALUES: &'static [SchemaValidationType] = &[
        SchemaValidationType::SCHEMA_VALIDATION_TYPE_UNKNOWN,
        SchemaValidationType::SCHEMA_VALIDATION_TYPE_JSONSCHEMA,
    ];
}

impl ::protobuf::EnumFull for SchemaValidationType {
    fn enum_descriptor() -> ::protobuf::reflect::EnumDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::EnumDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().enum_by_package_relative_name("SchemaValidationType").unwrap()).clone()
    }

    fn descriptor(&self) -> ::protobuf::reflect::EnumValueDescriptor {
        let index = *self as usize;
        Self::enum_descriptor().value_by_index(index)
    }
}

impl ::std::default::Default for SchemaValidationType {
    fn default() -> Self {
        SchemaValidationType::SCHEMA_VALIDATION_TYPE_UNKNOWN
    }
}

impl SchemaValidationType {
    fn generated_enum_descriptor_data() -> ::protobuf::reflect::GeneratedEnumDescriptorData {
        ::protobuf::reflect::GeneratedEnumDescriptorData::new::<SchemaValidationType>("SchemaValidationType")
    }
}

#[derive(Clone,Copy,PartialEq,Eq,Debug,Hash)]
// @@protoc_insertion_point(enum:protos.steps.SchemaValidationCondition)
pub enum SchemaValidationCondition {
    // @@protoc_insertion_point(enum_value:protos.steps.SchemaValidationCondition.SCHEMA_VALIDATION_CONDITION_UNKNOWN)
    SCHEMA_VALIDATION_CONDITION_UNKNOWN = 0,
    // @@protoc_insertion_point(enum_value:protos.steps.SchemaValidationCondition.SCHEMA_VALIDATION_CONDITION_MATCH)
    SCHEMA_VALIDATION_CONDITION_MATCH = 1,
    // @@protoc_insertion_point(enum_value:protos.steps.SchemaValidationCondition.SCHEMA_VALIDATION_CONDITION_NOT_MATCH)
    SCHEMA_VALIDATION_CONDITION_NOT_MATCH = 2,
}

impl ::protobuf::Enum for SchemaValidationCondition {
    const NAME: &'static str = "SchemaValidationCondition";

    fn value(&self) -> i32 {
        *self as i32
    }

    fn from_i32(value: i32) -> ::std::option::Option<SchemaValidationCondition> {
        match value {
            0 => ::std::option::Option::Some(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_UNKNOWN),
            1 => ::std::option::Option::Some(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_MATCH),
            2 => ::std::option::Option::Some(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_NOT_MATCH),
            _ => ::std::option::Option::None
        }
    }

    fn from_str(str: &str) -> ::std::option::Option<SchemaValidationCondition> {
        match str {
            "SCHEMA_VALIDATION_CONDITION_UNKNOWN" => ::std::option::Option::Some(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_UNKNOWN),
            "SCHEMA_VALIDATION_CONDITION_MATCH" => ::std::option::Option::Some(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_MATCH),
            "SCHEMA_VALIDATION_CONDITION_NOT_MATCH" => ::std::option::Option::Some(SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_NOT_MATCH),
            _ => ::std::option::Option::None
        }
    }

    const VALUES: &'static [SchemaValidationCondition] = &[
        SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_UNKNOWN,
        SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_MATCH,
        SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_NOT_MATCH,
    ];
}

impl ::protobuf::EnumFull for SchemaValidationCondition {
    fn enum_descriptor() -> ::protobuf::reflect::EnumDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::EnumDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().enum_by_package_relative_name("SchemaValidationCondition").unwrap()).clone()
    }

    fn descriptor(&self) -> ::protobuf::reflect::EnumValueDescriptor {
        let index = *self as usize;
        Self::enum_descriptor().value_by_index(index)
    }
}

impl ::std::default::Default for SchemaValidationCondition {
    fn default() -> Self {
        SchemaValidationCondition::SCHEMA_VALIDATION_CONDITION_UNKNOWN
    }
}

impl SchemaValidationCondition {
    fn generated_enum_descriptor_data() -> ::protobuf::reflect::GeneratedEnumDescriptorData {
        ::protobuf::reflect::GeneratedEnumDescriptorData::new::<SchemaValidationCondition>("SchemaValidationCondition")
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n&steps/sp_steps_schema_validation.proto\x12\x0cprotos.steps\"\xed\x01\
    \n\x14SchemaValidationStep\x126\n\x04type\x18\x01\x20\x01(\x0e2\".protos\
    .steps.SchemaValidationTypeR\x04type\x12E\n\tcondition\x18\x02\x20\x01(\
    \x0e2'.protos.steps.SchemaValidationConditionR\tcondition\x12K\n\x0bjson\
    _schema\x18e\x20\x01(\x0b2(.protos.steps.SchemaValidationJSONSchemaH\0R\
    \njsonSchemaB\t\n\x07options\"=\n\x1aSchemaValidationJSONSchema\x12\x1f\
    \n\x0bjson_schema\x18\x01\x20\x01(\x0cR\njsonSchema*a\n\x14SchemaValidat\
    ionType\x12\"\n\x1eSCHEMA_VALIDATION_TYPE_UNKNOWN\x10\0\x12%\n!SCHEMA_VA\
    LIDATION_TYPE_JSONSCHEMA\x10\x01*\x96\x01\n\x19SchemaValidationCondition\
    \x12'\n#SCHEMA_VALIDATION_CONDITION_UNKNOWN\x10\0\x12%\n!SCHEMA_VALIDATI\
    ON_CONDITION_MATCH\x10\x01\x12)\n%SCHEMA_VALIDATION_CONDITION_NOT_MATCH\
    \x10\x02BBZ@github.com/streamdal/streamdal/libs/protos/build/go/protos/s\
    tepsJ\xa7\x05\n\x06\x12\x04\0\0\x20\x01\n\x08\n\x01\x0c\x12\x03\0\0\x12\
    \n\x08\n\x01\x02\x12\x03\x02\0\x15\n\x08\n\x01\x08\x12\x03\x04\0W\n\t\n\
    \x02\x08\x0b\x12\x03\x04\0W\n3\n\x02\x05\0\x12\x04\x07\0\n\x01\x1a'\x20T\
    ODO:\x20expand\x20for\x20protobuf,\x20avro,\x20etc.\n\n\n\n\x03\x05\0\
    \x01\x12\x03\x07\x05\x19\n\x0b\n\x04\x05\0\x02\0\x12\x03\x08\x02%\n\x0c\
    \n\x05\x05\0\x02\0\x01\x12\x03\x08\x02\x20\n\x0c\n\x05\x05\0\x02\0\x02\
    \x12\x03\x08#$\n\x0b\n\x04\x05\0\x02\x01\x12\x03\t\x02(\n\x0c\n\x05\x05\
    \0\x02\x01\x01\x12\x03\t\x02#\n\x0c\n\x05\x05\0\x02\x01\x02\x12\x03\t&'\
    \n\n\n\x02\x05\x01\x12\x04\x0c\0\x12\x01\n\n\n\x03\x05\x01\x01\x12\x03\
    \x0c\x05\x1e\n\x0b\n\x04\x05\x01\x02\0\x12\x03\r\x02*\n\x0c\n\x05\x05\
    \x01\x02\0\x01\x12\x03\r\x02%\n\x0c\n\x05\x05\x01\x02\0\x02\x12\x03\r()\
    \n\x0b\n\x04\x05\x01\x02\x01\x12\x03\x0e\x02(\n\x0c\n\x05\x05\x01\x02\
    \x01\x01\x12\x03\x0e\x02#\n\x0c\n\x05\x05\x01\x02\x01\x02\x12\x03\x0e&'\
    \n3\n\x04\x05\x01\x02\x02\x12\x03\x0f\x02,\"&\x20TODO:\x20backwards\x20c\
    ompat,\x20evolve,\x20etc.\n\n\x0c\n\x05\x05\x01\x02\x02\x01\x12\x03\x0f\
    \x02'\n\x0c\n\x05\x05\x01\x02\x02\x02\x12\x03\x0f*+\n\n\n\x02\x04\0\x12\
    \x04\x14\0\x1c\x01\n\n\n\x03\x04\0\x01\x12\x03\x14\x08\x1c\n\x0b\n\x04\
    \x04\0\x02\0\x12\x03\x15\x02\x20\n\x0c\n\x05\x04\0\x02\0\x06\x12\x03\x15\
    \x02\x16\n\x0c\n\x05\x04\0\x02\0\x01\x12\x03\x15\x17\x1b\n\x0c\n\x05\x04\
    \0\x02\0\x03\x12\x03\x15\x1e\x1f\n\x0b\n\x04\x04\0\x02\x01\x12\x03\x17\
    \x02*\n\x0c\n\x05\x04\0\x02\x01\x06\x12\x03\x17\x02\x1b\n\x0c\n\x05\x04\
    \0\x02\x01\x01\x12\x03\x17\x1c%\n\x0c\n\x05\x04\0\x02\x01\x03\x12\x03\
    \x17()\n\x0c\n\x04\x04\0\x08\0\x12\x04\x19\x02\x1b\x03\n\x0c\n\x05\x04\0\
    \x08\0\x01\x12\x03\x19\x08\x0f\n\x0b\n\x04\x04\0\x02\x02\x12\x03\x1a\x04\
    1\n\x0c\n\x05\x04\0\x02\x02\x06\x12\x03\x1a\x04\x1e\n\x0c\n\x05\x04\0\
    \x02\x02\x01\x12\x03\x1a\x1f*\n\x0c\n\x05\x04\0\x02\x02\x03\x12\x03\x1a-\
    0\n\n\n\x02\x04\x01\x12\x04\x1e\0\x20\x01\n\n\n\x03\x04\x01\x01\x12\x03\
    \x1e\x08\"\n\x0b\n\x04\x04\x01\x02\0\x12\x03\x1f\x02\x18\n\x0c\n\x05\x04\
    \x01\x02\0\x05\x12\x03\x1f\x02\x07\n\x0c\n\x05\x04\x01\x02\0\x01\x12\x03\
    \x1f\x08\x13\n\x0c\n\x05\x04\x01\x02\0\x03\x12\x03\x1f\x16\x17b\x06proto\
    3\
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
            let mut deps = ::std::vec::Vec::with_capacity(0);
            let mut messages = ::std::vec::Vec::with_capacity(2);
            messages.push(SchemaValidationStep::generated_message_descriptor_data());
            messages.push(SchemaValidationJSONSchema::generated_message_descriptor_data());
            let mut enums = ::std::vec::Vec::with_capacity(2);
            enums.push(SchemaValidationType::generated_enum_descriptor_data());
            enums.push(SchemaValidationCondition::generated_enum_descriptor_data());
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
