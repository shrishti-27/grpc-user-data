package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shrishti-27/grpc-user-data/user"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "shrishti"
	var number int32
	number = 999999999
	cname := "Technogise"
	r, err := c.AddUserData(ctx, &pb.User{Name: name, Number: number, Company: cname})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf(`Response: %v`, r.GetBody())
}
