# gh-token-go

A simple Go module to use Github Application Tokens or Personal Access Tokens (PAT) to authenticate with the Github API.

## Usage

```go
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
```
