// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Types.txt

package DTO

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

type MsgTypes int32

const (
	MsgTypes_TYPE_LOGIN MsgTypes = 0
	MsgTypes_TYPE_USER  MsgTypes = 1
	MsgTypes_TYPE_MATCH MsgTypes = 2
	MsgTypes_TYPE_FIGHT MsgTypes = 3
)

var MsgTypes_name = map[int32]string{
	0: "TYPE_LOGIN",
	1: "TYPE_USER",
	2: "TYPE_MATCH",
	3: "TYPE_FIGHT",
}

var MsgTypes_value = map[string]int32{
	"TYPE_LOGIN": 0,
	"TYPE_USER":  1,
	"TYPE_MATCH": 2,
	"TYPE_FIGHT": 3,
}

func (x MsgTypes) String() string {
	return proto.EnumName(MsgTypes_name, int32(x))
}

func (MsgTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f8907c4c6e116672, []int{0}
}

func init() {
	proto.RegisterEnum("MsgTypes", MsgTypes_name, MsgTypes_value)
}

func init() { proto.RegisterFile("Types.txt", fileDescriptor_f8907c4c6e116672) }

var fileDescriptor_f8907c4c6e116672 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0xa9, 0x28, 0xd1, 0xf2, 0xe4, 0xe2, 0xf0, 0x2d, 0x4e, 0x07, 0xf3, 0x85, 0xf8,
	0xb8, 0xb8, 0x42, 0x22, 0x03, 0x5c, 0xe3, 0x7d, 0xfc, 0xdd, 0x3d, 0xfd, 0x04, 0x18, 0x84, 0x78,
	0xb9, 0x38, 0xc1, 0xfc, 0xd0, 0x60, 0xd7, 0x20, 0x01, 0x46, 0xb8, 0xb4, 0xaf, 0x63, 0x88, 0xb3,
	0x87, 0x00, 0x13, 0x9c, 0xef, 0xe6, 0xe9, 0xee, 0x11, 0x22, 0xc0, 0x9c, 0xc4, 0x56, 0x50, 0x94,
	0x5f, 0x92, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x35, 0x4e, 0xa0, 0x6b, 0x5e, 0x00, 0x00,
	0x00,
}
