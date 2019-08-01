package scheduler

import (
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
	ID string
	Image string
	SuccCount uint32
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
		SuccCount: 0,
	})

	fmt.Printf("Containers = %v\n", len(s.Containers))

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

	//删除保存的对应容器信息
	for index, c := range s.Containers {
		if c.ID == containerID {
			s.Containers = append(s.Containers[:index], s.Containers[index + 1:]...)
			break
		}
	}

	return nil
}