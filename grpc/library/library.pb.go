// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protocol/library.proto

package grpclibrary

import (
	reflect "reflect"
	sync "sync"

	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateLibraryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *CreateLibraryRequest) Reset() {
	*x = CreateLibraryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLibraryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLibraryRequest) ProtoMessage() {}

func (x *CreateLibraryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_library_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLibraryRequest.ProtoReflect.Descriptor instead.
func (*CreateLibraryRequest) Descriptor() ([]byte, []int) {
	return file_protocol_library_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLibraryRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UploadAuthorizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	LibraryID string `protobuf:"bytes,2,opt,name=LibraryID,proto3" json:"LibraryID,omitempty"`
}

func (x *UploadAuthorizationRequest) Reset() {
	*x = UploadAuthorizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_library_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAuthorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAuthorizationRequest) ProtoMessage() {}

func (x *UploadAuthorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_library_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAuthorizationRequest.ProtoReflect.Descriptor instead.
func (*UploadAuthorizationRequest) Descriptor() ([]byte, []int) {
	return file_protocol_library_proto_rawDescGZIP(), []int{1}
}

func (x *UploadAuthorizationRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UploadAuthorizationRequest) GetLibraryID() string {
	if x != nil {
		return x.LibraryID
	}
	return ""
}

type UploadAuthorizationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool `protobuf:"varint,1,opt,name=Authorized,proto3" json:"Authorized,omitempty"`
}

func (x *UploadAuthorizationReply) Reset() {
	*x = UploadAuthorizationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_library_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAuthorizationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAuthorizationReply) ProtoMessage() {}

func (x *UploadAuthorizationReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_library_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAuthorizationReply.ProtoReflect.Descriptor instead.
func (*UploadAuthorizationReply) Descriptor() ([]byte, []int) {
	return file_protocol_library_proto_rawDescGZIP(), []int{2}
}

func (x *UploadAuthorizationReply) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

type BookUploadedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *BookUploadedRequest) Reset() {
	*x = BookUploadedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_library_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookUploadedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookUploadedRequest) ProtoMessage() {}

func (x *BookUploadedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_library_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookUploadedRequest.ProtoReflect.Descriptor instead.
func (*BookUploadedRequest) Descriptor() ([]byte, []int) {
	return file_protocol_library_proto_rawDescGZIP(), []int{3}
}

func (x *BookUploadedRequest) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

func (x *BookUploadedRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CoverUploadedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID   string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
	CoverURL string `protobuf:"bytes,2,opt,name=CoverURL,proto3" json:"CoverURL,omitempty"`
}

func (x *CoverUploadedRequest) Reset() {
	*x = CoverUploadedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_library_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoverUploadedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoverUploadedRequest) ProtoMessage() {}

func (x *CoverUploadedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_library_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoverUploadedRequest.ProtoReflect.Descriptor instead.
func (*CoverUploadedRequest) Descriptor() ([]byte, []int) {
	return file_protocol_library_proto_rawDescGZIP(), []int{4}
}

func (x *CoverUploadedRequest) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

func (x *CoverUploadedRequest) GetCoverURL() string {
	if x != nil {
		return x.CoverURL
	}
	return ""
}

var File_protocol_library_proto protoreflect.FileDescriptor

var file_protocol_library_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x52,
	0x0a, 0x1a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x49, 0x44, 0x22, 0x3a, 0x0a, 0x18, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x41,
	0x0a, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x4a, 0x0a, 0x14, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x32, 0xc6, 0x02,
	0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x1d, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protocol_library_proto_rawDescOnce sync.Once
	file_protocol_library_proto_rawDescData = file_protocol_library_proto_rawDesc
)

func file_protocol_library_proto_rawDescGZIP() []byte {
	file_protocol_library_proto_rawDescOnce.Do(func() {
		file_protocol_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_library_proto_rawDescData)
	})
	return file_protocol_library_proto_rawDescData
}

var file_protocol_library_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protocol_library_proto_goTypes = []interface{}{
	(*CreateLibraryRequest)(nil),       // 0: library.CreateLibraryRequest
	(*UploadAuthorizationRequest)(nil), // 1: library.UploadAuthorizationRequest
	(*UploadAuthorizationReply)(nil),   // 2: library.UploadAuthorizationReply
	(*BookUploadedRequest)(nil),        // 3: library.BookUploadedRequest
	(*CoverUploadedRequest)(nil),       // 4: library.CoverUploadedRequest
	(*empty.Empty)(nil),                // 5: google.protobuf.Empty
}
var file_protocol_library_proto_depIdxs = []int32{
	0, // 0: library.Library.CreateLibrary:input_type -> library.CreateLibraryRequest
	1, // 1: library.Library.UploadAuthorization:input_type -> library.UploadAuthorizationRequest
	3, // 2: library.Library.BookUploaded:input_type -> library.BookUploadedRequest
	4, // 3: library.Library.CoverUploaded:input_type -> library.CoverUploadedRequest
	5, // 4: library.Library.CreateLibrary:output_type -> google.protobuf.Empty
	2, // 5: library.Library.UploadAuthorization:output_type -> library.UploadAuthorizationReply
	5, // 6: library.Library.BookUploaded:output_type -> google.protobuf.Empty
	5, // 7: library.Library.CoverUploaded:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_library_proto_init() }
func file_protocol_library_proto_init() {
	if File_protocol_library_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLibraryRequest); i {
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
		file_protocol_library_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAuthorizationRequest); i {
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
		file_protocol_library_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAuthorizationReply); i {
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
		file_protocol_library_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookUploadedRequest); i {
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
		file_protocol_library_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoverUploadedRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocol_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_library_proto_goTypes,
		DependencyIndexes: file_protocol_library_proto_depIdxs,
		MessageInfos:      file_protocol_library_proto_msgTypes,
	}.Build()
	File_protocol_library_proto = out.File
	file_protocol_library_proto_rawDesc = nil
	file_protocol_library_proto_goTypes = nil
	file_protocol_library_proto_depIdxs = nil
}
