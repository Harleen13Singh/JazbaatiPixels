package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	fePb "github.com/personal/jazbaatipixels/api/frontend"
	"github.com/personal/jazbaatipixels/api/vendorgateway"
	"github.com/personal/jazbaatipixels/frontend"
)

func main() {
	list, err := net.Listen("tcp", "localhost:3419")
	if err != nil {
		log.Panic("failed to listen,err: %w", err)
	}
	grpcServer := grpc.NewServer()
	reflection.RegisterV1(grpcServer)
	vgConn, vgConnErr := grpc.Dial("localhost:3418")
	if err != nil {
		log.Panic("failed to connect with vg service,err: %w", vgConnErr)
	}

	vgClient := vendorgateway.NewVendorGatewayClient(vgConn)
	feSvc := frontend.NewService(vgClient)
	fePb.RegisterFrontendServer(grpcServer, feSvc)
	serveErr := grpcServer.Serve(list)
	if serveErr != nil {
		log.Panic("failed to serve the requests, :%w", serveErr)
	}
}
