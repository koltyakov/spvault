private="./config/private.json"

install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

generate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc -I proto/ proto/spvault.proto --go_out=. --go-grpc_out=. # --experimental_allow_proto3_optional

server:
	go run ./cmd/spvault/...

client:
	go run ./samples/go_client/... -private=$(private)

# generate-dotnet-client:
# 	protoc -I proto/ proto/spvault.proto --csharp_out=./samples/dotnet_client --experimental_allow_proto3_optional
# 	mv ./samples/dotnet_client/Spvault.cs ./samples/dotnet_client/SPVault.cs