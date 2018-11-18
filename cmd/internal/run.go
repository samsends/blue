package internal

import (
	"context"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/davecgh/go-spew/spew"
)

// Run ...
func Run() {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// image := "78a84c946f9b"
	// resp, err := cli.ContainerCreate(ctx, &container.Config{
	// 	Image: image,
	// }, nil, nil, "test")

	// if err != nil {
	// 	panic(err)
	// }

	err = cli.ContainerStart(ctx, "test", types.ContainerStartOptions{})

	if err != nil {
		panic(err)
	}

	execConfig := types.ExecConfig{
		Cmd: []string{""},
	}
	resp, err := cli.ContainerExecCreate(ctx, "test", execConfig)
	if err != nil {
		panic(err)
	}
	spew.Dump(resp)
}
