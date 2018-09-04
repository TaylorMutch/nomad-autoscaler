package main

import (
	"github.com/TaylorMutch/nomad-autoscaler/pkg/nomadutil"
	"github.com/golang/glog"
	"github.com/mitchellh/cli"
)

type ScaleCommand struct {
	cli.Command
}

func (cmd *ScaleCommand) Run(args []string) int {

	jobID := args[0]
	count := args[1]

	c := nomadutil.NewClient()

	if err := c.Scale(jobID, count); err != nil {
		glog.Errorf("Failed to scale %v up to %v: %v", jobID, count, err)
		return 1
	}

	return 0
}
