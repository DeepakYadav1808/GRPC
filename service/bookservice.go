package service

import (
	"context"
	"errors"
	"io"
	"log"

	database "github.com/bookstore/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bookstore/model"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/bookstore/bookstore_pb"
)

type Bookserver struct {
	bookstore_pb.UnimplementedBookstoreServer
}

func handleError(err error) {
	if err != nil {
		log.Default().Println(err)
	}
}

func (s *Bookserver) CreateBook(ctx context.Context, booksRequest *bookstore_pb.BooksRequest) (*bookstore_pb.Id, error) {
	err := validate(booksRequest)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "request is empty")
	}

	serverReq := assignRequestFromclient(booksRequest)

	err = database.CreateBook(serverReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "error while creating bookstore")
	}

	createResponse := bookstore_pb.Id{
		Bookid: booksRequest.GetBooks().BookId,
	}

	return &createResponse, nil
}
func (s *Bookserver) GetBook(ctx context.Context, bookIDReq *bookstore_pb.Id) (*bookstore_pb.Book, error) {
	serReq := bookIDReq.GetBookid()
	if serReq == "" {
		return nil, status.Error(codes.InvalidArgument, "request is empty")
	}

	bookDetails, err := database.Getbook(serReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	booksinfo := assignBookRequestModel(bookDetails)

	return booksinfo, nil
}
func (s *Bookserver) DeleteBook(ctx context.Context, bookIDReq *bookstore_pb.Id) (*bookstore_pb.Response, error) {
	bookID := bookIDReq.GetBookid()
	if len(bookID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "bookid field is empty")
	}

	err := database.DeleteBook(bookID)
	if err != nil {
		return nil, status.Error(codes.Internal, "while deleting the record")
	}

	createResponse := createResponse("Book Created succesfully")

	return createResponse, nil
}
func (s *Bookserver) Upatebook(ctx context.Context, updateBookReq *bookstore_pb.UpdateBookRequest) (*bookstore_pb.Response, error) {
	err := validateUpdateRequest(updateBookReq)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "update request field is empty")
	}

	err = database.Updatebook(updateBookReq)
	if err != nil {
		return nil, status.Error(codes.Aborted, "error while updating the record")
	}

	createResponse := createResponse("Book updated succesfully")

	return createResponse, nil
}
func (s *Bookserver) GetContent(req *bookstore_pb.PageRequest, stream bookstore_pb.Bookstore_GetContentServer) error {
	if req.GetPagesize() < 0 || req.GetPgagenumber() < 0 {
		return status.Error(codes.InvalidArgument, "invalid request for pagenumber/pagesize")
	}
	pageSize := req.GetPagesize()
	pageNumber := req.GetPgagenumber()

	bookDetails, err := database.Pagination(pageSize, pageNumber)
	if err != nil {
		return status.Error(codes.Aborted, err.Error())
	}

	for _, v := range bookDetails {
		item := assignBookRequestModel(&v)

		err := stream.Send(item)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *Bookserver) GetBookdetails(stream bookstore_pb.Bookstore_GetBookdetailsServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		if err == io.EOF {
			return nil
		}
		option := req.GetRequest()
		switch optionvalue := option.(type) {
		case *bookstore_pb.Input_Search:
			data, err2 := database.RetrieveData(model.TABLE_BOOK_AUTHOR, optionvalue.Search)
			handleError(err2)

			err = stream.Send(data)

			handleError(err)

		case *bookstore_pb.Input_Bookid:
			data, err2 := database.RetrieveData(model.TABLE_BOOK_BOOKID, optionvalue.Bookid)
			handleError(err2)

			err = stream.Send(data)

			handleError(err)
		}
	}
}
func validate(booksRequest *bookstore_pb.BooksRequest) error {
	if len(booksRequest.GetBooks().BookId) == 0 ||
		len(booksRequest.GetBooks().Author) == 0 ||
		len(booksRequest.GetBooks().Title) == 0 ||
		len(booksRequest.GetBooks().Bookname) == 0 {
		return errors.New("book request field is empty")
	}

	return nil
}
func assignRequestFromclient(booksRequest *bookstore_pb.BooksRequest) model.Books {
	serverReq := model.Books{
		BookId:   booksRequest.GetBooks().BookId,
		Author:   booksRequest.GetBooks().Author,
		Title:    booksRequest.GetBooks().Title,
		BookName: booksRequest.GetBooks().Bookname,
	}

	return serverReq
}
func validateUpdateRequest(updateBookReq *bookstore_pb.UpdateBookRequest) error {
	if len(updateBookReq.Author) == 0 ||
		len(updateBookReq.Bookname) == 0 ||
		len(updateBookReq.Title) == 0 {
		return errors.New("Update book request field is empty")
	}

	return nil
}

func assignBookRequestModel(bookdetails *model.Books) *bookstore_pb.Book {
	booksinfo := bookstore_pb.Book{
		BookId:   bookdetails.BookId,
		Bookname: bookdetails.BookName,
		Title:    bookdetails.Title,
		Author:   bookdetails.Author,
	}

	return &booksinfo
}
func createResponse(errormsg string) *bookstore_pb.Response {
	createResponse := bookstore_pb.Response{
		Message: errormsg,
	}

	return &createResponse
}
