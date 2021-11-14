package cowswap

import (
	"context"
	"fmt"
	"testing"

	"github.com/gelfand/cowswap-go/types"
)

var client = NewClient("")

func TestClient_Fee(t *testing.T) {
	ctx := context.Background()
	_, err := client.Fee(ctx, types.FeeParams{
		SellToken: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		BuyToken:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Amount:    "100000000",
		Kind:      "buy",
	})
	if err != nil {
		t.Errorf("Client.Fee() error = %v, wantErr = false", err)
	}
}

func TestClient_Orders(t *testing.T) {
	ctx := context.Background()
	resp, err := client.Orders(ctx, types.OrdersParams{
		Owner:                      "0x7BB8536Bac7d6e5c4E352Ce50be8968d5d6cd445",
		IncludeFullyExecuted:       true,
		IncludeInvalidated:         true,
		IncludeInsufficientBalance: true,
		IncludePresignaturePending: true,
		IncludeUnsupportedTokens:   true,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)
}

func TestOrders(t *testing.T) {
	ctx := context.Background()
	resp, err := Orders(ctx, types.OrdersParams{
		Owner:                      "0x7BB8536Bac7d6e5c4E352Ce50be8968d5d6cd445",
		IncludeFullyExecuted:       true,
		IncludeInvalidated:         true,
		IncludeInsufficientBalance: true,
		IncludePresignaturePending: true,
		IncludeUnsupportedTokens:   true,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)
}
