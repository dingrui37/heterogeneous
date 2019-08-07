package scheduler

import (
	"log"
	"fmt"
	"context"
	"github.com/docker/docker/client"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

type ImagePool struct {
	WorkableImages []string
	ExceptionImages []string
}

type Container struct {
	ID string                       //容器ID
	Image string                    //镜像名称
	ServiceAddress string 		    //服务地址 IP：Port
	SuccCount uint32                //成功的次数
	ContinuousFailureCount uint32   //连续失败的次数
	TotalFailureCount uint32        //累计失败的次数
}

type Scheduler struct {
	Pool *ImagePool
	Containers []*Container
}

func (s *Scheduler) ContainerCreate(image, port, protocol string) error {
	ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }
    cli.NegotiateAPIVersion(ctx)

	//校验是否在WorkableImages中
	var isFind bool 
	for _, i := range s.Pool.WorkableImages {
		if i == image {
			isFind = true
			break
		}
	}

	if !isFind {
		return fmt.Errorf("Cannot find image: %s in workable images", image)
	}

	portProto := nat.Port(port + "/" + protocol)
	resp, err := cli.ContainerCreate(ctx, 
		&container.Config{
			Image: image,
			ExposedPorts: nat.PortSet{
				portProto : struct{}{},
			},
			Cmd:[]string{"-p", port},
		}, 

		&container.HostConfig{
			PortBindings:map[nat.Port][]nat.PortBinding{
				portProto:[]nat.PortBinding{
					{
						HostIP:"0.0.0.0", //简单处理，绑定本地所有地址
						HostPort:port,
					},
				},
			},
		}, nil, "")

    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
	}
	
	s.Containers = append(s.Containers, &Container{
		ID: resp.ID,
		Image: image,
		ServiceAddress: "0.0.0.0" + port,
	})

	log.Printf("Container started, id: %v, image: %v, ServicePort: %v", resp.ID, image, port)
	return nil
} 

func (s *Scheduler) ContanierRemove(containerID string) error {
	ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }
    cli.NegotiateAPIVersion(ctx)
	
	err = cli.ContainerRemove(ctx, containerID, 
			types.ContainerRemoveOptions{Force: true}) //强制停止
	if err != nil {
		return err
	}

	//删除保存的对应容器、镜像信息
	var containerIndex int
	var imageIndex int
	for index, c := range s.Containers {
		if c.ID == containerID {
			containerIndex = index
			break
		}
	}

	for index, i := range s.Pool.WorkableImages {
		if s.Containers[containerIndex].Image == i {
			imageIndex = index
			break
		}
	}

	s.Containers = append(s.Containers[:containerIndex], s.Containers[containerIndex + 1:]...)
	s.Pool.ExceptionImages = append(s.Pool.ExceptionImages, s.Pool.WorkableImages[imageIndex])
	s.Pool.WorkableImages = append(s.Pool.WorkableImages[:imageIndex], s.Pool.WorkableImages[imageIndex + 1:]...)
	log.Printf("Container removed, id: %v, workable images:%v, exception images:%v", 
				containerID, s.Pool.WorkableImages, s.Pool.ExceptionImages)
	return nil
}

func NewScheduler(images []string) *Scheduler {
	return &Scheduler {
		Pool:&ImagePool{
			WorkableImages:images,
			ExceptionImages:make([]string, 0),
		},
		Containers:make([]*Container, 0),
	}
}