# gh-token-go
![format](https://github.com/bombfork/gh-token-go/actions/workflows/format.yml/badge.svg)
![lint](https://github.com/bombfork/gh-token-go/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/bombfork/gh-token-go/actions/workflows/test.yml/badge.svg)

A simple Go module to use Github Application Tokens or Personal Access Tokens (PAT) to authenticate against the Github API.

## Usage

### With default configuration
The default configuration will use these environment variables:
* `GH_TKN_APP_ID` - The ID of the GitHub App.
* `GH_TKN_APP_INST_ID` - The installation ID of the GitHub App.
* `GH_TKN_APP_PRIVATE_KEY` - The private key of the GitHub App in PEM format.
* `GH_TKN` - A Personal Access Token (PAT) to use instead of the GitHub App.
If `GH_TKN` is unset, the stadnard variables `GITHUB_TOKEN` and `GH_TOKEN` will be tried as fallbacks, in that order.
```go
package main

import (
  "fmt"
  github.com/bombfork/gh-token-go/ghtoken
)

func main() {
  provider, err := ghtoken.NewGhTokenProviderDefault()
  if err != nil {
    panic(err)
  }

  token, err := provider.GetToken()
  if err != nil {
    panic(err)
  }

  // Use the token for GitHub API requests
  fmt.Println("GitHub Token:", token)
}
```

### With custom configuration
```go
package main

import (
  "fmt"
  github.com/bombfork/gh-token-go/ghtoken
)

func main() {
  var cfg = GhTokenProviderCfg{
	AppIdVarName:     "MY_OWN_APP_ID_VAR_NAME",
	AppInstIdVarName: "MY_OWN_APP_INST_ID_VAR_NAME",
	AppPemKeyVarName: "MY_OWN_APP_PEM_KEY_VAR_NAME",
	PatVarName:       "MY_OWN_PAT_VAR_NAME",
	GhApiUrlVarName:  "MY_OWN_GH_API_URL_VAR_NAME",
  }
  provider, err := ghtoken.NewGhTokenProvider(cfg)
  if err != nil {
    panic(err)
  }

  token, err := provider.GetToken()
  if err != nil {
    panic(err)
  }

  // Use the token for GitHub API requests
  fmt.Println("GitHub Token:", token)
}
```

If you only want to override some values, you can do:
```go
var cfg = ghtoken.DefaultGhTokenProviderCfg
cfg.PatVarName = "MY_OWN_PAT_VAR_NAME"
```
