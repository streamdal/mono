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

//! Generated file from `bus.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_3_2_0;

///  Type used by `snitch-server` for sending messages on its local bus.
#[derive(PartialEq,Clone,Default,Debug)]
// @@protoc_insertion_point(message:protos.BusEvent)
pub struct BusEvent {
    // message fields
    // @@protoc_insertion_point(field:protos.BusEvent.source)
    pub source: ::std::string::String,
    // @@protoc_insertion_point(field:protos.BusEvent._metadata)
    pub _metadata: ::std::collections::HashMap<::std::string::String, ::std::string::String>,
    // message oneof groups
    pub event: ::std::option::Option<bus_event::Event>,
    // special fields
    // @@protoc_insertion_point(special_field:protos.BusEvent.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a BusEvent {
    fn default() -> &'a BusEvent {
        <BusEvent as ::protobuf::Message>::default_instance()
    }
}

impl BusEvent {
    pub fn new() -> BusEvent {
        ::std::default::Default::default()
    }

    // .protos.Command command = 100;

    pub fn command(&self) -> &super::command::Command {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::Command(ref v)) => v,
            _ => <super::command::Command as ::protobuf::Message>::default_instance(),
        }
    }

    pub fn clear_command(&mut self) {
        self.event = ::std::option::Option::None;
    }

    pub fn has_command(&self) -> bool {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::Command(..)) => true,
            _ => false,
        }
    }

    // Param is passed by value, moved
    pub fn set_command(&mut self, v: super::command::Command) {
        self.event = ::std::option::Option::Some(bus_event::Event::Command(v))
    }

    // Mutable pointer to the field.
    pub fn mut_command(&mut self) -> &mut super::command::Command {
        if let ::std::option::Option::Some(bus_event::Event::Command(_)) = self.event {
        } else {
            self.event = ::std::option::Option::Some(bus_event::Event::Command(super::command::Command::new()));
        }
        match self.event {
            ::std::option::Option::Some(bus_event::Event::Command(ref mut v)) => v,
            _ => panic!(),
        }
    }

    // Take field
    pub fn take_command(&mut self) -> super::command::Command {
        if self.has_command() {
            match self.event.take() {
                ::std::option::Option::Some(bus_event::Event::Command(v)) => v,
                _ => panic!(),
            }
        } else {
            super::command::Command::new()
        }
    }

    // .protos.RegisterRequest register_request = 101;

    pub fn register_request(&self) -> &super::internal::RegisterRequest {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::RegisterRequest(ref v)) => v,
            _ => <super::internal::RegisterRequest as ::protobuf::Message>::default_instance(),
        }
    }

    pub fn clear_register_request(&mut self) {
        self.event = ::std::option::Option::None;
    }

    pub fn has_register_request(&self) -> bool {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::RegisterRequest(..)) => true,
            _ => false,
        }
    }

    // Param is passed by value, moved
    pub fn set_register_request(&mut self, v: super::internal::RegisterRequest) {
        self.event = ::std::option::Option::Some(bus_event::Event::RegisterRequest(v))
    }

    // Mutable pointer to the field.
    pub fn mut_register_request(&mut self) -> &mut super::internal::RegisterRequest {
        if let ::std::option::Option::Some(bus_event::Event::RegisterRequest(_)) = self.event {
        } else {
            self.event = ::std::option::Option::Some(bus_event::Event::RegisterRequest(super::internal::RegisterRequest::new()));
        }
        match self.event {
            ::std::option::Option::Some(bus_event::Event::RegisterRequest(ref mut v)) => v,
            _ => panic!(),
        }
    }

    // Take field
    pub fn take_register_request(&mut self) -> super::internal::RegisterRequest {
        if self.has_register_request() {
            match self.event.take() {
                ::std::option::Option::Some(bus_event::Event::RegisterRequest(v)) => v,
                _ => panic!(),
            }
        } else {
            super::internal::RegisterRequest::new()
        }
    }

    // .protos.DeregisterRequest deregister_request = 102;

    pub fn deregister_request(&self) -> &super::internal::DeregisterRequest {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::DeregisterRequest(ref v)) => v,
            _ => <super::internal::DeregisterRequest as ::protobuf::Message>::default_instance(),
        }
    }

    pub fn clear_deregister_request(&mut self) {
        self.event = ::std::option::Option::None;
    }

    pub fn has_deregister_request(&self) -> bool {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::DeregisterRequest(..)) => true,
            _ => false,
        }
    }

    // Param is passed by value, moved
    pub fn set_deregister_request(&mut self, v: super::internal::DeregisterRequest) {
        self.event = ::std::option::Option::Some(bus_event::Event::DeregisterRequest(v))
    }

    // Mutable pointer to the field.
    pub fn mut_deregister_request(&mut self) -> &mut super::internal::DeregisterRequest {
        if let ::std::option::Option::Some(bus_event::Event::DeregisterRequest(_)) = self.event {
        } else {
            self.event = ::std::option::Option::Some(bus_event::Event::DeregisterRequest(super::internal::DeregisterRequest::new()));
        }
        match self.event {
            ::std::option::Option::Some(bus_event::Event::DeregisterRequest(ref mut v)) => v,
            _ => panic!(),
        }
    }

    // Take field
    pub fn take_deregister_request(&mut self) -> super::internal::DeregisterRequest {
        if self.has_deregister_request() {
            match self.event.take() {
                ::std::option::Option::Some(bus_event::Event::DeregisterRequest(v)) => v,
                _ => panic!(),
            }
        } else {
            super::internal::DeregisterRequest::new()
        }
    }

    // .protos.HeartbeatRequest heartbeat_request = 103;

    pub fn heartbeat_request(&self) -> &super::internal::HeartbeatRequest {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(ref v)) => v,
            _ => <super::internal::HeartbeatRequest as ::protobuf::Message>::default_instance(),
        }
    }

    pub fn clear_heartbeat_request(&mut self) {
        self.event = ::std::option::Option::None;
    }

    pub fn has_heartbeat_request(&self) -> bool {
        match self.event {
            ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(..)) => true,
            _ => false,
        }
    }

    // Param is passed by value, moved
    pub fn set_heartbeat_request(&mut self, v: super::internal::HeartbeatRequest) {
        self.event = ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(v))
    }

    // Mutable pointer to the field.
    pub fn mut_heartbeat_request(&mut self) -> &mut super::internal::HeartbeatRequest {
        if let ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(_)) = self.event {
        } else {
            self.event = ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(super::internal::HeartbeatRequest::new()));
        }
        match self.event {
            ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(ref mut v)) => v,
            _ => panic!(),
        }
    }

    // Take field
    pub fn take_heartbeat_request(&mut self) -> super::internal::HeartbeatRequest {
        if self.has_heartbeat_request() {
            match self.event.take() {
                ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(v)) => v,
                _ => panic!(),
            }
        } else {
            super::internal::HeartbeatRequest::new()
        }
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(6);
        let mut oneofs = ::std::vec::Vec::with_capacity(1);
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "source",
            |m: &BusEvent| { &m.source },
            |m: &mut BusEvent| { &mut m.source },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_oneof_message_has_get_mut_set_accessor::<_, super::command::Command>(
            "command",
            BusEvent::has_command,
            BusEvent::command,
            BusEvent::mut_command,
            BusEvent::set_command,
        ));
        fields.push(::protobuf::reflect::rt::v2::make_oneof_message_has_get_mut_set_accessor::<_, super::internal::RegisterRequest>(
            "register_request",
            BusEvent::has_register_request,
            BusEvent::register_request,
            BusEvent::mut_register_request,
            BusEvent::set_register_request,
        ));
        fields.push(::protobuf::reflect::rt::v2::make_oneof_message_has_get_mut_set_accessor::<_, super::internal::DeregisterRequest>(
            "deregister_request",
            BusEvent::has_deregister_request,
            BusEvent::deregister_request,
            BusEvent::mut_deregister_request,
            BusEvent::set_deregister_request,
        ));
        fields.push(::protobuf::reflect::rt::v2::make_oneof_message_has_get_mut_set_accessor::<_, super::internal::HeartbeatRequest>(
            "heartbeat_request",
            BusEvent::has_heartbeat_request,
            BusEvent::heartbeat_request,
            BusEvent::mut_heartbeat_request,
            BusEvent::set_heartbeat_request,
        ));
        fields.push(::protobuf::reflect::rt::v2::make_map_simpler_accessor::<_, _, _>(
            "_metadata",
            |m: &BusEvent| { &m._metadata },
            |m: &mut BusEvent| { &mut m._metadata },
        ));
        oneofs.push(bus_event::Event::generated_oneof_descriptor_data());
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<BusEvent>(
            "BusEvent",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for BusEvent {
    const NAME: &'static str = "BusEvent";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                10 => {
                    self.source = is.read_string()?;
                },
                802 => {
                    self.event = ::std::option::Option::Some(bus_event::Event::Command(is.read_message()?));
                },
                810 => {
                    self.event = ::std::option::Option::Some(bus_event::Event::RegisterRequest(is.read_message()?));
                },
                818 => {
                    self.event = ::std::option::Option::Some(bus_event::Event::DeregisterRequest(is.read_message()?));
                },
                826 => {
                    self.event = ::std::option::Option::Some(bus_event::Event::HeartbeatRequest(is.read_message()?));
                },
                8002 => {
                    let len = is.read_raw_varint32()?;
                    let old_limit = is.push_limit(len as u64)?;
                    let mut key = ::std::default::Default::default();
                    let mut value = ::std::default::Default::default();
                    while let Some(tag) = is.read_raw_tag_or_eof()? {
                        match tag {
                            10 => key = is.read_string()?,
                            18 => value = is.read_string()?,
                            _ => ::protobuf::rt::skip_field_for_tag(tag, is)?,
                        };
                    }
                    is.pop_limit(old_limit);
                    self._metadata.insert(key, value);
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
        if !self.source.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.source);
        }
        for (k, v) in &self._metadata {
            let mut entry_size = 0;
            entry_size += ::protobuf::rt::string_size(1, &k);
            entry_size += ::protobuf::rt::string_size(2, &v);
            my_size += 2 + ::protobuf::rt::compute_raw_varint64_size(entry_size) + entry_size
        };
        if let ::std::option::Option::Some(ref v) = self.event {
            match v {
                &bus_event::Event::Command(ref v) => {
                    let len = v.compute_size();
                    my_size += 2 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
                },
                &bus_event::Event::RegisterRequest(ref v) => {
                    let len = v.compute_size();
                    my_size += 2 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
                },
                &bus_event::Event::DeregisterRequest(ref v) => {
                    let len = v.compute_size();
                    my_size += 2 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
                },
                &bus_event::Event::HeartbeatRequest(ref v) => {
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
        if !self.source.is_empty() {
            os.write_string(1, &self.source)?;
        }
        for (k, v) in &self._metadata {
            let mut entry_size = 0;
            entry_size += ::protobuf::rt::string_size(1, &k);
            entry_size += ::protobuf::rt::string_size(2, &v);
            os.write_raw_varint32(8002)?; // Tag.
            os.write_raw_varint32(entry_size as u32)?;
            os.write_string(1, &k)?;
            os.write_string(2, &v)?;
        };
        if let ::std::option::Option::Some(ref v) = self.event {
            match v {
                &bus_event::Event::Command(ref v) => {
                    ::protobuf::rt::write_message_field_with_cached_size(100, v, os)?;
                },
                &bus_event::Event::RegisterRequest(ref v) => {
                    ::protobuf::rt::write_message_field_with_cached_size(101, v, os)?;
                },
                &bus_event::Event::DeregisterRequest(ref v) => {
                    ::protobuf::rt::write_message_field_with_cached_size(102, v, os)?;
                },
                &bus_event::Event::HeartbeatRequest(ref v) => {
                    ::protobuf::rt::write_message_field_with_cached_size(103, v, os)?;
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

    fn new() -> BusEvent {
        BusEvent::new()
    }

    fn clear(&mut self) {
        self.source.clear();
        self.event = ::std::option::Option::None;
        self.event = ::std::option::Option::None;
        self.event = ::std::option::Option::None;
        self.event = ::std::option::Option::None;
        self._metadata.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static BusEvent {
        static instance: ::protobuf::rt::Lazy<BusEvent> = ::protobuf::rt::Lazy::new();
        instance.get(BusEvent::new)
    }
}

impl ::protobuf::MessageFull for BusEvent {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("BusEvent").unwrap()).clone()
    }
}

impl ::std::fmt::Display for BusEvent {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for BusEvent {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

/// Nested message and enums of message `BusEvent`
pub mod bus_event {

    #[derive(Clone,PartialEq,Debug)]
    #[non_exhaustive]
    // @@protoc_insertion_point(oneof:protos.BusEvent.event)
    pub enum Event {
        // @@protoc_insertion_point(oneof_field:protos.BusEvent.command)
        Command(super::super::command::Command),
        // @@protoc_insertion_point(oneof_field:protos.BusEvent.register_request)
        RegisterRequest(super::super::internal::RegisterRequest),
        // @@protoc_insertion_point(oneof_field:protos.BusEvent.deregister_request)
        DeregisterRequest(super::super::internal::DeregisterRequest),
        // @@protoc_insertion_point(oneof_field:protos.BusEvent.heartbeat_request)
        HeartbeatRequest(super::super::internal::HeartbeatRequest),
    }

    impl ::protobuf::Oneof for Event {
    }

    impl ::protobuf::OneofFull for Event {
        fn descriptor() -> ::protobuf::reflect::OneofDescriptor {
            static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::OneofDescriptor> = ::protobuf::rt::Lazy::new();
            descriptor.get(|| <super::BusEvent as ::protobuf::MessageFull>::descriptor().oneof_by_name("event").unwrap()).clone()
        }
    }

    impl Event {
        pub(in super) fn generated_oneof_descriptor_data() -> ::protobuf::reflect::GeneratedOneofDescriptorData {
            ::protobuf::reflect::GeneratedOneofDescriptorData::new::<Event>("event")
        }
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n\tbus.proto\x12\x06protos\x1a\rcommand.proto\x1a\x0einternal.proto\"\
    \xae\x03\n\x08BusEvent\x12\x16\n\x06source\x18\x01\x20\x01(\tR\x06source\
    \x12+\n\x07command\x18d\x20\x01(\x0b2\x0f.protos.CommandH\0R\x07command\
    \x12D\n\x10register_request\x18e\x20\x01(\x0b2\x17.protos.RegisterReques\
    tH\0R\x0fregisterRequest\x12J\n\x12deregister_request\x18f\x20\x01(\x0b2\
    \x19.protos.DeregisterRequestH\0R\x11deregisterRequest\x12G\n\x11heartbe\
    at_request\x18g\x20\x01(\x0b2\x18.protos.HeartbeatRequestH\0R\x10heartbe\
    atRequest\x12<\n\t_metadata\x18\xe8\x07\x20\x03(\x0b2\x1e.protos.BusEven\
    t.MetadataEntryR\x08Metadata\x1a;\n\rMetadataEntry\x12\x10\n\x03key\x18\
    \x01\x20\x01(\tR\x03key\x12\x14\n\x05value\x18\x02\x20\x01(\tR\x05value:\
    \x028\x01B\x07\n\x05eventB4Z2github.com/streamdal/snitch-protos/build/go\
    /protosJ\xd6\x08\n\x06\x12\x04\0\0\x1f\x01\n\x08\n\x01\x0c\x12\x03\0\0\
    \x12\n\x08\n\x01\x02\x12\x03\x02\0\x0f\n\t\n\x02\x03\0\x12\x03\x04\0\x17\
    \n\t\n\x02\x03\x01\x12\x03\x05\0\x18\n\x08\n\x01\x08\x12\x03\x07\0I\n\t\
    \n\x02\x08\x0b\x12\x03\x07\0I\nQ\n\x02\x04\0\x12\x04\n\0\x1f\x01\x1aE\
    \x20Type\x20used\x20by\x20`snitch-server`\x20for\x20sending\x20messages\
    \x20on\x20its\x20local\x20bus.\n\n\n\n\x03\x04\0\x01\x12\x03\n\x08\x10\n\
    \x0b\n\x04\x04\0\x02\0\x12\x03\x0b\x02\x14\n\x0c\n\x05\x04\0\x02\0\x05\
    \x12\x03\x0b\x02\x08\n\x0c\n\x05\x04\0\x02\0\x01\x12\x03\x0b\t\x0f\n\x0c\
    \n\x05\x04\0\x02\0\x03\x12\x03\x0b\x12\x13\n\x0c\n\x04\x04\0\x08\0\x12\
    \x04\r\x02\x12\x03\n\x0c\n\x05\x04\0\x08\0\x01\x12\x03\r\x08\r\n\x0b\n\
    \x04\x04\0\x02\x01\x12\x03\x0e\x04!\n\x0c\n\x05\x04\0\x02\x01\x06\x12\
    \x03\x0e\x04\x12\n\x0c\n\x05\x04\0\x02\x01\x01\x12\x03\x0e\x13\x1a\n\x0c\
    \n\x05\x04\0\x02\x01\x03\x12\x03\x0e\x1d\x20\n\x0b\n\x04\x04\0\x02\x02\
    \x12\x03\x0f\x042\n\x0c\n\x05\x04\0\x02\x02\x06\x12\x03\x0f\x04\x1a\n\
    \x0c\n\x05\x04\0\x02\x02\x01\x12\x03\x0f\x1b+\n\x0c\n\x05\x04\0\x02\x02\
    \x03\x12\x03\x0f.1\n\x0b\n\x04\x04\0\x02\x03\x12\x03\x10\x046\n\x0c\n\
    \x05\x04\0\x02\x03\x06\x12\x03\x10\x04\x1c\n\x0c\n\x05\x04\0\x02\x03\x01\
    \x12\x03\x10\x1d/\n\x0c\n\x05\x04\0\x02\x03\x03\x12\x03\x1025\n\x0b\n\
    \x04\x04\0\x02\x04\x12\x03\x11\x044\n\x0c\n\x05\x04\0\x02\x04\x06\x12\
    \x03\x11\x04\x1b\n\x0c\n\x05\x04\0\x02\x04\x01\x12\x03\x11\x1c-\n\x0c\n\
    \x05\x04\0\x02\x04\x03\x12\x03\x1103\n\xd4\x04\n\x04\x04\0\x02\x05\x12\
    \x03\x1e\x02(\x1a\x8f\x04\x20All\x20gRPC\x20metadata\x20is\x20stored\x20\
    in\x20ctx;\x20when\x20request\x20goes\x20outside\x20of\x20gRPC\n\x20boun\
    ds,\x20we\x20will\x20translate\x20ctx\x20metadata\x20into\x20this\x20fie\
    ld.\n\n\x20Example:\n\x201.\x20Request\x20comes\x20into\x20snitch-server\
    \x20via\x20external\x20gRPC\x20to\x20set\x20new\x20pipeline\n\x202.\x20s\
    nitch-server\x20has\x20to\x20send\x20SetPipeline\x20cmd\x20to\x20SDK\x20\
    via\x20gRPC\x20-\x20it\x20passes\n\x20\x20\x20\x20on\x20original\x20meta\
    data\x20in\x20request.\n\x203.\x20snitch-server\x20has\x20to\x20broadcas\
    t\x20SetPipeline\x20cmd\x20to\x20other\x20services\x20via\x20bus\n\x204.\
    \x20Since\x20this\x20is\x20not\x20a\x20gRPC\x20call,\x20snitch-server\
    \x20translates\x20ctx\x20metadata\x20to\n\x20\x20\x20\x20this\x20field\
    \x20and\x20includes\x20it\x20in\x20the\x20bus\x20event.\n\"5\x20protolin\
    t:disable:this\x20FIELD_NAMES_LOWER_SNAKE_CASE\n\n\x0c\n\x05\x04\0\x02\
    \x05\x06\x12\x03\x1e\x02\x16\n\x0c\n\x05\x04\0\x02\x05\x01\x12\x03\x1e\
    \x17\x20\n\x0c\n\x05\x04\0\x02\x05\x03\x12\x03\x1e#'b\x06proto3\
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
            deps.push(super::command::file_descriptor().clone());
            deps.push(super::internal::file_descriptor().clone());
            let mut messages = ::std::vec::Vec::with_capacity(1);
            messages.push(BusEvent::generated_message_descriptor_data());
            let mut enums = ::std::vec::Vec::with_capacity(0);
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
