﻿using System;
using System.Threading.Tasks;
using Grpc.Core;
using SPVault;

namespace SP.Client
{
    // ReSharper disable once ClassNeverInstantiated.Global
    public class Program
    {
        private static async Task Main(string[] args)
        {
            var vaultServer = args[0]; // "localhost:50051"
            var vaultToken = args[1];

            var channel = new Channel(vaultServer, ChannelCredentials.Insecure);
            var client = new Vault.VaultClient(channel);

            try {
                var reply = await client.AuthenticateWithTokenAsync(new TokenAuthRequest { VaultToken = vaultToken });

                var exriration = new DateTime(1970, 1, 1, 0, 0, 0, 0, System.DateTimeKind.Utc);
                exriration = exriration.AddSeconds(reply.Expiration).ToLocalTime();

                Console.WriteLine("Token: " + reply.AuthToken);
                Console.WriteLine("Token type: " + reply.TokenType);
                Console.WriteLine("Expires on: " + exriration);
            }
            catch(RpcException e)
            {
                Console.WriteLine($"gRPC error: {e.Status.Detail}");
                Console.WriteLine($"{e}");
            }
            catch
            {
                Console.WriteLine($"Unexpected error calling gRPC");
                throw;
            }

            await channel.ShutdownAsync();
        }
    }
}
