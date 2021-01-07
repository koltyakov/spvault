import grpc
import datetime
import argparse
from proto import spvault_pb2
from proto import spvault_pb2_grpc

parser = argparse.ArgumentParser()
parser.add_argument("server", help="SPVault gRPC server, e.g. localhost:50051", type=str, default="localhost:50051")
parser.add_argument("token", help="authentication registration token", type=str, default="")
args = parser.parse_args()

def run():
  with grpc.insecure_channel(args.server) as channel:

    client = spvault_pb2_grpc.VaultStub(channel)
    tokenAuthRequest = spvault_pb2.TokenAuthRequest(regToken=args.token)
    response = client.AuthenticateWithToken(tokenAuthRequest)

    expiration = datetime.datetime.fromtimestamp(response.expiration)

    print('Token: ' + response.token)
    print('Token type: ' + ['Bearer', 'Cookie', 'Custom'][response.tokenType])
    print('Expiration: ' + expiration.strftime('%Y-%m-%d %H:%M:%S'))

if __name__ == '__main__':
  run()