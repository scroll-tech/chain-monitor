package logic

import (
	"context"

	"chain-monitor/orm"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L1"
	l1gateway "chain-monitor/bytecode/scroll/L1/gateway"
	"chain-monitor/bytecode/scroll/L1/rollup"
	"chain-monitor/internal/config"
)

type L1Contracts struct {
	db     *gorm.DB
	tx     *gorm.DB
	client *ethclient.Client

	ETHGateway           *l1gateway.L1ETHGateway
	WETHGateway          *l1gateway.L1WETHGateway
	DAIGateway           *l1gateway.L1DAIGateway
	StandardERC20Gateway *l1gateway.L1StandardERC20Gateway
	CustomERC20Gateway   *l1gateway.L1CustomERC20Gateway
	ERC721Gateway        *l1gateway.L1ERC721Gateway
	ERC1155Gateway       *l1gateway.L1ERC1155Gateway

	ScrollChain     *rollup.ScrollChain
	ScrollMessenger *L1.L1ScrollMessenger

	filter *bytecode.ContractsFilter
}

func NewL1Contracts(client *ethclient.Client, db *gorm.DB, cfg *config.L1Contracts) (*L1Contracts, error) {
	var (
		cts = &L1Contracts{
			client: client,
			db:     db,
		}
		err error
	)
	cts.ScrollMessenger, err = L1.NewL1ScrollMessenger(cfg.ScrollMessenger, client)
	if err != nil {
		return nil, err
	}
	cts.ETHGateway, err = l1gateway.NewL1ETHGateway(cfg.ETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.WETHGateway, err = l1gateway.NewL1WETHGateway(cfg.WETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.StandardERC20Gateway, err = l1gateway.NewL1StandardERC20Gateway(cfg.StandardERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.CustomERC20Gateway, err = l1gateway.NewL1CustomERC20Gateway(cfg.CustomERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC721Gateway, err = l1gateway.NewL1ERC721Gateway(cfg.ERC721Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC1155Gateway, err = l1gateway.NewL1ERC1155Gateway(cfg.ERC1155Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.DAIGateway, err = l1gateway.NewL1DAIGateway(cfg.DAIGateway, client)
	if err != nil {
		return nil, err
	}
	cts.ScrollChain, err = rollup.NewScrollChain(cfg.L1ScrollChain, client)
	if err != nil {
		return nil, err
	}

	cts.filter = bytecode.NewContractsFilter([]bytecode.ContractAPI{
		cts.ScrollMessenger,

		cts.ETHGateway,
		cts.DAIGateway,
		cts.WETHGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,

		cts.ScrollChain,
	}...)

	cts.registerEventHandlers()

	return cts, nil
}

func (l1 *L1Contracts) ParseL1Events(ctx context.Context, start, end uint64) error {
	l1.tx = l1.db.Begin()
	count, err := l1.filter.ParseLogs(ctx, l1.client, start, end)
	if err != nil {
		l1.tx.Rollback()
		return err
	}

	err = l1.tx.Save(&orm.L1Block{
		Number:     end,
		EventCount: count,
	}).Error
	if err != nil {
		l1.tx.Rollback()
		return err
	}

	if err = l1.tx.Commit().Error; err != nil {
		l1.tx.Rollback()
		return err
	}
	return nil
}

func (l1 *L1Contracts) registerEventHandlers() {
	l1.ETHGateway.RegisterDepositETH(func(vLog *types.Log, data *l1gateway.L1ETHGatewayDepositETHEvent) error {
		return orm.SaveL1ETHEvent(l1.tx, orm.L1DepositETH, vLog, data.From, data.To, data.Amount)
	})
	l1.ETHGateway.RegisterFinalizeWithdrawETH(func(vLog *types.Log, data *l1gateway.L1ETHGatewayFinalizeWithdrawETHEvent) error {
		return orm.SaveL1ETHEvent(l1.tx, orm.L1FinalizeWithdrawETH, vLog, data.From, data.To, data.Amount)
	})
	l1.WETHGateway.RegisterDepositERC20(func(vLog *types.Log, data *l1gateway.L1WETHGatewayDepositERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1DepositWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l1.WETHGateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *l1gateway.L1WETHGatewayFinalizeWithdrawERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1FinalizeWithdrawWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l1.DAIGateway.RegisterDepositERC20(func(vLog *types.Log, data *l1gateway.L1DAIGatewayDepositERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1DepositDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l1.DAIGateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *l1gateway.L1DAIGatewayFinalizeWithdrawERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1FinalizeWithdrawDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l1.StandardERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *l1gateway.L1StandardERC20GatewayDepositERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1DepositStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l1.StandardERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *l1gateway.L1StandardERC20GatewayFinalizeWithdrawERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1FinalizeWithdrawStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l1.CustomERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *l1gateway.L1CustomERC20GatewayDepositERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1DepositCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})
	l1.CustomERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *l1gateway.L1CustomERC20GatewayFinalizeWithdrawERC20Event) error {
		return orm.SaveL1ETH20Event(l1.tx, orm.L1FinalizeWithdrawCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount)
	})

	l1.ERC721Gateway.RegisterDepositERC721(func(vLog *types.Log, data *l1gateway.L1ERC721GatewayDepositERC721Event) error {
		return orm.SaveL1ERC721Event(l1.tx, orm.L1DepositERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId)
	})
	l1.ERC721Gateway.RegisterFinalizeWithdrawERC721(func(vLog *types.Log, data *l1gateway.L1ERC721GatewayFinalizeWithdrawERC721Event) error {
		return orm.SaveL1ERC721Event(l1.tx, orm.L1FinalizeWithdrawERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId)
	})
	l1.ERC1155Gateway.RegisterDepositERC1155(func(vLog *types.Log, data *l1gateway.L1ERC1155GatewayDepositERC1155Event) error {
		return orm.SaveL1ERC1155Event(l1.tx, orm.L1DepositERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount)
	})
	l1.ERC1155Gateway.RegisterFinalizeWithdrawERC1155(func(vLog *types.Log, data *l1gateway.L1ERC1155GatewayFinalizeWithdrawERC1155Event) error {
		return orm.SaveL1ERC1155Event(l1.tx, orm.L1FinalizeWithdrawERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount)
	})

	l1.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerSentMessageEvent) error {
		return orm.SaveL1Messenger(l1.tx, orm.L1SentMessage, vLog, crypto.Keccak256Hash(data.Message))
	})
	l1.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerRelayedMessageEvent) error {
		return orm.SaveL1Messenger(l1.tx, orm.L1RelayedMessage, vLog, common.BytesToHash(data.MessageHash[:]))
	})
	l1.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerFailedRelayedMessageEvent) error {
		return orm.SaveL1Messenger(l1.tx, orm.L1FailedRelayedMessage, vLog, common.BytesToHash(data.MessageHash[:]))
	})

	l1.ScrollChain.RegisterCommitBatch(func(vLog *types.Log, data *rollup.ScrollChainCommitBatchEvent) error {
		l1Tx, _, err := l1.client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			return err
		}
		scrollBatch, err := bytecode.GetBatchRangeFromCalldataV2(l1.ScrollChain.ABI, l1Tx.Data())
		if err != nil {
			return err
		}
		return l1.tx.Create(&orm.ScrollChainEvent{
			BatchIndex:    data.BatchIndex.Uint64(),
			BatchHash:     common.BytesToHash(data.BatchHash[:]).String(),
			CommitNumber:  vLog.BlockNumber,
			CommitHash:    vLog.TxHash.String(),
			L2StartNumber: scrollBatch.L2StartNumber,
			L2EndNumber:   scrollBatch.L2EndNumber,
			Status:        orm.BatchCommit,
		}).Error
	})
	l1.ScrollChain.RegisterFinalizeBatch(func(vLog *types.Log, data *rollup.ScrollChainFinalizeBatchEvent) error {
		eventMsg := orm.ScrollChainEvent{
			BatchHash:      common.BytesToHash(data.BatchHash[:]).String(),
			FinalizeNumber: vLog.BlockNumber,
			FinalizeHash:   vLog.TxHash.String(),
			Status:         orm.BatchFinalize,
		}
		return l1.tx.Select("finalize_number", "finalize_hash", "status").Updates(&eventMsg).Error
	})
	l1.ScrollChain.RegisterRevertBatch(func(vLog *types.Log, data *rollup.ScrollChainRevertBatchEvent) error {
		eventMsg := orm.ScrollChainEvent{
			BatchHash:    common.BytesToHash(data.BatchHash[:]).String(),
			RevertNumber: vLog.BlockNumber,
			RevertHash:   vLog.TxHash.String(),
			Status:       orm.BatchRevert,
		}
		return l1.tx.Select("revert_number", "revert_hash", "status").Updates(&eventMsg).Error
	})
}
