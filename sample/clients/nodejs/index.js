//@ts-check

const grpc = require('@grpc/grpc-js');

const { VaultClient } = require('./proto/spvault_grpc_pb.js');
const { TokenAuthRequest } = require('./proto/spvault_pb.js');

const address = process.argv[2] || 'localhost:50051';
const token = process.argv[3] || 'unknown';

const client = new VaultClient(address, grpc.credentials.createInsecure());

const tokenRequest = new TokenAuthRequest();
tokenRequest.setVaulttoken(token);

client.authenticateWithToken(tokenRequest, (err, resp) => {
  if (err) {
    console.error(err.message);
    process.exit(1);
  }

  console.log(`Token: ${resp.getAuthtoken()}`);
  console.log(`Token type: ${['Bearer', 'Cookie', 'Custom'][resp.getTokentype()]}`);
  console.log(`Expiration: ${new Date(resp.getExpiration() * 1000)}`);
});