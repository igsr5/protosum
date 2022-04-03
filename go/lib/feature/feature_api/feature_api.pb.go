// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: feature/feature_api.proto

package featureApiPb

import (
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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_feature_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_feature_feature_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_feature_feature_api_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_feature_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_feature_feature_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_feature_feature_api_proto_rawDescGZIP(), []int{1}
}

func (x *Feature) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_feature_feature_api_proto protoreflect.FileDescriptor

var file_feature_feature_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x70, 0x62, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x60, 0x0a, 0x0e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x21, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x3b,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x70, 0x69, 0x50, 0x62, 0xea, 0x02, 0x1f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x75, 0x6d, 0x3a, 0x3a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x3a, 0x3a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x70, 0x69, 0x50, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feature_feature_api_proto_rawDescOnce sync.Once
	file_feature_feature_api_proto_rawDescData = file_feature_feature_api_proto_rawDesc
)

func file_feature_feature_api_proto_rawDescGZIP() []byte {
	file_feature_feature_api_proto_rawDescOnce.Do(func() {
		file_feature_feature_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_feature_feature_api_proto_rawDescData)
	})
	return file_feature_feature_api_proto_rawDescData
}

var file_feature_feature_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feature_feature_api_proto_goTypes = []interface{}{
	(*Point)(nil),   // 0: feature.feature_api_pb.Point
	(*Feature)(nil), // 1: feature.feature_api_pb.Feature
}
var file_feature_feature_api_proto_depIdxs = []int32{
	0, // 0: feature.feature_api_pb.FeatureService.GetFeature:input_type -> feature.feature_api_pb.Point
	1, // 1: feature.feature_api_pb.FeatureService.GetFeature:output_type -> feature.feature_api_pb.Feature
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feature_feature_api_proto_init() }
func file_feature_feature_api_proto_init() {
	if File_feature_feature_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feature_feature_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_feature_feature_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
			RawDescriptor: file_feature_feature_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feature_feature_api_proto_goTypes,
		DependencyIndexes: file_feature_feature_api_proto_depIdxs,
		MessageInfos:      file_feature_feature_api_proto_msgTypes,
	}.Build()
	File_feature_feature_api_proto = out.File
	file_feature_feature_api_proto_rawDesc = nil
	file_feature_feature_api_proto_goTypes = nil
	file_feature_feature_api_proto_depIdxs = nil
}
