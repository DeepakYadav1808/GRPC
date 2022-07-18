package database_test

import (
	"context"
	"testing"

	"github.com/bookstore/bookstore_pb"
	"github.com/bookstore/database"
	"github.com/google/uuid"

	"github.com/bookstore/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	db "github.com/bookstore/database"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

// Test for CreateItem method when input has empty name

// Ge

func TestCreateBook(t *testing.T) {
	id := uuid.NewString()
	books := model.Books{
		Author:   "adads",
		Title:    "magic1",
		BookId:   id,
		BookName: "real",
	}
	err := db.CreateBook(books)
	require.NoError(t, err)

	resp, err1 := db.Getbook(id)
	assert.Equal(t, id, resp.BookId, "should be equal")
	assert.Equal(t, "adads", resp.Author, "should be equal")
	assert.Equal(t, "magic1", resp.Title, "should be equal")
	assert.Equal(t, "real", resp.BookName, "should be equal")
	require.NoError(t, err1)

	err2 := database.DeleteBook(id)
	require.NoError(t, err2)
	_, err1 = db.Getbook(id)
	if err1 == nil {
		t.Error("Record not deleted")
	}
}

func TestDelete(t *testing.T) {
	id := uuid.NewString()
	books := model.Books{
		Author:   "adads",
		Title:    "magic1",
		BookId:   id,
		BookName: "real",
	}
	err := db.CreateBook(books)
	require.NoError(t, err)

	err = database.DeleteBook("@3213")
	if err != nil {
		t.Error(err.Error())
	}

	_, err1 := database.Getbook(id)
	if err1 == nil {
		t.Error("record not deleted")
	}

}

func TestUpdatebook(t *testing.T) {
	id := uuid.NewString()
	books := model.Books{
		Author:   "adads",
		Title:    "magic1",
		BookId:   id,
		BookName: "real",
	}
	err := db.CreateBook(books)
	require.NoError(t, err)

	err = database.Updatebook(&bookstore_pb.UpdateBookRequest{
		Author:   "newauth",
		Title:    "newtitle",
		Bookname: "newbookname",
	})
	require.NoError(t, err)

	resp, err1 := database.Getbook(id)
	require.NoError(t, err1)
	assert.Equal(t, "newauth", resp.Author, "author should be equal")
	assert.Equal(t, "newtitle", resp.Title, "title should be equal")
	assert.Equal(t, "newbookname", resp.BookName, "bookanme should be equal")

	err = database.DeleteBook(id)
	require.NoError(t, err)

}
func TestGetDetails(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()
	c := bookstore_pb.NewBookstoreClient(conn)
   
	stream, err1 := c.GetBookdetails(context.Background())
	require.NoError(t, err1)
	b := bookstore_pb.Input{
		Request: &bookstore_pb.Input_Bookid{
			Bookid: "2",
		}}
	err1 = stream.Send(&b)
    require.NoError(t, err1)
	
	resp, err2 := stream.Recv()
	require.NoError(t, err2)
	assert.Equal(t, "2", resp.GetBookrResp().BookId, "should be equal")
	b = bookstore_pb.Input{
		Request: &bookstore_pb.Input_Search{
			Search: "auth",
		}}
	err2 = stream.Send(&b)
	require.NoError(t, err2)
	assert.Equal(t, "newauth", resp.GetBookrResp().Author, "should be equal")
}

// func makeStreamMock() *StreamMock {
// 	return &StreamMock{
// 		ctx:            context.Background(),
// 		recvToServer:   make(chan *bookstore_pb.Input, 10),
// 		sentFromServer: make(chan *bookstore_pb.Streamresponse, 10),
// 	}
// }

// type StreamMock struct {
// 	grpc.ServerStream
// 	ctx            context.Context
// 	recvToServer   chan *bookstore_pb.Input
// 	sentFromServer chan *bookstore_pb.Streamresponse
// }

// func (m *StreamMock) Context() context.Context {
// 	return m.ctx
// }
// func (m *StreamMock) Send(resp *bookstore_pb.Streamresponse) error {
// 	m.sentFromServer <- resp
// 	return nil
// }
// func (m *StreamMock) Recv() (*bookstore_pb.Input, error) {
// 	req, more := <-m.recvToServer
// 	if !more {
// 		return nil, errors.New("empty")
// 	}
// 	return req, nil
// }
// func (m *StreamMock) SendFromClient(req *bookstore_pb.Input) error {
// 	m.recvToServer <- req
// 	return nil
// }
// func (m *StreamMock) RecvToClient() (*bookstore_pb.Streamresponse, error) {
// 	response, more := <-m.sentFromServer
// 	if !more {
// 		return nil, errors.New("empty")
// 	}
// 	return response, nil
// }

// func createStream(t *testing.T) *StreamMock {
// 	stream := makeStreamMock()
// 	go func() {
// 		api := bookstore_pb.RegisterBookstoreServer()
// 		err := api.SumStream(stream)
// 		if err != nil {
// 			t.Errorf(err.Error())
// 		}
// 		close(stream.sentFromServer)
// 		close(stream.recvToServer)
// 	}()
// 	return stream

// }

// type Api struct {
// }

// func RegisterApiServer() *Api {
// 	return &Api{}
// }
// func ( Api) getbok(stream bookstore_pb.Bookstore_GetBookdetailsServer) error {
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {

// 			return nil

// 		}
// 		u := req.GetRequest()
// 		switch v := u.(type) {
// 		case *bookstore_pb.Input_Author:
// 			var book bookstore_pb.Book
// 			var a bookstore_pb.Streamresponse
// 			var flag bool
// 			row, err := DBService.GetDB().Table(request.TABLE_BOOK).Select("*").Where(request.TABLE_BOOK_AUTHOR+"= ?", v.Author).Rows()
// 			if err != nil {
// 				return err
// 				//stream.Send(streamerr(err)
// 			}
// 			for row.Next() {
// 				if err = row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
// 					return err
// 					//stream.Send(streamerr(err)

// 				}
// 			}
// 			if flag != true {
// 				a = bookstore_pb.Streamresponse{
// 					Resp: &bookstore_pb.Streamresponse_BookrResp{
// 						BookrResp: &book,
// 					},
// 				}
// 			}

// 			err = stream.Send(&a)
// 			handleError(err)

// 		case *bookstore_pb.Input_Bookname:
// 			var book bookstore_pb.Book
// 			row, err := s.DBService.GetDB().Table(request.TABLE_BOOK).Select("*").Where(request.TABLE_BOOK_BOOKNAME+"= ?", v.Bookname).Rows()
// 			if err != nil {
// 				return err
// 			}
// 			for row.Next() {
// 				if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
// 					return err
// 				}
// 			}
// 			a := bookstore_pb.Streamresponse{
// 				Resp: &bookstore_pb.Streamresponse_BookrResp{
// 					BookrResp: &book,
// 				},
// 			}

// 			stream.Send(&a)
// 		}
// 		var book bookstore_pb.Book
// 		row, err := s.DBService.GetDB().Table(request.TABLE_BOOK).Select("*").Where(request.TABLE_BOOK_BOOKID+"= ?", req.BookID).Rows()
// 		if err != nil {
// 			handleError(err)
// 			//stream.Send(streamerr(err)
// 		}
// 		for row.Next() {
// 			if err := row.Scan(&book.BookId, &book.Author, &book.Bookname, &book.Title); err != nil {
// 				handleError(err)
// 				//stream.Send(streamerr(err)

// 			}
// 		}
// 		a := bookstore_pb.Streamresponse{
// 			Resp: &bookstore_pb.Streamresponse_BookrResp{
// 				BookrResp: &book,
// 			},
// 		}
// 		stream.Send(&a)

// 	}
// }
