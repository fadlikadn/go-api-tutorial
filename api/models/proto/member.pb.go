// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xdd, 0x6a, 0xb3, 0x40,
	0x10, 0x55, 0xe3, 0xe7, 0x67, 0xa6, 0x49, 0x48, 0xa7, 0x50, 0xc4, 0xde, 0x58, 0xaf, 0xa4, 0x85,
	0x15, 0xf4, 0x09, 0x0a, 0xd9, 0x86, 0x42, 0x92, 0x0b, 0xa1, 0x0f, 0x10, 0x71, 0x2b, 0x0b, 0x31,
	0x2b, 0xee, 0x86, 0xd2, 0xb7, 0x2f, 0xce, 0xb6, 0xa5, 0x5e, 0xf4, 0x6a, 0x67, 0xce, 0x9c, 0x9f,
	0xe5, 0xc0, 0xa2, 0x13, 0x5d, 0x2d, 0x06, 0xd6, 0x0f, 0xca, 0x28, 0xfc, 0x47, 0x4f, 0x7c, 0xd7,
	0x2a, 0xd5, 0x9e, 0x44, 0x4e, 0x5b, 0x7d, 0x79, 0xcb, 0x45, 0xd7, 0x9b, 0x0f, 0xcb, 0x49, 0x2f,
	0x10, 0xec, 0x49, 0x83, 0x2b, 0xf0, 0x64, 0x13, 0xb9, 0x89, 0x9b, 0xcd, 0x2b, 0x4f, 0x36, 0x88,
	0xe0, 0x9f, 0x8f, 0x9d, 0x88, 0x3c, 0x42, 0x68, 0xc6, 0x18, 0xc2, 0xfe, 0xa8, 0xf5, 0xbb, 0x1a,
	0x9a, 0x68, 0x46, 0xf8, 0xcf, 0x8e, 0x8f, 0x10, 0xb4, 0xe2, 0xdc, 0x88, 0x21, 0xf2, 0x13, 0x37,
	0x5b, 0x15, 0x37, 0x36, 0x81, 0x59, 0xfb, 0x2d, 0x9d, 0xaa, 0x2f, 0x4a, 0x9a, 0x03, 0x58, 0x7c,
	0x27, 0xb5, 0xc1, 0x7b, 0xf0, 0x4f, 0x52, 0x9b, 0xc8, 0x4d, 0x66, 0xd9, 0x55, 0xb1, 0x9c, 0x08,
	0x2b, 0x3a, 0x3d, 0x94, 0xb0, 0xf8, 0x6d, 0x84, 0x4b, 0x98, 0xbf, 0x1e, 0x36, 0x7c, 0xfb, 0x72,
	0xe0, 0x9b, 0xb5, 0x83, 0x21, 0xf8, 0xfb, 0xa7, 0x1d, 0x5f, 0xbb, 0x08, 0x10, 0x3c, 0x73, 0x9a,
	0xbd, 0x42, 0xc3, 0x7f, 0x2b, 0xd2, 0x58, 0x42, 0x58, 0x89, 0x56, 0x6a, 0x33, 0x6a, 0x27, 0x01,
	0xf1, 0x2d, 0xb3, 0x05, 0xb1, 0xef, 0x82, 0x18, 0x1f, 0x0b, 0x4a, 0x1d, 0x2c, 0xc1, 0xa7, 0xff,
	0xfd, 0xc1, 0x88, 0xaf, 0x27, 0x46, 0x23, 0x35, 0x75, 0xea, 0x80, 0xb0, 0xf2, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x7d, 0xc7, 0x89, 0x1e, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MembersClient is the client API for Members service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MembersClient interface {
	Register(ctx context.Context, in *Member, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MemberList, error)
}

type membersClient struct {
	cc *grpc.ClientConn
}

func NewMembersClient(cc *grpc.ClientConn) MembersClient {
	return &membersClient{cc}
}

func (c *membersClient) Register(ctx context.Context, in *Member, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Members/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membersClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MemberList, error) {
	out := new(MemberList)
	err := c.cc.Invoke(ctx, "/proto.Members/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MembersServer is the server API for Members service.
type MembersServer interface {
	Register(context.Context, *Member) (*empty.Empty, error)
	List(context.Context, *empty.Empty) (*MemberList, error)
}

// UnimplementedMembersServer can be embedded to have forward compatible implementations.
type UnimplementedMembersServer struct {
}

func (*UnimplementedMembersServer) Register(ctx context.Context, req *Member) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedMembersServer) List(ctx context.Context, req *empty.Empty) (*MemberList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterMembersServer(s *grpc.Server, srv MembersServer) {
	s.RegisterService(&_Members_serviceDesc, srv)
}

func _Members_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Members/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServer).Register(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _Members_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Members/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Members_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Members",
	HandlerType: (*MembersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Members_Register_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Members_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member.proto",
}
