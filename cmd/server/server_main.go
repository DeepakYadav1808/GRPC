package main 

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/bookstore/server_func"
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
	// ser := server_func.Server{}
	// ser.Conn = database.Initdb()
	server := server_func.NewRepository()

	bookstore_pb.RegisterBookstoreServer(s, server)
	handleError(err)
	go func() {
		if err := s.Serve(lis); err != nil {
			handleError(err)
		}
	}()

	fmt.Println("Postgress connected")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("\nclosing the server")
}
