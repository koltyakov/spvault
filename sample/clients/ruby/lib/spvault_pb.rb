# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spvault.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("spvault.proto", :syntax => :proto3) do
    add_message "spvault.AuthRequest" do
      optional :siteUrl, :string, 1
      optional :strategy, :enum, 2, "spvault.Strategy"
      optional :credentials, :string, 3
    end
    add_message "spvault.TokenAuthRequest" do
      optional :vaultToken, :string, 1
    end
    add_message "spvault.AuthReply" do
      optional :authToken, :string, 1
      optional :tokenType, :enum, 2, "spvault.TokenType"
      optional :expiration, :int64, 3
    end
    add_message "spvault.RegRequest" do
      optional :authRequest, :message, 1, "spvault.AuthRequest"
      optional :vaultToken, :string, 2
    end
    add_message "spvault.RegReply" do
      optional :vaultToken, :string, 1
    end
    add_message "spvault.DeRegRequest" do
      optional :vaultToken, :string, 1
    end
    add_message "spvault.Empty" do
    end
    add_enum "spvault.Strategy" do
      value :addin, 0
      value :adfs, 1
      value :fba, 2
      value :saml, 3
      value :tmg, 4
      value :azurecert, 5
      value :azurecreds, 6
    end
    add_enum "spvault.TokenType" do
      value :Bearer, 0
      value :Cookie, 1
      value :Custom, 3
    end
  end
end

module Spvault
  AuthRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.AuthRequest").msgclass
  TokenAuthRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.TokenAuthRequest").msgclass
  AuthReply = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.AuthReply").msgclass
  RegRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.RegRequest").msgclass
  RegReply = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.RegReply").msgclass
  DeRegRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.DeRegRequest").msgclass
  Empty = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.Empty").msgclass
  Strategy = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.Strategy").enummodule
  TokenType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("spvault.TokenType").enummodule
end
