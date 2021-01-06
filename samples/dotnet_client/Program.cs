using System;
using System.Net.Http;
using System.Threading.Tasks;
using Grpc.Net.Client;
using SPVault;

namespace SP.Client
{
    // ReSharper disable once ClassNeverInstantiated.Global
    public class Program
    {
        private static async Task Main(string[] args)
        {
            using var channel = GrpcChannel.ForAddress("https://localhost:50051");

            var client = new Vault.VaultClient(channel);

            var reply = await client.AuthenticateWithTokenAsync(new TokenAuthRequest
            {
                RegToken = "4f9e296e-3389-4cbc-a5a1-331b0d35771d" // copy from `make client` output
            });

            Console.WriteLine("Token: " + reply.Token);
            Console.WriteLine("Token type: " + reply.TokenType);
            Console.WriteLine("Expires on: " + reply.Expiration);
        }
    }
}
