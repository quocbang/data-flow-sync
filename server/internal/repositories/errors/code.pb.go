// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: code.proto

package errors

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

type Code int32

const (
	Code_NONE                   Code = 0
	Code_WRONG_PASSWORD         Code = 100100
	Code_UNREAL_EMAIL           Code = 1000000
	Code_WRONG_OPT              Code = 1000001
	Code_ALREADY_IN_TRANSACTION Code = 1000002
	Code_MISSING_TOKEN          Code = 100200
	Code_MISSING_SECRET_KEY     Code = 100300
	Code_INVALID_TOKEN          Code = 100400
	Code_FORBIDDEN              Code = 100500
	Code_TOKEN_EXPIRED          Code = 100600
	Code_TOKEN_BLOCKED          Code = 100800
	Code_ACCOUNT_NOT_FOUND      Code = 100700
	Code_FILE_CONFLICTED        Code = 100900
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:       "NONE",
		100100:  "WRONG_PASSWORD",
		1000000: "UNREAL_EMAIL",
		1000001: "WRONG_OPT",
		1000002: "ALREADY_IN_TRANSACTION",
		100200:  "MISSING_TOKEN",
		100300:  "MISSING_SECRET_KEY",
		100400:  "INVALID_TOKEN",
		100500:  "FORBIDDEN",
		100600:  "TOKEN_EXPIRED",
		100800:  "TOKEN_BLOCKED",
		100700:  "ACCOUNT_NOT_FOUND",
		100900:  "FILE_CONFLICTED",
	}
	Code_value = map[string]int32{
		"NONE":                   0,
		"WRONG_PASSWORD":         100100,
		"UNREAL_EMAIL":           1000000,
		"WRONG_OPT":              1000001,
		"ALREADY_IN_TRANSACTION": 1000002,
		"MISSING_TOKEN":          100200,
		"MISSING_SECRET_KEY":     100300,
		"INVALID_TOKEN":          100400,
		"FORBIDDEN":              100500,
		"TOKEN_EXPIRED":          100600,
		"TOKEN_BLOCKED":          100800,
		"ACCOUNT_NOT_FOUND":      100700,
		"FILE_CONFLICTED":        100900,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_code_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_code_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_code_proto_rawDescGZIP(), []int{0}
}

var File_code_proto protoreflect.FileDescriptor

var file_code_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2a, 0x98, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x0e, 0x57, 0x52, 0x4f, 0x4e, 0x47,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x84, 0x8e, 0x06, 0x12, 0x12, 0x0a,
	0x0c, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x4c, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0xc0, 0x84,
	0x3d, 0x12, 0x0f, 0x0a, 0x09, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x54, 0x10, 0xc1,
	0x84, 0x3d, 0x12, 0x1c, 0x0a, 0x16, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x49, 0x4e,
	0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xc2, 0x84, 0x3d,
	0x12, 0x13, 0x0a, 0x0d, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x10, 0xe8, 0x8e, 0x06, 0x12, 0x18, 0x0a, 0x12, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0xcc, 0x8f, 0x06, 0x12,
	0x13, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x10, 0xb0, 0x90, 0x06, 0x12, 0x0f, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45,
	0x4e, 0x10, 0x94, 0x91, 0x06, 0x12, 0x13, 0x0a, 0x0d, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45,
	0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0xf8, 0x91, 0x06, 0x12, 0x13, 0x0a, 0x0d, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0xc0, 0x93, 0x06, 0x12,
	0x17, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0xdc, 0x92, 0x06, 0x12, 0x15, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x45, 0x44, 0x10, 0xa4, 0x94, 0x06, 0x42,
	0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75,
	0x6f, 0x63, 0x62, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x66, 0x6c, 0x6f, 0x77,
	0x2d, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_code_proto_rawDescOnce sync.Once
	file_code_proto_rawDescData = file_code_proto_rawDesc
)

func file_code_proto_rawDescGZIP() []byte {
	file_code_proto_rawDescOnce.Do(func() {
		file_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_proto_rawDescData)
	})
	return file_code_proto_rawDescData
}

var file_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_code_proto_goTypes = []interface{}{
	(Code)(0), // 0: errors.Code
}
var file_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_code_proto_init() }
func file_code_proto_init() {
	if File_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_code_proto_goTypes,
		DependencyIndexes: file_code_proto_depIdxs,
		EnumInfos:         file_code_proto_enumTypes,
	}.Build()
	File_code_proto = out.File
	file_code_proto_rawDesc = nil
	file_code_proto_goTypes = nil
	file_code_proto_depIdxs = nil
}
