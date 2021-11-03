package sdk

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// see API docs here: https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/sandbox/post_sandbox_register.

// REST_API_URL - common API path
const REST_API_URL = "https://api-invest.tinkoff.ru/openapi"

// TIMEOUT - by default 30 secons wait
const TIMEOUT = time.Second * 30

// RestClient provide to rest methods from tinkoff invest api.
type RestClient struct {
	provider Provider
	token    string
	url      string
}

// BuildOption build options for rest client.
type BuildOption func(*RestClient)

// WithURL build rest client by custom api url.
func WithURL(url string) BuildOption {
	return func(client *RestClient) {
		client.url = url
	}
}

// NewRestClient - constructor
func NewRestClient(token string, options ...BuildOption) *RestClient {
	client := &RestClient{
		provider: &defaultHTTP{
			client: &http.Client{
				Transport: http.DefaultTransport,
				Timeout:   TIMEOUT,
			},
		},
		token: token,
		url:   REST_API_URL,
	}

	for i := range options {
		options[i](client)
	}

	return client
}

// Orders see docs https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/orders/get_orders.
func (c *RestClient) Orders(ctx context.Context, accountID string) ([]Order, error) {

	path := c.url + "/orders"
	if accountID != DefaultAccount {
		path += "?brokerAccountId=" + accountID
	}
	var response struct {
		Payload []Order `json:"payload"`
	}

	err := c.provider.Get(ctx, path, c.token, &response)
	if err != nil {
		return nil, fmt.Errorf("orders getting: %w", err)
	}

	return response.Payload, nil
}
