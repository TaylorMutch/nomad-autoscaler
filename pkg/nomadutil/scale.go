package nomadutil

import (
	"fmt"

	"github.com/pkg/errors"
)

// Scale updates the job count for a Nomad job.
func (c *Client) Scale(jobid string, count int) error {

	job, _, err := c.Jobs.Info(jobid, nil)
	if err != nil {
		return errors.Wrap(err, "Error retrieving job info")
	}

	taskGroup := job.TaskGroups[0]

	if count == *taskGroup.Count {
		fmt.Println("Count is the same, nothing to be done.")
		return nil
	}

	job.TaskGroups[0].Count = &count

	res, _, err := c.Jobs.Register(job, nil)
	if err != nil {
		return errors.Wrap(err, "Error registering updated job")
	}

	fmt.Printf("%v\n", res)

	return nil
}
