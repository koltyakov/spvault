private="./config/private.json"

install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

generate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc -I proto/ proto/auth.proto --go_out=. --go-grpc_out=.

server:
	go run ./cmd/auth_server/...

client:
	go run ./cmd/auth_client/... -private=$(private)