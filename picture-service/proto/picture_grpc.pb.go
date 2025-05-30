// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/picture.proto

package picture

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PictureService_CreatePicture_FullMethodName = "/picture.PictureService/CreatePicture"
	PictureService_GetAll_FullMethodName        = "/picture.PictureService/GetAll"
	PictureService_GetPicture_FullMethodName    = "/picture.PictureService/GetPicture"
	PictureService_UpdatePicture_FullMethodName = "/picture.PictureService/UpdatePicture"
	PictureService_DeletePicture_FullMethodName = "/picture.PictureService/DeletePicture"
)

// PictureServiceClient is the client API for PictureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PictureServiceClient interface {
	CreatePicture(ctx context.Context, in *CreatePictureRequest, opts ...grpc.CallOption) (*CreatePictureResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetPicture(ctx context.Context, in *GetPictureRequest, opts ...grpc.CallOption) (*GetPictureResponse, error)
	UpdatePicture(ctx context.Context, in *UpdatePictureRequest, opts ...grpc.CallOption) (*UpdatePictureResponse, error)
	DeletePicture(ctx context.Context, in *DeletePictureRequest, opts ...grpc.CallOption) (*DeletePictureResponse, error)
}

type pictureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPictureServiceClient(cc grpc.ClientConnInterface) PictureServiceClient {
	return &pictureServiceClient{cc}
}

func (c *pictureServiceClient) CreatePicture(ctx context.Context, in *CreatePictureRequest, opts ...grpc.CallOption) (*CreatePictureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePictureResponse)
	err := c.cc.Invoke(ctx, PictureService_CreatePicture_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pictureServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, PictureService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pictureServiceClient) GetPicture(ctx context.Context, in *GetPictureRequest, opts ...grpc.CallOption) (*GetPictureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPictureResponse)
	err := c.cc.Invoke(ctx, PictureService_GetPicture_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pictureServiceClient) UpdatePicture(ctx context.Context, in *UpdatePictureRequest, opts ...grpc.CallOption) (*UpdatePictureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePictureResponse)
	err := c.cc.Invoke(ctx, PictureService_UpdatePicture_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pictureServiceClient) DeletePicture(ctx context.Context, in *DeletePictureRequest, opts ...grpc.CallOption) (*DeletePictureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePictureResponse)
	err := c.cc.Invoke(ctx, PictureService_DeletePicture_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PictureServiceServer is the server API for PictureService service.
// All implementations must embed UnimplementedPictureServiceServer
// for forward compatibility.
type PictureServiceServer interface {
	CreatePicture(context.Context, *CreatePictureRequest) (*CreatePictureResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetPicture(context.Context, *GetPictureRequest) (*GetPictureResponse, error)
	UpdatePicture(context.Context, *UpdatePictureRequest) (*UpdatePictureResponse, error)
	DeletePicture(context.Context, *DeletePictureRequest) (*DeletePictureResponse, error)
	mustEmbedUnimplementedPictureServiceServer()
}

// UnimplementedPictureServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPictureServiceServer struct{}

func (UnimplementedPictureServiceServer) CreatePicture(context.Context, *CreatePictureRequest) (*CreatePictureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePicture not implemented")
}
func (UnimplementedPictureServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPictureServiceServer) GetPicture(context.Context, *GetPictureRequest) (*GetPictureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPicture not implemented")
}
func (UnimplementedPictureServiceServer) UpdatePicture(context.Context, *UpdatePictureRequest) (*UpdatePictureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePicture not implemented")
}
func (UnimplementedPictureServiceServer) DeletePicture(context.Context, *DeletePictureRequest) (*DeletePictureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePicture not implemented")
}
func (UnimplementedPictureServiceServer) mustEmbedUnimplementedPictureServiceServer() {}
func (UnimplementedPictureServiceServer) testEmbeddedByValue()                        {}

// UnsafePictureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PictureServiceServer will
// result in compilation errors.
type UnsafePictureServiceServer interface {
	mustEmbedUnimplementedPictureServiceServer()
}

func RegisterPictureServiceServer(s grpc.ServiceRegistrar, srv PictureServiceServer) {
	// If the following call pancis, it indicates UnimplementedPictureServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PictureService_ServiceDesc, srv)
}

func _PictureService_CreatePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PictureServiceServer).CreatePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PictureService_CreatePicture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PictureServiceServer).CreatePicture(ctx, req.(*CreatePictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PictureService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PictureServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PictureService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PictureServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PictureService_GetPicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PictureServiceServer).GetPicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PictureService_GetPicture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PictureServiceServer).GetPicture(ctx, req.(*GetPictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PictureService_UpdatePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PictureServiceServer).UpdatePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PictureService_UpdatePicture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PictureServiceServer).UpdatePicture(ctx, req.(*UpdatePictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PictureService_DeletePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PictureServiceServer).DeletePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PictureService_DeletePicture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PictureServiceServer).DeletePicture(ctx, req.(*DeletePictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PictureService_ServiceDesc is the grpc.ServiceDesc for PictureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PictureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "picture.PictureService",
	HandlerType: (*PictureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePicture",
			Handler:    _PictureService_CreatePicture_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PictureService_GetAll_Handler,
		},
		{
			MethodName: "GetPicture",
			Handler:    _PictureService_GetPicture_Handler,
		},
		{
			MethodName: "UpdatePicture",
			Handler:    _PictureService_UpdatePicture_Handler,
		},
		{
			MethodName: "DeletePicture",
			Handler:    _PictureService_DeletePicture_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/picture.proto",
}
