using System;
using System.Collections.Concurrent;
using System.ComponentModel.DataAnnotations;
using System.Net;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Microsoft.SharePoint.Client;
using Grpc.Core;
using SPVault;

namespace SP.Client
{
    public class AuthManager: IDisposable
    {
        private readonly ConcurrentDictionary<string, AuthReply> tokenCache = new ConcurrentDictionary<string, AuthReply>();
        private readonly Channel channel;
        private readonly Vault.VaultClient client;
        
        public AuthManager(string vaultServer)
        {
            channel = new Channel(vaultServer, ChannelCredentials.Insecure);
            client = new Vault.VaultClient(channel);
        }

        public ClientContext GetContext(Uri web, string vaultToken)
        {
            var context = new ClientContext(web);

            context.ExecutingWebRequest += (sender, e) =>
            {
                var authReply = AcquireTokenAsync(vaultToken).GetAwaiter().GetResult();
                switch (authReply.TokenType)
                {
                    case TokenType.Bearer:
                        e.WebRequestExecutor.RequestHeaders["Authorization"] = "Bearer " + authReply.Token;
                        break;
                    case TokenType.Cookie:
                    {
                        e.WebRequestExecutor.RequestHeaders["Cookie"] = authReply.Token;
                        e.WebRequestExecutor.RequestHeaders.Add("Origin", web.OriginalString);
                        break;
                    }
                    case TokenType.Custom:
                        break;
                    default:
                        throw new ArgumentOutOfRangeException();
                }
            };

            return context;
        }

        private async Task<AuthReply> AcquireTokenAsync(string vaultToken)
        {
            var authReply = TokenFromCache(vaultToken);
            if (authReply != null && !HasTokenExpired(authReply))
            {
                return authReply;
            }
            
            authReply = await client.AuthenticateWithTokenAsync(new TokenAuthRequest { RegToken = vaultToken });
            AddTokenToCache(vaultToken, authReply);

            return authReply;
        }

        private AuthReply TokenFromCache(string vaultToken)
        {
            return tokenCache.TryGetValue(vaultToken, out var authReply) ? authReply : null;
        }

        private void AddTokenToCache(string vaultToken, AuthReply authReply)
        {
            if (tokenCache.TryGetValue(vaultToken, out var currentAuthReply))
            {
                tokenCache.TryUpdate(vaultToken, authReply, currentAuthReply);
            }
            else
            {
                tokenCache.TryAdd(vaultToken, authReply);
            }
        }

        private static bool HasTokenExpired(AuthReply authReply)
        {
            var exriration = new DateTime(1970, 1, 1, 0, 0, 0, 0, System.DateTimeKind.Utc);
            exriration = exriration.AddSeconds(authReply.Expiration);
            return DateTime.Now >= exriration;
        }

        public void Dispose()
        {
            channel.ShutdownAsync().GetAwaiter().GetResult();;
            GC.SuppressFinalize(this);
        }
    }
}