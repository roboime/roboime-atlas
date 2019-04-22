// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages_robocup_ssl_refbox_log.proto

package ssl

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

type Log_Frame struct {
	Frame                *SSL_DetectionFrame `protobuf:"bytes,1,req,name=frame" json:"frame,omitempty"`
	RefboxCmd            *string             `protobuf:"bytes,2,req,name=refbox_cmd,json=refboxCmd" json:"refbox_cmd,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Log_Frame) Reset()         { *m = Log_Frame{} }
func (m *Log_Frame) String() string { return proto.CompactTextString(m) }
func (*Log_Frame) ProtoMessage()    {}
func (*Log_Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a10902219dbca77, []int{0}
}

func (m *Log_Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log_Frame.Unmarshal(m, b)
}
func (m *Log_Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log_Frame.Marshal(b, m, deterministic)
}
func (m *Log_Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log_Frame.Merge(m, src)
}
func (m *Log_Frame) XXX_Size() int {
	return xxx_messageInfo_Log_Frame.Size(m)
}
func (m *Log_Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Log_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Log_Frame proto.InternalMessageInfo

func (m *Log_Frame) GetFrame() *SSL_DetectionFrame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *Log_Frame) GetRefboxCmd() string {
	if m != nil && m.RefboxCmd != nil {
		return *m.RefboxCmd
	}
	return ""
}

type Refbox_Log struct {
	Log                  []*Log_Frame `protobuf:"bytes,1,rep,name=log" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Refbox_Log) Reset()         { *m = Refbox_Log{} }
func (m *Refbox_Log) String() string { return proto.CompactTextString(m) }
func (*Refbox_Log) ProtoMessage()    {}
func (*Refbox_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a10902219dbca77, []int{1}
}

func (m *Refbox_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Refbox_Log.Unmarshal(m, b)
}
func (m *Refbox_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Refbox_Log.Marshal(b, m, deterministic)
}
func (m *Refbox_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Refbox_Log.Merge(m, src)
}
func (m *Refbox_Log) XXX_Size() int {
	return xxx_messageInfo_Refbox_Log.Size(m)
}
func (m *Refbox_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Refbox_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Refbox_Log proto.InternalMessageInfo

func (m *Refbox_Log) GetLog() []*Log_Frame {
	if m != nil {
		return m.Log
	}
	return nil
}

func init() {
	proto.RegisterType((*Log_Frame)(nil), "Log_Frame")
	proto.RegisterType((*Refbox_Log)(nil), "Refbox_Log")
}

func init() {
	proto.RegisterFile("messages_robocup_ssl_refbox_log.proto", fileDescriptor_7a10902219dbca77)
}

var fileDescriptor_7a10902219dbca77 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0x8e, 0x2f, 0xca, 0x4f, 0xca, 0x4f, 0x2e, 0x2d, 0x88, 0x2f, 0x2e, 0xce,
	0x89, 0x2f, 0x4a, 0x4d, 0x4b, 0xca, 0xaf, 0x88, 0xcf, 0xc9, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x97, 0x52, 0xc1, 0xaa, 0x2c, 0x25, 0xb5, 0x24, 0x35, 0xb9, 0x24, 0x33, 0x3f, 0x0f, 0xa2,
	0x4a, 0x29, 0x94, 0x8b, 0xd3, 0x27, 0x3f, 0x3d, 0xde, 0xad, 0x28, 0x31, 0x37, 0x55, 0x48, 0x93,
	0x8b, 0x35, 0x0d, 0xc4, 0x90, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x36, 0x12, 0xd6, 0x0b, 0x0e, 0xf6,
	0x89, 0x77, 0x81, 0xe9, 0x00, 0xab, 0x09, 0x82, 0xa8, 0x10, 0x92, 0xe5, 0xe2, 0x82, 0xda, 0x98,
	0x9c, 0x9b, 0x22, 0xc1, 0xa4, 0xc0, 0xa4, 0xc1, 0x19, 0xc4, 0x09, 0x11, 0x71, 0xce, 0x4d, 0x51,
	0xd2, 0xe2, 0xe2, 0x0a, 0x82, 0x48, 0xfb, 0xe4, 0xa7, 0x0b, 0xc9, 0x70, 0x31, 0xe7, 0xe4, 0xa7,
	0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x71, 0xe9, 0xc1, 0x2d, 0x0c, 0x02, 0x09, 0x3b, 0xb1,
	0x46, 0x31, 0x17, 0x17, 0xe7, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x9a, 0xe7, 0x15, 0xd7,
	0x00, 0x00, 0x00,
}