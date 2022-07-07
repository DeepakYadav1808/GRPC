package server_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/bookstore/bookstore_pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// Test for CreateItem method when input has empty name
func TestCreateBook(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := bookstore_pb.NewBookstoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err = c.CreateBook(ctx, &bookstore_pb.BooksRequest{
		Books: &bookstore_pb.Book{
			Author:   "Yadav",
			BookId:   "22",
			Bookname: "Stranger1",
			Title:    "Thing3",
		},
	})

	require.NoError(t, err)
	book := bookstore_pb.ID{
		BookID: "22",
	}
	resp, err1 := c.GetBook(context.Background(), &book)

	// ID's should be equal as insertion succeeds
	assert.Equal(t, "22", resp.BookId, "should be equal")
	require.NoError(t, err1)

	bb := bookstore_pb.ID{
		BookID: "22",
	}
	_, err2 := c.DeleteBook(context.Background(), &bb)
	require.NoError(t, err2)
}
func TestDelete(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := bookstore_pb.NewBookstoreClient(conn)
	book := bookstore_pb.ID{
		BookID: "11",
	}
	_, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err1 := c.DeleteBook(context.Background(), &book)
	require.NoError(t, err1)

}
func TestGet(t *testing.T) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := bookstore_pb.NewBookstoreClient(conn)
	book := bookstore_pb.ID{
		BookID: "11",
	}
	_, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err1 := c.GetBook(context.Background(), &book)
	require.NoError(t, err1)
	defer cancel()
}
