package arbitrator

import (
	"sync"
	"strings"
	"math/rand"
	"time"
	"fmt"
	"log"
	"context"
	"heterogeneous/scheduler"
	"google.golang.org/grpc"
	pb "heterogeneous/proto"
)

type ArbitratePolicy struct {
	Threshhold uint32
	MaxFailures uint32
	RestartImage string
	IsUseTimePriority bool
}

//裁决器定义
type Arbitrator struct {
	Scheduler *scheduler.Scheduler  //调度器
	Servers []string                //服务端信息
	Policy	*ArbitratePolicy         //裁决策略信息
}

//通用裁决数据结构
type Element struct {
	Value interface{}  //map健值，不可以是slice、map以及func类型
	ServerID string
}

//结果统计
type Statistics struct {
	count uint32
	indexs []int
}

var once sync.Once

//裁决器裁决部分的算法是通用的，针对所有RPC
func (a *Arbitrator) arbitrate(elements []*Element) (int, error) {
	keys := make(map[interface{}]*Statistics)

	//统计结果分布
	for i, v := range elements {
		if _, ok := keys[v.Value]; ok {
			keys[v.Value].count++
			keys[v.Value].indexs = append(keys[v.Value].indexs, i)
		} else {
			s := &Statistics {
				count: 1,
				indexs: []int{i},
			}
			keys[v.Value] = s
		}
	}

	fmt.Printf("-----elements------\n")
	for i,v := range elements {
		fmt.Printf("i = %v, v = %v/%v\n", i, v.Value, v.ServerID )
	}

	fmt.Printf("-----keys------\n")
	for i,v := range keys {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}

	//如果结果全部一致，则成功
	//如果结果不一致
	//  找出结果最多的出现次数，计算出占比
	//    a. 如果达到阈值，则成功
	//    b. 如果达不到阈值，则失败
	if len(keys) == 1 {
		for _, c := range a.Scheduler.Containers {
			c.SuccCount++                
			c.ContinuousFailureCount = 0
		}
		return 0, nil
	} else {
		var maxCount uint32
		var key interface{}
		for k, v := range keys {
			if v.count > maxCount {
				maxCount = v.count
				key = k
			}
		}

		fmt.Printf("maxCount = %v, key = %v, keys = %v\n", maxCount, key, keys)

		//时间因素的考虑：运行时间越长，即RPC成功的次数越多的，越可靠，权重越高，反之新版本稳定性可能越差？
		//失败次数分为连续失败次数、累计失败次数，使用连续失败次数具备统计效应，裁决更准确
		if float32(maxCount) / float32(len(elements)) >= float32(a.Policy.Threshhold) {
			//容器状态更新
			for _, c := range a.Scheduler.Containers {
				isExisted := false
				for _, v := range keys[key].indexs {
					if  strings.Contains(c.ID, elements[v].ServerID) {   //长短ID
						isExisted = true
						break
					}
				}

				if isExisted {
					c.SuccCount++                  //累计成功次数
					c.ContinuousFailureCount = 0   //成功一次，连续失败次数清0
					log.Printf("c.ID = %v success = %v/%v\n", c.ID, c.SuccCount,c.ContinuousFailureCount)
				} else {
					c.TotalFailureCount++          //累计失败次数
					c.ContinuousFailureCount++	   //连续失败次数++	
					log.Printf("c.ID = %v failure = %v/%v\n", c.ID, c.TotalFailureCount,c.ContinuousFailureCount)	
				}
			}

			//判断异常容器
			for _, c := range a.Scheduler.Containers {
				if c.ContinuousFailureCount >= a.Policy.MaxFailures {
					if err := a.Scheduler.ContanierRemove(c.ID); err != nil {
						log.Printf("Remove container failed, reason:%s", err)
					}

					log.Printf(`Container:%v be removed, Image:%v, 
								ServiceAddress:%v, SuccCount:%v, 
								ContinuousFailureCount:%v, TotalFailureCount:%v`, 
								c.ID, c.Image, 
								c.ServiceAddress, c.SuccCount, 
								c.ContinuousFailureCount, c.TotalFailureCount)
				}
			}
			return keys[key].indexs[0], nil
		} else {
			return 0, fmt.Errorf("Results cannot be arbitrated")
		}
	}
}

func (a *Arbitrator) Add(param1, param2 int32) (int32, error) {
	//结果从该channel输出
	resultChan := make(chan *pb.AddResponse)
			
	//并发执行RPC调用
	for _, addr := range a.Servers {
		go func(address string) {
			log.Printf("Connect to %v", address)

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Printf("Cann not connect: %v", err)
				return
			}
			defer conn.Close()

			c := pb.NewCalculateClient(conn)
			
			//每次RPC超时时间设置为5s
			ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second) 
			defer cancel()
			
			//配置WaitForReady，如果失败grpc会多次retry直至超时
			r, err := c.Add(ctx, &pb.AddRequest{A: param1, B: param2}, grpc.WaitForReady(true)) 
			if err != nil {
				log.Printf("Could not execut add RPC: %v", err)
				return
			}
		
			log.Printf("Result: %v, Server Type: %v, Server Id: %v", 
						r.Result, r.ServerType, r.ServerId)
		
			resultChan <- r //RPC执行的结果传递给channel，供其他goroutine读取
		}(addr)
	}

	elements := make([]*Element, 0)
	// 从channel中读取RPC的结果
	isTimeOut := false
	for {
		//超时或者已经读取完所有结果则直接结束
		if isTimeOut || len(elements) == len(a.Servers) {
			break
		}

		select {
		case msg := <- resultChan:
			log.Printf("Receive result from channel: %v, %v,%v", 
				msg.Result, msg.ServerType, msg.ServerId)
			e := &Element{
				Value: msg.Result,
				ServerID:msg.ServerId.Id,
			}
			elements = append(elements, e)	
		case <- time.After(5 * time.Second): //防止读超时
			isTimeOut = true
		}
	}
   
	//有结果丢失或者异常
	if len(elements) < len(a.Servers) {
		return 0, fmt.Errorf("Result miss or exception")
	}

	index, error := a.arbitrate(elements)
	if error != nil {
		return 0, error
	}

	return elements[index].Value.(int32), nil
}

func (a *Arbitrator) Init() {
	for _, s := range a.Servers {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := r.Intn(len(a.Scheduler.Pool.WorkableImages))
		image := a.Scheduler.Pool.WorkableImages[i]
		port := strings.Split(s, ":")[1]
		if err := a.Scheduler.ContainerCreate(image, port, "tcp"); err != nil {
			log.Fatalf("Cannot create container, image: %v, port: %v, reason:%v\n", 
				image, port, err)
		}
	}
}

func NewArbitrator(images []string, servers []string, Policy *ArbitratePolicy) *Arbitrator {
	var a *Arbitrator
	once.Do(func(){ //单例模式
		a = &Arbitrator{
			Scheduler:&scheduler.Scheduler {
				Pool:&scheduler.ImagePool{
					WorkableImages:images,
					ExceptionImages:make([]string, 0),
				},
				Containers:make([]*scheduler.Container, 0),
			},
			Servers:servers,
			Policy:Policy,
		}
	})

	log.Printf("New sigleton Arbitrator success")
	return a
}

