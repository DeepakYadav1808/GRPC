package server_func_test

import (
	"context"
	"testing"

	"github.com/bookstore/bookstore_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bookstore/server_func"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

// Test for CreateItem method when input has empty name

// Ge

func TestCreateBook(t *testing.T) {

	// newServer := server.Server{}
	// newServer.Conn = database.Initdb()
	Server := server_func.NewRepository()
	_, err := Server.CreateBook(context.TODO(), &bookstore_pb.BooksRequest{
		Books: &bookstore_pb.Book{
			Author:   "Yadav",
			BookId:   "25",
			Bookname: "Stranger1",
			Title:    "Thing3",
		},
	})

	require.NoError(t, err)
	book := bookstore_pb.ID{
		BookID: "22",
	}
	resp, err1 := Server.GetBook(context.TODO(), &book)

	// ID's should be equal as insertion succeeds
	assert.Equal(t, "22", resp.BookId, "should be equal")
	require.NoError(t, err1)

	bb := bookstore_pb.ID{
		BookID: "22",
	}
	_, err2 := Server.DeleteBook(context.TODO(), &bb)
	require.NoError(t, err2)
}

func TestDelete(t *testing.T) {

	newserver := server_func.NewRepository()

	_, err := newserver.DeleteBook(context.Background(), &bookstore_pb.ID{
		BookID: "23",
	})
	if err != nil {
		t.Error("error while creating deleting the record")
	}
	require.NoError(t, err)

}
func TestGet(t *testing.T) {

	newserver := server_func.NewRepository()
	resp, err1 := newserver.GetBook(context.Background(), &bookstore_pb.ID{
		BookID: "1",
	})
	require.NoError(t, err1)
	assert.Equal(t, "1", resp.BookId, "should be equal")

}

func TestGetDetails(t *testing.T) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()
	c := bookstore_pb.NewBookstoreClient(conn)

	stream, err1 := c.GetBookdetails(context.Background())
	require.NoError(t, err1)
	b := bookstore_pb.Input{
		Request: &bookstore_pb.Input_Search{

			Search: "neauth",
		},
	}

	_ = server_func.NewRepository()

	err3 := stream.Send(&b)

	require.NoError(t, err3)
	resp, err2 := stream.Recv()
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
