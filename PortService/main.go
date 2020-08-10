package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/x1n13y84issmd42/portobello/shared/proto"
	"google.golang.org/grpc"
)

// PortsServer ...
type PortsServer struct {
}

// AddPort ...
func (server *PortsServer) AddPort(ctx context.Context, port *proto.Port) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

// GetPort ...
func (server *PortsServer) GetPort(ctx context.Context, port *proto.GetPortRequest) (*proto.Port, error) {
	return &proto.Port{
		ID:   "TEST",
		Name: "Test port",
	}, nil
}

func newPortsServer() *PortsServer {
	s := &PortsServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 80))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterPortsServer(grpcServer, newPortsServer())
	grpcServer.Serve(lis)
}
