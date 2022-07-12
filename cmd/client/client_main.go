package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/bookstore/bookstore_pb"
	client "github.com/bookstore/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
		client.CreateBook(c)
	case "2\n":
		fmt.Println("Get Book entry")
		client.GetBook(c)
	case "3\n":
		fmt.Println("Delete Book entry")
		client.DeleteBook(c)
	case "4\n":
		fmt.Println("Update book entry")
		client.Upatebook(c)
	case "5\n":
		fmt.Println("client side streaming")
		client.PutContent(c)

	case "6\n":
		fmt.Println("Server side streaming")
		client.GetContent(c)
	case "7\n":
		fmt.Println("Getting requested book details")
		client.GetBookdetails(c)
	default:
		fmt.Println("Wrong input")

	}

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
