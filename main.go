package main

import (
	"flag"
	"log"
	"os"

	"github.com/TaylorMutch/nomad-autoscaler/command"
	"github.com/TaylorMutch/nomad-autoscaler/version"
	"github.com/mitchellh/cli"
)

func main() {
	flag.Parse()

	c := cli.NewCLI("nomad-autoscaler", version.GetVersion().FullVersionNumber(true))
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"scale": func() (cli.Command, error) {
			return &command.ScaleCommand{}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version: version.GetVersion(),
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
