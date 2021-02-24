# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: metadata.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='metadata.proto',
  package='auth',
  syntax='proto3',
  serialized_options=b'Z,github.com/alexandr-io/backend/grpc/grpcauth',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0emetadata.proto\x12\x04\x61uth\"\x1a\n\x0b\x41uthRequest\x12\x0b\n\x03JWT\x18\x01 \x01(\t\"8\n\tAuthReply\x12\n\n\x02ID\x18\x01 \x01(\t\x12\x10\n\x08username\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t24\n\x04\x41uth\x12,\n\x04\x41uth\x12\x11.auth.AuthRequest\x1a\x0f.auth.AuthReply\"\x00\x42.Z,github.com/alexandr-io/backend/grpc/grpcauthb\x06proto3'
)




_AUTHREQUEST = _descriptor.Descriptor(
  name='AuthRequest',
  full_name='auth.AuthRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='JWT', full_name='auth.AuthRequest.JWT', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=50,
)


_AUTHREPLY = _descriptor.Descriptor(
  name='AuthReply',
  full_name='auth.AuthReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='auth.AuthReply.ID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='username', full_name='auth.AuthReply.username', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='email', full_name='auth.AuthReply.email', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=52,
  serialized_end=108,
)

DESCRIPTOR.message_types_by_name['AuthRequest'] = _AUTHREQUEST
DESCRIPTOR.message_types_by_name['AuthReply'] = _AUTHREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AuthRequest = _reflection.GeneratedProtocolMessageType('AuthRequest', (_message.Message,), {
  'DESCRIPTOR' : _AUTHREQUEST,
  '__module__' : 'metadata_pb2'
  # @@protoc_insertion_point(class_scope:auth.AuthRequest)
  })
_sym_db.RegisterMessage(AuthRequest)

AuthReply = _reflection.GeneratedProtocolMessageType('AuthReply', (_message.Message,), {
  'DESCRIPTOR' : _AUTHREPLY,
  '__module__' : 'metadata_pb2'
  # @@protoc_insertion_point(class_scope:auth.AuthReply)
  })
_sym_db.RegisterMessage(AuthReply)


DESCRIPTOR._options = None

_AUTH = _descriptor.ServiceDescriptor(
  name='Auth',
  full_name='auth.Auth',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=110,
  serialized_end=162,
  methods=[
  _descriptor.MethodDescriptor(
    name='Auth',
    full_name='auth.Auth.Auth',
    index=0,
    containing_service=None,
    input_type=_AUTHREQUEST,
    output_type=_AUTHREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_AUTH)

DESCRIPTOR.services_by_name['Auth'] = _AUTH

# @@protoc_insertion_point(module_scope)
