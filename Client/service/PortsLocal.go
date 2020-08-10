package service

import (
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

type portsmap map[string]*models.Port

// PortsLocal is a local in-memory Ports service.
// Used for testing.
type PortsLocal struct {
	data portsmap
}

// NewPortsLocal creates a new PortsLocal instance.
func NewPortsLocal() Ports {
	return &PortsLocal{
		data: make(portsmap),
	}
}

// AddPort adds a port (or updates an existing) to the service.
func (ports *PortsLocal) AddPort(port *models.Port) {
	ports.data[port.ID] = port
}

// GetPort looks up a port by the provided port ID.
func (ports *PortsLocal) GetPort(id string) (*models.Port, error) {
	if port, ok := ports.data[id]; ok {
		return port, nil
	}

	return nil, PortNotFound(id, "local")
}
