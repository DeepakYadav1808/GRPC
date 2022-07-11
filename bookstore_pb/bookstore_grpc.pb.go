// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: bookstore_pb/bookstore.proto

package bookstore_pb

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

// BookstoreClient is the client API for Bookstore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookstoreClient interface {
	CreateBook(ctx context.Context, in *BooksRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteBook(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error)
	GetBook(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error)
	Upatebook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Response, error)
	PutContent(ctx context.Context, opts ...grpc.CallOption) (Bookstore_PutContentClient, error)
	GetContent(ctx context.Context, in *Pagerequest, opts ...grpc.CallOption) (Bookstore_GetContentClient, error)
	GetBookdetails(ctx context.Context, opts ...grpc.CallOption) (Bookstore_GetBookdetailsClient, error)
}

type bookstoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBookstoreClient(cc grpc.ClientConnInterface) BookstoreClient {
	return &bookstoreClient{cc}
}

func (c *bookstoreClient) CreateBook(ctx context.Context, in *BooksRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteBook(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBook(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) Upatebook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/Upatebook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) PutContent(ctx context.Context, opts ...grpc.CallOption) (Bookstore_PutContentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bookstore_ServiceDesc.Streams[0], "/bookstore.Bookstore/PutContent", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookstorePutContentClient{stream}
	return x, nil
}

type Bookstore_PutContentClient interface {
	Send(*PageInfoRequest) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type bookstorePutContentClient struct {
	grpc.ClientStream
}

func (x *bookstorePutContentClient) Send(m *PageInfoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookstorePutContentClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookstoreClient) GetContent(ctx context.Context, in *Pagerequest, opts ...grpc.CallOption) (Bookstore_GetContentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bookstore_ServiceDesc.Streams[1], "/bookstore.Bookstore/GetContent", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookstoreGetContentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bookstore_GetContentClient interface {
	Recv() (*Pagecontent, error)
	grpc.ClientStream
}

type bookstoreGetContentClient struct {
	grpc.ClientStream
}

func (x *bookstoreGetContentClient) Recv() (*Pagecontent, error) {
	m := new(Pagecontent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookstoreClient) GetBookdetails(ctx context.Context, opts ...grpc.CallOption) (Bookstore_GetBookdetailsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bookstore_ServiceDesc.Streams[2], "/bookstore.Bookstore/GetBookdetails", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookstoreGetBookdetailsClient{stream}
	return x, nil
}

type Bookstore_GetBookdetailsClient interface {
	Send(*Input) error
	Recv() (*Streamresponse, error)
	grpc.ClientStream
}

type bookstoreGetBookdetailsClient struct {
	grpc.ClientStream
}

func (x *bookstoreGetBookdetailsClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookstoreGetBookdetailsClient) Recv() (*Streamresponse, error) {
	m := new(Streamresponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookstoreServer is the server API for Bookstore service.
// All implementations must embed UnimplementedBookstoreServer
// for forward compatibility
type BookstoreServer interface {
	CreateBook(context.Context, *BooksRequest) (*Response, error)
	DeleteBook(context.Context, *ID) (*Response, error)
	GetBook(context.Context, *ID) (*Book, error)
	Upatebook(context.Context, *UpdateBookRequest) (*Response, error)
	PutContent(Bookstore_PutContentServer) error
	GetContent(*Pagerequest, Bookstore_GetContentServer) error
	GetBookdetails(Bookstore_GetBookdetailsServer) error
	mustEmbedUnimplementedBookstoreServer()
}

// UnimplementedBookstoreServer must be embedded to have forward compatible implementations.
type UnimplementedBookstoreServer struct {
}

func (UnimplementedBookstoreServer) CreateBook(context.Context, *BooksRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookstoreServer) DeleteBook(context.Context, *ID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookstoreServer) GetBook(context.Context, *ID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookstoreServer) Upatebook(context.Context, *UpdateBookRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upatebook not implemented")
}
func (UnimplementedBookstoreServer) PutContent(Bookstore_PutContentServer) error {
	return status.Errorf(codes.Unimplemented, "method PutContent not implemented")
}
func (UnimplementedBookstoreServer) GetContent(*Pagerequest, Bookstore_GetContentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedBookstoreServer) GetBookdetails(Bookstore_GetBookdetailsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBookdetails not implemented")
}
func (UnimplementedBookstoreServer) mustEmbedUnimplementedBookstoreServer() {}

// UnsafeBookstoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookstoreServer will
// result in compilation errors.
type UnsafeBookstoreServer interface {
	mustEmbedUnimplementedBookstoreServer()
}

func RegisterBookstoreServer(s grpc.ServiceRegistrar, srv BookstoreServer) {
	s.RegisterService(&Bookstore_ServiceDesc, srv)
}

func _Bookstore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBook(ctx, req.(*BooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteBook(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBook(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_Upatebook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).Upatebook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/Upatebook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).Upatebook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_PutContent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookstoreServer).PutContent(&bookstorePutContentServer{stream})
}

type Bookstore_PutContentServer interface {
	SendAndClose(*Response) error
	Recv() (*PageInfoRequest, error)
	grpc.ServerStream
}

type bookstorePutContentServer struct {
	grpc.ServerStream
}

func (x *bookstorePutContentServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookstorePutContentServer) Recv() (*PageInfoRequest, error) {
	m := new(PageInfoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Bookstore_GetContent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pagerequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookstoreServer).GetContent(m, &bookstoreGetContentServer{stream})
}

type Bookstore_GetContentServer interface {
	Send(*Pagecontent) error
	grpc.ServerStream
}

type bookstoreGetContentServer struct {
	grpc.ServerStream
}

func (x *bookstoreGetContentServer) Send(m *Pagecontent) error {
	return x.ServerStream.SendMsg(m)
}

func _Bookstore_GetBookdetails_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookstoreServer).GetBookdetails(&bookstoreGetBookdetailsServer{stream})
}

type Bookstore_GetBookdetailsServer interface {
	Send(*Streamresponse) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type bookstoreGetBookdetailsServer struct {
	grpc.ServerStream
}

func (x *bookstoreGetBookdetailsServer) Send(m *Streamresponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookstoreGetBookdetailsServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Bookstore_ServiceDesc is the grpc.ServiceDesc for Bookstore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookstore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookstore.Bookstore",
	HandlerType: (*BookstoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _Bookstore_CreateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Bookstore_DeleteBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Bookstore_GetBook_Handler,
		},
		{
			MethodName: "Upatebook",
			Handler:    _Bookstore_Upatebook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutContent",
			Handler:       _Bookstore_PutContent_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetContent",
			Handler:       _Bookstore_GetContent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBookdetails",
			Handler:       _Bookstore_GetBookdetails_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bookstore_pb/bookstore.proto",
}
