package sdk

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestSandboxRegister - register in sandox
func TestSandboxRegister(t *testing.T) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env variables:", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tests := []struct {
		name, token string
		isAllow     bool
	}{
		{"Does not allow registration with an empty token", "", false},
		{"Does not allow registration with wrong token", "this-is-wrong-token", false},
		{"Allow registration with right token", os.Getenv("API_TOKEN"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewSandboxRestClient(tt.token)
			account, err := client.Register(ctx, AccountTinkoff)

			if tt.isAllow {
				assert.Equal(t, nil, err)
				assert.Equal(t, AccountTinkoff, account.Type)
				assert.NotEmpty(t, account.ID)
			} else {
				assert.NotEqual(t, nil, err)
			}

		})
	}
}
