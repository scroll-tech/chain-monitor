package assembler

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestTransferEventMatcher_checkTokenInCoinGecko_Debug(t *testing.T) {
	matcher := &TransferEventMatcher{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// Test with a random address that definitely doesn't exist
	testTokenAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")

	// Let's check what actually gets returned
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/ethereum/contract/%s", testTokenAddr.Hex())

	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "chain-monitor/1.0")

	resp, err := matcher.httpClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	t.Logf("Status Code: %d", resp.StatusCode)
	t.Logf("Response Body: %s", string(body))

	// Also test with a known existing token
	existingTokenAddr := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") // USDT
	url2 := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/ethereum/contract/%s", existingTokenAddr.Hex())

	req2, err := http.NewRequest("GET", url2, nil)
	require.NoError(t, err)

	req2.Header.Set("Accept", "application/json")
	req2.Header.Set("User-Agent", "chain-monitor/1.0")

	resp2, err := matcher.httpClient.Do(req2)
	require.NoError(t, err)
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	require.NoError(t, err)

	t.Logf("Existing Token Status Code: %d", resp2.StatusCode)
	t.Logf("Existing Token Response Body: %s", string(body2))
}
