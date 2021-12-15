package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"golang-grpc/user"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := user.Server{}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
