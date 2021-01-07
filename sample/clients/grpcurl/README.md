# gRPCurl - like curl but for gRPC

[gRPCurl project page](https://github.com/fullstorydev/grpcurl)

## Install

On a mac:

```bash
brew install grpcurl
```

or from sources:

```bash
go get github.com/fullstorydev/grpcurl/...
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

## Usage

```bash
grpcurl \
  -d '{"regToken": "e23c1bb2-1334-4bba-8961-38090905f726"}' \
  -plaintext -emit-defaults localhost:50051 \
  spvault.Vault/AuthenticateWithToken
```

or with `-import-path ./proto -proto spvault.proto` if API reflection is not enabled:

```bash
grpcurl \
  -d '{"regToken": "e23c1bb2-1334-4bba-8961-38090905f726"}' \
  -plaintext -emit-defaults \
  -import-path ./proto -proto spvault.proto \
  localhost:50051 spvault.Vault/AuthenticateWithToken
```

Where, `e23c1bb2-1334-4bba-8961-38090905f726` is a registration token received previously using `.Register()` method.

`-plaintext` is needed in a dev mode when no TLS is configured.

`-emit-defaults` is [needed](https://github.com/fullstorydev/grpcurl/issues/95) to return 0 position enum values.

`-import-path ./proto -proto spvault.proto` is required for gRPC servers with disabled API reflection. Arguments values should correspond to actual path and `.proto` file name.