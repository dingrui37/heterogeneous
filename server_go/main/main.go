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

	//伪造边界条件，当参数A是2的倍数，参数B是3的倍数时，返回错误结果
    //build有bug的镜像时，放开下面的代码
	var result int32
	if in.A % 2 == 0 && in.B % 3 == 0 {
		result = 0
	} else {
		result = in.A + in.B
	}

	//正常镜像代码
	//result := in.A + in.B
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

//1. 编译生成二进制文件
//     go build -o calculate main/main.go
//2. 制作docker镜像
//    docker build -t heterogeneous_cpp:v1.0.0 .