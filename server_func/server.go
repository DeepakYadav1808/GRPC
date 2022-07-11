package server_func

import (
	"context"
	"fmt"
	"io"
	"log"

	database "github.com/bookstore/Database"
	"github.com/bookstore/request"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/bookstore/bookstore_pb"
)

type IserverRepository interface {
	CreateBook(ctx context.Context, booksRequest *bookstore_pb.BooksRequest) (*bookstore_pb.Response, error)
	GetBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Book, error)
	DeleteBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Response, error)
	Upatebook(ctx context.Context, updateBookReq *bookstore_pb.UpdateBookRequest) (*bookstore_pb.Response, error)
	PutContent(stream bookstore_pb.Bookstore_PutContentServer) error
	GetContent(req *bookstore_pb.Pagerequest, stream bookstore_pb.Bookstore_GetContentServer) error
	GetBookdetails(stream bookstore_pb.Bookstore_GetBookdetailsServer) error
}

type SeverRepositoryImpl struct {
	DBService database.DBService
	bookstore_pb.UnimplementedBookstoreServer
}

func NewRepository() *SeverRepositoryImpl {
	return &SeverRepositoryImpl{
		DBService: database.DBService{},
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SeverRepositoryImpl) CreateBook(ctx context.Context, booksRequest *bookstore_pb.BooksRequest) (*bookstore_pb.Response, error) {

	serverReq := request.Book{
		BookID:   booksRequest.GetBooks().BookId,
		Author:   booksRequest.GetBooks().Author,
		Title:    booksRequest.GetBooks().Title,
		BookName: booksRequest.GetBooks().Bookname,
	}

	err := s.DBService.GetDB().Table(request.TABLE_BOOK).Create(&serverReq)
	handleError(err.Error)

	createResponse := bookstore_pb.Response{
		Message: "Book Created succesfully",
	}
	return &createResponse, nil
}
func (s *SeverRepositoryImpl) GetBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Book, error) {
	serReq := bookIDReq.GetBookID()
	var book bookstore_pb.Book
	row, err := s.DBService.GetDB().Table(request.TABLE_BOOK).Where(request.TABLE_BOOK_BOOKID+"=?", serReq).Rows()
	handleError(err)
	for row.Next() {

		if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
			handleError(err)
		}

	}

	return &book, nil

}
func (s *SeverRepositoryImpl) DeleteBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Response, error) {

	query := fmt.Sprintf("DELETE FROM  %s WHERE BOOK_ID= '%s'", request.TABLE_BOOK, bookIDReq.GetBookID())

	tx := s.DBService.GetDB().Begin()
	defer tx.Rollback()
	err := tx.Exec(query)
	handleError(err.Error)
	tx.Commit()
	createResponse := bookstore_pb.Response{
		Message: "Book Created succesfully",
	}
	return &createResponse, nil
}
func (s *SeverRepositoryImpl) Upatebook(ctx context.Context, updateBookReq *bookstore_pb.UpdateBookRequest) (*bookstore_pb.Response, error) {

	stmt := fmt.Sprintf("UPDATE book SET title = '%s', author = '%s', book_name='%s' WHERE book_id='%s'", updateBookReq.Title, updateBookReq.Author, updateBookReq.Bookname, updateBookReq.PreviousBookID)
	err := s.DBService.GetDB().Exec(stmt)
	handleError(err.Error)
	createResponse := bookstore_pb.Response{
		Message: "Book updated succesfully",
	}
	return &createResponse, nil
}
func (s *SeverRepositoryImpl) PutContent(stream bookstore_pb.Bookstore_PutContentServer) error {

	serverResponse := ""
	for {

		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&bookstore_pb.Response{
				Message: serverResponse,
			})
		}
		handleError(err)
		s.PageContent(context.TODO(), msg)
		serverResponse += fmt.Sprintf("Page number = %v Page Content =%v of book ID = %v", msg.GetPageNumber(), msg.GetPagecontent(), msg.GetBookID())

	}

}

