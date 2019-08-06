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
	log.Printf("Parse config file:%s success, configs:%v", cfgFile, configs)
	a.Init()
}

//export ArbitratorAdd
func ArbitratorAdd(param1, param2 int32) {
	if r, err := a.Add(param1, param2); err != nil {
		log.Println("Call add RPC failed")
	} else {
		log.Println("Result = ", r)
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
		log.Println("Result = ", r)
	}
}



//Usage:
//
// 1. 编译go文件生成动态库(workdir main)
//    go build -o cgo/libarbitrator.so -buildmode=c-shared main.go
// 2. 编写C代码，include编译生成的头文件，指定裁决配置文件，然后裁决器初始化、执行RPC调用
// 3. 编译C程序生成可执行文件(workdir main/cgo)
//    gcc main.c -o main ./libarbitrator.so
// 4. 执行C程序编译后的二进制文件
//     a. 杀死所有相关容器（可以手动操作，下面命令会杀死 全部 容器）
//        docker rm -f $(docker ps -aq)
//     b. ./main
//        dingrui@ubuntu:~/projects/heterogeneous/client/main/cgo$ ./main 
//        2019/08/06 10:18:39 New sigleton Arbitrator success
//        2019/08/06 10:18:39 Parse config file:../config.json success, configs:&{0xc000326600 0xc000291540}
//        2019/08/06 10:18:40 Container started, id: ..., image: heterogeneous_java:v1.0.0, ServicePort: 50051
//        2019/08/06 10:18:40 Container started, id: ..., image: heterogeneous_go:v1.0.0, ServicePort: 50052
//        2019/08/06 10:18:41 Container started, id: ..., image: heterogeneous_python:v1.0.0, ServicePort: 50053
//        2019/08/06 10:18:42 Container started, id: ..., image: heterogeneous_cpp:v1.0.0, ServicePort: 50054
//        2019/08/06 10:18:42 Connect to 127.0.0.1:50052
//        2019/08/06 10:18:42 Connect to 127.0.0.1:50054
//        2019/08/06 10:18:42 Connect to 127.0.0.1:50053
//        2019/08/06 10:18:42 Connect to 127.0.0.1:50051
//        2019/08/06 10:18:42 Result: 3, Server Type: PYTHON, Server Id: id:"ec6ac88c0d24" 
//        2019/08/06 10:18:42 Result: 3, Server Type: PYTHON, Server Id: id:"fc3894ec96a1" 
//        2019/08/06 10:18:42 Result: 3, Server Type: GOLANG, Server Id: id:"ea0b19e2afa1" 
//        2019/08/06 10:18:42 Result: 3, Server Type: GOLANG, Server Id: id:"657056cb84b6" 
//        2019/08/06 10:18:42 Receive result from channel: 3, PYTHON,id:"ec6ac88c0d24" 
//        2019/08/06 10:18:42 Receive result from channel: 3, PYTHON,id:"fc3894ec96a1" 
//        2019/08/06 10:18:42 Receive result from channel: 3, GOLANG,id:"ea0b19e2afa1" 
//        2019/08/06 10:18:42 Receive result from channel: 3, GOLANG,id:"657056cb84b6" 
//        2019/08/06 10:18:47 Result =  3  //fixed bug, not need to wait timeout
//  5. 添加其他RPC

