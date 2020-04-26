// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Fundamentals/Protobuf/Uuid.proto

package protobuf

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

type Uuid struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uuid) Reset()         { *m = Uuid{} }
func (m *Uuid) String() string { return proto.CompactTextString(m) }
func (*Uuid) ProtoMessage()    {}
func (*Uuid) Descriptor() ([]byte, []int) {
	return fileDescriptor_403cfe90b4adb20d, []int{0}
}

func (m *Uuid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uuid.Unmarshal(m, b)
}
func (m *Uuid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uuid.Marshal(b, m, deterministic)
}
func (m *Uuid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uuid.Merge(m, src)
}
func (m *Uuid) XXX_Size() int {
	return xxx_messageInfo_Uuid.Size(m)
}
func (m *Uuid) XXX_DiscardUnknown() {
	xxx_messageInfo_Uuid.DiscardUnknown(m)
}

var xxx_messageInfo_Uuid proto.InternalMessageInfo

func (m *Uuid) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Uuid)(nil), "dolittle.protobuf.Uuid")
}

func init() { proto.RegisterFile("Fundamentals/Protobuf/Uuid.proto", fileDescriptor_403cfe90b4adb20d) }

var fileDescriptor_403cfe90b4adb20d = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x70, 0x2b, 0xcd, 0x4b,
	0x49, 0xcc, 0x4d, 0xcd, 0x2b, 0x49, 0xcc, 0x29, 0xd6, 0x0f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a,
	0x4d, 0xd3, 0x0f, 0x2d, 0xcd, 0x4c, 0xd1, 0x2b, 0x00, 0xf1, 0x84, 0x04, 0x53, 0xf2, 0x73, 0x32,
	0x4b, 0x4a, 0x72, 0x52, 0x21, 0xfc, 0xa4, 0xd2, 0x34, 0x25, 0x19, 0x2e, 0x16, 0x90, 0x02, 0x21,
	0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x08,
	0xc7, 0xc9, 0x27, 0x4a, 0x27, 0x3d, 0x5f, 0x0f, 0xae, 0x2b, 0x33, 0x5f, 0x3f, 0x39, 0x23, 0x31,
	0xbf, 0xb8, 0x24, 0xb5, 0xb8, 0xa4, 0x58, 0x3f, 0x39, 0x3f, 0xaf, 0xa4, 0x28, 0x31, 0xb9, 0xa4,
	0x58, 0x1f, 0x66, 0xda, 0x2a, 0x26, 0x69, 0x17, 0x98, 0x5a, 0x98, 0xfd, 0x7a, 0xce, 0x30, 0x65,
	0x49, 0x6c, 0x60, 0x75, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xbb, 0x1c, 0x8e, 0xa9,
	0x00, 0x00, 0x00,
}
