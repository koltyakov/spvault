package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	pb "github.com/koltyakov/spvault/proto"
)

// Server struct
type Server struct {
	pb.UnsafeVaultServer
	regs sync.Map
}

// NewServer Server constructor
func NewServer() *Server { return &Server{} }

// AuthenticateWithCreds authenticates using credentials
func (s *Server) AuthenticateWithCreds(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
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

// AuthenticateWithToken authenticates using registration token (previously uploaded credentials stored in memory)
func (s *Server) AuthenticateWithToken(ctx context.Context, in *pb.TokenAuthRequest) (*pb.AuthReply, error) {
	reg, ok := s.regs.Load(in.RegToken)
	if !ok {
		return nil, fmt.Errorf("no registration found")
	}
	authRequest, ok := reg.(*pb.AuthRequest)
	if !ok {
		return nil, fmt.Errorf("can't cast auth registration")
	}
	return s.AuthenticateWithCreds(ctx, authRequest)
}

// Register registers authentication in server's memory
func (s *Server) Register(ctx context.Context, in *pb.RegRequest) (*pb.RegReply, error) {
	if in.RegToken == nil {
		regToken := uuid.New().String()
		in.RegToken = &regToken
	}
	s.regs.Store(*in.RegToken, in.AuthRequest)
	res := &pb.RegReply{
		RegToken: *in.RegToken,
	}
	return res, nil
}

// DeRegister removes authentication registeration from server's memory
func (s *Server) DeRegister(ctx context.Context, in *pb.DeRegRequest) (*pb.Empty, error) {
	s.regs.Delete(in.RegToken)
	return &pb.Empty{}, nil
}
