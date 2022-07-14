package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bookstore/bookstore_pb"
	"golang.org/x/sync/errgroup"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBook(c bookstore_pb.BookstoreClient) error {
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
		return err

	}
	fmt.Println("book created with bookdid ", resp.BookID)
	return nil
}

func GetBook(c bookstore_pb.BookstoreClient) error {

	book := bookstore_pb.ID{
		BookID: "115",
	}
	resp, err := c.GetBook(context.Background(), &book)
	if err != nil {
		return err
	}
	fmt.Println("Response -> ", resp)
	return nil
}
func DeleteBook(c bookstore_pb.BookstoreClient) error {
	book := bookstore_pb.ID{
		BookID: "1",
	}
	resp, err := c.DeleteBook(context.Background(), &book)
	if err != nil {
		return nil
	}
	fmt.Println("Response -> ", resp)
	return nil

}
func Upatebook(c bookstore_pb.BookstoreClient) error {

	book := bookstore_pb.UpdateBookRequest{

		Author:   "harry potter",
		Title:    "stone",
		Bookname: "magic",
	}
	resp, err := c.Upatebook(context.Background(), &book)

	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil

}
func GetContent(c bookstore_pb.BookstoreClient) error {
	RequestPage := bookstore_pb.Pagerequest{
		Pagesize:    6,
		Pgagenumber: 1,
	}
	res_Stream, err := c.GetContent(context.Background(), &RequestPage)
	handleError(err)
	for {
		content, err := res_Stream.Recv()
		if err != nil {
			return err
		}
		if err == io.EOF {
			break
		}

		fmt.Println(content)
	}
	return nil
}

func GetBookdetails(c bookstore_pb.BookstoreClient) error {

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
		if err != nil {
			return err
		}

	case "2\n":
		bookReq := &bookstore_pb.Input{
			Request: &bookstore_pb.Input_Bookid{
				Bookid: input},
		}

		err = stream.Send(bookReq)
		if err != nil {
			return err
		}

	}
	stream.CloseSend()

	errs, _ := errgroup.WithContext(context.Background())
	errs.Go(func() error {
		for {
			resp, err1 := stream.Recv()
			handleError(err1)
			if err != nil {
				return err
			}
			if err == io.EOF {

				break
			}

			if resp.GetBookrResp() == nil {
				fmt.Println("No book Found")
				return nil
			}
			fmt.Println("Revieved Book info ", resp.GetBookrResp())
		}

		return nil
	})
	return errs.Wait()

}
