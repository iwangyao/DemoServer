// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dto/FightLeaveDTO.txt

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

type FightLeaveDTO struct {
	Roomid               int32    `protobuf:"varint,1,opt,name=Roomid,proto3" json:"Roomid,omitempty"`
	Seat                 int32    `protobuf:"varint,2,opt,name=Seat,proto3" json:"Seat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FightLeaveDTO) Reset()         { *m = FightLeaveDTO{} }
func (m *FightLeaveDTO) String() string { return proto.CompactTextString(m) }
func (*FightLeaveDTO) ProtoMessage()    {}
func (*FightLeaveDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_f351908de38f2ef6, []int{0}
}

func (m *FightLeaveDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FightLeaveDTO.Unmarshal(m, b)
}
func (m *FightLeaveDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FightLeaveDTO.Marshal(b, m, deterministic)
}
func (m *FightLeaveDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FightLeaveDTO.Merge(m, src)
}
func (m *FightLeaveDTO) XXX_Size() int {
	return xxx_messageInfo_FightLeaveDTO.Size(m)
}
func (m *FightLeaveDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_FightLeaveDTO.DiscardUnknown(m)
}

var xxx_messageInfo_FightLeaveDTO proto.InternalMessageInfo

func (m *FightLeaveDTO) GetRoomid() int32 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *FightLeaveDTO) GetSeat() int32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

func init() {
	proto.RegisterType((*FightLeaveDTO)(nil), "FightLeaveDTO")
}

func init() { proto.RegisterFile("proto/dto/FightLeaveDTO.txt", fileDescriptor_f351908de38f2ef6) }

var fileDescriptor_f351908de38f2ef6 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x29, 0xc9, 0xd7, 0x77, 0xcb, 0x4c, 0xcf, 0x28, 0xf1, 0x49, 0x4d, 0x2c, 0x4b,
	0x75, 0x09, 0xf1, 0xd7, 0x2b, 0xa9, 0x28, 0x51, 0xb2, 0xe6, 0xe2, 0x45, 0x11, 0x14, 0x12, 0xe3,
	0x62, 0x0b, 0xca, 0xcf, 0xcf, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2,
	0x84, 0x84, 0xb8, 0x58, 0x82, 0x53, 0x13, 0x4b, 0x24, 0x98, 0xc0, 0xa2, 0x60, 0x76, 0x12, 0x1b,
	0xd8, 0x64, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xda, 0x81, 0x53, 0x62, 0x00, 0x00,
	0x00,
}
