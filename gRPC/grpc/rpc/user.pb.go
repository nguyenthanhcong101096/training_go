// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: protobuf/user.proto

package rpc

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{0}
}

func (x *AccessToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{1}
}

func (x *Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type FromRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Passwotd  string `protobuf:"bytes,4,opt,name=passwotd,proto3" json:"passwotd,omitempty"`
}

func (x *FromRegister) Reset() {
	*x = FromRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromRegister) ProtoMessage() {}

func (x *FromRegister) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromRegister.ProtoReflect.Descriptor instead.
func (*FromRegister) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{2}
}

func (x *FromRegister) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *FromRegister) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *FromRegister) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FromRegister) GetPasswotd() string {
	if x != nil {
		return x.Passwotd
	}
	return ""
}

type RegisterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok   bool         `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *AccessToken `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RegisterResult) Reset() {
	*x = RegisterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResult) ProtoMessage() {}

func (x *RegisterResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResult.ProtoReflect.Descriptor instead.
func (*RegisterResult) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResult) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *RegisterResult) GetData() *AccessToken {
	if x != nil {
		return x.Data
	}
	return nil
}

type LoginResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok   bool         `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *AccessToken `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LoginResult) Reset() {
	*x = LoginResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResult) ProtoMessage() {}

func (x *LoginResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResult.ProtoReflect.Descriptor instead.
func (*LoginResult) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResult) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *LoginResult) GetData() *AccessToken {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protobuf_user_proto protoreflect.FileDescriptor

var file_protobuf_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x30, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x45, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x7c, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x74,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x74,
	0x64, 0x22, 0x46, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x0b, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x6f,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x72,
	0x6f, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_user_proto_rawDescOnce sync.Once
	file_protobuf_user_proto_rawDescData = file_protobuf_user_proto_rawDesc
)

func file_protobuf_user_proto_rawDescGZIP() []byte {
	file_protobuf_user_proto_rawDescOnce.Do(func() {
		file_protobuf_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_user_proto_rawDescData)
	})
	return file_protobuf_user_proto_rawDescData
}

var file_protobuf_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_user_proto_goTypes = []interface{}{
	(*AccessToken)(nil),    // 0: rpc.AccessToken
	(*Credentials)(nil),    // 1: rpc.Credentials
	(*FromRegister)(nil),   // 2: rpc.FromRegister
	(*RegisterResult)(nil), // 3: rpc.RegisterResult
	(*LoginResult)(nil),    // 4: rpc.LoginResult
}
var file_protobuf_user_proto_depIdxs = []int32{
	0, // 0: rpc.RegisterResult.data:type_name -> rpc.AccessToken
	0, // 1: rpc.LoginResult.data:type_name -> rpc.AccessToken
	1, // 2: rpc.User.UserLogin:input_type -> rpc.Credentials
	2, // 3: rpc.User.UserRegister:input_type -> rpc.FromRegister
	4, // 4: rpc.User.UserLogin:output_type -> rpc.LoginResult
	3, // 5: rpc.User.UserRegister:output_type -> rpc.RegisterResult
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_user_proto_init() }
func file_protobuf_user_proto_init() {
	if File_protobuf_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromRegister); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_user_proto_goTypes,
		DependencyIndexes: file_protobuf_user_proto_depIdxs,
		MessageInfos:      file_protobuf_user_proto_msgTypes,
	}.Build()
	File_protobuf_user_proto = out.File
	file_protobuf_user_proto_rawDesc = nil
	file_protobuf_user_proto_goTypes = nil
	file_protobuf_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	UserLogin(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginResult, error)
	UserRegister(ctx context.Context, in *FromRegister, opts ...grpc.CallOption) (*RegisterResult, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserLogin(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginResult, error) {
	out := new(LoginResult)
	err := c.cc.Invoke(ctx, "/rpc.User/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserRegister(ctx context.Context, in *FromRegister, opts ...grpc.CallOption) (*RegisterResult, error) {
	out := new(RegisterResult)
	err := c.cc.Invoke(ctx, "/rpc.User/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	UserLogin(context.Context, *Credentials) (*LoginResult, error)
	UserRegister(context.Context, *FromRegister) (*RegisterResult, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) UserLogin(context.Context, *Credentials) (*LoginResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserServer) UserRegister(context.Context, *FromRegister) (*RegisterResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FromRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRegister(ctx, req.(*FromRegister))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _User_UserRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/user.proto",
}
