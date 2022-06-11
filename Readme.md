# Azure DevOps - Tokens
Small commandline tool for viewing status of tokens in Azure DevOps.
## Docs
To get a list of the flags that can be used.
``` shell
ado-tokens --help
```
For example.
``` shell
ado-tokens --token xxxxxxx --org cotv --filter my-test-token --output table
> DisplayName    IsValid  Expiration                     Scope                  TargetAccounts
> my-test-token  true     2022-05-20T19:33:01.27Z        vso.code               [00000000-0000-0000-0000-000000000000]
> my-test-token  true     2022-08-11T19:41:59.23Z        vso.code vso.project   [00000000-0000-0000-0000-000000000000]
> my-test-token  true     2022-03-18T11:48:24.1266667Z   vso.packaging_write    [00000000-0000-0000-0000-000000000000]
> my-test-token  true     2022-03-09T18:03:30.36Z        app_token              [00000000-0000-0000-0000-000000000000]
```
## Notes and Links
- https://golangbyexample.com/print-struct-variables-golang/
- https://gobyexample.com/command-line-flags