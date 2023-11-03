package contracts

import (
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc20"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type l2Contracts struct {
	client *ethclient.Client

	ERC20Gateways       map[types.ERC20]*il2erc20gateway.Il2erc20gateway
	ERC20Transfer       map[types.ERC20]*iscrollerc20.Iscrollerc20
	ERC20CategoryTokens []types.ERC20
}

func newL2Contracts(c *ethclient.Client) *l2Contracts {
	return &l2Contracts{
		client:              c,
		ERC20Gateways:       make(map[types.ERC20]*il2erc20gateway.Il2erc20gateway),
		ERC20Transfer:       make(map[types.ERC20]*iscrollerc20.Iscrollerc20),
		ERC20CategoryTokens: []types.ERC20{},
	}
}

func (l *l2Contracts) register(conf config.Config) error {
	// @todo: handle error.
	l.registerERC20Gateway(conf.L2Config.L2Contracts.WETHGateway, types.WETH)
	l.registerERC20Gateway(conf.L2Config.L2Contracts.StandardERC20Gateway, types.StandardERC20)
	l.registerERC20Gateway(conf.L2Config.L2Contracts.CustomERC20Gateway, types.CustomERC20)
	l.registerERC20Gateway(conf.L2Config.L2Contracts.DAIGateway, types.DAI)
	l.registerERC20Gateway(conf.L2Config.L2Contracts.USDCGateway, types.USDC)
	l.registerERC20Gateway(conf.L2Config.L2Contracts.LIDOGateway, types.LIDO)

	// TODO add others
	return nil
}

func (l *l2Contracts) registerERC20Gateway(gatewayAddress common.Address, tokenType types.ERC20) error {
	if gatewayAddress == (common.Address{}) {
		log.Warn("gateway unconfigured", "address", gatewayAddress, "token type", tokenType)
		return nil
	}
	erc20Gateway, err := il2erc20gateway.NewIl2erc20gateway(gatewayAddress, l.client)
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
