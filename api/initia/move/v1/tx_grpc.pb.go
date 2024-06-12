// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: initia/move/v1/tx.proto

package movev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Msg_Publish_FullMethodName        = "/initia.move.v1.Msg/Publish"
	Msg_Execute_FullMethodName        = "/initia.move.v1.Msg/Execute"
	Msg_ExecuteJSON_FullMethodName    = "/initia.move.v1.Msg/ExecuteJSON"
	Msg_Script_FullMethodName         = "/initia.move.v1.Msg/Script"
	Msg_ScriptJSON_FullMethodName     = "/initia.move.v1.Msg/ScriptJSON"
	Msg_GovPublish_FullMethodName     = "/initia.move.v1.Msg/GovPublish"
	Msg_GovExecute_FullMethodName     = "/initia.move.v1.Msg/GovExecute"
	Msg_GovExecuteJSON_FullMethodName = "/initia.move.v1.Msg/GovExecuteJSON"
	Msg_GovScript_FullMethodName      = "/initia.move.v1.Msg/GovScript"
	Msg_GovScriptJSON_FullMethodName  = "/initia.move.v1.Msg/GovScriptJSON"
	Msg_Whitelist_FullMethodName      = "/initia.move.v1.Msg/Whitelist"
	Msg_Delist_FullMethodName         = "/initia.move.v1.Msg/Delist"
	Msg_UpdateParams_FullMethodName   = "/initia.move.v1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Msg defines the move Msg service.
