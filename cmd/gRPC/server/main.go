package main

import (
	"fmt"
	"log"
	"net"

	"gosearch/internal/gRPC/service"
	"gosearch/internal/gRPC/impl"

	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(8888)
	gRPCServer := grpc.NewServer()

	searchService := impl.NewSearchService()
	service.RegisterSearchServiceServer(gRPCServer, searchService)

	// Start the server
	
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve gRPC server: %s", err)
	}
}

func getNetListener(port uint) net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		panic(fmt.Sprintf("Failed to listen: %v", err))
	} 

	return listener
}