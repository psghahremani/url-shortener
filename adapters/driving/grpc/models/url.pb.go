// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: adapters/driving/grpc/models/url.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ShortenedUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handle      string `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	CreatedAt   int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt   int64  `protobuf:"varint,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *ShortenedUrl) Reset() {
	*x = ShortenedUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapters_driving_grpc_models_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenedUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenedUrl) ProtoMessage() {}

func (x *ShortenedUrl) ProtoReflect() protoreflect.Message {
	mi := &file_adapters_driving_grpc_models_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenedUrl.ProtoReflect.Descriptor instead.
func (*ShortenedUrl) Descriptor() ([]byte, []int) {
	return file_adapters_driving_grpc_models_url_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenedUrl) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *ShortenedUrl) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *ShortenedUrl) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ShortenedUrl) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type ShortenUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url       string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ExpiresAt int64  `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *ShortenUrlRequest) Reset() {
	*x = ShortenUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapters_driving_grpc_models_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenUrlRequest) ProtoMessage() {}

func (x *ShortenUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adapters_driving_grpc_models_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenUrlRequest.ProtoReflect.Descriptor instead.
func (*ShortenUrlRequest) Descriptor() ([]byte, []int) {
	return file_adapters_driving_grpc_models_url_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenUrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortenUrlRequest) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type LookupUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handle string `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"`
}

func (x *LookupUrlRequest) Reset() {
	*x = LookupUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapters_driving_grpc_models_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupUrlRequest) ProtoMessage() {}

func (x *LookupUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adapters_driving_grpc_models_url_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupUrlRequest.ProtoReflect.Descriptor instead.
func (*LookupUrlRequest) Descriptor() ([]byte, []int) {
	return file_adapters_driving_grpc_models_url_proto_rawDescGZIP(), []int{2}
}

func (x *LookupUrlRequest) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

var File_adapters_driving_grpc_models_url_proto protoreflect.FileDescriptor

var file_adapters_driving_grpc_models_url_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75,
	0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x87, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22,
	0x2a, 0x0a, 0x10, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x32, 0x8d, 0x01, 0x0a, 0x13,
	0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x12, 0x39, 0x0a, 0x09, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x73, 0x67, 0x68, 0x61, 0x68,
	0x72, 0x65, 0x6d, 0x61, 0x6e, 0x69, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_adapters_driving_grpc_models_url_proto_rawDescOnce sync.Once
	file_adapters_driving_grpc_models_url_proto_rawDescData = file_adapters_driving_grpc_models_url_proto_rawDesc
)

func file_adapters_driving_grpc_models_url_proto_rawDescGZIP() []byte {
	file_adapters_driving_grpc_models_url_proto_rawDescOnce.Do(func() {
		file_adapters_driving_grpc_models_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_adapters_driving_grpc_models_url_proto_rawDescData)
	})
	return file_adapters_driving_grpc_models_url_proto_rawDescData
}

var file_adapters_driving_grpc_models_url_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_adapters_driving_grpc_models_url_proto_goTypes = []interface{}{
	(*ShortenedUrl)(nil),      // 0: proto.ShortenedUrl
	(*ShortenUrlRequest)(nil), // 1: proto.ShortenUrlRequest
	(*LookupUrlRequest)(nil),  // 2: proto.LookupUrlRequest
}
var file_adapters_driving_grpc_models_url_proto_depIdxs = []int32{
	1, // 0: proto.UrlShortenerService.ShortenUrl:input_type -> proto.ShortenUrlRequest
	2, // 1: proto.UrlShortenerService.LookupUrl:input_type -> proto.LookupUrlRequest
	0, // 2: proto.UrlShortenerService.ShortenUrl:output_type -> proto.ShortenedUrl
	0, // 3: proto.UrlShortenerService.LookupUrl:output_type -> proto.ShortenedUrl
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adapters_driving_grpc_models_url_proto_init() }
func file_adapters_driving_grpc_models_url_proto_init() {
	if File_adapters_driving_grpc_models_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adapters_driving_grpc_models_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenedUrl); i {
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
		file_adapters_driving_grpc_models_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenUrlRequest); i {
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
		file_adapters_driving_grpc_models_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupUrlRequest); i {
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
			RawDescriptor: file_adapters_driving_grpc_models_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adapters_driving_grpc_models_url_proto_goTypes,
		DependencyIndexes: file_adapters_driving_grpc_models_url_proto_depIdxs,
		MessageInfos:      file_adapters_driving_grpc_models_url_proto_msgTypes,
	}.Build()
	File_adapters_driving_grpc_models_url_proto = out.File
	file_adapters_driving_grpc_models_url_proto_rawDesc = nil
	file_adapters_driving_grpc_models_url_proto_goTypes = nil
	file_adapters_driving_grpc_models_url_proto_depIdxs = nil
}
