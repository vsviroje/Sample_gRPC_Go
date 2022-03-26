package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	userProto "gRPCHandler/user/protos"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := userProto.NewUserServiceClient(conn)

	response, err := c.GetUserByID(context.Background(), &userProto.RequestByID{Id: 2})
	if err != nil {
		log.Fatalf("Error when calling GetUserByID: %s", err)
	}
	log.Printf("Response from server: %s", response.User)

	resp, err := c.GetUsersByIDs(context.Background(), &userProto.RequestByIDs{Ids: []int32{3, 4, 5}})
	if err != nil {
		log.Fatalf("Error when calling GetUserByIDs: %s", err)
	}

	log.Printf("Response from server: %s", resp.Users)

}
