// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sp_notify.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotificationType int32

const (
	NotificationType_NOTIFICATION_TYPE_UNSET     NotificationType = 0
	NotificationType_NOTIFICATION_TYPE_SLACK     NotificationType = 1
	NotificationType_NOTIFICATION_TYPE_EMAIL     NotificationType = 2
	NotificationType_NOTIFICATION_TYPE_PAGERDUTY NotificationType = 3
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "NOTIFICATION_TYPE_UNSET",
		1: "NOTIFICATION_TYPE_SLACK",
		2: "NOTIFICATION_TYPE_EMAIL",
		3: "NOTIFICATION_TYPE_PAGERDUTY",
	}
	NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_UNSET":     0,
		"NOTIFICATION_TYPE_SLACK":     1,
		"NOTIFICATION_TYPE_EMAIL":     2,
		"NOTIFICATION_TYPE_PAGERDUTY": 3,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_notify_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_sp_notify_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{0}
}

type NotificationEmail_Type int32

const (
	NotificationEmail_TYPE_UNSET NotificationEmail_Type = 0
	NotificationEmail_TYPE_SMTP  NotificationEmail_Type = 1
	NotificationEmail_TYPE_SES   NotificationEmail_Type = 2
)

// Enum value maps for NotificationEmail_Type.
var (
	NotificationEmail_Type_name = map[int32]string{
		0: "TYPE_UNSET",
		1: "TYPE_SMTP",
		2: "TYPE_SES",
	}
	NotificationEmail_Type_value = map[string]int32{
		"TYPE_UNSET": 0,
		"TYPE_SMTP":  1,
		"TYPE_SES":   2,
	}
)

func (x NotificationEmail_Type) Enum() *NotificationEmail_Type {
	p := new(NotificationEmail_Type)
	*p = x
	return p
}

func (x NotificationEmail_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationEmail_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_notify_proto_enumTypes[1].Descriptor()
}

func (NotificationEmail_Type) Type() protoreflect.EnumType {
	return &file_sp_notify_proto_enumTypes[1]
}

func (x NotificationEmail_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationEmail_Type.Descriptor instead.
func (NotificationEmail_Type) EnumDescriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{2, 0}
}

type NotificationPagerDuty_Urgency int32

const (
	NotificationPagerDuty_URGENCY_UNSET NotificationPagerDuty_Urgency = 0
	NotificationPagerDuty_URGENCY_LOW   NotificationPagerDuty_Urgency = 1
	NotificationPagerDuty_URGENCY_HIGH  NotificationPagerDuty_Urgency = 2
)

// Enum value maps for NotificationPagerDuty_Urgency.
var (
	NotificationPagerDuty_Urgency_name = map[int32]string{
		0: "URGENCY_UNSET",
		1: "URGENCY_LOW",
		2: "URGENCY_HIGH",
	}
	NotificationPagerDuty_Urgency_value = map[string]int32{
		"URGENCY_UNSET": 0,
		"URGENCY_LOW":   1,
		"URGENCY_HIGH":  2,
	}
)

func (x NotificationPagerDuty_Urgency) Enum() *NotificationPagerDuty_Urgency {
	p := new(NotificationPagerDuty_Urgency)
	*p = x
	return p
}

func (x NotificationPagerDuty_Urgency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationPagerDuty_Urgency) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_notify_proto_enumTypes[2].Descriptor()
}

func (NotificationPagerDuty_Urgency) Type() protoreflect.EnumType {
	return &file_sp_notify_proto_enumTypes[2]
}

func (x NotificationPagerDuty_Urgency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationPagerDuty_Urgency.Descriptor instead.
func (NotificationPagerDuty_Urgency) EnumDescriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{5, 0}
}

type NotificationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *string          `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type NotificationType `protobuf:"varint,3,opt,name=type,proto3,enum=protos.NotificationType" json:"type,omitempty"`
	// Types that are assignable to Config:
	//
	//	*NotificationConfig_Slack
	//	*NotificationConfig_Email
	//	*NotificationConfig_Pagerduty
	Config isNotificationConfig_Config `protobuf_oneof:"config"`
}

func (x *NotificationConfig) Reset() {
	*x = NotificationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationConfig) ProtoMessage() {}

func (x *NotificationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sp_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationConfig.ProtoReflect.Descriptor instead.
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationConfig) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *NotificationConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationConfig) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_NOTIFICATION_TYPE_UNSET
}

func (m *NotificationConfig) GetConfig() isNotificationConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *NotificationConfig) GetSlack() *NotificationSlack {
	if x, ok := x.GetConfig().(*NotificationConfig_Slack); ok {
		return x.Slack
	}
	return nil
}

func (x *NotificationConfig) GetEmail() *NotificationEmail {
	if x, ok := x.GetConfig().(*NotificationConfig_Email); ok {
		return x.Email
	}
	return nil
}

func (x *NotificationConfig) GetPagerduty() *NotificationPagerDuty {
	if x, ok := x.GetConfig().(*NotificationConfig_Pagerduty); ok {
		return x.Pagerduty
	}
	return nil
}

type isNotificationConfig_Config interface {
	isNotificationConfig_Config()
}

type NotificationConfig_Slack struct {
	Slack *NotificationSlack `protobuf:"bytes,1000,opt,name=slack,proto3,oneof"`
}

type NotificationConfig_Email struct {
	Email *NotificationEmail `protobuf:"bytes,1001,opt,name=email,proto3,oneof"`
}

type NotificationConfig_Pagerduty struct {
	Pagerduty *NotificationPagerDuty `protobuf:"bytes,1002,opt,name=pagerduty,proto3,oneof"`
}

func (*NotificationConfig_Slack) isNotificationConfig_Config() {}

func (*NotificationConfig_Email) isNotificationConfig_Config() {}

func (*NotificationConfig_Pagerduty) isNotificationConfig_Config() {}

type NotificationSlack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotToken string `protobuf:"bytes,1,opt,name=bot_token,json=botToken,proto3" json:"bot_token,omitempty"`
	Channel  string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *NotificationSlack) Reset() {
	*x = NotificationSlack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationSlack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationSlack) ProtoMessage() {}

func (x *NotificationSlack) ProtoReflect() protoreflect.Message {
	mi := &file_sp_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationSlack.ProtoReflect.Descriptor instead.
func (*NotificationSlack) Descriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationSlack) GetBotToken() string {
	if x != nil {
		return x.BotToken
	}
	return ""
}

func (x *NotificationSlack) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type NotificationEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        NotificationEmail_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.NotificationEmail_Type" json:"type,omitempty"`
	Recipients  []string               `protobuf:"bytes,2,rep,name=recipients,proto3" json:"recipients,omitempty"`
	FromAddress string                 `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// Types that are assignable to Config:
	//
	//	*NotificationEmail_Smtp
	//	*NotificationEmail_Ses
	Config isNotificationEmail_Config `protobuf_oneof:"config"`
}

func (x *NotificationEmail) Reset() {
	*x = NotificationEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationEmail) ProtoMessage() {}

func (x *NotificationEmail) ProtoReflect() protoreflect.Message {
	mi := &file_sp_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationEmail.ProtoReflect.Descriptor instead.
func (*NotificationEmail) Descriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationEmail) GetType() NotificationEmail_Type {
	if x != nil {
		return x.Type
	}
	return NotificationEmail_TYPE_UNSET
}

func (x *NotificationEmail) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *NotificationEmail) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (m *NotificationEmail) GetConfig() isNotificationEmail_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *NotificationEmail) GetSmtp() *NotificationEmailSMTP {
	if x, ok := x.GetConfig().(*NotificationEmail_Smtp); ok {
		return x.Smtp
	}
	return nil
}

func (x *NotificationEmail) GetSes() *NotificationEmailSES {
	if x, ok := x.GetConfig().(*NotificationEmail_Ses); ok {
		return x.Ses
	}
	return nil
}

type isNotificationEmail_Config interface {
	isNotificationEmail_Config()
}

type NotificationEmail_Smtp struct {
	Smtp *NotificationEmailSMTP `protobuf:"bytes,1000,opt,name=smtp,proto3,oneof"`
}

type NotificationEmail_Ses struct {
	Ses *NotificationEmailSES `protobuf:"bytes,1001,opt,name=ses,proto3,oneof"`
}

func (*NotificationEmail_Smtp) isNotificationEmail_Config() {}

func (*NotificationEmail_Ses) isNotificationEmail_Config() {}

type NotificationEmailSMTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	User     string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	UseTls   bool   `protobuf:"varint,5,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
}

