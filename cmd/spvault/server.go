package main

import (
	"context"
	"fmt"
	"log"
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
	log.Printf("Auth %s using %s", in.SiteUrl, in.Strategy)

	authCnfg, err := resolveAuthCnfg(in)
	if err != nil {
		return nil, err
	}

	authToken, expiration, err := authCnfg.GetAuth()
	if err != nil {
		return nil, err
	}

	res := &pb.AuthReply{
		AuthToken:  authToken,
		TokenType:  detectTokenType(in),
		Expiration: expiration,
	}
	return res, nil
}

// AuthenticateWithToken authenticates using registration token (previously uploaded credentials stored in memory)
func (s *VaultServer) AuthenticateWithToken(ctx context.Context, in *pb.TokenAuthRequest) (*pb.AuthReply, error) {
	log.Printf("Auth with token %s", in.VaultToken)
	reg, ok := s.regs.Load(in.VaultToken)
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
	if len(in.VaultToken) == 0 {
		vaultToken := uuid.New().String()
		in.VaultToken = vaultToken
	}
	log.Printf("Register vault token %s", in.VaultToken)
	s.regs.Store(in.VaultToken, in.AuthRequest)
	res := &pb.RegReply{
		VaultToken: in.VaultToken,
	}
	return res, nil
}

// DeRegister removes authentication registeration from server's memory
func (s *VaultServer) DeRegister(ctx context.Context, in *pb.DeRegRequest) (*pb.Empty, error) {
	log.Printf("De-register vault token %s", in.VaultToken)
	s.regs.Delete(in.VaultToken)
	return &pb.Empty{}, nil
}
