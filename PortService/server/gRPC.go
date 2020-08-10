package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/x1n13y84issmd42/portobello/PortService/storage"
	"github.com/x1n13y84issmd42/portobello/shared/models"
	"github.com/x1n13y84issmd42/portobello/shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PortsServer ...
type PortsServer struct {
	Ports storage.Ports
}

// New creates a new PortsServer instance.
func New(store storage.Ports) *PortsServer {
	return &PortsServer{
		Ports: store,
	}
}

// Listen opens a port and listens for gRPC calls.
func (server *PortsServer) Listen(host string, port uint) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to listen @ %s:%d because of %v", host, port, err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterPortsServer(grpcServer, server)

	fmt.Printf("Starting the gRPC server @ %s:%d...\n", host, port)

	grpcServer.Serve(lis)
}

// AddPort ...
func (server *PortsServer) AddPort(ctx context.Context, protoPort *proto.Port) (*proto.Empty, error) {
	fmt.Println("AddPort")
	server.Ports.Add(models.NewPortFromProto(protoPort))
	return &proto.Empty{}, nil
}

// GetPort ...
func (server *PortsServer) GetPort(ctx context.Context, portReq *proto.GetPortRequest) (*proto.Port, error) {
	fmt.Printf("GetPort %s\n", portReq.ID)
	port, err := server.Ports.Get(portReq.ID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return port.Proto(), nil
}
