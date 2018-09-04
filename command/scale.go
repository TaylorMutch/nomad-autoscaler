package command

import (
	"strconv"

	"github.com/TaylorMutch/nomad-autoscaler/pkg/nomadutil"
	"github.com/golang/glog"
	"github.com/mitchellh/cli"
)

type ScaleCommand struct {
	cli.Command
}

func (cmd *ScaleCommand) Help() string {
	return ""
}

func (cmd *ScaleCommand) Synopsis() string {
	return "Manually scale a deployment"
}

func (cmd *ScaleCommand) Run(args []string) int {

	// Parse argument inputs
	jobID := args[0]
	count, err := strconv.ParseInt(args[1], 10, 64)

	// Ensure job count is a number
	if err != nil {
		glog.Error("count must be a number")
		return 1
	}

	// Create a new nomad client
	c := nomadutil.NewClient()

	// Perform the job scaling
	if err := c.Scale(jobID, int(count)); err != nil {
		glog.Errorf("Failed to scale %v up to %v: %v", jobID, count, err)
		return 1
	}

	return 0
}
