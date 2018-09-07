package nomadutil

import (
	"fmt"

	"github.com/pkg/errors"
)

// Scale updates the job count for a Nomad job.
func (c *Client) Scale(jobid string, count int) error {

	// Retrieve the job info
	job, _, err := c.Jobs.Info(jobid, nil)
	if err != nil {
		return errors.Wrap(err, "Error retrieving job info")
	}

	// Check if the job count is the same
	if count == *job.TaskGroups[0].Count {
		fmt.Println("Count is the same, nothing to be done.")
		return nil
	}

	// Update the job's Count
	// TODO - how do split counts by task groups?
	job.TaskGroups[0].Count = &count

	res, _, err := c.Jobs.Register(job, nil)
	if err != nil {
		return errors.Wrap(err, "Error registering updated job")
	}

	// Report any warnings about the updated job to the user
	if res.Warnings != "" {
		fmt.Printf("Updated job registered with warnings: %v\n", res.Warnings)
	}

	return nil
}
