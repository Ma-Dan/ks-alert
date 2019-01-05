// Code generated by protoc-gen-go. DO NOT EDIT.
// source: silence.proto

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

// silence
type SilenceID struct {
	AlertConfigId        string   `protobuf:"bytes,1,opt,name=alert_config_id,proto3" json:"alert_config_id,omitempty"`
	AlertRuleId          string   `protobuf:"bytes,2,opt,name=alert_rule_id,proto3" json:"alert_rule_id,omitempty"`
	ResourceId           string   `protobuf:"bytes,3,opt,name=resource_id,proto3" json:"resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SilenceID) Reset()         { *m = SilenceID{} }
func (m *SilenceID) String() string { return proto.CompactTextString(m) }
func (*SilenceID) ProtoMessage()    {}
func (*SilenceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fc56058cf68dbd8, []int{0}
}

func (m *SilenceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SilenceID.Unmarshal(m, b)
}
func (m *SilenceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SilenceID.Marshal(b, m, deterministic)
}
func (m *SilenceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SilenceID.Merge(m, src)
}
func (m *SilenceID) XXX_Size() int {
	return xxx_messageInfo_SilenceID.Size(m)
}
func (m *SilenceID) XXX_DiscardUnknown() {
	xxx_messageInfo_SilenceID.DiscardUnknown(m)
}

var xxx_messageInfo_SilenceID proto.InternalMessageInfo

func (m *SilenceID) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *SilenceID) GetAlertRuleId() string {
	if m != nil {
		return m.AlertRuleId
	}
	return ""
}

func (m *SilenceID) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

type Silence struct {
	SilenceId            *SilenceID `protobuf:"bytes,1,opt,name=silence_id,proto3" json:"silence_id,omitempty"`
	Dutation             int32      `protobuf:"varint,2,opt,name=dutation,proto3" json:"dutation,omitempty"`
	StartTimestamp       int64      `protobuf:"varint,3,opt,name=start_timestamp,proto3" json:"start_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Silence) Reset()         { *m = Silence{} }
func (m *Silence) String() string { return proto.CompactTextString(m) }
func (*Silence) ProtoMessage()    {}
func (*Silence) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fc56058cf68dbd8, []int{1}
}

func (m *Silence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Silence.Unmarshal(m, b)
}
func (m *Silence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Silence.Marshal(b, m, deterministic)
}
func (m *Silence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Silence.Merge(m, src)
}
func (m *Silence) XXX_Size() int {
	return xxx_messageInfo_Silence.Size(m)
}
func (m *Silence) XXX_DiscardUnknown() {
	xxx_messageInfo_Silence.DiscardUnknown(m)
}

var xxx_messageInfo_Silence proto.InternalMessageInfo

func (m *Silence) GetSilenceId() *SilenceID {
	if m != nil {
		return m.SilenceId
	}
	return nil
}

func (m *Silence) GetDutation() int32 {
	if m != nil {
		return m.Dutation
	}
	return 0
}

func (m *Silence) GetStartTimestamp() int64 {
	if m != nil {
		return m.StartTimestamp
	}
	return 0
}

