// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: bookstore_pb/bookstore.proto

package bookstore_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId   string `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	Author   string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Bookname string `protobuf:"bytes,3,opt,name=bookname,proto3" json:"bookname,omitempty"`
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetBookname() string {
	if x != nil {
		return x.Bookname
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author   string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Bookname string `protobuf:"bytes,3,opt,name=bookname,proto3" json:"bookname,omitempty"`
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpdateBookRequest) GetBookname() string {
	if x != nil {
		return x.Bookname
	}
	return ""
}

func (x *UpdateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type BooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books *Book `protobuf:"bytes,1,opt,name=books,proto3" json:"books,omitempty"`
}

func (x *BooksRequest) Reset() {
	*x = BooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksRequest) ProtoMessage() {}

func (x *BooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksRequest.ProtoReflect.Descriptor instead.
func (*BooksRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{2}
}

func (x *BooksRequest) GetBooks() *Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *ID) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PageInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber  string `protobuf:"bytes,1,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	BookID      string `protobuf:"bytes,2,opt,name=BookID,proto3" json:"BookID,omitempty"`
	Pagesize    string `protobuf:"bytes,3,opt,name=Pagesize,proto3" json:"Pagesize,omitempty"`
	Pagecontent string `protobuf:"bytes,4,opt,name=pagecontent,proto3" json:"pagecontent,omitempty"`
}

func (x *PageInfoRequest) Reset() {
	*x = PageInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoRequest) ProtoMessage() {}

func (x *PageInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoRequest.ProtoReflect.Descriptor instead.
func (*PageInfoRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{5}
}

func (x *PageInfoRequest) GetPageNumber() string {
	if x != nil {
		return x.PageNumber
	}
	return ""
}

func (x *PageInfoRequest) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

func (x *PageInfoRequest) GetPagesize() string {
	if x != nil {
		return x.Pagesize
	}
	return ""
}

func (x *PageInfoRequest) GetPagecontent() string {
	if x != nil {
		return x.Pagecontent
	}
	return ""
}

type Pagerequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pgagenumber int64 `protobuf:"varint,1,opt,name=pgagenumber,proto3" json:"pgagenumber,omitempty"`
	Pagesize    int64 `protobuf:"varint,2,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
}

func (x *Pagerequest) Reset() {
	*x = Pagerequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagerequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagerequest) ProtoMessage() {}

func (x *Pagerequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagerequest.ProtoReflect.Descriptor instead.
func (*Pagerequest) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{6}
}

func (x *Pagerequest) GetPgagenumber() int64 {
	if x != nil {
		return x.Pgagenumber
	}
	return 0
}

func (x *Pagerequest) GetPagesize() int64 {
	if x != nil {
		return x.Pagesize
	}
	return 0
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*Input_Search
	//	*Input_Bookid
	Request isInput_Request `protobuf_oneof:"request"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{7}
}

func (m *Input) GetRequest() isInput_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Input) GetSearch() string {
	if x, ok := x.GetRequest().(*Input_Search); ok {
		return x.Search
	}
	return ""
}

func (x *Input) GetBookid() string {
	if x, ok := x.GetRequest().(*Input_Bookid); ok {
		return x.Bookid
	}
	return ""
}

type isInput_Request interface {
	isInput_Request()
}

type Input_Search struct {
	Search string `protobuf:"bytes,1,opt,name=search,proto3,oneof"`
}

type Input_Bookid struct {
	Bookid string `protobuf:"bytes,2,opt,name=bookid,proto3,oneof"`
}

func (*Input_Search) isInput_Request() {}

func (*Input_Bookid) isInput_Request() {}

type Streamresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resp:
	//	*Streamresponse_Errormsg
	//	*Streamresponse_BookrResp
	Resp isStreamresponse_Resp `protobuf_oneof:"resp"`
}

func (x *Streamresponse) Reset() {
	*x = Streamresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_pb_bookstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Streamresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Streamresponse) ProtoMessage() {}

func (x *Streamresponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_pb_bookstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Streamresponse.ProtoReflect.Descriptor instead.
func (*Streamresponse) Descriptor() ([]byte, []int) {
	return file_bookstore_pb_bookstore_proto_rawDescGZIP(), []int{8}
}

func (m *Streamresponse) GetResp() isStreamresponse_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *Streamresponse) GetErrormsg() string {
	if x, ok := x.GetResp().(*Streamresponse_Errormsg); ok {
		return x.Errormsg
	}
	return ""
}

func (x *Streamresponse) GetBookrResp() *Book {
	if x, ok := x.GetResp().(*Streamresponse_BookrResp); ok {
		return x.BookrResp
	}
	return nil
}

type isStreamresponse_Resp interface {
	isStreamresponse_Resp()
}

