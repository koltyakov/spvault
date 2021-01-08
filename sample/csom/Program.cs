using System;
using System.Threading.Tasks;

namespace SP.Client
{
    // ReSharper disable once ClassNeverInstantiated.Global
    public class Program
    {
        private static async Task Main(string[] args)
        {
            var vaultServer = args[0]; // "localhost:50051"
            var vaultToken = args[1];
            var siteUrl = args[2];

            using var authManager = new AuthManager(vaultServer);
            using var ctx = authManager.GetContext(new Uri(siteUrl), vaultToken);

            ctx.Load(ctx.Web, p => p.Title);
            await ctx.ExecuteQueryAsync();
            Console.WriteLine($"Title: {ctx.Web.Title}");
        }
    }
}
