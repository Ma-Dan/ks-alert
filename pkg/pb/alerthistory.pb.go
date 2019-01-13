// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alerthistory.proto

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

// alert history
type AlertHistory struct {
	AlertHistoryId            string       `protobuf:"bytes,1,opt,name=alert_history_id,json=alertHistoryId,proto3" json:"alert_history_id,omitempty"`
	AlertConfigId             string       `protobuf:"bytes,2,opt,name=alert_config_id,json=alertConfigId,proto3" json:"alert_config_id,omitempty"`
	AlertName                 string       `protobuf:"bytes,3,opt,name=alert_name,json=alertName,proto3" json:"alert_name,omitempty"`
	SeverityId                string       `protobuf:"bytes,4,opt,name=severity_id,json=severityId,proto3" json:"severity_id,omitempty"`
	SeverityCh                string       `protobuf:"bytes,5,opt,name=severity_ch,json=severityCh,proto3" json:"severity_ch,omitempty"`
	ResourceGroupId           string       `protobuf:"bytes,6,opt,name=resource_group_id,json=resourceGroupId,proto3" json:"resource_group_id,omitempty"`
	ResourceGroupName         string       `protobuf:"bytes,7,opt,name=resource_group_name,json=resourceGroupName,proto3" json:"resource_group_name,omitempty"`
	ResourceTypeId            string       `protobuf:"bytes,8,opt,name=resource_type_id,json=resourceTypeId,proto3" json:"resource_type_id,omitempty"`
	ResourceType              string       `protobuf:"bytes,9,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	AlertedResource           string       `protobuf:"bytes,10,opt,name=alerted_resource,json=alertedResource,proto3" json:"alerted_resource,omitempty"`
	ReceiverGroupId           string       `protobuf:"bytes,11,opt,name=receiver_group_id,json=receiverGroupId,proto3" json:"receiver_group_id,omitempty"`
	ReceiverGroupName         string       `protobuf:"bytes,12,opt,name=receiver_group_name,json=receiverGroupName,proto3" json:"receiver_group_name,omitempty"`
	Receivers                 []*Receiver  `protobuf:"bytes,13,rep,name=receivers,proto3" json:"receivers,omitempty"`
	AlertRuleGroupId          string       `protobuf:"bytes,14,opt,name=alert_rule_group_id,json=alertRuleGroupId,proto3" json:"alert_rule_group_id,omitempty"`
	TriggerAlertRule          string       `protobuf:"bytes,15,opt,name=trigger_alert_rule,json=triggerAlertRule,proto3" json:"trigger_alert_rule,omitempty"`
	AlertRules                []*AlertRule `protobuf:"bytes,16,rep,name=alert_rules,json=alertRules,proto3" json:"alert_rules,omitempty"`
	RepeatSend                *RepeatSend  `protobuf:"bytes,17,opt,name=repeat_send,json=repeatSend,proto3" json:"repeat_send,omitempty"`
	RequestNotificationStatus string       `protobuf:"bytes,18,opt,name=request_notification_status,json=requestNotificationStatus,proto3" json:"request_notification_status,omitempty"`
	EventTime                 string       `protobuf:"bytes,19,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}     `json:"-"`
	XXX_unrecognized          []byte       `json:"-"`
	XXX_sizecache             int32        `json:"-"`
}

func (m *AlertHistory) Reset()         { *m = AlertHistory{} }
func (m *AlertHistory) String() string { return proto.CompactTextString(m) }
func (*AlertHistory) ProtoMessage()    {}
func (*AlertHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb773c7ee97f8e6, []int{0}
}

func (m *AlertHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertHistory.Unmarshal(m, b)
}
func (m *AlertHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertHistory.Marshal(b, m, deterministic)
}
func (m *AlertHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertHistory.Merge(m, src)
}
func (m *AlertHistory) XXX_Size() int {
	return xxx_messageInfo_AlertHistory.Size(m)
}
func (m *AlertHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertHistory.DiscardUnknown(m)
}

var xxx_messageInfo_AlertHistory proto.InternalMessageInfo

func (m *AlertHistory) GetAlertHistoryId() string {
	if m != nil {
		return m.AlertHistoryId
	}
	return ""
}

