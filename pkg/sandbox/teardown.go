package sandbox

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/enescakir/emoji"
	sandboxCmdConfig "github.com/flyteorg/flytectl/cmd/config/subcommand/sandbox"
	"github.com/flyteorg/flytectl/pkg/configutil"
	"github.com/flyteorg/flytectl/pkg/docker"
	"github.com/flyteorg/flytectl/pkg/k8s"
)

func Teardown(ctx context.Context, cli docker.Docker, teardownFlags *sandboxCmdConfig.TeardownFlags) error {
	c, err := docker.GetSandbox(ctx, cli)
	if err != nil {
		return err
	}
	if c != nil {
		if err := cli.ContainerRemove(context.Background(), c.ID, types.ContainerRemoveOptions{
			Force: true,
		}); err != nil {
			return err
		}
	}
	if err := configutil.ConfigCleanup(); err != nil {
		fmt.Printf("Config cleanup failed. Which Failed due to %v\n", err)
	}
	if err := removeSandboxKubeContext(); err != nil {
		fmt.Printf("Kubecontext cleanup failed. Which Failed due to %v\n", err)
	}
	// Teardown volume if option is specified
	if teardownFlags.Volume {
		if err := cli.VolumeRemove(ctx, docker.FlyteSandboxVolumeName, true); err != nil {
			fmt.Printf("Volume cleanup failed. Which Failed due to %v\n", err)
		}
	}

	fmt.Printf("%v %v Sandbox cluster is removed successfully.\n", emoji.Broom, emoji.Broom)
	return nil
}

func removeSandboxKubeContext() error {
	k8sCtxMgr := k8s.NewK8sContextManager()
	return k8sCtxMgr.RemoveContext(sandboxContextName)
}
