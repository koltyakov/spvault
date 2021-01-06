package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	pb "github.com/koltyakov/spvault/proto"
)

// VaultServer struct
type VaultServer struct {
	pb.VaultServer
	regs sync.Map
}

// NewVaultServer VaultServer constructor
func NewVaultServer() *VaultServer { return &VaultServer{} }

// AuthenticateWithCreds authenticates using credentials
func (s *VaultServer) AuthenticateWithCreds(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
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
func (s *VaultServer) AuthenticateWithToken(ctx context.Context, in *pb.TokenAuthRequest) (*pb.AuthReply, error) {
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
func (s *VaultServer) Register(ctx context.Context, in *pb.RegRequest) (*pb.RegReply, error) {
	if len(in.RegToken) == 0 {
		regToken := uuid.New().String()
		in.RegToken = regToken
	}
	s.regs.Store(in.RegToken, in.AuthRequest)
	res := &pb.RegReply{
		RegToken: in.RegToken,
	}
	return res, nil
}

// DeRegister removes authentication registeration from server's memory
func (s *VaultServer) DeRegister(ctx context.Context, in *pb.DeRegRequest) (*pb.Empty, error) {
	s.regs.Delete(in.RegToken)
	return &pb.Empty{}, nil
}
