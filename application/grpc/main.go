package main

import (
	"google.golang.org/grpc"
	"grpc-demo/application/grpc/grpc_simple_operation"
	"grpc-demo/application/grpc/grpc_stream_operation"
	"log"
	"net"
)

const ADDRESS = ":8080"
const PROTOCOL = "tcp"

func main() {
	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Attach the handlers to gRPC handler
	grpc_simple_operation.RegisterSimpleOperationServer(grpcServer, &grpc_simple_operation.Server{})
	grpc_stream_operation.RegisterStreamOperationServer(grpcServer, &grpc_stream_operation.Server{})

	// Acquiring a port
	lis, err := net.Listen(PROTOCOL, ADDRESS)

	if err != nil {
		log.Fatalf("Could not acquire port: %v", err)
	}

	// Start serving
	log.Printf("Listening: %v", ADDRESS)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Server could not server: %v", err)
	}
}
