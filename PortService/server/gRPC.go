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
func (server *PortsServer) Listen(host string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", host))
	if err != nil {
		log.Fatalf("Failed to listen @ %s because of %v", host, err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterPortsServer(grpcServer, server)

	fmt.Printf("Starting the gRPC server @ %s...\n", host)

	grpcServer.Serve(lis)
}

// AddPort adds port data to the storage.
func (server *PortsServer) AddPort(ctx context.Context, protoPort *proto.Port) (*proto.Empty, error) {
	err := server.Ports.Add(models.NewPortFromProto(protoPort))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

// GetPort looks up port data by it's ID from portReq.
func (server *PortsServer) GetPort(ctx context.Context, portReq *proto.GetPortRequest) (*proto.Port, error) {
	port, err := server.Ports.Get(portReq.ID)
	if err != nil {
		//TODO: handle other errors, i.e. from DBs and alike.
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return port.Proto(), nil
}
