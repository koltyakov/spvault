// package: SPVault
// file: spvault.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as spvault_pb from "./spvault_pb";

interface IVaultService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    authenticateWithCreds: IVaultService_IAuthenticateWithCreds;
    authenticateWithToken: IVaultService_IAuthenticateWithToken;
    register: IVaultService_IRegister;
    deRegister: IVaultService_IDeRegister;
}

interface IVaultService_IAuthenticateWithCreds extends grpc.MethodDefinition<spvault_pb.AuthRequest, spvault_pb.AuthReply> {
    path: "/SPVault.Vault/AuthenticateWithCreds";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<spvault_pb.AuthRequest>;
    requestDeserialize: grpc.deserialize<spvault_pb.AuthRequest>;
    responseSerialize: grpc.serialize<spvault_pb.AuthReply>;
    responseDeserialize: grpc.deserialize<spvault_pb.AuthReply>;
}
interface IVaultService_IAuthenticateWithToken extends grpc.MethodDefinition<spvault_pb.TokenAuthRequest, spvault_pb.AuthReply> {
    path: "/SPVault.Vault/AuthenticateWithToken";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<spvault_pb.TokenAuthRequest>;
    requestDeserialize: grpc.deserialize<spvault_pb.TokenAuthRequest>;
    responseSerialize: grpc.serialize<spvault_pb.AuthReply>;
    responseDeserialize: grpc.deserialize<spvault_pb.AuthReply>;
}
interface IVaultService_IRegister extends grpc.MethodDefinition<spvault_pb.RegRequest, spvault_pb.RegReply> {
    path: "/SPVault.Vault/Register";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<spvault_pb.RegRequest>;
    requestDeserialize: grpc.deserialize<spvault_pb.RegRequest>;
    responseSerialize: grpc.serialize<spvault_pb.RegReply>;
    responseDeserialize: grpc.deserialize<spvault_pb.RegReply>;
}
interface IVaultService_IDeRegister extends grpc.MethodDefinition<spvault_pb.DeRegRequest, spvault_pb.Empty> {
    path: "/SPVault.Vault/DeRegister";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<spvault_pb.DeRegRequest>;
    requestDeserialize: grpc.deserialize<spvault_pb.DeRegRequest>;
    responseSerialize: grpc.serialize<spvault_pb.Empty>;
    responseDeserialize: grpc.deserialize<spvault_pb.Empty>;
}

export const VaultService: IVaultService;

export interface IVaultServer {
    authenticateWithCreds: grpc.handleUnaryCall<spvault_pb.AuthRequest, spvault_pb.AuthReply>;
    authenticateWithToken: grpc.handleUnaryCall<spvault_pb.TokenAuthRequest, spvault_pb.AuthReply>;
    register: grpc.handleUnaryCall<spvault_pb.RegRequest, spvault_pb.RegReply>;
    deRegister: grpc.handleUnaryCall<spvault_pb.DeRegRequest, spvault_pb.Empty>;
}

export interface IVaultClient {
    authenticateWithCreds(request: spvault_pb.AuthRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    authenticateWithCreds(request: spvault_pb.AuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    authenticateWithCreds(request: spvault_pb.AuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    authenticateWithToken(request: spvault_pb.TokenAuthRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    authenticateWithToken(request: spvault_pb.TokenAuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    authenticateWithToken(request: spvault_pb.TokenAuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    register(request: spvault_pb.RegRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.RegReply) => void): grpc.ClientUnaryCall;
    register(request: spvault_pb.RegRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.RegReply) => void): grpc.ClientUnaryCall;
    register(request: spvault_pb.RegRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.RegReply) => void): grpc.ClientUnaryCall;
    deRegister(request: spvault_pb.DeRegRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.Empty) => void): grpc.ClientUnaryCall;
    deRegister(request: spvault_pb.DeRegRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.Empty) => void): grpc.ClientUnaryCall;
    deRegister(request: spvault_pb.DeRegRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class VaultClient extends grpc.Client implements IVaultClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public authenticateWithCreds(request: spvault_pb.AuthRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    public authenticateWithCreds(request: spvault_pb.AuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    public authenticateWithCreds(request: spvault_pb.AuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    public authenticateWithToken(request: spvault_pb.TokenAuthRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    public authenticateWithToken(request: spvault_pb.TokenAuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    public authenticateWithToken(request: spvault_pb.TokenAuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.AuthReply) => void): grpc.ClientUnaryCall;
    public register(request: spvault_pb.RegRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.RegReply) => void): grpc.ClientUnaryCall;
    public register(request: spvault_pb.RegRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.RegReply) => void): grpc.ClientUnaryCall;
    public register(request: spvault_pb.RegRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.RegReply) => void): grpc.ClientUnaryCall;
    public deRegister(request: spvault_pb.DeRegRequest, callback: (error: grpc.ServiceError | null, response: spvault_pb.Empty) => void): grpc.ClientUnaryCall;
    public deRegister(request: spvault_pb.DeRegRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: spvault_pb.Empty) => void): grpc.ClientUnaryCall;
    public deRegister(request: spvault_pb.DeRegRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: spvault_pb.Empty) => void): grpc.ClientUnaryCall;
}
