// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alertconfig.proto

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

// alert config
type AlertConfig struct {
	AlertConfigId   string          `protobuf:"bytes,1,opt,name=alert_config_id,json=alertConfigId,proto3" json:"alert_config_id,omitempty"`
	AlertConfigName string          `protobuf:"bytes,2,opt,name=alert_config_name,json=alertConfigName,proto3" json:"alert_config_name,omitempty"`
	AlertRuleGroup  *AlertRuleGroup `protobuf:"bytes,3,opt,name=alert_rule_group,json=alertRuleGroup,proto3" json:"alert_rule_group,omitempty"`
	ResourceGroup   *ResourceGroup  `protobuf:"bytes,4,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	ReceiverGroup   *ReceiverGroup  `protobuf:"bytes,5,opt,name=receiver_group,json=receiverGroup,proto3" json:"receiver_group,omitempty"`
	SeverityId      string          `protobuf:"bytes,6,opt,name=severity_id,json=severityId,proto3" json:"severity_id,omitempty"`
	SeverityCh      string          `protobuf:"bytes,7,opt,name=severity_ch,json=severityCh,proto3" json:"severity_ch,omitempty"`
	//    enum UpdateType {
	//        RECEIVER = 0;
	//        ALERTRULE = 1;
	//        RESOURCE = 2;
	//        SEVERITY = 3;
	//        EFFECTIVETIME = 4;
	//        ENABLE = 5;
	//        DESC = 6;
	//    }
	//
	//    UpdateType update_type = 8;
	EnableStart string `protobuf:"bytes,8,opt,name=enable_start,json=enableStart,proto3" json:"enable_start,omitempty"`
	EnableEnd   string `protobuf:"bytes,9,opt,name=enable_end,json=enableEnd,proto3" json:"enable_end,omitempty"`
	//    google.protobuf.Timestamp enable_start = 8;
	//    google.protobuf.Timestamp enable_end = 9;
	Desc                 string   `protobuf:"bytes,10,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertConfig) Reset()         { *m = AlertConfig{} }
func (m *AlertConfig) String() string { return proto.CompactTextString(m) }
func (*AlertConfig) ProtoMessage()    {}
func (*AlertConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f45947961a9298d9, []int{0}
}

func (m *AlertConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertConfig.Unmarshal(m, b)
}
func (m *AlertConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertConfig.Marshal(b, m, deterministic)
}
func (m *AlertConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertConfig.Merge(m, src)
}
func (m *AlertConfig) XXX_Size() int {
	return xxx_messageInfo_AlertConfig.Size(m)
}
func (m *AlertConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AlertConfig proto.InternalMessageInfo

func (m *AlertConfig) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *AlertConfig) GetAlertConfigName() string {
	if m != nil {
		return m.AlertConfigName
	}
	return ""
}

func (m *AlertConfig) GetAlertRuleGroup() *AlertRuleGroup {
	if m != nil {
		return m.AlertRuleGroup
	}
	return nil
}

func (m *AlertConfig) GetResourceGroup() *ResourceGroup {
	if m != nil {
		return m.ResourceGroup
	}
	return nil
}

func (m *AlertConfig) GetReceiverGroup() *ReceiverGroup {
	if m != nil {
		return m.ReceiverGroup
	}
	return nil
}

func (m *AlertConfig) GetSeverityId() string {
	if m != nil {
		return m.SeverityId
	}
	return ""
}

func (m *AlertConfig) GetSeverityCh() string {
	if m != nil {
		return m.SeverityCh
	}
	return ""
}

func (m *AlertConfig) GetEnableStart() string {
	if m != nil {
		return m.EnableStart
	}
	return ""
}

func (m *AlertConfig) GetEnableEnd() string {
	if m != nil {
		return m.EnableEnd
	}
	return ""
}

func (m *AlertConfig) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type AlertConfigSpec struct {
	AlertConfigId        string   `protobuf:"bytes,1,opt,name=alert_config_id,json=alertConfigId,proto3" json:"alert_config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertConfigSpec) Reset()         { *m = AlertConfigSpec{} }
func (m *AlertConfigSpec) String() string { return proto.CompactTextString(m) }
func (*AlertConfigSpec) ProtoMessage()    {}
func (*AlertConfigSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f45947961a9298d9, []int{1}
}

func (m *AlertConfigSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertConfigSpec.Unmarshal(m, b)
}
func (m *AlertConfigSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertConfigSpec.Marshal(b, m, deterministic)
}
func (m *AlertConfigSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertConfigSpec.Merge(m, src)
}
func (m *AlertConfigSpec) XXX_Size() int {
	return xxx_messageInfo_AlertConfigSpec.Size(m)
}
func (m *AlertConfigSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertConfigSpec.DiscardUnknown(m)
}

var xxx_messageInfo_AlertConfigSpec proto.InternalMessageInfo

func (m *AlertConfigSpec) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

type AlertConfigResponse struct {
	AlertConfig          *AlertConfig `protobuf:"bytes,1,opt,name=alert_config,json=alertConfig,proto3" json:"alert_config,omitempty"`
	Error                *Error       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AlertConfigResponse) Reset()         { *m = AlertConfigResponse{} }
func (m *AlertConfigResponse) String() string { return proto.CompactTextString(m) }
func (*AlertConfigResponse) ProtoMessage()    {}
func (*AlertConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f45947961a9298d9, []int{2}
}

func (m *AlertConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertConfigResponse.Unmarshal(m, b)
}
func (m *AlertConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertConfigResponse.Marshal(b, m, deterministic)
}
func (m *AlertConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertConfigResponse.Merge(m, src)
}
func (m *AlertConfigResponse) XXX_Size() int {
	return xxx_messageInfo_AlertConfigResponse.Size(m)
}
func (m *AlertConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlertConfigResponse proto.InternalMessageInfo

func (m *AlertConfigResponse) GetAlertConfig() *AlertConfig {
	if m != nil {
		return m.AlertConfig
	}
	return nil
}

func (m *AlertConfigResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*AlertConfig)(nil), "pb.AlertConfig")
	proto.RegisterType((*AlertConfigSpec)(nil), "pb.AlertConfigSpec")
	proto.RegisterType((*AlertConfigResponse)(nil), "pb.AlertConfigResponse")
}

func init() { proto.RegisterFile("alertconfig.proto", fileDescriptor_f45947961a9298d9) }

var fileDescriptor_f45947961a9298d9 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x80, 0x49, 0x9a, 0x16, 0x32, 0xdb, 0xc6, 0x78, 0x7a, 0xc0, 0x8a, 0x84, 0x5a, 0x7c, 0x40,
	0x15, 0x87, 0x1c, 0xcc, 0x05, 0x24, 0xa4, 0x02, 0xa6, 0x2a, 0xbd, 0x70, 0x70, 0xc5, 0x85, 0x8b,
	0xb5, 0xf6, 0x0e, 0xad, 0x91, 0x6b, 0xaf, 0xd6, 0x76, 0x25, 0x1e, 0x86, 0x27, 0xe0, 0x25, 0xd1,
	0xfe, 0x24, 0x5d, 0xe7, 0x80, 0x2a, 0xe5, 0x96, 0x99, 0xfd, 0xbe, 0x99, 0xc9, 0xce, 0x1a, 0x42,
	0x5e, 0x93, 0xea, 0xcb, 0xb6, 0xf9, 0x59, 0xdd, 0xac, 0xa4, 0x6a, 0xfb, 0x16, 0xa7, 0xb2, 0x58,
	0x32, 0x52, 0xaa, 0x55, 0x36, 0xb1, 0x5c, 0x28, 0x2a, 0xa9, 0xba, 0xa7, 0x75, 0x1c, 0x18, 0x47,
	0x0d, 0x35, 0x3d, 0x00, 0x5d, 0x3b, 0xa8, 0xd2, 0xc5, 0xf1, 0xdf, 0x3d, 0x60, 0x9f, 0x34, 0x93,
	0x9a, 0xba, 0xf8, 0x1a, 0xac, 0x92, 0xdb, 0x3e, 0x79, 0x25, 0xa2, 0xc9, 0xe9, 0xe4, 0x6c, 0x9e,
	0x1d, 0xf1, 0x07, 0xea, 0x4a, 0xe0, 0x1b, 0x37, 0xce, 0x9a, 0x6b, 0xf8, 0x1d, 0x45, 0x53, 0x43,
	0x06, 0x1e, 0xf9, 0x8d, 0xdf, 0x11, 0x7e, 0x80, 0xe7, 0x96, 0xd5, 0x73, 0xe4, 0x37, 0xaa, 0x1d,
	0x64, 0xb4, 0x77, 0x3a, 0x39, 0x63, 0x09, 0xae, 0x64, 0xb1, 0x32, 0xed, 0xb3, 0xa1, 0xa6, 0x4b,
	0x7d, 0x92, 0x2d, 0xf8, 0x28, 0xc6, 0x77, 0xb0, 0x99, 0xd9, 0xb9, 0x33, 0xe3, 0x86, 0xda, 0xcd,
	0xdc, 0x89, 0x55, 0x8f, 0x94, 0x1f, 0x5a, 0xd3, 0x5e, 0x87, 0x33, 0xf7, 0x7d, 0xd3, 0x9e, 0x6c,
	0x4c, 0x2f, 0xc4, 0x13, 0x60, 0x1d, 0xdd, 0x93, 0xaa, 0xfa, 0xdf, 0xfa, 0x06, 0x0e, 0xcc, 0xff,
	0x82, 0x75, 0xea, 0x4a, 0x8c, 0x80, 0xf2, 0x36, 0x7a, 0x3a, 0x06, 0xd2, 0x5b, 0x7c, 0x05, 0x87,
	0xd4, 0xf0, 0xa2, 0xa6, 0xbc, 0xeb, 0xb9, 0xea, 0xa3, 0x67, 0x86, 0x60, 0x36, 0x77, 0xad, 0x53,
	0xf8, 0x12, 0xc0, 0x21, 0xd4, 0x88, 0x68, 0x6e, 0x80, 0xb9, 0xcd, 0x5c, 0x34, 0x02, 0x11, 0x66,
	0x82, 0xba, 0x32, 0x02, 0x73, 0x60, 0x7e, 0xc7, 0xef, 0x21, 0xf0, 0x96, 0x75, 0x2d, 0xa9, 0x7c,
	0xec, 0xc2, 0xe2, 0x5f, 0x70, 0xec, 0xa9, 0x19, 0x75, 0xb2, 0x6d, 0x3a, 0xc2, 0x04, 0x0e, 0x7d,
	0xdd, 0xb8, 0x2c, 0x09, 0x36, 0x7b, 0x71, 0x38, 0xf3, 0x8a, 0xe1, 0x09, 0xec, 0x9b, 0x37, 0x67,
	0xf6, 0xcd, 0x92, 0xb9, 0x86, 0x2f, 0x74, 0x22, 0xb3, 0xf9, 0xe4, 0xcf, 0x14, 0xd0, 0xb3, 0xbf,
	0xf2, 0x46, 0xd4, 0xa4, 0xf0, 0x1c, 0xc2, 0x54, 0x11, 0xef, 0xc9, 0x7f, 0x70, 0xdb, 0xad, 0x96,
	0x2f, 0xb6, 0x7b, 0xbb, 0x51, 0xe3, 0x27, 0x98, 0x42, 0xf8, 0x85, 0x6a, 0x1a, 0x17, 0x38, 0xde,
	0xe2, 0xf5, 0xad, 0xfc, 0xaf, 0xc8, 0x39, 0x84, 0xdf, 0xa5, 0xd8, 0x61, 0x8a, 0x8f, 0xb0, 0xb8,
	0xa4, 0x7e, 0x87, 0x11, 0x3e, 0xcf, 0x7e, 0x4c, 0x65, 0x51, 0x1c, 0x98, 0x2f, 0xf0, 0xed, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xb3, 0x0b, 0x97, 0xd8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertConfigHandlerClient is the client API for AlertConfigHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertConfigHandlerClient interface {
	// alert
	CreateAlertConfig(ctx context.Context, in *AlertConfig, opts ...grpc.CallOption) (*AlertConfigResponse, error)
	DeleteAlertConfig(ctx context.Context, in *AlertConfigSpec, opts ...grpc.CallOption) (*AlertConfigResponse, error)
	UpdateAlertConfig(ctx context.Context, in *AlertConfig, opts ...grpc.CallOption) (*AlertConfigResponse, error)
	GetAlertConfig(ctx context.Context, in *AlertConfigSpec, opts ...grpc.CallOption) (*AlertConfigResponse, error)
}

type alertConfigHandlerClient struct {
	cc *grpc.ClientConn
}

func NewAlertConfigHandlerClient(cc *grpc.ClientConn) AlertConfigHandlerClient {
	return &alertConfigHandlerClient{cc}
}

func (c *alertConfigHandlerClient) CreateAlertConfig(ctx context.Context, in *AlertConfig, opts ...grpc.CallOption) (*AlertConfigResponse, error) {
	out := new(AlertConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.AlertConfigHandler/CreateAlertConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConfigHandlerClient) DeleteAlertConfig(ctx context.Context, in *AlertConfigSpec, opts ...grpc.CallOption) (*AlertConfigResponse, error) {
	out := new(AlertConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.AlertConfigHandler/DeleteAlertConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConfigHandlerClient) UpdateAlertConfig(ctx context.Context, in *AlertConfig, opts ...grpc.CallOption) (*AlertConfigResponse, error) {
	out := new(AlertConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.AlertConfigHandler/UpdateAlertConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConfigHandlerClient) GetAlertConfig(ctx context.Context, in *AlertConfigSpec, opts ...grpc.CallOption) (*AlertConfigResponse, error) {
	out := new(AlertConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.AlertConfigHandler/GetAlertConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertConfigHandlerServer is the server API for AlertConfigHandler service.
type AlertConfigHandlerServer interface {
	// alert
	CreateAlertConfig(context.Context, *AlertConfig) (*AlertConfigResponse, error)
	DeleteAlertConfig(context.Context, *AlertConfigSpec) (*AlertConfigResponse, error)
	UpdateAlertConfig(context.Context, *AlertConfig) (*AlertConfigResponse, error)
	GetAlertConfig(context.Context, *AlertConfigSpec) (*AlertConfigResponse, error)
}

func RegisterAlertConfigHandlerServer(s *grpc.Server, srv AlertConfigHandlerServer) {
	s.RegisterService(&_AlertConfigHandler_serviceDesc, srv)
}

func _AlertConfigHandler_CreateAlertConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConfigHandlerServer).CreateAlertConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlertConfigHandler/CreateAlertConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConfigHandlerServer).CreateAlertConfig(ctx, req.(*AlertConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConfigHandler_DeleteAlertConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertConfigSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConfigHandlerServer).DeleteAlertConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlertConfigHandler/DeleteAlertConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConfigHandlerServer).DeleteAlertConfig(ctx, req.(*AlertConfigSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConfigHandler_UpdateAlertConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConfigHandlerServer).UpdateAlertConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlertConfigHandler/UpdateAlertConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConfigHandlerServer).UpdateAlertConfig(ctx, req.(*AlertConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConfigHandler_GetAlertConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertConfigSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConfigHandlerServer).GetAlertConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlertConfigHandler/GetAlertConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConfigHandlerServer).GetAlertConfig(ctx, req.(*AlertConfigSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlertConfigHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AlertConfigHandler",
	HandlerType: (*AlertConfigHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlertConfig",
			Handler:    _AlertConfigHandler_CreateAlertConfig_Handler,
		},
		{
			MethodName: "DeleteAlertConfig",
			Handler:    _AlertConfigHandler_DeleteAlertConfig_Handler,
		},
		{
			MethodName: "UpdateAlertConfig",
			Handler:    _AlertConfigHandler_UpdateAlertConfig_Handler,
		},
		{
			MethodName: "GetAlertConfig",
			Handler:    _AlertConfigHandler_GetAlertConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alertconfig.proto",
}
