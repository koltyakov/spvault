#!/usr/bin/env sh

vaultServer=$1
vaultToken=$2
siteUrl=$3

auth=$(grpcurl -d '{"vaultToken": "'$vaultToken'"}' -plaintext -emit-defaults $vaultServer spvault.Vault/AuthenticateWithToken)

authToken=$(jq -r '.authToken' <<< $auth)
tokenType=$(jq -r '.tokenType' <<< $auth)
expiration=$(jq -r '.expiration' <<< $auth)

echo "Token: $authToken"
echo "Token type: $tokenType"
echo "Expiration: $(date -r $expiration)"

authHeader=""
if [ "$tokenType" = "Bearer" ]; then
  authHeader="Authentication: Bearer $authToken"
fi
if [ "$tokenType" = "Cookie" ]; then
  authHeader="Cookie: $authToken"
fi

title=$(curl -s -H "Accept: application/json" -H "$authHeader" $siteUrl/_api/web?\$select=Title | jq -r '.Title')

echo "Site title: $title"