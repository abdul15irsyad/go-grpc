package lib

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	address string
}

func NewGRPCServer(address string) *gRPCServer {
	return &gRPCServer{address: address}
}

func (gs *gRPCServer) Run() error {
	listen, err := net.Listen("tcp", gs.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Println("Starting gRPC server on", gs.address)

	return grpcServer.Serve(listen)
}
