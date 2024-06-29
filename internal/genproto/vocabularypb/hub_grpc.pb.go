// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: vocabularypb/hub.proto

package vocabularypb

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
	VocabularyService_SearchVocabulary_FullMethodName               = "/vocabularypb.VocabularyService/SearchVocabulary"
	VocabularyService_CreateCollection_FullMethodName               = "/vocabularypb.VocabularyService/CreateCollection"
	VocabularyService_UpdateCollection_FullMethodName               = "/vocabularypb.VocabularyService/UpdateCollection"
	VocabularyService_AddVocabularyToCollection_FullMethodName      = "/vocabularypb.VocabularyService/AddVocabularyToCollection"
	VocabularyService_RemoveVocabularyFromCollection_FullMethodName = "/vocabularypb.VocabularyService/RemoveVocabularyFromCollection"
	VocabularyService_CreateCommunitySentenceDraft_FullMethodName   = "/vocabularypb.VocabularyService/CreateCommunitySentenceDraft"
	VocabularyService_UpdateCommunitySentenceDraft_FullMethodName   = "/vocabularypb.VocabularyService/UpdateCommunitySentenceDraft"
	VocabularyService_PromoteCommunitySentenceDraft_FullMethodName  = "/vocabularypb.VocabularyService/PromoteCommunitySentenceDraft"
)

// VocabularyServiceClient is the client API for VocabularyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VocabularyServiceClient interface {
	SearchVocabulary(ctx context.Context, in *SearchVocabularyRequest, opts ...grpc.CallOption) (*SearchVocabularyResponse, error)
	CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	UpdateCollection(ctx context.Context, in *UpdateCollectionRequest, opts ...grpc.CallOption) (*UpdateCollectionResponse, error)
	AddVocabularyToCollection(ctx context.Context, in *AddVocabularyToCollectionRequest, opts ...grpc.CallOption) (*AddVocabularyToCollectionResponse, error)
	RemoveVocabularyFromCollection(ctx context.Context, in *RemoveVocabularyFromCollectionRequest, opts ...grpc.CallOption) (*RemoveVocabularyFromCollectionResponse, error)
	CreateCommunitySentenceDraft(ctx context.Context, in *CreateCommunitySentenceDraftRequest, opts ...grpc.CallOption) (*CreateCommunitySentenceDraftResponse, error)
	UpdateCommunitySentenceDraft(ctx context.Context, in *UpdateCommunitySentenceDraftRequest, opts ...grpc.CallOption) (*UpdateCommunitySentenceDraftResponse, error)
	PromoteCommunitySentenceDraft(ctx context.Context, in *PromoteCommunitySentenceDraftRequest, opts ...grpc.CallOption) (*PromoteCommunitySentenceDraftResponse, error)
}

type vocabularyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVocabularyServiceClient(cc grpc.ClientConnInterface) VocabularyServiceClient {
	return &vocabularyServiceClient{cc}
}

func (c *vocabularyServiceClient) SearchVocabulary(ctx context.Context, in *SearchVocabularyRequest, opts ...grpc.CallOption) (*SearchVocabularyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchVocabularyResponse)
	err := c.cc.Invoke(ctx, VocabularyService_SearchVocabulary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, VocabularyService_CreateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) UpdateCollection(ctx context.Context, in *UpdateCollectionRequest, opts ...grpc.CallOption) (*UpdateCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCollectionResponse)
	err := c.cc.Invoke(ctx, VocabularyService_UpdateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) AddVocabularyToCollection(ctx context.Context, in *AddVocabularyToCollectionRequest, opts ...grpc.CallOption) (*AddVocabularyToCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVocabularyToCollectionResponse)
	err := c.cc.Invoke(ctx, VocabularyService_AddVocabularyToCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) RemoveVocabularyFromCollection(ctx context.Context, in *RemoveVocabularyFromCollectionRequest, opts ...grpc.CallOption) (*RemoveVocabularyFromCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveVocabularyFromCollectionResponse)
	err := c.cc.Invoke(ctx, VocabularyService_RemoveVocabularyFromCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) CreateCommunitySentenceDraft(ctx context.Context, in *CreateCommunitySentenceDraftRequest, opts ...grpc.CallOption) (*CreateCommunitySentenceDraftResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCommunitySentenceDraftResponse)
	err := c.cc.Invoke(ctx, VocabularyService_CreateCommunitySentenceDraft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) UpdateCommunitySentenceDraft(ctx context.Context, in *UpdateCommunitySentenceDraftRequest, opts ...grpc.CallOption) (*UpdateCommunitySentenceDraftResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCommunitySentenceDraftResponse)
	err := c.cc.Invoke(ctx, VocabularyService_UpdateCommunitySentenceDraft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) PromoteCommunitySentenceDraft(ctx context.Context, in *PromoteCommunitySentenceDraftRequest, opts ...grpc.CallOption) (*PromoteCommunitySentenceDraftResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromoteCommunitySentenceDraftResponse)
	err := c.cc.Invoke(ctx, VocabularyService_PromoteCommunitySentenceDraft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VocabularyServiceServer is the server API for VocabularyService service.
// All implementations should embed UnimplementedVocabularyServiceServer
// for forward compatibility
type VocabularyServiceServer interface {
	SearchVocabulary(context.Context, *SearchVocabularyRequest) (*SearchVocabularyResponse, error)
	CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	UpdateCollection(context.Context, *UpdateCollectionRequest) (*UpdateCollectionResponse, error)
	AddVocabularyToCollection(context.Context, *AddVocabularyToCollectionRequest) (*AddVocabularyToCollectionResponse, error)
	RemoveVocabularyFromCollection(context.Context, *RemoveVocabularyFromCollectionRequest) (*RemoveVocabularyFromCollectionResponse, error)
	CreateCommunitySentenceDraft(context.Context, *CreateCommunitySentenceDraftRequest) (*CreateCommunitySentenceDraftResponse, error)
	UpdateCommunitySentenceDraft(context.Context, *UpdateCommunitySentenceDraftRequest) (*UpdateCommunitySentenceDraftResponse, error)
	PromoteCommunitySentenceDraft(context.Context, *PromoteCommunitySentenceDraftRequest) (*PromoteCommunitySentenceDraftResponse, error)
}

// UnimplementedVocabularyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVocabularyServiceServer struct {
}

