// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

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

type Message_Signal int32

const (
	Message_CREATE Message_Signal = 0
	Message_STOP   Message_Signal = 1
	Message_RELOAD Message_Signal = 2
	Message_OTHER  Message_Signal = 3
)

var Message_Signal_name = map[int32]string{
	0: "CREATE",
	1: "STOP",
	2: "RELOAD",
	3: "OTHER",
}

var Message_Signal_value = map[string]int32{
	"CREATE": 0,
	"STOP":   1,
	"RELOAD": 2,
	"OTHER":  3,
}

func (x Message_Signal) String() string {
	return proto.EnumName(Message_Signal_name, int32(x))
}

func (Message_Signal) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0, 0}
}

type Error_ErrorCode int32

const (
	Error_SUCCESS           Error_ErrorCode = 0
	Error_INVALID_PARAM     Error_ErrorCode = 1
	Error_ACCESS_DENIED     Error_ErrorCode = 2
	Error_THRESHOLD_REACHED Error_ErrorCode = 3
)

var Error_ErrorCode_name = map[int32]string{
	0: "SUCCESS",
	1: "INVALID_PARAM",
	2: "ACCESS_DENIED",
	3: "THRESHOLD_REACHED",
}

var Error_ErrorCode_value = map[string]int32{
	"SUCCESS":           0,
	"INVALID_PARAM":     1,
	"ACCESS_DENIED":     2,
	"THRESHOLD_REACHED": 3,
}

func (x Error_ErrorCode) String() string {
	return proto.EnumName(Error_ErrorCode_name, int32(x))
}

func (Error_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{1, 0}
}

type Message struct {
	AlertConfigId        string         `protobuf:"bytes,1,opt,name=alert_config_id,json=alertConfigId,proto3" json:"alert_config_id,omitempty"`
	Signal               Message_Signal `protobuf:"varint,2,opt,name=signal,proto3,enum=executor.Message_Signal" json:"signal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *Message) GetSignal() Message_Signal {
	if m != nil {
		return m.Signal
	}
	return Message_CREATE
}

// Error
type Error struct {
	Code                 Error_ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=executor.Error_ErrorCode" json:"code,omitempty"`
	Text                 string          `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() Error_ErrorCode {
	if m != nil {
		return m.Code
	}
	return Error_SUCCESS
}

func (m *Error) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterEnum("executor.Message_Signal", Message_Signal_name, Message_Signal_value)
	proto.RegisterEnum("executor.Error_ErrorCode", Error_ErrorCode_name, Error_ErrorCode_value)
	proto.RegisterType((*Message)(nil), "executor.Message")
	proto.RegisterType((*Error)(nil), "executor.Error")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f) }

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x86, 0x5d, 0x8d, 0xd1, 0xcc, 0x87, 0xba, 0x0e, 0x7c, 0x60, 0x7b, 0x92, 0x1c, 0x8a, 0x97,
	0xda, 0x62, 0xe9, 0xa9, 0xa7, 0x6d, 0x32, 0x90, 0x80, 0x1a, 0xd9, 0xc4, 0x1e, 0x7a, 0x09, 0x6a,
	0xb6, 0x22, 0x88, 0x2b, 0x31, 0x05, 0x7f, 0x4b, 0x6f, 0xfd, 0xa7, 0xc5, 0x35, 0xb5, 0xd0, 0x5e,
	0xc2, 0xe4, 0x79, 0x5f, 0x66, 0x1f, 0x76, 0xa1, 0xad, 0x8e, 0x6a, 0xf5, 0x5e, 0xe8, 0x7c, 0xb8,
	0xcf, 0x75, 0xa1, 0xb1, 0xf9, 0xfd, 0xef, 0x7e, 0x30, 0x68, 0x4c, 0xd4, 0xe1, 0xb0, 0x58, 0x2b,
	0xbc, 0x81, 0xce, 0x62, 0xab, 0xf2, 0x22, 0x5d, 0xe9, 0xdd, 0xdb, 0x66, 0x9d, 0x6e, 0xb2, 0x1e,
	0xeb, 0xb3, 0x81, 0x23, 0x5b, 0x06, 0x7b, 0x86, 0x86, 0x19, 0xde, 0x83, 0x7d, 0xd8, 0xac, 0x77,
	0x8b, 0x6d, 0xaf, 0xda, 0x67, 0x83, 0xf6, 0xa8, 0x37, 0xbc, 0xac, 0x2f, 0x57, 0x0d, 0x63, 0x93,
	0xcb, 0xb2, 0xe7, 0x3e, 0x82, 0x7d, 0x26, 0x08, 0x60, 0x7b, 0x92, 0x44, 0x42, 0xbc, 0x82, 0x4d,
	0xb0, 0xe2, 0x24, 0x9a, 0x71, 0x76, 0xa2, 0x92, 0xc6, 0x91, 0xf0, 0x79, 0x15, 0x1d, 0xa8, 0x47,
	0x49, 0x40, 0x92, 0xd7, 0xdc, 0x4f, 0x06, 0x75, 0xca, 0x73, 0x9d, 0xe3, 0x2d, 0x58, 0x2b, 0x9d,
	0x29, 0xe3, 0xd3, 0x1e, 0x5d, 0xfd, 0x1c, 0x68, 0xe2, 0xf3, 0xd7, 0xd3, 0x99, 0x92, 0xa6, 0x86,
	0x08, 0x56, 0xa1, 0x8e, 0x85, 0xf1, 0x73, 0xa4, 0x99, 0xdd, 0x39, 0x38, 0x97, 0x1a, 0xfe, 0x83,
	0x46, 0x3c, 0xf7, 0x3c, 0x8a, 0x63, 0x5e, 0xc1, 0x2e, 0xb4, 0xc2, 0xe9, 0x8b, 0x18, 0x87, 0x7e,
	0x3a, 0x13, 0x52, 0x4c, 0x38, 0x3b, 0x21, 0x61, 0xe2, 0xd4, 0xa7, 0x69, 0x48, 0x27, 0xaf, 0xff,
	0xd0, 0x4d, 0x02, 0x49, 0x71, 0x10, 0x8d, 0xfd, 0x54, 0x92, 0xf0, 0x02, 0xf2, 0x79, 0x6d, 0xf4,
	0x04, 0x4d, 0x2a, 0x65, 0xf0, 0x0e, 0x1a, 0xe7, 0x59, 0x61, 0xf7, 0xcf, 0x9d, 0x5c, 0x77, 0x7e,
	0x59, 0xbb, 0x95, 0x67, 0xeb, 0xb5, 0xba, 0x5f, 0x2e, 0x6d, 0xf3, 0x28, 0x0f, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x53, 0x12, 0xa8, 0xe8, 0xa6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	Execute(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Error, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Execute(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/executor.Executor/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	Execute(context.Context, *Message) (*Error, error)
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executor.Executor/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Execute(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "executor.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Executor_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executor.proto",
}
