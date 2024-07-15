package controller

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/assembler"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	messagematch "github.com/scroll-tech/chain-monitor/internal/logic/message_match"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

const maxBlockFetchSize uint64 = 49

// ContractController is a struct that manages the interaction with contracts on Layer 1 and Layer 2.
type ContractController struct {
	l1Client              *rpc.Client
	l2Client              *rpc.Client
	conf                  *config.Config
	eventGatherLogic      *events.EventGather
	contractsLogic        *contracts.Contracts
	messageMatchAssembler *assembler.MessageMatchAssembler
	messageMatchLogic     *messagematch.LogicMessageMatch

	stopContractChan    chan struct{}
	l1EventCategoryList []types.EventCategory
	l2EventCategoryList []types.EventCategory

	contractControllerRunningTotal                           *prometheus.CounterVec
	contractControllerBlockNumber                            *prometheus.GaugeVec
	contractControllerFilterGatewayIteratorFailureTotal      *prometheus.CounterVec
	contractControllerFilterTransferIteratorFailureTotal     *prometheus.CounterVec
	contractControllerGatewayCheckFailureTotal               *prometheus.CounterVec
	contractControllerUpdateOrInsertMessageMatchFailureTotal *prometheus.CounterVec
	contractControllerCheckWithdrawRootFailureTotal          *prometheus.CounterVec

	db                       *gorm.DB
	messengerMessageMatchOrm *orm.MessengerMessageMatch
	gatewayMessageMatchOrm   *orm.GatewayMessageMatch
}

// NewContractController creates a new ContractController object.
func NewContractController(conf *config.Config, db *gorm.DB, l1Client, l2Client *rpc.Client) *ContractController {
	c := &ContractController{
		l1Client:                 l1Client,
		l2Client:                 l2Client,
		conf:                     conf,
		eventGatherLogic:         events.NewEventGather(),
		contractsLogic:           contracts.NewContracts(ethclient.NewClient(l1Client), ethclient.NewClient(l2Client)),
		messageMatchAssembler:    assembler.NewMessageMatchAssembler(db),
		messageMatchLogic:        messagematch.NewMessageMatchLogic(conf, db),
		stopContractChan:         make(chan struct{}),
		db:                       db,
		messengerMessageMatchOrm: orm.NewMessengerMessageMatch(db),
		gatewayMessageMatchOrm:   orm.NewGatewayMessageMatch(db),
	}

	if err := c.contractsLogic.Register(c.conf); err != nil {
		log.Crit("contract register failure", "error", err)
		return nil
	}

	// eth balance is checked by other means.
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC20EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC721EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC1155EventCategory)

	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC20EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC721EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC1155EventCategory)

	reg := prometheus.DefaultRegisterer
	c.contractControllerRunningTotal = promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
		Name: "contract_controller_running_total",
		Help: "The total number of controllers running.",
	}, []string{"layer"})
	c.contractControllerBlockNumber = promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
		Name: "contract_controller_block_number",
		Help: "The block number of controller running.",
	}, []string{"layer"})
	c.contractControllerFilterGatewayIteratorFailureTotal = promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
		Name: "contract_controller_filter_gateway_iterator_failure_total",
		Help: "The total number of controller filter gateway iterator failure total.",
	}, []string{"layer", "event_category_type"})
	c.contractControllerFilterTransferIteratorFailureTotal = promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
		Name: "contract_controller_filter_transfer_iterator_failure_total",
		Help: "The total number of controller filter transfer iterator failure total.",
	}, []string{"layer", "event_category_type"})
	c.contractControllerGatewayCheckFailureTotal = promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
		Name: "contract_controller_gateway_check_failure_total",
		Help: "The total number of controller gateway check failure total.",
	}, []string{"layer"})
	c.contractControllerUpdateOrInsertMessageMatchFailureTotal = promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
		Name: "contract_controller_update_or_insert_failure_total",
		Help: "The total number of controller update or insert message match failure total.",
	}, []string{"layer"})
	c.contractControllerCheckWithdrawRootFailureTotal = promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
		Name: "contract_controller_check_l2_withdraw_root_failure_total",
		Help: "The total number of controller check l2 withdraw root failure total.",
	}, []string{"layer"})

	return c
}