func (UnimplementedVocabularyServiceServer) SearchVocabulary(context.Context, *SearchVocabularyRequest) (*SearchVocabularyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchVocabulary not implemented")
}
func (UnimplementedVocabularyServiceServer) CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedVocabularyServiceServer) UpdateCollection(context.Context, *UpdateCollectionRequest) (*UpdateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollection not implemented")
}
func (UnimplementedVocabularyServiceServer) AddVocabularyToCollection(context.Context, *AddVocabularyToCollectionRequest) (*AddVocabularyToCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVocabularyToCollection not implemented")
}
func (UnimplementedVocabularyServiceServer) RemoveVocabularyFromCollection(context.Context, *RemoveVocabularyFromCollectionRequest) (*RemoveVocabularyFromCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVocabularyFromCollection not implemented")
}
func (UnimplementedVocabularyServiceServer) CreateCommunitySentenceDraft(context.Context, *CreateCommunitySentenceDraftRequest) (*CreateCommunitySentenceDraftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunitySentenceDraft not implemented")
}
func (UnimplementedVocabularyServiceServer) UpdateCommunitySentenceDraft(context.Context, *UpdateCommunitySentenceDraftRequest) (*UpdateCommunitySentenceDraftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommunitySentenceDraft not implemented")
}
func (UnimplementedVocabularyServiceServer) PromoteCommunitySentenceDraft(context.Context, *PromoteCommunitySentenceDraftRequest) (*PromoteCommunitySentenceDraftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteCommunitySentenceDraft not implemented")
}

// UnsafeVocabularyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VocabularyServiceServer will
// result in compilation errors.
type UnsafeVocabularyServiceServer interface {
	mustEmbedUnimplementedVocabularyServiceServer()
}

func RegisterVocabularyServiceServer(s grpc.ServiceRegistrar, srv VocabularyServiceServer) {
	s.RegisterService(&VocabularyService_ServiceDesc, srv)
}

func _VocabularyService_SearchVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchVocabularyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).SearchVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_SearchVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).SearchVocabulary(ctx, req.(*SearchVocabularyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_CreateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).CreateCollection(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_UpdateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).UpdateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_UpdateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).UpdateCollection(ctx, req.(*UpdateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_AddVocabularyToCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVocabularyToCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).AddVocabularyToCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_AddVocabularyToCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).AddVocabularyToCollection(ctx, req.(*AddVocabularyToCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_RemoveVocabularyFromCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveVocabularyFromCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).RemoveVocabularyFromCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_RemoveVocabularyFromCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).RemoveVocabularyFromCollection(ctx, req.(*RemoveVocabularyFromCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_CreateCommunitySentenceDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunitySentenceDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).CreateCommunitySentenceDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_CreateCommunitySentenceDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).CreateCommunitySentenceDraft(ctx, req.(*CreateCommunitySentenceDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_UpdateCommunitySentenceDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommunitySentenceDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).UpdateCommunitySentenceDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_UpdateCommunitySentenceDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).UpdateCommunitySentenceDraft(ctx, req.(*UpdateCommunitySentenceDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_PromoteCommunitySentenceDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteCommunitySentenceDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).PromoteCommunitySentenceDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_PromoteCommunitySentenceDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).PromoteCommunitySentenceDraft(ctx, req.(*PromoteCommunitySentenceDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VocabularyService_ServiceDesc is the grpc.ServiceDesc for VocabularyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VocabularyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vocabularypb.VocabularyService",
	HandlerType: (*VocabularyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchVocabulary",
			Handler:    _VocabularyService_SearchVocabulary_Handler,
		},
		{
			MethodName: "CreateCollection",
			Handler:    _VocabularyService_CreateCollection_Handler,
		},
		{
			MethodName: "UpdateCollection",
			Handler:    _VocabularyService_UpdateCollection_Handler,
		},
		{
			MethodName: "AddVocabularyToCollection",
			Handler:    _VocabularyService_AddVocabularyToCollection_Handler,
		},
		{
			MethodName: "RemoveVocabularyFromCollection",
			Handler:    _VocabularyService_RemoveVocabularyFromCollection_Handler,
		},
		{
			MethodName: "CreateCommunitySentenceDraft",
			Handler:    _VocabularyService_CreateCommunitySentenceDraft_Handler,
		},
		{
			MethodName: "UpdateCommunitySentenceDraft",
			Handler:    _VocabularyService_UpdateCommunitySentenceDraft_Handler,
		},
		{
			MethodName: "PromoteCommunitySentenceDraft",
			Handler:    _VocabularyService_PromoteCommunitySentenceDraft_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vocabularypb/hub.proto",
}
