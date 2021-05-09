// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: sms-sender-service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_sms_sender_service_proto protoreflect.FileDescriptor

var file_sms_sender_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x6d, 0x73, 0x2d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x6d, 0x73, 0x2d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc2,
	0x01, 0x0a, 0x10, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x54, 0x6f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x12,
	0x0f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x41, 0x63, 0x6b, 0x54, 0x6f, 0x53, 0x6d, 0x73, 0x12, 0x10,
	0x2e, 0x41, 0x63, 0x6b, 0x54, 0x6f, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x41, 0x63, 0x6b, 0x54, 0x6f, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x30, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x61, 0x70, 0x61, 0x6e,
	0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x61, 0x76, 0x79,
	0x61, 0x70, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a,
	0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_sms_sender_service_proto_goTypes = []interface{}{
	(*ReadyToSendMessageRequest)(nil),  // 0: ReadyToSendMessageRequest
	(*SendSmsRequest)(nil),             // 1: SendSmsRequest
	(*AckToSmsRequest)(nil),            // 2: AckToSmsRequest
	(*ReadyToSendMessageResponse)(nil), // 3: ReadyToSendMessageResponse
	(*SendSmsResponse)(nil),            // 4: SendSmsResponse
	(*AckToSmsResponse)(nil),           // 5: AckToSmsResponse
}
var file_sms_sender_service_proto_depIdxs = []int32{
	0, // 0: SmsSenderService.ReadyToSendMessage:input_type -> ReadyToSendMessageRequest
	1, // 1: SmsSenderService.SendSms:input_type -> SendSmsRequest
	2, // 2: SmsSenderService.AckToSms:input_type -> AckToSmsRequest
	3, // 3: SmsSenderService.ReadyToSendMessage:output_type -> ReadyToSendMessageResponse
	4, // 4: SmsSenderService.SendSms:output_type -> SendSmsResponse
	5, // 5: SmsSenderService.AckToSms:output_type -> AckToSmsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_sender_service_proto_init() }
func file_sms_sender_service_proto_init() {
	if File_sms_sender_service_proto != nil {
		return
	}
	file_sms_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sms_sender_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_sender_service_proto_goTypes,
		DependencyIndexes: file_sms_sender_service_proto_depIdxs,
	}.Build()
	File_sms_sender_service_proto = out.File
	file_sms_sender_service_proto_rawDesc = nil
	file_sms_sender_service_proto_goTypes = nil
	file_sms_sender_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmsSenderServiceClient is the client API for SmsSenderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmsSenderServiceClient interface {
	ReadyToSendMessage(ctx context.Context, in *ReadyToSendMessageRequest, opts ...grpc.CallOption) (SmsSenderService_ReadyToSendMessageClient, error)
	SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error)
	AckToSms(ctx context.Context, in *AckToSmsRequest, opts ...grpc.CallOption) (*AckToSmsResponse, error)
}

type smsSenderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsSenderServiceClient(cc grpc.ClientConnInterface) SmsSenderServiceClient {
	return &smsSenderServiceClient{cc}
}

func (c *smsSenderServiceClient) ReadyToSendMessage(ctx context.Context, in *ReadyToSendMessageRequest, opts ...grpc.CallOption) (SmsSenderService_ReadyToSendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SmsSenderService_serviceDesc.Streams[0], "/SmsSenderService/ReadyToSendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &smsSenderServiceReadyToSendMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmsSenderService_ReadyToSendMessageClient interface {
	Recv() (*ReadyToSendMessageResponse, error)
	grpc.ClientStream
}

type smsSenderServiceReadyToSendMessageClient struct {
	grpc.ClientStream
}

func (x *smsSenderServiceReadyToSendMessageClient) Recv() (*ReadyToSendMessageResponse, error) {
	m := new(ReadyToSendMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smsSenderServiceClient) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error) {
	out := new(SendSmsResponse)
	err := c.cc.Invoke(ctx, "/SmsSenderService/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsSenderServiceClient) AckToSms(ctx context.Context, in *AckToSmsRequest, opts ...grpc.CallOption) (*AckToSmsResponse, error) {
	out := new(AckToSmsResponse)
	err := c.cc.Invoke(ctx, "/SmsSenderService/AckToSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsSenderServiceServer is the server API for SmsSenderService service.
type SmsSenderServiceServer interface {
	ReadyToSendMessage(*ReadyToSendMessageRequest, SmsSenderService_ReadyToSendMessageServer) error
	SendSms(context.Context, *SendSmsRequest) (*SendSmsResponse, error)
	AckToSms(context.Context, *AckToSmsRequest) (*AckToSmsResponse, error)
}

// UnimplementedSmsSenderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmsSenderServiceServer struct {
}

func (*UnimplementedSmsSenderServiceServer) ReadyToSendMessage(*ReadyToSendMessageRequest, SmsSenderService_ReadyToSendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadyToSendMessage not implemented")
}
func (*UnimplementedSmsSenderServiceServer) SendSms(context.Context, *SendSmsRequest) (*SendSmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (*UnimplementedSmsSenderServiceServer) AckToSms(context.Context, *AckToSmsRequest) (*AckToSmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckToSms not implemented")
}

func RegisterSmsSenderServiceServer(s *grpc.Server, srv SmsSenderServiceServer) {
	s.RegisterService(&_SmsSenderService_serviceDesc, srv)
}

func _SmsSenderService_ReadyToSendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadyToSendMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmsSenderServiceServer).ReadyToSendMessage(m, &smsSenderServiceReadyToSendMessageServer{stream})
}

type SmsSenderService_ReadyToSendMessageServer interface {
	Send(*ReadyToSendMessageResponse) error
	grpc.ServerStream
}

type smsSenderServiceReadyToSendMessageServer struct {
	grpc.ServerStream
}

func (x *smsSenderServiceReadyToSendMessageServer) Send(m *ReadyToSendMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SmsSenderService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsSenderServiceServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmsSenderService/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsSenderServiceServer).SendSms(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsSenderService_AckToSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckToSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsSenderServiceServer).AckToSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmsSenderService/AckToSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsSenderServiceServer).AckToSms(ctx, req.(*AckToSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmsSenderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SmsSenderService",
	HandlerType: (*SmsSenderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _SmsSenderService_SendSms_Handler,
		},
		{
			MethodName: "AckToSms",
			Handler:    _SmsSenderService_AckToSms_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadyToSendMessage",
			Handler:       _SmsSenderService_ReadyToSendMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sms-sender-service.proto",
}
