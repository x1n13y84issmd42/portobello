package storage

import "github.com/x1n13y84issmd42/portobello/shared/models"

// Ports is an interface to a ports storage.
type Ports interface {
	Add(port *models.Port) error
	Get(portID models.PortID) (*models.Port, error)
}
