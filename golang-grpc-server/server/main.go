package main

import (
	"context"
	"log"
	"net"

	pb "github.com/mahendrabagul/devsecops-meetup/golang-grpc-server/protos/employee"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedEmployeeServiceServer
}

func (s *server) GetDetails(ctx context.Context, in *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	log.Printf("Received: %v", in.GetId())
	ed := pb.EmployeeDetails{
		Id:        in.GetId(),
		Email:     "bagulm123@gmail.com",
		FirstName: "Mahendra",
		LastName:  "Bagul",
	}
	return &pb.EmployeeResponse{Message: &ed}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
