// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	api "github.com/tweag/trustix/packages/trustix-proto/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TrustixRPCClient is the client API for TrustixRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrustixRPCClient interface {
	// Get map[LogID]Log (all local logs)
	Logs(ctx context.Context, in *api.LogsRequest, opts ...grpc.CallOption) (*api.LogsResponse, error)
	GetLogEntries(ctx context.Context, in *api.GetLogEntriesRequest, opts ...grpc.CallOption) (*api.LogEntriesResponse, error)
	// Get map[LogID]OutputHash
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*EntriesResponse, error)
	// Compare(inputHash)
	Decide(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*DecisionResponse, error)
	// Get stored value by digest (TODO: Remove, it's a duplicate from api.proto
	GetValue(ctx context.Context, in *api.ValueRequest, opts ...grpc.CallOption) (*api.ValueResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
}

type trustixRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTrustixRPCClient(cc grpc.ClientConnInterface) TrustixRPCClient {
	return &trustixRPCClient{cc}
}

func (c *trustixRPCClient) Logs(ctx context.Context, in *api.LogsRequest, opts ...grpc.CallOption) (*api.LogsResponse, error) {
	out := new(api.LogsResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixRPC/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixRPCClient) GetLogEntries(ctx context.Context, in *api.GetLogEntriesRequest, opts ...grpc.CallOption) (*api.LogEntriesResponse, error) {
	out := new(api.LogEntriesResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixRPC/GetLogEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixRPCClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*EntriesResponse, error) {
	out := new(EntriesResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixRPCClient) Decide(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*DecisionResponse, error) {
	out := new(DecisionResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixRPC/Decide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixRPCClient) GetValue(ctx context.Context, in *api.ValueRequest, opts ...grpc.CallOption) (*api.ValueResponse, error) {
	out := new(api.ValueResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixRPC/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixRPCClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixRPC/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixRPCClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixRPC/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustixRPCServer is the server API for TrustixRPC service.
// All implementations must embed UnimplementedTrustixRPCServer
// for forward compatibility
type TrustixRPCServer interface {
	// Get map[LogID]Log (all local logs)
	Logs(context.Context, *api.LogsRequest) (*api.LogsResponse, error)
	GetLogEntries(context.Context, *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error)
	// Get map[LogID]OutputHash
	Get(context.Context, *KeyRequest) (*EntriesResponse, error)
	// Compare(inputHash)
	Decide(context.Context, *KeyRequest) (*DecisionResponse, error)
	// Get stored value by digest (TODO: Remove, it's a duplicate from api.proto
	GetValue(context.Context, *api.ValueRequest) (*api.ValueResponse, error)
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	mustEmbedUnimplementedTrustixRPCServer()
}

// UnimplementedTrustixRPCServer must be embedded to have forward compatible implementations.
type UnimplementedTrustixRPCServer struct {
}

func (UnimplementedTrustixRPCServer) Logs(context.Context, *api.LogsRequest) (*api.LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedTrustixRPCServer) GetLogEntries(context.Context, *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogEntries not implemented")
}
func (UnimplementedTrustixRPCServer) Get(context.Context, *KeyRequest) (*EntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTrustixRPCServer) Decide(context.Context, *KeyRequest) (*DecisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decide not implemented")
}
func (UnimplementedTrustixRPCServer) GetValue(context.Context, *api.ValueRequest) (*api.ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedTrustixRPCServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedTrustixRPCServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedTrustixRPCServer) mustEmbedUnimplementedTrustixRPCServer() {}

// UnsafeTrustixRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrustixRPCServer will
// result in compilation errors.
type UnsafeTrustixRPCServer interface {
	mustEmbedUnimplementedTrustixRPCServer()
}

func RegisterTrustixRPCServer(s grpc.ServiceRegistrar, srv TrustixRPCServer) {
	s.RegisterService(&TrustixRPC_ServiceDesc, srv)
}

func _TrustixRPC_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.LogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixRPCServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixRPC/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixRPCServer).Logs(ctx, req.(*api.LogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixRPC_GetLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixRPCServer).GetLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixRPC/GetLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixRPCServer).GetLogEntries(ctx, req.(*api.GetLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixRPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixRPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixRPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixRPCServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixRPC_Decide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixRPCServer).Decide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixRPC/Decide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixRPCServer).Decide(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixRPC_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixRPCServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixRPC/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixRPCServer).GetValue(ctx, req.(*api.ValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixRPC_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixRPCServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixRPC/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixRPCServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixRPC_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixRPCServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixRPC/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixRPCServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrustixRPC_ServiceDesc is the grpc.ServiceDesc for TrustixRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrustixRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trustix.TrustixRPC",
	HandlerType: (*TrustixRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logs",
			Handler:    _TrustixRPC_Logs_Handler,
		},
		{
			MethodName: "GetLogEntries",
			Handler:    _TrustixRPC_GetLogEntries_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TrustixRPC_Get_Handler,
		},
		{
			MethodName: "Decide",
			Handler:    _TrustixRPC_Decide_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _TrustixRPC_GetValue_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _TrustixRPC_Submit_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _TrustixRPC_Flush_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
