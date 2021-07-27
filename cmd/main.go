package main

import (
	"context"
	"log"
	"net"

	v1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/grpc"
)

const (
	port = ":4317"
)

type server struct {
	v1.UnimplementedTraceServiceServer
}

func (s server) Export(ctx context.Context, req *v1.ExportTraceServiceRequest) (*v1.ExportTraceServiceResponse, error) {
	log.Printf("%v\n", req)
	return &v1.ExportTraceServiceResponse{}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterTraceServiceServer(s, server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
