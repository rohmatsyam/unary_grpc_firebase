package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"unary_grpc/cmd/config"
	"unary_grpc/cmd/service"
	productPb "unary_grpc/pb/product"
)

const (
	port = "localhost:50051"
)

func main() {
	//	grpc server
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen : %v", err.Error())
	}

	client, err := config.ConnectFirestore()
	if err != nil {
		log.Fatal("can't connect firestore")
	}

	grpcServer := grpc.NewServer()
	productService := service.ProductService{Client: client}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server start on port : %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatal("Failed to serve : %v", err.Error())
	}
}
