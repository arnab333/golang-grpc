package main

import (
	"log"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"golang-grpc/user"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := user.NewUserServiceClient(conn)

	userID := user.UserID{
		UserID: 2,
	}

	em := user.EmptyParams{}

	ids := user.UserIDList{UserIDList: []int64{3, 2}}

	response, err := c.GetUser(context.Background(), &userID)
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}
	log.Printf("GetUser Response from Server: %s \n", response)

	response2, err2 := c.GetUsers(context.Background(), &em)
	if err2 != nil {
		log.Fatalf("Error when calling GetUsers: %s", err2)
	}
	log.Printf("GetUsers Response from Server: %s \n", response2)

	response3, err3 := c.GetUsersWithIds(context.Background(), &ids)
	if err3 != nil {
		log.Fatalf("Error when calling GetUsersWithIds: %s", err3)
	}
	log.Printf("GetUsersWithIds Response from Server: %s \n", response3)
}
