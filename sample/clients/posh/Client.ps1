#!/usr/bin/env pwsh

Param (
  [Parameter(Mandatory=$true)] [string] $siteUrl,
  [Parameter(Mandatory=$true)] [string] $vaultToken,

  [Parameter(Mandatory=$false)][string] $vaultServer = "localhost:50051"
)

$params = @{ "vaultToken" = $vaultToken }
$auth = $params | ConvertTo-Json -Compress `
  | grpcurl -d '@' -plaintext -emit-defaults $vaultServer spvault.Vault/AuthenticateWithToken `
  | ConvertFrom-Json

$expiration = New-Object -Type DateTime -ArgumentList 1970, 1, 1, 0, 0, 0, 0
$expiration = $expiration.AddSeconds($auth.expiration)

Write-Host "Token: $($auth.authToken)"
Write-Host "Token type: $($auth.tokenType)"
Write-Host "Expiration: $expiration"

$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"

if ("Bearer" -eq $auth.tokenType)
{
  $headers.Add("Autherization", "Bearer $($auth.authToken)")
}
if ("Cookie" -eq $auth.tokenType)
{
  $headers.Add("Cookie", $auth.authToken)
}

$headers.Add("Accept", "application/json")

$response = Invoke-RestMethod `
  $siteUrl'/_api/web?$select=Title' `
  -Method 'GET' -Headers $headers

Write-Host "Site title: $($response.Title)"