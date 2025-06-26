# gh-token-go
![format](https://github.com/bombfork/gh-token-go/actions/workflows/format.yml/badge.svg)
![lint](https://github.com/bombfork/gh-token-go/actions/workflows/lint.yml/badge.svg)
![test](https://github.com/bombfork/gh-token-go/actions/workflows/test.yml/badge.svg)

A simple Go module to use Github Application Tokens or Personal Access Tokens (PAT) to authenticate against the Github API.

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