func (m *AlertHistory) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *AlertHistory) GetAlertName() string {
	if m != nil {
		return m.AlertName
	}
	return ""
}

func (m *AlertHistory) GetSeverityId() string {
	if m != nil {
		return m.SeverityId
	}
	return ""
}

func (m *AlertHistory) GetSeverityCh() string {
	if m != nil {
		return m.SeverityCh
	}
	return ""
}

func (m *AlertHistory) GetResourceGroupId() string {
	if m != nil {
		return m.ResourceGroupId
	}
	return ""
}

func (m *AlertHistory) GetResourceGroupName() string {
	if m != nil {
		return m.ResourceGroupName
	}
	return ""
}

func (m *AlertHistory) GetResourceTypeId() string {
	if m != nil {
		return m.ResourceTypeId
	}
	return ""
}

func (m *AlertHistory) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *AlertHistory) GetAlertedResource() string {
	if m != nil {
		return m.AlertedResource
	}
	return ""
}

func (m *AlertHistory) GetReceiverGroupId() string {
	if m != nil {
		return m.ReceiverGroupId
	}
	return ""
}

func (m *AlertHistory) GetReceiverGroupName() string {
	if m != nil {
		return m.ReceiverGroupName
	}
	return ""
}

func (m *AlertHistory) GetReceivers() []*Receiver {
	if m != nil {
		return m.Receivers
	}
	return nil
}

func (m *AlertHistory) GetAlertRuleGroupId() string {
	if m != nil {
		return m.AlertRuleGroupId
	}
	return ""
}

func (m *AlertHistory) GetTriggerAlertRule() string {
	if m != nil {
		return m.TriggerAlertRule
	}
	return ""
}

func (m *AlertHistory) GetAlertRules() []*AlertRule {
	if m != nil {
		return m.AlertRules
	}
	return nil
}

func (m *AlertHistory) GetRepeatSend() *RepeatSend {
	if m != nil {
		return m.RepeatSend
	}
	return nil
}

func (m *AlertHistory) GetRequestNotificationStatus() string {
	if m != nil {
		return m.RequestNotificationStatus
	}
	return ""
}

func (m *AlertHistory) GetEventTime() string {
	if m != nil {
		return m.EventTime
	}
	return ""
}

type AlertHistoryRequest struct {
	AlertHistoryId       string   `protobuf:"bytes,1,opt,name=alert_history_id,json=alertHistoryId,proto3" json:"alert_history_id,omitempty"`
	AlertConfigId        string   `protobuf:"bytes,2,opt,name=alert_config_id,json=alertConfigId,proto3" json:"alert_config_id,omitempty"`
	AlertRuleId          string   `protobuf:"bytes,3,opt,name=alert_rule_id,json=alertRuleId,proto3" json:"alert_rule_id,omitempty"`
	ResourceId           string   `protobuf:"bytes,4,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Product              string   `protobuf:"bytes,5,opt,name=product,proto3" json:"product,omitempty"`
	Page                 int32    `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Field                string   `protobuf:"bytes,8,opt,name=field,proto3" json:"field,omitempty"`
	Fuzz                 string   `protobuf:"bytes,9,opt,name=fuzz,proto3" json:"fuzz,omitempty"`
	StartTime            int64    `protobuf:"varint,10,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,11,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertHistoryRequest) Reset()         { *m = AlertHistoryRequest{} }
func (m *AlertHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*AlertHistoryRequest) ProtoMessage()    {}
func (*AlertHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb773c7ee97f8e6, []int{1}
}

func (m *AlertHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertHistoryRequest.Unmarshal(m, b)
}
func (m *AlertHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertHistoryRequest.Marshal(b, m, deterministic)
}
func (m *AlertHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertHistoryRequest.Merge(m, src)
}
func (m *AlertHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_AlertHistoryRequest.Size(m)
}
func (m *AlertHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertHistoryRequest proto.InternalMessageInfo

func (m *AlertHistoryRequest) GetAlertHistoryId() string {
	if m != nil {
		return m.AlertHistoryId
	}
	return ""
}

func (m *AlertHistoryRequest) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *AlertHistoryRequest) GetAlertRuleId() string {
	if m != nil {
		return m.AlertRuleId
	}
	return ""
}

func (m *AlertHistoryRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *AlertHistoryRequest) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *AlertHistoryRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AlertHistoryRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AlertHistoryRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *AlertHistoryRequest) GetFuzz() string {
	if m != nil {
		return m.Fuzz
	}
	return ""
}

func (m *AlertHistoryRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AlertHistoryRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type AlertHistoryResponse struct {
	AlertHistory         []*AlertHistory `protobuf:"bytes,1,rep,name=alert_history,json=alertHistory,proto3" json:"alert_history,omitempty"`
	Error                []*Error        `protobuf:"bytes,2,rep,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AlertHistoryResponse) Reset()         { *m = AlertHistoryResponse{} }
func (m *AlertHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*AlertHistoryResponse) ProtoMessage()    {}
func (*AlertHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb773c7ee97f8e6, []int{2}
}

func (m *AlertHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertHistoryResponse.Unmarshal(m, b)
}
func (m *AlertHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertHistoryResponse.Marshal(b, m, deterministic)
}
func (m *AlertHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertHistoryResponse.Merge(m, src)
}
func (m *AlertHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_AlertHistoryResponse.Size(m)
}
func (m *AlertHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlertHistoryResponse proto.InternalMessageInfo

func (m *AlertHistoryResponse) GetAlertHistory() []*AlertHistory {
	if m != nil {
		return m.AlertHistory
	}
	return nil
}

func (m *AlertHistoryResponse) GetError() []*Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*AlertHistory)(nil), "pb.AlertHistory")
	proto.RegisterType((*AlertHistoryRequest)(nil), "pb.AlertHistoryRequest")
	proto.RegisterType((*AlertHistoryResponse)(nil), "pb.AlertHistoryResponse")
}

