// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package contentrpc

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

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	AddArticle(ctx context.Context, in *AddArticleRequest, opts ...grpc.CallOption) (*AddArticleResponse, error)
	AddQuote(ctx context.Context, in *AddQuoteRequest, opts ...grpc.CallOption) (*AddQuoteResponse, error)
	AddMeme(ctx context.Context, in *AddMemeRequest, opts ...grpc.CallOption) (*AddMemeResponse, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) AddArticle(ctx context.Context, in *AddArticleRequest, opts ...grpc.CallOption) (*AddArticleResponse, error) {
	out := new(AddArticleResponse)
	err := c.cc.Invoke(ctx, "/contentrpc.Content/AddArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) AddQuote(ctx context.Context, in *AddQuoteRequest, opts ...grpc.CallOption) (*AddQuoteResponse, error) {
	out := new(AddQuoteResponse)
	err := c.cc.Invoke(ctx, "/contentrpc.Content/AddQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) AddMeme(ctx context.Context, in *AddMemeRequest, opts ...grpc.CallOption) (*AddMemeResponse, error) {
	out := new(AddMemeResponse)
	err := c.cc.Invoke(ctx, "/contentrpc.Content/AddMeme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	AddArticle(context.Context, *AddArticleRequest) (*AddArticleResponse, error)
	AddQuote(context.Context, *AddQuoteRequest) (*AddQuoteResponse, error)
	AddMeme(context.Context, *AddMemeRequest) (*AddMemeResponse, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) AddArticle(context.Context, *AddArticleRequest) (*AddArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddArticle not implemented")
}
func (UnimplementedContentServer) AddQuote(context.Context, *AddQuoteRequest) (*AddQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQuote not implemented")
}
func (UnimplementedContentServer) AddMeme(context.Context, *AddMemeRequest) (*AddMemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMeme not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s grpc.ServiceRegistrar, srv ContentServer) {
	s.RegisterService(&Content_ServiceDesc, srv)
}

func _Content_AddArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentrpc.Content/AddArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddArticle(ctx, req.(*AddArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_AddQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentrpc.Content/AddQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddQuote(ctx, req.(*AddQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_AddMeme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddMeme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentrpc.Content/AddMeme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddMeme(ctx, req.(*AddMemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Content_ServiceDesc is the grpc.ServiceDesc for Content service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Content_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contentrpc.Content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddArticle",
			Handler:    _Content_AddArticle_Handler,
		},
		{
			MethodName: "AddQuote",
			Handler:    _Content_AddQuote_Handler,
		},
		{
			MethodName: "AddMeme",
			Handler:    _Content_AddMeme_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contentrpc/content.proto",
}
