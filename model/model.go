package model

const (
	TABLE_BOOK          = "books"
	TABLE_BOOK_BOOKID   = "book_id"
	TABLE_BOOK_BOOKNAME = "book_name"
	TABLE_BOOK_TITLE    = "title"
	TABLE_BOOK_AUTHOR   = "author"
)
const (
	TABLE_PAGE            = "page"
	TABLE_PAGE_BOOKID     = "book_id"
	TABLE_PAGE_CONTENT    = "page_content"
	TABLE_PAGE_PAGENUMBER = "page_number"
	TABLE_PAGE_PAGESIZE   = "page_size"
)

type Books struct {
	BookID   string `json:"book_id"`
	BookName string `json:"book_name"`
	Author   string `json:"author"`
	Title    string `json:"title"`
}

type Page struct {
	BookID      uint64 `json:"book_id"`
	PageNumber  string `json:"page_name"`
	PageSize    string `json:"page_size"`
	PageContent string `json:"page_content"`
}
