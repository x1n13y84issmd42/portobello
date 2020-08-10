FROM golang:latest

ADD PortClient /go/src/github.com/x1n13y84issmd42/portobello/PortClient
ADD shared /go/src/github.com/x1n13y84issmd42/portobello/shared
ADD PortClient/ports.json /go/src/github.com/x1n13y84issmd42/portobello/PortClient/ports.json
WORKDIR /go/src/github.com/x1n13y84issmd42/portobello/PortClient

ENV GO111MODULE "off"
RUN go get google.golang.org/grpc
