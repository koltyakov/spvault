# Ruby SPVault sample

## Prerequisites

- Ruby
- Protocol Buffers v3 (`protoc` compiler)
- Gems:
  - grpc
  - grpc-tools (installs grpc_tools_ruby_protoc, GEM_HOME/bin should be added to PATH environment variable)

## Install dependencies

```bash
make install
```

## Generate gRPC assets

```bash
make generate
```

## Run

```bash
./main.py localhost:50051 afc1e997-81a4-4943-9981-8972a0e665de
```