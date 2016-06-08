package agent

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/docker/swarmkit/agent/exec"
	"github.com/docker/swarmkit/api"
	"github.com/docker/swarmkit/picker"
	"google.golang.org/grpc"
)

// Config provides values for an Agent.
type Config struct {
	// Hostname the name of host for agent instance.
	Hostname string

	// Managers provides the manager backend used by the agent. It will be
	// updated with managers weights as observed by the agent.
	Managers picker.Remotes

	// Executor specifies the executor to use for the agent.
	Executor exec.Executor

	// Conn specifies the client connection Agent will use
	Conn *grpc.ClientConn

	// NotifyRoleChange channel receives new roles from session messages.
	NotifyRoleChange chan<- api.NodeRole

	DB *bolt.DB
}

func (c *Config) validate() error {
	if c.Conn == nil {
		return fmt.Errorf("config: Connection is required")
	}

	if c.SecurityConfig == nil {
		return fmt.Errorf("agent: SecurityConfig required")
	}

	if c.DB == nil {
		return fmt.Errorf("agent: database required")
	}

	return nil
}