func init() { proto.RegisterFile("alerthistory.proto", fileDescriptor_eeb773c7ee97f8e6) }

var fileDescriptor_eeb773c7ee97f8e6 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x69, 0xb7, 0x6e, 0xeb, 0xa5, 0xbf, 0xe6, 0x4e, 0xc2, 0x1b, 0x42, 0x9b, 0x8a, 0x84,
	0xca, 0x04, 0x45, 0x1a, 0xe2, 0x15, 0x69, 0x4c, 0xb0, 0xf5, 0x65, 0x0f, 0xd9, 0x9e, 0x90, 0x50,
	0x94, 0xd6, 0xd7, 0xce, 0x52, 0x9b, 0x04, 0xc7, 0x99, 0xd4, 0xfd, 0x6d, 0xfc, 0x19, 0xfc, 0x41,
	0xc8, 0x67, 0x3b, 0x4d, 0xd8, 0x33, 0x6f, 0xf1, 0xf7, 0x3e, 0xb6, 0xef, 0xbe, 0xbe, 0x0b, 0xb0,
	0x78, 0x85, 0x4a, 0x3f, 0xc8, 0x5c, 0xa7, 0x6a, 0x33, 0xc9, 0x54, 0xaa, 0x53, 0xd6, 0xcc, 0x66,
	0x27, 0x01, 0x2a, 0x95, 0x2a, 0x2b, 0x9c, 0xf4, 0x14, 0xce, 0x51, 0x3e, 0xa2, 0x5f, 0xf7, 0x69,
	0x93, 0x2a, 0x56, 0xe8, 0x84, 0x81, 0xc2, 0x0c, 0x63, 0x9d, 0x63, 0x22, 0xac, 0x32, 0xfa, 0xbd,
	0x07, 0x9d, 0x4b, 0x43, 0xdd, 0xd8, 0xa3, 0xd9, 0x18, 0x06, 0xb4, 0x2b, 0x72, 0x77, 0x45, 0x52,
	0xf0, 0xc6, 0x59, 0x63, 0xdc, 0x0e, 0x7b, 0x71, 0x85, 0x9b, 0x0a, 0xf6, 0x16, 0xec, 0xf9, 0xd1,
	0x3c, 0x4d, 0x16, 0x72, 0x69, 0xc0, 0x26, 0x81, 0x5d, 0x92, 0xaf, 0x48, 0x9d, 0x0a, 0xf6, 0x1a,
	0xc0, 0x72, 0x49, 0xbc, 0x46, 0xbe, 0x43, 0x48, 0x9b, 0x94, 0xdb, 0x78, 0x8d, 0xec, 0x14, 0x82,
	0x1c, 0x1f, 0x51, 0x49, 0x4d, 0x77, 0xed, 0x52, 0x1c, 0xbc, 0x34, 0x15, 0x35, 0x60, 0xfe, 0xc0,
	0x5b, 0x75, 0xe0, 0xea, 0x81, 0x9d, 0xc3, 0xa1, 0xc2, 0x3c, 0x2d, 0xd4, 0x1c, 0xa3, 0xa5, 0x4a,
	0x8b, 0xcc, 0x9c, 0xb3, 0x47, 0x58, 0xdf, 0x07, 0xae, 0x8d, 0x3e, 0x15, 0x6c, 0x02, 0xc3, 0x7f,
	0x58, 0xca, 0x6a, 0x9f, 0xe8, 0xc3, 0x1a, 0x4d, 0xd9, 0x8d, 0x61, 0x50, 0xf2, 0x7a, 0x93, 0xa1,
	0x39, 0xfa, 0xc0, 0xda, 0xe1, 0xf5, 0xfb, 0x4d, 0x86, 0x53, 0xc1, 0xde, 0x40, 0xb7, 0x46, 0xf2,
	0x36, 0x61, 0x9d, 0x2a, 0xc6, 0xde, 0x39, 0x77, 0x51, 0x44, 0x5e, 0xe7, 0x60, 0x33, 0x75, 0x7a,
	0xe8, 0x64, 0x5b, 0x95, 0x7d, 0xce, 0x6d, 0x55, 0x81, 0xaf, 0xca, 0x06, 0x6a, 0x55, 0xd5, 0x58,
	0xaa, 0xaa, 0xe3, 0xab, 0xaa, 0xd0, 0x54, 0xd5, 0x39, 0xb4, 0xbd, 0x98, 0xf3, 0xee, 0xd9, 0xce,
	0x38, 0xb8, 0xe8, 0x4c, 0xb2, 0xd9, 0x24, 0x74, 0x62, 0xb8, 0x0d, 0xb3, 0x0f, 0x30, 0xb4, 0xcf,
	0x67, 0xfa, 0x68, 0x9b, 0x49, 0x8f, 0xce, 0xb6, 0xd5, 0x84, 0xc5, 0xaa, 0x34, 0xf8, 0x3d, 0x30,
	0xad, 0xe4, 0x72, 0x89, 0x2a, 0xda, 0x6e, 0xe3, 0x7d, 0x4b, 0xbb, 0xc8, 0xa5, 0xdf, 0xc4, 0x26,
	0x10, 0x6c, 0xa9, 0x9c, 0x0f, 0x28, 0x95, 0xae, 0x49, 0xa5, 0x64, 0x42, 0x28, 0xef, 0xc8, 0xd9,
	0x47, 0x08, 0x6c, 0x0b, 0x47, 0xa6, 0x87, 0xf9, 0xe1, 0x59, 0x63, 0x1c, 0x5c, 0xf4, 0x6c, 0xea,
	0x46, 0xbe, 0xc3, 0x44, 0x84, 0xa0, 0xca, 0x6f, 0xf6, 0x05, 0x5e, 0x29, 0xfc, 0x55, 0x60, 0xae,
	0xa3, 0x24, 0xd5, 0x72, 0x21, 0xe7, 0xb1, 0x96, 0x69, 0x12, 0xe5, 0x3a, 0xd6, 0x45, 0xce, 0x19,
	0xe5, 0x75, 0xec, 0x90, 0xdb, 0x0a, 0x71, 0x47, 0x80, 0x69, 0x5e, 0x7c, 0xc4, 0x44, 0x47, 0x5a,
	0xae, 0x91, 0x0f, 0x6d, 0xf3, 0x92, 0x72, 0x2f, 0xd7, 0x38, 0xfa, 0xd3, 0x84, 0x61, 0x75, 0x7c,
	0x42, 0x7b, 0xd0, 0x7f, 0x98, 0xa2, 0x11, 0x74, 0x2b, 0xcf, 0x20, 0x85, 0x1b, 0xa4, 0xa0, 0x34,
	0xc7, 0x4e, 0x4a, 0xd9, 0x82, 0xdb, 0x51, 0xf2, 0xd2, 0x54, 0x30, 0x0e, 0xfb, 0x99, 0x4a, 0x45,
	0x31, 0xd7, 0x6e, 0x8c, 0xfc, 0x92, 0x31, 0xd8, 0xcd, 0xe2, 0x25, 0xd2, 0xd8, 0xb4, 0x42, 0xfa,
	0x66, 0x47, 0xd0, 0x5a, 0xc9, 0xb5, 0xd4, 0x34, 0x1d, 0xad, 0xd0, 0x2e, 0x8c, 0xba, 0x90, 0xb8,
	0xf2, 0x63, 0x60, 0x17, 0x66, 0xff, 0xa2, 0x78, 0x7a, 0x72, 0x4d, 0x4f, 0xdf, 0xc6, 0xbb, 0x5c,
	0xc7, 0xca, 0x79, 0x67, 0xda, 0x7c, 0x27, 0x6c, 0x93, 0x62, 0xbc, 0x63, 0xc7, 0x70, 0x80, 0x89,
	0xb0, 0xc1, 0x80, 0x82, 0xfb, 0x98, 0x08, 0xb2, 0x35, 0x81, 0xa3, 0xba, 0xab, 0x79, 0x96, 0x26,
	0x39, 0xb2, 0xcf, 0xde, 0x04, 0x67, 0x2b, 0x6f, 0x50, 0xc3, 0x0c, 0xca, 0x86, 0xf1, 0x1b, 0x3a,
	0x55, 0x97, 0xd9, 0x29, 0xb4, 0xe8, 0x37, 0xc9, 0x9b, 0x84, 0xb7, 0x0d, 0xfe, 0xcd, 0x08, 0xa1,
	0xd5, 0x2f, 0x7e, 0xd6, 0x5f, 0xf1, 0x26, 0x4e, 0xc4, 0x0a, 0x15, 0xfb, 0x0e, 0xfd, 0x6b, 0xd4,
	0xb5, 0xdf, 0xe3, 0xcb, 0x67, 0x57, 0xd9, 0x17, 0x3f, 0xe1, 0xcf, 0x03, 0x36, 0xe9, 0xd1, 0x8b,
	0xaf, 0xbb, 0x3f, 0x9a, 0xd9, 0x6c, 0xb6, 0x47, 0x7f, 0xdc, 0x4f, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xff, 0x63, 0xdc, 0x23, 0xcb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertHistoryHandlerClient is the client API for AlertHistoryHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertHistoryHandlerClient interface {
	// history
	GetAlertHistory(ctx context.Context, in *AlertHistoryRequest, opts ...grpc.CallOption) (*AlertHistoryResponse, error)
}

