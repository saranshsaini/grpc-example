package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"saranshsaini/arithmetic_pb"
	"saranshsaini/arithmetic_server"
	grpc "google.golang.org/grpc"
)



func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:50001"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	arithmetic_pb.RegisterArithmeticServiceServer(grpcServer, &arithmetic_server.ArithmeticServer{})

	fmt.Println("Starting")
	grpcServer.Serve(lis)
}
