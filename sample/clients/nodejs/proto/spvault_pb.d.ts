// package: spvault
// file: spvault.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class AuthRequest extends jspb.Message { 
    getSiteurl(): string;
    setSiteurl(value: string): AuthRequest;

    getStrategy(): Strategy;
    setStrategy(value: Strategy): AuthRequest;

    getCredentials(): string;
    setCredentials(value: string): AuthRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AuthRequest): AuthRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthRequest;
    static deserializeBinaryFromReader(message: AuthRequest, reader: jspb.BinaryReader): AuthRequest;
}

export namespace AuthRequest {
    export type AsObject = {
        siteurl: string,
        strategy: Strategy,
        credentials: string,
    }
}

export class TokenAuthRequest extends jspb.Message { 
    getVaulttoken(): string;
    setVaulttoken(value: string): TokenAuthRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TokenAuthRequest.AsObject;
    static toObject(includeInstance: boolean, msg: TokenAuthRequest): TokenAuthRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TokenAuthRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TokenAuthRequest;
    static deserializeBinaryFromReader(message: TokenAuthRequest, reader: jspb.BinaryReader): TokenAuthRequest;
}

export namespace TokenAuthRequest {
    export type AsObject = {
        vaulttoken: string,
    }
}

export class AuthReply extends jspb.Message { 
    getAuthtoken(): string;
    setAuthtoken(value: string): AuthReply;

    getTokentype(): TokenType;
    setTokentype(value: TokenType): AuthReply;

    getExpiration(): number;
    setExpiration(value: number): AuthReply;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthReply.AsObject;
    static toObject(includeInstance: boolean, msg: AuthReply): AuthReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthReply;
    static deserializeBinaryFromReader(message: AuthReply, reader: jspb.BinaryReader): AuthReply;
}

export namespace AuthReply {
    export type AsObject = {
        authtoken: string,
        tokentype: TokenType,
        expiration: number,
    }
}

export class RegRequest extends jspb.Message { 

    hasAuthrequest(): boolean;
    clearAuthrequest(): void;
    getAuthrequest(): AuthRequest | undefined;
    setAuthrequest(value?: AuthRequest): RegRequest;

    getVaulttoken(): string;
    setVaulttoken(value: string): RegRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RegRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RegRequest): RegRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RegRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RegRequest;
    static deserializeBinaryFromReader(message: RegRequest, reader: jspb.BinaryReader): RegRequest;
}

export namespace RegRequest {
    export type AsObject = {
        authrequest?: AuthRequest.AsObject,
        vaulttoken: string,
    }
}

export class RegReply extends jspb.Message { 
    getVaulttoken(): string;
    setVaulttoken(value: string): RegReply;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RegReply.AsObject;
    static toObject(includeInstance: boolean, msg: RegReply): RegReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RegReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RegReply;
    static deserializeBinaryFromReader(message: RegReply, reader: jspb.BinaryReader): RegReply;
}

export namespace RegReply {
    export type AsObject = {
        vaulttoken: string,
    }
}

export class DeRegRequest extends jspb.Message { 
    getVaulttoken(): string;
    setVaulttoken(value: string): DeRegRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeRegRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeRegRequest): DeRegRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeRegRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeRegRequest;
    static deserializeBinaryFromReader(message: DeRegRequest, reader: jspb.BinaryReader): DeRegRequest;
}

export namespace DeRegRequest {
    export type AsObject = {
        vaulttoken: string,
    }
}

export class Empty extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Empty.AsObject;
    static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Empty;
    static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
    export type AsObject = {
    }
}

export enum Strategy {
    ADDIN = 0,
    ADFS = 1,
    FBA = 2,
    SAML = 3,
    TMG = 4,
}

export enum TokenType {
    BEARER = 0,
    COOKIE = 1,
    CUSTOM = 3,
}
