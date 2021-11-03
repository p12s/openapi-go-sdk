package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var _ Provider = &defaultHTTP{}

// Provider
type Provider interface {
	Get(ctx context.Context, url string, token string, unmarshal interface{}) error
	Post(ctx context.Context, url string, token string, payload, unmarshal interface{}) error
}

// defaultHTTP
type defaultHTTP struct {
	client *http.Client
}

// Post - getting POST request
func (c *defaultHTTP) Post(ctx context.Context, url string, token string, payload, unmarshal interface{}) error {
	var body io.ReadWriter
	if payload != nil {
		buf, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("post-request json marshal: %w", err)
		}
		body = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("post-request building: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.do(req)
	if err != nil {
		return fmt.Errorf("post-request doing: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println("post-request response body close error:", err.Error())
		}
	}()

	if unmarshal != nil {
		err = json.NewDecoder(resp.Body).Decode(unmarshal)
		if err != nil {
			return fmt.Errorf("post-response decode: %w", err)
		}
	}

	return nil
}

// Get - getting GET request
func (c *defaultHTTP) Get(ctx context.Context, url string, token string, unmarshal interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("get-request building: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.do(req)
	if err != nil {
		return fmt.Errorf("get-request doing: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println("get-request response body close error:", err.Error())
		}
	}()

	err = json.NewDecoder(resp.Body).Decode(unmarshal)
	if err != nil {
		return fmt.Errorf("get-response decode: %w", err)
	}

	return nil
}

// do - doing request directly
func (c *defaultHTTP) do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail client do: %w", err)
	}

	return resp, nil
}
