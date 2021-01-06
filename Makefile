private="./config/private.json"

install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

generate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc -I proto/ proto/spvault.proto --go_out=. --go-grpc_out=. --experimental_allow_proto3_optional

server:
	go run ./cmd/auth_server/...

client:
	go run ./cmd/auth_client/... -private=$(private)