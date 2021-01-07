package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/koltyakov/spvault/proto"
	"google.golang.org/grpc"
)

func main() {
	server := flag.String("server", "localhost:50051", "SPVault server address, e.g. server:50051")
	privatePath := flag.String("private", "./config/private.json", "Path to private JSON file")
	scenario := flag.String("scenario", "auth:creds", "Sample scenation: auth:creds, auth:token, register, deregister")
	token := flag.String("token", "", "Auth registration token")
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVaultClient(conn)

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

	authRequest := &pb.AuthRequest{
		SiteUrl:     creds.SiteURL,
		Strategy:    pb.Strategy(strategy),
		Credentials: string(dat),
	}

	if *scenario == "register" {
		// Registering an authentication
		// with registration when authentication can be done via `AuthenticateWithToken`
		// the client doesn't know any creds in that case
		reg, err := c.Register(ctx, &pb.RegRequest{AuthRequest: authRequest})
		if err != nil {
			log.Fatalf("could not register authentication: %v", err)
		}
		fmt.Printf("Registration token: %s\n", reg.GetRegToken())
	}

	if *scenario == "auth:token" {
		if *token == "" {
			log.Fatalf("token must be provided")
		}

		// Getting authentication header/cookie which then can be injected to HTTP requests
		auth, err := c.AuthenticateWithToken(ctx, &pb.TokenAuthRequest{RegToken: *token}) // auth using token
		if err != nil {
			log.Fatalf("could not authenticate: %v", err)
		}

		fmt.Printf("Token: %s\n", auth.GetToken())
		fmt.Printf("Token type: %s\n", auth.GetTokenType())
		fmt.Printf("Expires on: %s\n", time.Unix(auth.GetExpiration(), 0))
	}

	if *scenario == "auth:creds" {
		// Getting authentication header/cookie which then can be injected to HTTP requests
		auth, err := c.AuthenticateWithCreds(ctx, authRequest) // auth using creds
		if err != nil {
			log.Fatalf("could not authenticate: %v", err)
		}

		fmt.Printf("Token: %s\n", auth.GetToken())
		fmt.Printf("Token type: %s\n", auth.GetTokenType())
		fmt.Printf("Expires on: %s\n", time.Unix(auth.GetExpiration(), 0))
	}

	if *scenario == "deregister" {
		if *token == "" {
			log.Fatalf("token must be provided")
		}

		// Deregister authentication
		if _, err := c.DeRegister(ctx, &pb.DeRegRequest{RegToken: *token}); err != nil {
			fmt.Printf("error de-registering authentication: %s\n", err)
		}
	}
}
