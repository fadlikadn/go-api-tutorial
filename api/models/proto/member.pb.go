// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member.proto

package proto

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

type MemberGender int32

const (
	MemberGender_UNDEGINED MemberGender = 0
	MemberGender_MALE      MemberGender = 1
	MemberGender_FEMALE    MemberGender = 2
)

var MemberGender_name = map[int32]string{
	0: "UNDEGINED",
	1: "MALE",
	2: "FEMALE",
}

var MemberGender_value = map[string]int32{
	"UNDEGINED": 0,
	"MALE":      1,
	"FEMALE":    2,
}

func (x MemberGender) String() string {
	return proto.EnumName(MemberGender_name, int32(x))
}

func (MemberGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{0}
}

type Member struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string       `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Gender               MemberGender `protobuf:"varint,4,opt,name=gender,proto3,enum=proto.MemberGender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{0}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Member) GetGender() MemberGender {
	if m != nil {
		return m.Gender
	}
	return MemberGender_UNDEGINED
}

type MemberList struct {
	List                 []*Member `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberList) Reset()         { *m = MemberList{} }
func (m *MemberList) String() string { return proto.CompactTextString(m) }
func (*MemberList) ProtoMessage()    {}
func (*MemberList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{1}
}

func (m *MemberList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberList.Unmarshal(m, b)
}
func (m *MemberList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberList.Marshal(b, m, deterministic)
}
func (m *MemberList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberList.Merge(m, src)
}
func (m *MemberList) XXX_Size() int {
	return xxx_messageInfo_MemberList.Size(m)
}
func (m *MemberList) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberList.DiscardUnknown(m)
}

var xxx_messageInfo_MemberList proto.InternalMessageInfo

func (m *MemberList) GetList() []*Member {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.MemberGender", MemberGender_name, MemberGender_value)
	proto.RegisterType((*Member)(nil), "proto.Member")
	proto.RegisterType((*MemberList)(nil), "proto.MemberList")
}

func init() { proto.RegisterFile("member.proto", fileDescriptor_9b9836b7e13de206) }

var fileDescriptor_9b9836b7e13de206 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0xcd, 0x4d,
	0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xa5, 0x5c, 0x6c,
	0xbe, 0x60, 0x61, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x08, 0x98,
	0x2d, 0x24, 0xc5, 0xc5, 0x51, 0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x94, 0x22, 0xc1, 0x0c, 0x16,
	0x87, 0xf3, 0x85, 0xb4, 0xb9, 0xd8, 0xd2, 0x53, 0xf3, 0x52, 0x52, 0x8b, 0x24, 0x58, 0x14, 0x18,
	0x35, 0xf8, 0x8c, 0x84, 0x21, 0x16, 0xe9, 0x41, 0x8c, 0x77, 0x07, 0x4b, 0x05, 0x41, 0x95, 0x28,
	0xe9, 0x73, 0x71, 0x41, 0xc4, 0x7d, 0x32, 0x8b, 0x4b, 0x84, 0x14, 0xb9, 0x58, 0x72, 0x32, 0x8b,
	0x4b, 0x24, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0x51, 0x34, 0x06, 0x81, 0xa5, 0xb4, 0x8c,
	0xb9, 0x78, 0x90, 0x0d, 0x12, 0xe2, 0xe5, 0xe2, 0x0c, 0xf5, 0x73, 0x71, 0x75, 0xf7, 0xf4, 0x73,
	0x75, 0x11, 0x60, 0x10, 0xe2, 0xe0, 0x62, 0xf1, 0x75, 0xf4, 0x71, 0x15, 0x60, 0x14, 0xe2, 0xe2,
	0x62, 0x73, 0x73, 0x05, 0xb3, 0x99, 0x92, 0xd8, 0xc0, 0x06, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x96, 0x8e, 0x88, 0x72, 0xfa, 0x00, 0x00, 0x00,
}