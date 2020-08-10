package service

import (
	"github.com/x1n13y84issmd42/portobello/shared/errors"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

type portsmap map[models.PortID]*models.Port

// MemPorts is a local in-memory Ports service.
// Used for testing & debugging when you want to exclude gRPC layer whatsoever.
type MemPorts struct {
	data portsmap
}

// NewMemPorts creates a new MemPorts instance.
func NewMemPorts() Ports {
	return &MemPorts{
		data: make(portsmap),
	}
}

// Close does nothing for memory-backed service.
func (ports *MemPorts) Close() {
	//
}

// AddPort adds a port (or updates an existing) to the service.
func (ports *MemPorts) AddPort(port *models.Port) error {
	ports.data[port.ID] = port
	return nil
}

// GetPort looks up a port by the provided port ID.
func (ports *MemPorts) GetPort(id models.PortID) (*models.Port, error) {
	if port, ok := ports.data[id]; ok {
		return port, nil
	}

	return nil, errors.PortNotFound(id, "local")
}
