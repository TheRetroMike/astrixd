// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: astrixwalletd.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AstrixwalletdClient is the client API for Astrixwalletd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AstrixwalletdClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type astrixwalletdClient struct {
	cc grpc.ClientConnInterface
}

func NewAstrixwalletdClient(cc grpc.ClientConnInterface) AstrixwalletdClient {
	return &astrixwalletdClient{cc}
}

func (c *astrixwalletdClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error) {
	out := new(GetExternalSpendableUTXOsResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/GetExternalSpendableUTXOs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error) {
	out := new(CreateUnsignedTransactionsResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/CreateUnsignedTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error) {
	out := new(ShowAddressesResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/ShowAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error) {
	out := new(NewAddressResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/NewAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *astrixwalletdClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/astrixwalletd.astrixwalletd/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AstrixwalletdServer is the server API for Astrixwalletd service.
// All implementations must embed UnimplementedAstrixwalletdServer
// for forward compatibility
type AstrixwalletdServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedAstrixwalletdServer()
}

// UnimplementedAstrixwalletdServer must be embedded to have forward compatible implementations.
type UnimplementedAstrixwalletdServer struct {
}

func (UnimplementedAstrixwalletdServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedAstrixwalletdServer) GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalSpendableUTXOs not implemented")
}
func (UnimplementedAstrixwalletdServer) CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnsignedTransactions not implemented")
}
func (UnimplementedAstrixwalletdServer) ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAddresses not implemented")
}
func (UnimplementedAstrixwalletdServer) NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddress not implemented")
}
func (UnimplementedAstrixwalletdServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedAstrixwalletdServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedAstrixwalletdServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedAstrixwalletdServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedAstrixwalletdServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedAstrixwalletdServer) mustEmbedUnimplementedAstrixwalletdServer() {}

// UnsafeAstrixwalletdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AstrixwalletdServer will
// result in compilation errors.
type UnsafeAstrixwalletdServer interface {
	mustEmbedUnimplementedAstrixwalletdServer()
}

func RegisterAstrixwalletdServer(s grpc.ServiceRegistrar, srv AstrixwalletdServer) {
	s.RegisterService(&Astrixwalletd_ServiceDesc, srv)
}

func _Astrixwalletd_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_GetExternalSpendableUTXOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExternalSpendableUTXOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).GetExternalSpendableUTXOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/GetExternalSpendableUTXOs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).GetExternalSpendableUTXOs(ctx, req.(*GetExternalSpendableUTXOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_CreateUnsignedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUnsignedTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).CreateUnsignedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/CreateUnsignedTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).CreateUnsignedTransactions(ctx, req.(*CreateUnsignedTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_ShowAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).ShowAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/ShowAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).ShowAddresses(ctx, req.(*ShowAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/NewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).NewAddress(ctx, req.(*NewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Astrixwalletd_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstrixwalletdServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/astrixwalletd.astrixwalletd/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstrixwalletdServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Astrixwalletd_ServiceDesc is the grpc.ServiceDesc for Astrixwalletd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Astrixwalletd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "astrixwalletd.astrixwalletd",
	HandlerType: (*AstrixwalletdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Astrixwalletd_GetBalance_Handler,
		},
		{
			MethodName: "GetExternalSpendableUTXOs",
			Handler:    _Astrixwalletd_GetExternalSpendableUTXOs_Handler,
		},
		{
			MethodName: "CreateUnsignedTransactions",
			Handler:    _Astrixwalletd_CreateUnsignedTransactions_Handler,
		},
		{
			MethodName: "ShowAddresses",
			Handler:    _Astrixwalletd_ShowAddresses_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _Astrixwalletd_NewAddress_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Astrixwalletd_Shutdown_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Astrixwalletd_Broadcast_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Astrixwalletd_Send_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Astrixwalletd_Sign_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Astrixwalletd_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "astrixwalletd.proto",
}
