// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages_robocup_ssl_wrapper.proto

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

type SSL_WrapperPacket struct {
	Detection            *SSL_DetectionFrame `protobuf:"bytes,1,opt,name=detection" json:"detection,omitempty"`
	Geometry             *SSL_GeometryData   `protobuf:"bytes,2,opt,name=geometry" json:"geometry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SSL_WrapperPacket) Reset()         { *m = SSL_WrapperPacket{} }
func (m *SSL_WrapperPacket) String() string { return proto.CompactTextString(m) }
func (*SSL_WrapperPacket) ProtoMessage()    {}
func (*SSL_WrapperPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf05487903c52d35, []int{0}
}

func (m *SSL_WrapperPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_WrapperPacket.Unmarshal(m, b)
}
func (m *SSL_WrapperPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_WrapperPacket.Marshal(b, m, deterministic)
}
func (m *SSL_WrapperPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_WrapperPacket.Merge(m, src)
}
func (m *SSL_WrapperPacket) XXX_Size() int {
	return xxx_messageInfo_SSL_WrapperPacket.Size(m)
}
func (m *SSL_WrapperPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_WrapperPacket.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_WrapperPacket proto.InternalMessageInfo

func (m *SSL_WrapperPacket) GetDetection() *SSL_DetectionFrame {
	if m != nil {
		return m.Detection
	}
	return nil
}

func (m *SSL_WrapperPacket) GetGeometry() *SSL_GeometryData {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func init() {
	proto.RegisterType((*SSL_WrapperPacket)(nil), "SSL_WrapperPacket")
}

func init() { proto.RegisterFile("messages_robocup_ssl_wrapper.proto", fileDescriptor_bf05487903c52d35) }

var fileDescriptor_bf05487903c52d35 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0x8e, 0x2f, 0xca, 0x4f, 0xca, 0x4f, 0x2e, 0x2d, 0x88, 0x2f, 0x2e, 0xce,
	0x89, 0x2f, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x52,
	0xc1, 0xaa, 0x26, 0x25, 0xb5, 0x24, 0x35, 0xb9, 0x24, 0x33, 0x3f, 0x0f, 0xaa, 0x4a, 0x19, 0xab,
	0xaa, 0xf4, 0xd4, 0xfc, 0xdc, 0xd4, 0x92, 0xa2, 0x4a, 0x88, 0x22, 0xa5, 0x52, 0x2e, 0xc1, 0xe0,
	0x60, 0x9f, 0xf8, 0x70, 0x88, 0xf9, 0x01, 0x89, 0xc9, 0xd9, 0xa9, 0x25, 0x42, 0x86, 0x5c, 0x9c,
	0x70, 0xc3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0x40, 0xca, 0x5c, 0x60, 0xa2,
	0x6e, 0x45, 0x89, 0xb9, 0xa9, 0x41, 0x08, 0x55, 0x42, 0xba, 0x5c, 0x1c, 0x30, 0x93, 0x25, 0x98,
	0xc0, 0x3a, 0x04, 0xc1, 0x3a, 0xdc, 0xa1, 0x82, 0x2e, 0x89, 0x25, 0x89, 0x41, 0x70, 0x25, 0x4e,
	0xac, 0x51, 0xcc, 0xc5, 0xc5, 0x39, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x09, 0xe7, 0x65,
	0xed, 0x00, 0x00, 0x00,
}