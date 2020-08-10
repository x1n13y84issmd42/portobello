package main

import (
	"os"

	"github.com/x1n13y84issmd42/portobello/PortService/server"
	"github.com/x1n13y84issmd42/portobello/PortService/storage"
)

func main() {
	server.New(storage.NewMemPorts()).Listen(os.Getenv("LISTEN_ADDR"))
}
