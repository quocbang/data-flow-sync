// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.proto

package errors

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Code int32

const (
	Code_NONE               Code = 0
	Code_WRONG_PASSWORD     Code = 100100
	Code_MISSING_TOKEN      Code = 100200
	Code_MISSING_SECRET_KEY Code = 100300
	Code_INVALID_TOKEN      Code = 100400
	Code_FORBIDDEN          Code = 100500
	Code_TOKEN_EXPIRED      Code = 100600
	Code_TOKEN_BLOCKED      Code = 100800
	Code_ACCOUNT_NOT_FOUND  Code = 100700
)

var Code_name = map[int32]string{
	0:      "NONE",
	100100: "WRONG_PASSWORD",
	100200: "MISSING_TOKEN",
	100300: "MISSING_SECRET_KEY",
	100400: "INVALID_TOKEN",
	100500: "FORBIDDEN",
	100600: "TOKEN_EXPIRED",
	100800: "TOKEN_BLOCKED",
	100700: "ACCOUNT_NOT_FOUND",
}

var Code_value = map[string]int32{
	"NONE":               0,
	"WRONG_PASSWORD":     100100,
	"MISSING_TOKEN":      100200,
	"MISSING_SECRET_KEY": 100300,
	"INVALID_TOKEN":      100400,
	"FORBIDDEN":          100500,
	"TOKEN_EXPIRED":      100600,
	"TOKEN_BLOCKED":      100800,
	"ACCOUNT_NOT_FOUND":  100700,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e9b0151640170c3, []int{0}
}

func init() {
	proto.RegisterEnum("errors.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("code.proto", fileDescriptor_6e9b0151640170c3) }

var fileDescriptor_6e9b0151640170c3 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x5f, 0x4a, 0xc3, 0x30,
	0x00, 0xc7, 0x71, 0x85, 0x11, 0x34, 0xa0, 0xc6, 0x28, 0xe8, 0x19, 0x84, 0x2d, 0x0f, 0x9e, 0xa0,
	0x6d, 0xd2, 0x19, 0x3a, 0x93, 0xd1, 0x76, 0x4e, 0x7d, 0x09, 0xfd, 0x13, 0x67, 0x61, 0x36, 0x33,
	0xcd, 0x14, 0xdf, 0x7d, 0x9e, 0x7f, 0x0f, 0xe2, 0x09, 0xc4, 0x03, 0x78, 0x04, 0x0f, 0xe0, 0x11,
	0x7c, 0x14, 0xaa, 0x82, 0xaf, 0xbf, 0xdf, 0xe7, 0xe9, 0x0b, 0x61, 0x61, 0x4a, 0xdd, 0x9b, 0x59,
	0xe3, 0x0c, 0x06, 0xda, 0x5a, 0x63, 0x9b, 0xbd, 0xd7, 0x65, 0xd8, 0x09, 0x4c, 0xa9, 0xf1, 0x0a,
	0xec, 0x08, 0x29, 0x18, 0x5a, 0xc2, 0xdb, 0x70, 0x7d, 0x1c, 0x4b, 0xd1, 0x57, 0x43, 0x2f, 0x49,
	0xc6, 0x32, 0xa6, 0xe8, 0x76, 0x01, 0xf0, 0x16, 0x5c, 0x3b, 0xe4, 0x49, 0xc2, 0x45, 0x5f, 0xa5,
	0x32, 0x62, 0x02, 0x7d, 0x2e, 0x00, 0xde, 0x85, 0xf8, 0x6f, 0x4c, 0x58, 0x10, 0xb3, 0x54, 0x45,
	0xec, 0x04, 0xbd, 0xdf, 0xb5, 0x9c, 0x8b, 0x23, 0x6f, 0xc0, 0xe9, 0x2f, 0x7f, 0xb9, 0x07, 0x78,
	0x03, 0xae, 0x86, 0x32, 0xf6, 0x39, 0xa5, 0x4c, 0xa0, 0xe7, 0x87, 0x56, 0xb5, 0xaf, 0x62, 0xc7,
	0x43, 0x1e, 0x33, 0x8a, 0xbe, 0xfe, 0x8f, 0xfe, 0x40, 0x06, 0x11, 0xa3, 0xe8, 0xed, 0x09, 0xe0,
	0x1d, 0xb8, 0xe9, 0x05, 0x81, 0x1c, 0x89, 0x54, 0x09, 0x99, 0xaa, 0x50, 0x8e, 0x04, 0x45, 0x1f,
	0x8f, 0xc0, 0x3f, 0x38, 0x0d, 0x27, 0x95, 0x3b, 0x9f, 0xe7, 0xbd, 0xc2, 0x5c, 0x90, 0xcb, 0xb9,
	0x29, 0xf2, 0xac, 0x9e, 0x90, 0x32, 0x73, 0x59, 0xf7, 0x6c, 0x6a, 0xae, 0xbb, 0xcd, 0x4d, 0x5d,
	0x90, 0x46, 0xdb, 0x2b, 0x6d, 0x49, 0x55, 0x3b, 0x6d, 0xeb, 0x6c, 0x4a, 0xac, 0x9e, 0x99, 0xa6,
	0x72, 0xc6, 0x56, 0xba, 0x21, 0x3f, 0x29, 0x72, 0xd0, 0x96, 0xd9, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x72, 0xf1, 0x30, 0xb4, 0x27, 0x01, 0x00, 0x00,
}
