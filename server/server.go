package main

import (
	"context"
	"fmt"
	pb "github.com/shrishti-27/grpc-user-data/user"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	port = ":50051"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) AddUserData(ctx context.Context, in *pb.User) (*pb.Message, error) {
	log.Printf("Recieved: %v", in.GetName())

	userData := fmt.Sprintf("Username: %v , CompanyName: %v , Phone Number: %v", in.GetName(), in.GetCompany(), fmt.Sprintf("%d", in.GetNumber()))
	file, err := os.Create("userData.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.WriteString(userData)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("saved data in file")

	return &pb.Message{Body: "Stored data in file"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