type MsgClient interface {
	// Publish stores compiled Move module
	Publish(ctx context.Context, in *MsgPublish, opts ...grpc.CallOption) (*MsgPublishResponse, error)
	// Deprecated: Use ExecuteJSON instead
	// Execute runs a entry function with the given message
	Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error)
	// ExecuteJSON runs a entry function with the given message
	ExecuteJSON(ctx context.Context, in *MsgExecuteJSON, opts ...grpc.CallOption) (*MsgExecuteJSONResponse, error)
	// Deprecated: Use ScriptJSON instead
	// Script runs a scripts with the given message
	Script(ctx context.Context, in *MsgScript, opts ...grpc.CallOption) (*MsgScriptResponse, error)
	// ScriptJSON runs a scripts with the given message
	ScriptJSON(ctx context.Context, in *MsgScriptJSON, opts ...grpc.CallOption) (*MsgScriptJSONResponse, error)
	// GovPublish stores compiled Move module via gov proposal
	GovPublish(ctx context.Context, in *MsgGovPublish, opts ...grpc.CallOption) (*MsgGovPublishResponse, error)
	// Deprecated: Use GovExecuteJSON instead
	// GovExecute runs a entry function with the given message via gov proposal
	GovExecute(ctx context.Context, in *MsgGovExecute, opts ...grpc.CallOption) (*MsgGovExecuteResponse, error)
	// GovExecuteJSON runs a entry function with the given message via gov proposal
	GovExecuteJSON(ctx context.Context, in *MsgGovExecuteJSON, opts ...grpc.CallOption) (*MsgGovExecuteJSONResponse, error)
	// Deprecated: Use GovScriptJSON instead
	// GovScript runs a scripts with the given message via gov proposal
	GovScript(ctx context.Context, in *MsgGovScript, opts ...grpc.CallOption) (*MsgGovScriptResponse, error)
	// GovScriptJSON runs a scripts with the given message via gov proposal
	GovScriptJSON(ctx context.Context, in *MsgGovScriptJSON, opts ...grpc.CallOption) (*MsgGovScriptJSONResponse, error)
	// Whitelist registers a dex pair to whitelist of various features.
	// - whitelist from coin register operation
	// - allow counter party denom can be used as gas fee
	// - register lp denom as staking denom
	Whitelist(ctx context.Context, in *MsgWhitelist, opts ...grpc.CallOption) (*MsgWhitelistResponse, error)
	// Delist unregisters a dex pair from the whitelist.
	Delist(ctx context.Context, in *MsgDelist, opts ...grpc.CallOption) (*MsgDelistResponse, error)
	// UpdateParams defines an operation for updating the x/move module
	// parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Publish(ctx context.Context, in *MsgPublish, opts ...grpc.CallOption) (*MsgPublishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgPublishResponse)
	err := c.cc.Invoke(ctx, Msg_Publish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgExecuteResponse)
	err := c.cc.Invoke(ctx, Msg_Execute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExecuteJSON(ctx context.Context, in *MsgExecuteJSON, opts ...grpc.CallOption) (*MsgExecuteJSONResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgExecuteJSONResponse)
	err := c.cc.Invoke(ctx, Msg_ExecuteJSON_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Script(ctx context.Context, in *MsgScript, opts ...grpc.CallOption) (*MsgScriptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgScriptResponse)
	err := c.cc.Invoke(ctx, Msg_Script_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ScriptJSON(ctx context.Context, in *MsgScriptJSON, opts ...grpc.CallOption) (*MsgScriptJSONResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgScriptJSONResponse)
	err := c.cc.Invoke(ctx, Msg_ScriptJSON_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovPublish(ctx context.Context, in *MsgGovPublish, opts ...grpc.CallOption) (*MsgGovPublishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgGovPublishResponse)
	err := c.cc.Invoke(ctx, Msg_GovPublish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovExecute(ctx context.Context, in *MsgGovExecute, opts ...grpc.CallOption) (*MsgGovExecuteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgGovExecuteResponse)
	err := c.cc.Invoke(ctx, Msg_GovExecute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovExecuteJSON(ctx context.Context, in *MsgGovExecuteJSON, opts ...grpc.CallOption) (*MsgGovExecuteJSONResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgGovExecuteJSONResponse)
	err := c.cc.Invoke(ctx, Msg_GovExecuteJSON_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovScript(ctx context.Context, in *MsgGovScript, opts ...grpc.CallOption) (*MsgGovScriptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgGovScriptResponse)
	err := c.cc.Invoke(ctx, Msg_GovScript_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovScriptJSON(ctx context.Context, in *MsgGovScriptJSON, opts ...grpc.CallOption) (*MsgGovScriptJSONResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgGovScriptJSONResponse)
	err := c.cc.Invoke(ctx, Msg_GovScriptJSON_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Whitelist(ctx context.Context, in *MsgWhitelist, opts ...grpc.CallOption) (*MsgWhitelistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgWhitelistResponse)
	err := c.cc.Invoke(ctx, Msg_Whitelist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Delist(ctx context.Context, in *MsgDelist, opts ...grpc.CallOption) (*MsgDelistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgDelistResponse)
	err := c.cc.Invoke(ctx, Msg_Delist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
//
// Msg defines the move Msg service.
type MsgServer interface {
	// Publish stores compiled Move module
	Publish(context.Context, *MsgPublish) (*MsgPublishResponse, error)
	// Deprecated: Use ExecuteJSON instead
	// Execute runs a entry function with the given message
	Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error)
	// ExecuteJSON runs a entry function with the given message
	ExecuteJSON(context.Context, *MsgExecuteJSON) (*MsgExecuteJSONResponse, error)
	// Deprecated: Use ScriptJSON instead
	// Script runs a scripts with the given message
	Script(context.Context, *MsgScript) (*MsgScriptResponse, error)
	// ScriptJSON runs a scripts with the given message
	ScriptJSON(context.Context, *MsgScriptJSON) (*MsgScriptJSONResponse, error)
	// GovPublish stores compiled Move module via gov proposal
	GovPublish(context.Context, *MsgGovPublish) (*MsgGovPublishResponse, error)
	// Deprecated: Use GovExecuteJSON instead
	// GovExecute runs a entry function with the given message via gov proposal
	GovExecute(context.Context, *MsgGovExecute) (*MsgGovExecuteResponse, error)
	// GovExecuteJSON runs a entry function with the given message via gov proposal
	GovExecuteJSON(context.Context, *MsgGovExecuteJSON) (*MsgGovExecuteJSONResponse, error)
	// Deprecated: Use GovScriptJSON instead
	// GovScript runs a scripts with the given message via gov proposal
	GovScript(context.Context, *MsgGovScript) (*MsgGovScriptResponse, error)
	// GovScriptJSON runs a scripts with the given message via gov proposal
	GovScriptJSON(context.Context, *MsgGovScriptJSON) (*MsgGovScriptJSONResponse, error)
	// Whitelist registers a dex pair to whitelist of various features.
	// - whitelist from coin register operation
	// - allow counter party denom can be used as gas fee
	// - register lp denom as staking denom
	Whitelist(context.Context, *MsgWhitelist) (*MsgWhitelistResponse, error)
	// Delist unregisters a dex pair from the whitelist.
	Delist(context.Context, *MsgDelist) (*MsgDelistResponse, error)
	// UpdateParams defines an operation for updating the x/move module
	// parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Publish(context.Context, *MsgPublish) (*MsgPublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedMsgServer) Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedMsgServer) ExecuteJSON(context.Context, *MsgExecuteJSON) (*MsgExecuteJSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteJSON not implemented")
}
func (UnimplementedMsgServer) Script(context.Context, *MsgScript) (*MsgScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Script not implemented")
}
func (UnimplementedMsgServer) ScriptJSON(context.Context, *MsgScriptJSON) (*MsgScriptJSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScriptJSON not implemented")
}
func (UnimplementedMsgServer) GovPublish(context.Context, *MsgGovPublish) (*MsgGovPublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovPublish not implemented")
}
func (UnimplementedMsgServer) GovExecute(context.Context, *MsgGovExecute) (*MsgGovExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovExecute not implemented")
}
func (UnimplementedMsgServer) GovExecuteJSON(context.Context, *MsgGovExecuteJSON) (*MsgGovExecuteJSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovExecuteJSON not implemented")
}
func (UnimplementedMsgServer) GovScript(context.Context, *MsgGovScript) (*MsgGovScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovScript not implemented")
}
func (UnimplementedMsgServer) GovScriptJSON(context.Context, *MsgGovScriptJSON) (*MsgGovScriptJSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovScriptJSON not implemented")
}
func (UnimplementedMsgServer) Whitelist(context.Context, *MsgWhitelist) (*MsgWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Whitelist not implemented")
}
func (UnimplementedMsgServer) Delist(context.Context, *MsgDelist) (*MsgDelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delist not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
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

func _Msg_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPublish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Publish(ctx, req.(*MsgPublish))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Execute(ctx, req.(*MsgExecute))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExecuteJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteJSON)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ExecuteJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteJSON(ctx, req.(*MsgExecuteJSON))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Script_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgScript)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Script(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Script_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Script(ctx, req.(*MsgScript))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ScriptJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgScriptJSON)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ScriptJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ScriptJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ScriptJSON(ctx, req.(*MsgScriptJSON))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovPublish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GovPublish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovPublish(ctx, req.(*MsgGovPublish))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GovExecute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovExecute(ctx, req.(*MsgGovExecute))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovExecuteJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovExecuteJSON)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovExecuteJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GovExecuteJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovExecuteJSON(ctx, req.(*MsgGovExecuteJSON))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovScript)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GovScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovScript(ctx, req.(*MsgGovScript))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovScriptJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovScriptJSON)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovScriptJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GovScriptJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovScriptJSON(ctx, req.(*MsgGovScriptJSON))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Whitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Whitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Whitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Whitelist(ctx, req.(*MsgWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Delist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Delist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Delist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Delist(ctx, req.(*MsgDelist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "initia.move.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Msg_Publish_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _Msg_Execute_Handler,
		},
		{
			MethodName: "ExecuteJSON",
			Handler:    _Msg_ExecuteJSON_Handler,
		},
		{
			MethodName: "Script",
			Handler:    _Msg_Script_Handler,
		},
		{
			MethodName: "ScriptJSON",
			Handler:    _Msg_ScriptJSON_Handler,
		},
		{
			MethodName: "GovPublish",
			Handler:    _Msg_GovPublish_Handler,
		},
		{
			MethodName: "GovExecute",
			Handler:    _Msg_GovExecute_Handler,
		},
		{
			MethodName: "GovExecuteJSON",
			Handler:    _Msg_GovExecuteJSON_Handler,
		},
		{
			MethodName: "GovScript",
			Handler:    _Msg_GovScript_Handler,
		},
		{
			MethodName: "GovScriptJSON",
			Handler:    _Msg_GovScriptJSON_Handler,
		},
		{
			MethodName: "Whitelist",
			Handler:    _Msg_Whitelist_Handler,
		},
		{
			MethodName: "Delist",
			Handler:    _Msg_Delist_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "initia/move/v1/tx.proto",
}
