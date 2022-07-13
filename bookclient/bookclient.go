package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/bookstore/bookstore_pb"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBook(c bookstore_pb.BookstoreClient) {
	bookrequest := bookstore_pb.BooksRequest{

		Books: &bookstore_pb.Book{
			Author:   "Yadav1",
			BookId:   "1011",
			Bookname: "Stranger1",
			Title:    "Thing31",
		},
	}
	resp, err := c.CreateBook(context.Background(), &bookrequest)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("book created with bookdid ", resp.BookID)
}

func GetBook(c bookstore_pb.BookstoreClient) {

	book := bookstore_pb.ID{
		BookID: "115",
	}
	resp, err := c.GetBook(context.Background(), &book)
	handleError(err)
	fmt.Println("Response -> ", resp)

}
func DeleteBook(c bookstore_pb.BookstoreClient) {
	book := bookstore_pb.ID{
		BookID: "1",
	}
	resp, err := c.DeleteBook(context.Background(), &book)
	handleError(err)
	fmt.Println("Response -> ", resp)

}
func Upatebook(c bookstore_pb.BookstoreClient) {

	book := bookstore_pb.UpdateBookRequest{

		Author:   "harry potter",
		Title:    "stone",
		Bookname: "magic",
	}
	resp, err := c.Upatebook(context.Background(), &book)

	handleError(err)
	fmt.Println(resp)

}
func PutContent(c bookstore_pb.BookstoreClient) {
	stream, err := c.PutContent(context.Background())
	handleError(err)

	neRequest := &bookstore_pb.PageInfoRequest{
		BookID:      "1",
		PageNumber:  "1",
		Pagesize:    "10",
		Pagecontent: "hello"}
	stream.Send(neRequest)
	neRequest = &bookstore_pb.PageInfoRequest{
		BookID:      "2",
		PageNumber:  "1",
		Pagesize:    "10",
		Pagecontent: "hello"}
	stream.Send(neRequest)
	neRequest = &bookstore_pb.PageInfoRequest{
		BookID:      "3",
		PageNumber:  "1",
		Pagesize:    "10",
		Pagecontent: "hello"}
	stream.Send(neRequest)
	neRequest = &bookstore_pb.PageInfoRequest{
		BookID:      "1",
		PageNumber:  "2",
		Pagesize:    "10",
		Pagecontent: "world"}
	err = stream.Send(neRequest)
	handleError(err)
	stream.CloseAndRecv()
}
func GetContent(c bookstore_pb.BookstoreClient) {
	RequestPage := bookstore_pb.Pagerequest{
		Pagesize:    6,
		Pgagenumber: 1,
	}
	res_Stream, err := c.GetContent(context.Background(), &RequestPage)
	handleError(err)
	for {
		content, err := res_Stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			handleError(err)
			return
		}
		fmt.Println(content)
	}
}

func GetBookdetails(c bookstore_pb.BookstoreClient) {
	var wg sync.WaitGroup
	stream, err := c.GetBookdetails(context.Background())
	fmt.Println("1. search author 2. book details with bookID")
	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')
	var input string
	fmt.Scanln(&input)

	switch text {
	case "1\n":
		authReq := &bookstore_pb.Input{
			Request: &bookstore_pb.Input_Search{
				Search: input},
		}
		err := stream.Send(authReq)
		handleError(err)

	case "2\n":
		bookReq := &bookstore_pb.Input{
			Request: &bookstore_pb.Input_Bookid{
				Bookid: input},
		}

		err = stream.Send(bookReq)
		handleError(err)

	}
	stream.CloseSend()
	wg.Add(1)
	go func() {
		for {
			resp, err1 := stream.Recv()
			handleError(err1)
			if err == io.EOF {
				wg.Done()
				break
			}
			if err != nil {
				handleError(err)
			}

			fmt.Println("Revieved Book info ", resp.GetBookrResp())

		}
	}()
	wg.Wait()

}
