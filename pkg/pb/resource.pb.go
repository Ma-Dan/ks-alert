// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource.proto

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

// resource group
type ResourceGroup struct {
	ResourceGroupId      string           `protobuf:"bytes,1,opt,name=resource_group_id,json=resourceGroupId,proto3" json:"resource_group_id,omitempty"`
	ResourceGroupName    string           `protobuf:"bytes,2,opt,name=resource_group_name,json=resourceGroupName,proto3" json:"resource_group_name,omitempty"`
	ResourceUriTmpl      *ResourceUriTmpl `protobuf:"bytes,3,opt,name=resource_uri_tmpl,json=resourceUriTmpl,proto3" json:"resource_uri_tmpl,omitempty"`
	Desc                 string           `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	ResourceTypeId       string           `protobuf:"bytes,5,opt,name=resource_type_id,json=resourceTypeId,proto3" json:"resource_type_id,omitempty"`
	Resources            []*Resource      `protobuf:"bytes,6,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ResourceGroup) Reset()         { *m = ResourceGroup{} }
func (m *ResourceGroup) String() string { return proto.CompactTextString(m) }
func (*ResourceGroup) ProtoMessage()    {}
func (*ResourceGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b72f771c35e3b8, []int{0}
}

func (m *ResourceGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceGroup.Unmarshal(m, b)
}
func (m *ResourceGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceGroup.Marshal(b, m, deterministic)
}
func (m *ResourceGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceGroup.Merge(m, src)
}
func (m *ResourceGroup) XXX_Size() int {
	return xxx_messageInfo_ResourceGroup.Size(m)
}
func (m *ResourceGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceGroup proto.InternalMessageInfo

func (m *ResourceGroup) GetResourceGroupId() string {
	if m != nil {
		return m.ResourceGroupId
	}
	return ""
}

func (m *ResourceGroup) GetResourceGroupName() string {
	if m != nil {
		return m.ResourceGroupName
	}
	return ""
}

func (m *ResourceGroup) GetResourceUriTmpl() *ResourceUriTmpl {
	if m != nil {
		return m.ResourceUriTmpl
	}
	return nil
}

func (m *ResourceGroup) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ResourceGroup) GetResourceTypeId() string {
	if m != nil {
		return m.ResourceTypeId
	}
	return ""
}

func (m *ResourceGroup) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type ResourceGroupSpec struct {
	ResourceGroupId      string   `protobuf:"bytes,1,opt,name=resource_group_id,json=resourceGroupId,proto3" json:"resource_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceGroupSpec) Reset()         { *m = ResourceGroupSpec{} }
func (m *ResourceGroupSpec) String() string { return proto.CompactTextString(m) }
func (*ResourceGroupSpec) ProtoMessage()    {}
func (*ResourceGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b72f771c35e3b8, []int{1}
}

func (m *ResourceGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceGroupSpec.Unmarshal(m, b)
}
func (m *ResourceGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceGroupSpec.Marshal(b, m, deterministic)
}
func (m *ResourceGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceGroupSpec.Merge(m, src)
}
func (m *ResourceGroupSpec) XXX_Size() int {
	return xxx_messageInfo_ResourceGroupSpec.Size(m)
}
func (m *ResourceGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceGroupSpec proto.InternalMessageInfo

func (m *ResourceGroupSpec) GetResourceGroupId() string {
	if m != nil {
		return m.ResourceGroupId
	}
	return ""
}

type Resource struct {
	ResourceId           string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceName         string   `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceGroupId      string   `protobuf:"bytes,3,opt,name=resource_group_id,json=resourceGroupId,proto3" json:"resource_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b72f771c35e3b8, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *Resource) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Resource) GetResourceGroupId() string {
	if m != nil {
		return m.ResourceGroupId
	}
	return ""
}

type ResourceGroupResponse struct {
	ResourceGroup        *ResourceGroup `protobuf:"bytes,1,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	Error                *Error         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResourceGroupResponse) Reset()         { *m = ResourceGroupResponse{} }
func (m *ResourceGroupResponse) String() string { return proto.CompactTextString(m) }
func (*ResourceGroupResponse) ProtoMessage()    {}
func (*ResourceGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b72f771c35e3b8, []int{3}
}

func (m *ResourceGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceGroupResponse.Unmarshal(m, b)
}
func (m *ResourceGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceGroupResponse.Marshal(b, m, deterministic)
}
func (m *ResourceGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceGroupResponse.Merge(m, src)
}
func (m *ResourceGroupResponse) XXX_Size() int {
	return xxx_messageInfo_ResourceGroupResponse.Size(m)
}
func (m *ResourceGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceGroupResponse proto.InternalMessageInfo

func (m *ResourceGroupResponse) GetResourceGroup() *ResourceGroup {
	if m != nil {
		return m.ResourceGroup
	}
	return nil
}

func (m *ResourceGroupResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceGroup)(nil), "pb.ResourceGroup")
	proto.RegisterType((*ResourceGroupSpec)(nil), "pb.ResourceGroupSpec")
	proto.RegisterType((*Resource)(nil), "pb.Resource")
	proto.RegisterType((*ResourceGroupResponse)(nil), "pb.ResourceGroupResponse")
}

func init() { proto.RegisterFile("resource.proto", fileDescriptor_d1b72f771c35e3b8) }

var fileDescriptor_d1b72f771c35e3b8 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x3d, 0x6f, 0xb2, 0x50,
	0x14, 0x7e, 0xc5, 0x8f, 0xbc, 0x1e, 0x14, 0xcb, 0x35, 0x26, 0xd4, 0x45, 0x43, 0x17, 0xe2, 0xc0,
	0x40, 0x97, 0x6e, 0xda, 0xd6, 0xc6, 0xba, 0x74, 0xa0, 0xba, 0x74, 0x31, 0x20, 0x27, 0x0d, 0x09,
	0xc8, 0xcd, 0x05, 0x06, 0x93, 0xfe, 0x8e, 0x8e, 0xfd, 0xad, 0x0d, 0x17, 0x2f, 0x42, 0x35, 0x4d,
	0xda, 0x6e, 0xe6, 0x39, 0xcf, 0x57, 0x1e, 0x2f, 0xa0, 0x30, 0x8c, 0xa3, 0x94, 0x6d, 0xd1, 0xa4,
	0x2c, 0x4a, 0x22, 0x22, 0x51, 0x77, 0x28, 0x23, 0x63, 0x11, 0xcb, 0x81, 0xe1, 0x40, 0x10, 0x52,
	0xe6, 0x27, 0x21, 0x0d, 0x72, 0x58, 0x7f, 0x97, 0xa0, 0x6b, 0x1f, 0x2e, 0x0b, 0x16, 0xa5, 0x94,
	0x4c, 0x40, 0x15, 0xd4, 0xcd, 0x6b, 0x86, 0x6c, 0x7c, 0x4f, 0xab, 0x8d, 0x6b, 0x46, 0xdb, 0xee,
	0xb1, 0x32, 0x73, 0xe9, 0x11, 0x13, 0xfa, 0x5f, 0xb8, 0x3b, 0x27, 0x44, 0x4d, 0xe2, 0x6c, 0xb5,
	0xc2, 0x7e, 0x72, 0x42, 0x24, 0xd3, 0x92, 0x77, 0xca, 0xfc, 0x4d, 0x56, 0x44, 0xab, 0x8f, 0x6b,
	0x86, 0x6c, 0xf5, 0x4d, 0xea, 0x9a, 0xa2, 0xc9, 0x9a, 0xf9, 0xab, 0x90, 0x06, 0xc7, 0xc0, 0x03,
	0x40, 0x08, 0x34, 0x3c, 0x8c, 0xb7, 0x5a, 0x83, 0x27, 0xf0, 0xdf, 0xc4, 0x80, 0x8b, 0xc2, 0x34,
	0xd9, 0x53, 0xcc, 0xfa, 0x36, 0xf9, 0xbd, 0x18, 0x65, 0xb5, 0xa7, 0xb8, 0xf4, 0xc8, 0x04, 0xda,
	0x02, 0x89, 0xb5, 0xd6, 0xb8, 0x6e, 0xc8, 0x56, 0xa7, 0x1c, 0x6b, 0x1f, 0xcf, 0xfa, 0x14, 0xd4,
	0xca, 0x2e, 0xcf, 0x14, 0xb7, 0x3f, 0xd9, 0x46, 0x7f, 0x83, 0xff, 0xc2, 0x80, 0x8c, 0x40, 0x2e,
	0x74, 0x85, 0x02, 0x04, 0xb4, 0xf4, 0xc8, 0x15, 0x74, 0x0b, 0x42, 0x69, 0xc2, 0x8e, 0x00, 0xf9,
	0x7a, 0x67, 0xd3, 0xeb, 0xe7, 0xd3, 0x19, 0x0c, 0x2a, 0xf5, 0x6d, 0x8c, 0x69, 0xb4, 0x8b, 0x91,
	0xdc, 0x1c, 0x9f, 0x4a, 0x6e, 0xc2, 0xdb, 0xc8, 0x96, 0x5a, 0x1e, 0x22, 0x97, 0x74, 0x2b, 0xa6,
	0x64, 0x04, 0x4d, 0xfe, 0xa0, 0x78, 0x37, 0xd9, 0x6a, 0x67, 0x82, 0x87, 0x0c, 0xb0, 0x73, 0xdc,
	0xfa, 0x90, 0xa0, 0x27, 0x1c, 0x1e, 0x9d, 0x9d, 0x17, 0x20, 0x23, 0x33, 0x50, 0xee, 0x19, 0x3a,
	0x09, 0x16, 0x5b, 0x9c, 0x06, 0x0d, 0x2f, 0x4f, 0xb3, 0x0f, 0x75, 0xf5, 0x7f, 0xe4, 0x16, 0xe4,
	0x05, 0x26, 0x85, 0x7c, 0x70, 0xc2, 0xcd, 0xfe, 0x99, 0xef, 0x2d, 0x66, 0xa0, 0xac, 0xa9, 0xf7,
	0x97, 0x12, 0x73, 0x50, 0xe6, 0x18, 0x60, 0xc9, 0xe1, 0x17, 0x3d, 0xee, 0x1a, 0x2f, 0x12, 0x75,
	0xdd, 0x16, 0xff, 0xf2, 0xae, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xab, 0xf8, 0x54, 0x6b, 0xb3,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceHandlerClient is the client API for ResourceHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceHandlerClient interface {
	// resource
	CreateResource(ctx context.Context, in *ResourceGroup, opts ...grpc.CallOption) (*ResourceGroupResponse, error)
	GetResource(ctx context.Context, in *ResourceGroupSpec, opts ...grpc.CallOption) (*ResourceGroupResponse, error)
	UpdateResource(ctx context.Context, in *ResourceGroup, opts ...grpc.CallOption) (*ResourceGroupResponse, error)
	DeleteResource(ctx context.Context, in *ResourceGroupSpec, opts ...grpc.CallOption) (*ResourceGroupResponse, error)
}

type resourceHandlerClient struct {
	cc *grpc.ClientConn
}

func NewResourceHandlerClient(cc *grpc.ClientConn) ResourceHandlerClient {
	return &resourceHandlerClient{cc}
}

func (c *resourceHandlerClient) CreateResource(ctx context.Context, in *ResourceGroup, opts ...grpc.CallOption) (*ResourceGroupResponse, error) {
	out := new(ResourceGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceHandler/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceHandlerClient) GetResource(ctx context.Context, in *ResourceGroupSpec, opts ...grpc.CallOption) (*ResourceGroupResponse, error) {
	out := new(ResourceGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceHandler/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceHandlerClient) UpdateResource(ctx context.Context, in *ResourceGroup, opts ...grpc.CallOption) (*ResourceGroupResponse, error) {
	out := new(ResourceGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceHandler/UpdateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceHandlerClient) DeleteResource(ctx context.Context, in *ResourceGroupSpec, opts ...grpc.CallOption) (*ResourceGroupResponse, error) {
	out := new(ResourceGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ResourceHandler/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceHandlerServer is the server API for ResourceHandler service.
type ResourceHandlerServer interface {
	// resource
	CreateResource(context.Context, *ResourceGroup) (*ResourceGroupResponse, error)
	GetResource(context.Context, *ResourceGroupSpec) (*ResourceGroupResponse, error)
	UpdateResource(context.Context, *ResourceGroup) (*ResourceGroupResponse, error)
	DeleteResource(context.Context, *ResourceGroupSpec) (*ResourceGroupResponse, error)
}

func RegisterResourceHandlerServer(s *grpc.Server, srv ResourceHandlerServer) {
	s.RegisterService(&_ResourceHandler_serviceDesc, srv)
}

func _ResourceHandler_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHandlerServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHandler/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHandlerServer).CreateResource(ctx, req.(*ResourceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceHandler_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGroupSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHandlerServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHandler/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHandlerServer).GetResource(ctx, req.(*ResourceGroupSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceHandler_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHandlerServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHandler/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHandlerServer).UpdateResource(ctx, req.(*ResourceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceHandler_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGroupSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHandlerServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHandler/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHandlerServer).DeleteResource(ctx, req.(*ResourceGroupSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ResourceHandler",
	HandlerType: (*ResourceHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _ResourceHandler_CreateResource_Handler,
		},
		{
			MethodName: "GetResource",
			Handler:    _ResourceHandler_GetResource_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _ResourceHandler_UpdateResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _ResourceHandler_DeleteResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}
