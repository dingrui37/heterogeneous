package main

import (
	"encoding/json"
	"io/ioutil"
	"flag"
	"fmt"
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	pb "heterogeneous/proto"
	"heterogeneous/scheduler"
)

type InstanceT struct {
	Count uint32
	IsUseTimePriority bool
    Images []string
}

type ExceptionRuleT struct {
	Threshold uint32
	MaxFailures uint32
	RestartImage string
}

type ConfigInfo struct {
	Instance *InstanceT
	ExceptionRule *ExceptionRuleT
}

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


func main() {
	flag.Parse()

	configs := &ConfigInfo{}
	parseConfig(cfgFile, configs)

	s := &scheduler.Scheduler{
		Pool: &scheduler.ImagePool {
			WorkableImages:configs.Instance.Images,
		},
	}

	s.ContainerCreate("heterogeneous_go:v1.0.0", "50051", "tcp")
	s.ContainerCreate("heterogeneous_python:v1.0.0", "50052", "tcp")
	time.Sleep(10 * time.Second)
	s.ContanierRemove(s.Containers[0].ID)
	s.ContanierRemove(s.Containers[0].ID)
	fmt.Println("dddddd")
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
	
}


