package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/koltyakov/spvault/pkg/auth"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthenticatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dat, err := ioutil.ReadFile("./config/private.json")
	if err != nil {
		log.Fatalf("could not read creds file: %v", err)
	}

	creds := &struct {
		SiteURL  string `json:"siteUrl"`
		Strategy string `json:"strategy"`
	}{}

	if err := json.Unmarshal(dat, &creds); err != nil {
		log.Fatalf("could not parse creds: %v", err)
	}

	strategy, ok := pb.Strategy_value[creds.Strategy]
	if !ok {
		log.Fatalf("can't resolve the strategy, %s", creds.Strategy)
	}

	r, err := c.Authenticate(ctx, &pb.AuthRequest{
		SiteUrl:     creds.SiteURL,
		Strategy:    pb.Strategy(strategy),
		Credentials: string(dat),
	})
	if err != nil {
		log.Fatalf("could not authenticate: %v", err)
	}

	log.Printf("Token: %s", r.GetToken())
}
