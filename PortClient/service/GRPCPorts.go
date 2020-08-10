package service

import (
	"context"
	"fmt"
	"log"

	"github.com/x1n13y84issmd42/portobello/shared/errors"
	"github.com/x1n13y84issmd42/portobello/shared/models"
	"github.com/x1n13y84issmd42/portobello/shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCPorts is a remote gRPC Ports service.
// The main production implementation.
type GRPCPorts struct {
	Connection *grpc.ClientConn
	Client     proto.PortsClient
}

// NewGRPCPorts creates a new GRPCPorts instance.
func NewGRPCPorts(grpcHost string) Ports {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	fmt.Printf("Dialing the gRPC server @ %s...\n", grpcHost)

	conn, err := grpc.Dial(fmt.Sprintf("%s", grpcHost), opts...)
	if err != nil {
		log.Fatalf("Failed to dial the gRPC server @ %s because of this: %s", grpcHost, err.Error())
	}

	return &GRPCPorts{
		Connection: conn,
		Client:     proto.NewPortsClient(conn),
	}
}

// Close closes the connection to the gRPC server.
func (ports *GRPCPorts) Close() {
	ports.Connection.Close()
}

// AddPort adds a port (or updates an existing) to the service.
func (ports *GRPCPorts) AddPort(port *models.Port) {
	_, err := ports.Client.AddPort(context.Background(), port.Proto())
	if err != nil {
		panic(err)
	}
}

// GetPort looks up a port by the provided port ID.
func (ports *GRPCPorts) GetPort(id models.PortID) (*models.Port, error) {
	protoPort, err := ports.Client.GetPort(context.Background(), &proto.GetPortRequest{ID: string(id)})
	if err != nil {
		if statusCode := status.Code(err); statusCode == codes.NotFound {
			return nil, errors.PortNotFound(id, "gRPC port service")
		}

		return nil, err
	}

	return models.NewPortFromProto(protoPort), nil
}
