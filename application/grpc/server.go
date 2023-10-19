package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func StartServer(db *gorm.DB, port int) {
	server := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:%d", port)

	logError := "cannot start gRPC server"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(logError, err)
	}

	log.Printf("gRPC server has been start on port %d", port)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(logError, err)
	}
}
