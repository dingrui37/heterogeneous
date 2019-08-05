package main

import (
	"flag"
	"os"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "heterogeneous/proto"
)

type server struct{}

var listenPort string

func init() {
	flag.StringVar(&listenPort, "p", "", "Server listen port")
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received: %v %v", in.A, in.B)
	result := in.A + in.B
	hostname, _ := os.Hostname()
	return &pb.AddResponse{
			Result: result,
			ServerType: pb.AddResponse_GOLANG,
			ServerId: &pb.AddResponse_ServerId {Id: hostname},
		}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":"+ listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterCalculateServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

