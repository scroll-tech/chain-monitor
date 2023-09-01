package config

import (
	"encoding/json"
	"os"

	"github.com/scroll-tech/go-ethereum/common"
)

// Gateway address list.
type Gateway struct {
	MessageQueue         common.Address `json:"message_queue"`
	ScrollMessenger      common.Address `json:"scroll_messenger"`
	ETHGateway           common.Address `json:"eth_gateway"`
	DAIGateway           common.Address `json:"dai_gateway"`
	WETHGateway          common.Address `json:"weth_gateway"`
	StandardERC20Gateway common.Address `json:"standard_erc20_gateway"`
	CustomERC20Gateway   common.Address `json:"custom_erc20_gateway"`
	ERC721Gateway        common.Address `json:"erc721_gateway"`
	ERC1155Gateway       common.Address `json:"erc1155_gateway"`
}

type L1Contracts struct {
	*Gateway
	L1ScrollChain common.Address `json:"scroll_chain"`
}

// L1Config l1 chain config.
type L1Config struct {
	L1Gateways  *L1Contracts `json:"l1_gateways"`
	L1ChainURL  string       `json:"l1chain_url"`
	Confirm     uint64
	StartNumber uint64 `json:"start_number"`
}

// L2Config l1 chain config.
type L2Config struct {
	L2gateways *Gateway `json:"l2_gateways"`
	L2ChainURL string   `json:"l2chain_url"`
	Confirm    uint64
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

type Config struct {
	L1Config *L1Config `json:"l1_config"`
	L2Config *L2Config `json:"l2_config"`
	DBConfig *DBConfig `json:"db_config"`
}

func NewConfig(file string) (*Config, error) {
	data, err := os.ReadFile(file)
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