func (s *SeverRepositoryImpl) GetContent(req *bookstore_pb.Pagerequest, stream bookstore_pb.Bookstore_GetContentServer) error {
	bookID := req.BooKID
	pageNumber := req.Pagenumber
	stmt := fmt.Sprintf("select page_content from page where book_id='%s' and page_number='%s'", bookID, pageNumber)
	row, err := s.DBService.GetDB().DB().Query(stmt)
	handleError(err)
	var content []string
	for row.Next() {
		var s string
		if err := row.Scan(&s); err != nil {
			handleError(err)
		}
		content = append(content, s)
	}
	for _, v := range content {
		item := &bookstore_pb.Pagecontent{
			Content: v,
		}
		stream.Send(item)
	}
	return nil
}
func (s *SeverRepositoryImpl) PageContent(ctx context.Context, value *bookstore_pb.PageInfoRequest) {

	streamReq := request.Page{
		BookID:      value.GetBookID(),
		PageNumber:  value.GetPageNumber(),
		PageSize:    value.GetPagesize(),
		PageContent: value.GetPagecontent(),
	}
	// page contetent insertation
	//   for {
	// 	      if (ExistPage(pagenumber)==nil) { msg="book is full" break}
	//       r=getcurrentpageContent(pagenumber, bookid ) ,r2=getPagesize(page number, bookid)
	//             if r<r2 (CreateContent )
	//                            else { pagenumber+1 continue   }
	//             return content added

	err := s.DBService.GetDB().Table(request.TABLE_PAGE).Create(&streamReq)
	handleError(err.Error)

}

func (s *SeverRepositoryImpl) GetBookdetails(stream bookstore_pb.Bookstore_GetBookdetailsServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {

			return nil

		}
		u := req.GetRequest()
		switch v := u.(type) {
		case *bookstore_pb.Input_Author:
			var book bookstore_pb.Book
			var a bookstore_pb.Streamresponse
			var flag bool
			row, err := s.DBService.GetDB().Table(request.TABLE_BOOK).Select("*").Where(request.TABLE_BOOK_AUTHOR+"= ?", v.Author).Rows()
			if err != nil {
				return err
				//stream.Send(streamerr(err)
			}
			for row.Next() {
				if err = row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
					return err
					//stream.Send(streamerr(err)

				}
			}
			if flag != true {
				a = bookstore_pb.Streamresponse{
					Resp: &bookstore_pb.Streamresponse_BookrResp{
						BookrResp: &book,
					},
				}
			}

			err = stream.Send(&a)
			handleError(err)

		case *bookstore_pb.Input_Bookname:
			var book bookstore_pb.Book
			row, err := s.DBService.GetDB().Table(request.TABLE_BOOK).Select("*").Where(request.TABLE_BOOK_BOOKNAME+"= ?", v.Bookname).Rows()
			if err != nil {
				return err
			}
			for row.Next() {
				if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
					return err
				}
			}
			a := bookstore_pb.Streamresponse{
				Resp: &bookstore_pb.Streamresponse_BookrResp{
					BookrResp: &book,
				},
			}

			stream.Send(&a)
		}
		var book bookstore_pb.Book
		row, err := s.DBService.GetDB().Table(request.TABLE_BOOK).Select("*").Where(request.TABLE_BOOK_BOOKID+"= ?", req.BookID).Rows()
		if err != nil {
			handleError(err)
			//stream.Send(streamerr(err)
		}
		for row.Next() {
			if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
				handleError(err)
				//stream.Send(streamerr(err)

			}
		}
		a := bookstore_pb.Streamresponse{
			Resp: &bookstore_pb.Streamresponse_BookrResp{
				BookrResp: &book,
			},
		}
		stream.Send(&a)

	}

}

func streamerr(err error) *bookstore_pb.Streamresponse {

	streamerr := bookstore_pb.Streamresponse{

		Resp: &bookstore_pb.Streamresponse_Errormsg{
			Errormsg: err.Error(),
		},
	}
	return &streamerr
}
