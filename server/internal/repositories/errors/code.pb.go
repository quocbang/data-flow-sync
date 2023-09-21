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
)

var Code_name = map[int32]string{
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
}

var Code_value = map[string]int32{
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
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x4d, 0x4e, 0xe3, 0x30,
	0x18, 0xc6, 0xf1, 0xd1, 0xa8, 0xb2, 0x66, 0xac, 0xf9, 0xf0, 0x78, 0x46, 0x03, 0x0b, 0x4e, 0x80,
	0xd4, 0x66, 0xc1, 0xba, 0x0b, 0x37, 0x76, 0x8b, 0xd5, 0xd4, 0xae, 0x9c, 0x94, 0x52, 0x36, 0x56,
	0x9a, 0x9a, 0x12, 0xa9, 0xc4, 0xc5, 0x49, 0x41, 0x2c, 0x91, 0xb2, 0x2e, 0x9f, 0x07, 0xe1, 0x08,
	0x85, 0x35, 0x47, 0xe0, 0x00, 0x1c, 0x81, 0x25, 0x6a, 0x0b, 0x12, 0xdb, 0xe7, 0xfd, 0x49, 0xaf,
	0xf4, 0x87, 0x30, 0xb1, 0x23, 0x53, 0x9b, 0x3a, 0x5b, 0x58, 0x0c, 0x8c, 0x73, 0xd6, 0xe5, 0xdb,
	0x17, 0x5f, 0x61, 0xc5, 0xb7, 0x23, 0x83, 0xbf, 0xc1, 0x8a, 0x90, 0x82, 0xa1, 0x2f, 0xf8, 0x1f,
	0xfc, 0xd5, 0x57, 0x52, 0xb4, 0x74, 0x97, 0x84, 0x61, 0x5f, 0x2a, 0x8a, 0xca, 0x39, 0xc0, 0x18,
	0xfe, 0xe8, 0x09, 0xc5, 0x48, 0xa0, 0x59, 0x87, 0xf0, 0x00, 0x2d, 0xca, 0x3a, 0xfe, 0x0d, 0xbf,
	0xaf, 0xa5, 0xec, 0x46, 0xe8, 0xa1, 0xac, 0xe3, 0x2d, 0xf8, 0x9f, 0x04, 0x8a, 0x11, 0x3a, 0xd0,
	0x5c, 0xe8, 0x48, 0x11, 0x11, 0x12, 0x3f, 0xe2, 0x52, 0xa0, 0xc7, 0xb2, 0x8e, 0xff, 0xc2, 0x9f,
	0x1d, 0x1e, 0x86, 0x5c, 0xb4, 0x74, 0x24, 0xdb, 0x4c, 0xa0, 0x97, 0x39, 0xc0, 0x9b, 0x10, 0x7f,
	0x8c, 0x21, 0xf3, 0x15, 0x8b, 0x74, 0x9b, 0x0d, 0xd0, 0xd3, 0x25, 0x58, 0x72, 0x2e, 0xf6, 0x48,
	0xc0, 0xe9, 0x3b, 0xbf, 0xbf, 0x02, 0xcb, 0x97, 0x4d, 0xa9, 0x1a, 0x9c, 0x52, 0x26, 0xd0, 0xdd,
	0xf5, 0x4a, 0xad, 0xae, 0x9a, 0xed, 0x77, 0xb9, 0x62, 0x14, 0xbd, 0x7e, 0x1e, 0x1b, 0x81, 0xf4,
	0xdb, 0x8c, 0xa2, 0xc5, 0x2d, 0xc0, 0x1b, 0xf0, 0x0f, 0xf1, 0x7d, 0xd9, 0x13, 0x91, 0x16, 0x32,
	0xd2, 0x4d, 0xd9, 0x13, 0x14, 0x3d, 0xdf, 0x80, 0xc6, 0xee, 0x41, 0x73, 0x9c, 0x16, 0x47, 0xb3,
	0x61, 0x2d, 0xb1, 0xc7, 0xde, 0xc9, 0xcc, 0x26, 0xc3, 0x38, 0x1b, 0x7b, 0xa3, 0xb8, 0x88, 0xab,
	0x87, 0x13, 0x7b, 0x56, 0xcd, 0xcf, 0xb3, 0xc4, 0xcb, 0x8d, 0x3b, 0x35, 0xce, 0x4b, 0xb3, 0xc2,
	0xb8, 0x2c, 0x9e, 0x78, 0xce, 0x4c, 0x6d, 0x9e, 0x16, 0xd6, 0xa5, 0x26, 0xf7, 0xd6, 0x35, 0x87,
	0x60, 0x15, 0x77, 0xe7, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x16, 0x47, 0xba, 0x86, 0x6a, 0x01, 0x00,
	0x00,
}