func (x *NotificationEmailSMTP) Reset() {
	*x = NotificationEmailSMTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationEmailSMTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationEmailSMTP) ProtoMessage() {}

func (x *NotificationEmailSMTP) ProtoReflect() protoreflect.Message {
	mi := &file_sp_notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationEmailSMTP.ProtoReflect.Descriptor instead.
func (*NotificationEmailSMTP) Descriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationEmailSMTP) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *NotificationEmailSMTP) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *NotificationEmailSMTP) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *NotificationEmailSMTP) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NotificationEmailSMTP) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

type NotificationEmailSES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SesRegion          string `protobuf:"bytes,1,opt,name=ses_region,json=sesRegion,proto3" json:"ses_region,omitempty"`
	SesAccessKeyId     string `protobuf:"bytes,2,opt,name=ses_access_key_id,json=sesAccessKeyId,proto3" json:"ses_access_key_id,omitempty"`
	SesSecretAccessKey string `protobuf:"bytes,3,opt,name=ses_secret_access_key,json=sesSecretAccessKey,proto3" json:"ses_secret_access_key,omitempty"`
}

func (x *NotificationEmailSES) Reset() {
	*x = NotificationEmailSES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_notify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationEmailSES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationEmailSES) ProtoMessage() {}

func (x *NotificationEmailSES) ProtoReflect() protoreflect.Message {
	mi := &file_sp_notify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationEmailSES.ProtoReflect.Descriptor instead.
func (*NotificationEmailSES) Descriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{4}
}

func (x *NotificationEmailSES) GetSesRegion() string {
	if x != nil {
		return x.SesRegion
	}
	return ""
}

func (x *NotificationEmailSES) GetSesAccessKeyId() string {
	if x != nil {
		return x.SesAccessKeyId
	}
	return ""
}

func (x *NotificationEmailSES) GetSesSecretAccessKey() string {
	if x != nil {
		return x.SesSecretAccessKey
	}
	return ""
}

