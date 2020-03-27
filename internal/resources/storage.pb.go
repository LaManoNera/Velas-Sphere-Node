// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package resources

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

type FileStorageRequest struct {
	// TODO: contract link/id/whatever with the same public key and with another verifiable signature
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	RequesterPublicKey   string   `protobuf:"bytes,3,opt,name=requesterPublicKey,proto3" json:"requesterPublicKey,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	MerkleTreeRoot       string   `protobuf:"bytes,5,opt,name=merkleTreeRoot,proto3" json:"merkleTreeRoot,omitempty"`
	Signature            string   `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileStorageRequest) Reset()         { *m = FileStorageRequest{} }
func (m *FileStorageRequest) String() string { return proto.CompactTextString(m) }
func (*FileStorageRequest) ProtoMessage()    {}
func (*FileStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *FileStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileStorageRequest.Unmarshal(m, b)
}
func (m *FileStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileStorageRequest.Marshal(b, m, deterministic)
}
func (m *FileStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileStorageRequest.Merge(m, src)
}
func (m *FileStorageRequest) XXX_Size() int {
	return xxx_messageInfo_FileStorageRequest.Size(m)
}
func (m *FileStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileStorageRequest proto.InternalMessageInfo

func (m *FileStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileStorageRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *FileStorageRequest) GetRequesterPublicKey() string {
	if m != nil {
		return m.RequesterPublicKey
	}
	return ""
}

func (m *FileStorageRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FileStorageRequest) GetMerkleTreeRoot() string {
	if m != nil {
		return m.MerkleTreeRoot
	}
	return ""
}

func (m *FileStorageRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type FileStorageResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GetBackToken         string   `protobuf:"bytes,2,opt,name=getBackToken,proto3" json:"getBackToken,omitempty"`
	VerificationToken    string   `protobuf:"bytes,3,opt,name=verificationToken,proto3" json:"verificationToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileStorageResponse) Reset()         { *m = FileStorageResponse{} }
func (m *FileStorageResponse) String() string { return proto.CompactTextString(m) }
func (*FileStorageResponse) ProtoMessage()    {}
func (*FileStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *FileStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileStorageResponse.Unmarshal(m, b)
}
func (m *FileStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileStorageResponse.Marshal(b, m, deterministic)
}
func (m *FileStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileStorageResponse.Merge(m, src)
}
func (m *FileStorageResponse) XXX_Size() int {
	return xxx_messageInfo_FileStorageResponse.Size(m)
}
func (m *FileStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileStorageResponse proto.InternalMessageInfo

func (m *FileStorageResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FileStorageResponse) GetGetBackToken() string {
	if m != nil {
		return m.GetBackToken
	}
	return ""
}

func (m *FileStorageResponse) GetVerificationToken() string {
	if m != nil {
		return m.VerificationToken
	}
	return ""
}

type FileStorageVerificationRequest struct {
	Id                         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Challenge                  string   `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	VerificationTokenSignature string   `protobuf:"bytes,3,opt,name=verificationTokenSignature,proto3" json:"verificationTokenSignature,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *FileStorageVerificationRequest) Reset()         { *m = FileStorageVerificationRequest{} }
func (m *FileStorageVerificationRequest) String() string { return proto.CompactTextString(m) }
func (*FileStorageVerificationRequest) ProtoMessage()    {}
func (*FileStorageVerificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *FileStorageVerificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileStorageVerificationRequest.Unmarshal(m, b)
}
func (m *FileStorageVerificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileStorageVerificationRequest.Marshal(b, m, deterministic)
}
func (m *FileStorageVerificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileStorageVerificationRequest.Merge(m, src)
}
func (m *FileStorageVerificationRequest) XXX_Size() int {
	return xxx_messageInfo_FileStorageVerificationRequest.Size(m)
}
func (m *FileStorageVerificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileStorageVerificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileStorageVerificationRequest proto.InternalMessageInfo

func (m *FileStorageVerificationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FileStorageVerificationRequest) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

func (m *FileStorageVerificationRequest) GetVerificationTokenSignature() string {
	if m != nil {
		return m.VerificationTokenSignature
	}
	return ""
}

type FileStorageVerificationResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	NewVerificationToken string   `protobuf:"bytes,3,opt,name=newVerificationToken,proto3" json:"newVerificationToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileStorageVerificationResponse) Reset()         { *m = FileStorageVerificationResponse{} }
func (m *FileStorageVerificationResponse) String() string { return proto.CompactTextString(m) }
func (*FileStorageVerificationResponse) ProtoMessage()    {}
func (*FileStorageVerificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}

func (m *FileStorageVerificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileStorageVerificationResponse.Unmarshal(m, b)
}
func (m *FileStorageVerificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileStorageVerificationResponse.Marshal(b, m, deterministic)
}
func (m *FileStorageVerificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileStorageVerificationResponse.Merge(m, src)
}
func (m *FileStorageVerificationResponse) XXX_Size() int {
	return xxx_messageInfo_FileStorageVerificationResponse.Size(m)
}
func (m *FileStorageVerificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileStorageVerificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileStorageVerificationResponse proto.InternalMessageInfo

func (m *FileStorageVerificationResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FileStorageVerificationResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileStorageVerificationResponse) GetNewVerificationToken() string {
	if m != nil {
		return m.NewVerificationToken
	}
	return ""
}

type GetFileBackRequest struct {
	Id                    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GetBackTokenSignature string   `protobuf:"bytes,2,opt,name=getBackTokenSignature,proto3" json:"getBackTokenSignature,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *GetFileBackRequest) Reset()         { *m = GetFileBackRequest{} }
func (m *GetFileBackRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileBackRequest) ProtoMessage()    {}
func (*GetFileBackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{4}
}

func (m *GetFileBackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileBackRequest.Unmarshal(m, b)
}
func (m *GetFileBackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileBackRequest.Marshal(b, m, deterministic)
}
func (m *GetFileBackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileBackRequest.Merge(m, src)
}
func (m *GetFileBackRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileBackRequest.Size(m)
}
func (m *GetFileBackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileBackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileBackRequest proto.InternalMessageInfo

func (m *GetFileBackRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetFileBackRequest) GetGetBackTokenSignature() string {
	if m != nil {
		return m.GetBackTokenSignature
	}
	return ""
}

type GetFileBackResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InvoiceID            string   `protobuf:"bytes,2,opt,name=invoiceID,proto3" json:"invoiceID,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileBackResponse) Reset()         { *m = GetFileBackResponse{} }
func (m *GetFileBackResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileBackResponse) ProtoMessage()    {}
func (*GetFileBackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{5}
}

func (m *GetFileBackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileBackResponse.Unmarshal(m, b)
}
func (m *GetFileBackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileBackResponse.Marshal(b, m, deterministic)
}
func (m *GetFileBackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileBackResponse.Merge(m, src)
}
func (m *GetFileBackResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileBackResponse.Size(m)
}
func (m *GetFileBackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileBackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileBackResponse proto.InternalMessageInfo

func (m *GetFileBackResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetFileBackResponse) GetInvoiceID() string {
	if m != nil {
		return m.InvoiceID
	}
	return ""
}

func (m *GetFileBackResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*FileStorageRequest)(nil), "resources.FileStorageRequest")
	proto.RegisterType((*FileStorageResponse)(nil), "resources.FileStorageResponse")
	proto.RegisterType((*FileStorageVerificationRequest)(nil), "resources.FileStorageVerificationRequest")
	proto.RegisterType((*FileStorageVerificationResponse)(nil), "resources.FileStorageVerificationResponse")
	proto.RegisterType((*GetFileBackRequest)(nil), "resources.GetFileBackRequest")
	proto.RegisterType((*GetFileBackResponse)(nil), "resources.GetFileBackResponse")
}

func init() {
	proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb)
}

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x5e, 0xb2, 0x52, 0xc8, 0x0b, 0x4c, 0xda, 0x3b, 0x40, 0x51, 0x55, 0x06, 0xf2, 0x01, 0x01,
	0x42, 0x39, 0x0c, 0xce, 0x1c, 0x10, 0x02, 0xa1, 0x49, 0x08, 0xb5, 0xd3, 0x90, 0xb8, 0x79, 0xe9,
	0x4b, 0x66, 0x35, 0xb3, 0x8b, 0xe3, 0xb4, 0xea, 0x9f, 0xe0, 0xc4, 0xaf, 0xe2, 0x57, 0xa1, 0xd8,
	0x69, 0x3e, 0x9a, 0x04, 0xed, 0xe6, 0xbe, 0xcf, 0x63, 0x3f, 0x1f, 0xb5, 0x03, 0x0f, 0x33, 0xa3,
	0x34, 0x4f, 0x28, 0x5a, 0x69, 0x65, 0x14, 0x06, 0x9a, 0x32, 0x95, 0xeb, 0x98, 0x32, 0xf6, 0xd7,
	0x03, 0xfc, 0x24, 0x52, 0x9a, 0x3b, 0xc2, 0x8c, 0x7e, 0xe5, 0x94, 0x19, 0x44, 0x18, 0x49, 0x7e,
	0x43, 0xa1, 0xf7, 0xdc, 0x7b, 0x19, 0xcc, 0xec, 0x1a, 0x9f, 0xc0, 0xd8, 0x70, 0x9d, 0x90, 0x09,
	0x7d, 0x3b, 0x2d, 0x7f, 0x61, 0x04, 0xa8, 0xdd, 0x36, 0xd2, 0xdf, 0xf2, 0xab, 0x54, 0xc4, 0xe7,
	0xb4, 0x0d, 0x0f, 0x2d, 0xa7, 0x07, 0x29, 0xce, 0x5e, 0x70, 0xc3, 0xc3, 0x91, 0x3b, 0xbb, 0x58,
	0xe3, 0x0b, 0x38, 0xba, 0x21, 0xbd, 0x4c, 0xe9, 0x42, 0x13, 0xcd, 0x94, 0x32, 0xe1, 0x1d, 0x8b,
	0xee, 0x4d, 0x71, 0x0a, 0x41, 0x26, 0x12, 0xc9, 0x4d, 0xae, 0x29, 0x1c, 0x5b, 0x4a, 0x3d, 0x60,
	0x1b, 0x38, 0x69, 0x65, 0xc9, 0x56, 0x4a, 0x66, 0x84, 0x47, 0xe0, 0x8b, 0x45, 0x19, 0xc5, 0x17,
	0x0b, 0x64, 0xf0, 0x20, 0x21, 0xf3, 0x81, 0xc7, 0xcb, 0x0b, 0xb5, 0x24, 0x59, 0xc6, 0x69, 0xcd,
	0xf0, 0x0d, 0x1c, 0xaf, 0x49, 0x8b, 0x9f, 0x22, 0xe6, 0x46, 0x28, 0xe9, 0x88, 0x2e, 0x53, 0x17,
	0x60, 0xbf, 0x3d, 0x38, 0x6d, 0x28, 0x5f, 0x36, 0x08, 0xbb, 0x46, 0xf7, 0x4d, 0x4c, 0x21, 0x88,
	0xaf, 0x79, 0x9a, 0x92, 0x4c, 0xa8, 0x74, 0x50, 0x0f, 0xf0, 0x3d, 0x4c, 0x3a, 0x2a, 0xf3, 0x2a,
	0xb8, 0xf3, 0xf1, 0x1f, 0x06, 0xdb, 0xc2, 0xb3, 0x41, 0x3f, 0x03, 0xad, 0x20, 0x8c, 0x56, 0xdc,
	0x5c, 0x97, 0x5e, 0xec, 0x1a, 0xcf, 0xe0, 0x91, 0xa4, 0xcd, 0xe5, 0x40, 0x11, 0xbd, 0x18, 0xfb,
	0x01, 0xf8, 0x99, 0x4c, 0xa1, 0x5e, 0xb4, 0x39, 0x14, 0xff, 0x1d, 0x3c, 0x6e, 0xf6, 0x5d, 0x67,
	0x73, 0xf2, 0xfd, 0x20, 0xfb, 0x0e, 0x27, 0xad, 0xb3, 0x07, 0xa2, 0x4c, 0x21, 0x10, 0x72, 0xad,
	0x44, 0x4c, 0x5f, 0x3e, 0xee, 0xba, 0xad, 0x06, 0xd5, 0xfd, 0x3b, 0xac, 0xef, 0xdf, 0xd9, 0x1f,
	0x1f, 0xee, 0x96, 0x65, 0xe1, 0x39, 0xdc, 0x9b, 0xf3, 0x35, 0x15, 0x2a, 0xf8, 0x34, 0xaa, 0x9e,
	0x4a, 0xd4, 0x7d, 0x26, 0x93, 0xd3, 0x21, 0xd8, 0x19, 0x63, 0x07, 0x98, 0xc2, 0xb1, 0xad, 0x68,
	0xdb, 0x80, 0xf1, 0x55, 0xff, 0xb6, 0x9e, 0x6b, 0x33, 0x79, 0x7d, 0x1b, 0x6a, 0xa5, 0xf6, 0x15,
	0xee, 0x37, 0xfa, 0x69, 0xb9, 0xef, 0xfe, 0x27, 0x2d, 0xf7, 0x3d, 0xb5, 0xb2, 0x83, 0xab, 0xb1,
	0xfd, 0x5e, 0xbc, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x85, 0x33, 0x76, 0x9e, 0x40, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	SaveFile(ctx context.Context, in *FileStorageRequest, opts ...grpc.CallOption) (*FileStorageResponse, error)
	VerifyFileStorage(ctx context.Context, in *FileStorageVerificationRequest, opts ...grpc.CallOption) (*FileStorageVerificationResponse, error)
	GetFileBack(ctx context.Context, in *GetFileBackRequest, opts ...grpc.CallOption) (*GetFileBackResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SaveFile(ctx context.Context, in *FileStorageRequest, opts ...grpc.CallOption) (*FileStorageResponse, error) {
	out := new(FileStorageResponse)
	err := c.cc.Invoke(ctx, "/resources.Storage/SaveFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) VerifyFileStorage(ctx context.Context, in *FileStorageVerificationRequest, opts ...grpc.CallOption) (*FileStorageVerificationResponse, error) {
	out := new(FileStorageVerificationResponse)
	err := c.cc.Invoke(ctx, "/resources.Storage/VerifyFileStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetFileBack(ctx context.Context, in *GetFileBackRequest, opts ...grpc.CallOption) (*GetFileBackResponse, error) {
	out := new(GetFileBackResponse)
	err := c.cc.Invoke(ctx, "/resources.Storage/GetFileBack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	SaveFile(context.Context, *FileStorageRequest) (*FileStorageResponse, error)
	VerifyFileStorage(context.Context, *FileStorageVerificationRequest) (*FileStorageVerificationResponse, error)
	GetFileBack(context.Context, *GetFileBackRequest) (*GetFileBackResponse, error)
}

// UnimplementedStorageServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (*UnimplementedStorageServer) SaveFile(ctx context.Context, req *FileStorageRequest) (*FileStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (*UnimplementedStorageServer) VerifyFileStorage(ctx context.Context, req *FileStorageVerificationRequest) (*FileStorageVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyFileStorage not implemented")
}
func (*UnimplementedStorageServer) GetFileBack(ctx context.Context, req *GetFileBackRequest) (*GetFileBackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileBack not implemented")
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_SaveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Storage/SaveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveFile(ctx, req.(*FileStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_VerifyFileStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileStorageVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).VerifyFileStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Storage/VerifyFileStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).VerifyFileStorage(ctx, req.(*FileStorageVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetFileBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileBackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetFileBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Storage/GetFileBack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetFileBack(ctx, req.(*GetFileBackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resources.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveFile",
			Handler:    _Storage_SaveFile_Handler,
		},
		{
			MethodName: "VerifyFileStorage",
			Handler:    _Storage_VerifyFileStorage_Handler,
		},
		{
			MethodName: "GetFileBack",
			Handler:    _Storage_GetFileBack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}