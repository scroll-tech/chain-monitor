package contracts

import (
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il1erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc20"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type l1Contracts struct {
	client *ethclient.Client

	ERC20Gateways       map[types.ERC20]*il1erc20gateway.Il1erc20gateway
	ERC20Transfer       map[types.ERC20]*iscrollerc20.Iscrollerc20
	ERC20CategoryTokens []types.ERC20
}

func newL1Contracts(c *ethclient.Client) *l1Contracts {
	return &l1Contracts{
		client:              c,
		ERC20Gateways:       make(map[types.ERC20]*il1erc20gateway.Il1erc20gateway),
		ERC20Transfer:       make(map[types.ERC20]*iscrollerc20.Iscrollerc20),
		ERC20CategoryTokens: []types.ERC20{},
	}
}

func (l *l1Contracts) register(conf config.Config) error {
	// @todo: handle error.
	l.registerERC20Gateway(conf.L1Config.L1Contracts.WETHGateway, types.WETH)
	l.registerERC20Gateway(conf.L1Config.L1Contracts.StandardERC20Gateway, types.StandardERC20)
	l.registerERC20Gateway(conf.L1Config.L1Contracts.CustomERC20Gateway, types.CustomERC20)
	l.registerERC20Gateway(conf.L1Config.L1Contracts.DAIGateway, types.DAI)
	l.registerERC20Gateway(conf.L1Config.L1Contracts.USDCGateway, types.USDC)
	l.registerERC20Gateway(conf.L1Config.L1Contracts.LIDOGateway, types.LIDO)

	// TODO add others
	return nil
}

func (l *l1Contracts) registerERC20Gateway(gatewayAddress common.Address, tokenType types.ERC20) error {
	if gatewayAddress == (common.Address{}) {
		log.Warn("gateway unconfigured", "address", gatewayAddress, "token type", tokenType)
		return nil
	}
	erc20Gateway, err := il1erc20gateway.NewIl1erc20gateway(gatewayAddress, l.client)
	if err != nil {
		log.Error("register erc20 gateway contract failed", "address", gatewayAddress, "error", err)
		return fmt.Errorf("register erc20 gateway contract failed, err:%w", err)
	}

	l.ERC20Gateways[tokenType] = erc20Gateway
	l.ERC20Transfer[tokenType], err = iscrollerc20.NewIscrollerc20(gatewayAddress, l.client)
	if err != nil {
		log.Error("register erc20 transfer contract failed", "address", gatewayAddress, "error", err)
		return fmt.Errorf("register erc20 transfer contract failed, err:%w", err)
	}
	l.ERC20CategoryTokens = append(l.ERC20CategoryTokens, tokenType)

	return nil
}
