package main

import (
	"flag"
	"log"

	nomad "github.com/TaylorMutch/nomad-autoscaler/pkg/nomadutil"
)

func main() {
	flag.Parse()

	_ = nomad.NewClient()

	log.Print("Initialized!")

}
