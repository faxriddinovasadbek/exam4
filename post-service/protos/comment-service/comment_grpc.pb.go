// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: protos/comment-service/comment.proto

package __

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
	CommentService_GetAllUsers_FullMethodName          = "/comment.CommentService/GetAllUsers"
	CommentService_GetPostById_FullMethodName          = "/comment.CommentService/GetPostById"
	CommentService_GetUserById_FullMethodName          = "/comment.CommentService/GetUserById"
	CommentService_CreateComment_FullMethodName        = "/comment.CommentService/CreateComment"
	CommentService_UpdateComment_FullMethodName        = "/comment.CommentService/UpdateComment"
	CommentService_DeleteComment_FullMethodName        = "/comment.CommentService/DeleteComment"
	CommentService_GetComment_FullMethodName           = "/comment.CommentService/GetComment"
	CommentService_GetAllComment_FullMethodName        = "/comment.CommentService/GetAllComment"
	CommentService_GetCommentsByPostId_FullMethodName  = "/comment.CommentService/GetCommentsByPostId"
	CommentService_GetCommentsByOwnerId_FullMethodName = "/comment.CommentService/GetCommentsByOwnerId"
	CommentService_GetCommentById_FullMethodName       = "/comment.CommentService/GetCommentById"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	GetAllUsers(ctx context.Context, in *GetAllCommentsRequest, opts ...grpc.CallOption) (*GetAllCommentsResponse, error)
	GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error)
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error)
	UpdateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error)
	DeleteComment(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetComment(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*Comment, error)
	GetAllComment(ctx context.Context, in *GetAllCommentsRequest, opts ...grpc.CallOption) (*GetAllCommentResponse, error)
	GetCommentsByPostId(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*GetAllCommentResponse, error)
	GetCommentsByOwnerId(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*GetAllCommentResponse, error)
	GetCommentById(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*Comment, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) GetAllUsers(ctx context.Context, in *GetAllCommentsRequest, opts ...grpc.CallOption) (*GetAllCommentsResponse, error) {
	out := new(GetAllCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetAllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error) {
	out := new(GetPostByIdResponse)
	err := c.cc.Invoke(ctx, CommentService_GetPostById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, CommentService_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, CommentService_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) UpdateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, CommentService_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) DeleteComment(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, CommentService_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetComment(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, CommentService_GetComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetAllComment(ctx context.Context, in *GetAllCommentsRequest, opts ...grpc.CallOption) (*GetAllCommentResponse, error) {
	out := new(GetAllCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_GetAllComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentsByPostId(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*GetAllCommentResponse, error) {
	out := new(GetAllCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_GetCommentsByPostId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentsByOwnerId(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*GetAllCommentResponse, error) {
	out := new(GetAllCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_GetCommentsByOwnerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentById(ctx context.Context, in *IdRequst, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, CommentService_GetCommentById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	GetAllUsers(context.Context, *GetAllCommentsRequest) (*GetAllCommentsResponse, error)
	GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	CreateComment(context.Context, *Comment) (*Comment, error)
	UpdateComment(context.Context, *Comment) (*Comment, error)
	DeleteComment(context.Context, *IdRequst) (*DeleteResponse, error)
	GetComment(context.Context, *IdRequst) (*Comment, error)
	GetAllComment(context.Context, *GetAllCommentsRequest) (*GetAllCommentResponse, error)
	GetCommentsByPostId(context.Context, *IdRequst) (*GetAllCommentResponse, error)
	GetCommentsByOwnerId(context.Context, *IdRequst) (*GetAllCommentResponse, error)
	GetCommentById(context.Context, *IdRequst) (*Comment, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) GetAllUsers(context.Context, *GetAllCommentsRequest) (*GetAllCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedCommentServiceServer) GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedCommentServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedCommentServiceServer) CreateComment(context.Context, *Comment) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) UpdateComment(context.Context, *Comment) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentServiceServer) DeleteComment(context.Context, *IdRequst) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentServiceServer) GetComment(context.Context, *IdRequst) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentServiceServer) GetAllComment(context.Context, *GetAllCommentsRequest) (*GetAllCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllComment not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentsByPostId(context.Context, *IdRequst) (*GetAllCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByPostId not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentsByOwnerId(context.Context, *IdRequst) (*GetAllCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByOwnerId not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentById(context.Context, *IdRequst) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetAllUsers(ctx, req.(*GetAllCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetPostById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetPostById(ctx, req.(*GetPostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).UpdateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).DeleteComment(ctx, req.(*IdRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetComment(ctx, req.(*IdRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetAllComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetAllComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetAllComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetAllComment(ctx, req.(*GetAllCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentsByPostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentsByPostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetCommentsByPostId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentsByPostId(ctx, req.(*IdRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentsByOwnerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentsByOwnerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetCommentsByOwnerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentsByOwnerId(ctx, req.(*IdRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetCommentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentById(ctx, req.(*IdRequst))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUsers",
			Handler:    _CommentService_GetAllUsers_Handler,
		},
		{
			MethodName: "GetPostById",
			Handler:    _CommentService_GetPostById_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _CommentService_GetUserById_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _CommentService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentService_DeleteComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _CommentService_GetComment_Handler,
		},
		{
			MethodName: "GetAllComment",
			Handler:    _CommentService_GetAllComment_Handler,
		},
		{
			MethodName: "GetCommentsByPostId",
			Handler:    _CommentService_GetCommentsByPostId_Handler,
		},
		{
			MethodName: "GetCommentsByOwnerId",
			Handler:    _CommentService_GetCommentsByOwnerId_Handler,
		},
		{
			MethodName: "GetCommentById",
			Handler:    _CommentService_GetCommentById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/comment-service/comment.proto",
}
