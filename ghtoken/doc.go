package ghtoken

/* Package ghtoken provides functionality to retrieve GitHub tokens for authentication.
It supports both GitHub App tokens and Personal Access Tokens (PATs) through environment variables.
It includes a `GhTokenProvider` interface that defines a method to get the token, and implementations for both GitHub App and PAT providers.
It automatically handles token expiration and refreshing for GitHub App tokens.
*/

/* To use this package, set the appropriate environment variables and call `NewGhTokenProviderDefault()` to get an instance of `GhTokenProvider`.
- For GitHub App tokens:
GH_TKN_APP_ID: The ID of the GitHub App.
GH_TKN_APP_INST_ID: The installation ID of the GitHub App.
GH_TKN_APP_PRIVATE_KEY: The private key of the GitHub App in PEM format.
- For Personal Access Tokens (PAT):
GH_TKN: The Personal Access Token.
GITHUB_TOKEN: An alternative name for the Personal Access Token.
GH_TOKEN: Another alternative name for the Personal Access Token.
If neither set of variables is provided, an error will be returned indicating that no credentials were provided.
If both GitHub App and PAT variables are set, the GitHub App token will be used.
For PAT usage, GH_TKN variable is precedent over GITHUB_TOKEN, which is precedent over GH_TOKEN.

// Example usage:
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
*/

/* You can set GH_TKN_API_URL to specify a custom GitHub API URL. */

/* You can override the default environment variable names by creating a custom `GhTokenProviderCfg` struct and passing it to `NewGhTokenProvider()`. For example:
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
*/
