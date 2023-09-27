package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/rpc"
)

// Gateway address list.
type Gateway struct {
	ETHGateway           common.Address `json:"eth_gateway"`
	DAIGateway           common.Address `json:"dai_gateway"`
	WETHGateway          common.Address `json:"weth_gateway"`
	StandardERC20Gateway common.Address `json:"standard_erc20_gateway"`
	CustomERC20Gateway   common.Address `json:"custom_erc20_gateway"`
	ERC721Gateway        common.Address `json:"erc721_gateway"`
	ERC1155Gateway       common.Address `json:"erc1155_gateway"`
}

// L1Contracts l1chain config.
type L1Contracts struct {
	*Gateway        `json:"l1_gateways"`
	MessageQueue    common.Address `json:"message_queue"`
	ScrollMessenger common.Address `json:"scroll_messenger"`
	L1ScrollChain   common.Address `json:"scroll_chain"`
}

// L1Config l1 chain config.
type L1Config struct {
	L1Contracts *L1Contracts `json:"l1_contracts"`
	L1URL       string       `json:"l1_url"`
	Confirm     rpc.BlockNumber
	StartNumber uint64 `json:"start_number"`
}

// L2Contracts l1chain config.
type L2Contracts struct {
	*Gateway        `json:"l2_gateways"`
	MessageQueue    common.Address `json:"message_queue"`
	ScrollMessenger common.Address `json:"scroll_messenger"`
}

// L2Config l1 chain config.
type L2Config struct {
	L2Contracts *L2Contracts `json:"l2_contracts"`
	L2URL       string       `json:"l2_url"`
	Confirm     rpc.BlockNumber
	StartNumber uint64 `json:"start_number,omitempty"`
}

// DBConfig db config
type DBConfig struct {
	// data source name
	DSN        string `json:"dsn"`
	DriverName string `json:"driver_name"`

	MaxOpenNum int `json:"maxOpenNum"`
	MaxIdleNum int `json:"maxIdleNum"`
	LogLevel   int `json:"logLevel,omitempty"`
}

// SlackWebhookConfig slack webhook config.
type SlackWebhookConfig struct {
	Channel    string `json:"channel"`
	UserName   string `json:"user_name"`
	WebhookURL string `json:"webhook_url,omitempty"`
}

// Config chain-monitor main config.
type Config struct {
	L1Config    *L1Config           `json:"l1_config"`
	L2Config    *L2Config           `json:"l2_config"`
	AlertConfig *SlackWebhookConfig `json:"slack_webhook_config"`
	DBConfig    *DBConfig           `json:"db_config"`
}

// NewConfig return a unmarshalled config instance.
func NewConfig(file string) (*Config, error) {
	data, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