type Streamresponse_Errormsg struct {
	Errormsg string `protobuf:"bytes,1,opt,name=errormsg,proto3,oneof"`
}

type Streamresponse_BookrResp struct {
	BookrResp *Book `protobuf:"bytes,2,opt,name=bookrResp,proto3,oneof"`
}

func (*Streamresponse_Errormsg) isStreamresponse_Resp() {}

func (*Streamresponse_BookrResp) isStreamresponse_Resp() {}

var File_bookstore_pb_bookstore_proto protoreflect.FileDescriptor

var file_bookstore_pb_bookstore_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x62, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x68, 0x0a, 0x04, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x5d, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0x35, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x1c, 0x0a, 0x02, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x22, 0x24, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x87, 0x01,
	0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x67, 0x61, 0x67, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x67, 0x61,
	0x67, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x46, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x0e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x72, 0x52, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x0a,
	0x04, 0x72, 0x65, 0x73, 0x70, 0x32, 0xe2, 0x02, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0d, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09,
	0x55, 0x70, 0x61, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bookstore_pb_bookstore_proto_rawDescOnce sync.Once
	file_bookstore_pb_bookstore_proto_rawDescData = file_bookstore_pb_bookstore_proto_rawDesc
)

func file_bookstore_pb_bookstore_proto_rawDescGZIP() []byte {
	file_bookstore_pb_bookstore_proto_rawDescOnce.Do(func() {
		file_bookstore_pb_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookstore_pb_bookstore_proto_rawDescData)
	})
	return file_bookstore_pb_bookstore_proto_rawDescData
}

var file_bookstore_pb_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bookstore_pb_bookstore_proto_goTypes = []interface{}{
	(*Book)(nil),              // 0: bookstore.Book
	(*UpdateBookRequest)(nil), // 1: bookstore.updateBookRequest
	(*BooksRequest)(nil),      // 2: bookstore.BooksRequest
	(*ID)(nil),                // 3: bookstore.ID
	(*Response)(nil),          // 4: bookstore.response
	(*PageInfoRequest)(nil),   // 5: bookstore.PageInfoRequest
	(*Pagerequest)(nil),       // 6: bookstore.pagerequest
	(*Input)(nil),             // 7: bookstore.input
	(*Streamresponse)(nil),    // 8: bookstore.streamresponse
}
var file_bookstore_pb_bookstore_proto_depIdxs = []int32{
	0, // 0: bookstore.BooksRequest.books:type_name -> bookstore.Book
	0, // 1: bookstore.streamresponse.bookrResp:type_name -> bookstore.Book
	2, // 2: bookstore.Bookstore.CreateBook:input_type -> bookstore.BooksRequest
	3, // 3: bookstore.Bookstore.DeleteBook:input_type -> bookstore.ID
	3, // 4: bookstore.Bookstore.GetBook:input_type -> bookstore.ID
	1, // 5: bookstore.Bookstore.Upatebook:input_type -> bookstore.updateBookRequest
	6, // 6: bookstore.Bookstore.GetContent:input_type -> bookstore.pagerequest
	7, // 7: bookstore.Bookstore.GetBookdetails:input_type -> bookstore.input
	3, // 8: bookstore.Bookstore.CreateBook:output_type -> bookstore.ID
	4, // 9: bookstore.Bookstore.DeleteBook:output_type -> bookstore.response
	0, // 10: bookstore.Bookstore.GetBook:output_type -> bookstore.Book
	4, // 11: bookstore.Bookstore.Upatebook:output_type -> bookstore.response
	0, // 12: bookstore.Bookstore.GetContent:output_type -> bookstore.Book
	8, // 13: bookstore.Bookstore.GetBookdetails:output_type -> bookstore.streamresponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bookstore_pb_bookstore_proto_init() }
func file_bookstore_pb_bookstore_proto_init() {
	if File_bookstore_pb_bookstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookstore_pb_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagerequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_pb_bookstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Streamresponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_bookstore_pb_bookstore_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Input_Search)(nil),
		(*Input_Bookid)(nil),
	}
	file_bookstore_pb_bookstore_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Streamresponse_Errormsg)(nil),
		(*Streamresponse_BookrResp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bookstore_pb_bookstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookstore_pb_bookstore_proto_goTypes,
		DependencyIndexes: file_bookstore_pb_bookstore_proto_depIdxs,
		MessageInfos:      file_bookstore_pb_bookstore_proto_msgTypes,
	}.Build()
	File_bookstore_pb_bookstore_proto = out.File
	file_bookstore_pb_bookstore_proto_rawDesc = nil
	file_bookstore_pb_bookstore_proto_goTypes = nil
	file_bookstore_pb_bookstore_proto_depIdxs = nil
}