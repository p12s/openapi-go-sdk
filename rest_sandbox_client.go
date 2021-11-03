package sdk

import (
	"context"
	"fmt"
)

// see API docs here: https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/sandbox/post_sandbox_register.

// SandboxRestClient - client for working with sandbox
type SandboxRestClient struct {
	*RestClient
}

// NewSandboxRestClient returns new SandboxRestClient by token
func NewSandboxRestClient(token string) *SandboxRestClient {
	return &SandboxRestClient{
		RestClient: NewRestClient(token, WithURL(REST_API_URL+"/sandbox")),
	}
}

// Register in sandbox
func (c *SandboxRestClient) Register(ctx context.Context, accountType AccountType) (Account, error) {

	path := c.url + "/sandbox/register"
	payload := struct {
		AccountType AccountType `json:"brokerAccountType"`
	}{AccountType: accountType}
	var response struct {
		Payload Account `json:"payload"`
	}

	err := c.provider.Post(ctx, path, c.token, payload, &response)
	if err != nil {
		return Account{}, fmt.Errorf("provider post: %w", err)
	}

	return response.Payload, nil
}
