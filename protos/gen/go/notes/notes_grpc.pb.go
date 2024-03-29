// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: notes/notes.proto

package notes

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

// NotesClient is the client API for Notes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotesClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	GetNoteById(ctx context.Context, in *GetNoteByIdRequest, opts ...grpc.CallOption) (*Note, error)
	GetNotes(ctx context.Context, in *GetNotesRequest, opts ...grpc.CallOption) (*GetNotesResponse, error)
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*Note, error)
	DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*Note, error)
}

type notesClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesClient(cc grpc.ClientConnInterface) NotesClient {
	return &notesClient{cc}
}

func (c *notesClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, "/notes.Notes/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) GetNoteById(ctx context.Context, in *GetNoteByIdRequest, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/GetNoteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) GetNotes(ctx context.Context, in *GetNotesRequest, opts ...grpc.CallOption) (*GetNotesResponse, error) {
	out := new(GetNotesResponse)
	err := c.cc.Invoke(ctx, "/notes.Notes/GetNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServer is the server API for Notes service.
// All implementations must embed UnimplementedNotesServer
// for forward compatibility
type NotesServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	GetNoteById(context.Context, *GetNoteByIdRequest) (*Note, error)
	GetNotes(context.Context, *GetNotesRequest) (*GetNotesResponse, error)
	UpdateNote(context.Context, *UpdateNoteRequest) (*Note, error)
	DeleteNote(context.Context, *DeleteNoteRequest) (*Note, error)
	mustEmbedUnimplementedNotesServer()
}

// UnimplementedNotesServer must be embedded to have forward compatible implementations.
type UnimplementedNotesServer struct {
}

func (UnimplementedNotesServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNotesServer) GetNoteById(context.Context, *GetNoteByIdRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteById not implemented")
}
func (UnimplementedNotesServer) GetNotes(context.Context, *GetNotesRequest) (*GetNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotes not implemented")
}
func (UnimplementedNotesServer) UpdateNote(context.Context, *UpdateNoteRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedNotesServer) DeleteNote(context.Context, *DeleteNoteRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedNotesServer) mustEmbedUnimplementedNotesServer() {}

// UnsafeNotesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotesServer will
// result in compilation errors.
type UnsafeNotesServer interface {
	mustEmbedUnimplementedNotesServer()
}

func RegisterNotesServer(s grpc.ServiceRegistrar, srv NotesServer) {
	s.RegisterService(&Notes_ServiceDesc, srv)
}

func _Notes_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_GetNoteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).GetNoteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/GetNoteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).GetNoteById(ctx, req.(*GetNoteByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_GetNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).GetNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/GetNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).GetNotes(ctx, req.(*GetNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).UpdateNote(ctx, req.(*UpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).DeleteNote(ctx, req.(*DeleteNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notes_ServiceDesc is the grpc.ServiceDesc for Notes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notes.Notes",
	HandlerType: (*NotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _Notes_CreateNote_Handler,
		},
		{
			MethodName: "GetNoteById",
			Handler:    _Notes_GetNoteById_Handler,
		},
		{
			MethodName: "GetNotes",
			Handler:    _Notes_GetNotes_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _Notes_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _Notes_DeleteNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notes/notes.proto",
}
