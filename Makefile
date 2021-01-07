private=./config/private.json
server=localhost
port=:50051
scenario=auth:creds
token=

install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

generate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc -I proto/ proto/spvault.proto --go_out=. --go-grpc_out=. # --experimental_allow_proto3_optional

# generate-dotnet-client:
# 	protoc -I proto/ proto/spvault.proto --csharp_out=./samples/dotnet_client --experimental_allow_proto3_optional
# 	mv ./samples/dotnet_client/Spvault.cs ./samples/dotnet_client/SPVault.cs

server:
	go run ./cmd/spvault/... port=$(port)

client-go:
	go run ./sample/clients/go/... -server=$(server)$(port) -private=$(private) -scenario=$(scenario) -token=$(token)

client-dotnet:
	dotnet run --project ./sample/clients/dotnet $(server)$(port) $(token)

client-nodejs:
	cd ./sample/clients/nodejs && npm run client -- $(server)$(port) $(token)

client-python:
	cd ./sample/clients/python && python main.py $(server)$(port) $(token)

client-cli:
	grpcurl -d '{"regToken": "$(token)"}' \
		-plaintext -emit-defaults -import-path ./proto -proto spvault.proto \
		$(server)$(port) spvault.Vault/AuthenticateWithToken