package main

import (
	"encoding/json"
	"io/ioutil"
	"flag"
	"fmt"
	"context"
	"log"
	"time"
	"io"
	"os"
	"google.golang.org/grpc"
	pb "heterogeneous/proto"

	"github.com/docker/docker/client"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/pkg/stdcopy"
)

/*
InstanceT represents the type of configuration information for Instance
*/
type InstanceT struct {
	Count uint32
	IsUseTimePriority bool
    Images []string
}

/*
ExceptionRuleT represents the type of configuration information for ExceptionRule
*/
type ExceptionRuleT struct {
	Threshold uint32
	MaxFailures uint32
	RestartImage string
}

/*
ConfigInfo represents the type of the whole configuration file
*/
type ConfigInfo struct {
	Instance *InstanceT
	ExceptionRule *ExceptionRuleT
}

/*
Servers addrsss
*/
var Servers = []string {"localhost:50051" , "localhost:50052"}

var (
	cfgFile string
)


func init() {
	flag.StringVar(&cfgFile, "c", "", "Config file")
}

func parseConfig(file string, v interface{}) {
	var data []byte
	if file != "" {
		var err error
		if data, err = ioutil.ReadFile(file); err != nil {
			log.Fatalf("Read config file: %s failed, reson: %v \n", file, err)
		}
	} else {
		log.Fatal("Config file path is empty")
	}

	if err := json.Unmarshal(data, v); err != nil {
		log.Fatalf("Parse config file: %s failed, reason: %v \n", file, err)
	}
}

/* 
Add encapsulate RPC call process
*/
func Add(address string, resultChan chan *pb.AddResponse) {
	fmt.Println("connect to ", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cann not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculateClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{A: 1, B: 2})
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}

	log.Printf("Result: %v, Server Type: %v, Server Id: %v", 
				r.Result, r.ServerType, r.ServerId)

	resultChan <- r //RPC执行的结果传递给channel，供main routine读取
}

func main() {
	flag.Parse()

	configs := &ConfigInfo{}
	parseConfig(cfgFile, configs)

	// c := make(chan *pb.AddResponse, 2)
	
	// //并发执行RPC调用
	// for _, addr := range Servers {
	// 	go func(addr string) {
	// 		Add(addr, c)
	// 	}(addr)
	// }

	// // 从channel中读取RPC的结果
	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg := <- c:
	// 		log.Printf("Receive result from channel: %v, %v,%v", 
	// 			msg.Result, msg.ServerType, msg.ServerId)
	// 	}
	// }
	ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }
    cli.NegotiateAPIVersion(ctx)

    reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "alpine",
        Cmd:   []string{"echo", "hello world"},
    }, nil, nil, "")
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        if err != nil {
            panic(err)
        }
    case <-statusCh:
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        panic(err)
    }

    stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}


