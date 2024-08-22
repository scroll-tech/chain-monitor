package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"

	"github.com/scroll-tech/chain-monitor/internal/utils/database"
)

// Gateway address list.
type Gateway struct {
	// erc20
	WETHGateway          common.Address `json:"weth_gateway"`
	StandardERC20Gateway common.Address `json:"standard_erc20_gateway"`
	CustomERC20Gateway   common.Address `json:"custom_erc20_gateway"`
	DAIGateway           common.Address `json:"dai_gateway"`
	USDCGateway          common.Address `json:"usdc_gateway"`
	LIDOGateway          common.Address `json:"lido_gateway"`
	PufferGateway        common.Address `json:"puffer_gateway"`
	GasTokenGateway      common.Address `json:"gas_token_gateway"`

	// erc721
	ERC721Gateway common.Address `json:"erc721_gateway"`

	// erc1155
	ERC1155Gateway common.Address `json:"erc1155_gateway"`
}

// L1Contracts l1chain config.
type L1Contracts struct {
	Gateway         `json:"l1_gateways"`
	ScrollMessenger common.Address `json:"scroll_messenger"`
	GasToken        common.Address `json:"gas_token"`
}

// L1Config l1 chain config.
type L1Config struct {
	L1URL                 string `json:"l1_url"`
	Confirm               rpc.BlockNumber
	L1Contracts           *L1Contracts     `json:"l1_contracts"`
	StartNumber           uint64           `json:"start_number"`
	StartMessengerBalance uint64           `json:"start_messenger_balance"`
	IgnoredTokens         []common.Address `json:"ignored_tokens"`
}

// L2Contracts l1chain config.
type L2Contracts struct {
	Gateway         `json:"l2_gateways"`
	ScrollMessenger common.Address `json:"scroll_messenger"`
	MessageQueue    common.Address `json:"message_queue"`
}

// L2Config l1 chain config.
type L2Config struct {
	L2URL         string `json:"l2_url"`
	Confirm       rpc.BlockNumber
	L2Contracts   *L2Contracts     `json:"l2_contracts"`
	IgnoredTokens []common.Address `json:"ignored_tokens"`
}

// SlackWebhookConfig slack webhook config.
type SlackWebhookConfig struct {
	WebhookURL       string `json:"webhook_url,omitempty"`
	WorkerCount      int    `json:"worker_count"`
	WorkerBufferSize int    `json:"worker_buffer_size"`
}

// Config chain-monitor main config.
type Config struct {
	L1Config    *L1Config           `json:"l1_config"`
	L2Config    *L2Config           `json:"l2_config"`
	AlertConfig *SlackWebhookConfig `json:"slack_webhook_config"`
	DBConfig    *database.Config    `json:"db_config"`
}

// NewConfig return an unmarshalled config instance.
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

	// Override config with environment variables
	err = overrideConfigWithEnv(&cfg, "SCROLL_CHAIN_MONITOR")
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// overrideConfigWithEnv recursively overrides config values with environment variables
func overrideConfigWithEnv(cfg interface{}, prefix string) error {
	v := reflect.ValueOf(cfg)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil
	}
	v = v.Elem()

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if !fieldValue.CanSet() {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}

		envKey := prefix + "_" + strings.ToUpper(tag)

		switch fieldValue.Kind() {
		case reflect.Ptr:
			if !fieldValue.IsNil() {
				err := overrideConfigWithEnv(fieldValue.Interface(), envKey)
				if err != nil {
					return err
				}
			}
		case reflect.Struct:
			err := overrideConfigWithEnv(fieldValue.Addr().Interface(), envKey)
			if err != nil {
				return err
			}
		default:
			if envValue, exists := os.LookupEnv(envKey); exists {
				err := setField(fieldValue, envValue)
				log.Info("Overriding config with env var", "key", envKey)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// setField sets the value of a field based on the environment variable value
func setField(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(uintValue)
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
	default:
		return fmt.Errorf("unsupported type: %v", field.Kind())
	}
	return nil
}
