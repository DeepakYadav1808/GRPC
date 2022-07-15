package database

import (
	"github.com/bookstore/bookstore_pb"
	"github.com/bookstore/model"
)

func CreateBook(req model.Books) error {

	err := Connectdb().Model(&model.Books{}).Create(&req)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
func Getbook(serReq string) (*model.Books, error) {
	var bookDetail model.Books

	result := Connectdb().Model(&model.Books{}).Where(model.TABLE_BOOK_BOOKID+"=?", serReq).Scan(&bookDetail)

	if result.Error != nil {
		return nil, result.Error
	}
	return &bookDetail, nil
}
func DeleteBook(bookId string) error {

	result := Connectdb().Where(model.TABLE_PAGE_BOOKID+"=?", bookId).Delete(&model.Books{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func Updatebook(updateBookReq *bookstore_pb.UpdateBookRequest) error {

	newCol := make(map[string]interface{})
	newCol["title"] = updateBookReq.Title
	newCol["author"] = updateBookReq.Author
	newCol["book_name"] = updateBookReq.Bookname
	result := Connectdb().Model(&model.Books{}).Updates(newCol)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func Pagination(pageSize int64, pageNumber int64) ([]model.Books, error) {
	var book []model.Books

	stmt := Connectdb().Limit(pageSize).Offset((pageNumber - 1) * pageSize) // , order(asc)
	result := stmt.Model(&model.Books{}).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil

}
func RetrieveData(column_name string, column_value string) (*bookstore_pb.Streamresponse, error) {
	var book model.Books

	result := Connectdb().Model(&model.Books{}).Select("*").Where(column_name+"= ?", column_value).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	bookDetail := bookstore_pb.Streamresponse{
		Resp: &bookstore_pb.Streamresponse_BookrResp{
			BookrResp: &bookstore_pb.Book{
				BookId:   book.BookId,
				Bookname: book.BookName,
				Title:    book.Title,
				Author:   book.Author,
			},
		},
	}
	return &bookDetail, nil
}
