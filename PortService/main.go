package main

import (
	"github.com/x1n13y84issmd42/portobello/PortService/server"
	"github.com/x1n13y84issmd42/portobello/PortService/storage"
)

func main() {
	server.New(storage.NewMemPorts()).Listen("localhost", 8080)
}
