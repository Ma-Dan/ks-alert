// Code generated by protoc-gen-go. DO NOT EDIT.
// source: receiver.proto

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

// receiver group
type ReceiverGroup struct {
	ReceiverGroupId      string      `protobuf:"bytes,1,opt,name=receiver_group_id,proto3" json:"receiver_group_id,omitempty"`
	ReceiverGroupName    string      `protobuf:"bytes,2,opt,name=receiver_group_name,proto3" json:"receiver_group_name,omitempty"`
	Webhook              string      `protobuf:"bytes,3,opt,name=webhook,proto3" json:"webhook,omitempty"`
	WebhookEnable        bool        `protobuf:"varint,4,opt,name=webhook_enable,proto3" json:"webhook_enable,omitempty"`
	Desc                 string      `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Receivers            []*Receiver `protobuf:"bytes,6,rep,name=receivers,proto3" json:"receivers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReceiverGroup) Reset()         { *m = ReceiverGroup{} }
func (m *ReceiverGroup) String() string { return proto.CompactTextString(m) }
func (*ReceiverGroup) ProtoMessage()    {}
func (*ReceiverGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{0}
}

func (m *ReceiverGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiverGroup.Unmarshal(m, b)
}
func (m *ReceiverGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiverGroup.Marshal(b, m, deterministic)
}
func (m *ReceiverGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiverGroup.Merge(m, src)
}
func (m *ReceiverGroup) XXX_Size() int {
	return xxx_messageInfo_ReceiverGroup.Size(m)
}
func (m *ReceiverGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiverGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiverGroup proto.InternalMessageInfo

func (m *ReceiverGroup) GetReceiverGroupId() string {
	if m != nil {
		return m.ReceiverGroupId
	}
	return ""
}

func (m *ReceiverGroup) GetReceiverGroupName() string {
	if m != nil {
		return m.ReceiverGroupName
	}
	return ""
}

func (m *ReceiverGroup) GetWebhook() string {
	if m != nil {
		return m.Webhook
	}
	return ""
}

func (m *ReceiverGroup) GetWebhookEnable() bool {
	if m != nil {
		return m.WebhookEnable
	}
	return false
}

func (m *ReceiverGroup) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ReceiverGroup) GetReceivers() []*Receiver {
	if m != nil {
		return m.Receivers
	}
	return nil
}

type Receiver struct {
	ReceiverId           string   `protobuf:"bytes,1,opt,name=receiver_id,proto3" json:"receiver_id,omitempty"`
	ReceiverName         string   `protobuf:"bytes,2,opt,name=receiver_name,proto3" json:"receiver_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Wechat               string   `protobuf:"bytes,5,opt,name=wechat,proto3" json:"wechat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Receiver) Reset()         { *m = Receiver{} }
func (m *Receiver) String() string { return proto.CompactTextString(m) }
func (*Receiver) ProtoMessage()    {}
func (*Receiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{1}
}

func (m *Receiver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receiver.Unmarshal(m, b)
}
func (m *Receiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receiver.Marshal(b, m, deterministic)
}
func (m *Receiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receiver.Merge(m, src)
}
func (m *Receiver) XXX_Size() int {
	return xxx_messageInfo_Receiver.Size(m)
}
func (m *Receiver) XXX_DiscardUnknown() {
	xxx_messageInfo_Receiver.DiscardUnknown(m)
}

var xxx_messageInfo_Receiver proto.InternalMessageInfo

func (m *Receiver) GetReceiverId() string {
	if m != nil {
		return m.ReceiverId
	}
	return ""
}

func (m *Receiver) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *Receiver) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Receiver) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Receiver) GetWechat() string {
	if m != nil {
		return m.Wechat
	}
	return ""
}

