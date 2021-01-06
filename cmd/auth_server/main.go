package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/koltyakov/spvault/pkg/auth"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeAuthenticatorServer
}

func main() {
	port := flag.String("port", ":50051", "SPVault server port, e.g. :50051")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthenticatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Authenticate(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
	authCnfg, err := resolveAuthCnfg(in)
	if err != nil {
		return nil, err
	}

	token, exp, err := authCnfg.GetAuth()
	if err != nil {
		return nil, err
	}

	res := &pb.AuthReply{
		Token:      token,
		TokenType:  detectTokenType(in),
		Expiration: exp,
	}
	return res, nil
}
