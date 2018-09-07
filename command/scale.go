package command

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/TaylorMutch/nomad-autoscaler/pkg/nomadutil"
	"github.com/golang/glog"
	nomad "github.com/hashicorp/nomad/api"
	"github.com/mitchellh/cli"
)

// ScaleCommand is a CLI command for scaling a Nomad job
type ScaleCommand struct {
	cli.Command

	flagAddress string
	flagRegion  string
}

// Name of the Scale command
func (cmd *ScaleCommand) Name() string { return "scale" }

// Help text for the Scale command
func (cmd *ScaleCommand) Help() string {
	return ""
}

// Synopsis for the Scale command
func (cmd *ScaleCommand) Synopsis() string {
	return "Manually scale a deployment"
}

// Run executes the Scale command
func (cmd *ScaleCommand) Run(args []string) int {

	// Create a FlagSet for flags we want to parse
	// TODO - we will want to make these flags common for all commands
	flags := flag.NewFlagSet("scale", flag.ContinueOnError)
	flags.StringVar(&cmd.flagAddress, "address", "", "")
	flags.StringVar(&cmd.flagRegion, "region", "", "")

	var err error

	if err = flags.Parse(args); err != nil {
		glog.Errorf("Error parsing commands: %v", err)
		return 1
	}

	args = flags.Args()

	if l := len(args); l != 2 {
		glog.Errorf("Not enough arguments; Requires a job ID and a non-negative integer.")
	}

	// Parse argument inputs
	jobID := args[0]
	var count int64
	if count, err = strconv.ParseInt(args[1], 10, 64); err != nil {
		glog.Error("Number to scale must be a number!")
		return 1
	}

	if count < 0 {
		glog.Error("Number to scale must be non-negative!")
		return 1
	}

	// Create a new nomad client and scale the job
	c := nomadutil.NewClient(&nomad.Config{
		Address: cmd.flagAddress,
		Region:  cmd.flagRegion,
	})

	if err := c.Scale(jobID, int(count)); err != nil {
		glog.Errorf("Failed to scale %v up to %v: %v", jobID, count, err)
		return 1
	}

	fmt.Printf("Successfully scaled %v to %v\n", jobID, count)

	return 0
}
