// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: initia/bank/v1/tx.proto

package bankv1

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
	Msg_SetDenomMetadata_FullMethodName = "/initia.bank.v1.Msg/SetDenomMetadata"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// SetDenomMetadata defines a governance operation for updating the x/bank
	// denom metadata. The authority is defined in the keeper.
	SetDenomMetadata(ctx context.Context, in *MsgSetDenomMetadata, opts ...grpc.CallOption) (*MsgSetDenomMetadataResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetDenomMetadata(ctx context.Context, in *MsgSetDenomMetadata, opts ...grpc.CallOption) (*MsgSetDenomMetadataResponse, error) {
	out := new(MsgSetDenomMetadataResponse)
	err := c.cc.Invoke(ctx, Msg_SetDenomMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// SetDenomMetadata defines a governance operation for updating the x/bank
	// denom metadata. The authority is defined in the keeper.
	SetDenomMetadata(context.Context, *MsgSetDenomMetadata) (*MsgSetDenomMetadataResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) SetDenomMetadata(context.Context, *MsgSetDenomMetadata) (*MsgSetDenomMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDenomMetadata not implemented")
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

func _Msg_SetDenomMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetDenomMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetDenomMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetDenomMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetDenomMetadata(ctx, req.(*MsgSetDenomMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "initia.bank.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDenomMetadata",
			Handler:    _Msg_SetDenomMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "initia/bank/v1/tx.proto",
}
