// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package fsm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	packet "github.com/vx-labs/mqtt-protocol/packet"
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

type RetainedMessageStored struct {
	Publish              *packet.Publish `protobuf:"bytes,1,opt,name=Publish,proto3" json:"Publish,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RetainedMessageStored) Reset()         { *m = RetainedMessageStored{} }
func (m *RetainedMessageStored) String() string { return proto.CompactTextString(m) }
func (*RetainedMessageStored) ProtoMessage()    {}
func (*RetainedMessageStored) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *RetainedMessageStored) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetainedMessageStored.Unmarshal(m, b)
}
func (m *RetainedMessageStored) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetainedMessageStored.Marshal(b, m, deterministic)
}
func (m *RetainedMessageStored) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetainedMessageStored.Merge(m, src)
}
func (m *RetainedMessageStored) XXX_Size() int {
	return xxx_messageInfo_RetainedMessageStored.Size(m)
}
func (m *RetainedMessageStored) XXX_DiscardUnknown() {
	xxx_messageInfo_RetainedMessageStored.DiscardUnknown(m)
}

var xxx_messageInfo_RetainedMessageStored proto.InternalMessageInfo

func (m *RetainedMessageStored) GetPublish() *packet.Publish {
	if m != nil {
		return m.Publish
	}
	return nil
}

type RetainedMessageDeleted struct {
	Topic                []byte   `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetainedMessageDeleted) Reset()         { *m = RetainedMessageDeleted{} }
