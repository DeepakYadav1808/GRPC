package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	database "github.com/bookstore/Database"
	"github.com/bookstore/request"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/bookstore/bookstore_pb"
	"google.golang.org/grpc"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	bookstore_pb.UnimplementedBookstoreServer
	conn *gorm.DB
}

func newServer() *server {
	return &(server{})
}
func initserver() *server {

	c := new(database.Service)
	serv := newServer()
	serv.conn = c.Init()
	return serv
}
func (s *server) CreateBook(ctx context.Context, booksRequest *bookstore_pb.BooksRequest) (*bookstore_pb.Response, error) {

	serverReq := request.Book{
		BookID:   booksRequest.GetBooks().BookId,
		Author:   booksRequest.GetBooks().Author,
		Title:    booksRequest.GetBooks().Title,
		BookName: booksRequest.GetBooks().Bookname,
	}

	err := s.conn.Table(request.TABLE_BOOK).Create(&serverReq)
	handleError(err.Error)

	createResponse := bookstore_pb.Response{
		Message: "Book Created succesfully",
	}
	return &createResponse, nil
}
func (s *server) GetBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Book, error) {
	serReq := bookIDReq.GetBookID()
	var book bookstore_pb.Book
	row, err := s.conn.Table(request.TABLE_BOOK).Where(request.TABLE_BOOK_BOOKID+"=?", serReq).Rows()
	handleError(err)
	for row.Next() {

		if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
			handleError(err)
		}

	}

	return &book, nil

}
func (s *server) DeleteBook(ctx context.Context, bookIDReq *bookstore_pb.ID) (*bookstore_pb.Response, error) {

	query := fmt.Sprintf("DELETE FROM  %s WHERE BOOK_ID= '%s'", request.TABLE_BOOK, bookIDReq.GetBookID())

	tx := s.conn.Begin()
	defer tx.Rollback()
	err := tx.Exec(query)
	handleError(err.Error)
	tx.Commit()
	createResponse := bookstore_pb.Response{
		Message: "Book Created succesfully",
	}
	return &createResponse, nil
}
func (s *server) Upatebook(ctx context.Context, updateBookReq *bookstore_pb.UpdateBookRequest) (*bookstore_pb.Response, error) {

	stmt := fmt.Sprintf("UPDATE book SET title = '%s', author = '%s', book_name='%s' WHERE book_id='%s'", updateBookReq.Title, updateBookReq.Author, updateBookReq.Bookname, updateBookReq.PreviousBookID)
	err := s.conn.Exec(stmt)
	handleError(err.Error)
	createResponse := bookstore_pb.Response{
		Message: "Book updated succesfully",
	}
	return &createResponse, nil
}
func (s *server) PutContent(stream bookstore_pb.Bookstore_PutContentServer) error {

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

func (s *server) GetContent(req *bookstore_pb.Pagerequest, stream bookstore_pb.Bookstore_GetContentServer) error {
	bookID := req.BooKID
	pageNumber := req.Pagenumber
	stmt := fmt.Sprintf("select page_content from page where book_id='%s' and page_number='%s'", bookID, pageNumber)
	row, err := s.conn.DB().Query(stmt)
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
func (s *server) PageContent(ctx context.Context, value *bookstore_pb.PageInfoRequest) {

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

	err := s.conn.Table(request.TABLE_PAGE).Create(&streamReq)
	handleError(err.Error)

}

func (s *server) GetBookdetails(stream bookstore_pb.Bookstore_GetBookdetailsServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		var book bookstore_pb.Book
		row, err := s.conn.Table(request.TABLE_BOOK).Where(request.TABLE_BOOK_BOOKID+"=?", req.BookID).Rows()
		handleError(err)
		for row.Next() {

			if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
				handleError(err)
			}

		}
		stream.Send(&book)
	}
}
func main() {

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	newServer := initserver()

	bookstore_pb.RegisterBookstoreServer(s, newServer)
	handleError(err)
	go func() {
		if err := s.Serve(lis); err != nil {
			handleError(err)
		}
	}()

	fmt.Println("Postgress connected")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("\nclosing the server")
}
