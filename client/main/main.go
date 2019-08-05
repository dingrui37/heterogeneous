package main

import (
	"log"
	"flag"
	"heterogeneous/arbitrator"
	// "google.golang.org/grpc"
	// pb "heterogeneous/proto"
	// "context"
	// "time"
)

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "c", "", "Config file")
}


func main() {
	//解析命令行参数
	flag.Parse()

	//解析配置文件
	configs := &arbitrator.ConfigInfo{}
	arbitrator.NewParser().Parse(cfgFile, configs)

	//构建裁决器
	a := arbitrator.NewArbitrator(configs.Instance.Images,
		configs.Instance.Addresses,
		&arbitrator.ArbitratePolicy {
			Threshhold: configs.ExceptionRule.Threshold,
			MaxFailures: configs.ExceptionRule.MaxFailures,
			RestartImage: configs.ExceptionRule.RestartImage,
			IsUseTimePriority:configs.Instance.IsUseTimePriority,
		})
	
    //裁决器初始化
	a.Init()

	//RPC调用
	if r, err := a.Add(1, 2); err != nil {
		log.Println("Add rpc failed")
	} else {
		log.Println("result = ", r)
	}
}