// Watch is an exported function that starts watching the Layer 1 and Layer 2 events, which include gateways events, transfer events, and messenger events.
func (c *ContractController) Watch(ctx context.Context) {
	go c.watcherStart(ctx, ethclient.NewClient(c.l1Client), types.Layer1, c.conf.L1Config.Confirm, 2)
	go c.watcherStart(ctx, ethclient.NewClient(c.l2Client), types.Layer2, c.conf.L2Config.Confirm, 2)
}

// Stop the contract controller
func (c *ContractController) Stop() {
	c.stopContractChan <- struct{}{}
}

func (c *ContractController) watcherStart(ctx context.Context, client *ethclient.Client, layer types.LayerType, confirmation rpc.BlockNumber, concurrency int) {
	log.Info("contract controller start successful", "layer", layer.String(), "confirmation", confirmation)

	// 1. get the max l1_number and l2_number
	blockNumberInDB, getLastBlockErr := c.messageMatchLogic.GetLatestBlockNumber(ctx, layer)
	if getLastBlockErr != nil {
		log.Error("ContractController.Watch get latest block number failed", "layer", layer, "err", getLastBlockErr)
		return
	}
	log.Info("Block process height in db", "layer", layer, "block number", blockNumberInDB)
	start := blockNumberInDB + 1

	if layer == types.Layer2 {
		l2CurrentMaxBlockNumber.Store(blockNumberInDB)
	}

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Error("ContractController canceled with error", "error", ctx.Err())
			}
			return
		case <-c.stopContractChan:
			log.Info("ContractController l1 watch & l2 watch the run loop exit")
			return
		default:
		}

		c.contractControllerRunningTotal.WithLabelValues(layer.String()).Inc()

		// 2. get latest chain confirmation number
		confirmationNumber, latestConfirmedBlockErr := utils.GetLatestConfirmedBlockNumber(ctx, client, confirmation)
		if latestConfirmedBlockErr != nil {
			log.Error("ContractController.watcherStart get latest confirmation block number failed", "layer", layer.String(), "err", latestConfirmedBlockErr)
			time.Sleep(time.Second)
			continue
		}

		// three cases.
		// for example : concurrency = 3
		// case 1: confirmationNumber 500    start: 71
		//		g0:[71 - 120] g1:[121-170] g2:[171-220]
		// case 2: confirmationNumber 160    start: 71
		// 		g0:[71 - 120] g1:[121-160] g3: break
		// case 3: confirmationNumber 70，start: 71
		//		g0: break
		// case 4: confirmationNumber 1 start: 1
		//		g0 break
		var eg errgroup.Group
		loopStart := start
		loopEnd := loopStart - 1
		var mux sync.Mutex
		var gatewayMessageMatches []orm.GatewayMessageMatch
		var messengerMessageMatches []orm.MessengerMessageMatch
		for i := 0; i < concurrency; i++ {
			if loopStart > confirmationNumber {
				log.Info("Watcher loop start block number > ConfirmationNumber",
					"layer", layer.String(),
					"startBlockNumber", loopStart,
					"confirmationNumber", confirmationNumber,
				)
				time.Sleep(time.Second)
				break
			}

			// 3. get the max fetch number
			loopEnd = loopStart + maxBlockFetchSize
			if loopStart+maxBlockFetchSize > confirmationNumber {
				loopEnd = confirmationNumber
			}

			currentStart := loopStart
			currentEnd := loopEnd
			loopStart = loopEnd + 1

			c.contractControllerBlockNumber.WithLabelValues(layer.String()).Set(float64(currentEnd))

			eg.Go(func() error {
				var retGatewayMessageMatches []orm.GatewayMessageMatch
				var retMessengerMessageMatches []orm.MessengerMessageMatch
				var watchErr error
				switch layer {
				case types.Layer1:
					retGatewayMessageMatches, retMessengerMessageMatches, watchErr = c.l1Watch(ctx, currentStart, currentEnd)
					if watchErr != nil {
						return watchErr
					}
				case types.Layer2:
					retGatewayMessageMatches, retMessengerMessageMatches, watchErr = c.l2Watch(ctx, currentStart, currentEnd)
					if watchErr != nil {
						return watchErr
					}
				}
				mux.Lock()
				gatewayMessageMatches = append(gatewayMessageMatches, retGatewayMessageMatches...)
				messengerMessageMatches = append(messengerMessageMatches, retMessengerMessageMatches...)
				mux.Unlock()
				return nil
			})
		}

		if egErr := eg.Wait(); egErr != nil {
			log.Error("error in watcher goroutine", "layer", layer, "err", egErr)
			continue
		}

		if loopEnd >= start {
			var lastMessage *orm.MessengerMessageMatch
			if layer == types.Layer2 {
				var checkErr error
				lastMessage, checkErr = c.messageMatchAssembler.L2WithdrawRootsValidator(ctx, start, loopEnd, c.l2Client, c.conf.L2Config.L2Contracts.MessageQueue)
				if checkErr != nil {
					c.contractControllerCheckWithdrawRootFailureTotal.WithLabelValues(types.Layer2.String()).Inc()
					log.Error("check withdraw roots failed", "layer", types.Layer2, "start", start, "end", loopEnd, "error", checkErr)
					continue
				}
			}

			// Update last valid message's withdraw trie proof and block status after check.
			updateErr := c.db.Transaction(func(tx *gorm.DB) error {
				if layer == types.Layer2 {
					if updateMsgProofErr := c.messengerMessageMatchOrm.UpdateMsgProofAndStatus(ctx, lastMessage, tx); updateMsgProofErr != nil {
						return fmt.Errorf("insert or update msg proof and status failed, err: %w, message: %+v", updateMsgProofErr, lastMessage)
					}
				}

				if insertEventErr := c.messageMatchLogic.InsertOrUpdateMessageMatches(ctx, layer, gatewayMessageMatches, messengerMessageMatches); insertEventErr != nil {
					c.contractControllerUpdateOrInsertMessageMatchFailureTotal.WithLabelValues(layer.String()).Inc()
					log.Error("insert message events failed", "layer", layer.String(), "error", insertEventErr)
					return insertEventErr
				}
				return nil
			})
			if updateErr != nil {
				log.Error("update db status after check failed", "layer", layer, "from", start, "end", loopEnd, "err", updateErr)
				continue
			}

			if layer == types.Layer2 {
				l2CurrentMaxBlockNumber.Store(loopEnd)
			}
		}

		// Update start after all handlings are successful.
		start = loopEnd + 1
	}
}

