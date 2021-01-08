// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var spvault_pb = require('./spvault_pb.js');

function serialize_spvault_AuthReply(arg) {
  if (!(arg instanceof spvault_pb.AuthReply)) {
    throw new Error('Expected argument of type spvault.AuthReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_spvault_AuthReply(buffer_arg) {
  return spvault_pb.AuthReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_spvault_AuthRequest(arg) {
  if (!(arg instanceof spvault_pb.AuthRequest)) {
    throw new Error('Expected argument of type spvault.AuthRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_spvault_AuthRequest(buffer_arg) {
  return spvault_pb.AuthRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_spvault_DeRegRequest(arg) {
  if (!(arg instanceof spvault_pb.DeRegRequest)) {
    throw new Error('Expected argument of type spvault.DeRegRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_spvault_DeRegRequest(buffer_arg) {
  return spvault_pb.DeRegRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_spvault_Empty(arg) {
  if (!(arg instanceof spvault_pb.Empty)) {
    throw new Error('Expected argument of type spvault.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_spvault_Empty(buffer_arg) {
  return spvault_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_spvault_RegReply(arg) {
  if (!(arg instanceof spvault_pb.RegReply)) {
    throw new Error('Expected argument of type spvault.RegReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_spvault_RegReply(buffer_arg) {
  return spvault_pb.RegReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_spvault_RegRequest(arg) {
  if (!(arg instanceof spvault_pb.RegRequest)) {
    throw new Error('Expected argument of type spvault.RegRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_spvault_RegRequest(buffer_arg) {
  return spvault_pb.RegRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_spvault_TokenAuthRequest(arg) {
  if (!(arg instanceof spvault_pb.TokenAuthRequest)) {
    throw new Error('Expected argument of type spvault.TokenAuthRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_spvault_TokenAuthRequest(buffer_arg) {
  return spvault_pb.TokenAuthRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var VaultService = exports.VaultService = {
  authenticateWithCreds: {
    path: '/spvault.Vault/AuthenticateWithCreds',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.AuthRequest,
    responseType: spvault_pb.AuthReply,
    requestSerialize: serialize_spvault_AuthRequest,
    requestDeserialize: deserialize_spvault_AuthRequest,
    responseSerialize: serialize_spvault_AuthReply,
    responseDeserialize: deserialize_spvault_AuthReply,
  },
  authenticateWithToken: {
    path: '/spvault.Vault/AuthenticateWithToken',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.TokenAuthRequest,
    responseType: spvault_pb.AuthReply,
    requestSerialize: serialize_spvault_TokenAuthRequest,
    requestDeserialize: deserialize_spvault_TokenAuthRequest,
    responseSerialize: serialize_spvault_AuthReply,
    responseDeserialize: deserialize_spvault_AuthReply,
  },
  register: {
    path: '/spvault.Vault/Register',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.RegRequest,
    responseType: spvault_pb.RegReply,
    requestSerialize: serialize_spvault_RegRequest,
    requestDeserialize: deserialize_spvault_RegRequest,
    responseSerialize: serialize_spvault_RegReply,
    responseDeserialize: deserialize_spvault_RegReply,
  },
  deRegister: {
    path: '/spvault.Vault/DeRegister',
    requestStream: false,
    responseStream: false,
    requestType: spvault_pb.DeRegRequest,
    responseType: spvault_pb.Empty,
    requestSerialize: serialize_spvault_DeRegRequest,
    requestDeserialize: deserialize_spvault_DeRegRequest,
    responseSerialize: serialize_spvault_Empty,
    responseDeserialize: deserialize_spvault_Empty,
  },
};

exports.VaultClient = grpc.makeGenericClientConstructor(VaultService);
