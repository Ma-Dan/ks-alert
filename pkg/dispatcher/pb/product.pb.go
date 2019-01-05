// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

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

type Product struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,proto3" json:"product_id,omitempty"`
	ProductName          string   `protobuf:"bytes,2,opt,name=product_name,proto3" json:"product_name,omitempty"`
	EnterpriseId         string   `protobuf:"bytes,3,opt,name=enterprise_id,proto3" json:"enterprise_id,omitempty"`
	EnterpriseName       string   `protobuf:"bytes,4,opt,name=enterprise_name,proto3" json:"enterprise_name,omitempty"`
	HomePage             string   `protobuf:"bytes,5,opt,name=home_page,proto3" json:"home_page,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Phone                string   `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	MonitorCenterHost    string   `protobuf:"bytes,9,opt,name=monitor_center_host,proto3" json:"monitor_center_host,omitempty"`
	MonitorCenterPort    int32    `protobuf:"varint,10,opt,name=monitor_center_port,proto3" json:"monitor_center_port,omitempty"`
	Contacts             string   `protobuf:"bytes,11,opt,name=contacts,proto3" json:"contacts,omitempty"`
	Desc                 string   `protobuf:"bytes,12,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *Product) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *Product) GetEnterpriseId() string {
	if m != nil {
		return m.EnterpriseId
	}
	return ""
}

func (m *Product) GetEnterpriseName() string {
	if m != nil {
		return m.EnterpriseName
	}
	return ""
}

func (m *Product) GetHomePage() string {
	if m != nil {
		return m.HomePage
	}
	return ""
}

func (m *Product) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Product) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Product) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Product) GetMonitorCenterHost() string {
	if m != nil {
		return m.MonitorCenterHost
	}
	return ""
}

func (m *Product) GetMonitorCenterPort() int32 {
	if m != nil {
		return m.MonitorCenterPort
	}
	return 0
}

func (m *Product) GetContacts() string {
	if m != nil {
		return m.Contacts
	}
	return ""
}

func (m *Product) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ProductSpec struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,proto3" json:"product_id,omitempty"`
	ProductName          string   `protobuf:"bytes,2,opt,name=product_name,proto3" json:"product_name,omitempty"`
	EnterpriseId         string   `protobuf:"bytes,3,opt,name=enterprise_id,proto3" json:"enterprise_id,omitempty"`
	EnterpriseName       string   `protobuf:"bytes,4,opt,name=enterprise_name,proto3" json:"enterprise_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSpec) Reset()         { *m = ProductSpec{} }
func (m *ProductSpec) String() string { return proto.CompactTextString(m) }
func (*ProductSpec) ProtoMessage()    {}
func (*ProductSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *ProductSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSpec.Unmarshal(m, b)
}
func (m *ProductSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSpec.Marshal(b, m, deterministic)
}
func (m *ProductSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSpec.Merge(m, src)
}
func (m *ProductSpec) XXX_Size() int {
	return xxx_messageInfo_ProductSpec.Size(m)
}
func (m *ProductSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSpec proto.InternalMessageInfo

func (m *ProductSpec) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *ProductSpec) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductSpec) GetEnterpriseId() string {
	if m != nil {
		return m.EnterpriseId
	}
	return ""
}

func (m *ProductSpec) GetEnterpriseName() string {
	if m != nil {
		return m.EnterpriseName
	}
	return ""
}

type ProductResponse struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *ProductResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "pb.Product")
	proto.RegisterType((*ProductSpec)(nil), "pb.ProductSpec")
	proto.RegisterType((*ProductResponse)(nil), "pb.ProductResponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x49, 0x69, 0x9a, 0xe6, 0xd2, 0x10, 0x70, 0x41, 0x58, 0x85, 0xa1, 0xca, 0xd4, 0xa9,
	0x43, 0x8b, 0xc4, 0xce, 0x1f, 0x01, 0x1b, 0x02, 0xb1, 0xb0, 0x44, 0x4e, 0x7c, 0x6a, 0x23, 0x25,
	0xb6, 0x65, 0x9b, 0x87, 0xe6, 0x29, 0x40, 0x71, 0x13, 0x91, 0xa5, 0x43, 0x07, 0xc6, 0xfb, 0xee,
	0x77, 0xfe, 0xa4, 0x3b, 0x43, 0xac, 0xb4, 0xe4, 0x5f, 0x85, 0x5d, 0x2a, 0x2d, 0xad, 0x24, 0x03,
	0x95, 0xcf, 0x22, 0xd4, 0x5a, 0xea, 0x1d, 0x48, 0x7f, 0x3c, 0x08, 0x5e, 0x77, 0x11, 0x42, 0x00,
	0xda, 0x74, 0x56, 0x72, 0xea, 0xcd, 0xbd, 0x45, 0x48, 0xce, 0x61, 0xd2, 0x31, 0xc1, 0x6a, 0xa4,
	0x03, 0x47, 0x2f, 0x20, 0x46, 0x61, 0x51, 0x2b, 0x5d, 0x1a, 0x6c, 0xc2, 0xc7, 0x0e, 0x5f, 0x42,
	0xd2, 0xc3, 0x2e, 0x3f, 0x74, 0x8d, 0x33, 0x08, 0xb7, 0xb2, 0xc6, 0x4c, 0xb1, 0x0d, 0x52, 0xdf,
	0xa1, 0x04, 0x02, 0xc6, 0xb9, 0x46, 0x63, 0xe8, 0xc8, 0x81, 0x18, 0x7c, 0xb5, 0x95, 0x02, 0x69,
	0xd0, 0x95, 0x58, 0xb3, 0xb2, 0xa2, 0x63, 0x57, 0x5e, 0xc1, 0xb4, 0x96, 0xa2, 0xb4, 0x52, 0x67,
	0x85, 0x73, 0x64, 0x5b, 0x69, 0x2c, 0x0d, 0xf7, 0x34, 0x95, 0xd4, 0x96, 0xc2, 0xdc, 0x5b, 0xf8,
	0xe4, 0x14, 0xc6, 0x85, 0x14, 0x96, 0x15, 0xd6, 0xd0, 0xc8, 0xc5, 0x27, 0x30, 0xe4, 0x68, 0x0a,
	0x3a, 0x69, 0xaa, 0x74, 0x03, 0x51, 0xbb, 0x80, 0x77, 0x85, 0xc5, 0xff, 0x2d, 0x21, 0x7d, 0x81,
	0xa4, 0x15, 0xbd, 0xa1, 0x51, 0x52, 0x18, 0x24, 0xd7, 0x10, 0xb4, 0x0f, 0x3b, 0x53, 0xb4, 0x8a,
	0x96, 0x2a, 0x5f, 0x76, 0xf7, 0xa0, 0xe0, 0xbb, 0x53, 0x39, 0x5f, 0xb4, 0x0a, 0x9b, 0xde, 0x63,
	0x03, 0x56, 0xdf, 0x1e, 0x9c, 0xb4, 0xa9, 0x67, 0x26, 0x78, 0x85, 0x9a, 0xac, 0x21, 0xbe, 0xd7,
	0xc8, 0x2c, 0x76, 0xd3, 0xfd, 0xa7, 0x66, 0xd3, 0x5e, 0xd1, 0xd9, 0xd3, 0x23, 0x72, 0x0b, 0xf1,
	0x03, 0x56, 0xf8, 0x37, 0x94, 0xf4, 0x72, 0xcd, 0x3a, 0xf6, 0x0d, 0xae, 0x21, 0xfe, 0x50, 0xfc,
	0x40, 0xdb, 0x0d, 0xc0, 0x13, 0xda, 0x03, 0x55, 0x77, 0xc3, 0xcf, 0x81, 0xca, 0xf3, 0x91, 0xfb,
	0xae, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x35, 0x5c, 0xe9, 0xd0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductHandlerClient is the client API for ProductHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductHandlerClient interface {
	// product
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error)
	GetProduct(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productHandlerClient struct {
	cc *grpc.ClientConn
}

func NewProductHandlerClient(cc *grpc.ClientConn) ProductHandlerClient {
	return &productHandlerClient{cc}
}

func (c *productHandlerClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductHandler/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHandlerClient) DeleteProduct(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductHandler/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHandlerClient) UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductHandler/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHandlerClient) GetProduct(ctx context.Context, in *ProductSpec, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductHandler/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductHandlerServer is the server API for ProductHandler service.
type ProductHandlerServer interface {
	// product
	CreateProduct(context.Context, *Product) (*ProductResponse, error)
	DeleteProduct(context.Context, *ProductSpec) (*ProductResponse, error)
	UpdateProduct(context.Context, *Product) (*ProductResponse, error)
	GetProduct(context.Context, *ProductSpec) (*ProductResponse, error)
}

func RegisterProductHandlerServer(s *grpc.Server, srv ProductHandlerServer) {
	s.RegisterService(&_ProductHandler_serviceDesc, srv)
}

func _ProductHandler_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHandlerServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductHandler/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHandlerServer).CreateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductHandler_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHandlerServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductHandler/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHandlerServer).DeleteProduct(ctx, req.(*ProductSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductHandler_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHandlerServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductHandler/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHandlerServer).UpdateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductHandler_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHandlerServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductHandler/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHandlerServer).GetProduct(ctx, req.(*ProductSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductHandler",
	HandlerType: (*ProductHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductHandler_CreateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductHandler_DeleteProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductHandler_UpdateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductHandler_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
