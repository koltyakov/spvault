package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/koltyakov/spvault/proto"
	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", ":50051", "SPVault server port, e.g. :50051")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVaultServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
