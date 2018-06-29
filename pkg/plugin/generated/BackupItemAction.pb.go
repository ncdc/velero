// Code generated by protoc-gen-go. DO NOT EDIT.
// source: BackupItemAction.proto

/*
Package generated is a generated protocol buffer package.

It is generated from these files:
	BackupItemAction.proto
	BlockStore.proto
	ObjectStore.proto
	RestoreItemAction.proto
	Shared.proto

It has these top-level messages:
	ExecuteRequest
	ExecuteResponse
	ResourceIdentifier
	CreateVolumeRequest
	CreateVolumeResponse
	GetVolumeInfoRequest
	GetVolumeInfoResponse
	IsVolumeReadyRequest
	IsVolumeReadyResponse
	CreateSnapshotRequest
	CreateSnapshotResponse
	DeleteSnapshotRequest
	GetVolumeIDRequest
	GetVolumeIDResponse
	SetVolumeIDRequest
	SetVolumeIDResponse
	PutObjectRequest
	ObjectExistsRequest
	ObjectExistsResponse
	GetObjectRequest
	Bytes
	ListCommonPrefixesRequest
	ListCommonPrefixesResponse
	ListObjectsRequest
	ListObjectsResponse
	DeleteObjectRequest
	CreateSignedURLRequest
	CreateSignedURLResponse
	RestoreExecuteRequest
	RestoreExecuteResponse
	Empty
	InitRequest
	AppliesToResponse
*/
package generated

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

type ExecuteRequest struct {
	Item   []byte `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Backup []byte `protobuf:"bytes,2,opt,name=backup,proto3" json:"backup,omitempty"`
}

func (m *ExecuteRequest) Reset()                    { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()               {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExecuteRequest) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ExecuteRequest) GetBackup() []byte {
	if m != nil {
		return m.Backup
	}
	return nil
}

type ExecuteResponse struct {
	Item            []byte                `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	AdditionalItems []*ResourceIdentifier `protobuf:"bytes,2,rep,name=additionalItems" json:"additionalItems,omitempty"`
}

func (m *ExecuteResponse) Reset()                    { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()               {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecuteResponse) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ExecuteResponse) GetAdditionalItems() []*ResourceIdentifier {
	if m != nil {
		return m.AdditionalItems
	}
	return nil
}

type ResourceIdentifier struct {
	Group     string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Resource  string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *ResourceIdentifier) Reset()                    { *m = ResourceIdentifier{} }
func (m *ResourceIdentifier) String() string            { return proto.CompactTextString(m) }
func (*ResourceIdentifier) ProtoMessage()               {}
func (*ResourceIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResourceIdentifier) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *ResourceIdentifier) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *ResourceIdentifier) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ResourceIdentifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecuteRequest)(nil), "generated.ExecuteRequest")
	proto.RegisterType((*ExecuteResponse)(nil), "generated.ExecuteResponse")
	proto.RegisterType((*ResourceIdentifier)(nil), "generated.ResourceIdentifier")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BackupItemAction service

type BackupItemActionClient interface {
	AppliesTo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppliesToResponse, error)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type backupItemActionClient struct {
	cc *grpc.ClientConn
}

func NewBackupItemActionClient(cc *grpc.ClientConn) BackupItemActionClient {
	return &backupItemActionClient{cc}
}

