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

	pb "github.com/koltyakov/spvault/pkg/auth"
)

func resolveAuthCnfg(r *pb.AuthRequest) (gosip.AuthCnfg, error) {
	var authCnfg gosip.AuthCnfg
	c := []byte(r.Credentials)

	switch r.Strategy {
	case pb.Strategy_addin:
		a := &addin.AuthCnfg{}
		if err := json.Unmarshal(c, &a); err != nil {
			return nil, err
		}
		a.SiteURL = r.SiteUrl
		authCnfg = a
	case pb.Strategy_adfs:
		a := &adfs.AuthCnfg{}
		if err := json.Unmarshal(c, &a); err != nil {
			return nil, err
		}
		a.SiteURL = r.SiteUrl
		authCnfg = a
	case pb.Strategy_fba:
		a := &fba.AuthCnfg{}
		if err := json.Unmarshal(c, &a); err != nil {
			return nil, err
		}
		a.SiteURL = r.SiteUrl
		authCnfg = a
	case pb.Strategy_saml:
		a := &saml.AuthCnfg{}
		if err := json.Unmarshal(c, &a); err != nil {
			return nil, err
		}
		a.SiteURL = r.SiteUrl
		authCnfg = a
	case pb.Strategy_tmg:
		a := &tmg.AuthCnfg{}
		if err := json.Unmarshal(c, &a); err != nil {
			return nil, err
		}
		a.SiteURL = r.SiteUrl
		authCnfg = a
	default:
		return nil, fmt.Errorf("can't resolve auth strategy")
	}

	return authCnfg, nil
}
