// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/sacloud/libsacloud/v2/grpc/sacloud/auth_status.proto

package sacloud // import "github.com/sacloud/libsacloud/v2/grpc/sacloud"

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	types "github.com/sacloud/libsacloud/v2/grpc/sacloud/types"

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

type AuthStatus struct {
	AccountID            int64                  `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	AccountName          string                 `protobuf:"bytes,2,opt,name=accountName,proto3" json:"accountName,omitempty"`
	AccountCode          string                 `protobuf:"bytes,3,opt,name=accountCode,proto3" json:"accountCode,omitempty"`
	AccountClass         string                 `protobuf:"bytes,4,opt,name=accountClass,proto3" json:"accountClass,omitempty"`
	MemberCode           string                 `protobuf:"bytes,5,opt,name=memberCode,proto3" json:"memberCode,omitempty"`
	MemberClass          string                 `protobuf:"bytes,6,opt,name=memberClass,proto3" json:"memberClass,omitempty"`
	AuthClass            types.AuthClass        `protobuf:"varint,7,opt,name=authClass,proto3,enum=types.AuthClass" json:"authClass,omitempty"`
	AuthMethod           types.AuthMethod       `protobuf:"varint,8,opt,name=authMethod,proto3,enum=types.AuthMethod" json:"authMethod,omitempty"`
	IsAPIKey             bool                   `protobuf:"varint,9,opt,name=isAPIKey,proto3" json:"isAPIKey,omitempty"`
	ExternalPermission   string                 `protobuf:"bytes,10,opt,name=externalPermission,proto3" json:"externalPermission,omitempty"`
	OperationPenalty     types.OperationPenalty `protobuf:"varint,11,opt,name=operationPenalty,proto3,enum=types.OperationPenalty" json:"operationPenalty,omitempty"`
	Permission           types.APIKeyPermission `protobuf:"varint,12,opt,name=permission,proto3,enum=types.APIKeyPermission" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AuthStatus) Reset()         { *m = AuthStatus{} }
func (m *AuthStatus) String() string { return proto.CompactTextString(m) }
func (*AuthStatus) ProtoMessage()    {}
func (*AuthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_status_9b1472e91fe2e9a3, []int{0}
}
func (m *AuthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStatus.Unmarshal(m, b)
}
func (m *AuthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStatus.Marshal(b, m, deterministic)
}
func (dst *AuthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatus.Merge(dst, src)
}
func (m *AuthStatus) XXX_Size() int {
	return xxx_messageInfo_AuthStatus.Size(m)
}
func (m *AuthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatus proto.InternalMessageInfo

func (m *AuthStatus) GetAccountID() int64 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *AuthStatus) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *AuthStatus) GetAccountCode() string {
	if m != nil {
		return m.AccountCode
	}
	return ""
}

func (m *AuthStatus) GetAccountClass() string {
	if m != nil {
		return m.AccountClass
	}
	return ""
}

func (m *AuthStatus) GetMemberCode() string {
	if m != nil {
		return m.MemberCode
	}
	return ""
}

func (m *AuthStatus) GetMemberClass() string {
	if m != nil {
		return m.MemberClass
	}
	return ""
}

func (m *AuthStatus) GetAuthClass() types.AuthClass {
	if m != nil {
		return m.AuthClass
	}
	return types.AuthClass_AUTH_CLASS_UNSPECIFIED
}

func (m *AuthStatus) GetAuthMethod() types.AuthMethod {
	if m != nil {
		return m.AuthMethod
	}
	return types.AuthMethod_AUTH_METHOD_UNSPECIFIED
}

func (m *AuthStatus) GetIsAPIKey() bool {
	if m != nil {
		return m.IsAPIKey
	}
	return false
}

func (m *AuthStatus) GetExternalPermission() string {
	if m != nil {
		return m.ExternalPermission
	}
	return ""
}

func (m *AuthStatus) GetOperationPenalty() types.OperationPenalty {
	if m != nil {
		return m.OperationPenalty
	}
	return types.OperationPenalty_OPERATION_PENALTY_UNSPECIFIED
}

func (m *AuthStatus) GetPermission() types.APIKeyPermission {
	if m != nil {
		return m.Permission
	}
	return types.APIKeyPermission_API_KEY_PERMISSION_UNSPECIFIED
}

type AuthStatusReadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthStatusReadRequest) Reset()         { *m = AuthStatusReadRequest{} }
func (m *AuthStatusReadRequest) String() string { return proto.CompactTextString(m) }
func (*AuthStatusReadRequest) ProtoMessage()    {}
func (*AuthStatusReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_status_9b1472e91fe2e9a3, []int{1}
}
func (m *AuthStatusReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStatusReadRequest.Unmarshal(m, b)
}
func (m *AuthStatusReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStatusReadRequest.Marshal(b, m, deterministic)
}
func (dst *AuthStatusReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusReadRequest.Merge(dst, src)
}
func (m *AuthStatusReadRequest) XXX_Size() int {
	return xxx_messageInfo_AuthStatusReadRequest.Size(m)
}
func (m *AuthStatusReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusReadRequest proto.InternalMessageInfo

type AuthStatusReadResult struct {
	AuthStatus           *AuthStatus `protobuf:"bytes,1,opt,name=authStatus,proto3" json:"authStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AuthStatusReadResult) Reset()         { *m = AuthStatusReadResult{} }
func (m *AuthStatusReadResult) String() string { return proto.CompactTextString(m) }
func (*AuthStatusReadResult) ProtoMessage()    {}
func (*AuthStatusReadResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_status_9b1472e91fe2e9a3, []int{2}
}
func (m *AuthStatusReadResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStatusReadResult.Unmarshal(m, b)
}
func (m *AuthStatusReadResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStatusReadResult.Marshal(b, m, deterministic)
}
func (dst *AuthStatusReadResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusReadResult.Merge(dst, src)
}
func (m *AuthStatusReadResult) XXX_Size() int {
	return xxx_messageInfo_AuthStatusReadResult.Size(m)
}
func (m *AuthStatusReadResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusReadResult.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusReadResult proto.InternalMessageInfo

func (m *AuthStatusReadResult) GetAuthStatus() *AuthStatus {
	if m != nil {
		return m.AuthStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthStatus)(nil), "sacloud.AuthStatus")
	proto.RegisterType((*AuthStatusReadRequest)(nil), "sacloud.AuthStatusReadRequest")
	proto.RegisterType((*AuthStatusReadResult)(nil), "sacloud.AuthStatusReadResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthStatusAPIClient is the client API for AuthStatusAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthStatusAPIClient interface {
	Read(ctx context.Context, in *AuthStatusReadRequest, opts ...grpc.CallOption) (*AuthStatusReadResult, error)
}

type authStatusAPIClient struct {
	cc *grpc.ClientConn
}

func NewAuthStatusAPIClient(cc *grpc.ClientConn) AuthStatusAPIClient {
	return &authStatusAPIClient{cc}
}

func (c *authStatusAPIClient) Read(ctx context.Context, in *AuthStatusReadRequest, opts ...grpc.CallOption) (*AuthStatusReadResult, error) {
	out := new(AuthStatusReadResult)
	err := c.cc.Invoke(ctx, "/sacloud.AuthStatusAPI/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthStatusAPIServer is the server API for AuthStatusAPI service.
type AuthStatusAPIServer interface {
	Read(context.Context, *AuthStatusReadRequest) (*AuthStatusReadResult, error)
}

func RegisterAuthStatusAPIServer(s *grpc.Server, srv AuthStatusAPIServer) {
	s.RegisterService(&_AuthStatusAPI_serviceDesc, srv)
}

func _AuthStatusAPI_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthStatusReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthStatusAPIServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sacloud.AuthStatusAPI/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthStatusAPIServer).Read(ctx, req.(*AuthStatusReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthStatusAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sacloud.AuthStatusAPI",
	HandlerType: (*AuthStatusAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _AuthStatusAPI_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/sacloud/libsacloud/v2/grpc/sacloud/auth_status.proto",
}

func init() {
	proto.RegisterFile("github.com/sacloud/libsacloud/v2/grpc/sacloud/auth_status.proto", fileDescriptor_auth_status_9b1472e91fe2e9a3)
}

var fileDescriptor_auth_status_9b1472e91fe2e9a3 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0x58, 0xd9, 0x9a, 0xd7, 0x81, 0x86, 0x01, 0xcd, 0xaa, 0x60, 0x8a, 0x72, 0xca, 0x85,
	0x44, 0x64, 0x07, 0x8e, 0x28, 0x0c, 0x0e, 0xd5, 0x04, 0x44, 0x46, 0xe2, 0xc0, 0x05, 0x39, 0x89,
	0xd5, 0x44, 0x4a, 0xe2, 0x10, 0xdb, 0x88, 0x1e, 0xf9, 0xcf, 0x51, 0xec, 0xb4, 0x71, 0x4b, 0x77,
	0xe8, 0xa5, 0xea, 0xfb, 0x7e, 0xbd, 0xaf, 0xd6, 0x2b, 0xbc, 0x5f, 0x57, 0xb2, 0x54, 0x59, 0x98,
	0xf3, 0x26, 0x12, 0x34, 0xaf, 0xb9, 0x2a, 0xa2, 0xba, 0xca, 0xb6, 0x5f, 0x7f, 0xc7, 0xd1, 0xba,
	0xef, 0xf2, 0x1d, 0x45, 0x95, 0x2c, 0x7f, 0x0a, 0x49, 0xa5, 0x12, 0x61, 0xd7, 0x73, 0xc9, 0xd1,
	0xc5, 0x48, 0x2d, 0x4f, 0x4c, 0x92, 0x9b, 0x8e, 0x09, 0xf3, 0x69, 0x92, 0xfc, 0xbf, 0x33, 0x80,
	0x44, 0xc9, 0xf2, 0x9b, 0x8e, 0x47, 0xaf, 0xc0, 0xa5, 0x79, 0xce, 0x55, 0x2b, 0x57, 0x1f, 0xb1,
	0xe3, 0x39, 0xc1, 0x19, 0x99, 0x00, 0xe4, 0xc1, 0x62, 0x1c, 0xbe, 0xd0, 0x86, 0xe1, 0x47, 0x9e,
	0x13, 0xb8, 0xc4, 0x86, 0x2c, 0xc5, 0x1d, 0x2f, 0x18, 0x3e, 0xdb, 0x53, 0x0c, 0x10, 0xf2, 0xe1,
	0x72, 0x3b, 0xd6, 0x54, 0x08, 0x3c, 0xd3, 0x92, 0x3d, 0x0c, 0xdd, 0x00, 0x34, 0xac, 0xc9, 0x58,
	0xaf, 0x43, 0x1e, 0x6b, 0x85, 0x85, 0x0c, 0x5b, 0xc6, 0x49, 0x47, 0x9c, 0x9b, 0x2d, 0x16, 0x84,
	0x42, 0x70, 0x87, 0x57, 0x33, 0xfc, 0x85, 0xe7, 0x04, 0x4f, 0xe3, 0xab, 0xd0, 0xfc, 0xee, 0x64,
	0x8b, 0x93, 0x49, 0x82, 0xde, 0x02, 0x0c, 0xc3, 0x67, 0x26, 0x4b, 0x5e, 0xe0, 0xb9, 0x36, 0x3c,
	0xb3, 0x0c, 0x86, 0x20, 0x96, 0x08, 0x2d, 0x61, 0x5e, 0x89, 0x24, 0x5d, 0xdd, 0xb3, 0x0d, 0x76,
	0x3d, 0x27, 0x98, 0x93, 0xdd, 0x8c, 0x42, 0x40, 0xec, 0x8f, 0x64, 0x7d, 0x4b, 0xeb, 0x94, 0xf5,
	0x4d, 0x25, 0x44, 0xc5, 0x5b, 0x0c, 0xba, 0xe7, 0x11, 0x06, 0xdd, 0xc1, 0x15, 0xef, 0x58, 0x4f,
	0x65, 0xc5, 0xdb, 0x94, 0xb5, 0xb4, 0x96, 0x1b, 0xbc, 0xd0, 0x25, 0xae, 0xc7, 0x12, 0x5f, 0x0f,
	0x68, 0xf2, 0x9f, 0x01, 0xbd, 0x03, 0xe8, 0xa6, 0x65, 0x97, 0x7b, 0x76, 0xd3, 0x6b, 0xda, 0x48,
	0x2c, 0xa9, 0x7f, 0x0d, 0x2f, 0xa7, 0x13, 0x20, 0x8c, 0x16, 0x84, 0xfd, 0x52, 0x4c, 0x48, 0xff,
	0x1e, 0x5e, 0x1c, 0x12, 0x42, 0xd5, 0x12, 0xdd, 0x9a, 0xd7, 0x32, 0xb8, 0x3e, 0x93, 0x45, 0xfc,
	0x3c, 0x1c, 0x8f, 0x2c, 0xb4, 0x2c, 0x96, 0x2c, 0xfe, 0x0e, 0x4f, 0x26, 0x26, 0x49, 0x57, 0xe8,
	0x13, 0xcc, 0x86, 0x4c, 0x74, 0x73, 0xcc, 0x39, 0xb5, 0x58, 0xbe, 0x7e, 0x90, 0x1f, 0xca, 0x7c,
	0x88, 0x7e, 0xbc, 0x39, 0xe9, 0x4f, 0x90, 0x9d, 0xeb, 0xcb, 0xbf, 0xfd, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x42, 0xae, 0xbe, 0x86, 0x03, 0x00, 0x00,
}