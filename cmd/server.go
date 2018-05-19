//go:generate protoc -I echo --go_out=plugins=grpc:echo echo/echo.proto

// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	mwgrpc "github.com/grpc-ecosystem/go-grpc-middleware"
	otgrpc "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	pb "github.com/kublr/workshop-microservice-build-pipeline-colorer/pkg/colorer"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 10000, "The server port")
)

func main() {
	// parse flags
	flag.Parse()

	// prepare requested port listener
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// prepare server options including interceptors
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(mwgrpc.ChainStreamServer(
			otgrpc.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(mwgrpc.ChainUnaryServer(
			otgrpc.UnaryServerInterceptor(),
		)),
	}

	// ... including TLS options
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	// create gRPC server
	grpcServer := grpc.NewServer(opts...)

	// register gRPC procedures handler
	pb.RegisterColorerServer(grpcServer, pb.NewServer())

	// start server
	grpcServer.Serve(lis)
}