type NotificationPagerDuty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auth token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Must be a valid email for a PagerDuty user
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Must be a valid PagerDuty service
	ServiceId string                        `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Urgency   NotificationPagerDuty_Urgency `protobuf:"varint,4,opt,name=urgency,proto3,enum=protos.NotificationPagerDuty_Urgency" json:"urgency,omitempty"`
}

func (x *NotificationPagerDuty) Reset() {
	*x = NotificationPagerDuty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_notify_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationPagerDuty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationPagerDuty) ProtoMessage() {}

func (x *NotificationPagerDuty) ProtoReflect() protoreflect.Message {
	mi := &file_sp_notify_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationPagerDuty.ProtoReflect.Descriptor instead.
func (*NotificationPagerDuty) Descriptor() ([]byte, []int) {
	return file_sp_notify_proto_rawDescGZIP(), []int{5}
}

func (x *NotificationPagerDuty) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *NotificationPagerDuty) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NotificationPagerDuty) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *NotificationPagerDuty) GetUrgency() NotificationPagerDuty_Urgency {
	if x != nil {
		return x.Urgency
	}
	return NotificationPagerDuty_URGENCY_UNSET
}

var File_sp_notify_proto protoreflect.FileDescriptor

var file_sp_notify_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xa4, 0x02, 0x0a, 0x12, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6c, 0x61,
	0x63, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x3e, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x18, 0xea, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x22, 0x4a, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xb2, 0x02, 0x0a,
	0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x6d, 0x74,
	0x70, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x4d, 0x54, 0x50, 0x48, 0x00, 0x52, 0x04, 0x73, 0x6d, 0x74, 0x70, 0x12,
	0x31, 0x0a, 0x03, 0x73, 0x65, 0x73, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x45, 0x53, 0x48, 0x00, 0x52, 0x03, 0x73,
	0x65, 0x73, 0x22, 0x33, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x4d, 0x54, 0x50, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x45, 0x53, 0x10, 0x02, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x4d, 0x54, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x73, 0x65, 0x54, 0x6c, 0x73, 0x22, 0x93, 0x01, 0x0a,
	0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x45, 0x53, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x11, 0x73, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x65, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x15, 0x73, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x73, 0x65, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x22, 0xe4, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x2e, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x07, 0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x3f, 0x0a, 0x07, 0x55, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43,
	0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x52, 0x47, 0x45, 0x4e,
	0x43, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x2a, 0x8a, 0x01, 0x0a, 0x10, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4e,
	0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x52,
	0x44, 0x55, 0x54, 0x59, 0x10, 0x03, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73,
	0x6e, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_notify_proto_rawDescOnce sync.Once
	file_sp_notify_proto_rawDescData = file_sp_notify_proto_rawDesc
)

func file_sp_notify_proto_rawDescGZIP() []byte {
	file_sp_notify_proto_rawDescOnce.Do(func() {
		file_sp_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_notify_proto_rawDescData)
	})
	return file_sp_notify_proto_rawDescData
}

var file_sp_notify_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_sp_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sp_notify_proto_goTypes = []interface{}{
	(NotificationType)(0),              // 0: protos.NotificationType
	(NotificationEmail_Type)(0),        // 1: protos.NotificationEmail.Type
	(NotificationPagerDuty_Urgency)(0), // 2: protos.NotificationPagerDuty.Urgency
	(*NotificationConfig)(nil),         // 3: protos.NotificationConfig
	(*NotificationSlack)(nil),          // 4: protos.NotificationSlack
	(*NotificationEmail)(nil),          // 5: protos.NotificationEmail
	(*NotificationEmailSMTP)(nil),      // 6: protos.NotificationEmailSMTP
	(*NotificationEmailSES)(nil),       // 7: protos.NotificationEmailSES
	(*NotificationPagerDuty)(nil),      // 8: protos.NotificationPagerDuty
}
var file_sp_notify_proto_depIdxs = []int32{
	0, // 0: protos.NotificationConfig.type:type_name -> protos.NotificationType
	4, // 1: protos.NotificationConfig.slack:type_name -> protos.NotificationSlack
	5, // 2: protos.NotificationConfig.email:type_name -> protos.NotificationEmail
	8, // 3: protos.NotificationConfig.pagerduty:type_name -> protos.NotificationPagerDuty
	1, // 4: protos.NotificationEmail.type:type_name -> protos.NotificationEmail.Type
	6, // 5: protos.NotificationEmail.smtp:type_name -> protos.NotificationEmailSMTP
	7, // 6: protos.NotificationEmail.ses:type_name -> protos.NotificationEmailSES
	2, // 7: protos.NotificationPagerDuty.urgency:type_name -> protos.NotificationPagerDuty.Urgency
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_sp_notify_proto_init() }
func file_sp_notify_proto_init() {
	if File_sp_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sp_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationSlack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationEmail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationEmailSMTP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_notify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationEmailSES); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_notify_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationPagerDuty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_sp_notify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NotificationConfig_Slack)(nil),
		(*NotificationConfig_Email)(nil),
		(*NotificationConfig_Pagerduty)(nil),
	}
	file_sp_notify_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*NotificationEmail_Smtp)(nil),
		(*NotificationEmail_Ses)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_notify_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_notify_proto_goTypes,
		DependencyIndexes: file_sp_notify_proto_depIdxs,
		EnumInfos:         file_sp_notify_proto_enumTypes,
		MessageInfos:      file_sp_notify_proto_msgTypes,
	}.Build()
	File_sp_notify_proto = out.File
	file_sp_notify_proto_rawDesc = nil
	file_sp_notify_proto_goTypes = nil
	file_sp_notify_proto_depIdxs = nil
}
