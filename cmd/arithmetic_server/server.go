package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"saranshsaini/arithmetic_pb"
	"saranshsaini/arithmetic_server"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:50051"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	arithmetic_pb.RegisterArithmeticServiceServer(grpcServer, &arithmetic_server.ArithmeticServer{})

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	fmt.Println("Starting")
	grpcServer.Serve(lis)
}
