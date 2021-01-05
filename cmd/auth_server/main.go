package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/koltyakov/spvault/pkg/auth"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnsafeAuthenticatorServer
}

func (s *server) Authenticate(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
	authCnfg, err := resolveAuthCnfg(in)
	if err != nil {
		return nil, err
	}

	expiration := int64(60 * time.Minute) // ToDo: Expose getting expiration in GetAuth flow
	token, err := authCnfg.GetAuth()
	if err != nil {
		return nil, err
	}

	res := &pb.AuthReply{
		Token:      token,
		Expiration: expiration,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthenticatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
