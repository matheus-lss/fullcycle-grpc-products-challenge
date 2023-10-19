package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/application/grpc/pb"
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/application/usecase"
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/infra/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer(db *gorm.DB, port int) {
	server := grpc.NewServer()
	reflection.Register(server)

	productRepository := repository.ProductRepositoryDb{Db: db}
	productUseCase := usecase.ProductUseCase{ProductRepositoryInterface: productRepository}
	productGrpcService := NewProductGrpcService(productUseCase)
	pb.RegisterProductServiceServer(server, productGrpcService)

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
