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

	ERC20Gateways        map[types.ERC20]*il1erc20gateway.Il1erc20gateway
	ERC20Transfer        map[types.ERC20]*iscrollerc20.Iscrollerc20
	ERC20CategoryTokens  map[types.TxEventCategory][]types.ERC20
	ERC20TokenEventTypes map[types.ERC20]map[string]types.EventType
}

func newL1Contracts(c *ethclient.Client) *l1Contracts {
	l1c := &l1Contracts{
		client:              c,
		ERC20CategoryTokens: make(map[types.TxEventCategory][]types.ERC20),
		ERC20Gateways:       make(map[types.ERC20]*il1erc20gateway.Il1erc20gateway),
		ERC20Transfer:       make(map[types.ERC20]*iscrollerc20.Iscrollerc20),
	}

	wethFilterEventType := make(map[string]types.EventType)
	wethFilterEventType["FilterDepositERC20"] = types.L1DepositWETH
	wethFilterEventType["FilterFinalizeWithdrawERC20"] = types.L1FinalizeWithdrawWETH
	wethFilterEventType["FilterRefundERC20"] = types.L1RefundWETH
	l1c.ERC20TokenEventTypes = make(map[types.ERC20]map[string]types.EventType)
	l1c.ERC20TokenEventTypes[types.WETH] = wethFilterEventType

	return l1c
}

func (l *l1Contracts) register(conf config.Config) error {
	if conf.L1Config.L1Contracts.WETHGateway != (common.Address{}) {
		wethGateway, err := il1erc20gateway.NewIl1erc20gateway(conf.L1Config.L1Contracts.WETHGateway, l.client)
		if err != nil {
			log.Error("register erc20 gateway contract failed", "address", conf.L1Config.L1Contracts.WETHGateway, "error", err)
			return fmt.Errorf("register erc20 gateway contract failed, err:%w", err)
		}

		l.ERC20Gateways[types.WETH] = wethGateway
		l.ERC20Transfer[types.WETH], err = iscrollerc20.NewIscrollerc20(conf.L1Config.L1Contracts.WETHGateway, l.client)
		if err != nil {
			log.Error("register erc20 transfer contract failed", "address", conf.L1Config.L1Contracts.WETHGateway, "error", err)
			return fmt.Errorf("register erc20 transfer contract failed, err:%w", err)
		}
		l.ERC20CategoryTokens[types.ERC20EventCategory] = append(l.ERC20CategoryTokens[types.ERC20EventCategory], types.WETH)
	}

	// TODO add others
	return nil
}
