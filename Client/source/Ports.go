package source

import (
	"github.com/x1n13y84issmd42/portobello/Client/service"
	"github.com/x1n13y84issmd42/portobello/shared/models"
)

// PortsChannel is a channel to receive ports during data import.
type PortsChannel chan *models.Port

// PortsReader is a function to read ports data from a source file.
type PortsReader func(string) (PortsChannel, error)

// ImportPorts imports ports from the provided reader into the provided ports service.
func ImportPorts(filePath string, reader PortsReader, ports service.Ports) {
	ch, err := reader(filePath)
	if err != nil {
		panic(err)
	}

	for port := range ch {
		ports.AddPort(port)
	}
}
