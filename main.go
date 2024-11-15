package main

import (
	"learn-grpc/lib"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := lib.NewGRPCServer(":8010")
	grpcServer.Run()
}

func TestClient(address *string) {
	conn, err := grpc.NewClient(":" + *address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
