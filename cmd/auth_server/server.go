package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	pb "github.com/koltyakov/spvault/proto"
)

type Server struct {
	pb.UnsafeVaultServer
	regs map[string]*pb.AuthRequest // ToDo: add thread safety
}

func NewServer() *Server {
	return &Server{
		regs: map[string]*pb.AuthRequest{},
	}
}

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

func (s *Server) AuthenticateWithToken(ctx context.Context, in *pb.TokenAuthRequest) (*pb.AuthReply, error) {
	authRequest, ok := s.regs[in.RegToken]
	if !ok {
		return nil, fmt.Errorf("no registration found")
	}
	return s.AuthenticateWithCreds(ctx, authRequest)
}

func (s *Server) Register(ctx context.Context, in *pb.RegRequest) (*pb.RegReply, error) {
	if in.RegToken == nil {
		regToken := uuid.New().String()
		in.RegToken = &regToken
	}

	s.regs[*in.RegToken] = in.AuthRequest
	res := &pb.RegReply{
		RegToken: *in.RegToken,
	}

	return res, nil
}

func (s *Server) DeRegister(ctx context.Context, in *pb.DeRegRequest) (*pb.Empty, error) {
	if _, ok := s.regs[in.RegToken]; ok {
		delete(s.regs, in.RegToken)
	}

	return &pb.Empty{}, nil
}
