package command

import (
	"fmt"

	"github.com/TaylorMutch/nomad-autoscaler/version"
)

// VersionCommand prints the version for the CLI
type VersionCommand struct {
	Version *version.VersionInfo
}

func (c *VersionCommand) Help() string {
	return ""
}

func (c *VersionCommand) Name() string { return "version" }

func (c *VersionCommand) Run(_ []string) int {
	fmt.Println(c.Version.FullVersionNumber(true))
	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Prints the Nomad version"
}
