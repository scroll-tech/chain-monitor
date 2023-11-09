package contracts

import (
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2scrollmessenger"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type l2Contracts struct {
	client *ethclient.Client

	Messenger *il2scrollmessenger.Il2scrollmessenger

	ERC20Gateways      map[types.ERC20]*il2erc20gateway.Il2erc20gateway
	ERC20GatewayTokens []ERC20GatewayMapping
}

func newL2Contracts(c *ethclient.Client) *l2Contracts {
	return &l2Contracts{
		client:        c,
		ERC20Gateways: make(map[types.ERC20]*il2erc20gateway.Il2erc20gateway),
	}
}

func (l *l2Contracts) register(conf config.Config) error {
	var err error
	l.Messenger, err = il2scrollmessenger.NewIl2scrollmessenger(conf.L2Config.L2Contracts.ScrollMessenger, l.client)
	if err != nil {
		log.Error("registerERC20Gateway failed", "address", conf.L2Config.L2Contracts.ScrollMessenger, "err", err)
		return fmt.Errorf("register l2 scroll messenger contract failed, address:%v, err:%w", conf.L2Config.L2Contracts.ScrollMessenger.Hex(), err)
	}

	gateways := []struct {
		Address common.Address
		Token   types.ERC20
	}{
		{conf.L2Config.L2Contracts.WETHGateway, types.WETH},
		{conf.L2Config.L2Contracts.StandardERC20Gateway, types.StandardERC20},
		{conf.L2Config.L2Contracts.CustomERC20Gateway, types.CustomERC20},
		{conf.L2Config.L2Contracts.DAIGateway, types.DAI},
		{conf.L2Config.L2Contracts.USDCGateway, types.USDC},
		{conf.L2Config.L2Contracts.LIDOGateway, types.LIDO},
	}

	for _, gw := range gateways {
		if err := l.registerERC20Gateway(gw.Address, gw.Token); err != nil {
			log.Error("registerERC20Gateway failed", "address", gw.Address, "token", gw.Token, "err", err)
			return err
		}
	}
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
	l.ERC20GatewayTokens = append(l.ERC20GatewayTokens, ERC20GatewayMapping{TokenType: tokenType, Address: gatewayAddress})

	return nil
}
