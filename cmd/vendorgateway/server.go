package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	vgPb "github.com/personal/jazbaatipixels/api/vendorgateway"
	vg "github.com/personal/jazbaatipixels/vendorgateway"
)

func main() {
	list, err := net.Listen("tcp", "localhost:3418")
	if err != nil {
		log.Panic("failed to listen,err: %w", err)
	}
	grpcServer := grpc.NewServer()
	reflection.RegisterV1(grpcServer)
	vgSvc := vg.NewService()
	vgPb.RegisterVendorGatewayServer(grpcServer, vgSvc)
	serveErr := grpcServer.Serve(list)
	if serveErr != nil {
		log.Panic("failed to serve the requests, :%w", serveErr)
	}
}
