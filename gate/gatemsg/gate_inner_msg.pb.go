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

func init() {
	proto.RegisterEnum("gatemsg.MsgTypeGateInnerNode_GateProtocol", MsgTypeGateInnerNode_GateProtocol_name, MsgTypeGateInnerNode_GateProtocol_value)
	proto.RegisterType((*MsgTypeGateInnerNode)(nil), "gatemsg.MsgTypeGateInnerNode")
	proto.RegisterType((*RegServerReq)(nil), "gatemsg.RegServerReq")
	proto.RegisterType((*RegServerResp)(nil), "gatemsg.RegServerResp")
	proto.RegisterType((*ServerBeatReq)(nil), "gatemsg.ServerBeatReq")
	proto.RegisterType((*ServerBeatResp)(nil), "gatemsg.ServerBeatResp")
}

func init() { proto.RegisterFile("gate_inner_msg.proto", fileDescriptor_139cee9231f95161) }

var fileDescriptor_139cee9231f95161 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdb, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xbf, 0x34, 0x3d, 0x4e, 0x0f, 0x5f, 0x98, 0xaf, 0x1f, 0x44, 0x44, 0x2c, 0xf1, 0xa6,
	0x78, 0xe1, 0x8d, 0x6f, 0xd0, 0x22, 0x12, 0xd1, 0x22, 0xb1, 0x5e, 0x87, 0x68, 0x86, 0xb0, 0x90,
	0xc3, 0x76, 0x77, 0x91, 0xf6, 0x21, 0x04, 0xcf, 0xb7, 0x3e, 0x86, 0xaf, 0xe1, 0x13, 0xe4, 0x59,
	0x64, 0xb3, 0x0a, 0x6d, 0x2f, 0xbc, 0xdb, 0xf9, 0xcd, 0x8f, 0x61, 0xe6, 0xcf, 0xc2, 0x30, 0x89,
	0x14, 0x85, 0x2c, 0xcf, 0x49, 0x84, 0x99, 0x4c, 0x8e, 0xb8, 0x28, 0x54, 0x81, 0x2d, 0x4d, 0x33,
	0x99, 0x78, 0x9f, 0x16, 0x0c, 0x2f, 0x64, 0x32, 0x5f, 0x71, 0x3a, 0x8d, 0x14, 0xf9, 0xda, 0x9b,
	0x15, 0x31, 0x79, 0x1f, 0x16, 0xf4, 0x34, 0xb9, 0xd4, 0xfe, 0x6d, 0x91, 0x22, 0x40, 0x73, 0x56,
	0xa8, 0x6b, 0x49, 0xce, 0x1f, 0xec, 0x43, 0xcb, 0x8f, 0x27, 0x94, 0xb0, 0xdc, 0x79, 0x28, 0x6d,
	0x44, 0xe8, 0x05, 0x94, 0x5c, 0x91, 0xb8, 0x23, 0x11, 0xd0, 0xc2, 0x79, 0x2c, 0x6d, 0xfc, 0x07,
	0xfd, 0x35, 0x26, 0xb9, 0xf3, 0x64, 0xa0, 0x21, 0x13, 0x8a, 0x94, 0x36, 0x9f, 0x4b, 0x1b, 0x87,
	0x30, 0x58, 0x87, 0x92, 0x3b, 0x2f, 0x86, 0x4e, 0x53, 0x46, 0xb9, 0x9a, 0xa6, 0x85, 0x24, 0xed,
	0xbe, 0x96, 0x36, 0xfe, 0x87, 0xbf, 0x1b, 0x54, 0x72, 0xe7, 0xad, 0xb4, 0xb1, 0x0b, 0x0d, 0x3f,
	0x3e, 0xc9, 0x63, 0xe7, 0xbd, 0xb4, 0xbd, 0x7b, 0x6b, 0x73, 0x1d, 0xdc, 0x87, 0xae, 0xac, 0x8a,
	0x50, 0xad, 0x38, 0xb9, 0xd6, 0xc8, 0x1a, 0x37, 0x02, 0x30, 0x48, 0x1f, 0x8e, 0xbb, 0xd0, 0xf9,
	0x16, 0x58, 0xec, 0xd6, 0xaa, 0x76, 0xdb, 0x00, 0x3f, 0xc6, 0x01, 0xd4, 0x18, 0x77, 0xed, 0x91,
	0x35, 0xee, 0x04, 0x35, 0xc6, 0x11, 0xa1, 0xce, 0x0b, 0xa1, 0xdc, 0x7a, 0xe5, 0x55, 0x6f, 0x3d,
	0x20, 0x8b, 0x96, 0x61, 0xca, 0x32, 0xa6, 0xdc, 0x86, 0x19, 0x90, 0x45, 0xcb, 0x73, 0x5d, 0x7b,
	0x07, 0x5b, 0x49, 0xe8, 0x09, 0x8a, 0x65, 0x3f, 0x8b, 0x54, 0x6f, 0xef, 0x70, 0x2b, 0x19, 0xdc,
	0x81, 0xb6, 0xa0, 0x45, 0xb8, 0x26, 0xb6, 0x04, 0x2d, 0xe6, 0xda, 0x3d, 0xdb, 0x0e, 0xec, 0x17,
	0x19, 0xf7, 0x00, 0x04, 0xf1, 0x74, 0x65, 0x9a, 0xe6, 0xb8, 0x4e, 0x45, 0x74, 0xfb, 0xa6, 0x59,
	0xfd, 0x87, 0xe3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xda, 0x07, 0x93, 0x27, 0x02, 0x00,
	0x00,
}
