package main

import "C"
import (
	"log"
	"flag"
	"heterogeneous/arbitrator"
)

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "c", "", "Config file")
}

var a *arbitrator.Arbitrator

//export ArbitratorInit
func ArbitratorInit(cfgFile string) {
	configs := &arbitrator.ConfigInfo{}
	arbitrator.NewParser().Parse(cfgFile, configs)
	a = arbitrator.NewArbitrator(configs.Instance.Images,
		configs.Instance.Addresses,
		&arbitrator.ArbitratePolicy {
			Threshhold: configs.ExceptionRule.Threshold,
			MaxFailures: configs.ExceptionRule.MaxFailures,
			RestartImage: configs.ExceptionRule.RestartImage,
			IsUseTimePriority:configs.Instance.IsUseTimePriority,
		})
	a.Init()
}

//export ArbitratorAdd
func ArbitratorAdd(param1, param2 int32) {
	if r, err := a.Add(param1, param2); err != nil {
		log.Println("Call add RPC failed")
	} else {
		log.Println("result = ", r)
	}
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
		log.Println("Call add RPC failed")
	} else {
		log.Println("result = ", r)
	}
}

