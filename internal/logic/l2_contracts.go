package logic

import (
	"context"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/internal/config"
	"chain-monitor/orm"
)

type L2Contracts struct {
	db     *gorm.DB
	tx     *gorm.DB
	client *ethclient.Client

	ETHGateway           *gateway.L2ETHGateway
	WETHGateway          *gateway.L2WETHGateway
	DAIGateway           *gateway.L2DAIGateway
	StandardERC20Gateway *gateway.L2StandardERC20Gateway
	CustomERC20Gateway   *gateway.L2CustomERC20Gateway
	ERC721Gateway        *gateway.L2ERC721Gateway
	ERC1155Gateway       *gateway.L2ERC1155Gateway

	ScrollMessenger *L2.L2ScrollMessenger

	filter *bytecode.ContractsFilter
}

func NewL2Contracts(client *ethclient.Client, db *gorm.DB, cfg *config.Gateway) (*L2Contracts, error) {
	var (
		cts = &L2Contracts{
			client: client,
			db:     db,
		}
		err error
	)
	cts.ETHGateway, err = gateway.NewL2ETHGateway(cfg.ETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.WETHGateway, err = gateway.NewL2WETHGateway(cfg.WETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.DAIGateway, err = gateway.NewL2DAIGateway(cfg.DAIGateway, client)
	if err != nil {
		return nil, err
	}
	cts.StandardERC20Gateway, err = gateway.NewL2StandardERC20Gateway(cfg.StandardERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.CustomERC20Gateway, err = gateway.NewL2CustomERC20Gateway(cfg.CustomERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC721Gateway, err = gateway.NewL2ERC721Gateway(cfg.ERC721Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC1155Gateway, err = gateway.NewL2ERC1155Gateway(cfg.ERC1155Gateway, client)
	if err != nil {
		return nil, err
	}

	cts.ScrollMessenger, err = L2.NewL2ScrollMessenger(cfg.ScrollMessenger, client)
	if err != nil {
		return nil, err
	}

	cts.filter = bytecode.NewContractsFilter([]bytecode.ContractAPI{
		cts.ScrollMessenger,

		cts.ETHGateway,
		cts.WETHGateway,
		cts.DAIGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,
	}...)

	cts.registerDepositHandlers()

	return cts, nil
}

func (l2 *L2Contracts) ParseL2Events(ctx context.Context, start, end uint64) error {
	l2.tx = l2.db.Begin()
	count, err := l2.filter.ParseLogs(ctx, l2.client, start, end)
	if err != nil {
		l2.tx.Rollback()
		return err
	}

	err = l2.tx.Save(&orm.L2Block{
		Number:     end,
		EventCount: count,
	}).Error
	if err != nil {
		l2.tx.Rollback()
		return err
	}

	if err = l2.tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (l2 *L2Contracts) registerDepositHandlers() {
	l2.ETHGateway.RegisterWithdrawETH(func(vLog *types.Log, data *gateway.L2ETHGatewayWithdrawETHEvent) error {
		return orm.SaveL2ETHEvent(l2.tx, orm.L2WithdrawETH, vLog, data.From, data.To, data.Amount)
	})
	l2.ETHGateway.RegisterFinalizeDepositETH(func(vLog *types.Log, data *gateway.L2ETHGatewayFinalizeDepositETHEvent) error {
		return orm.SaveL2ETHEvent(l2.tx, orm.L2FinalizeDepositETH, vLog, data.From, data.To, data.Amount)
	})
	l2.WETHGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayWithdrawERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2WithdrawWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l2.WETHGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayFinalizeDepositERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2FinalizeDepositWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l2.DAIGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayWithdrawERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2WithdrawDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l2.DAIGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayFinalizeDepositERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2FinalizeDepositDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})

	l2.StandardERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayWithdrawERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2WithdrawStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l2.StandardERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayFinalizeDepositERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2FinalizeDepositStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})

	l2.CustomERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayWithdrawERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2WithdrawCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l2.CustomERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayFinalizeDepositERC20Event) error {
		return orm.SaveL2ETH20Event(l2.tx, orm.L2FinalizeDepositCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})

	l2.ERC721Gateway.RegisterWithdrawERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayWithdrawERC721Event) error {
		return orm.SaveL2ERC721Event(l2.tx, orm.L2WithdrawERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId)
	})
	l2.ERC721Gateway.RegisterFinalizeDepositERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayFinalizeDepositERC721Event) error {
		return orm.SaveL2ERC721Event(l2.tx, orm.L2FinalizeDepositERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId)
	})
	l2.ERC1155Gateway.RegisterWithdrawERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayWithdrawERC1155Event) error {
		return orm.SaveL2ERC1155Event(l2.tx, orm.L2WithdrawERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount)
	})
	l2.ERC1155Gateway.RegisterFinalizeDepositERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayFinalizeDepositERC1155Event) error {
		return orm.SaveL2ERC1155Event(l2.tx, orm.L2FinalizeDepositERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount)
	})

	l2.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerSentMessageEvent) error {
		return orm.SaveL2Messenger(l2.tx, orm.L2SentMessage, vLog, crypto.Keccak256Hash(data.Message))
	})

	l2.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerRelayedMessageEvent) error {
		return orm.SaveL2Messenger(l2.tx, orm.L2RelayedMessage, vLog, common.BytesToHash(data.MessageHash[:]))
	})

	l2.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerFailedRelayedMessageEvent) error {
		return orm.SaveL2Messenger(l2.tx, orm.L2FailedRelayedMessage, vLog, common.BytesToHash(data.MessageHash[:]))
	})
}
