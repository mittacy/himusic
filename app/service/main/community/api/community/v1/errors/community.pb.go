// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api/community/v1/errors/community.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/api/kratos/api"
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

type Community int32

const (
	Community_NameExists       Community = 0
	Community_CommunityNoFound Community = 1 // 社区不存在
	Community_MysqlErr         Community = 2 // Mysql操作错误
	Community_FieldInvalid     Community = 3 // 字段不符合规定
)

// Enum value maps for Community.
var (
	Community_name = map[int32]string{
		0: "NameExists",
		1: "CommunityNoFound",
		2: "MysqlErr",
		3: "FieldInvalid",
	}
	Community_value = map[string]int32{
		"NameExists":       0,
		"CommunityNoFound": 1,
		"MysqlErr":         2,
		"FieldInvalid":     3,
	}
)

func (x Community) Enum() *Community {
	p := new(Community)
	*p = x
	return p
}

func (x Community) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Community) Descriptor() protoreflect.EnumDescriptor {
	return file_api_community_v1_errors_community_proto_enumTypes[0].Descriptor()
}

func (Community) Type() protoreflect.EnumType {
	return &file_api_community_v1_errors_community_proto_enumTypes[0]
}

func (x Community) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Community.Descriptor instead.
func (Community) EnumDescriptor() ([]byte, []int) {
	return file_api_community_v1_errors_community_proto_rawDescGZIP(), []int{0}
}

var File_api_community_v1_errors_community_proto protoreflect.FileDescriptor

var file_api_community_v1_errors_community_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x1a, 0x1c, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x56, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x45, 0x72, 0x72, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x10, 0x03, 0x1a, 0x03, 0xc0, 0x3e, 0x01, 0x42, 0x67, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69,
	0x74, 0x74, 0x61, 0x63, 0x79, 0x2f, 0x68, 0x69, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x74, 0x69, 0x6e, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x12, 0x41,
	0x50, 0x49, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_community_v1_errors_community_proto_rawDescOnce sync.Once
	file_api_community_v1_errors_community_proto_rawDescData = file_api_community_v1_errors_community_proto_rawDesc
)

func file_api_community_v1_errors_community_proto_rawDescGZIP() []byte {
	file_api_community_v1_errors_community_proto_rawDescOnce.Do(func() {
		file_api_community_v1_errors_community_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_community_v1_errors_community_proto_rawDescData)
	})
	return file_api_community_v1_errors_community_proto_rawDescData
}

var file_api_community_v1_errors_community_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_community_v1_errors_community_proto_goTypes = []interface{}{
	(Community)(0), // 0: api.community.v1.errors.Community
}
var file_api_community_v1_errors_community_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_community_v1_errors_community_proto_init() }
func file_api_community_v1_errors_community_proto_init() {
	if File_api_community_v1_errors_community_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_community_v1_errors_community_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_community_v1_errors_community_proto_goTypes,
		DependencyIndexes: file_api_community_v1_errors_community_proto_depIdxs,
		EnumInfos:         file_api_community_v1_errors_community_proto_enumTypes,
	}.Build()
	File_api_community_v1_errors_community_proto = out.File
	file_api_community_v1_errors_community_proto_rawDesc = nil
	file_api_community_v1_errors_community_proto_goTypes = nil
	file_api_community_v1_errors_community_proto_depIdxs = nil
}