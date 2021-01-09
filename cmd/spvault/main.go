package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/koltyakov/spvault/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.String("port", ":50051", "SPVault server port, e.g. :50051")
	certFile := flag.String("cert", "", "Certificate file, e.g. ./certs/service.pem")
	keyFile := flag.String("key", "", "Certificate key file, e.g. ./certs/service.key")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var option grpc.ServerOption = grpc.EmptyServerOption{}

	if *certFile != "" && *keyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to setup TLS: %v", err)
		}
		option = grpc.Creds(creds)
	}

	s := grpc.NewServer(option)

	pb.RegisterVaultServer(s, NewVaultServer())
	reflection.Register(s) // Register API reflection

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