func (m *RetainedMessageDeleted) String() string { return proto.CompactTextString(m) }
func (*RetainedMessageDeleted) ProtoMessage()    {}
func (*RetainedMessageDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *RetainedMessageDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetainedMessageDeleted.Unmarshal(m, b)
}
func (m *RetainedMessageDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetainedMessageDeleted.Marshal(b, m, deterministic)
}
func (m *RetainedMessageDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetainedMessageDeleted.Merge(m, src)
}
func (m *RetainedMessageDeleted) XXX_Size() int {
	return xxx_messageInfo_RetainedMessageDeleted.Size(m)
}
func (m *RetainedMessageDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_RetainedMessageDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_RetainedMessageDeleted proto.InternalMessageInfo

func (m *RetainedMessageDeleted) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

type SessionCreated struct {
	SessionID            string          `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	ClientID             string          `protobuf:"bytes,2,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ConnectedAt          int64           `protobuf:"varint,3,opt,name=ConnectedAt,proto3" json:"ConnectedAt,omitempty"`
	Peer                 uint64          `protobuf:"varint,4,opt,name=Peer,proto3" json:"Peer,omitempty"`
	LWT                  *packet.Publish `protobuf:"bytes,5,opt,name=LWT,proto3" json:"LWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SessionCreated) Reset()         { *m = SessionCreated{} }
func (m *SessionCreated) String() string { return proto.CompactTextString(m) }
func (*SessionCreated) ProtoMessage()    {}
func (*SessionCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *SessionCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionCreated.Unmarshal(m, b)
}
func (m *SessionCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionCreated.Marshal(b, m, deterministic)
}
func (m *SessionCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionCreated.Merge(m, src)
}
func (m *SessionCreated) XXX_Size() int {
	return xxx_messageInfo_SessionCreated.Size(m)
}
func (m *SessionCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SessionCreated proto.InternalMessageInfo

func (m *SessionCreated) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SessionCreated) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *SessionCreated) GetConnectedAt() int64 {
	if m != nil {
		return m.ConnectedAt
	}
	return 0
}

func (m *SessionCreated) GetPeer() uint64 {
	if m != nil {
		return m.Peer
	}
	return 0
}

func (m *SessionCreated) GetLWT() *packet.Publish {
	if m != nil {
		return m.LWT
	}
	return nil
}

type SessionDeleted struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Peer                 uint64   `protobuf:"varint,2,opt,name=Peer,proto3" json:"Peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionDeleted) Reset()         { *m = SessionDeleted{} }
func (m *SessionDeleted) String() string { return proto.CompactTextString(m) }
func (*SessionDeleted) ProtoMessage()    {}
func (*SessionDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *SessionDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDeleted.Unmarshal(m, b)
}
func (m *SessionDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDeleted.Marshal(b, m, deterministic)
}
func (m *SessionDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDeleted.Merge(m, src)
}
func (m *SessionDeleted) XXX_Size() int {
	return xxx_messageInfo_SessionDeleted.Size(m)
}
func (m *SessionDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDeleted proto.InternalMessageInfo

func (m *SessionDeleted) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SessionDeleted) GetPeer() uint64 {
	if m != nil {
		return m.Peer
	}
	return 0
}

type SubscriptionCreated struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Peer                 uint64   `protobuf:"varint,2,opt,name=Peer,proto3" json:"Peer,omitempty"`
	Pattern              []byte   `protobuf:"bytes,3,opt,name=Pattern,proto3" json:"Pattern,omitempty"`
	Qos                  int32    `protobuf:"varint,4,opt,name=Qos,proto3" json:"Qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriptionCreated) Reset()         { *m = SubscriptionCreated{} }
func (m *SubscriptionCreated) String() string { return proto.CompactTextString(m) }
func (*SubscriptionCreated) ProtoMessage()    {}
func (*SubscriptionCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}

func (m *SubscriptionCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionCreated.Unmarshal(m, b)
}
func (m *SubscriptionCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionCreated.Marshal(b, m, deterministic)
}
func (m *SubscriptionCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionCreated.Merge(m, src)
}
func (m *SubscriptionCreated) XXX_Size() int {
	return xxx_messageInfo_SubscriptionCreated.Size(m)
}
func (m *SubscriptionCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionCreated proto.InternalMessageInfo

func (m *SubscriptionCreated) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SubscriptionCreated) GetPeer() uint64 {
	if m != nil {
		return m.Peer
	}
	return 0
}

func (m *SubscriptionCreated) GetPattern() []byte {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *SubscriptionCreated) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

type SubscriptionDeleted struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Pattern              []byte   `protobuf:"bytes,2,opt,name=Pattern,proto3" json:"Pattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriptionDeleted) Reset()         { *m = SubscriptionDeleted{} }
func (m *SubscriptionDeleted) String() string { return proto.CompactTextString(m) }
func (*SubscriptionDeleted) ProtoMessage()    {}
func (*SubscriptionDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}

func (m *SubscriptionDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionDeleted.Unmarshal(m, b)
}
func (m *SubscriptionDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionDeleted.Marshal(b, m, deterministic)
}
func (m *SubscriptionDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionDeleted.Merge(m, src)
}
func (m *SubscriptionDeleted) XXX_Size() int {
	return xxx_messageInfo_SubscriptionDeleted.Size(m)
}
func (m *SubscriptionDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionDeleted proto.InternalMessageInfo

func (m *SubscriptionDeleted) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SubscriptionDeleted) GetPattern() []byte {
	if m != nil {
		return m.Pattern
	}
	return nil
}

type PeerLost struct {
	Peer                 uint64   `protobuf:"varint,1,opt,name=Peer,proto3" json:"Peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerLost) Reset()         { *m = PeerLost{} }
func (m *PeerLost) String() string { return proto.CompactTextString(m) }
func (*PeerLost) ProtoMessage()    {}
func (*PeerLost) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{6}
}

func (m *PeerLost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerLost.Unmarshal(m, b)
}
func (m *PeerLost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerLost.Marshal(b, m, deterministic)
}
func (m *PeerLost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerLost.Merge(m, src)
}
func (m *PeerLost) XXX_Size() int {
	return xxx_messageInfo_PeerLost.Size(m)
}
func (m *PeerLost) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerLost.DiscardUnknown(m)
}

var xxx_messageInfo_PeerLost proto.InternalMessageInfo

func (m *PeerLost) GetPeer() uint64 {
	if m != nil {
		return m.Peer
	}
	return 0
}

type StateTransitionSet struct {
	Events               []*StateTransition `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StateTransitionSet) Reset()         { *m = StateTransitionSet{} }
func (m *StateTransitionSet) String() string { return proto.CompactTextString(m) }
func (*StateTransitionSet) ProtoMessage()    {}
func (*StateTransitionSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{7}
}

func (m *StateTransitionSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateTransitionSet.Unmarshal(m, b)
}
func (m *StateTransitionSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateTransitionSet.Marshal(b, m, deterministic)
}
func (m *StateTransitionSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateTransitionSet.Merge(m, src)
}
func (m *StateTransitionSet) XXX_Size() int {
	return xxx_messageInfo_StateTransitionSet.Size(m)
}
func (m *StateTransitionSet) XXX_DiscardUnknown() {
	xxx_messageInfo_StateTransitionSet.DiscardUnknown(m)
}

var xxx_messageInfo_StateTransitionSet proto.InternalMessageInfo

func (m *StateTransitionSet) GetEvents() []*StateTransition {
	if m != nil {
		return m.Events
	}
	return nil
}

type StateTransition struct {
	// Types that are valid to be assigned to Event:
	//	*StateTransition_RetainedMessageStored
	//	*StateTransition_RetainedMessageDeleted
	//	*StateTransition_SessionSubscribed
	//	*StateTransition_SessionUnsubscribed
	//	*StateTransition_PeerLost
	//	*StateTransition_SessionCreated
	//	*StateTransition_SessionDeleted
	Event                isStateTransition_Event `protobuf_oneof:"Event"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StateTransition) Reset()         { *m = StateTransition{} }
func (m *StateTransition) String() string { return proto.CompactTextString(m) }
func (*StateTransition) ProtoMessage()    {}
func (*StateTransition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{8}
}

func (m *StateTransition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateTransition.Unmarshal(m, b)
}
func (m *StateTransition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateTransition.Marshal(b, m, deterministic)
}
func (m *StateTransition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateTransition.Merge(m, src)
}
func (m *StateTransition) XXX_Size() int {
	return xxx_messageInfo_StateTransition.Size(m)
}
func (m *StateTransition) XXX_DiscardUnknown() {
	xxx_messageInfo_StateTransition.DiscardUnknown(m)
}

var xxx_messageInfo_StateTransition proto.InternalMessageInfo

type isStateTransition_Event interface {
	isStateTransition_Event()
}

type StateTransition_RetainedMessageStored struct {
	RetainedMessageStored *RetainedMessageStored `protobuf:"bytes,1,opt,name=RetainedMessageStored,proto3,oneof"`
}

type StateTransition_RetainedMessageDeleted struct {
	RetainedMessageDeleted *RetainedMessageDeleted `protobuf:"bytes,2,opt,name=RetainedMessageDeleted,proto3,oneof"`
}

type StateTransition_SessionSubscribed struct {
	SessionSubscribed *SubscriptionCreated `protobuf:"bytes,3,opt,name=SessionSubscribed,proto3,oneof"`
}

type StateTransition_SessionUnsubscribed struct {
	SessionUnsubscribed *SubscriptionDeleted `protobuf:"bytes,4,opt,name=SessionUnsubscribed,proto3,oneof"`
}

type StateTransition_PeerLost struct {
	PeerLost *PeerLost `protobuf:"bytes,5,opt,name=PeerLost,proto3,oneof"`
}

type StateTransition_SessionCreated struct {
	SessionCreated *SessionCreated `protobuf:"bytes,6,opt,name=SessionCreated,proto3,oneof"`
}

type StateTransition_SessionDeleted struct {
	SessionDeleted *SessionDeleted `protobuf:"bytes,7,opt,name=SessionDeleted,proto3,oneof"`
}

func (*StateTransition_RetainedMessageStored) isStateTransition_Event() {}

func (*StateTransition_RetainedMessageDeleted) isStateTransition_Event() {}

func (*StateTransition_SessionSubscribed) isStateTransition_Event() {}

func (*StateTransition_SessionUnsubscribed) isStateTransition_Event() {}

func (*StateTransition_PeerLost) isStateTransition_Event() {}

func (*StateTransition_SessionCreated) isStateTransition_Event() {}

func (*StateTransition_SessionDeleted) isStateTransition_Event() {}

func (m *StateTransition) GetEvent() isStateTransition_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StateTransition) GetRetainedMessageStored() *RetainedMessageStored {
	if x, ok := m.GetEvent().(*StateTransition_RetainedMessageStored); ok {
		return x.RetainedMessageStored
	}
	return nil
}

func (m *StateTransition) GetRetainedMessageDeleted() *RetainedMessageDeleted {
	if x, ok := m.GetEvent().(*StateTransition_RetainedMessageDeleted); ok {
		return x.RetainedMessageDeleted
	}
	return nil
}

func (m *StateTransition) GetSessionSubscribed() *SubscriptionCreated {
	if x, ok := m.GetEvent().(*StateTransition_SessionSubscribed); ok {
		return x.SessionSubscribed
	}
	return nil
}

func (m *StateTransition) GetSessionUnsubscribed() *SubscriptionDeleted {
	if x, ok := m.GetEvent().(*StateTransition_SessionUnsubscribed); ok {
		return x.SessionUnsubscribed
	}
	return nil
}

func (m *StateTransition) GetPeerLost() *PeerLost {
	if x, ok := m.GetEvent().(*StateTransition_PeerLost); ok {
		return x.PeerLost
	}
	return nil
}

func (m *StateTransition) GetSessionCreated() *SessionCreated {
	if x, ok := m.GetEvent().(*StateTransition_SessionCreated); ok {
		return x.SessionCreated
	}
	return nil
}

func (m *StateTransition) GetSessionDeleted() *SessionDeleted {
	if x, ok := m.GetEvent().(*StateTransition_SessionDeleted); ok {
		return x.SessionDeleted
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StateTransition) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StateTransition_RetainedMessageStored)(nil),
		(*StateTransition_RetainedMessageDeleted)(nil),
		(*StateTransition_SessionSubscribed)(nil),
		(*StateTransition_SessionUnsubscribed)(nil),
		(*StateTransition_PeerLost)(nil),
		(*StateTransition_SessionCreated)(nil),
		(*StateTransition_SessionDeleted)(nil),
	}
}

func init() {
	proto.RegisterType((*RetainedMessageStored)(nil), "fsm.RetainedMessageStored")
	proto.RegisterType((*RetainedMessageDeleted)(nil), "fsm.RetainedMessageDeleted")
	proto.RegisterType((*SessionCreated)(nil), "fsm.SessionCreated")
	proto.RegisterType((*SessionDeleted)(nil), "fsm.SessionDeleted")
	proto.RegisterType((*SubscriptionCreated)(nil), "fsm.SubscriptionCreated")
	proto.RegisterType((*SubscriptionDeleted)(nil), "fsm.SubscriptionDeleted")
	proto.RegisterType((*PeerLost)(nil), "fsm.PeerLost")
	proto.RegisterType((*StateTransitionSet)(nil), "fsm.StateTransitionSet")
	proto.RegisterType((*StateTransition)(nil), "fsm.StateTransition")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xe3, 0x7c, 0xb4, 0xe3, 0x42, 0x61, 0x53, 0xd0, 0x2a, 0x20, 0x14, 0x7c, 0x0a, 0x82,
	0x3a, 0x28, 0x9c, 0x7b, 0x20, 0x09, 0x92, 0x2b, 0xa5, 0x52, 0xd8, 0xa4, 0xe2, 0x6c, 0x3b, 0xd3,
	0xd6, 0xc2, 0xf1, 0x1a, 0xef, 0xa6, 0x82, 0x7f, 0xc0, 0xff, 0xe0, 0x8f, 0x22, 0xaf, 0xed, 0x3a,
	0x71, 0xec, 0xaa, 0x37, 0xef, 0xcc, 0x9b, 0x37, 0xf3, 0x76, 0xdf, 0x18, 0x0c, 0xf9, 0x27, 0x42,
	0x61, 0x45, 0x31, 0x97, 0x9c, 0xe8, 0x37, 0x62, 0xd3, 0xff, 0x7c, 0xeb, 0xcb, 0xbb, 0xad, 0x6b,
	0x79, 0x7c, 0x33, 0xba, 0xff, 0x7d, 0x1e, 0x38, 0xae, 0x18, 0x6d, 0x7e, 0x49, 0x79, 0xae, 0x30,
	0x1e, 0x0f, 0x46, 0x91, 0xe3, 0xfd, 0x44, 0x39, 0x8a, 0xdc, 0xb4, 0xcc, 0x9c, 0xc0, 0x2b, 0x86,
	0xd2, 0xf1, 0x43, 0x5c, 0x5f, 0xa1, 0x10, 0xce, 0x2d, 0x2e, 0x25, 0x8f, 0x71, 0x4d, 0x3e, 0x40,
	0x77, 0xb1, 0x75, 0x03, 0x5f, 0xdc, 0x51, 0x6d, 0xa0, 0x0d, 0x8d, 0xf1, 0xa9, 0x95, 0xd6, 0x5a,
	0x59, 0x98, 0xe5, 0x79, 0xd3, 0x82, 0xd7, 0x25, 0x8e, 0x19, 0x06, 0x28, 0x71, 0x4d, 0xce, 0xa0,
	0xbd, 0xe2, 0x91, 0xef, 0x29, 0x8a, 0x13, 0x96, 0x1e, 0xcc, 0x7f, 0x1a, 0x3c, 0x5f, 0xa2, 0x10,
	0x3e, 0x0f, 0xa7, 0x31, 0x3a, 0x09, 0xf0, 0x2d, 0x1c, 0x67, 0x91, 0xcb, 0x99, 0x02, 0x1f, 0xb3,
	0x22, 0x40, 0xfa, 0x70, 0x34, 0x0d, 0x7c, 0x0c, 0xe5, 0xe5, 0x8c, 0x36, 0x55, 0xf2, 0xe1, 0x4c,
	0x06, 0x60, 0x4c, 0x79, 0x18, 0xa2, 0x27, 0x71, 0xfd, 0x55, 0x52, 0x7d, 0xa0, 0x0d, 0x75, 0xb6,
	0x1b, 0x22, 0x04, 0x5a, 0x0b, 0xc4, 0x98, 0xb6, 0x06, 0xda, 0xb0, 0xc5, 0xd4, 0x37, 0x79, 0x0f,
	0xfa, 0xfc, 0xc7, 0x8a, 0xb6, 0xab, 0x95, 0x25, 0x39, 0x73, 0xf2, 0x30, 0x64, 0xae, 0xe6, 0xf1,
	0x21, 0xf3, 0x36, 0xcd, 0xa2, 0x8d, 0x29, 0xa0, 0xb7, 0xdc, 0xba, 0xc2, 0x8b, 0xfd, 0x48, 0x3e,
	0x59, 0x6d, 0x05, 0x11, 0xa1, 0xd0, 0x5d, 0x38, 0x52, 0x62, 0x1c, 0x2a, 0x85, 0x27, 0x2c, 0x3f,
	0x92, 0x17, 0xa0, 0x7f, 0xe7, 0x42, 0x89, 0x6b, 0xb3, 0xe4, 0xd3, 0xbc, 0xda, 0x6f, 0xfa, 0xb4,
	0xe9, 0x77, 0x1a, 0x34, 0xf7, 0x1a, 0x98, 0xef, 0xe0, 0x28, 0x19, 0x61, 0xce, 0x45, 0x71, 0x95,
	0xda, 0x8e, 0xc6, 0x09, 0x90, 0xa5, 0x74, 0x24, 0xae, 0x62, 0x27, 0x14, 0x7e, 0xd2, 0x71, 0x89,
	0x92, 0x7c, 0x82, 0x0e, 0xde, 0x63, 0x28, 0x05, 0xd5, 0x06, 0xfa, 0xd0, 0x18, 0x9f, 0x59, 0x37,
	0x62, 0x63, 0x95, 0x80, 0x2c, 0xc3, 0x98, 0x7f, 0x5b, 0x70, 0x5a, 0xca, 0x11, 0x56, 0xe3, 0xcc,
	0xcc, 0x8e, 0x7d, 0x45, 0x58, 0x89, 0xb0, 0x1b, 0xac, 0xc6, 0xd4, 0xd7, 0x75, 0x4e, 0x55, 0xa2,
	0x8d, 0xf1, 0x9b, 0x2a, 0xd2, 0x0c, 0x62, 0x37, 0x58, 0x9d, 0xcd, 0x6d, 0x78, 0x99, 0xdd, 0x64,
	0x76, 0xf1, 0x2e, 0xae, 0xd5, 0x3b, 0x19, 0x63, 0x9a, 0xea, 0x3e, 0x34, 0x81, 0xdd, 0x60, 0x87,
	0x45, 0x64, 0x0e, 0xbd, 0x2c, 0x78, 0x1d, 0x8a, 0x82, 0xab, 0x55, 0xc3, 0x55, 0x8c, 0x56, 0x55,
	0x46, 0x3e, 0x16, 0x4f, 0x97, 0x59, 0xfd, 0x99, 0xa2, 0xc8, 0x83, 0x76, 0x83, 0x15, 0x6f, 0x7b,
	0x51, 0x5e, 0x4a, 0xda, 0x51, 0x25, 0xbd, 0xb4, 0xeb, 0x5e, 0xca, 0x6e, 0xb0, 0xf2, 0x06, 0x5f,
	0x94, 0xd7, 0x85, 0x76, 0x0f, 0xcb, 0x8b, 0x79, 0x4b, 0xe0, 0x49, 0x17, 0xda, 0xdf, 0x12, 0x2f,
	0xb8, 0x1d, 0xf5, 0x5f, 0xfa, 0xf2, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xe5, 0x3b, 0x41, 0xdd,
	0x04, 0x00, 0x00,
}
