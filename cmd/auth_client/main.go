package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/koltyakov/spvault/pkg/auth"
	"google.golang.org/grpc"
)

func main() {
	server := flag.String("server", "localhost:50051", "SPVault server address, e.g. server:50051")
	privatePath := flag.String("private", "./config/private.json", "Path to private JSON file")
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthenticatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dat, err := ioutil.ReadFile(*privatePath)
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

	fmt.Printf("Token: %s\n", r.GetToken())
	fmt.Printf("Token type: %s\n", r.GetTokenType())
	fmt.Printf("Expires on: %s\n", time.Unix(r.GetExpiration(), 0))
}
