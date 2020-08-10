package main

import (
	"github.com/x1n13y84issmd42/portobello/Client/server"
	"github.com/x1n13y84issmd42/portobello/Client/service"
)

/*
func main() {
	fmt.Printf("Importing the %s file.\n", os.Args[1])
	ports := service.NewPortsLocal()
	source.ImportPorts(os.Args[1], source.PortsStreamJSONReader, ports)
}*/
//}*/

func main() {
	server.New(service.NewGRPCPorts("localhost", 8080)).Serve("localhost", 80)
}
