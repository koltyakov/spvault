package main

import (
	"encoding/json"
	"fmt"

	"github.com/koltyakov/gosip"

	"github.com/koltyakov/gosip/auth/addin"
	"github.com/koltyakov/gosip/auth/adfs"
	"github.com/koltyakov/gosip/auth/fba"
	"github.com/koltyakov/gosip/auth/saml"
	"github.com/koltyakov/gosip/auth/tmg"

	// more strateges can be added, the exception is only ntml and ondemand which won't work that way

	pb "github.com/koltyakov/spvault/pkg/auth"
)

func resolveAuthCnfg(r *pb.AuthRequest) (gosip.AuthCnfg, error) {
	var authCnfg gosip.AuthCnfg
	jsonBytes := []byte(r.Credentials)

	// Inject siteURL into config JSON
	var conf map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &conf); err != nil {
		return nil, err
	}
	conf["siteUrl"] = r.SiteUrl
	jsonBytes, _ = json.Marshal(conf)

	// Get demanded auth strategy
	switch r.Strategy {
	case pb.Strategy_addin:
		authCnfg = &addin.AuthCnfg{}
	case pb.Strategy_adfs:
		authCnfg = &adfs.AuthCnfg{}
	case pb.Strategy_fba:
		authCnfg = &fba.AuthCnfg{}
	case pb.Strategy_saml:
		authCnfg = &saml.AuthCnfg{}
	case pb.Strategy_tmg:
		authCnfg = &tmg.AuthCnfg{}
	default:
		return nil, fmt.Errorf("can't resolve auth strategy")
	}

	if err := authCnfg.ParseConfig(jsonBytes); err != nil {
		return nil, fmt.Errorf("can't parse the config: %s", err)
	}

	return authCnfg, nil
}

func detectTokenType(r *pb.AuthRequest) pb.TokenType {
	s := map[pb.Strategy]pb.TokenType{
		pb.Strategy_addin: pb.TokenType_Bearer,
		pb.Strategy_adfs:  pb.TokenType_Cookie,
		pb.Strategy_fba:   pb.TokenType_Cookie,
		pb.Strategy_saml:  pb.TokenType_Cookie,
		pb.Strategy_tmg:   pb.TokenType_Cookie,
	}
	t, ok := s[r.Strategy]
	if !ok {
		return pb.TokenType_Custom
	}
	return t
}
