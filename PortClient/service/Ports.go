package service

import (
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// Ports is an interface to a Ports service.
type Ports interface {

	// AddPort adds a port (or updates an existing) to the service.
	AddPort(port *models.Port) error

	// GetPort looks up a port by the provided port ID.
	GetPort(id models.PortID) (*models.Port, error)

	// Close closes the service.
	Close()
}
