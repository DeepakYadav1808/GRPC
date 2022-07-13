package database

import (
	"github.com/bookstore/bookstore_pb"
	"github.com/bookstore/model"
)

func CreateBook(req model.Books) error {

	err := Connectdb().Model(&model.Books{}).Create(&req)
	if err != nil {
		return err.Error
	}

	return nil
}
func Getbook(serReq string) (*model.Books, error) {
	var bookDetail model.Books
	row := Connectdb().Model(&model.Books{}).Where(model.TABLE_BOOK_BOOKID+"=?", serReq).Scan(&bookDetail)
	if row.Error != nil {
		return &model.Books{}, row.Error
	}
	return &bookDetail, nil
}
func DeleteBook(bookId string) error {

	tx := Connectdb().Begin()
	defer tx.Rollback()
	err := tx.Where(model.TABLE_PAGE_BOOKID+"=?", bookId).Delete(&model.Books{})
	if err != nil {
		return err.Error
	}
	tx.Commit()
	return nil
}
func Updatebook(updateBookReq *bookstore_pb.UpdateBookRequest) error {

	tx := Connectdb().Begin()
	defer tx.Rollback()
	newCol := make(map[string]interface{})
	newCol["title"] = updateBookReq.Title
	newCol["author"] = updateBookReq.Author
	newCol["book_name"] = updateBookReq.Bookname
	err := tx.Model(&model.Books{}).Updates(newCol)
	if err.Error != nil {
		return err.Error
	}
	tx.Commit()
	return nil
}
func Pagination(pageSize int64, pageNumber int64) ([]model.Books, error) {
	var book []model.Books
	tx := Connectdb().Begin()
	defer tx.Rollback()

	stmt := tx.Limit(pageSize).Offset((pageNumber - 1) * pageSize) // , order(asc)
	result := stmt.Model(&model.Books{}).Find(&book)
	if result.Error != nil {
		return []model.Books{}, result.Error
	}
	tx.Commit()
	return book, nil

}
func RetrieveData(column_name string, column_value string) (*bookstore_pb.Streamresponse, error) {
	var book model.Books
	tx := Connectdb().Begin()
	defer tx.Rollback()
	err := tx.Model(&model.Books{}).Select("*").Where(column_name+"= ?", column_value).Find(&book)
	if err.Error != nil {
		return &bookstore_pb.Streamresponse{}, err.Error
	}
	tx.Commit()
	bookDetail := bookstore_pb.Streamresponse{
		Resp: &bookstore_pb.Streamresponse_BookrResp{
			BookrResp: &bookstore_pb.Book{
				BookId:   book.BookId,
				Bookname: book.Bookname,
				Title:    book.Title,
				Author:   book.Author,
			},
		},
	}
	return &bookDetail, nil
}
