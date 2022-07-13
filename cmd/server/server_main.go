package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	bookservice "github.com/bookstore/service"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/bookstore/bookstore_pb"
	"google.golang.org/grpc"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	bookstore_pb.RegisterBookstoreServer(s, &bookservice.Bookserver{})
	handleError(err)
	go func() {
		err := s.Serve(lis)
		handleError(err)

	}()

	fmt.Println("server started")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("\nclosing the server")
}
