// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: search/search.proto

package searchv1

import (
	context "context"
	common "github.com/market-4-test/contract-go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Search_GetEntities_FullMethodName       = "/search.Search/GetEntities"
	Search_UpsertEntities_FullMethodName    = "/search.Search/UpsertEntities"
	Search_SetActiveEntities_FullMethodName = "/search.Search/SetActiveEntities"
)

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	GetEntities(ctx context.Context, in *GetEntitiesParams, opts ...grpc.CallOption) (*GetEntitiesResponse, error)
	UpsertEntities(ctx context.Context, in *GetEntitiesParams, opts ...grpc.CallOption) (*GetEntitiesResponse, error)
	SetActiveEntities(ctx context.Context, in *SetActiveEntitiesParams, opts ...grpc.CallOption) (*common.Status, error)
}

type searchClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchClient(cc grpc.ClientConnInterface) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) GetEntities(ctx context.Context, in *GetEntitiesParams, opts ...grpc.CallOption) (*GetEntitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEntitiesResponse)
	err := c.cc.Invoke(ctx, Search_GetEntities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) UpsertEntities(ctx context.Context, in *GetEntitiesParams, opts ...grpc.CallOption) (*GetEntitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEntitiesResponse)
	err := c.cc.Invoke(ctx, Search_UpsertEntities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) SetActiveEntities(ctx context.Context, in *SetActiveEntitiesParams, opts ...grpc.CallOption) (*common.Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Status)
	err := c.cc.Invoke(ctx, Search_SetActiveEntities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
// All implementations must embed UnimplementedSearchServer
// for forward compatibility.
type SearchServer interface {
	GetEntities(context.Context, *GetEntitiesParams) (*GetEntitiesResponse, error)
	UpsertEntities(context.Context, *GetEntitiesParams) (*GetEntitiesResponse, error)
	SetActiveEntities(context.Context, *SetActiveEntitiesParams) (*common.Status, error)
	mustEmbedUnimplementedSearchServer()
}

// UnimplementedSearchServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchServer struct{}

func (UnimplementedSearchServer) GetEntities(context.Context, *GetEntitiesParams) (*GetEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntities not implemented")
}
func (UnimplementedSearchServer) UpsertEntities(context.Context, *GetEntitiesParams) (*GetEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertEntities not implemented")
}
func (UnimplementedSearchServer) SetActiveEntities(context.Context, *SetActiveEntitiesParams) (*common.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetActiveEntities not implemented")
}
func (UnimplementedSearchServer) mustEmbedUnimplementedSearchServer() {}
func (UnimplementedSearchServer) testEmbeddedByValue()                {}

// UnsafeSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServer will
// result in compilation errors.
type UnsafeSearchServer interface {
	mustEmbedUnimplementedSearchServer()
}

func RegisterSearchServer(s grpc.ServiceRegistrar, srv SearchServer) {
	// If the following call pancis, it indicates UnimplementedSearchServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Search_ServiceDesc, srv)
}

func _Search_GetEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntitiesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).GetEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Search_GetEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).GetEntities(ctx, req.(*GetEntitiesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_UpsertEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntitiesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).UpsertEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Search_UpsertEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).UpsertEntities(ctx, req.(*GetEntitiesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_SetActiveEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetActiveEntitiesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).SetActiveEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Search_SetActiveEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).SetActiveEntities(ctx, req.(*SetActiveEntitiesParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Search_ServiceDesc is the grpc.ServiceDesc for Search service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Search_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntities",
			Handler:    _Search_GetEntities_Handler,
		},
		{
			MethodName: "UpsertEntities",
			Handler:    _Search_UpsertEntities_Handler,
		},
		{
			MethodName: "SetActiveEntities",
			Handler:    _Search_SetActiveEntities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search/search.proto",
}
