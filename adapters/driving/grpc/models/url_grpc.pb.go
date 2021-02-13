// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package models

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

// UrlShortenerServiceClient is the client API for UrlShortenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortenerServiceClient interface {
	ShortenUrl(ctx context.Context, in *ShortenUrlRequest, opts ...grpc.CallOption) (*ShortenedUrl, error)
	LookupUrl(ctx context.Context, in *LookupUrlRequest, opts ...grpc.CallOption) (*ShortenedUrl, error)
}

type urlShortenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortenerServiceClient(cc grpc.ClientConnInterface) UrlShortenerServiceClient {
	return &urlShortenerServiceClient{cc}
}

func (c *urlShortenerServiceClient) ShortenUrl(ctx context.Context, in *ShortenUrlRequest, opts ...grpc.CallOption) (*ShortenedUrl, error) {
	out := new(ShortenedUrl)
	err := c.cc.Invoke(ctx, "/proto.UrlShortenerService/ShortenUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerServiceClient) LookupUrl(ctx context.Context, in *LookupUrlRequest, opts ...grpc.CallOption) (*ShortenedUrl, error) {
	out := new(ShortenedUrl)
	err := c.cc.Invoke(ctx, "/proto.UrlShortenerService/LookupUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortenerServiceServer is the server API for UrlShortenerService service.
// All implementations must embed UnimplementedUrlShortenerServiceServer
// for forward compatibility
type UrlShortenerServiceServer interface {
	ShortenUrl(context.Context, *ShortenUrlRequest) (*ShortenedUrl, error)
	LookupUrl(context.Context, *LookupUrlRequest) (*ShortenedUrl, error)
	mustEmbedUnimplementedUrlShortenerServiceServer()
}

// UnimplementedUrlShortenerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUrlShortenerServiceServer struct {
}

func (UnimplementedUrlShortenerServiceServer) ShortenUrl(context.Context, *ShortenUrlRequest) (*ShortenedUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenUrl not implemented")
}
func (UnimplementedUrlShortenerServiceServer) LookupUrl(context.Context, *LookupUrlRequest) (*ShortenedUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupUrl not implemented")
}
func (UnimplementedUrlShortenerServiceServer) mustEmbedUnimplementedUrlShortenerServiceServer() {}

// UnsafeUrlShortenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortenerServiceServer will
// result in compilation errors.
type UnsafeUrlShortenerServiceServer interface {
	mustEmbedUnimplementedUrlShortenerServiceServer()
}

func RegisterUrlShortenerServiceServer(s grpc.ServiceRegistrar, srv UrlShortenerServiceServer) {
	s.RegisterService(&UrlShortenerService_ServiceDesc, srv)
}

func _UrlShortenerService_ShortenUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServiceServer).ShortenUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UrlShortenerService/ShortenUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServiceServer).ShortenUrl(ctx, req.(*ShortenUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortenerService_LookupUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServiceServer).LookupUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UrlShortenerService/LookupUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServiceServer).LookupUrl(ctx, req.(*LookupUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortenerService_ServiceDesc is the grpc.ServiceDesc for UrlShortenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UrlShortenerService",
	HandlerType: (*UrlShortenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShortenUrl",
			Handler:    _UrlShortenerService_ShortenUrl_Handler,
		},
		{
			MethodName: "LookupUrl",
			Handler:    _UrlShortenerService_LookupUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adapters/driving/grpc/models/url.proto",
}