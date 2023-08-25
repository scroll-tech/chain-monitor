package monitor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/scroll-tech/go-ethereum/common"
)

// Response the response schema
type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Data    string `json:"data"`
}

func (c *ChainMonitor) getWithdrawRoot(batchIndex uint64) (common.Hash, error) {
	var actualRoot common.Hash
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/withdraw_root?batch_index=%d", c.bridgeHistoryURL, batchIndex), nil)
	if err != nil {
		return actualRoot, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return actualRoot, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return actualRoot, err
	}

	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return actualRoot, err
	}
	if response.ErrMsg != "" {
		return actualRoot, errors.New(response.ErrMsg)
	}
	actualRoot = common.HexToHash(response.Data)
	return actualRoot, nil
}

func (c *ChainMonitor) confirmWithdrawRoot(batchIndex uint64, endNumber uint64) (bool, error) {
	expectRoot, err := c.l2API.WithdrawRoot(context.Background(), endNumber)
	if err != nil {
		return false, err
	}

	actualRoot, err := c.getWithdrawRoot(batchIndex)
	if err != nil {
		return false, err
	}

	return expectRoot == actualRoot, nil
}
