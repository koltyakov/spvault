private=./config/private.json
server=localhost
port=:50051
scenario=auth:creds
token=73f4fa39-03a5-46f9-b7a1-07c5104bd339
siteUrl=

install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

generate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc -I proto/ proto/spvault.proto --go_out=. --go-grpc_out=. # --experimental_allow_proto3_optional

certs:
	chmod +x ./certs.sh && ./certs.sh

server:
	go run ./cmd/spvault/... port=$(port)

server-tls: certs
	go run ./cmd/spvault/... port=$(port) cert=./certs/service.pem key=./certs/service.key

client-go:
	go run ./sample/clients/go/... -server=$(server)$(port) -private=$(private) -scenario=$(scenario) -token=$(token)

client-dotnet:
	dotnet run --project ./sample/clients/dotnet $(server)$(port) $(token)

client-csom:
	dotnet run --project ./sample/csom $(server)$(port) $(token) $(siteUrl)

client-nodejs:
	cd ./sample/clients/nodejs && npm run client -- $(server)$(port) $(token)

client-python:
	cd ./sample/clients/python && python main.py $(server)$(port) $(token)

client-cli:
	grpcurl -d '{"vaultToken": "$(token)"}' \
		-plaintext -emit-defaults -import-path ./proto -proto spvault.proto \
		$(server)$(port) spvault.Vault/AuthenticateWithToken