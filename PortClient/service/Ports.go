package service

import (
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// Ports is an interface to a Ports service.
type Ports interface {
	AddPort(port *models.Port)
	GetPort(id models.PortID) (*models.Port, error)

	Close()
}
