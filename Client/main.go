package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/x1n13y84issmd42/portobello/shared/proto"
	"google.golang.org/grpc"
)

/*
func main() {
	fmt.Printf("Importing the %s file.\n", os.Args[1])
	ports := service.NewPortsLocal()
	source.ImportPorts(os.Args[1], source.PortsStreamJSONReader, ports)
}*/
//}*/

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:80", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewPortsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	port, err := client.GetPort(ctx, &proto.GetPortRequest{ID: "yolo"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Client got a port %#v", port)
}