type ReceiverGroupSpec struct {
	ReceiverGroupId      string   `protobuf:"bytes,1,opt,name=receiver_group_id,proto3" json:"receiver_group_id,omitempty"`
	ReceiverId           string   `protobuf:"bytes,2,opt,name=receiver_id,proto3" json:"receiver_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiverGroupSpec) Reset()         { *m = ReceiverGroupSpec{} }
func (m *ReceiverGroupSpec) String() string { return proto.CompactTextString(m) }
func (*ReceiverGroupSpec) ProtoMessage()    {}
func (*ReceiverGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{2}
}

func (m *ReceiverGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiverGroupSpec.Unmarshal(m, b)
}
func (m *ReceiverGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiverGroupSpec.Marshal(b, m, deterministic)
}
func (m *ReceiverGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiverGroupSpec.Merge(m, src)
}
func (m *ReceiverGroupSpec) XXX_Size() int {
	return xxx_messageInfo_ReceiverGroupSpec.Size(m)
}
func (m *ReceiverGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiverGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiverGroupSpec proto.InternalMessageInfo

func (m *ReceiverGroupSpec) GetReceiverGroupId() string {
	if m != nil {
		return m.ReceiverGroupId
	}
	return ""
}

func (m *ReceiverGroupSpec) GetReceiverId() string {
	if m != nil {
		return m.ReceiverId
	}
	return ""
}

type ReceiverGroupResponse struct {
	ReceiverGroup        *ReceiverGroup `protobuf:"bytes,1,opt,name=receiver_group,proto3" json:"receiver_group,omitempty"`
	Error                *Error         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReceiverGroupResponse) Reset()         { *m = ReceiverGroupResponse{} }
func (m *ReceiverGroupResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiverGroupResponse) ProtoMessage()    {}
func (*ReceiverGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{3}
}

func (m *ReceiverGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiverGroupResponse.Unmarshal(m, b)
}
func (m *ReceiverGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiverGroupResponse.Marshal(b, m, deterministic)
}
func (m *ReceiverGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiverGroupResponse.Merge(m, src)
}
func (m *ReceiverGroupResponse) XXX_Size() int {
	return xxx_messageInfo_ReceiverGroupResponse.Size(m)
}
func (m *ReceiverGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiverGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiverGroupResponse proto.InternalMessageInfo

func (m *ReceiverGroupResponse) GetReceiverGroup() *ReceiverGroup {
	if m != nil {
		return m.ReceiverGroup
	}
	return nil
}

func (m *ReceiverGroupResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ReceiverGroup)(nil), "pb.ReceiverGroup")
	proto.RegisterType((*Receiver)(nil), "pb.Receiver")
	proto.RegisterType((*ReceiverGroupSpec)(nil), "pb.ReceiverGroupSpec")
	proto.RegisterType((*ReceiverGroupResponse)(nil), "pb.ReceiverGroupResponse")
}

func init() { proto.RegisterFile("receiver.proto", fileDescriptor_4b7296e1d2b388c5) }

var fileDescriptor_4b7296e1d2b388c5 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xbf, 0x16, 0xca, 0x47, 0x6f, 0x69, 0x09, 0x43, 0x30, 0x03, 0x2e, 0x24, 0x5d, 0xe1,
	0x86, 0x45, 0x7d, 0x01, 0x14, 0x0c, 0xae, 0x31, 0x6e, 0x8c, 0x09, 0xe9, 0x9f, 0x1b, 0x21, 0x96,
	0xce, 0x64, 0x5a, 0xe5, 0x6d, 0xf4, 0x55, 0x4d, 0xa7, 0x9d, 0x0a, 0x34, 0x61, 0xa1, 0xbb, 0xe6,
	0xdc, 0x73, 0x4f, 0x7f, 0x3d, 0xb7, 0xe0, 0x08, 0x0c, 0x71, 0xfb, 0x81, 0x62, 0xca, 0x05, 0xcb,
	0x18, 0xd1, 0x79, 0x30, 0xb2, 0x50, 0x08, 0x56, 0x0a, 0xee, 0x97, 0x06, 0xf6, 0xaa, 0xf4, 0x2c,
	0x05, 0x7b, 0xe7, 0x64, 0x08, 0x3d, 0xb5, 0xb4, 0x7e, 0xcd, 0x95, 0xf5, 0x36, 0xa2, 0xda, 0x58,
	0x9b, 0x98, 0xe4, 0x12, 0xfa, 0x27, 0xa3, 0xc4, 0xdf, 0x21, 0xd5, 0xe5, 0xb0, 0x0b, 0xff, 0xf7,
	0x18, 0x6c, 0x18, 0x7b, 0xa3, 0x0d, 0x29, 0x5c, 0x80, 0x53, 0x0a, 0x6b, 0x4c, 0xfc, 0x20, 0x46,
	0xda, 0x1c, 0x6b, 0x93, 0x36, 0xe9, 0x40, 0x33, 0xc2, 0x34, 0xa4, 0x86, 0x74, 0x5d, 0x81, 0xa9,
	0x32, 0x53, 0xda, 0x1a, 0x37, 0x26, 0x96, 0xd7, 0x99, 0xf2, 0x60, 0xaa, 0xa0, 0xdc, 0x08, 0xda,
	0xea, 0x99, 0xf4, 0xc1, 0xaa, 0x00, 0x2a, 0xaa, 0x01, 0xd8, 0x95, 0x78, 0xc0, 0x63, 0x83, 0x81,
	0x3b, 0x7f, 0x1b, 0x97, 0x34, 0x36, 0x18, 0x7c, 0xc3, 0x92, 0x02, 0xc2, 0x24, 0x0e, 0xb4, 0xf6,
	0x18, 0x6e, 0xfc, 0xac, 0xc0, 0x70, 0xe7, 0xd0, 0x3b, 0xaa, 0xe1, 0x91, 0x63, 0x78, 0xae, 0x8a,
	0x13, 0x12, 0xf9, 0x4a, 0xf7, 0x05, 0x06, 0x47, 0x21, 0x2b, 0x4c, 0x39, 0x4b, 0x52, 0x24, 0xd7,
	0x3f, 0x87, 0x28, 0x82, 0x64, 0x8a, 0xe5, 0xf5, 0x0e, 0xbf, 0xb4, 0xa8, 0x9f, 0x82, 0x21, 0xef,
	0x23, 0x23, 0x2d, 0xcf, 0xcc, 0x1d, 0xf7, 0xb9, 0xe0, 0x7d, 0xea, 0xd0, 0x55, 0xde, 0x07, 0x3f,
	0x89, 0x62, 0x14, 0x64, 0x06, 0xce, 0x5c, 0xa0, 0x9f, 0x61, 0x55, 0x51, 0x3d, 0x72, 0x34, 0xac,
	0x49, 0x0a, 0xcc, 0xfd, 0x47, 0x6e, 0xc1, 0x5a, 0x62, 0x56, 0xad, 0x0f, 0x6a, 0xde, 0xbc, 0x89,
	0xf3, 0x11, 0x33, 0x70, 0x9e, 0x78, 0xf4, 0x17, 0x88, 0x05, 0x38, 0x0b, 0x8c, 0xf1, 0x20, 0xe1,
	0x17, 0x1c, 0x77, 0xcd, 0x67, 0x9d, 0x07, 0x41, 0x4b, 0xfe, 0xd8, 0x37, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x77, 0x3d, 0x0a, 0xfb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReceiverHandlerClient is the client API for ReceiverHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReceiverHandlerClient interface {
	// Receiver
	CreateReceiver(ctx context.Context, in *ReceiverGroup, opts ...grpc.CallOption) (*ReceiverGroupResponse, error)
	GetReceiver(ctx context.Context, in *ReceiverGroupSpec, opts ...grpc.CallOption) (*ReceiverGroupResponse, error)
	UpdateReceiver(ctx context.Context, in *ReceiverGroup, opts ...grpc.CallOption) (*ReceiverGroupResponse, error)
	DeleteReceiver(ctx context.Context, in *ReceiverGroupSpec, opts ...grpc.CallOption) (*ReceiverGroupResponse, error)
}

type receiverHandlerClient struct {
	cc *grpc.ClientConn
}

func NewReceiverHandlerClient(cc *grpc.ClientConn) ReceiverHandlerClient {
	return &receiverHandlerClient{cc}
}

func (c *receiverHandlerClient) CreateReceiver(ctx context.Context, in *ReceiverGroup, opts ...grpc.CallOption) (*ReceiverGroupResponse, error) {
	out := new(ReceiverGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ReceiverHandler/CreateReceiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiverHandlerClient) GetReceiver(ctx context.Context, in *ReceiverGroupSpec, opts ...grpc.CallOption) (*ReceiverGroupResponse, error) {
	out := new(ReceiverGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ReceiverHandler/GetReceiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiverHandlerClient) UpdateReceiver(ctx context.Context, in *ReceiverGroup, opts ...grpc.CallOption) (*ReceiverGroupResponse, error) {
	out := new(ReceiverGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ReceiverHandler/UpdateReceiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiverHandlerClient) DeleteReceiver(ctx context.Context, in *ReceiverGroupSpec, opts ...grpc.CallOption) (*ReceiverGroupResponse, error) {
	out := new(ReceiverGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ReceiverHandler/DeleteReceiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiverHandlerServer is the server API for ReceiverHandler service.
type ReceiverHandlerServer interface {
	// Receiver
	CreateReceiver(context.Context, *ReceiverGroup) (*ReceiverGroupResponse, error)
	GetReceiver(context.Context, *ReceiverGroupSpec) (*ReceiverGroupResponse, error)
	UpdateReceiver(context.Context, *ReceiverGroup) (*ReceiverGroupResponse, error)
	DeleteReceiver(context.Context, *ReceiverGroupSpec) (*ReceiverGroupResponse, error)
}

func RegisterReceiverHandlerServer(s *grpc.Server, srv ReceiverHandlerServer) {
	s.RegisterService(&_ReceiverHandler_serviceDesc, srv)
}

func _ReceiverHandler_CreateReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiverGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverHandlerServer).CreateReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReceiverHandler/CreateReceiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverHandlerServer).CreateReceiver(ctx, req.(*ReceiverGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReceiverHandler_GetReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiverGroupSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverHandlerServer).GetReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReceiverHandler/GetReceiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverHandlerServer).GetReceiver(ctx, req.(*ReceiverGroupSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReceiverHandler_UpdateReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiverGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverHandlerServer).UpdateReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReceiverHandler/UpdateReceiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverHandlerServer).UpdateReceiver(ctx, req.(*ReceiverGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReceiverHandler_DeleteReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiverGroupSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverHandlerServer).DeleteReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReceiverHandler/DeleteReceiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverHandlerServer).DeleteReceiver(ctx, req.(*ReceiverGroupSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReceiverHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReceiverHandler",
	HandlerType: (*ReceiverHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReceiver",
			Handler:    _ReceiverHandler_CreateReceiver_Handler,
		},
		{
			MethodName: "GetReceiver",
			Handler:    _ReceiverHandler_GetReceiver_Handler,
		},
		{
			MethodName: "UpdateReceiver",
			Handler:    _ReceiverHandler_UpdateReceiver_Handler,
		},
		{
			MethodName: "DeleteReceiver",
			Handler:    _ReceiverHandler_DeleteReceiver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiver.proto",
}
