package main

import (
	"context"
	"fmt"
	"log"
	"net"

	gHelper "Sample_gRPC/helper"
	"Sample_gRPC/models"
	svc "Sample_gRPC/services"
	userProto "gRPCHandler/user/protos"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Server Started!")
	defer fmt.Println("Server ended!")

	gHelper.PrepareUserRec()

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := UserServer{}

	grpcServer := grpc.NewServer()

	userProto.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

type UserServer struct {
	userProto.UnimplementedUserServiceServer
}

func (s *UserServer) GetUserByID(ctx context.Context, in *userProto.RequestByID) (*userProto.RespUser, error) {
	log.Printf("Received: %v", in.GetId())

	user, err := svc.GetUserByID(int(in.GetId()))
	if err != nil {
		return nil, err
	}

	resp := userProto.RespUser{}
	resp.User = ModelToProto(user)

	return &resp, nil
}

func (s *UserServer) GetUsersByIDs(ctx context.Context, in *userProto.RequestByIDs) (*userProto.RespUsers, error) {
	log.Printf("Received: %v", in.GetIds())

	userList, err := svc.GetUserByIDs(in.GetIds())
	if err != nil {
		return nil, err
	}
	respList := []*userProto.User{}
	for i := range userList {
		respList = append(respList, ModelToProto(&userList[i]))
	}

	resp := userProto.RespUsers{}
	resp.Users = respList

	return &resp, nil
}

func ModelToProto(in *models.User) *userProto.User {
	if in == nil {
		return nil
	}

	out := userProto.User{}
	out.Id = int32(in.ID)
	out.City = in.City
	out.Fname = in.FName
	out.Phone = in.Phone
	out.Height = in.Height
	out.Married = in.Married

	return &out
}
