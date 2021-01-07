# Python SPVault sample

## Prerequisites

- Python
- Protocol Buffers v3 (`protoc` compiler)
- grpcio-tools

## Install dependencies

```bash
make install
```

## Generate gRPC assets (stubs)

```bash
make generate
```

## Run

```bash
python main.py localhost:50051 afc1e997-81a4-4943-9981-8972a0e665de
```