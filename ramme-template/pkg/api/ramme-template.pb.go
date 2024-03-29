// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: ramme-template.proto

package api

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ramme_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ramme_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_ramme_template_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ramme_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ramme_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_ramme_template_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PingResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_ramme_template_proto protoreflect.FileDescriptor

var file_ramme_template_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x5f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x67, 0x0a, 0x0d, 0x52, 0x61, 0x6d, 0x6d, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x56, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x72,
	0x61, 0x6d, 0x6d, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x6d, 0x6d,
	0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22,
	0x08, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ramme_template_proto_rawDescOnce sync.Once
	file_ramme_template_proto_rawDescData = file_ramme_template_proto_rawDesc
)

func file_ramme_template_proto_rawDescGZIP() []byte {
	file_ramme_template_proto_rawDescOnce.Do(func() {
		file_ramme_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_ramme_template_proto_rawDescData)
	})
	return file_ramme_template_proto_rawDescData
}

var file_ramme_template_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ramme_template_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: ramme_template.PingRequest
	(*PingResponse)(nil), // 1: ramme_template.PingResponse
}
var file_ramme_template_proto_depIdxs = []int32{
	0, // 0: ramme_template.RammeTemplate.Ping:input_type -> ramme_template.PingRequest
	1, // 1: ramme_template.RammeTemplate.Ping:output_type -> ramme_template.PingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ramme_template_proto_init() }
func file_ramme_template_proto_init() {
	if File_ramme_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ramme_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_ramme_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_ramme_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ramme_template_proto_goTypes,
		DependencyIndexes: file_ramme_template_proto_depIdxs,
		MessageInfos:      file_ramme_template_proto_msgTypes,
	}.Build()
	File_ramme_template_proto = out.File
	file_ramme_template_proto_rawDesc = nil
	file_ramme_template_proto_goTypes = nil
	file_ramme_template_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RammeTemplateClient is the client API for RammeTemplate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RammeTemplateClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type rammeTemplateClient struct {
	cc grpc.ClientConnInterface
}

func NewRammeTemplateClient(cc grpc.ClientConnInterface) RammeTemplateClient {
	return &rammeTemplateClient{cc}
}

func (c *rammeTemplateClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/ramme_template.RammeTemplate/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RammeTemplateServer is the server API for RammeTemplate service.
type RammeTemplateServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedRammeTemplateServer can be embedded to have forward compatible implementations.
type UnimplementedRammeTemplateServer struct {
}

func (*UnimplementedRammeTemplateServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterRammeTemplateServer(s *grpc.Server, srv RammeTemplateServer) {
	s.RegisterService(&_RammeTemplate_serviceDesc, srv)
}

func _RammeTemplate_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RammeTemplateServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ramme_template.RammeTemplate/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RammeTemplateServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RammeTemplate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ramme_template.RammeTemplate",
	HandlerType: (*RammeTemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RammeTemplate_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ramme-template.proto",
}
