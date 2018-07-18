// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FolderInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FolderInfo) Reset()         { *m = FolderInfo{} }
func (m *FolderInfo) String() string { return proto.CompactTextString(m) }
func (*FolderInfo) ProtoMessage()    {}
func (*FolderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_b67e7327b0ef4312, []int{0}
}
func (m *FolderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FolderInfo.Unmarshal(m, b)
}
func (m *FolderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FolderInfo.Marshal(b, m, deterministic)
}
func (dst *FolderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FolderInfo.Merge(dst, src)
}
func (m *FolderInfo) XXX_Size() int {
	return xxx_messageInfo_FolderInfo.Size(m)
}
func (m *FolderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FolderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FolderInfo proto.InternalMessageInfo

func (m *FolderInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FileInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Mode                 uint32   `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_b67e7327b0ef4312, []int{1}
}
func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (dst *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(dst, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileInfo) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type FileData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileData) Reset()         { *m = FileData{} }
func (m *FileData) String() string { return proto.CompactTextString(m) }
func (*FileData) ProtoMessage()    {}
func (*FileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_b67e7327b0ef4312, []int{2}
}
func (m *FileData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileData.Unmarshal(m, b)
}
func (m *FileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileData.Marshal(b, m, deterministic)
}
func (dst *FileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileData.Merge(dst, src)
}
func (m *FileData) XXX_Size() int {
	return xxx_messageInfo_FileData.Size(m)
}
func (m *FileData) XXX_DiscardUnknown() {
	xxx_messageInfo_FileData.DiscardUnknown(m)
}

var xxx_messageInfo_FileData proto.InternalMessageInfo

func (m *FileData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Res struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_b67e7327b0ef4312, []int{3}
}
func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (dst *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(dst, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FolderInfo)(nil), "FolderInfo")
	proto.RegisterType((*FileInfo)(nil), "FileInfo")
	proto.RegisterType((*FileData)(nil), "FileData")
	proto.RegisterType((*Res)(nil), "Res")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileTransferServiceClient is the client API for FileTransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileTransferServiceClient interface {
	GetFolderInfo(ctx context.Context, in *FolderInfo, opts ...grpc.CallOption) (*Res, error)
	GetFileInfo(ctx context.Context, in *FileInfo, opts ...grpc.CallOption) (*Res, error)
	TransferFile(ctx context.Context, opts ...grpc.CallOption) (FileTransferService_TransferFileClient, error)
}

type fileTransferServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileTransferServiceClient(cc *grpc.ClientConn) FileTransferServiceClient {
	return &fileTransferServiceClient{cc}
}

func (c *fileTransferServiceClient) GetFolderInfo(ctx context.Context, in *FolderInfo, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/FileTransferService/getFolderInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTransferServiceClient) GetFileInfo(ctx context.Context, in *FileInfo, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/FileTransferService/getFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTransferServiceClient) TransferFile(ctx context.Context, opts ...grpc.CallOption) (FileTransferService_TransferFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileTransferService_serviceDesc.Streams[0], "/FileTransferService/transferFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileTransferServiceTransferFileClient{stream}
	return x, nil
}

type FileTransferService_TransferFileClient interface {
	Send(*FileData) error
	CloseAndRecv() (*Res, error)
	grpc.ClientStream
}

type fileTransferServiceTransferFileClient struct {
	grpc.ClientStream
}

func (x *fileTransferServiceTransferFileClient) Send(m *FileData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileTransferServiceTransferFileClient) CloseAndRecv() (*Res, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileTransferServiceServer is the server API for FileTransferService service.
type FileTransferServiceServer interface {
	GetFolderInfo(context.Context, *FolderInfo) (*Res, error)
	GetFileInfo(context.Context, *FileInfo) (*Res, error)
	TransferFile(FileTransferService_TransferFileServer) error
}

func RegisterFileTransferServiceServer(s *grpc.Server, srv FileTransferServiceServer) {
	s.RegisterService(&_FileTransferService_serviceDesc, srv)
}

func _FileTransferService_GetFolderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTransferServiceServer).GetFolderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileTransferService/GetFolderInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTransferServiceServer).GetFolderInfo(ctx, req.(*FolderInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTransferService_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTransferServiceServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileTransferService/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTransferServiceServer).GetFileInfo(ctx, req.(*FileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTransferService_TransferFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileTransferServiceServer).TransferFile(&fileTransferServiceTransferFileServer{stream})
}

type FileTransferService_TransferFileServer interface {
	SendAndClose(*Res) error
	Recv() (*FileData, error)
	grpc.ServerStream
}

type fileTransferServiceTransferFileServer struct {
	grpc.ServerStream
}

func (x *fileTransferServiceTransferFileServer) SendAndClose(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileTransferServiceTransferFileServer) Recv() (*FileData, error) {
	m := new(FileData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileTransferService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileTransferService",
	HandlerType: (*FileTransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getFolderInfo",
			Handler:    _FileTransferService_GetFolderInfo_Handler,
		},
		{
			MethodName: "getFileInfo",
			Handler:    _FileTransferService_GetFileInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "transferFile",
			Handler:       _FileTransferService_TransferFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_b67e7327b0ef4312) }

var fileDescriptor_service_b67e7327b0ef4312 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xc5, 0x20,
	0x14, 0x85, 0xc5, 0x36, 0x6a, 0x6f, 0xdb, 0x05, 0x17, 0xe2, 0xa0, 0x04, 0x1d, 0x98, 0x18, 0xf4,
	0x2f, 0x98, 0x26, 0xae, 0xe8, 0x1f, 0x40, 0x7b, 0xdb, 0x34, 0x69, 0x8b, 0x01, 0xe2, 0xe0, 0xe8,
	0x2f, 0x37, 0xd0, 0xf2, 0xde, 0x5b, 0xde, 0xf6, 0xdd, 0x7b, 0x4e, 0x0e, 0x9c, 0x0b, 0xad, 0x47,
	0xf7, 0x33, 0x7d, 0xa1, 0xfa, 0x76, 0x36, 0x58, 0xc1, 0x01, 0x3a, 0x3b, 0xf7, 0xe8, 0xde, 0xd6,
	0xc1, 0x52, 0x0a, 0xe5, 0x6a, 0x16, 0x64, 0x84, 0x13, 0x59, 0xe9, 0xc4, 0xa2, 0x83, 0x9b, 0x6e,
	0x9a, 0xf1, 0x9c, 0x1e, 0x77, 0x7e, 0xfa, 0x45, 0x76, 0xc9, 0x89, 0x2c, 0x74, 0xe2, 0xb8, 0x5b,
	0x6c, 0x8f, 0xac, 0xe0, 0x44, 0xb6, 0x3a, 0xb1, 0xb8, 0xdf, 0x72, 0x5e, 0x4d, 0x30, 0x51, 0xef,
	0x4d, 0x30, 0x29, 0xa7, 0xd1, 0x89, 0xc5, 0x03, 0x14, 0x1a, 0x3d, 0x65, 0x70, 0xbd, 0xa0, 0xf7,
	0x66, 0xcc, 0xaf, 0xe4, 0xf1, 0xf9, 0x8f, 0xc0, 0x6d, 0x4c, 0xf8, 0x70, 0x66, 0xf5, 0x03, 0xba,
	0xf7, 0xad, 0x08, 0x7d, 0x82, 0x76, 0xc4, 0x70, 0xd2, 0xa2, 0x56, 0xc7, 0xe1, 0xae, 0x54, 0x1a,
	0xbd, 0xb8, 0xa0, 0x1c, 0xea, 0xe8, 0xca, 0x4d, 0x2a, 0x95, 0xf1, 0xe0, 0x78, 0x84, 0x26, 0xec,
	0xd1, 0x51, 0xdb, 0x2d, 0xf1, 0xbf, 0xd9, 0x22, 0xc9, 0xe7, 0x55, 0x3a, 0xdb, 0xcb, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x05, 0xad, 0x55, 0x78, 0x47, 0x01, 0x00, 0x00,
}
