package ghtoken

/* Package ghtoken provides functionality to retrieve GitHub tokens for authentication.
It supports both GitHub App tokens and Personal Access Tokens (PATs) through environment variables.
It includes a `GhTokenProvider` interface that defines a method to get the token, and implementations for both GitHub App and PAT providers.
It automatically handles token expiration and refreshing for GitHub App tokens.
*/

/* To use this package, set the appropriate environment variables and call `NewGhTokenProvider()` to get an instance of `GhTokenProvider`.
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
  provider, err := ghtoken.NewGhTokenProvider()
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
