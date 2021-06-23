/// Defines simple messages to run the logsink server and client

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: logsink.proto

package logsink

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

/// The single line message
type LineMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line     string `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
	Priority int32  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *LineMessage) Reset() {
	*x = LineMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logsink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineMessage) ProtoMessage() {}

func (x *LineMessage) ProtoReflect() protoreflect.Message {
	mi := &file_logsink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineMessage.ProtoReflect.Descriptor instead.
func (*LineMessage) Descriptor() ([]byte, []int) {
	return file_logsink_proto_rawDescGZIP(), []int{0}
}

func (x *LineMessage) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

func (x *LineMessage) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

/// Result of the log send
type LineResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LineResult) Reset() {
	*x = LineResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logsink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineResult) ProtoMessage() {}

func (x *LineResult) ProtoReflect() protoreflect.Message {
	mi := &file_logsink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineResult.ProtoReflect.Descriptor instead.
func (*LineResult) Descriptor() ([]byte, []int) {
	return file_logsink_proto_rawDescGZIP(), []int{1}
}

func (x *LineResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_logsink_proto protoreflect.FileDescriptor

var file_logsink_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6c, 0x6f, 0x67, 0x73, 0x69, 0x6e, 0x6b, 0x22, 0x3d, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x24, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x46, 0x0a,
	0x0b, 0x4c, 0x6f, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x69,
	0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13,
	0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logsink_proto_rawDescOnce sync.Once
	file_logsink_proto_rawDescData = file_logsink_proto_rawDesc
)

func file_logsink_proto_rawDescGZIP() []byte {
	file_logsink_proto_rawDescOnce.Do(func() {
		file_logsink_proto_rawDescData = protoimpl.X.CompressGZIP(file_logsink_proto_rawDescData)
	})
	return file_logsink_proto_rawDescData
}

var file_logsink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logsink_proto_goTypes = []interface{}{
	(*LineMessage)(nil), // 0: logsink.LineMessage
	(*LineResult)(nil),  // 1: logsink.LineResult
}
var file_logsink_proto_depIdxs = []int32{
	0, // 0: logsink.LogTransfer.SendLine:input_type -> logsink.LineMessage
	1, // 1: logsink.LogTransfer.SendLine:output_type -> logsink.LineResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logsink_proto_init() }
func file_logsink_proto_init() {
	if File_logsink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logsink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineMessage); i {
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
		file_logsink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineResult); i {
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
			RawDescriptor: file_logsink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logsink_proto_goTypes,
		DependencyIndexes: file_logsink_proto_depIdxs,
		MessageInfos:      file_logsink_proto_msgTypes,
	}.Build()
	File_logsink_proto = out.File
	file_logsink_proto_rawDesc = nil
	file_logsink_proto_goTypes = nil
	file_logsink_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogTransferClient is the client API for LogTransfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogTransferClient interface {
	/// SendLine sends a log line to the server
	SendLine(ctx context.Context, in *LineMessage, opts ...grpc.CallOption) (*LineResult, error)
}

type logTransferClient struct {
	cc grpc.ClientConnInterface
}

func NewLogTransferClient(cc grpc.ClientConnInterface) LogTransferClient {
	return &logTransferClient{cc}
}

func (c *logTransferClient) SendLine(ctx context.Context, in *LineMessage, opts ...grpc.CallOption) (*LineResult, error) {
	out := new(LineResult)
	err := c.cc.Invoke(ctx, "/logsink.LogTransfer/SendLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogTransferServer is the server API for LogTransfer service.
type LogTransferServer interface {
	/// SendLine sends a log line to the server
	SendLine(context.Context, *LineMessage) (*LineResult, error)
}

// UnimplementedLogTransferServer can be embedded to have forward compatible implementations.
type UnimplementedLogTransferServer struct {
}

func (*UnimplementedLogTransferServer) SendLine(context.Context, *LineMessage) (*LineResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLine not implemented")
}

func RegisterLogTransferServer(s *grpc.Server, srv LogTransferServer) {
	s.RegisterService(&_LogTransfer_serviceDesc, srv)
}

func _LogTransfer_SendLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogTransferServer).SendLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logsink.LogTransfer/SendLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogTransferServer).SendLine(ctx, req.(*LineMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogTransfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logsink.LogTransfer",
	HandlerType: (*LogTransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLine",
			Handler:    _LogTransfer_SendLine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logsink.proto",
}
