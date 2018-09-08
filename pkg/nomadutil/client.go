package nomadutil

import (
	"github.com/golang/glog"
	nomad "github.com/hashicorp/nomad/api"
)

// Client is a wrapper around the Nomad Client api
type Client struct {
	API         *nomad.Client
	Jobs        *nomad.Jobs
	Allocations *nomad.Allocations
}

// NewClient creates a new nomad autoscaling client
func NewClient(config *nomad.Config) *Client {

	c, err := nomad.NewClient(config)
	if err != nil {
		glog.Fatalf("Could not initialize Nomad client: %v", err)
	}

	return &Client{API: c, Jobs: c.Jobs(), Allocations: c.Allocations()}
}
