// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gate_inner_msg.proto

package gatemsg

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

type MsgTypeGateInnerNode_GateProtocol int32

const (
	MsgTypeGateInnerNode_NotUse          MsgTypeGateInnerNode_GateProtocol = 0
	MsgTypeGateInnerNode_IdBegin         MsgTypeGateInnerNode_GateProtocol = 61200
	MsgTypeGateInnerNode_RegServerReq    MsgTypeGateInnerNode_GateProtocol = 61201
	MsgTypeGateInnerNode_RegServerResp   MsgTypeGateInnerNode_GateProtocol = 61202
	MsgTypeGateInnerNode_ServerBeatReq   MsgTypeGateInnerNode_GateProtocol = 61203
	MsgTypeGateInnerNode_ServerBeatResp  MsgTypeGateInnerNode_GateProtocol = 61204
	MsgTypeGateInnerNode_ClientCloseReq  MsgTypeGateInnerNode_GateProtocol = 61205
	MsgTypeGateInnerNode_ClientCloseResp MsgTypeGateInnerNode_GateProtocol = 61206
	MsgTypeGateInnerNode_KickOffClientNt MsgTypeGateInnerNode_GateProtocol = 61207
	MsgTypeGateInnerNode_IdEnd           MsgTypeGateInnerNode_GateProtocol = 61215
)

var MsgTypeGateInnerNode_GateProtocol_name = map[int32]string{
	0:     "NotUse",
	61200: "IdBegin",
	61201: "RegServerReq",
	61202: "RegServerResp",
	61203: "ServerBeatReq",
	61204: "ServerBeatResp",
	61205: "ClientCloseReq",
	61206: "ClientCloseResp",
	61207: "KickOffClientNt",
	61215: "IdEnd",
}

var MsgTypeGateInnerNode_GateProtocol_value = map[string]int32{
	"NotUse":          0,
	"IdBegin":         61200,
	"RegServerReq":    61201,
	"RegServerResp":   61202,
	"ServerBeatReq":   61203,
	"ServerBeatResp":  61204,
	"ClientCloseReq":  61205,
	"ClientCloseResp": 61206,
	"KickOffClientNt": 61207,
	"IdEnd":           61215,
}

func (x MsgTypeGateInnerNode_GateProtocol) String() string {
	return proto.EnumName(MsgTypeGateInnerNode_GateProtocol_name, int32(x))
}

func (MsgTypeGateInnerNode_GateProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{0, 0}
}

type MsgTypeGateInnerNode struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgTypeGateInnerNode) Reset()         { *m = MsgTypeGateInnerNode{} }
func (m *MsgTypeGateInnerNode) String() string { return proto.CompactTextString(m) }
func (*MsgTypeGateInnerNode) ProtoMessage()    {}
func (*MsgTypeGateInnerNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{0}
}

func (m *MsgTypeGateInnerNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTypeGateInnerNode.Unmarshal(m, b)
}
func (m *MsgTypeGateInnerNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTypeGateInnerNode.Marshal(b, m, deterministic)
}
func (m *MsgTypeGateInnerNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTypeGateInnerNode.Merge(m, src)
}
func (m *MsgTypeGateInnerNode) XXX_Size() int {
	return xxx_messageInfo_MsgTypeGateInnerNode.Size(m)
}
func (m *MsgTypeGateInnerNode) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTypeGateInnerNode.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTypeGateInnerNode proto.InternalMessageInfo

