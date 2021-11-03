![License](https://img.shields.io/github/license/p12s/openapi-go-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/openapi-go-sdk?style=plastic)
[![Codecov](https://codecov.io/gh/p12s/openapi-go-sdk/branch/master/graph/badge.svg?token=320DFY1M5Q)](https://codecov.io/gh/p12s/openapi-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/openapi-go-sdk)](https://goreportcard.com/report/github.com/p12s/openapi-go-sdk)
<img src="https://github.com/p12s/openapi-go-sdk/workflows/lint-build-test/badge.svg?branch=master">

# SDK create training
Task description is [here](task.md)

## Implementation
As example for the trainig sdk implementation I took sandbox 
[tinkoffcreditsystems OpenAPI](https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/sandbox/post_sandbox_register)   
and implemented 2 REST methods:  
- Register (POST, sandbox register)
- Orders (GET, sandbox getting orders)

## Where can I get an authentication token?
In the investment section of your personal tinkoff account.  
In your personal account:
- Go to settings;
- Check that the function “Confirmation of transactions by code” is disabled;
- Issue a token for exchange trading and sandbox mode;
- Copy the token and save, the token is displayed only once, you will not be able to view it later, however, you can issue an unlimited number of tokens.

## Example
Installation: 
```
go get -u github.com/p12s/openapi-go-sdk
```
Using:
```
package main

import (
  "context"
  "time"
  "fmt"
  
  sdk "github.com/p12s/openapi-go-sdk"
)

func main() {
  var token := "your token"

  client := sdk.NewSandboxRestClient(token)
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  fmt.Println("Registering a regular account in the sandbox:")
  account, err := client.Register(ctx, sdk.AccountTinkoff)
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(account)
  
  fmt.Println("Getting orders (empty list by default):")
  orders, err := client.Orders(ctx, account.ID)
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(orders)
}
```