// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: protocol/metadata.proto

package grpcmetadata

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *MetadataRequest) Reset() {
	*x = MetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataRequest) ProtoMessage() {}

func (x *MetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataRequest.ProtoReflect.Descriptor instead.
func (*MetadataRequest) Descriptor() ([]byte, []int) {
	return file_protocol_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MetadataRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type MetadataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title          string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Author         string `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	Publisher      string `protobuf:"bytes,3,opt,name=Publisher,proto3" json:"Publisher,omitempty"`
	PublishedDate  string `protobuf:"bytes,4,opt,name=PublishedDate,proto3" json:"PublishedDate,omitempty"`
	PageCount      string `protobuf:"bytes,5,opt,name=PageCount,proto3" json:"PageCount,omitempty"`
	Categories     string `protobuf:"bytes,6,opt,name=Categories,proto3" json:"Categories,omitempty"`
	MaturityRating string `protobuf:"bytes,7,opt,name=MaturityRating,proto3" json:"MaturityRating,omitempty"`
	Language       string `protobuf:"bytes,8,opt,name=Language,proto3" json:"Language,omitempty"`
	ImageLinks     string `protobuf:"bytes,9,opt,name=ImageLinks,proto3" json:"ImageLinks,omitempty"`
	Descriptions   string `protobuf:"bytes,10,opt,name=Descriptions,proto3" json:"Descriptions,omitempty"`
}

func (x *MetadataReply) Reset() {
	*x = MetadataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataReply) ProtoMessage() {}

func (x *MetadataReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataReply.ProtoReflect.Descriptor instead.
func (*MetadataReply) Descriptor() ([]byte, []int) {
	return file_protocol_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MetadataReply) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *MetadataReply) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *MetadataReply) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *MetadataReply) GetPageCount() string {
	if x != nil {
		return x.PageCount
	}
	return ""
}

func (x *MetadataReply) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

func (x *MetadataReply) GetMaturityRating() string {
	if x != nil {
		return x.MaturityRating
	}
	return ""
}

func (x *MetadataReply) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *MetadataReply) GetImageLinks() string {
	if x != nil {
		return x.ImageLinks
	}
	return ""
}

func (x *MetadataReply) GetDescriptions() string {
	if x != nil {
		return x.Descriptions
	}
	return ""
}

var File_protocol_metadata_proto protoreflect.FileDescriptor

var file_protocol_metadata_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x22, 0xc7, 0x02, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x4d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x4c,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78, 0x61,
	0x6e, 0x64, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_metadata_proto_rawDescOnce sync.Once
	file_protocol_metadata_proto_rawDescData = file_protocol_metadata_proto_rawDesc
)

func file_protocol_metadata_proto_rawDescGZIP() []byte {
	file_protocol_metadata_proto_rawDescOnce.Do(func() {
		file_protocol_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_metadata_proto_rawDescData)
	})
	return file_protocol_metadata_proto_rawDescData
}

var file_protocol_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_metadata_proto_goTypes = []interface{}{
	(*MetadataRequest)(nil), // 0: metadata.MetadataRequest
	(*MetadataReply)(nil),   // 1: metadata.MetadataReply
}
var file_protocol_metadata_proto_depIdxs = []int32{
	0, // 0: metadata.Metadata.Metadata:input_type -> metadata.MetadataRequest
	1, // 1: metadata.Metadata.Metadata:output_type -> metadata.MetadataReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_metadata_proto_init() }
func file_protocol_metadata_proto_init() {
	if File_protocol_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataRequest); i {
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
		file_protocol_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataReply); i {
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
			RawDescriptor: file_protocol_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_metadata_proto_goTypes,
		DependencyIndexes: file_protocol_metadata_proto_depIdxs,
		MessageInfos:      file_protocol_metadata_proto_msgTypes,
	}.Build()
	File_protocol_metadata_proto = out.File
	file_protocol_metadata_proto_rawDesc = nil
	file_protocol_metadata_proto_goTypes = nil
	file_protocol_metadata_proto_depIdxs = nil
}
