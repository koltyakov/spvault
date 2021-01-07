PROTO_DIR=../../../proto

install:
	pip install grpcio-tools

generate:
	python -m grpc_tools.protoc \
		--proto_path=$(PROTO_DIR)/ \
		--python_out=./proto --grpc_python_out=./proto \
		spvault.proto
