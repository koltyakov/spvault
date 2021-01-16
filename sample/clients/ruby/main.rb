#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'date'
require 'warning'

Warning.ignore(//, /.*spvault_.*/)

require 'spvault_services_pb'

def main
  vault_server = ARGV.size > 0 ?  ARGV[0] : 'localhost:50051'
  vault_token = ARGV.size > 1 ?  ARGV[1] : ''

  stub = Spvault::Vault::Stub.new(vault_server, :this_channel_is_insecure)

  begin
    auth_reply = stub.authenticate_with_token(Spvault::TokenAuthRequest.new(vaultToken: vault_token))
    print "Token: #{auth_reply.authToken}\n"
    print "Token type: #{auth_reply.tokenType}\n"
    print "Expiration: #{DateTime.strptime("#{auth_reply.expiration}",'%s')}\n"
  rescue GRPC::BadStatus => e
    abort "ERROR: #{e.message}"
  end
end

main