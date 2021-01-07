// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var spvault_pb = require('./spvault_pb.js');

function serialize_SPVault_AuthReply(arg) {
  if (!(arg instanceof spvault_pb.AuthReply)) {
    throw new Error('Expected argument of type SPVault.AuthReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SPVault_AuthReply(buffer_arg) {
  return spvault_pb.AuthReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SPVault_AuthRequest(arg) {
  if (!(arg instanceof spvault_pb.AuthRequest)) {
    throw new Error('Expected argument of type SPVault.AuthRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SPVault_AuthRequest(buffer_arg) {
  return spvault_pb.AuthRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SPVault_DeRegRequest(arg) {
  if (!(arg instanceof spvault_pb.DeRegRequest)) {
    throw new Error('Expected argument of type SPVault.DeRegRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SPVault_DeRegRequest(buffer_arg) {
  return spvault_pb.DeRegRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SPVault_Empty(arg) {
  if (!(arg instanceof spvault_pb.Empty)) {
    throw new Error('Expected argument of type SPVault.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SPVault_Empty(buffer_arg) {
  return spvault_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SPVault_RegReply(arg) {
  if (!(arg instanceof spvault_pb.RegReply)) {
    throw new Error('Expected argument of type SPVault.RegReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SPVault_RegReply(buffer_arg) {
  return spvault_pb.RegReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SPVault_RegRequest(arg) {
  if (!(arg instanceof spvault_pb.RegRequest)) {
    throw new Error('Expected argument of type SPVault.RegRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SPVault_RegRequest(buffer_arg) {
  return spvault_pb.RegRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SPVault_TokenAuthRequest(arg) {
  if (!(arg instanceof spvault_pb.TokenAuthRequest)) {
    throw new Error('Expected argument of type SPVault.TokenAuthRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_SPVault_TokenAuthRequest(buffer_arg) {
  return spvault_pb.TokenAuthRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var VaultService = exports.VaultService = {
  authenticateWithCreds: {
    path: '/SPVault.Vault/AuthenticateWithCreds',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.AuthRequest,
    responseType: spvault_pb.AuthReply,
    requestSerialize: serialize_SPVault_AuthRequest,
    requestDeserialize: deserialize_SPVault_AuthRequest,
    responseSerialize: serialize_SPVault_AuthReply,
    responseDeserialize: deserialize_SPVault_AuthReply,
  },
  authenticateWithToken: {
    path: '/SPVault.Vault/AuthenticateWithToken',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.TokenAuthRequest,
    responseType: spvault_pb.AuthReply,
    requestSerialize: serialize_SPVault_TokenAuthRequest,
    requestDeserialize: deserialize_SPVault_TokenAuthRequest,
    responseSerialize: serialize_SPVault_AuthReply,
    responseDeserialize: deserialize_SPVault_AuthReply,
  },
  register: {
    path: '/SPVault.Vault/Register',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.RegRequest,
    responseType: spvault_pb.RegReply,
    requestSerialize: serialize_SPVault_RegRequest,
    requestDeserialize: deserialize_SPVault_RegRequest,
    responseSerialize: serialize_SPVault_RegReply,
    responseDeserialize: deserialize_SPVault_RegReply,
  },
  deRegister: {
    path: '/SPVault.Vault/DeRegister',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.DeRegRequest,
    responseType: spvault_pb.Empty,
    requestSerialize: serialize_SPVault_DeRegRequest,
    requestDeserialize: deserialize_SPVault_DeRegRequest,
    responseSerialize: serialize_SPVault_Empty,
    responseDeserialize: deserialize_SPVault_Empty,
  },
};

exports.VaultClient = grpc.makeGenericClientConstructor(VaultService);