type alertHistoryHandlerClient struct {
	cc *grpc.ClientConn
}

func NewAlertHistoryHandlerClient(cc *grpc.ClientConn) AlertHistoryHandlerClient {
	return &alertHistoryHandlerClient{cc}
}

func (c *alertHistoryHandlerClient) GetAlertHistory(ctx context.Context, in *AlertHistoryRequest, opts ...grpc.CallOption) (*AlertHistoryResponse, error) {
	out := new(AlertHistoryResponse)
	err := c.cc.Invoke(ctx, "/pb.AlertHistoryHandler/GetAlertHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertHistoryHandlerServer is the server API for AlertHistoryHandler service.
type AlertHistoryHandlerServer interface {
	// history
	GetAlertHistory(context.Context, *AlertHistoryRequest) (*AlertHistoryResponse, error)
}

func RegisterAlertHistoryHandlerServer(s *grpc.Server, srv AlertHistoryHandlerServer) {
	s.RegisterService(&_AlertHistoryHandler_serviceDesc, srv)
}

func _AlertHistoryHandler_GetAlertHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertHistoryHandlerServer).GetAlertHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlertHistoryHandler/GetAlertHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertHistoryHandlerServer).GetAlertHistory(ctx, req.(*AlertHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlertHistoryHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AlertHistoryHandler",
	HandlerType: (*AlertHistoryHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlertHistory",
			Handler:    _AlertHistoryHandler_GetAlertHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alerthistory.proto",
}