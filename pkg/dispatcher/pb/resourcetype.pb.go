// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resourcetype.proto

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

type ResourceType struct {
	ResourceTypeId       string            `protobuf:"bytes,1,opt,name=resource_type_id,proto3" json:"resource_type_id,omitempty"`
	ResourceTypeName     string            `protobuf:"bytes,2,opt,name=resource_type_name,proto3" json:"resource_type_name,omitempty"`
	ProductId            string            `protobuf:"bytes,3,opt,name=product_id,proto3" json:"product_id,omitempty"`
	MonitorCenterHost    string            `protobuf:"bytes,4,opt,name=monitor_center_host,proto3" json:"monitor_center_host,omitempty"`
	MonitorCenterPort    int32             `protobuf:"varint,5,opt,name=monitor_center_port,proto3" json:"monitor_center_port,omitempty"`
	ResourceUriTmpl      *ResourceUriTmpls `protobuf:"bytes,6,opt,name=resource_uri_tmpl,proto3" json:"resource_uri_tmpl,omitempty"`
	Enable               bool              `protobuf:"varint,7,opt,name=enable,proto3" json:"enable,omitempty"`
	Desc                 string            `protobuf:"bytes,8,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResourceType) Reset()         { *m = ResourceType{} }
func (m *ResourceType) String() string { return proto.CompactTextString(m) }
func (*ResourceType) ProtoMessage()    {}
func (*ResourceType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a19646b98e5042c5, []int{0}
}

func (m *ResourceType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceType.Unmarshal(m, b)
}
func (m *ResourceType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceType.Marshal(b, m, deterministic)
}
func (m *ResourceType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceType.Merge(m, src)
}
func (m *ResourceType) XXX_Size() int {
	return xxx_messageInfo_ResourceType.Size(m)
}
func (m *ResourceType) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceType.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceType proto.InternalMessageInfo

func (m *ResourceType) GetResourceTypeId() string {
	if m != nil {
		return m.ResourceTypeId
	}
	return ""
}

func (m *ResourceType) GetResourceTypeName() string {
	if m != nil {
		return m.ResourceTypeName
	}
	return ""
}

func (m *ResourceType) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *ResourceType) GetMonitorCenterHost() string {
	if m != nil {
		return m.MonitorCenterHost
	}
	return ""
}

func (m *ResourceType) GetMonitorCenterPort() int32 {
	if m != nil {
		return m.MonitorCenterPort
	}
	return 0
}

func (m *ResourceType) GetResourceUriTmpl() *ResourceUriTmpls {
	if m != nil {
		return m.ResourceUriTmpl
	}
	return nil
}

func (m *ResourceType) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *ResourceType) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ResourceTypeSpec struct {
	ResourceTypeId string `protobuf:"bytes,1,opt,name=resource_type_id,proto3" json:"resource_type_id,omitempty"`
	// resource_type_name + product_name + enterprise_name
	ResourceTypeName     string   `protobuf:"bytes,2,opt,name=resource_type_name,proto3" json:"resource_type_name,omitempty"`
	ProductId            string   `protobuf:"bytes,3,opt,name=product_id,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceTypeSpec) Reset()         { *m = ResourceTypeSpec{} }
func (m *ResourceTypeSpec) String() string { return proto.CompactTextString(m) }
func (*ResourceTypeSpec) ProtoMessage()    {}
func (*ResourceTypeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a19646b98e5042c5, []int{1}
}

func (m *ResourceTypeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceTypeSpec.Unmarshal(m, b)
}
func (m *ResourceTypeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceTypeSpec.Marshal(b, m, deterministic)
}
func (m *ResourceTypeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceTypeSpec.Merge(m, src)
}
func (m *ResourceTypeSpec) XXX_Size() int {
	return xxx_messageInfo_ResourceTypeSpec.Size(m)
}
func (m *ResourceTypeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceTypeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceTypeSpec proto.InternalMessageInfo

func (m *ResourceTypeSpec) GetResourceTypeId() string {
	if m != nil {
		return m.ResourceTypeId
	}
	return ""
}

func (m *ResourceTypeSpec) GetResourceTypeName() string {
	if m != nil {
		return m.ResourceTypeName
	}
	return ""
}

func (m *ResourceTypeSpec) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

type ResourceTypeResponse struct {
	ResourceType         *ResourceType `protobuf:"bytes,1,opt,name=resource_type,proto3" json:"resource_type,omitempty"`
	Error                *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResourceTypeResponse) Reset()         { *m = ResourceTypeResponse{} }
func (m *ResourceTypeResponse) String() string { return proto.CompactTextString(m) }
func (*ResourceTypeResponse) ProtoMessage()    {}
func (*ResourceTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a19646b98e5042c5, []int{2}
}

func (m *ResourceTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceTypeResponse.Unmarshal(m, b)
}
func (m *ResourceTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceTypeResponse.Marshal(b, m, deterministic)
}
func (m *ResourceTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceTypeResponse.Merge(m, src)
}
func (m *ResourceTypeResponse) XXX_Size() int {
	return xxx_messageInfo_ResourceTypeResponse.Size(m)
}
func (m *ResourceTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceTypeResponse proto.InternalMessageInfo

func (m *ResourceTypeResponse) GetResourceType() *ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return nil
}

func (m *ResourceTypeResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceType)(nil), "pb.ResourceType")
	proto.RegisterType((*ResourceTypeSpec)(nil), "pb.ResourceTypeSpec")
	proto.RegisterType((*ResourceTypeResponse)(nil), "pb.ResourceTypeResponse")
}