//
type RegServerReq struct {
	ServerType           int32    `protobuf:"varint,1,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
	ServerId             int32    `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	MaxLimit             int32    `protobuf:"varint,5,opt,name=max_limit,json=maxLimit,proto3" json:"max_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegServerReq) Reset()         { *m = RegServerReq{} }
func (m *RegServerReq) String() string { return proto.CompactTextString(m) }
func (*RegServerReq) ProtoMessage()    {}
func (*RegServerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{1}
}

func (m *RegServerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegServerReq.Unmarshal(m, b)
}
func (m *RegServerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegServerReq.Marshal(b, m, deterministic)
}
func (m *RegServerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegServerReq.Merge(m, src)
}
func (m *RegServerReq) XXX_Size() int {
	return xxx_messageInfo_RegServerReq.Size(m)
}
func (m *RegServerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegServerReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegServerReq proto.InternalMessageInfo

func (m *RegServerReq) GetServerType() int32 {
	if m != nil {
		return m.ServerType
	}
	return 0
}

func (m *RegServerReq) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *RegServerReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RegServerReq) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RegServerReq) GetMaxLimit() int32 {
	if m != nil {
		return m.MaxLimit
	}
	return 0
}

//
type RegServerResp struct {
	Time                 int32    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegServerResp) Reset()         { *m = RegServerResp{} }
func (m *RegServerResp) String() string { return proto.CompactTextString(m) }
func (*RegServerResp) ProtoMessage()    {}
func (*RegServerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{2}
}

func (m *RegServerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegServerResp.Unmarshal(m, b)
}
func (m *RegServerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegServerResp.Marshal(b, m, deterministic)
}
func (m *RegServerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegServerResp.Merge(m, src)
}
func (m *RegServerResp) XXX_Size() int {
	return xxx_messageInfo_RegServerResp.Size(m)
}
func (m *RegServerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegServerResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegServerResp proto.InternalMessageInfo

func (m *RegServerResp) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type ServerBeatReq struct {
	ReqTime              int32    `protobuf:"varint,1,opt,name=req_time,json=reqTime,proto3" json:"req_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerBeatReq) Reset()         { *m = ServerBeatReq{} }
func (m *ServerBeatReq) String() string { return proto.CompactTextString(m) }
func (*ServerBeatReq) ProtoMessage()    {}
func (*ServerBeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{3}
}

func (m *ServerBeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerBeatReq.Unmarshal(m, b)
}
func (m *ServerBeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerBeatReq.Marshal(b, m, deterministic)
}
func (m *ServerBeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerBeatReq.Merge(m, src)
}
func (m *ServerBeatReq) XXX_Size() int {
	return xxx_messageInfo_ServerBeatReq.Size(m)
}
func (m *ServerBeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerBeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_ServerBeatReq proto.InternalMessageInfo

func (m *ServerBeatReq) GetReqTime() int32 {
	if m != nil {
		return m.ReqTime
	}
	return 0
}

type ServerBeatResp struct {
	ReqTime              int32    `protobuf:"varint,1,opt,name=req_time,json=reqTime,proto3" json:"req_time,omitempty"`
	ReplyTime            int32    `protobuf:"varint,2,opt,name=reply_time,json=replyTime,proto3" json:"reply_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerBeatResp) Reset()         { *m = ServerBeatResp{} }
func (m *ServerBeatResp) String() string { return proto.CompactTextString(m) }
func (*ServerBeatResp) ProtoMessage()    {}
func (*ServerBeatResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{4}
}

func (m *ServerBeatResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerBeatResp.Unmarshal(m, b)
}
func (m *ServerBeatResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerBeatResp.Marshal(b, m, deterministic)
}
func (m *ServerBeatResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerBeatResp.Merge(m, src)
}
func (m *ServerBeatResp) XXX_Size() int {
	return xxx_messageInfo_ServerBeatResp.Size(m)
}
func (m *ServerBeatResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerBeatResp.DiscardUnknown(m)
}

var xxx_messageInfo_ServerBeatResp proto.InternalMessageInfo

func (m *ServerBeatResp) GetReqTime() int32 {
	if m != nil {
		return m.ReqTime
	}
	return 0
}

func (m *ServerBeatResp) GetReplyTime() int32 {
	if m != nil {
		return m.ReplyTime
	}
	return 0
}

type KickOffClientNt struct {
	SessionId            uint64   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	RepeatedLogin        bool     `protobuf:"varint,3,opt,name=repeated_login,json=repeatedLogin,proto3" json:"repeated_login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickOffClientNt) Reset()         { *m = KickOffClientNt{} }
func (m *KickOffClientNt) String() string { return proto.CompactTextString(m) }
func (*KickOffClientNt) ProtoMessage()    {}
func (*KickOffClientNt) Descriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{5}
}

func (m *KickOffClientNt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickOffClientNt.Unmarshal(m, b)
}
func (m *KickOffClientNt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickOffClientNt.Marshal(b, m, deterministic)
}
func (m *KickOffClientNt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickOffClientNt.Merge(m, src)
}
func (m *KickOffClientNt) XXX_Size() int {
	return xxx_messageInfo_KickOffClientNt.Size(m)
}
func (m *KickOffClientNt) XXX_DiscardUnknown() {
	xxx_messageInfo_KickOffClientNt.DiscardUnknown(m)
}

var xxx_messageInfo_KickOffClientNt proto.InternalMessageInfo

func (m *KickOffClientNt) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *KickOffClientNt) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *KickOffClientNt) GetRepeatedLogin() bool {
	if m != nil {
		return m.RepeatedLogin
	}
	return false
}

