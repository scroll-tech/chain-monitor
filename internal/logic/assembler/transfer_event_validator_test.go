package assembler

import (
	"net/http"
	"testing"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTransferEventMatcher_checkTokenInCoinGecko(t *testing.T) {
	// Create TransferEventMatcher with real HTTP client
	matcher := &TransferEventMatcher{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	tests := []struct {
		name        string
		tokenAddr   string
		shouldExist bool
	}{
		{
			name:        "USDT - should exist",
			tokenAddr:   "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			shouldExist: true,
		},
		{
			name:        "USDC - should exist",
			tokenAddr:   "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			shouldExist: true,
		},
		{
			name:        "WETH - should exist",
			tokenAddr:   "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			shouldExist: true,
		},
		{
			name:        "Random address - should not exist",
			tokenAddr:   "0x1111111111111111111111111111111111111111",
			shouldExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenAddr := common.HexToAddress(tt.tokenAddr)
			exists, err := matcher.checkTokenInCoinGecko(tokenAddr)

			assert.NoError(t, err)
			assert.Equal(t, tt.shouldExist, exists, "Token %s existence check failed: expected %v, got %v", tt.tokenAddr, tt.shouldExist, exists)

			// Add a small delay between requests to be respectful to the API
			time.Sleep(1000 * time.Millisecond)
		})
	}
}