func init() { proto.RegisterFile("resourcetype.proto", fileDescriptor_a19646b98e5042c5) }

var fileDescriptor_a19646b98e5042c5 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4e, 0xfa, 0x40,
	0x14, 0xc5, 0xff, 0xed, 0x1f, 0x10, 0x6e, 0x51, 0x71, 0xc0, 0x64, 0x52, 0x37, 0x4d, 0x37, 0x76,
	0x85, 0x49, 0x7d, 0x03, 0xf0, 0x6b, 0x8d, 0xb0, 0xd0, 0x98, 0x34, 0xfd, 0xb8, 0x89, 0x4d, 0xda,
	0xce, 0xe4, 0x76, 0xba, 0xe0, 0x95, 0x7c, 0x2b, 0xdf, 0xc4, 0x4c, 0xa1, 0x5a, 0x30, 0xba, 0xc1,
	0xed, 0xb9, 0xa7, 0xbf, 0x73, 0xee, 0xed, 0x00, 0x23, 0x2c, 0x45, 0x45, 0x31, 0xaa, 0xb5, 0xc4,
	0xa9, 0x24, 0xa1, 0x04, 0x33, 0x65, 0x64, 0x5b, 0x48, 0x24, 0x68, 0x23, 0xd8, 0xe7, 0x8d, 0xa9,
	0xa2, 0x54, 0xe5, 0x32, 0xdb, 0xc8, 0xee, 0xbb, 0x01, 0xc3, 0xc5, 0x76, 0xb2, 0x5c, 0x4b, 0x64,
	0x1c, 0x46, 0x8d, 0x33, 0xd0, 0xbc, 0x20, 0x4d, 0xb8, 0xe1, 0x18, 0xde, 0x80, 0xd9, 0x5f, 0x41,
	0x9b, 0x49, 0x11, 0xe6, 0xc8, 0xcd, 0x7a, 0xc6, 0x00, 0x24, 0x89, 0xa4, 0x8a, 0x95, 0xf6, 0xff,
	0xaf, 0xb5, 0x0b, 0x18, 0xe7, 0xa2, 0x48, 0x95, 0xa0, 0x20, 0xc6, 0x42, 0x21, 0x05, 0xaf, 0xa2,
	0x54, 0xbc, 0xf3, 0xc3, 0x50, 0x0a, 0x52, 0xbc, 0xeb, 0x18, 0x5e, 0x97, 0x5d, 0xc1, 0xd9, 0x67,
	0x52, 0x45, 0x69, 0xa0, 0xfb, 0xf2, 0x9e, 0x63, 0x78, 0x96, 0x3f, 0x99, 0xca, 0x68, 0xda, 0x14,
	0x5e, 0x51, 0xba, 0xcc, 0x65, 0x56, 0xb2, 0x13, 0xe8, 0x61, 0x11, 0x46, 0x19, 0xf2, 0x23, 0xc7,
	0xf0, 0xfa, 0x6c, 0x08, 0x9d, 0x04, 0xcb, 0x98, 0xf7, 0x75, 0x96, 0xfb, 0x02, 0xa3, 0xf6, 0x8a,
	0x8f, 0x12, 0xe3, 0xbf, 0x5b, 0xd3, 0x7d, 0x82, 0x49, 0x9b, 0xbe, 0xc0, 0x52, 0x8a, 0xa2, 0x44,
	0x76, 0x09, 0xc7, 0x3b, 0x9c, 0x1a, 0x6f, 0xf9, 0xa3, 0xf6, 0x02, 0xdb, 0x8b, 0x77, 0xeb, 0x1f,
	0x55, 0x67, 0x58, 0xfe, 0x40, 0x1b, 0x6e, 0xb5, 0xe0, 0xbf, 0x99, 0x30, 0x6e, 0x5b, 0x1f, 0xc2,
	0x22, 0xc9, 0x90, 0xd8, 0x0c, 0xd8, 0x9c, 0x30, 0x54, 0xb8, 0xc3, 0xf9, 0x46, 0xb6, 0xf9, 0xbe,
	0xd2, 0x94, 0x73, 0xff, 0xb1, 0x3b, 0x60, 0x37, 0x98, 0xe1, 0x1e, 0x63, 0xb2, 0xff, 0x85, 0x3e,
	0xd6, 0xaf, 0x9c, 0x19, 0xb0, 0x95, 0x4c, 0x0e, 0xeb, 0x32, 0x87, 0xd3, 0x7b, 0x54, 0x87, 0x15,
	0x99, 0x75, 0x9e, 0x4d, 0x19, 0x45, 0xbd, 0xfa, 0x59, 0x5f, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x64, 0x6d, 0x6b, 0xa8, 0x14, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceTypeHandlerClient is the client API for ResourceTypeHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceTypeHandlerClient interface {
	// resource type
	CreateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
	DeleteResourceType(ctx context.Context, in *ResourceTypeSpec, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
	UpdateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
	GetResourceType(ctx context.Context, in *ResourceTypeSpec, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
}

type resourceTypeHandlerClient struct {
	cc *grpc.ClientConn
}

func NewResourceTypeHandlerClient(cc *grpc.ClientConn) ResourceTypeHandlerClient {
	return &resourceTypeHandlerClient{cc}
}

func (c *resourceTypeHandlerClient) CreateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceTypeHandler/CreateResourceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTypeHandlerClient) DeleteResourceType(ctx context.Context, in *ResourceTypeSpec, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceTypeHandler/DeleteResourceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTypeHandlerClient) UpdateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceTypeHandler/UpdateResourceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTypeHandlerClient) GetResourceType(ctx context.Context, in *ResourceTypeSpec, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceTypeHandler/GetResourceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceTypeHandlerServer is the server API for ResourceTypeHandler service.
type ResourceTypeHandlerServer interface {
	// resource type
	CreateResourceType(context.Context, *ResourceType) (*ResourceTypeResponse, error)
	DeleteResourceType(context.Context, *ResourceTypeSpec) (*ResourceTypeResponse, error)
	UpdateResourceType(context.Context, *ResourceType) (*ResourceTypeResponse, error)
	GetResourceType(context.Context, *ResourceTypeSpec) (*ResourceTypeResponse, error)
}

func RegisterResourceTypeHandlerServer(s *grpc.Server, srv ResourceTypeHandlerServer) {
	s.RegisterService(&_ResourceTypeHandler_serviceDesc, srv)
}

func _ResourceTypeHandler_CreateResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).CreateResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/CreateResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).CreateResourceType(ctx, req.(*ResourceType))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTypeHandler_DeleteResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTypeSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).DeleteResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/DeleteResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).DeleteResourceType(ctx, req.(*ResourceTypeSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTypeHandler_UpdateResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).UpdateResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/UpdateResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).UpdateResourceType(ctx, req.(*ResourceType))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTypeHandler_GetResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTypeSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).GetResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/GetResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).GetResourceType(ctx, req.(*ResourceTypeSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceTypeHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ResourceTypeHandler",
	HandlerType: (*ResourceTypeHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResourceType",
			Handler:    _ResourceTypeHandler_CreateResourceType_Handler,
		},
		{
			MethodName: "DeleteResourceType",
			Handler:    _ResourceTypeHandler_DeleteResourceType_Handler,
		},
		{
			MethodName: "UpdateResourceType",
			Handler:    _ResourceTypeHandler_UpdateResourceType_Handler,
		},
		{
			MethodName: "GetResourceType",
			Handler:    _ResourceTypeHandler_GetResourceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resourcetype.proto",
}
