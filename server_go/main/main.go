package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "heterogeneous/proto"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received: %v %v", in.A, in.B)
	result := in.A + in.B
	return &pb.AddResponse{
			Result: result,
			ServerType: pb.AddResponse_GOLANG,
			ServerId: &pb.AddResponse_ServerId {Id: 1},
		}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterCalculateServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

