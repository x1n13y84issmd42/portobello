package storage

import (
	"github.com/x1n13y84issmd42/portobello/shared/errors"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

type portsmap map[models.PortID]*models.Port

// MemPorts is a memory port storage.
type MemPorts struct {
	data portsmap
}

// NewMemPorts creates a new MemPorts instance.
func NewMemPorts() Ports {
	return &MemPorts{
		data: portsmap{
			"yolo": &models.Port{
				ID:   "YOLO",
				Name: "You only visit this port once.",
			},
		},
	}
}

// Add adds a port to the storage.
func (ports *MemPorts) Add(port *models.Port) error {
	ports.data[port.ID] = port
	return nil
}

// Get retrieves a port by it's ID.
func (ports *MemPorts) Get(portID models.PortID) (*models.Port, error) {
	if port, ok := ports.data[portID]; ok {
		return port, nil
	}

	return nil, errors.PortNotFound(portID, "memory storage")
}
