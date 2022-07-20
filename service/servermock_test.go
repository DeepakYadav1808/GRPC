package service

import (
	"context"
	"testing"

	"github.com/bookstore/bookstore_pb"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test bookservice by id", func(t *testing.T) {
		BookStoreMock := NewMockBookstoreServer(mockCtrl)
		BookStoreMock.EXPECT().CreateBook(context.Background(), &bookstore_pb.BooksRequest{
			Books: &bookstore_pb.Book{
				BookId:   "1",
				Bookname: "new",
				Author:   "new",
				Title:    "best",
			},
		}).Return(&bookstore_pb.Id{
			Bookid: "1",
		}, nil).AnyTimes()

		id, errr := BookStoreMock.CreateBook(context.Background(), &bookstore_pb.BooksRequest{
			Books: &bookstore_pb.Book{
				BookId:   "1",
				Bookname: "new",
				Author:   "new",
				Title:    "best",
			},
		})
		assert.NoError(t, errr)
		assert.Equal(t, "1", id.Bookid)
	})

}

func TestDeleteBookService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test delete book by id", func(t *testing.T) {
		BookStoreMock := NewMockBookstoreServer(mockCtrl)
		BookStoreMock.EXPECT().DeleteBook(context.Background(), &bookstore_pb.Id{
			Bookid: "1",
		}).Return(&bookstore_pb.Response{
			Message: "deleted",
		}, nil).AnyTimes()

		_, errr := BookStoreMock.DeleteBook(context.Background(), &bookstore_pb.Id{
			Bookid: "1",
		},
		)
		assert.NoError(t, errr)

	})

}
