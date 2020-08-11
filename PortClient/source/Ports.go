package source

import (
	"io"

	"github.com/x1n13y84issmd42/portobello/PortClient/service"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// PortsChannel is a channel to receive ports during data import.
type PortsChannel chan *models.Port

// PortsReader is a function to read ports data from a source file.
type PortsReader func(io.Reader) (PortsChannel, error)

// ImportPorts imports ports from the provided reader into the provided ports service.
func ImportPorts(r io.Reader, reader PortsReader, ports service.Ports) (chan uint, chan error, error) {
	ch, err := reader(r)
	if err != nil {
		return nil, nil, err
	}

	progressChannel := make(chan uint)
	errorChannel := make(chan error)

	var progress uint = 0

	go func() {
		defer close(progressChannel)
		defer close(errorChannel)

		for port := range ch {
			err = ports.AddPort(port)
			if err != nil {
				errorChannel <- err
				continue
			}

			progress++
			progressChannel <- progress
		}
	}()

	return progressChannel, errorChannel, nil
}