func (c *ContractController) l1Watch(ctx context.Context, start uint64, end uint64) ([]orm.GatewayMessageMatch, []orm.MessengerMessageMatch, error) {
	log.Info("watching block number", "layer", types.Layer1, "start", start, "end", end)
	opts := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	messengerIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer1, types.MessengerEventCategory)
	if err != nil {
		c.contractControllerFilterGatewayIteratorFailureTotal.WithLabelValues(types.Layer1.String(), types.MessengerEventCategory.String()).Inc()
		log.Error("get messenger iterator failed", "layer", types.Layer1, "eventCategory", types.MessengerEventCategory, "error", err)
		return nil, nil, err
	}
	messengerEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer1, types.MessengerEventCategory, messengerIterList)
	messengerMessageMatches, err := c.messageMatchAssembler.MessageMatchAssembler(messengerEvents)
	if err != nil {
		log.Error("generate messenger message match failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
		return nil, nil, err
	}

	if len(messengerMessageMatches) == 0 {
		return nil, nil, nil
	}

	var l1GatewayMessageMatches []orm.GatewayMessageMatch
	for _, eventCategory := range c.l1EventCategoryList {
		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer1, eventCategory)
		if err != nil {
			c.contractControllerFilterGatewayIteratorFailureTotal.WithLabelValues(types.Layer1.String(), eventCategory.String()).Inc()
			log.Error("get contract iterator failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", err)
			return nil, nil, err
		}

		transferEvents, err := c.contractsLogic.GetGatewayTransfer(ctx, start, end, types.Layer1, eventCategory)
		if err != nil {
			c.contractControllerFilterTransferIteratorFailureTotal.WithLabelValues(types.Layer1.String(), "transfer").Inc()
			log.Error("get gateway related transfer events failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", err)
			return nil, nil, err
		}

		// parse the gateway and messenger event data
		gatewayEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer1, eventCategory, wrapIterList)
		if gatewayEvents == nil {
			log.Debug("event gather deal event return empty data", "layer", types.Layer1, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		retL1MessageMatches, checkErr := c.messageMatchAssembler.GatewayMessageAssembler(eventCategory, gatewayEvents, messengerEvents, transferEvents)
		l1GatewayMessageMatches = append(l1GatewayMessageMatches, retL1MessageMatches...)
		if checkErr != nil {
			c.contractControllerGatewayCheckFailureTotal.WithLabelValues(types.Layer1.String()).Inc()
			log.Error("event matcher deal failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", checkErr)
			return nil, nil, checkErr
		}
	}

	return l1GatewayMessageMatches, messengerMessageMatches, nil
}

