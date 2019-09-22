// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages_robocup_ssl_wrapper.proto

package ssl

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FrameRequest struct {
	MatchId              int32    `protobuf:"varint,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrameRequest) Reset()         { *m = FrameRequest{} }
func (m *FrameRequest) String() string { return proto.CompactTextString(m) }
func (*FrameRequest) ProtoMessage()    {}
func (*FrameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf05487903c52d35, []int{0}
}

func (m *FrameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameRequest.Unmarshal(m, b)
}
func (m *FrameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameRequest.Marshal(b, m, deterministic)
}
func (m *FrameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameRequest.Merge(m, src)
}
func (m *FrameRequest) XXX_Size() int {
	return xxx_messageInfo_FrameRequest.Size(m)
}
func (m *FrameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FrameRequest proto.InternalMessageInfo

func (m *FrameRequest) GetMatchId() int32 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

type ActiveMatchesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveMatchesRequest) Reset()         { *m = ActiveMatchesRequest{} }
func (m *ActiveMatchesRequest) String() string { return proto.CompactTextString(m) }
func (*ActiveMatchesRequest) ProtoMessage()    {}
func (*ActiveMatchesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf05487903c52d35, []int{1}
}

func (m *ActiveMatchesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveMatchesRequest.Unmarshal(m, b)
}
func (m *ActiveMatchesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveMatchesRequest.Marshal(b, m, deterministic)
}
func (m *ActiveMatchesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveMatchesRequest.Merge(m, src)
}
func (m *ActiveMatchesRequest) XXX_Size() int {
	return xxx_messageInfo_ActiveMatchesRequest.Size(m)
}
func (m *ActiveMatchesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveMatchesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveMatchesRequest proto.InternalMessageInfo

type MatchInfoRequest struct {
	MatchId              int32    `protobuf:"varint,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchInfoRequest) Reset()         { *m = MatchInfoRequest{} }
func (m *MatchInfoRequest) String() string { return proto.CompactTextString(m) }
func (*MatchInfoRequest) ProtoMessage()    {}
func (*MatchInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf05487903c52d35, []int{2}
}

func (m *MatchInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchInfoRequest.Unmarshal(m, b)
}
func (m *MatchInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchInfoRequest.Marshal(b, m, deterministic)
}
func (m *MatchInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchInfoRequest.Merge(m, src)
}
func (m *MatchInfoRequest) XXX_Size() int {
	return xxx_messageInfo_MatchInfoRequest.Size(m)
}
func (m *MatchInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MatchInfoRequest proto.InternalMessageInfo

func (m *MatchInfoRequest) GetMatchId() int32 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

type MatchesPacket struct {
	Match                []*MatchData `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MatchesPacket) Reset()         { *m = MatchesPacket{} }
func (m *MatchesPacket) String() string { return proto.CompactTextString(m) }
func (*MatchesPacket) ProtoMessage()    {}
func (*MatchesPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf05487903c52d35, []int{3}
}

func (m *MatchesPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchesPacket.Unmarshal(m, b)
}
func (m *MatchesPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchesPacket.Marshal(b, m, deterministic)
}
func (m *MatchesPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchesPacket.Merge(m, src)
}
func (m *MatchesPacket) XXX_Size() int {
	return xxx_messageInfo_MatchesPacket.Size(m)
}
func (m *MatchesPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchesPacket.DiscardUnknown(m)
}

var xxx_messageInfo_MatchesPacket proto.InternalMessageInfo

func (m *MatchesPacket) GetMatch() []*MatchData {
	if m != nil {
		return m.Match
	}
	return nil
}

type MatchData struct {
	MatchId              int32    `protobuf:"varint,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	MatchName            string   `protobuf:"bytes,2,opt,name=match_name,json=matchName,proto3" json:"match_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchData) Reset()         { *m = MatchData{} }
func (m *MatchData) String() string { return proto.CompactTextString(m) }
func (*MatchData) ProtoMessage()    {}
func (*MatchData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf05487903c52d35, []int{4}
}

func (m *MatchData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchData.Unmarshal(m, b)
}
func (m *MatchData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchData.Marshal(b, m, deterministic)
}
func (m *MatchData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchData.Merge(m, src)
}
func (m *MatchData) XXX_Size() int {
	return xxx_messageInfo_MatchData.Size(m)
}
func (m *MatchData) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchData.DiscardUnknown(m)
}

var xxx_messageInfo_MatchData proto.InternalMessageInfo

func (m *MatchData) GetMatchId() int32 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *MatchData) GetMatchName() string {
	if m != nil {
		return m.MatchName
	}
	return ""
}

