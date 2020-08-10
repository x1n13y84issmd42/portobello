package main

import (
	"fmt"
	"os"

	"github.com/x1n13y84issmd42/portobello/Client/service"
	"github.com/x1n13y84issmd42/portobello/Client/source"
)

func main() {
	fmt.Printf("Importing the %s file.\n", os.Args[1])
	ports := service.NewPortsLocal()
	source.ImportPorts(os.Args[1], source.PortsStreamJSONReader, ports)
}