func (c *ContractController) l2Watch(ctx context.Context, start uint64, end uint64) ([]orm.GatewayMessageMatch, []orm.MessengerMessageMatch, error) {
	log.Info("watching block number", "layer", types.Layer2, "start", start, "end", end)
	opts := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	messengerIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer2, types.MessengerEventCategory)
	if err != nil {
		c.contractControllerFilterGatewayIteratorFailureTotal.WithLabelValues(types.Layer2.String(), types.MessengerEventCategory.String()).Inc()
		log.Error("get messenger iterator failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
		return nil, nil, err
	}
	messengerEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, types.MessengerEventCategory, messengerIterList)
	messengerMessageMatches, err := c.messageMatchAssembler.MessageMatchAssembler(messengerEvents)
	if err != nil {
		log.Error("generate messenger message match failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
		return nil, nil, err
	}

	if len(messengerMessageMatches) == 0 {
		return nil, nil, nil
	}

	var l2GatewayMessageMatches []orm.GatewayMessageMatch
	for _, eventCategory := range c.l2EventCategoryList {
		var wrapIterList []types.WrapIterator
		wrapIterList, err = c.contractsLogic.Iterator(ctx, &opts, types.Layer2, eventCategory)
		if err != nil {
			c.contractControllerFilterGatewayIteratorFailureTotal.WithLabelValues(types.Layer2.String(), eventCategory.String()).Inc()
			log.Error("get contract iterator failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", err)
			return nil, nil, err
		}

		var transferEvents []events.EventUnmarshaler
		transferEvents, err = c.contractsLogic.GetGatewayTransfer(ctx, start, end, types.Layer2, eventCategory)
		if err != nil {
			c.contractControllerFilterTransferIteratorFailureTotal.WithLabelValues(types.Layer2.String(), "transfer").Inc()
			log.Error("get gateway related transfer events failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", err)
			return nil, nil, err
		}

		// parse the event data
		gatewayEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, eventCategory, wrapIterList)
		if gatewayEvents == nil {
			log.Debug("dispatch gateway events returns empty data", "layer", types.Layer2, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		retL2MessageMatches, checkErr := c.messageMatchAssembler.GatewayMessageAssembler(eventCategory, gatewayEvents, messengerEvents, transferEvents)
		l2GatewayMessageMatches = append(l2GatewayMessageMatches, retL2MessageMatches...)
		if checkErr != nil {
			c.contractControllerGatewayCheckFailureTotal.WithLabelValues(types.Layer2.String()).Inc()
			log.Error("event matcher deal failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", checkErr)
			return nil, nil, checkErr
		}
	}
	return l2GatewayMessageMatches, messengerMessageMatches, nil
}
