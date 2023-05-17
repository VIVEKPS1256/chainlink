// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: median.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PluginMedianClient is the client API for PluginMedian service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginMedianClient interface {
	NewMedianFactory(ctx context.Context, in *NewMedianFactoryRequest, opts ...grpc.CallOption) (*NewMedianFactoryReply, error)
}

type pluginMedianClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginMedianClient(cc grpc.ClientConnInterface) PluginMedianClient {
	return &pluginMedianClient{cc}
}

func (c *pluginMedianClient) NewMedianFactory(ctx context.Context, in *NewMedianFactoryRequest, opts ...grpc.CallOption) (*NewMedianFactoryReply, error) {
	out := new(NewMedianFactoryReply)
	err := c.cc.Invoke(ctx, "/loop.PluginMedian/NewMedianFactory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginMedianServer is the server API for PluginMedian service.
// All implementations must embed UnimplementedPluginMedianServer
// for forward compatibility
type PluginMedianServer interface {
	NewMedianFactory(context.Context, *NewMedianFactoryRequest) (*NewMedianFactoryReply, error)
	mustEmbedUnimplementedPluginMedianServer()
}

// UnimplementedPluginMedianServer must be embedded to have forward compatible implementations.
type UnimplementedPluginMedianServer struct {
}

func (UnimplementedPluginMedianServer) NewMedianFactory(context.Context, *NewMedianFactoryRequest) (*NewMedianFactoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMedianFactory not implemented")
}
func (UnimplementedPluginMedianServer) mustEmbedUnimplementedPluginMedianServer() {}

// UnsafePluginMedianServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginMedianServer will
// result in compilation errors.
type UnsafePluginMedianServer interface {
	mustEmbedUnimplementedPluginMedianServer()
}

func RegisterPluginMedianServer(s grpc.ServiceRegistrar, srv PluginMedianServer) {
	s.RegisterService(&PluginMedian_ServiceDesc, srv)
}

func _PluginMedian_NewMedianFactory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMedianFactoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginMedianServer).NewMedianFactory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.PluginMedian/NewMedianFactory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginMedianServer).NewMedianFactory(ctx, req.(*NewMedianFactoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginMedian_ServiceDesc is the grpc.ServiceDesc for PluginMedian service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginMedian_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.PluginMedian",
	HandlerType: (*PluginMedianServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewMedianFactory",
			Handler:    _PluginMedian_NewMedianFactory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "median.proto",
}

// ErrorLogClient is the client API for ErrorLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErrorLogClient interface {
	SaveError(ctx context.Context, in *SaveErrorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type errorLogClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorLogClient(cc grpc.ClientConnInterface) ErrorLogClient {
	return &errorLogClient{cc}
}

func (c *errorLogClient) SaveError(ctx context.Context, in *SaveErrorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/loop.ErrorLog/SaveError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorLogServer is the server API for ErrorLog service.
// All implementations must embed UnimplementedErrorLogServer
// for forward compatibility
type ErrorLogServer interface {
	SaveError(context.Context, *SaveErrorRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedErrorLogServer()
}

// UnimplementedErrorLogServer must be embedded to have forward compatible implementations.
type UnimplementedErrorLogServer struct {
}

func (UnimplementedErrorLogServer) SaveError(context.Context, *SaveErrorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveError not implemented")
}
func (UnimplementedErrorLogServer) mustEmbedUnimplementedErrorLogServer() {}

// UnsafeErrorLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ErrorLogServer will
// result in compilation errors.
type UnsafeErrorLogServer interface {
	mustEmbedUnimplementedErrorLogServer()
}

func RegisterErrorLogServer(s grpc.ServiceRegistrar, srv ErrorLogServer) {
	s.RegisterService(&ErrorLog_ServiceDesc, srv)
}

func _ErrorLog_SaveError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorLogServer).SaveError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.ErrorLog/SaveError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorLogServer).SaveError(ctx, req.(*SaveErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ErrorLog_ServiceDesc is the grpc.ServiceDesc for ErrorLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ErrorLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.ErrorLog",
	HandlerType: (*ErrorLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveError",
			Handler:    _ErrorLog_SaveError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "median.proto",
}

// ReportCodecClient is the client API for ReportCodec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportCodecClient interface {
	BuildReport(ctx context.Context, in *BuildReportRequest, opts ...grpc.CallOption) (*BuildReportReply, error)
	MedianFromReport(ctx context.Context, in *MedianFromReportRequest, opts ...grpc.CallOption) (*MedianFromReportReply, error)
	MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error)
}

type reportCodecClient struct {
	cc grpc.ClientConnInterface
}

func NewReportCodecClient(cc grpc.ClientConnInterface) ReportCodecClient {
	return &reportCodecClient{cc}
}

func (c *reportCodecClient) BuildReport(ctx context.Context, in *BuildReportRequest, opts ...grpc.CallOption) (*BuildReportReply, error) {
	out := new(BuildReportReply)
	err := c.cc.Invoke(ctx, "/loop.ReportCodec/BuildReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecClient) MedianFromReport(ctx context.Context, in *MedianFromReportRequest, opts ...grpc.CallOption) (*MedianFromReportReply, error) {
	out := new(MedianFromReportReply)
	err := c.cc.Invoke(ctx, "/loop.ReportCodec/MedianFromReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecClient) MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error) {
	out := new(MaxReportLengthReply)
	err := c.cc.Invoke(ctx, "/loop.ReportCodec/MaxReportLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportCodecServer is the server API for ReportCodec service.
// All implementations must embed UnimplementedReportCodecServer
// for forward compatibility
type ReportCodecServer interface {
	BuildReport(context.Context, *BuildReportRequest) (*BuildReportReply, error)
	MedianFromReport(context.Context, *MedianFromReportRequest) (*MedianFromReportReply, error)
	MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error)
	mustEmbedUnimplementedReportCodecServer()
}

// UnimplementedReportCodecServer must be embedded to have forward compatible implementations.
type UnimplementedReportCodecServer struct {
}

func (UnimplementedReportCodecServer) BuildReport(context.Context, *BuildReportRequest) (*BuildReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildReport not implemented")
}
func (UnimplementedReportCodecServer) MedianFromReport(context.Context, *MedianFromReportRequest) (*MedianFromReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedianFromReport not implemented")
}
func (UnimplementedReportCodecServer) MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxReportLength not implemented")
}
func (UnimplementedReportCodecServer) mustEmbedUnimplementedReportCodecServer() {}

// UnsafeReportCodecServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportCodecServer will
// result in compilation errors.
type UnsafeReportCodecServer interface {
	mustEmbedUnimplementedReportCodecServer()
}

func RegisterReportCodecServer(s grpc.ServiceRegistrar, srv ReportCodecServer) {
	s.RegisterService(&ReportCodec_ServiceDesc, srv)
}

func _ReportCodec_BuildReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecServer).BuildReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.ReportCodec/BuildReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecServer).BuildReport(ctx, req.(*BuildReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodec_MedianFromReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedianFromReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecServer).MedianFromReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.ReportCodec/MedianFromReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecServer).MedianFromReport(ctx, req.(*MedianFromReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodec_MaxReportLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaxReportLengthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecServer).MaxReportLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.ReportCodec/MaxReportLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecServer).MaxReportLength(ctx, req.(*MaxReportLengthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportCodec_ServiceDesc is the grpc.ServiceDesc for ReportCodec service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportCodec_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.ReportCodec",
	HandlerType: (*ReportCodecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildReport",
			Handler:    _ReportCodec_BuildReport_Handler,
		},
		{
			MethodName: "MedianFromReport",
			Handler:    _ReportCodec_MedianFromReport_Handler,
		},
		{
			MethodName: "MaxReportLength",
			Handler:    _ReportCodec_MaxReportLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "median.proto",
}

// MedianContractClient is the client API for MedianContract service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedianContractClient interface {
	LatestTransmissionDetails(ctx context.Context, in *LatestTransmissionDetailsRequest, opts ...grpc.CallOption) (*LatestTransmissionDetailsReply, error)
	LatestRoundRequested(ctx context.Context, in *LatestRoundRequestedRequest, opts ...grpc.CallOption) (*LatestRoundRequestedReply, error)
}

type medianContractClient struct {
	cc grpc.ClientConnInterface
}

func NewMedianContractClient(cc grpc.ClientConnInterface) MedianContractClient {
	return &medianContractClient{cc}
}

func (c *medianContractClient) LatestTransmissionDetails(ctx context.Context, in *LatestTransmissionDetailsRequest, opts ...grpc.CallOption) (*LatestTransmissionDetailsReply, error) {
	out := new(LatestTransmissionDetailsReply)
	err := c.cc.Invoke(ctx, "/loop.MedianContract/LatestTransmissionDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medianContractClient) LatestRoundRequested(ctx context.Context, in *LatestRoundRequestedRequest, opts ...grpc.CallOption) (*LatestRoundRequestedReply, error) {
	out := new(LatestRoundRequestedReply)
	err := c.cc.Invoke(ctx, "/loop.MedianContract/LatestRoundRequested", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedianContractServer is the server API for MedianContract service.
// All implementations must embed UnimplementedMedianContractServer
// for forward compatibility
type MedianContractServer interface {
	LatestTransmissionDetails(context.Context, *LatestTransmissionDetailsRequest) (*LatestTransmissionDetailsReply, error)
	LatestRoundRequested(context.Context, *LatestRoundRequestedRequest) (*LatestRoundRequestedReply, error)
	mustEmbedUnimplementedMedianContractServer()
}

// UnimplementedMedianContractServer must be embedded to have forward compatible implementations.
type UnimplementedMedianContractServer struct {
}

func (UnimplementedMedianContractServer) LatestTransmissionDetails(context.Context, *LatestTransmissionDetailsRequest) (*LatestTransmissionDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestTransmissionDetails not implemented")
}
func (UnimplementedMedianContractServer) LatestRoundRequested(context.Context, *LatestRoundRequestedRequest) (*LatestRoundRequestedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestRoundRequested not implemented")
}
func (UnimplementedMedianContractServer) mustEmbedUnimplementedMedianContractServer() {}

// UnsafeMedianContractServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedianContractServer will
// result in compilation errors.
type UnsafeMedianContractServer interface {
	mustEmbedUnimplementedMedianContractServer()
}

func RegisterMedianContractServer(s grpc.ServiceRegistrar, srv MedianContractServer) {
	s.RegisterService(&MedianContract_ServiceDesc, srv)
}

func _MedianContract_LatestTransmissionDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestTransmissionDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedianContractServer).LatestTransmissionDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.MedianContract/LatestTransmissionDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedianContractServer).LatestTransmissionDetails(ctx, req.(*LatestTransmissionDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedianContract_LatestRoundRequested_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestRoundRequestedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedianContractServer).LatestRoundRequested(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.MedianContract/LatestRoundRequested",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedianContractServer).LatestRoundRequested(ctx, req.(*LatestRoundRequestedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MedianContract_ServiceDesc is the grpc.ServiceDesc for MedianContract service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedianContract_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.MedianContract",
	HandlerType: (*MedianContractServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LatestTransmissionDetails",
			Handler:    _MedianContract_LatestTransmissionDetails_Handler,
		},
		{
			MethodName: "LatestRoundRequested",
			Handler:    _MedianContract_LatestRoundRequested_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "median.proto",
}

// OnchainConfigCodecClient is the client API for OnchainConfigCodec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnchainConfigCodecClient interface {
	Encode(ctx context.Context, in *EncodeRequest, opts ...grpc.CallOption) (*EncodeReply, error)
	Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeReply, error)
}

type onchainConfigCodecClient struct {
	cc grpc.ClientConnInterface
}

func NewOnchainConfigCodecClient(cc grpc.ClientConnInterface) OnchainConfigCodecClient {
	return &onchainConfigCodecClient{cc}
}

func (c *onchainConfigCodecClient) Encode(ctx context.Context, in *EncodeRequest, opts ...grpc.CallOption) (*EncodeReply, error) {
	out := new(EncodeReply)
	err := c.cc.Invoke(ctx, "/loop.OnchainConfigCodec/Encode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onchainConfigCodecClient) Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeReply, error) {
	out := new(DecodeReply)
	err := c.cc.Invoke(ctx, "/loop.OnchainConfigCodec/Decode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnchainConfigCodecServer is the server API for OnchainConfigCodec service.
// All implementations must embed UnimplementedOnchainConfigCodecServer
// for forward compatibility
type OnchainConfigCodecServer interface {
	Encode(context.Context, *EncodeRequest) (*EncodeReply, error)
	Decode(context.Context, *DecodeRequest) (*DecodeReply, error)
	mustEmbedUnimplementedOnchainConfigCodecServer()
}

// UnimplementedOnchainConfigCodecServer must be embedded to have forward compatible implementations.
type UnimplementedOnchainConfigCodecServer struct {
}

func (UnimplementedOnchainConfigCodecServer) Encode(context.Context, *EncodeRequest) (*EncodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encode not implemented")
}
func (UnimplementedOnchainConfigCodecServer) Decode(context.Context, *DecodeRequest) (*DecodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decode not implemented")
}
func (UnimplementedOnchainConfigCodecServer) mustEmbedUnimplementedOnchainConfigCodecServer() {}

// UnsafeOnchainConfigCodecServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnchainConfigCodecServer will
// result in compilation errors.
type UnsafeOnchainConfigCodecServer interface {
	mustEmbedUnimplementedOnchainConfigCodecServer()
}

func RegisterOnchainConfigCodecServer(s grpc.ServiceRegistrar, srv OnchainConfigCodecServer) {
	s.RegisterService(&OnchainConfigCodec_ServiceDesc, srv)
}

func _OnchainConfigCodec_Encode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnchainConfigCodecServer).Encode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.OnchainConfigCodec/Encode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnchainConfigCodecServer).Encode(ctx, req.(*EncodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnchainConfigCodec_Decode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnchainConfigCodecServer).Decode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loop.OnchainConfigCodec/Decode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnchainConfigCodecServer).Decode(ctx, req.(*DecodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OnchainConfigCodec_ServiceDesc is the grpc.ServiceDesc for OnchainConfigCodec service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnchainConfigCodec_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.OnchainConfigCodec",
	HandlerType: (*OnchainConfigCodecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encode",
			Handler:    _OnchainConfigCodec_Encode_Handler,
		},
		{
			MethodName: "Decode",
			Handler:    _OnchainConfigCodec_Decode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "median.proto",
}
