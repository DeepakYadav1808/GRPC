package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/bookstore/bookstore_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBook(c bookstore_pb.BookstoreClient) {
	bookrequest := bookstore_pb.BooksRequest{

		Books: &bookstore_pb.Book{
			Author:   "Yadav",
			BookId:   "1",
			Bookname: "Stranger",
			Title:    "Thing3",
		},
	}
	resp, err := c.CreateBook(context.Background(), &bookrequest)
	handleError(err)
	fmt.Println("Response -> ", resp)
}

func GetBook(c bookstore_pb.BookstoreClient) {

	book := bookstore_pb.ID{
		BookID: "2",
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
		PreviousBookID: "3",
		Author:         "harry potter",
		Title:          "stone",
		Bookname:       "magic",
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
	stream.Send(neRequest)
	stream.CloseAndRecv()
}
func GetContent(c bookstore_pb.BookstoreClient) {
	RequestPage := bookstore_pb.Pagerequest{
		BooKID:     "1",
		Pagenumber: "1",
	}
	res_Stream, err := c.GetContent(context.Background(), &RequestPage)
	handleError(err)
	for {
		content, err := res_Stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Printf("Page Number=%v Content %v\n", RequestPage.Pagenumber, content.Content)
	}

}
func GetBookdetails(c bookstore_pb.BookstoreClient) {
	var wg sync.WaitGroup
	stream, err := c.GetBookdetails(context.Background())
	fmt.Println("1. search author 2. search bookname ")
	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')
	var input string
	fmt.Scanln(&input)

	switch text {
	case "1\n":
		authReq := &bookstore_pb.Input{
			Request: &bookstore_pb.Input_Author{
				Author: input},
		}
		err := stream.Send(authReq)

		handleError(err)

	case "2\n":
		bookReq := &bookstore_pb.Input{
			Request: &bookstore_pb.Input_Bookname{
				Bookname: input},
		}

		err = stream.Send(bookReq)
		handleError(err)

	}
	fmt.Println("Search with bookID")
	bokId := bookstore_pb.Input{
		BookID: "2",
	}
	stream.Send(&bokId)

	handleError(err)

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

			fmt.Println("Revieved Book info ", resp.GetBookrResp())

		}
	}()
	wg.Wait()

}
func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	handleError(err)

	defer conn.Close()
	c := bookstore_pb.NewBookstoreClient(conn)
	fmt.Println("Enter the one of the folliwing choices below:")
	menu := fmt.Sprintf(`       1.create an book entry:
        2.getbook from book id:
        3.remove book entry:
        4.update book entry:
        5.Put content using client side streaming:
        6.Get Content using server side streaming:
        7.Bidirectional -> get requested book details
	  `)
	fmt.Println(menu)

	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')

	switch text {
	case "1\n":
		fmt.Println("Creating Book entry")
		CreateBook(c)
	case "2\n":
		fmt.Println("Get Book entry")
		GetBook(c)
	case "3\n":
		fmt.Println("Delete Book entry")
		DeleteBook(c)
	case "4\n":
		fmt.Println("Update book entry")
		Upatebook(c)
	case "5\n":
		fmt.Println("client side streaming")
		PutContent(c)

	case "6\n":
		fmt.Println("Server side streaming")
		GetContent(c)
	case "7\n":
		fmt.Println("Getting requested book details")
		GetBookdetails(c)
	default:
		fmt.Println("Wrong input")

	}

}
