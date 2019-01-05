// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resourceuritmpl.proto

package pb

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

// resource type
type ResourceUriTmpls struct {
	ResourceUriTmpl      []*ResourceUriTmpl `protobuf:"bytes,1,rep,name=resource_uri_tmpl,proto3" json:"resource_uri_tmpl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ResourceUriTmpls) Reset()         { *m = ResourceUriTmpls{} }
func (m *ResourceUriTmpls) String() string { return proto.CompactTextString(m) }
func (*ResourceUriTmpls) ProtoMessage()    {}
func (*ResourceUriTmpls) Descriptor() ([]byte, []int) {
	return fileDescriptor_046a52f93a75bc7a, []int{0}
}

func (m *ResourceUriTmpls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceUriTmpls.Unmarshal(m, b)
}
func (m *ResourceUriTmpls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceUriTmpls.Marshal(b, m, deterministic)
}
func (m *ResourceUriTmpls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceUriTmpls.Merge(m, src)
}
func (m *ResourceUriTmpls) XXX_Size() int {
	return xxx_messageInfo_ResourceUriTmpls.Size(m)
}
func (m *ResourceUriTmpls) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceUriTmpls.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceUriTmpls proto.InternalMessageInfo

func (m *ResourceUriTmpls) GetResourceUriTmpl() []*ResourceUriTmpl {
	if m != nil {
		return m.ResourceUriTmpl
	}
	return nil
}

type ResourceUriTmpl struct {
	UriTmpl              string            `protobuf:"bytes,1,opt,name=uri_tmpl,proto3" json:"uri_tmpl,omitempty"`
	ResourceName         []string          `protobuf:"bytes,2,rep,name=resource_name,proto3" json:"resource_name,omitempty"`
	Params               map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResourceUriTmpl) Reset()         { *m = ResourceUriTmpl{} }
func (m *ResourceUriTmpl) String() string { return proto.CompactTextString(m) }
func (*ResourceUriTmpl) ProtoMessage()    {}
func (*ResourceUriTmpl) Descriptor() ([]byte, []int) {
	return fileDescriptor_046a52f93a75bc7a, []int{1}
}

func (m *ResourceUriTmpl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceUriTmpl.Unmarshal(m, b)
}
func (m *ResourceUriTmpl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceUriTmpl.Marshal(b, m, deterministic)
}
func (m *ResourceUriTmpl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceUriTmpl.Merge(m, src)
}
func (m *ResourceUriTmpl) XXX_Size() int {
	return xxx_messageInfo_ResourceUriTmpl.Size(m)
}
func (m *ResourceUriTmpl) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceUriTmpl.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceUriTmpl proto.InternalMessageInfo

func (m *ResourceUriTmpl) GetUriTmpl() string {
	if m != nil {
		return m.UriTmpl
	}
	return ""
}

func (m *ResourceUriTmpl) GetResourceName() []string {
	if m != nil {
		return m.ResourceName
	}
	return nil
}

func (m *ResourceUriTmpl) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceUriTmpls)(nil), "pb.ResourceUriTmpls")
	proto.RegisterType((*ResourceUriTmpl)(nil), "pb.ResourceUriTmpl")
	proto.RegisterMapType((map[string]string)(nil), "pb.ResourceUriTmpl.ParamsEntry")
}

func init() { proto.RegisterFile("resourceuritmpl.proto", fileDescriptor_046a52f93a75bc7a) }

var fileDescriptor_046a52f93a75bc7a = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0x2d, 0xca, 0x2c, 0xc9, 0x2d, 0xc8, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe2, 0x12, 0x08, 0x82, 0x4a, 0x86, 0x16, 0x65, 0x86,
	0xe4, 0x16, 0xe4, 0x14, 0x0b, 0xe9, 0x71, 0x09, 0xc2, 0x34, 0xc4, 0x97, 0x16, 0x65, 0xc6, 0x83,
	0xb4, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xeb, 0x15, 0x24, 0xe9, 0xa1, 0x69, 0x50,
	0x9a, 0xc5, 0xc8, 0xc5, 0x8f, 0x26, 0x26, 0x24, 0xc0, 0xc5, 0x81, 0xa4, 0x95, 0x51, 0x83, 0x53,
	0x48, 0x94, 0x8b, 0x17, 0x6e, 0x6a, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xa7,
	0x90, 0x3e, 0x17, 0x5b, 0x41, 0x62, 0x51, 0x62, 0x6e, 0xb1, 0x04, 0x33, 0xd8, 0x06, 0x79, 0x2c,
	0x36, 0xe8, 0x05, 0x80, 0x55, 0xb8, 0xe6, 0x95, 0x14, 0x55, 0x4a, 0xe9, 0x72, 0x71, 0x23, 0x71,
	0x85, 0xb8, 0xb9, 0x98, 0xb3, 0x53, 0x2b, 0xa1, 0x76, 0xf0, 0x72, 0xb1, 0x96, 0x25, 0xe6, 0x94,
	0x82, 0xcc, 0x66, 0xd4, 0xe0, 0xb4, 0x62, 0xb2, 0x60, 0x74, 0x62, 0x89, 0x62, 0x2a, 0x48, 0x4a,
	0x62, 0x03, 0xfb, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x62, 0x2c, 0x22, 0x0a, 0x01,
	0x00, 0x00,
}
