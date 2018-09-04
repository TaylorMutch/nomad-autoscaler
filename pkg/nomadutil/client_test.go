package nomadutil

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Explicitly call main test run
	flag.Parse()
	m.Run()
}

func TestNewClient(t *testing.T) {

	c := NewClient()

	status := c.API.Status()

	peers, err := status.Peers()
	assert.NoError(t, err)
	fmt.Printf("Peers: %v\n", peers)
	leader, err := c.API.Status().Leader()
	assert.NoError(t, err)
	fmt.Printf("Leader: %v\n", leader)

}
