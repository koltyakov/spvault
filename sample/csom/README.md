# SharePoint Client Side Object Model (CSOM)

With SPVault you can use not only AAD but also legacy auth flows with CSOM Standard library.

Another side effect is that CSOM Standard (which is shipped only for SPO) also can be used together with SharePoint On-Prem if it uses token based authentication (not NTLM). Some compatibility issues might be expected, but the majority of CSOM will just work fine.

## Reference

- [Using CSOM for .NET Standard](https://docs.microsoft.com/en-us/sharepoint/dev/sp-add-ins/using-csom-for-dotnet-standard)