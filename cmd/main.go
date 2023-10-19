package main

import (
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/application/grpc"
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/infra/db"
)

func main() {
	database := db.ConnectDB()
	grpc.StartServer(database, 50051)
}
