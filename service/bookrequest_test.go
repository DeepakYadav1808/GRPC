package service

import (
	"context"
	"testing"

	"github.com/bookstore/bookstore_pb"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateBook(t *testing.T) {
	bookservice := Bookserver{}

	id := uuid.NewString()
	_, err := bookservice.CreateBook(context.TODO(), &bookstore_pb.BooksRequest{
		Books: &bookstore_pb.Book{
			Author:   "adads",
			Title:    "magic1",
			BookId:   id,
			Bookname: "real",
		},
	},
	)
	require.NoError(t, err)

	resp, err2 := bookservice.GetBook(context.TODO(), &bookstore_pb.Id{
		Bookid: id,
	})
	require.NoError(t, err2)
	assert.Equal(t, id, resp.BookId, "should be equal")
	assert.Equal(t, "adads", resp.Author, "should be equal")
	assert.Equal(t, "magic1", resp.Title, "should be equal")
	assert.Equal(t, "real", resp.Bookname, "should be equal")
	require.NoError(t, err2)

	_, err2 = bookservice.DeleteBook(context.TODO(), &bookstore_pb.Id{
		Bookid : id,
	})
	require.NoError(t, err2)
}
func TestUpdatebook(t *testing.T) {
	bookservice := Bookserver{}

	id := uuid.NewString()
	_, err := bookservice.CreateBook(context.TODO(), &bookstore_pb.BooksRequest{
		Books: &bookstore_pb.Book{
			Author:   "adads",
			Title:    "magic1",
			BookId:   id,
			Bookname: "real",
		},
	},
	)
	require.NoError(t, err)

	_, err = bookservice.Upatebook(context.TODO(), &bookstore_pb.UpdateBookRequest{
		Title:    "new",
		Bookname: "newmagic",
		Author:   "bestauth",
	})
	require.NoError(t,err)

	resp, err2 := bookservice.GetBook(context.TODO(), &bookstore_pb.Id{
		Bookid: id,
	})
	require.NoError(t, err2)
	assert.Equal(t, "new", resp.Title, "should be equal")
	assert.Equal(t, "newmagic", resp.Bookname, "should be equal")
	assert.Equal(t, "bestauth", resp.Author, "should be equal")
	
	_, err2 = bookservice.DeleteBook(context.TODO(), &bookstore_pb.Id{
		Bookid: id,
	})
	require.NoError(t, err2)
}
