package contracts

import (
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc1155gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc721gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2scrollmessenger"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type l2Contracts struct {
	client *ethclient.Client

	messenger *il2scrollmessenger.Il2scrollmessenger

	erc20Gateways      map[types.ERC20]*il2erc20gateway.Il2erc20gateway
	erc20GatewayTokens []erc20GatewayMapping

	erc721Gateway         *il2erc721gateway.Il2erc721gateway
	erc721GatewayAddress  common.Address
	ERC1155Gateway        *il2erc1155gateway.Il2erc1155gateway
	ERC1155GatewayAddress common.Address
}

func newL2Contracts(c *ethclient.Client) *l2Contracts {
	return &l2Contracts{
		client:        c,
		erc20Gateways: make(map[types.ERC20]*il2erc20gateway.Il2erc20gateway),
	}
}

func (l *l2Contracts) register(conf *config.Config) error {
	var err error
	l.messenger, err = il2scrollmessenger.NewIl2scrollmessenger(conf.L2Config.L2Contracts.ScrollMessenger, l.client)
	if err != nil {
		log.Error("registerERC20Gateway failed", "address", conf.L2Config.L2Contracts.ScrollMessenger, "err", err)
		return fmt.Errorf("register l2 scroll messenger contract failed, address:%v, err:%w", conf.L2Config.L2Contracts.ScrollMessenger.Hex(), err)
	}

	erc20Gateways := []struct {
		Address common.Address
		Token   types.ERC20
	}{
		{conf.L2Config.L2Contracts.WETHGateway, types.WETH},
		{conf.L2Config.L2Contracts.StandardERC20Gateway, types.StandardERC20},
		{conf.L2Config.L2Contracts.CustomERC20Gateway, types.CustomERC20},
		{conf.L2Config.L2Contracts.DAIGateway, types.DAI},
		{conf.L2Config.L2Contracts.USDCGateway, types.USDC},
		{conf.L2Config.L2Contracts.LIDOGateway, types.LIDO},
		{conf.L2Config.L2Contracts.PufferGateway, types.PUFETH},
	}

	for _, gw := range erc20Gateways {
		if err := l.registerERC20Gateway(gw.Address, gw.Token); err != nil {
			log.Error("registerERC20Gateway failed", "address", gw.Address, "token", gw.Token, "err", err)
			return err
		}
	}

	erc721GatewayAddress := conf.L2Config.L2Contracts.ERC721Gateway
	if err := l.registerERC721Gateway(erc721GatewayAddress); err != nil {
		log.Error("registerERC721Gateway failed", "address", erc721GatewayAddress, "err", err)
		return err
	}

	erc1155GatewayAddress := conf.L2Config.L2Contracts.ERC1155Gateway
	if err := l.registerERC1155Gateway(erc1155GatewayAddress); err != nil {
		log.Error("registerERC1155Gateway failed", "address", erc1155GatewayAddress, "err", err)
		return err
	}

	return nil
}

func (l *l2Contracts) registerERC20Gateway(gatewayAddress common.Address, tokenType types.ERC20) error {
	if gatewayAddress == (common.Address{}) {
		log.Warn("l2 erc20 gateway unconfigured", "address", gatewayAddress, "token type", tokenType.String())
		return nil
	}
	erc20Gateway, err := il2erc20gateway.NewIl2erc20gateway(gatewayAddress, l.client)
	if err != nil {
		log.Error("l2 register erc20 gateway contract failed", "address", gatewayAddress, "error", err)
		return fmt.Errorf("register erc20 gateway contract failed, err:%w", err)
	}

	l.erc20Gateways[tokenType] = erc20Gateway
	l.erc20GatewayTokens = append(l.erc20GatewayTokens, erc20GatewayMapping{tokenType: tokenType, address: gatewayAddress})

	return nil
}

func (l *l2Contracts) registerERC721Gateway(gatewayAddress common.Address) error {
	if gatewayAddress == (common.Address{}) {
		log.Warn("l2 erc721 gateway unconfigured", "address", gatewayAddress)
		return nil
	}

	l.erc721GatewayAddress = gatewayAddress

	erc721Gateways, err := il2erc721gateway.NewIl2erc721gateway(gatewayAddress, l.client)
	if err != nil {
		return fmt.Errorf("l2 register erc721 gateway contract failed, err:%w", err)
	}
	l.erc721Gateway = erc721Gateways
	return nil
}

func (l *l2Contracts) registerERC1155Gateway(gatewayAddress common.Address) error {
	if gatewayAddress == (common.Address{}) {
		log.Warn("l2 erc1155 gateway unconfigured", "address", gatewayAddress)
		return nil
	}

	l.ERC1155GatewayAddress = gatewayAddress

	erc1155Gateways, err := il2erc1155gateway.NewIl2erc1155gateway(gatewayAddress, l.client)
	if err != nil {
		return fmt.Errorf("l2 register erc1155 gateway contract failed, err:%w", err)
	}
	l.ERC1155Gateway = erc1155Gateways
	return nil
}
