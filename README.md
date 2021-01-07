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

or with a redefined path:

```bash
make client private="./config/private.addin.json"
```

Client output contains auth bearer/cookie:

```txt
Token: eyJ0eXAiOiJ...7OqF7sX2J3JfXKZH2keuqLs_boSDEa47vw
Token type: Bearer
Expires on: 2021-01-07 08:24:25 -0600 CST
```

#### Scenarios

Scenario | Command
---------|--------
Register authentication | ```make client-go scenario=register```
Auth with creds         | ```make client-go scenario=auth:creds```
Auth with token         | ```make client-go scenario=auth:token token=9375a36f-049a-41af-aacc-2caac9e20882```
De-Register auth        | ```make client-go scenario=deregister token=9375a36f-049a-41af-aacc-2caac9e20882```
