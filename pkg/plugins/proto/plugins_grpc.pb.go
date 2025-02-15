// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: plugins.proto

package proto

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

// SourceClient is the client API for Source service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceClient interface {
	Open(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error)
	Teardown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Validate(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error)
	Read(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Record, error)
	Ack(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Empty, error)
}

type sourceClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceClient(cc grpc.ClientConnInterface) SourceClient {
	return &sourceClient{cc}
}

func (c *sourceClient) Open(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Source/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Teardown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Source/Teardown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Validate(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Source/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Read(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Source/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Ack(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Source/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceServer is the server API for Source service.
// All implementations must embed UnimplementedSourceServer
// for forward compatibility
type SourceServer interface {
	Open(context.Context, *Config) (*Empty, error)
	Teardown(context.Context, *Empty) (*Empty, error)
	Validate(context.Context, *Config) (*Empty, error)
	Read(context.Context, *Position) (*Record, error)
	Ack(context.Context, *Position) (*Empty, error)
	mustEmbedUnimplementedSourceServer()
}

// UnimplementedSourceServer must be embedded to have forward compatible implementations.
type UnimplementedSourceServer struct {
}

func (UnimplementedSourceServer) Open(context.Context, *Config) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedSourceServer) Teardown(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teardown not implemented")
}
func (UnimplementedSourceServer) Validate(context.Context, *Config) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedSourceServer) Read(context.Context, *Position) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedSourceServer) Ack(context.Context, *Position) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (UnimplementedSourceServer) mustEmbedUnimplementedSourceServer() {}

// UnsafeSourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceServer will
// result in compilation errors.
type UnsafeSourceServer interface {
	mustEmbedUnimplementedSourceServer()
}

func RegisterSourceServer(s grpc.ServiceRegistrar, srv SourceServer) {
	s.RegisterService(&Source_ServiceDesc, srv)
}

func _Source_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Source/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Open(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Teardown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Teardown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Source/Teardown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Teardown(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Source/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Validate(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Source/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Read(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Source/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Ack(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

// Source_ServiceDesc is the grpc.ServiceDesc for Source service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Source_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "conduit.plugins.Source",
	HandlerType: (*SourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _Source_Open_Handler,
		},
		{
			MethodName: "Teardown",
			Handler:    _Source_Teardown_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Source_Validate_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Source_Read_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _Source_Ack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins.proto",
}

// DestinationClient is the client API for Destination service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DestinationClient interface {
	Open(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error)
	Teardown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Validate(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error)
	Write(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Position, error)
}

type destinationClient struct {
	cc grpc.ClientConnInterface
}

func NewDestinationClient(cc grpc.ClientConnInterface) DestinationClient {
	return &destinationClient{cc}
}

func (c *destinationClient) Open(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Destination/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) Teardown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Destination/Teardown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) Validate(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Destination/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) Write(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Destination/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DestinationServer is the server API for Destination service.
// All implementations must embed UnimplementedDestinationServer
// for forward compatibility
type DestinationServer interface {
	Open(context.Context, *Config) (*Empty, error)
	Teardown(context.Context, *Empty) (*Empty, error)
	Validate(context.Context, *Config) (*Empty, error)
	Write(context.Context, *Record) (*Position, error)
	mustEmbedUnimplementedDestinationServer()
}

// UnimplementedDestinationServer must be embedded to have forward compatible implementations.
type UnimplementedDestinationServer struct {
}

func (UnimplementedDestinationServer) Open(context.Context, *Config) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedDestinationServer) Teardown(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teardown not implemented")
}
func (UnimplementedDestinationServer) Validate(context.Context, *Config) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedDestinationServer) Write(context.Context, *Record) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedDestinationServer) mustEmbedUnimplementedDestinationServer() {}

// UnsafeDestinationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DestinationServer will
// result in compilation errors.
type UnsafeDestinationServer interface {
	mustEmbedUnimplementedDestinationServer()
}

func RegisterDestinationServer(s grpc.ServiceRegistrar, srv DestinationServer) {
	s.RegisterService(&Destination_ServiceDesc, srv)
}

func _Destination_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Destination/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).Open(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_Teardown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).Teardown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Destination/Teardown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).Teardown(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Destination/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).Validate(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Destination/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).Write(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

// Destination_ServiceDesc is the grpc.ServiceDesc for Destination service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Destination_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "conduit.plugins.Destination",
	HandlerType: (*DestinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _Destination_Open_Handler,
		},
		{
			MethodName: "Teardown",
			Handler:    _Destination_Teardown_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Destination_Validate_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Destination_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins.proto",
}

// SpecificationsClient is the client API for Specifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpecificationsClient interface {
	Specify(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Specification, error)
}

type specificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewSpecificationsClient(cc grpc.ClientConnInterface) SpecificationsClient {
	return &specificationsClient{cc}
}

func (c *specificationsClient) Specify(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Specification, error) {
	out := new(Specification)
	err := c.cc.Invoke(ctx, "/conduit.plugins.Specifications/Specify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpecificationsServer is the server API for Specifications service.
// All implementations must embed UnimplementedSpecificationsServer
// for forward compatibility
type SpecificationsServer interface {
	Specify(context.Context, *Empty) (*Specification, error)
	mustEmbedUnimplementedSpecificationsServer()
}

// UnimplementedSpecificationsServer must be embedded to have forward compatible implementations.
type UnimplementedSpecificationsServer struct {
}

func (UnimplementedSpecificationsServer) Specify(context.Context, *Empty) (*Specification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Specify not implemented")
}
func (UnimplementedSpecificationsServer) mustEmbedUnimplementedSpecificationsServer() {}

// UnsafeSpecificationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpecificationsServer will
// result in compilation errors.
type UnsafeSpecificationsServer interface {
	mustEmbedUnimplementedSpecificationsServer()
}

func RegisterSpecificationsServer(s grpc.ServiceRegistrar, srv SpecificationsServer) {
	s.RegisterService(&Specifications_ServiceDesc, srv)
}

func _Specifications_Specify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecificationsServer).Specify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.plugins.Specifications/Specify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecificationsServer).Specify(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Specifications_ServiceDesc is the grpc.ServiceDesc for Specifications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Specifications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "conduit.plugins.Specifications",
	HandlerType: (*SpecificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Specify",
			Handler:    _Specifications_Specify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins.proto",
}
