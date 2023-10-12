// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: initia/intertx/v1/tx.proto

package intertxv1

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

const (
	Msg_RegisterAccount_FullMethodName = "/initia.intertx.v1.Msg/RegisterAccount"
	Msg_SubmitTx_FullMethodName        = "/initia.intertx.v1.Msg/SubmitTx"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// Register defines a rpc handler for MsgRegisterAccount
	RegisterAccount(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error)
	// SubmitTx defines a rpc handler for MsgSubmitTx
	SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterAccount(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error) {
	out := new(MsgRegisterAccountResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error) {
	out := new(MsgSubmitTxResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// Register defines a rpc handler for MsgRegisterAccount
	RegisterAccount(context.Context, *MsgRegisterAccount) (*MsgRegisterAccountResponse, error)
	// SubmitTx defines a rpc handler for MsgSubmitTx
	SubmitTx(context.Context, *MsgSubmitTx) (*MsgSubmitTxResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RegisterAccount(context.Context, *MsgRegisterAccount) (*MsgRegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (UnimplementedMsgServer) SubmitTx(context.Context, *MsgSubmitTx) (*MsgSubmitTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTx not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterAccount(ctx, req.(*MsgRegisterAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitTx(ctx, req.(*MsgSubmitTx))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "initia.intertx.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccount",
			Handler:    _Msg_RegisterAccount_Handler,
		},
		{
			MethodName: "SubmitTx",
			Handler:    _Msg_SubmitTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "initia/intertx/v1/tx.proto",
}