type SilenceResponse struct {
	Silence              *Silence `protobuf:"bytes,1,opt,name=silence,proto3" json:"silence,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SilenceResponse) Reset()         { *m = SilenceResponse{} }
func (m *SilenceResponse) String() string { return proto.CompactTextString(m) }
func (*SilenceResponse) ProtoMessage()    {}
func (*SilenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fc56058cf68dbd8, []int{2}
}

func (m *SilenceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SilenceResponse.Unmarshal(m, b)
}
func (m *SilenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SilenceResponse.Marshal(b, m, deterministic)
}
func (m *SilenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SilenceResponse.Merge(m, src)
}
func (m *SilenceResponse) XXX_Size() int {
	return xxx_messageInfo_SilenceResponse.Size(m)
}
func (m *SilenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SilenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SilenceResponse proto.InternalMessageInfo

func (m *SilenceResponse) GetSilence() *Silence {
	if m != nil {
		return m.Silence
	}
	return nil
}

func (m *SilenceResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*SilenceID)(nil), "pb.SilenceID")
	proto.RegisterType((*Silence)(nil), "pb.Silence")
	proto.RegisterType((*SilenceResponse)(nil), "pb.SilenceResponse")
}

func init() { proto.RegisterFile("silence.proto", fileDescriptor_7fc56058cf68dbd8) }

var fileDescriptor_7fc56058cf68dbd8 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x6d, 0xe7, 0x9c, 0x7d, 0xa5, 0x4e, 0x32, 0xc4, 0x31, 0x3c, 0x68, 0x4f, 0x9e, 0x8a,
	0x74, 0xdf, 0x40, 0x27, 0xba, 0x9b, 0x28, 0x22, 0x78, 0x29, 0x69, 0xfb, 0x94, 0x42, 0x97, 0x84,
	0x97, 0xf4, 0x03, 0xfb, 0x4d, 0x24, 0x69, 0xca, 0x7a, 0xed, 0x31, 0xbf, 0xf7, 0xfe, 0xbf, 0x7f,
	0x42, 0x20, 0xd1, 0x4d, 0x8b, 0xa2, 0xc2, 0x4c, 0x91, 0x34, 0x92, 0x85, 0xaa, 0xdc, 0xc4, 0x48,
	0x24, 0xa9, 0x07, 0xe9, 0x1b, 0x44, 0x1f, 0xfd, 0xc6, 0x7e, 0xc7, 0xae, 0x61, 0xc9, 0x5b, 0x24,
	0x53, 0x54, 0x52, 0xfc, 0x34, 0xbf, 0x45, 0x53, 0xaf, 0x83, 0xdb, 0xe0, 0x3e, 0x62, 0x57, 0x90,
	0xf4, 0x03, 0xea, 0x5a, 0xb4, 0x38, 0x74, 0x78, 0x05, 0x31, 0xa1, 0x96, 0x1d, 0x55, 0x0e, 0xce,
	0x2c, 0x4c, 0xbf, 0x60, 0xe1, 0x8d, 0xec, 0x0e, 0xc0, 0xd7, 0x0f, 0xaa, 0x38, 0x4f, 0x32, 0x55,
	0x66, 0xc7, 0xca, 0x4b, 0x38, 0xaf, 0x3b, 0xc3, 0x4d, 0x23, 0x85, 0x93, 0xce, 0xed, 0x25, 0xb4,
	0xe1, 0x64, 0x0a, 0xd3, 0x1c, 0x50, 0x1b, 0x7e, 0x50, 0x4e, 0x3c, 0x4b, 0xf7, 0xb0, 0xf4, 0xb9,
	0x77, 0xd4, 0x4a, 0x0a, 0x8d, 0xec, 0x06, 0x16, 0xbe, 0xc0, 0xdb, 0xe3, 0x91, 0x9d, 0xad, 0x61,
	0xee, 0x9e, 0xea, 0xc4, 0x71, 0x1e, 0xd9, 0xd9, 0xb3, 0x05, 0xf9, 0x5f, 0x00, 0x17, 0x7e, 0xeb,
	0x95, 0x8b, 0xba, 0x45, 0x62, 0x5b, 0x48, 0x9e, 0x08, 0xb9, 0xc1, 0x21, 0x3d, 0x56, 0x6d, 0x56,
	0xa3, 0xc3, 0xd0, 0x9e, 0x9e, 0xd8, 0xd0, 0x0e, 0x5b, 0x9c, 0x1c, 0xfa, 0x54, 0xf5, 0xc4, 0xa6,
	0x07, 0x80, 0x17, 0x34, 0x13, 0x12, 0x8f, 0xa7, 0xdf, 0xa1, 0x2a, 0xcb, 0x33, 0xf7, 0xcd, 0xdb,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xf1, 0xec, 0xd3, 0x08, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SilenceHandlerClient is the client API for SilenceHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SilenceHandlerClient interface {
	// silence
	CreateSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error)
	DeleteSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error)
	UpdateSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error)
	GetSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error)
}

type silenceHandlerClient struct {
	cc *grpc.ClientConn
}

func NewSilenceHandlerClient(cc *grpc.ClientConn) SilenceHandlerClient {
	return &silenceHandlerClient{cc}
}

func (c *silenceHandlerClient) CreateSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error) {
	out := new(SilenceResponse)
	err := c.cc.Invoke(ctx, "/pb.SilenceHandler/CreateSilence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *silenceHandlerClient) DeleteSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error) {
	out := new(SilenceResponse)
	err := c.cc.Invoke(ctx, "/pb.SilenceHandler/DeleteSilence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *silenceHandlerClient) UpdateSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error) {
	out := new(SilenceResponse)
	err := c.cc.Invoke(ctx, "/pb.SilenceHandler/UpdateSilence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *silenceHandlerClient) GetSilence(ctx context.Context, in *Silence, opts ...grpc.CallOption) (*SilenceResponse, error) {
	out := new(SilenceResponse)
	err := c.cc.Invoke(ctx, "/pb.SilenceHandler/GetSilence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SilenceHandlerServer is the server API for SilenceHandler service.
type SilenceHandlerServer interface {
	// silence
	CreateSilence(context.Context, *Silence) (*SilenceResponse, error)
	DeleteSilence(context.Context, *Silence) (*SilenceResponse, error)
	UpdateSilence(context.Context, *Silence) (*SilenceResponse, error)
	GetSilence(context.Context, *Silence) (*SilenceResponse, error)
}

func RegisterSilenceHandlerServer(s *grpc.Server, srv SilenceHandlerServer) {
	s.RegisterService(&_SilenceHandler_serviceDesc, srv)
}

func _SilenceHandler_CreateSilence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Silence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SilenceHandlerServer).CreateSilence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SilenceHandler/CreateSilence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SilenceHandlerServer).CreateSilence(ctx, req.(*Silence))
	}
	return interceptor(ctx, in, info, handler)
}

func _SilenceHandler_DeleteSilence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Silence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SilenceHandlerServer).DeleteSilence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SilenceHandler/DeleteSilence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SilenceHandlerServer).DeleteSilence(ctx, req.(*Silence))
	}
	return interceptor(ctx, in, info, handler)
}

func _SilenceHandler_UpdateSilence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Silence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SilenceHandlerServer).UpdateSilence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SilenceHandler/UpdateSilence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SilenceHandlerServer).UpdateSilence(ctx, req.(*Silence))
	}
	return interceptor(ctx, in, info, handler)
}

func _SilenceHandler_GetSilence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Silence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SilenceHandlerServer).GetSilence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SilenceHandler/GetSilence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SilenceHandlerServer).GetSilence(ctx, req.(*Silence))
	}
	return interceptor(ctx, in, info, handler)
}

var _SilenceHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SilenceHandler",
	HandlerType: (*SilenceHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSilence",
			Handler:    _SilenceHandler_CreateSilence_Handler,
		},
		{
			MethodName: "DeleteSilence",
			Handler:    _SilenceHandler_DeleteSilence_Handler,
		},
		{
			MethodName: "UpdateSilence",
			Handler:    _SilenceHandler_UpdateSilence_Handler,
		},
		{
			MethodName: "GetSilence",
			Handler:    _SilenceHandler_GetSilence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "silence.proto",
}
