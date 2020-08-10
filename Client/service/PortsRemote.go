package service

import (
	"fmt"

	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// PortsRemote is a remote gRPC Ports service.
// The main production implementation.
type PortsRemote struct {
}

// AddPort adds a port (or updates an existing) to the service.
func (ports *PortsRemote) AddPort(port models.Port) {
}

// GetPort looks up a port by the provided port ID.
func (ports *PortsRemote) GetPort(id string) (*models.Port, error) {
	return nil, fmt.Errorf("not implemented yet")
}
