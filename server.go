package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"golang-grpc/user"
)

func main() {
	listener, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("Failed to listen on port 4000: %v", err)
	}

	s := user.Server{}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, &s)

	log.Println("Server is running!")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over port 4000: %v", err)
	}

}
