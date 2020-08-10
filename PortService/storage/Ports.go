package storage

import "github.com/x1n13y84issmd42/portobello/shared/models"

// Ports is an interface to a ports storage.
type Ports interface {

	// Add adds a port to the storage.
	Add(port *models.Port) error

	// GetPort looks up port data by it's ID from portReq.
	Get(portID models.PortID) (*models.Port, error)
}
