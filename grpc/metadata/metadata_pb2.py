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
  package='metadata',
  syntax='proto3',
  serialized_options=b'Z0github.com/alexandr-io/backend/grpc/grpcmetadata',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0emetadata.proto\x12\x08metadata\"1\n\x0fMetadataRequest\x12\r\n\x05Title\x18\x01 \x01(\t\x12\x0f\n\x07\x41uthors\x18\x02 \x01(\t\"\xd3\x01\n\rMetadataReply\x12\r\n\x05Title\x18\x01 \x01(\t\x12\x0f\n\x07\x41uthors\x18\x02 \x01(\t\x12\x11\n\tPublisher\x18\x03 \x01(\t\x12\x15\n\rPublishedDate\x18\x04 \x01(\t\x12\x11\n\tPageCount\x18\x05 \x01(\t\x12\x12\n\nCategories\x18\x06 \x01(\t\x12\x16\n\x0eMaturityRating\x18\x07 \x01(\t\x12\x10\n\x08Language\x18\x08 \x01(\t\x12\x12\n\nImageLinks\x18\t \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\n \x01(\t2L\n\x08Metadata\x12@\n\x08Metadata\x12\x19.metadata.MetadataRequest\x1a\x17.metadata.MetadataReply\"\x00\x42\x32Z0github.com/alexandr-io/backend/grpc/grpcmetadatab\x06proto3'
)




_METADATAREQUEST = _descriptor.Descriptor(
  name='MetadataRequest',
  full_name='metadata.MetadataRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Title', full_name='metadata.MetadataRequest.Title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Authors', full_name='metadata.MetadataRequest.Authors', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=28,
  serialized_end=77,
)


_METADATAREPLY = _descriptor.Descriptor(
  name='MetadataReply',
  full_name='metadata.MetadataReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Title', full_name='metadata.MetadataReply.Title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Authors', full_name='metadata.MetadataReply.Authors', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Publisher', full_name='metadata.MetadataReply.Publisher', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='PublishedDate', full_name='metadata.MetadataReply.PublishedDate', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='PageCount', full_name='metadata.MetadataReply.PageCount', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Categories', full_name='metadata.MetadataReply.Categories', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaturityRating', full_name='metadata.MetadataReply.MaturityRating', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Language', full_name='metadata.MetadataReply.Language', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ImageLinks', full_name='metadata.MetadataReply.ImageLinks', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Description', full_name='metadata.MetadataReply.Description', index=9,
      number=10, type=9, cpp_type=9, label=1,
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
  serialized_start=80,
  serialized_end=291,
)

DESCRIPTOR.message_types_by_name['MetadataRequest'] = _METADATAREQUEST
DESCRIPTOR.message_types_by_name['MetadataReply'] = _METADATAREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MetadataRequest = _reflection.GeneratedProtocolMessageType('MetadataRequest', (_message.Message,), {
  'DESCRIPTOR' : _METADATAREQUEST,
  '__module__' : 'metadata_pb2'
  # @@protoc_insertion_point(class_scope:metadata.MetadataRequest)
  })
_sym_db.RegisterMessage(MetadataRequest)

MetadataReply = _reflection.GeneratedProtocolMessageType('MetadataReply', (_message.Message,), {
  'DESCRIPTOR' : _METADATAREPLY,
  '__module__' : 'metadata_pb2'
  # @@protoc_insertion_point(class_scope:metadata.MetadataReply)
  })
_sym_db.RegisterMessage(MetadataReply)


DESCRIPTOR._options = None

_METADATA = _descriptor.ServiceDescriptor(
  name='Metadata',
  full_name='metadata.Metadata',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=293,
  serialized_end=369,
  methods=[
  _descriptor.MethodDescriptor(
    name='Metadata',
    full_name='metadata.Metadata.Metadata',
    index=0,
    containing_service=None,
    input_type=_METADATAREQUEST,
    output_type=_METADATAREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_METADATA)

DESCRIPTOR.services_by_name['Metadata'] = _METADATA

# @@protoc_insertion_point(module_scope)
