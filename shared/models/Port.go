package models

import "github.com/x1n13y84issmd42/portobello/shared/proto"

// PortID is a port ID, y'know.
type PortID = string

// Port describes an entry in the ports.json file.
// This is a canonical description of a port,
// the protocol-level as well as database-level
// port structures are derived from it.
type Port struct {
	ID          PortID
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

// Proto creates a proto.Port instance from the Port data.
func (port Port) Proto() *proto.Port {
	protoPort := &proto.Port{
		ID:       port.ID,
		Name:     port.Name,
		City:     port.City,
		Country:  port.Country,
		Alias:    port.Alias,
		Regions:  port.Regions,
		Province: port.Province,
		Timezone: port.Timezone,
		Unlocs:   port.Unlocs,
		Code:     port.Code,
	}

	for _, f64 := range port.Coordinates {
		protoPort.Coordinates = append(protoPort.Coordinates, float32(f64))
	}

	return protoPort
}

// NewPortFromProto creates a new Port instance from the provided proto.Port data.
func NewPortFromProto(protoPort *proto.Port) *Port {
	port := &Port{
		ID:       protoPort.ID,
		Name:     protoPort.Name,
		City:     protoPort.City,
		Country:  protoPort.Country,
		Alias:    protoPort.Alias,
		Regions:  protoPort.Regions,
		Province: protoPort.Province,
		Timezone: protoPort.Timezone,
		Unlocs:   protoPort.Unlocs,
		Code:     protoPort.Code,
	}

	for _, f32 := range protoPort.Coordinates {
		port.Coordinates = append(port.Coordinates, float64(f32))
	}

	return port
}
