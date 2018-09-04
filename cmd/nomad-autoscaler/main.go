package main

import (
	"flag"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	flag.Parse()

	c := cli.NewCLI("nomad-autoscaler", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"scale": func() (cli.Command, error) {
			return &ScaleCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
