package service

import (
	"fmt"

	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// Ports is an interface to a Ports service.
type Ports interface {
	AddPort(port *models.Port)
	GetPort(id string) (*models.Port, error)
}

// ErrPortNotFound is returned from Ports services when they're unable to find a port.
type ErrPortNotFound struct {
	ID        string
	ServiceID string
}

func (err ErrPortNotFound) Error() string {
	return fmt.Sprintf("Port \"%s\" is not found in the %s service.", err.ID, err.ServiceID)
}

// PortNotFound creaes a new ErrPortNotFound instance.
func PortNotFound(id string, serviceID string) error {
	return ErrPortNotFound{
		ID:        id,
		ServiceID: serviceID,
	}
}
