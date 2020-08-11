package errors

import (
	"fmt"

	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// ErrPortNotFound is returned from Ports services and storages when they're unable to find a port.
type ErrPortNotFound struct {
	ID       models.PortID
	SourceID string
}

func (err ErrPortNotFound) Error() string {
	return fmt.Sprintf("Port \"%s\" is not found in the %s.", err.ID, err.SourceID)
}

// PortNotFound creaes a new ErrPortNotFound instance.
func PortNotFound(id models.PortID, sourceID string) error {
	return ErrPortNotFound{
		ID:       id,
		SourceID: sourceID,
	}
}
