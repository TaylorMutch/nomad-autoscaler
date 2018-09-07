package main

import (
	"fmt"
	"os"

	"github.com/TaylorMutch/nomad-autoscaler/command"
	"github.com/TaylorMutch/nomad-autoscaler/version"
	"github.com/mitchellh/cli"
)

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	return RunCustom(args)
}

func RunCustom(args []string) int {

	commands := map[string]cli.CommandFactory{
		"scale": func() (cli.Command, error) {
			return &command.ScaleCommand{}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version: version.GetVersion(),
			}, nil
		},
	}

	/*
		ui := &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		}
	*/

	cli := &cli.CLI{
		Name:     "nomad-autoscaler",
		Version:  version.GetVersion().FullVersionNumber(true),
		Args:     os.Args[1:],
		Commands: commands,
	}

	exitStatus, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitStatus
}
