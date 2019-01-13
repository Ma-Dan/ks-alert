// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ent.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Enterprise struct {
	EnterpriseId         string   `protobuf:"bytes,1,opt,name=enterprise_id,json=enterpriseId,proto3" json:"enterprise_id,omitempty"`
	EnterpriseName       string   `protobuf:"bytes,2,opt,name=enterprise_name,json=enterpriseName,proto3" json:"enterprise_name,omitempty"`
	HomePage             string   `protobuf:"bytes,3,opt,name=home_page,json=homePage,proto3" json:"home_page,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Contacts             string   `protobuf:"bytes,7,opt,name=contacts,proto3" json:"contacts,omitempty"`
	Desc                 string   `protobuf:"bytes,8,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Enterprise) Reset()         { *m = Enterprise{} }
func (m *Enterprise) String() string { return proto.CompactTextString(m) }
func (*Enterprise) ProtoMessage()    {}
func (*Enterprise) Descriptor() ([]byte, []int) {
	return fileDescriptor_077dec8aae1e7756, []int{0}
}

func (m *Enterprise) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enterprise.Unmarshal(m, b)
}
func (m *Enterprise) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enterprise.Marshal(b, m, deterministic)
}
func (m *Enterprise) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enterprise.Merge(m, src)
}
func (m *Enterprise) XXX_Size() int {
	return xxx_messageInfo_Enterprise.Size(m)
}
func (m *Enterprise) XXX_DiscardUnknown() {
	xxx_messageInfo_Enterprise.DiscardUnknown(m)
}

var xxx_messageInfo_Enterprise proto.InternalMessageInfo

func (m *Enterprise) GetEnterpriseId() string {
	if m != nil {
		return m.EnterpriseId
	}
	return ""
}

func (m *Enterprise) GetEnterpriseName() string {
	if m != nil {
		return m.EnterpriseName
	}
	return ""
}

func (m *Enterprise) GetHomePage() string {
	if m != nil {
		return m.HomePage
	}
	return ""
}

func (m *Enterprise) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Enterprise) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Enterprise) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Enterprise) GetContacts() string {
	if m != nil {
		return m.Contacts
	}
	return ""
}