type ClientCloseReq struct {
	SessionId            uint64   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCloseReq) Reset()         { *m = ClientCloseReq{} }
func (m *ClientCloseReq) String() string { return proto.CompactTextString(m) }
func (*ClientCloseReq) ProtoMessage()    {}
func (*ClientCloseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_139cee9231f95161, []int{6}
}

func (m *ClientCloseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCloseReq.Unmarshal(m, b)
}
func (m *ClientCloseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCloseReq.Marshal(b, m, deterministic)
}
func (m *ClientCloseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCloseReq.Merge(m, src)
}
func (m *ClientCloseReq) XXX_Size() int {
	return xxx_messageInfo_ClientCloseReq.Size(m)
}
func (m *ClientCloseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCloseReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCloseReq proto.InternalMessageInfo

func (m *ClientCloseReq) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *ClientCloseReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterEnum("gatemsg.MsgTypeGateInnerNode_GateProtocol", MsgTypeGateInnerNode_GateProtocol_name, MsgTypeGateInnerNode_GateProtocol_value)
	proto.RegisterType((*MsgTypeGateInnerNode)(nil), "gatemsg.MsgTypeGateInnerNode")
	proto.RegisterType((*RegServerReq)(nil), "gatemsg.RegServerReq")
	proto.RegisterType((*RegServerResp)(nil), "gatemsg.RegServerResp")
	proto.RegisterType((*ServerBeatReq)(nil), "gatemsg.ServerBeatReq")
	proto.RegisterType((*ServerBeatResp)(nil), "gatemsg.ServerBeatResp")
	proto.RegisterType((*KickOffClientNt)(nil), "gatemsg.KickOffClientNt")
	proto.RegisterType((*ClientCloseReq)(nil), "gatemsg.ClientCloseReq")
}

func init() { proto.RegisterFile("gate_inner_msg.proto", fileDescriptor_139cee9231f95161) }

