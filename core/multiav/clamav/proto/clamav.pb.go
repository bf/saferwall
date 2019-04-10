// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multiav.clamav.proto

package clamav_api

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

// The scan file request message containing the file path to scan.
type ScanFileRequest struct {
	Filepath             string   `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanFileRequest) Reset()         { *m = ScanFileRequest{} }
func (m *ScanFileRequest) String() string { return proto.CompactTextString(m) }
func (*ScanFileRequest) ProtoMessage()    {}
func (*ScanFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6c358cfaa39550, []int{0}
}

func (m *ScanFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanFileRequest.Unmarshal(m, b)
}
func (m *ScanFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanFileRequest.Marshal(b, m, deterministic)
}
func (m *ScanFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanFileRequest.Merge(m, src)
}
func (m *ScanFileRequest) XXX_Size() int {
	return xxx_messageInfo_ScanFileRequest.Size(m)
}
func (m *ScanFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanFileRequest proto.InternalMessageInfo

func (m *ScanFileRequest) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

// The scan response message containing detection results of the AntiVirus.
type ScanResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Infected             bool     `protobuf:"varint,2,opt,name=infected,proto3" json:"infected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanResponse) Reset()         { *m = ScanResponse{} }
func (m *ScanResponse) String() string { return proto.CompactTextString(m) }
func (*ScanResponse) ProtoMessage()    {}
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6c358cfaa39550, []int{1}
}

func (m *ScanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanResponse.Unmarshal(m, b)
}
func (m *ScanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanResponse.Marshal(b, m, deterministic)
}
func (m *ScanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanResponse.Merge(m, src)
}
func (m *ScanResponse) XXX_Size() int {
	return xxx_messageInfo_ScanResponse.Size(m)
}
func (m *ScanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScanResponse proto.InternalMessageInfo

func (m *ScanResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *ScanResponse) GetInfected() bool {
	if m != nil {
		return m.Infected
	}
	return false
}

// The version request message ask for version.
type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6c358cfaa39550, []int{2}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

// The response message containing program/VPS version.
type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6c358cfaa39550, []int{3}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*ScanFileRequest)(nil), "clamav.api.ScanFileRequest")
	proto.RegisterType((*ScanResponse)(nil), "clamav.api.ScanResponse")
	proto.RegisterType((*VersionRequest)(nil), "clamav.api.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "clamav.api.VersionResponse")
}

func init() { proto.RegisterFile("multiav.clamav.proto", fileDescriptor_fe6c358cfaa39550) }

var fileDescriptor_fe6c358cfaa39550 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xad, 0x87, 0xb5, 0x0e, 0xea, 0x4a, 0x10, 0x29, 0xd9, 0xcb, 0x92, 0xd3, 0x82, 0x98,
	0x83, 0xfe, 0x02, 0x5d, 0x70, 0xef, 0x11, 0xf6, 0x1e, 0xeb, 0x2c, 0x06, 0xd2, 0x24, 0x36, 0x93,
	0xfe, 0x1b, 0xff, 0xab, 0x54, 0x93, 0xaa, 0xc5, 0xe3, 0x37, 0xf3, 0x78, 0xf3, 0xde, 0xc0, 0x55,
	0x97, 0x2c, 0x19, 0x3d, 0xc8, 0xd6, 0xea, 0x4e, 0x0f, 0x32, 0xf4, 0x9e, 0x3c, 0x83, 0x4c, 0x3a,
	0x18, 0x71, 0x0b, 0xcb, 0xe7, 0x56, 0xbb, 0x27, 0x63, 0x51, 0xe1, 0x7b, 0xc2, 0x48, 0x8c, 0x43,
	0x7d, 0x30, 0x16, 0x83, 0xa6, 0xb7, 0xa6, 0x5a, 0x57, 0x9b, 0x53, 0x35, 0xb1, 0x78, 0x84, 0xb3,
	0x51, 0xae, 0x30, 0x06, 0xef, 0x22, 0xb2, 0x6b, 0x58, 0xf8, 0x44, 0x21, 0x51, 0x56, 0x66, 0x1a,
	0x3d, 0x8c, 0x3b, 0x60, 0x4b, 0xf8, 0xda, 0x1c, 0xaf, 0xab, 0x4d, 0xad, 0x26, 0x16, 0x97, 0x70,
	0xb1, 0xc7, 0x3e, 0x1a, 0xef, 0xf2, 0x45, 0x71, 0x03, 0xcb, 0x69, 0x92, 0x8d, 0x1b, 0x38, 0x19,
	0xbe, 0x47, 0xd9, 0xb9, 0xe0, 0xdd, 0x47, 0x05, 0xe7, 0x5b, 0xab, 0xbb, 0x87, 0xfd, 0x98, 0xc4,
	0x61, 0xcf, 0xb6, 0x50, 0x97, 0x0e, 0x6c, 0x25, 0x7f, 0xca, 0xc9, 0x59, 0x33, 0xde, 0xcc, 0x97,
	0xe5, 0x9c, 0x38, 0x62, 0x3b, 0x80, 0x1d, 0x52, 0x8e, 0xc1, 0xf8, 0x6f, 0xe5, 0xdf, 0xb4, 0x7c,
	0xf5, 0xef, 0xae, 0x18, 0xbd, 0x2c, 0xbe, 0x9e, 0x7c, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x15, 0xa0, 0xcc, 0x7c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClamAVScannerClient is the client API for ClamAVScanner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClamAVScannerClient interface {
	// Scan a file
	ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error)
	// Get program version
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type clamAVScannerClient struct {
	cc *grpc.ClientConn
}

func NewClamAVScannerClient(cc *grpc.ClientConn) ClamAVScannerClient {
	return &clamAVScannerClient{cc}
}

func (c *clamAVScannerClient) ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error) {
	out := new(ScanResponse)
	err := c.cc.Invoke(ctx, "/clamav.api.ClamAVScanner/ScanFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clamAVScannerClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/clamav.api.ClamAVScanner/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClamAVScannerServer is the server API for ClamAVScanner service.
type ClamAVScannerServer interface {
	// Scan a file
	ScanFile(context.Context, *ScanFileRequest) (*ScanResponse, error)
	// Get program version
	GetVersion(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedClamAVScannerServer can be embedded to have forward compatible implementations.
type UnimplementedClamAVScannerServer struct {
}

func (*UnimplementedClamAVScannerServer) ScanFile(ctx context.Context, req *ScanFileRequest) (*ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanFile not implemented")
}
func (*UnimplementedClamAVScannerServer) GetVersion(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterClamAVScannerServer(s *grpc.Server, srv ClamAVScannerServer) {
	s.RegisterService(&_ClamAVScanner_serviceDesc, srv)
}

func _ClamAVScanner_ScanFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClamAVScannerServer).ScanFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clamav.api.ClamAVScanner/ScanFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClamAVScannerServer).ScanFile(ctx, req.(*ScanFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClamAVScanner_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClamAVScannerServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clamav.api.ClamAVScanner/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClamAVScannerServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClamAVScanner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clamav.api.ClamAVScanner",
	HandlerType: (*ClamAVScannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanFile",
			Handler:    _ClamAVScanner_ScanFile_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ClamAVScanner_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multiav.clamav.proto",
}
