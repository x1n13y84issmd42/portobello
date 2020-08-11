package source

import (
	"io"

	"github.com/x1n13y84issmd42/portobello/PortClient/service"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// PortsChannel is a channel to receive ports during data import.
type PortsChannel chan *models.Port

// ErrorChannel is a channel to feed errors through.
type ErrorChannel chan error

// PortsReader is a function to read ports data from a source file.
type PortsReader func(io.Reader, ErrorChannel) PortsChannel

// ImportPorts imports ports from the provided reader into the provided ports service.
func ImportPorts(r io.Reader, reader PortsReader, ports service.Ports) (chan uint, ErrorChannel, error) {
	errorChannel := make(chan error)
	progressChannel := make(chan uint)

	portsChan := reader(r, errorChannel)

	var progress uint = 0

	go func() {
		defer close(progressChannel)
		defer close(errorChannel)

		for port := range portsChan {
			err := ports.AddPort(port)
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
