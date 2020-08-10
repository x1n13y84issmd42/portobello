package main

import (
	"os"

	"github.com/x1n13y84issmd42/portobello/PortClient/server"
	"github.com/x1n13y84issmd42/portobello/PortClient/service"
)

func main() {
	server.New(service.NewGRPCPorts(os.Getenv("GRPCPORTS_ADDR"))).Serve(os.Getenv("LISTEN_ADDR"))
}