var fileDescriptor_139cee9231f95161 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0xd3, 0x6e, 0xdb, 0xb3, 0xdb, 0x3a, 0x1c, 0x2b, 0x44, 0x44, 0x5c, 0x22, 0xc2,
	0xe2, 0x85, 0x37, 0x3e, 0x81, 0xbb, 0x88, 0x44, 0xd7, 0x2a, 0x71, 0xbd, 0x0e, 0x71, 0xe7, 0x6c,
	0x18, 0x4c, 0x32, 0xd3, 0x99, 0x51, 0xb6, 0x0f, 0x21, 0xf8, 0xdf, 0x4b, 0xdf, 0xc8, 0x47, 0xc8,
	0xad, 0xaf, 0x21, 0x33, 0xd3, 0x85, 0x6e, 0x04, 0x61, 0xef, 0xce, 0xfc, 0xce, 0xaf, 0x5f, 0x27,
	0x5f, 0x02, 0x8b, 0xaa, 0xb4, 0x54, 0x88, 0xb6, 0x25, 0x5d, 0x34, 0xa6, 0x7a, 0xa8, 0xb4, 0xb4,
	0x12, 0xc7, 0x8e, 0x36, 0xa6, 0x4a, 0xff, 0x44, 0xb0, 0x78, 0x61, 0xaa, 0x93, 0xb5, 0xa2, 0xa7,
	0xa5, 0xa5, 0xcc, 0x79, 0x4b, 0xc9, 0x29, 0xfd, 0x1d, 0xc1, 0x9e, 0x23, 0xaf, 0x9c, 0x7f, 0x2a,
	0x6b, 0x04, 0xd8, 0x59, 0x4a, 0xfb, 0xc6, 0x10, 0xbb, 0x86, 0x33, 0x18, 0x67, 0xfc, 0x90, 0x2a,
	0xd1, 0xb2, 0x4f, 0x5d, 0x8c, 0x08, 0x7b, 0x39, 0x55, 0xaf, 0x49, 0x7f, 0x20, 0x9d, 0xd3, 0x8a,
	0x7d, 0xee, 0x62, 0xbc, 0x01, 0xb3, 0x2d, 0x66, 0x14, 0xfb, 0x12, 0x60, 0x20, 0x87, 0x54, 0x5a,
	0x67, 0x7e, 0xed, 0x62, 0x5c, 0xc0, 0x7c, 0x1b, 0x1a, 0xc5, 0xbe, 0x05, 0x7a, 0x54, 0x0b, 0x6a,
	0xed, 0x51, 0x2d, 0x0d, 0x39, 0xf7, 0x7b, 0x17, 0xe3, 0x4d, 0xb8, 0x7e, 0x89, 0x1a, 0xc5, 0x7e,
	0x04, 0xfc, 0x5c, 0x9c, 0xbe, 0x7b, 0x79, 0x76, 0x16, 0xb6, 0x4b, 0xcb, 0x7e, 0x76, 0x31, 0xee,
	0xc2, 0x28, 0xe3, 0x4f, 0x5a, 0xce, 0x7e, 0x75, 0x71, 0xfa, 0x31, 0xba, 0x7c, 0x4b, 0xbc, 0x0b,
	0xbb, 0xc6, 0x1f, 0x0a, 0xbb, 0x56, 0x94, 0x44, 0xfb, 0xd1, 0xc1, 0x28, 0x87, 0x80, 0x5c, 0x1f,
	0x78, 0x1b, 0xa6, 0x1b, 0x41, 0xf0, 0x64, 0xe0, 0xd7, 0x93, 0x00, 0x32, 0x8e, 0x73, 0x18, 0x08,
	0x95, 0xc4, 0xfb, 0xd1, 0xc1, 0x34, 0x1f, 0x08, 0x85, 0x08, 0x43, 0x25, 0xb5, 0x4d, 0x86, 0xde,
	0xf3, 0xb3, 0x0b, 0x68, 0xca, 0xf3, 0xa2, 0x16, 0x8d, 0xb0, 0xc9, 0x28, 0x04, 0x34, 0xe5, 0xf9,
	0xb1, 0x3b, 0xa7, 0xf7, 0x7a, 0x05, 0xb9, 0x04, 0x2b, 0x9a, 0x8b, 0x8b, 0xf8, 0x39, 0x7d, 0xd0,
	0x2b, 0x0c, 0x6f, 0xc1, 0x44, 0xd3, 0xaa, 0xd8, 0x12, 0xc7, 0x9a, 0x56, 0x27, 0xce, 0x7d, 0xd6,
	0xef, 0xf1, 0x3f, 0x32, 0xde, 0x01, 0xd0, 0xa4, 0xea, 0x75, 0x58, 0x86, 0x87, 0x9b, 0x7a, 0xe2,
	0xb3, 0xc4, 0x3f, 0x85, 0xba, 0x5f, 0x18, 0x32, 0x46, 0xc8, 0xd6, 0xd5, 0xe1, 0xe2, 0x86, 0xf9,
	0x74, 0x43, 0x32, 0x8e, 0x0c, 0xe2, 0xf7, 0x9b, 0x9a, 0x86, 0xb9, 0x1b, 0xf1, 0x3e, 0xcc, 0x35,
	0x29, 0x2a, 0x2d, 0xf1, 0xa2, 0x96, 0x95, 0x68, 0x7d, 0x5b, 0x93, 0x7c, 0x76, 0x41, 0x8f, 0x1d,
	0x4c, 0x1f, 0xf7, 0x5f, 0xf4, 0x95, 0xff, 0xe9, 0xed, 0x8e, 0xff, 0xa8, 0x1f, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x95, 0x4a, 0x58, 0xea, 0xec, 0x02, 0x00, 0x00,
}
