FROM golang:latest

ADD PortService /go/src/github.com/x1n13y84issmd42/portobello/PortService
ADD shared /go/src/github.com/x1n13y84issmd42/portobello/shared
WORKDIR /go/src/github.com/x1n13y84issmd42/portobello/PortService

ENV GO111MODULE "off"
RUN go get google.golang.org/grpc