func (m *Enterprise) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type EnterpriseSpec struct {
	EnterpriseId         string   `protobuf:"bytes,1,opt,name=enterprise_id,json=enterpriseId,proto3" json:"enterprise_id,omitempty"`
	EnterpriseName       string   `protobuf:"bytes,2,opt,name=enterprise_name,json=enterpriseName,proto3" json:"enterprise_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterpriseSpec) Reset()         { *m = EnterpriseSpec{} }
func (m *EnterpriseSpec) String() string { return proto.CompactTextString(m) }
func (*EnterpriseSpec) ProtoMessage()    {}
func (*EnterpriseSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_077dec8aae1e7756, []int{1}
}

func (m *EnterpriseSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseSpec.Unmarshal(m, b)
}
func (m *EnterpriseSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseSpec.Marshal(b, m, deterministic)
}
func (m *EnterpriseSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseSpec.Merge(m, src)
}
func (m *EnterpriseSpec) XXX_Size() int {
	return xxx_messageInfo_EnterpriseSpec.Size(m)
}
func (m *EnterpriseSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseSpec.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseSpec proto.InternalMessageInfo

func (m *EnterpriseSpec) GetEnterpriseId() string {
	if m != nil {
		return m.EnterpriseId
	}
	return ""
}

func (m *EnterpriseSpec) GetEnterpriseName() string {
	if m != nil {
		return m.EnterpriseName
	}
	return ""
}

type EnterpriseResponse struct {
	Enterprise           *Enterprise `protobuf:"bytes,1,opt,name=enterprise,proto3" json:"enterprise,omitempty"`
	Error                *Error      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EnterpriseResponse) Reset()         { *m = EnterpriseResponse{} }
func (m *EnterpriseResponse) String() string { return proto.CompactTextString(m) }
func (*EnterpriseResponse) ProtoMessage()    {}
func (*EnterpriseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_077dec8aae1e7756, []int{2}
}

func (m *EnterpriseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseResponse.Unmarshal(m, b)
}
func (m *EnterpriseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseResponse.Marshal(b, m, deterministic)
}
func (m *EnterpriseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseResponse.Merge(m, src)
}
func (m *EnterpriseResponse) XXX_Size() int {
	return xxx_messageInfo_EnterpriseResponse.Size(m)
}
func (m *EnterpriseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseResponse proto.InternalMessageInfo

func (m *EnterpriseResponse) GetEnterprise() *Enterprise {
	if m != nil {
		return m.Enterprise
	}
	return nil
}

func (m *EnterpriseResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Enterprise)(nil), "pb.Enterprise")
	proto.RegisterType((*EnterpriseSpec)(nil), "pb.EnterpriseSpec")
	proto.RegisterType((*EnterpriseResponse)(nil), "pb.EnterpriseResponse")
}

func init() { proto.RegisterFile("ent.proto", fileDescriptor_077dec8aae1e7756) }

var fileDescriptor_077dec8aae1e7756 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0xe8, 0x2d, 0xa7, 0xb4, 0x94, 0x23, 0x84, 0xac, 0x32, 0x80, 0xc2, 0x00, 0x53,
	0x86, 0xb2, 0x16, 0x09, 0x71, 0x11, 0xb0, 0x20, 0x54, 0xc4, 0xc2, 0x40, 0xe5, 0xc4, 0x47, 0x6d,
	0xa5, 0xc4, 0xb6, 0x6c, 0x3f, 0x08, 0x2f, 0xca, 0x3b, 0xa0, 0x38, 0x94, 0xa4, 0x1d, 0x18, 0x2a,
	0xb6, 0xfc, 0x5f, 0x3e, 0xfb, 0x8f, 0x73, 0x0c, 0x21, 0x49, 0x17, 0x6b, 0xa3, 0x9c, 0xc2, 0x40,
	0x27, 0xa3, 0x1e, 0x19, 0xa3, 0x4c, 0x09, 0xa2, 0xaf, 0x06, 0xc0, 0xbd, 0x74, 0x64, 0xb4, 0x59,
	0x5a, 0xc2, 0x33, 0xe8, 0xd3, 0x6f, 0x9a, 0x2d, 0x05, 0x6b, 0x9c, 0x36, 0x2e, 0xc2, 0xe9, 0x5e,
	0x05, 0x9f, 0x04, 0x9e, 0xc3, 0x7e, 0x4d, 0x92, 0x3c, 0x27, 0x16, 0x78, 0x6d, 0x50, 0xe1, 0x67,
	0x9e, 0x13, 0x1e, 0x43, 0xb8, 0x50, 0x39, 0xcd, 0x34, 0x9f, 0x13, 0xdb, 0xf5, 0x4a, 0xb7, 0x00,
	0x2f, 0x7c, 0x4e, 0xc8, 0xa0, 0xc3, 0x85, 0x30, 0x64, 0x2d, 0x6b, 0xfa, 0x57, 0xab, 0x88, 0x87,
	0xd0, 0xd2, 0x0b, 0x25, 0x89, 0xb5, 0x3c, 0x2f, 0x43, 0x41, 0x29, 0xe7, 0xcb, 0x8c, 0xb5, 0x4b,
	0xea, 0x03, 0x8e, 0xa0, 0x9b, 0x2a, 0xe9, 0x78, 0xea, 0x2c, 0xeb, 0x94, 0x0d, 0xab, 0x8c, 0x08,
	0x4d, 0x41, 0x36, 0x65, 0x5d, 0xcf, 0xfd, 0x73, 0xf4, 0x01, 0x83, 0xea, 0xb8, 0xaf, 0x9a, 0xd2,
	0xff, 0x3d, 0x72, 0x44, 0x80, 0xd5, 0xfe, 0x53, 0xb2, 0x5a, 0x49, 0x4b, 0x18, 0x03, 0x54, 0x9e,
	0x2f, 0xe8, 0x8d, 0x07, 0xb1, 0x4e, 0xe2, 0x9a, 0x5b, 0x33, 0xf0, 0x04, 0x5a, 0x7e, 0x48, 0xbe,
	0xa4, 0x37, 0x0e, 0xbd, 0x5a, 0x80, 0x69, 0xc9, 0xc7, 0x9f, 0x01, 0x1c, 0x54, 0x6b, 0x1f, 0xb9,
	0x14, 0x19, 0x19, 0x9c, 0xc0, 0xf0, 0xd6, 0x10, 0x77, 0x54, 0x9b, 0xe8, 0x46, 0xcd, 0xe8, 0x68,
	0xa3, 0xf6, 0xe7, 0x13, 0xa3, 0x1d, 0xbc, 0x86, 0xe1, 0x1d, 0x65, 0xb4, 0xb6, 0x1a, 0xd7, 0xed,
	0xe2, 0x87, 0xfd, 0xb1, 0xc3, 0x04, 0x86, 0x6f, 0x5a, 0x6c, 0xdb, 0x7f, 0x05, 0xfd, 0x07, 0x72,
	0xdb, 0x96, 0xdf, 0x34, 0xdf, 0x03, 0x9d, 0x24, 0x6d, 0x7f, 0xad, 0x2f, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x29, 0xdb, 0x48, 0xcc, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EnterpriseHandlerClient is the client API for EnterpriseHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnterpriseHandlerClient interface {
	// enterprise
	CreateEnterprise(ctx context.Context, in *Enterprise, opts ...grpc.CallOption) (*EnterpriseResponse, error)
	DeleteEnterprise(ctx context.Context, in *EnterpriseSpec, opts ...grpc.CallOption) (*EnterpriseResponse, error)
	UpdateEnterprise(ctx context.Context, in *Enterprise, opts ...grpc.CallOption) (*EnterpriseResponse, error)
	GetEnterprise(ctx context.Context, in *EnterpriseSpec, opts ...grpc.CallOption) (*EnterpriseResponse, error)
}

type enterpriseHandlerClient struct {
	cc *grpc.ClientConn
}

func NewEnterpriseHandlerClient(cc *grpc.ClientConn) EnterpriseHandlerClient {
	return &enterpriseHandlerClient{cc}
}

func (c *enterpriseHandlerClient) CreateEnterprise(ctx context.Context, in *Enterprise, opts ...grpc.CallOption) (*EnterpriseResponse, error) {
	out := new(EnterpriseResponse)
	err := c.cc.Invoke(ctx, "/pb.EnterpriseHandler/CreateEnterprise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enterpriseHandlerClient) DeleteEnterprise(ctx context.Context, in *EnterpriseSpec, opts ...grpc.CallOption) (*EnterpriseResponse, error) {
	out := new(EnterpriseResponse)
	err := c.cc.Invoke(ctx, "/pb.EnterpriseHandler/DeleteEnterprise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enterpriseHandlerClient) UpdateEnterprise(ctx context.Context, in *Enterprise, opts ...grpc.CallOption) (*EnterpriseResponse, error) {
	out := new(EnterpriseResponse)
	err := c.cc.Invoke(ctx, "/pb.EnterpriseHandler/UpdateEnterprise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enterpriseHandlerClient) GetEnterprise(ctx context.Context, in *EnterpriseSpec, opts ...grpc.CallOption) (*EnterpriseResponse, error) {
	out := new(EnterpriseResponse)
	err := c.cc.Invoke(ctx, "/pb.EnterpriseHandler/GetEnterprise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnterpriseHandlerServer is the server API for EnterpriseHandler service.
type EnterpriseHandlerServer interface {
	// enterprise
	CreateEnterprise(context.Context, *Enterprise) (*EnterpriseResponse, error)
	DeleteEnterprise(context.Context, *EnterpriseSpec) (*EnterpriseResponse, error)
	UpdateEnterprise(context.Context, *Enterprise) (*EnterpriseResponse, error)
	GetEnterprise(context.Context, *EnterpriseSpec) (*EnterpriseResponse, error)
}

func RegisterEnterpriseHandlerServer(s *grpc.Server, srv EnterpriseHandlerServer) {
	s.RegisterService(&_EnterpriseHandler_serviceDesc, srv)
}

func _EnterpriseHandler_CreateEnterprise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Enterprise)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnterpriseHandlerServer).CreateEnterprise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EnterpriseHandler/CreateEnterprise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnterpriseHandlerServer).CreateEnterprise(ctx, req.(*Enterprise))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnterpriseHandler_DeleteEnterprise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterpriseSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnterpriseHandlerServer).DeleteEnterprise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EnterpriseHandler/DeleteEnterprise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnterpriseHandlerServer).DeleteEnterprise(ctx, req.(*EnterpriseSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnterpriseHandler_UpdateEnterprise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Enterprise)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnterpriseHandlerServer).UpdateEnterprise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EnterpriseHandler/UpdateEnterprise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnterpriseHandlerServer).UpdateEnterprise(ctx, req.(*Enterprise))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnterpriseHandler_GetEnterprise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterpriseSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnterpriseHandlerServer).GetEnterprise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EnterpriseHandler/GetEnterprise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnterpriseHandlerServer).GetEnterprise(ctx, req.(*EnterpriseSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnterpriseHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EnterpriseHandler",
	HandlerType: (*EnterpriseHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEnterprise",
			Handler:    _EnterpriseHandler_CreateEnterprise_Handler,
		},
		{
			MethodName: "DeleteEnterprise",
			Handler:    _EnterpriseHandler_DeleteEnterprise_Handler,
		},
		{
			MethodName: "UpdateEnterprise",
			Handler:    _EnterpriseHandler_UpdateEnterprise_Handler,
		},
		{
			MethodName: "GetEnterprise",
			Handler:    _EnterpriseHandler_GetEnterprise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ent.proto",
}