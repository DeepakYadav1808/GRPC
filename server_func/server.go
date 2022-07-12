package server_func

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	database "github.com/bookstore/database"

	"github.com/bookstore/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/bookstore/bookstore_pb"
)

type IserverRepository interface {
	CreateBook(ctx context.Context, booksRequest *bookstore_pb.BooksRequest) (*bookstore_pb.Input, error)
	GetBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Book, error)
	DeleteBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Response, error)
	Upatebook(ctx context.Context, updateBookReq *bookstore_pb.UpdateBookRequest) (*bookstore_pb.Response, error)
	PutContent(stream bookstore_pb.Bookstore_PutContentServer) error
	GetContent(req *bookstore_pb.Pagerequest, stream bookstore_pb.Bookstore_GetContentServer) error
	GetBookdetails(stream bookstore_pb.Bookstore_GetBookdetailsServer) error
}

type SeverRepositoryImpl struct {
	DBService *gorm.DB
	bookstore_pb.UnimplementedBookstoreServer
}

func NewRepository() *SeverRepositoryImpl {
	return &SeverRepositoryImpl{
		DBService: database.Initdb(),
	}
}

func handleError(err error) {
	if err != nil {
		log.Default().Println(err)
	}
}

func (s *SeverRepositoryImpl) CreateBook(ctx context.Context, booksRequest *bookstore_pb.BooksRequest) (*bookstore_pb.ID, error) {

	serverReq := model.Books{
		BookID:   booksRequest.GetBooks().BookId,
		Author:   booksRequest.GetBooks().Author,
		Title:    booksRequest.GetBooks().Title,
		BookName: booksRequest.GetBooks().Bookname,
	}
	err := s.DBService.Model(&model.Books{}).Create(&serverReq)
	if err != nil {
		return &bookstore_pb.ID{}, err.Error
	}
	createResponse := bookstore_pb.ID{
		BookID: booksRequest.GetBooks().BookId,
	}
	return &createResponse, nil

}
func (s *SeverRepositoryImpl) GetBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Book, error) {

	serReq := bookIDReq.GetBookID()
	var book bookstore_pb.Book
	row, err := s.DBService.Model(&model.Books{}).Where(model.TABLE_BOOK_BOOKID+"=?", serReq).Rows()
	if err != nil {
		return &bookstore_pb.Book{}, err
	}
	for row.Next() {
		if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
			if err != nil {
				return &bookstore_pb.Book{}, err
			}
		}
	}
	return &book, nil

}
func (s *SeverRepositoryImpl) DeleteBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Response, error) {
	tx := s.DBService.Begin()
	defer tx.Rollback()
	err := tx.Where(model.TABLE_PAGE_BOOKID+"=?", bookIDReq.GetBookID()).Delete(&model.Books{})
	if err.Error != nil {
		createResponse := bookstore_pb.Response{
			Message: "error while deleting the record",
		}
		return &createResponse, err.Error
	}
	tx.Commit()
	createResponse := bookstore_pb.Response{
		Message: "Book Created succesfully",
	}
	return &createResponse, nil
}
func (s *SeverRepositoryImpl) Upatebook(ctx context.Context, updateBookReq *bookstore_pb.UpdateBookRequest) (*bookstore_pb.Response, error) {

	newCol := make(map[string]interface{})
	newCol["title"] = updateBookReq.Title
	newCol["author"] = updateBookReq.Author
	newCol["book_name"] = updateBookReq.Bookname
	err := s.DBService.Model(&model.Books{}).Updates(newCol)
	if err.Error != nil {
		createResponse := bookstore_pb.Response{
			Message: "Book did not updated succesfully",
		}
		return &createResponse, err.Error
	}

	createResponse := bookstore_pb.Response{
		Message: "Book updated succesfully",
	}
	return &createResponse, nil
}
func (s *SeverRepositoryImpl) PutContent(stream bookstore_pb.Bookstore_PutContentServer) error {

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

func (s *SeverRepositoryImpl) GetContent(req *bookstore_pb.Pagerequest, stream bookstore_pb.Bookstore_GetContentServer) error {

	if req.GetPagesize() == -1 || req.GetPgagenumber() == -1 {
		return errors.New("invalid request ")
	}
	pageSize := req.GetPagesize()
	pageNumber := req.GetPgagenumber()
	var book []model.Books

	stmt := s.DBService.Limit(pageSize).Offset((pageNumber - 1) * pageSize) // , order(asc)
	result := stmt.Model(&model.Books{}).Find(&book)

	if result.Error != nil {
		return result.Error
	}

	for _, v := range book {
		item := &bookstore_pb.Book{
			BookId:   v.BookID,
			Bookname: v.BookName,
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

func (s *SeverRepositoryImpl) GetBookdetails(stream bookstore_pb.Bookstore_GetBookdetailsServer) error {
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
			data, err2 := s.retrieveData(model.TABLE_BOOK_AUTHOR, v.Search)
			handleError(err2)
			err = stream.Send(data)
			handleError(err)

		case *bookstore_pb.Input_Bookid:
			data, err2 := s.retrieveData(model.TABLE_BOOK_BOOKID, v.Bookid)
			handleError(err2)
			err = stream.Send(data)
			handleError(err)
		}

	}
}

func (s *SeverRepositoryImpl) retrieveData(column_name string, column_value string) (*bookstore_pb.Streamresponse, error) {
	var book bookstore_pb.Book
	row, err := s.DBService.Model(&model.Books{}).Select("*").Where(column_name+"= ?", column_value).Rows()
	if err != nil {
		return &bookstore_pb.Streamresponse{}, err
	}
	for row.Next() {
		if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
			return &bookstore_pb.Streamresponse{}, err

		}
	}
	bookDetail := bookstore_pb.Streamresponse{
		Resp: &bookstore_pb.Streamresponse_BookrResp{
			BookrResp: &book,
		},
	}
	return &bookDetail, nil
}
