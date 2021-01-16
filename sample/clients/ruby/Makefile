PROTO_DIR=../../../proto

install:
	gem install grpc grpc-tools bundler
	bundler install

generate:
	@mkdir -p lib/
	grpc_tools_ruby_protoc -I $(PROTO_DIR)/ --ruby_out=lib --grpc_out=lib $(PROTO_DIR)/spvault.proto

run:
	# bundle exec ./main.rb
	./main.rb localhost:50051 73f4fa39-03a5-46f9-b7a1-07c5104bd339