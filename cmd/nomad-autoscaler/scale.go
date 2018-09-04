package main

import (
	"fmt"
	"strconv"

	"github.com/TaylorMutch/nomad-autoscaler/pkg/nomadutil"
	"github.com/golang/glog"
	"github.com/mitchellh/cli"
)

type ScaleCommand struct {
	cli.Command
}

func (cmd *ScaleCommand) Help() string {
	return "Help me!"
}

func (cmd *ScaleCommand) Synopsis() string {
	return "Synopsis!"
}

func (cmd *ScaleCommand) Run(args []string) int {

	fmt.Println(args)

	jobID := args[0]
	count, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		glog.Error("count must be a number")
		return 1
	}

	c := nomadutil.NewClient()

	if err := c.Scale(jobID, int(count)); err != nil {
		glog.Errorf("Failed to scale %v up to %v: %v", jobID, count, err)
		return 1
	}

	return 0
}
