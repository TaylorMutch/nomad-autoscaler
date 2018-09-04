package nomadutil

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	nomad "github.com/hashicorp/nomad/api"
)

var (
	nomadAddress = flag.String("address", os.Getenv("NOMAD_ADDR"), "The address for the Nomad HTTP API Endpoint. Default is `NOMAD_ADDR`.")
	nomadRegion  = flag.String("region", os.Getenv("NOMAD_REGION"), "The region for where we want to perform Nomad job scaling. Default is `NOMAD_REGION`.")
)

// Client is a wrapper around the Nomad Client api
type Client struct {
	API         *nomad.Client
	Jobs        *nomad.Jobs
	Allocations *nomad.Allocations
}

// NewClient creates a new nomad autoscaling client
func NewClient() *Client {

	config := nomad.Config{
		Address: *nomadAddress,
		Region:  *nomadRegion,
	}

	c, err := nomad.NewClient(&config)
	if err != nil {
		glog.Fatalf("Could not initialize Nomad client: %v", err)
	}

	return &Client{API: c, Jobs: c.Jobs(), Allocations: c.Allocations()}
}

func (c *Client) Stuff() {

	/*
		jobs, _, err := c.Jobs.List(&qo)
		if err != nil {
			glog.Errorf("Error retrieving jobs list: %v", err)
		}

		allocations, _, err := c.Allocations.List(&qo)
		if err != nil {
			glog.Errorf("Error retreiving allocations list: %v", err)
		}

		fmt.Println("Jobs")
		for _, x := range jobs {
			fmt.Println(x.ID)
		}

		fmt.Println("Allocations")
		for _, y := range allocations {
			fmt.Println(y.ID)
		}

		job1 := jobs[0]

		jobAllocations, _, err := c.Jobs.Allocations(job1.ID, true, &qo)
		if err != nil {
			glog.Errorf("Error retrieving allocations for job %v: %v", job1.ID, err)
		}

		fmt.Printf("Job allocations for %v:\n", job1.ID)
		for _, z := range jobAllocations {
			fmt.Println(z.ID)
		}

		// Get the detailed info for the first job
		job, _, err := c.Jobs.Info(job1.ID, &qo)
		if err != nil {
			glog.Errorf("Error retrieving info for job %v: %v", job1.ID, err)
		}

		println(&job.TaskGroups[0].Count)

		count := job.TaskGroups[0].Count

		fmt.Printf("%d", *count)

		//fmt.Printf("%v", job1.)
	*/

	// Scale up a job
	fmt.Println("Scaling thind up to 7")
	if err := c.Scale("thind", 7); err != nil {
		glog.Errorf("Error scaling up!")
	}

	// Sleep for 60 seconds
	time.Sleep(60 * time.Second)

	// Scale a job back down
	fmt.Println("Scaling thind down to 1")
	if err := c.Scale("thind", 1); err != nil {
		glog.Errorf("Error scaling down!")
	}
}
