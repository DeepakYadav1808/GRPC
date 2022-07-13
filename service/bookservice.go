package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	database "github.com/bookstore/database"

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

func (s *Bookserver) CreateBook(ctx context.Context, booksRequest *bookstore_pb.BooksRequest) (*bookstore_pb.ID, error) {

	if booksRequest == nil {
		return &bookstore_pb.ID{}, errors.New("book request is empty")
	}
	serverReq := model.Books{
		BookId:   booksRequest.GetBooks().BookId,
		Author:   booksRequest.GetBooks().Author,
		Title:    booksRequest.GetBooks().Title,
		Bookname: booksRequest.GetBooks().Bookname,
	}
	err := database.CreateBook(serverReq)
	if err != nil {
		return &bookstore_pb.ID{}, err
	}
	createResponse := bookstore_pb.ID{
		BookID: booksRequest.GetBooks().BookId,
	}
	return &createResponse, nil

}
func (s *Bookserver) GetBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Book, error) {

	serReq := bookIDReq.GetBookID()
	if serReq == "" {
		return &bookstore_pb.Book{}, errors.New("Bookid request is empty")

	}
	bookDetails, err := database.Getbook(serReq)
	handleError(err)
	booksinfo := bookstore_pb.Book{
		BookId:   bookDetails.BookId,
		Bookname: bookDetails.Bookname,
		Title:    bookDetails.Title,
		Author:   bookDetails.Author,
	}
	return &booksinfo, nil

}
func (s *Bookserver) DeleteBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Response, error) {

	bookID := bookIDReq.GetBookID()
	if bookID == "" {
		createResponse := bookstore_pb.Response{
			Message: "Invalid bookID ",
		}

		return &createResponse, nil
	}
	err := database.DeleteBook(bookID)
	if err != nil {
		createResponse := bookstore_pb.Response{
			Message: "error while deleting the record",
		}
		return &createResponse, err
	}
	createResponse := bookstore_pb.Response{
		Message: "Book Created succesfully",
	}
	return &createResponse, nil
}
func (s *Bookserver) Upatebook(ctx context.Context, updateBookReq *bookstore_pb.UpdateBookRequest) (*bookstore_pb.Response, error) {

	if updateBookReq == nil {
		return &bookstore_pb.Response{}, errors.New("Invalid update book request")

	}
	err := database.Updatebook(updateBookReq)
	if err != nil {
		createResponse := bookstore_pb.Response{
			Message: "error while updating the book record",
		}
		return &createResponse, err
	}

	createResponse := bookstore_pb.Response{
		Message: "Book updated succesfully",
	}
	return &createResponse, nil
}
func (s *Bookserver) PutContent(stream bookstore_pb.Bookstore_PutContentServer) error {

	serverResponse := ""
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		if err == io.EOF {
			return stream.SendAndClose(&bookstore_pb.Response{
				Message: serverResponse,
			})

		}
		handleError(err)
		//s.PageContent(context.TODO(), msg)
		serverResponse += fmt.Sprintf("Page number = %v Page Content =%v of book ID = %v", msg.GetPageNumber(), msg.GetPagecontent(), msg.GetBookID())

	}

}

func (s *Bookserver) GetContent(req *bookstore_pb.Pagerequest, stream bookstore_pb.Bookstore_GetContentServer) error {

	if req.GetPagesize() < 0 || req.GetPgagenumber() < 0 {
		return errors.New("invalid request ")
	}
	pageSize := req.GetPagesize()
	pageNumber := req.GetPgagenumber()
	bookDetails, err := database.Pagination(pageSize, pageNumber)
	if err != nil {
		return err
	}
	for _, v := range bookDetails {
		item := &bookstore_pb.Book{
			BookId:   v.BookId,
			Bookname: v.Bookname,
			Title:    v.Title,
			Author:   v.Author,
		}
		err := stream.Send(item)
		if err != nil {
			return err
		}
	}
	return nil

}

func (s *Bookserver) GetBookdetails(stream bookstore_pb.Bookstore_GetBookdetailsServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {

			return nil
		}
		if err != nil {

			return err
		}
		u := req.GetRequest()
		switch v := u.(type) {
		case *bookstore_pb.Input_Search:
			data, err2 := database.RetrieveData(model.TABLE_BOOK_AUTHOR, v.Search)
			handleError(err2)
			err = stream.Send(data)
			handleError(err)

		case *bookstore_pb.Input_Bookid:
			data, err2 := database.RetrieveData(model.TABLE_BOOK_BOOKID, v.Bookid)
			handleError(err2)
			err = stream.Send(data)
			handleError(err)
		}

	}
}