func (c *backupItemActionClient) AppliesTo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppliesToResponse, error) {
	out := new(AppliesToResponse)
	err := grpc.Invoke(ctx, "/generated.BackupItemAction/AppliesTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupItemActionClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := grpc.Invoke(ctx, "/generated.BackupItemAction/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BackupItemAction service

type BackupItemActionServer interface {
	AppliesTo(context.Context, *Empty) (*AppliesToResponse, error)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

func RegisterBackupItemActionServer(s *grpc.Server, srv BackupItemActionServer) {
	s.RegisterService(&_BackupItemAction_serviceDesc, srv)
}

func _BackupItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BackupItemAction/AppliesTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).AppliesTo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BackupItemAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupItemAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.BackupItemAction",
	HandlerType: (*BackupItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _BackupItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _BackupItemAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BackupItemAction.proto",
}

func init() { proto.RegisterFile("BackupItemAction.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x4d, 0x01, 0xd1, 0x8e, 0x44, 0xc8, 0xc4, 0x90, 0xda, 0x60, 0x42, 0x7a, 0xe2, 0xd4, 0x03,
	0x1e, 0xf5, 0x20, 0x26, 0xc4, 0x70, 0x5d, 0xfd, 0x81, 0xa5, 0x1d, 0x71, 0x23, 0xdd, 0x5d, 0x77,
	0xb7, 0x09, 0x7e, 0x86, 0x7f, 0x6c, 0xba, 0x6d, 0x6a, 0x45, 0x6e, 0xfb, 0xe6, 0xcd, 0x9b, 0x7d,
	0xf3, 0x06, 0xa6, 0x4f, 0x3c, 0xfb, 0x28, 0xf5, 0xc6, 0x51, 0xb1, 0xca, 0x9c, 0x50, 0x32, 0xd5,
	0x46, 0x39, 0x85, 0xe1, 0x8e, 0x24, 0x19, 0xee, 0x28, 0x8f, 0x47, 0x2f, 0xef, 0xdc, 0x50, 0x5e,
	0x13, 0xc9, 0x03, 0x5c, 0xad, 0x0f, 0x94, 0x95, 0x8e, 0x18, 0x7d, 0x96, 0x64, 0x1d, 0x22, 0x0c,
	0x84, 0xa3, 0x22, 0x0a, 0xe6, 0xc1, 0x62, 0xc4, 0xfc, 0x1b, 0xa7, 0x30, 0xdc, 0xfa, 0xc1, 0x51,
	0xcf, 0x57, 0x1b, 0x94, 0x48, 0x18, 0xb7, 0x6a, 0xab, 0x95, 0xb4, 0x74, 0x52, 0xfe, 0x0c, 0x63,
	0x9e, 0xe7, 0xa2, 0xf2, 0xc3, 0xf7, 0x95, 0x37, 0x1b, 0xf5, 0xe6, 0xfd, 0xc5, 0xe5, 0xf2, 0x36,
	0x6d, 0x7d, 0xa5, 0x8c, 0xac, 0x2a, 0x4d, 0x46, 0x9b, 0x9c, 0xa4, 0x13, 0x6f, 0x82, 0x0c, 0x3b,
	0x56, 0x25, 0x07, 0xc0, 0xff, 0x6d, 0x78, 0x0d, 0x67, 0x3b, 0xa3, 0x4a, 0xed, 0xff, 0x0c, 0x59,
	0x0d, 0x30, 0x86, 0x0b, 0xd3, 0xf4, 0x7a, 0xd7, 0x21, 0x6b, 0x31, 0xce, 0x20, 0x94, 0xbc, 0x20,
	0xab, 0x79, 0x46, 0x51, 0xdf, 0x93, 0xbf, 0x85, 0x6a, 0x85, 0x0a, 0x44, 0x03, 0x4f, 0xf8, 0xf7,
	0xf2, 0x3b, 0x80, 0xc9, 0x71, 0xb6, 0x78, 0x0f, 0xe1, 0x4a, 0xeb, 0xbd, 0x20, 0xfb, 0xaa, 0x70,
	0xd2, 0xd9, 0x65, 0x5d, 0x68, 0xf7, 0x15, 0xcf, 0x3a, 0x95, 0xb6, 0xaf, 0x0d, 0xea, 0x11, 0xce,
	0x9b, 0xec, 0xf0, 0xa6, 0x2b, 0xfd, 0x73, 0x8d, 0x38, 0x3e, 0x45, 0xd5, 0x13, 0xb6, 0x43, 0x7f,
	0xc2, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x09, 0x4d, 0x36, 0xf5, 0x01, 0x00, 0x00,
}