type SSL_WrapperPacket struct {
	Detection            *SSL_DetectionFrame `protobuf:"bytes,1,opt,name=detection,proto3" json:"detection,omitempty"`
	Geometry             *SSL_GeometryData   `protobuf:"bytes,2,opt,name=geometry,proto3" json:"geometry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SSL_WrapperPacket) Reset()         { *m = SSL_WrapperPacket{} }
func (m *SSL_WrapperPacket) String() string { return proto.CompactTextString(m) }
func (*SSL_WrapperPacket) ProtoMessage()    {}
func (*SSL_WrapperPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf05487903c52d35, []int{5}
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
	proto.RegisterType((*FrameRequest)(nil), "FrameRequest")
	proto.RegisterType((*ActiveMatchesRequest)(nil), "ActiveMatchesRequest")
	proto.RegisterType((*MatchInfoRequest)(nil), "MatchInfoRequest")
	proto.RegisterType((*MatchesPacket)(nil), "MatchesPacket")
	proto.RegisterType((*MatchData)(nil), "MatchData")
	proto.RegisterType((*SSL_WrapperPacket)(nil), "SSL_WrapperPacket")
}

func init() { proto.RegisterFile("messages_robocup_ssl_wrapper.proto", fileDescriptor_bf05487903c52d35) }

var fileDescriptor_bf05487903c52d35 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x45, 0xb4, 0x78, 0x30, 0x15, 0x6c, 0x3f, 0x44, 0x2d, 0x55, 0xb2, 0xb6, 0x3d, 0xd0,
	0x03, 0xab, 0xda, 0x39, 0xe6, 0x44, 0x04, 0xb1, 0x90, 0x42, 0x14, 0x99, 0x43, 0xa4, 0x5c, 0xac,
	0xc5, 0x0c, 0x04, 0x05, 0xb3, 0xce, 0xee, 0x92, 0x28, 0x3f, 0x2b, 0xff, 0x30, 0x62, 0x6d, 0x9c,
	0x90, 0xa0, 0x28, 0xc7, 0x79, 0xef, 0xcd, 0x9b, 0xd1, 0x9b, 0x01, 0x9a, 0xa2, 0x52, 0x7c, 0x81,
	0x2a, 0x96, 0x62, 0x2a, 0x92, 0x4d, 0x16, 0x2b, 0xb5, 0x8a, 0xef, 0x25, 0xcf, 0x32, 0x94, 0x2c,
	0x93, 0x42, 0x0b, 0xf7, 0xef, 0x41, 0xcd, 0x0c, 0x35, 0x26, 0x7a, 0x29, 0xd6, 0x85, 0xea, 0xcf,
	0x41, 0xd5, 0x02, 0x45, 0x8a, 0x5a, 0x3e, 0x14, 0xa2, 0xa6, 0xc4, 0x39, 0x4a, 0xc4, 0xbc, 0xa4,
	0xff, 0xc0, 0x39, 0x95, 0x3c, 0xc5, 0x08, 0x6f, 0x37, 0xa8, 0x34, 0xf9, 0x05, 0xf5, 0x94, 0xeb,
	0xe4, 0x3a, 0x5e, 0xce, 0x3a, 0x96, 0x67, 0x75, 0x6b, 0xd1, 0x17, 0x53, 0x8f, 0x66, 0xf4, 0x27,
	0x7c, 0xef, 0x27, 0x7a, 0x79, 0x87, 0xe3, 0x2d, 0x80, 0xaa, 0x68, 0xa1, 0x3d, 0x68, 0x19, 0x64,
	0xb4, 0x9e, 0x8b, 0x0f, 0xd8, 0xf8, 0xd0, 0x2c, 0x0c, 0x2e, 0x78, 0x72, 0x83, 0x9a, 0x78, 0x50,
	0x33, 0x5c, 0xc7, 0xf2, 0xaa, 0xdd, 0x46, 0x00, 0xcc, 0xd0, 0x03, 0xae, 0x79, 0x94, 0x13, 0x74,
	0x08, 0x76, 0x89, 0xbd, 0x63, 0x4d, 0x7e, 0x03, 0xe4, 0xd4, 0x9a, 0xa7, 0xd8, 0xf9, 0xe4, 0x59,
	0x5d, 0x3b, 0xb2, 0x0d, 0x72, 0xce, 0x53, 0xa4, 0x1b, 0x68, 0x4f, 0x26, 0x67, 0xf1, 0x65, 0x1e,
	0x6d, 0x31, 0xdd, 0x07, 0xbb, 0xcc, 0xd1, 0xf8, 0x35, 0x82, 0x6f, 0x6c, 0x2b, 0x1b, 0xec, 0xd0,
	0x3c, 0x9f, 0x67, 0x15, 0xe9, 0x41, 0x7d, 0x17, 0xaa, 0x19, 0xd2, 0x08, 0xda, 0xa6, 0x23, 0x2c,
	0x40, 0xb3, 0x7a, 0x29, 0x09, 0x1e, 0x2d, 0x70, 0x22, 0x31, 0x15, 0xa3, 0xf1, 0xb0, 0xaf, 0x57,
	0x5c, 0x11, 0x1f, 0xea, 0x21, 0x6a, 0x63, 0x4b, 0x9a, 0xec, 0x65, 0xfc, 0x2e, 0x61, 0x6f, 0x36,
	0xa4, 0x95, 0xff, 0x16, 0x39, 0x86, 0x56, 0x88, 0x7a, 0x2f, 0x7e, 0xf2, 0x83, 0x1d, 0x3a, 0x87,
	0xfb, 0x95, 0xed, 0xc5, 0x4b, 0x2b, 0xc4, 0x07, 0x27, 0x44, 0x5d, 0xde, 0x88, 0xb4, 0xd9, 0xeb,
	0x7b, 0xb9, 0x8e, 0x99, 0x1b, 0xe5, 0xaf, 0x41, 0x2b, 0x27, 0xb5, 0xab, 0xaa, 0x52, 0xab, 0xe9,
	0x67, 0xf3, 0x24, 0x47, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xab, 0xc6, 0x3b, 0xf1, 0xa4, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoboIMEAtlasClient is the client API for RoboIMEAtlas service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoboIMEAtlasClient interface {
	GetFrame(ctx context.Context, in *FrameRequest, opts ...grpc.CallOption) (RoboIMEAtlas_GetFrameClient, error)
	GetActiveMatches(ctx context.Context, in *ActiveMatchesRequest, opts ...grpc.CallOption) (*MatchesPacket, error)
	GetMatchInfo(ctx context.Context, in *MatchInfoRequest, opts ...grpc.CallOption) (*SSL_Referee, error)
}

type roboIMEAtlasClient struct {
	cc *grpc.ClientConn
}

func NewRoboIMEAtlasClient(cc *grpc.ClientConn) RoboIMEAtlasClient {
	return &roboIMEAtlasClient{cc}
}

func (c *roboIMEAtlasClient) GetFrame(ctx context.Context, in *FrameRequest, opts ...grpc.CallOption) (RoboIMEAtlas_GetFrameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoboIMEAtlas_serviceDesc.Streams[0], "/RoboIMEAtlas/GetFrame", opts...)
	if err != nil {
		return nil, err
	}
	x := &roboIMEAtlasGetFrameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoboIMEAtlas_GetFrameClient interface {
	Recv() (*SSL_WrapperPacket, error)
	grpc.ClientStream
}

type roboIMEAtlasGetFrameClient struct {
	grpc.ClientStream
}

func (x *roboIMEAtlasGetFrameClient) Recv() (*SSL_WrapperPacket, error) {
	m := new(SSL_WrapperPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *roboIMEAtlasClient) GetActiveMatches(ctx context.Context, in *ActiveMatchesRequest, opts ...grpc.CallOption) (*MatchesPacket, error) {
	out := new(MatchesPacket)
	err := c.cc.Invoke(ctx, "/RoboIMEAtlas/GetActiveMatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roboIMEAtlasClient) GetMatchInfo(ctx context.Context, in *MatchInfoRequest, opts ...grpc.CallOption) (*SSL_Referee, error) {
	out := new(SSL_Referee)
	err := c.cc.Invoke(ctx, "/RoboIMEAtlas/GetMatchInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoboIMEAtlasServer is the server API for RoboIMEAtlas service.
type RoboIMEAtlasServer interface {
	GetFrame(*FrameRequest, RoboIMEAtlas_GetFrameServer) error
	GetActiveMatches(context.Context, *ActiveMatchesRequest) (*MatchesPacket, error)
	GetMatchInfo(context.Context, *MatchInfoRequest) (*SSL_Referee, error)
}

// UnimplementedRoboIMEAtlasServer can be embedded to have forward compatible implementations.
type UnimplementedRoboIMEAtlasServer struct {
}

func (*UnimplementedRoboIMEAtlasServer) GetFrame(req *FrameRequest, srv RoboIMEAtlas_GetFrameServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFrame not implemented")
}
func (*UnimplementedRoboIMEAtlasServer) GetActiveMatches(ctx context.Context, req *ActiveMatchesRequest) (*MatchesPacket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveMatches not implemented")
}
func (*UnimplementedRoboIMEAtlasServer) GetMatchInfo(ctx context.Context, req *MatchInfoRequest) (*SSL_Referee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchInfo not implemented")
}

func RegisterRoboIMEAtlasServer(s *grpc.Server, srv RoboIMEAtlasServer) {
	s.RegisterService(&_RoboIMEAtlas_serviceDesc, srv)
}

func _RoboIMEAtlas_GetFrame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FrameRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoboIMEAtlasServer).GetFrame(m, &roboIMEAtlasGetFrameServer{stream})
}

type RoboIMEAtlas_GetFrameServer interface {
	Send(*SSL_WrapperPacket) error
	grpc.ServerStream
}

type roboIMEAtlasGetFrameServer struct {
	grpc.ServerStream
}

func (x *roboIMEAtlasGetFrameServer) Send(m *SSL_WrapperPacket) error {
	return x.ServerStream.SendMsg(m)
}

func _RoboIMEAtlas_GetActiveMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoboIMEAtlasServer).GetActiveMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoboIMEAtlas/GetActiveMatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoboIMEAtlasServer).GetActiveMatches(ctx, req.(*ActiveMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoboIMEAtlas_GetMatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoboIMEAtlasServer).GetMatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoboIMEAtlas/GetMatchInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoboIMEAtlasServer).GetMatchInfo(ctx, req.(*MatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoboIMEAtlas_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoboIMEAtlas",
	HandlerType: (*RoboIMEAtlasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActiveMatches",
			Handler:    _RoboIMEAtlas_GetActiveMatches_Handler,
		},
		{
			MethodName: "GetMatchInfo",
			Handler:    _RoboIMEAtlas_GetMatchInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFrame",
			Handler:       _RoboIMEAtlas_GetFrame_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messages_robocup_ssl_wrapper.proto",
}
