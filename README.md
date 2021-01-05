# SPVault - SharePoint Vault

> SharePoint Authentication Vault gPRC server

__Project status: WIP, early PoC__

## What the project intends solving

- Provide and abstract SharePoint authentication interface via gPRC.
- Issue SharePoint authentication headers/cookies using client tokens (when a client doesn't know actual credentials).
- Showcase some basic gRPC scenarios together with SharePoint ecosystem.

## Development

### Prerequisites

- Protocol Buffers v3 (`protoc` compiler)
- Go v1.16 or greater
- protoc-gen-go-grpc

#### Installing `protoc`

**On a mac:**

```bash
brew install protobuf
```

**In Windows:**

```bash
choco install protoc
```

#### Installing `protoc-gen-go-grpc`

```bash
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Verify `$GOPATH/bin` is in `PATH` environment variable (`export PATH=$PATH:$GOPATH/bin`).

### Gererate Protobuf

```bash
make generate
```

### Run a server

```bash
make server
```

### Run test client

Create `./config/private.json` corresponding to [gosip auth format](https://go.spflow.com/auth/overview). Add and extra field named "strategy". Use one of the possible strategies: addin, adfs, fba, saml, tmg.

Run client

```bash
make client
```

Client output contains auth token/cookie.