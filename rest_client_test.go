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

// TestOrders - testing Orders method
func TestOrders(t *testing.T) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env variables:", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token := os.Getenv("API_TOKEN")

	// client for getting real sanbox account id
	sandboxClient := NewSandboxRestClient(token)
	realSandboxAccount, err := sandboxClient.Register(ctx, AccountTinkoff)
	if err != nil {
		fmt.Println("error getting real sandbox account:", err.Error())
	}

	tests := []struct {
		name, token, accountID string
		expectedOrders         []Order
		isExpectedError        bool
	}{
		{"Returns empty order list by default if account id is empty", token, DefaultAccount, []Order{}, false},
		{"Returns empty order list by default if account id is right", token, realSandboxAccount.ID, []Order{}, false},
		{"Returns error if token is wrong", "", realSandboxAccount.ID, []Order{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ordersTestingClient := NewSandboxRestClient(tt.token)
			orders, err := ordersTestingClient.Orders(ctx, tt.accountID)

			if tt.isExpectedError {
				assert.NotEqual(t, nil, err)
			} else {
				assert.Equal(t, nil, err)
				assert.Equal(t, tt.expectedOrders, orders)
			}

		})
	}
}
