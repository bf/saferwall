// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multiav.avast.proto

package avast_api

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
	return fileDescriptor_0394137c802eaed6, []int{0}
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

// The scan file binary request message containing the binary data of the file
// to scan.
type ScanFileBinaryRequest struct {
	File                 []byte   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanFileBinaryRequest) Reset()         { *m = ScanFileBinaryRequest{} }
func (m *ScanFileBinaryRequest) String() string { return proto.CompactTextString(m) }
func (*ScanFileBinaryRequest) ProtoMessage()    {}
func (*ScanFileBinaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0394137c802eaed6, []int{1}
}

func (m *ScanFileBinaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanFileBinaryRequest.Unmarshal(m, b)
}
func (m *ScanFileBinaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanFileBinaryRequest.Marshal(b, m, deterministic)
}
func (m *ScanFileBinaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanFileBinaryRequest.Merge(m, src)
}
func (m *ScanFileBinaryRequest) XXX_Size() int {
	return xxx_messageInfo_ScanFileBinaryRequest.Size(m)
}
func (m *ScanFileBinaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanFileBinaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanFileBinaryRequest proto.InternalMessageInfo

func (m *ScanFileBinaryRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
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
	return fileDescriptor_0394137c802eaed6, []int{2}
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
	return fileDescriptor_0394137c802eaed6, []int{3}
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
	return fileDescriptor_0394137c802eaed6, []int{4}
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

// The license request message contains the license in binary format.
type LicenseActivationRequest struct {
	License              []byte   `protobuf:"bytes,1,opt,name=license,proto3" json:"license,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LicenseActivationRequest) Reset()         { *m = LicenseActivationRequest{} }
func (m *LicenseActivationRequest) String() string { return proto.CompactTextString(m) }
func (*LicenseActivationRequest) ProtoMessage()    {}
func (*LicenseActivationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0394137c802eaed6, []int{5}
}

func (m *LicenseActivationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LicenseActivationRequest.Unmarshal(m, b)
}
func (m *LicenseActivationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LicenseActivationRequest.Marshal(b, m, deterministic)
}
func (m *LicenseActivationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LicenseActivationRequest.Merge(m, src)
}
func (m *LicenseActivationRequest) XXX_Size() int {
	return xxx_messageInfo_LicenseActivationRequest.Size(m)
}
func (m *LicenseActivationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LicenseActivationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LicenseActivationRequest proto.InternalMessageInfo

func (m *LicenseActivationRequest) GetLicense() []byte {
	if m != nil {
		return m.License
	}
	return nil
}

// The license response message.
type LicenseActivationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LicenseActivationResponse) Reset()         { *m = LicenseActivationResponse{} }
func (m *LicenseActivationResponse) String() string { return proto.CompactTextString(m) }
func (*LicenseActivationResponse) ProtoMessage()    {}
func (*LicenseActivationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0394137c802eaed6, []int{6}
}

func (m *LicenseActivationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LicenseActivationResponse.Unmarshal(m, b)
}
func (m *LicenseActivationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LicenseActivationResponse.Marshal(b, m, deterministic)
}
func (m *LicenseActivationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LicenseActivationResponse.Merge(m, src)
}
func (m *LicenseActivationResponse) XXX_Size() int {
	return xxx_messageInfo_LicenseActivationResponse.Size(m)
}
func (m *LicenseActivationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LicenseActivationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LicenseActivationResponse proto.InternalMessageInfo

// The update vps request.
type UpdateVPSRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVPSRequest) Reset()         { *m = UpdateVPSRequest{} }
func (m *UpdateVPSRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateVPSRequest) ProtoMessage()    {}
func (*UpdateVPSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0394137c802eaed6, []int{7}
}

func (m *UpdateVPSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVPSRequest.Unmarshal(m, b)
}
func (m *UpdateVPSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVPSRequest.Marshal(b, m, deterministic)
}
func (m *UpdateVPSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVPSRequest.Merge(m, src)
}
func (m *UpdateVPSRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateVPSRequest.Size(m)
}
func (m *UpdateVPSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVPSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVPSRequest proto.InternalMessageInfo

// The pdate vps request response.
type UpdateVPSResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVPSResponse) Reset()         { *m = UpdateVPSResponse{} }
func (m *UpdateVPSResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateVPSResponse) ProtoMessage()    {}
func (*UpdateVPSResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0394137c802eaed6, []int{8}
}

func (m *UpdateVPSResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVPSResponse.Unmarshal(m, b)
}
func (m *UpdateVPSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVPSResponse.Marshal(b, m, deterministic)
}
func (m *UpdateVPSResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVPSResponse.Merge(m, src)
}
func (m *UpdateVPSResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateVPSResponse.Size(m)
}
func (m *UpdateVPSResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVPSResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVPSResponse proto.InternalMessageInfo

// The license status request.
type LicenseStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LicenseStatusRequest) Reset()         { *m = LicenseStatusRequest{} }
func (m *LicenseStatusRequest) String() string { return proto.CompactTextString(m) }
func (*LicenseStatusRequest) ProtoMessage()    {}
func (*LicenseStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0394137c802eaed6, []int{9}
}

func (m *LicenseStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LicenseStatusRequest.Unmarshal(m, b)
}
func (m *LicenseStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LicenseStatusRequest.Marshal(b, m, deterministic)
}
func (m *LicenseStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LicenseStatusRequest.Merge(m, src)
}
func (m *LicenseStatusRequest) XXX_Size() int {
	return xxx_messageInfo_LicenseStatusRequest.Size(m)
}
func (m *LicenseStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LicenseStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LicenseStatusRequest proto.InternalMessageInfo

// The license status response.
type LicenseStatusResponse struct {
	Expired              bool     `protobuf:"varint,1,opt,name=expired,proto3" json:"expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LicenseStatusResponse) Reset()         { *m = LicenseStatusResponse{} }
func (m *LicenseStatusResponse) String() string { return proto.CompactTextString(m) }
func (*LicenseStatusResponse) ProtoMessage()    {}
func (*LicenseStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0394137c802eaed6, []int{10}
}

func (m *LicenseStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LicenseStatusResponse.Unmarshal(m, b)
}
func (m *LicenseStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LicenseStatusResponse.Marshal(b, m, deterministic)
}
func (m *LicenseStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LicenseStatusResponse.Merge(m, src)
}
func (m *LicenseStatusResponse) XXX_Size() int {
	return xxx_messageInfo_LicenseStatusResponse.Size(m)
}
func (m *LicenseStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LicenseStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LicenseStatusResponse proto.InternalMessageInfo

func (m *LicenseStatusResponse) GetExpired() bool {
	if m != nil {
		return m.Expired
	}
	return false
}

func init() {
	proto.RegisterType((*ScanFileRequest)(nil), "avast.api.ScanFileRequest")
	proto.RegisterType((*ScanFileBinaryRequest)(nil), "avast.api.ScanFileBinaryRequest")
	proto.RegisterType((*ScanResponse)(nil), "avast.api.ScanResponse")
	proto.RegisterType((*VersionRequest)(nil), "avast.api.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "avast.api.VersionResponse")
	proto.RegisterType((*LicenseActivationRequest)(nil), "avast.api.LicenseActivationRequest")
	proto.RegisterType((*LicenseActivationResponse)(nil), "avast.api.LicenseActivationResponse")
	proto.RegisterType((*UpdateVPSRequest)(nil), "avast.api.UpdateVPSRequest")
	proto.RegisterType((*UpdateVPSResponse)(nil), "avast.api.UpdateVPSResponse")
	proto.RegisterType((*LicenseStatusRequest)(nil), "avast.api.LicenseStatusRequest")
	proto.RegisterType((*LicenseStatusResponse)(nil), "avast.api.LicenseStatusResponse")
}

func init() { proto.RegisterFile("multiav.avast.proto", fileDescriptor_0394137c802eaed6) }

var fileDescriptor_0394137c802eaed6 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0xa8, 0x2a, 0xc9, 0x28, 0x34, 0xe9, 0x96, 0x16, 0x77, 0x8b, 0x84, 0xb5, 0x70, 0xa8,
	0x54, 0x61, 0x89, 0x8f, 0x3f, 0xd0, 0x4a, 0x10, 0x0e, 0x45, 0xb2, 0x6c, 0x11, 0x6e, 0x48, 0x8b,
	0x33, 0xa5, 0x2b, 0xb9, 0xb6, 0xb1, 0xc7, 0x16, 0xfc, 0x3a, 0xfe, 0x1a, 0xb2, 0xbd, 0xbb, 0xb5,
	0x8d, 0x93, 0x4b, 0x6f, 0x79, 0x33, 0xef, 0xbd, 0x9d, 0xcc, 0x1b, 0xc3, 0xd1, 0x5d, 0x19, 0x93,
	0x92, 0x95, 0x27, 0x2b, 0x59, 0x90, 0x97, 0xe5, 0x29, 0xa5, 0x6c, 0xd6, 0x02, 0x99, 0x29, 0xf1,
	0x06, 0x16, 0x61, 0x24, 0x93, 0x4f, 0x2a, 0xc6, 0x00, 0x7f, 0x95, 0x58, 0x10, 0xe3, 0x30, 0xbd,
	0x51, 0x31, 0x66, 0x92, 0x6e, 0x9d, 0x89, 0x3b, 0x39, 0x9f, 0x05, 0x16, 0x8b, 0x0b, 0x38, 0x36,
	0xf4, 0x2b, 0x95, 0xc8, 0xfc, 0x8f, 0x11, 0x31, 0xd8, 0xab, 0x49, 0x8d, 0x60, 0x1e, 0x34, 0xbf,
	0xc5, 0x15, 0xcc, 0x6b, 0x72, 0x80, 0x45, 0x96, 0x26, 0x05, 0xb2, 0x13, 0xd8, 0x4f, 0x4b, 0xca,
	0x4a, 0xd2, 0xb6, 0x1a, 0xd5, 0x0f, 0xaa, 0xe4, 0x06, 0x23, 0xc2, 0x8d, 0xf3, 0xd8, 0x9d, 0x9c,
	0x4f, 0x03, 0x8b, 0xc5, 0x12, 0x0e, 0xd6, 0x98, 0x17, 0x2a, 0x4d, 0xf4, 0x4b, 0xe2, 0x02, 0x16,
	0xb6, 0xa2, 0x8d, 0x1d, 0x78, 0x52, 0xb5, 0x25, 0xed, 0x6c, 0xa0, 0xf8, 0x00, 0xce, 0xb5, 0x8a,
	0x30, 0x29, 0xf0, 0x32, 0x22, 0x55, 0x49, 0xba, 0x37, 0xaa, 0x55, 0x71, 0xdb, 0xd3, 0x53, 0x1b,
	0x28, 0xce, 0xe0, 0x74, 0x44, 0xd5, 0x3e, 0x26, 0x18, 0x2c, 0xbf, 0x66, 0x1b, 0x49, 0xb8, 0xf6,
	0x43, 0x33, 0xd3, 0x11, 0x1c, 0x76, 0x6a, 0x9a, 0x78, 0x02, 0xcf, 0xb4, 0x4b, 0x48, 0x92, 0xca,
	0xc2, 0x90, 0xdf, 0xc2, 0xf1, 0xa0, 0x7e, 0xff, 0x37, 0xf0, 0x77, 0xa6, 0x72, 0xdc, 0x34, 0x03,
	0x4d, 0x03, 0x03, 0xdf, 0xfd, 0xdd, 0x83, 0xf9, 0x65, 0x9d, 0x59, 0xbd, 0xcf, 0x04, 0x73, 0xf6,
	0xb1, 0x5d, 0x6d, 0x9d, 0x83, 0x2f, 0xe9, 0x96, 0x71, 0xcf, 0x46, 0xea, 0x0d, 0xf2, 0xe4, 0xcf,
	0x07, 0x3d, 0x3b, 0xe0, 0x23, 0xf6, 0x05, 0x0e, 0xfa, 0x71, 0x32, 0x77, 0xc4, 0xa8, 0x97, 0xf4,
	0x2e, 0xbb, 0x6b, 0x38, 0x5c, 0x21, 0xf9, 0x79, 0xfa, 0x33, 0x97, 0x77, 0x3a, 0x24, 0x76, 0xda,
	0xe1, 0xf7, 0xa3, 0xe4, 0x7c, 0xac, 0x65, 0xdd, 0x3e, 0xc3, 0xd3, 0x15, 0xd2, 0xda, 0x0f, 0x1f,
	0xec, 0xf4, 0x1d, 0x16, 0x3a, 0x48, 0xd4, 0x9b, 0x67, 0xaf, 0x3a, 0x82, 0x6d, 0x17, 0xc2, 0x5f,
	0xef, 0x26, 0x59, 0xff, 0x6f, 0xb0, 0x5c, 0x21, 0xf5, 0x42, 0x65, 0x2f, 0xff, 0xd7, 0xf6, 0xce,
	0x80, 0xbb, 0xdb, 0x09, 0x9d, 0x15, 0xcc, 0xec, 0x5d, 0xb1, 0xb3, 0x8e, 0x60, 0x78, 0x81, 0xfc,
	0xc5, 0x78, 0xd3, 0x38, 0xfd, 0xd8, 0x6f, 0xbe, 0xfc, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0x9e, 0x72, 0x82, 0x10, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AvastScannerClient is the client API for AvastScanner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AvastScannerClient interface {
	// Scan a file from a file path
	ScanFilePath(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error)
	// Scan a file from a binary blob
	ScanFileBinary(ctx context.Context, in *ScanFileBinaryRequest, opts ...grpc.CallOption) (*ScanResponse, error)
	// Get program version
	GetProgramVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Get VPS version
	GetVPSVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Activate product license
	ActivateLicense(ctx context.Context, in *LicenseActivationRequest, opts ...grpc.CallOption) (*LicenseActivationResponse, error)
	// Check license status
	GetLicenseStatus(ctx context.Context, in *LicenseStatusRequest, opts ...grpc.CallOption) (*LicenseStatusResponse, error)
	// Update VPS
	UpdateVPS(ctx context.Context, in *UpdateVPSRequest, opts ...grpc.CallOption) (*UpdateVPSResponse, error)
}

type avastScannerClient struct {
	cc *grpc.ClientConn
}

func NewAvastScannerClient(cc *grpc.ClientConn) AvastScannerClient {
	return &avastScannerClient{cc}
}

func (c *avastScannerClient) ScanFilePath(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error) {
	out := new(ScanResponse)
	err := c.cc.Invoke(ctx, "/avast.api.AvastScanner/ScanFilePath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avastScannerClient) ScanFileBinary(ctx context.Context, in *ScanFileBinaryRequest, opts ...grpc.CallOption) (*ScanResponse, error) {
	out := new(ScanResponse)
	err := c.cc.Invoke(ctx, "/avast.api.AvastScanner/ScanFileBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avastScannerClient) GetProgramVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/avast.api.AvastScanner/GetProgramVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avastScannerClient) GetVPSVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/avast.api.AvastScanner/GetVPSVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avastScannerClient) ActivateLicense(ctx context.Context, in *LicenseActivationRequest, opts ...grpc.CallOption) (*LicenseActivationResponse, error) {
	out := new(LicenseActivationResponse)
	err := c.cc.Invoke(ctx, "/avast.api.AvastScanner/ActivateLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avastScannerClient) GetLicenseStatus(ctx context.Context, in *LicenseStatusRequest, opts ...grpc.CallOption) (*LicenseStatusResponse, error) {
	out := new(LicenseStatusResponse)
	err := c.cc.Invoke(ctx, "/avast.api.AvastScanner/GetLicenseStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avastScannerClient) UpdateVPS(ctx context.Context, in *UpdateVPSRequest, opts ...grpc.CallOption) (*UpdateVPSResponse, error) {
	out := new(UpdateVPSResponse)
	err := c.cc.Invoke(ctx, "/avast.api.AvastScanner/UpdateVPS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvastScannerServer is the server API for AvastScanner service.
type AvastScannerServer interface {
	// Scan a file from a file path
	ScanFilePath(context.Context, *ScanFileRequest) (*ScanResponse, error)
	// Scan a file from a binary blob
	ScanFileBinary(context.Context, *ScanFileBinaryRequest) (*ScanResponse, error)
	// Get program version
	GetProgramVersion(context.Context, *VersionRequest) (*VersionResponse, error)
	// Get VPS version
	GetVPSVersion(context.Context, *VersionRequest) (*VersionResponse, error)
	// Activate product license
	ActivateLicense(context.Context, *LicenseActivationRequest) (*LicenseActivationResponse, error)
	// Check license status
	GetLicenseStatus(context.Context, *LicenseStatusRequest) (*LicenseStatusResponse, error)
	// Update VPS
	UpdateVPS(context.Context, *UpdateVPSRequest) (*UpdateVPSResponse, error)
}

// UnimplementedAvastScannerServer can be embedded to have forward compatible implementations.
type UnimplementedAvastScannerServer struct {
}

func (*UnimplementedAvastScannerServer) ScanFilePath(ctx context.Context, req *ScanFileRequest) (*ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanFilePath not implemented")
}
func (*UnimplementedAvastScannerServer) ScanFileBinary(ctx context.Context, req *ScanFileBinaryRequest) (*ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanFileBinary not implemented")
}
func (*UnimplementedAvastScannerServer) GetProgramVersion(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgramVersion not implemented")
}
func (*UnimplementedAvastScannerServer) GetVPSVersion(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVPSVersion not implemented")
}
func (*UnimplementedAvastScannerServer) ActivateLicense(ctx context.Context, req *LicenseActivationRequest) (*LicenseActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateLicense not implemented")
}
func (*UnimplementedAvastScannerServer) GetLicenseStatus(ctx context.Context, req *LicenseStatusRequest) (*LicenseStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLicenseStatus not implemented")
}
func (*UnimplementedAvastScannerServer) UpdateVPS(ctx context.Context, req *UpdateVPSRequest) (*UpdateVPSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVPS not implemented")
}

func RegisterAvastScannerServer(s *grpc.Server, srv AvastScannerServer) {
	s.RegisterService(&_AvastScanner_serviceDesc, srv)
}

func _AvastScanner_ScanFilePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvastScannerServer).ScanFilePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avast.api.AvastScanner/ScanFilePath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvastScannerServer).ScanFilePath(ctx, req.(*ScanFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvastScanner_ScanFileBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanFileBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvastScannerServer).ScanFileBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avast.api.AvastScanner/ScanFileBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvastScannerServer).ScanFileBinary(ctx, req.(*ScanFileBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvastScanner_GetProgramVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvastScannerServer).GetProgramVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avast.api.AvastScanner/GetProgramVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvastScannerServer).GetProgramVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvastScanner_GetVPSVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvastScannerServer).GetVPSVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avast.api.AvastScanner/GetVPSVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvastScannerServer).GetVPSVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvastScanner_ActivateLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LicenseActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvastScannerServer).ActivateLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avast.api.AvastScanner/ActivateLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvastScannerServer).ActivateLicense(ctx, req.(*LicenseActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvastScanner_GetLicenseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LicenseStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvastScannerServer).GetLicenseStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avast.api.AvastScanner/GetLicenseStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvastScannerServer).GetLicenseStatus(ctx, req.(*LicenseStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvastScanner_UpdateVPS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVPSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvastScannerServer).UpdateVPS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avast.api.AvastScanner/UpdateVPS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvastScannerServer).UpdateVPS(ctx, req.(*UpdateVPSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AvastScanner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "avast.api.AvastScanner",
	HandlerType: (*AvastScannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanFilePath",
			Handler:    _AvastScanner_ScanFilePath_Handler,
		},
		{
			MethodName: "ScanFileBinary",
			Handler:    _AvastScanner_ScanFileBinary_Handler,
		},
		{
			MethodName: "GetProgramVersion",
			Handler:    _AvastScanner_GetProgramVersion_Handler,
		},
		{
			MethodName: "GetVPSVersion",
			Handler:    _AvastScanner_GetVPSVersion_Handler,
		},
		{
			MethodName: "ActivateLicense",
			Handler:    _AvastScanner_ActivateLicense_Handler,
		},
		{
			MethodName: "GetLicenseStatus",
			Handler:    _AvastScanner_GetLicenseStatus_Handler,
		},
		{
			MethodName: "UpdateVPS",
			Handler:    _AvastScanner_UpdateVPS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multiav.avast.proto",
}
