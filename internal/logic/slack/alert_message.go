package slack

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

var (
	withdrawRootNotMatchTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_withdraw_root_not_match_total",
		Help: "The total number of alert withdraw root not match total.",
	})

	gatewayTransferEventNotMatchTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_gateway_transfer_not_match_total",
		Help: "The total number of alert gateway transfer event not match total.",
	})

	crossChainGatewayEventNotMatchTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_cross_chain_gateway_event_not_match_total",
		Help: "The total number of alert cross chain gateway event not match total.",
	})

	crossChainETHEventNotMatchTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_cross_chain_eth_event_not_match_total",
		Help: "The total number of alert cross chain eth event not match total.",
	})

	crossChainETHEventBalanceNotMatchTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_cross_chain_eth_event_balance_not_match_total",
		Help: "The total number of alert cross chain eth event balance not match total.",
	})

	gatewayEventDuplicatedTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_gateway_event_duplicated_total",
		Help: "The total number of alert gateway event duplicated.",
	})

	messengerEventDuplicatedTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_messenger_event_duplicated_total",
		Help: "The total number of alert messenger event duplicated.",
	})

	nodeSyncHeightDiffAlertTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_node_sync_height_diff_total",
		Help: "The total number of node sync height difference alerts.",
	})

	nodeSyncHashMismatchAlertTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "slack_alert_node_sync_hash_mismatch_total",
		Help: "The total number of node sync hash mismatch alerts.",
	})
)

// GatewayTransferInfo the alert message of gateway and transfer event
type GatewayTransferInfo struct {
	TokenAddress    common.Address
	TokenType       types.TokenType
	Layer           types.LayerType
	EventType       types.EventType
	BlockNumber     uint64
	TxHash          common.Hash
	MessageHash     common.Hash
	Error           string
	TransferBalance *big.Int
	GatewayBalance  *big.Int
	TokenIgnored    bool
	CoinGeckoStatus string
}

// WithdrawRootInfo the alert message of withdraw root info
type WithdrawRootInfo struct {
	BlockNumber          uint64
	LastWithdrawRoot     common.Hash
	ExpectedWithdrawRoot common.Hash
}

// NodeSyncHeightDiffInfo the alert message of node sync height difference
type NodeSyncHeightDiffInfo struct {
	RethHeight uint64
	GethHeight uint64
	Difference uint64
	Threshold  uint64
}

// NodeSyncHashMismatchInfo the alert message of node sync hash mismatch
type NodeSyncHashMismatchInfo struct {
	Height   uint64
	RethHash common.Hash
	GethHash common.Hash
}

// MrkDwnWithdrawRootMessage make the markdown message of withdraw root alert message
func MrkDwnWithdrawRootMessage(info WithdrawRootInfo) string {
	withdrawRootNotMatchTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:bangbang: ")
	buffer.WriteString("*L2 withdraw root check failed*\n")
	buffer.WriteString(fmt.Sprintf("• block number: %d\n", info.BlockNumber))
	buffer.WriteString(fmt.Sprintf("• got withdraw root: %s\n", info.LastWithdrawRoot.Hex()))
	buffer.WriteString(fmt.Sprintf("• excepted withdraw root: %s\n", info.ExpectedWithdrawRoot.Hex()))
	return buffer.String()
}

// MrkDwnGatewayTransferMessage make the markdown message of gateway and transfer alert message
func MrkDwnGatewayTransferMessage(info GatewayTransferInfo) string {
	gatewayTransferEventNotMatchTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:bangbang: ")
	buffer.WriteString("*Gateway event and transfer event check failed*\n")
	buffer.WriteString(fmt.Sprintf("• token type: %s\n", info.TokenType.String()))
	buffer.WriteString(fmt.Sprintf("• layer type: %s\n", info.Layer.String()))
	buffer.WriteString(fmt.Sprintf("• event type: %s\n", info.EventType.String()))
	buffer.WriteString(fmt.Sprintf("• block number: %d\n", info.BlockNumber))
	buffer.WriteString(fmt.Sprintf("• tx_hash: %s\n", info.TxHash.Hex()))
	buffer.WriteString(fmt.Sprintf("• msg_hash: %s\n", info.MessageHash.Hex()))
	buffer.WriteString(fmt.Sprintf("• transfer balance: %s\n", info.TransferBalance.String()))
	buffer.WriteString(fmt.Sprintf("• gateway balance: %s\n", info.GatewayBalance.String()))
	buffer.WriteString(fmt.Sprintf("• token address: %s\n", info.TokenAddress.Hex()))
	buffer.WriteString(fmt.Sprintf("• token ignored: %t\n", info.TokenIgnored))
	buffer.WriteString(fmt.Sprintf("• err info:%s\n", info.Error))
	return buffer.String()
}

// MrkDwnGatewayCrossChainMessage make the markdown message of cross chain alert message
func MrkDwnGatewayCrossChainMessage(message orm.GatewayMessageMatch, checkResult types.MismatchType) string {
	crossChainGatewayEventNotMatchTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:bangbang: ")
	buffer.WriteString("*Cross chain gateway event check failed*\n")
	buffer.WriteString(fmt.Sprintf("• database id: %d\n", message.ID))
	buffer.WriteString(fmt.Sprintf("• token type: %s\n", types.TokenType(message.TokenType).String()))
	buffer.WriteString(fmt.Sprintf("• l1 event type: %s\n", types.EventType(message.L1EventType).String()))
	buffer.WriteString(fmt.Sprintf("• l2 event type: %s\n", types.EventType(message.L2EventType).String()))
	buffer.WriteString(fmt.Sprintf("• mismatch type: %s\n", checkResult.String()))
	buffer.WriteString(fmt.Sprintf("• l1 block number: %d\n", message.L1BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l2 block number: %d\n", message.L2BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l1 mount: %s\n", message.L1Amounts))
	buffer.WriteString(fmt.Sprintf("• l2 mount: %s\n", message.L2Amounts))
	buffer.WriteString(fmt.Sprintf("• l1 token: %s\n", message.L1TokenIds))
	buffer.WriteString(fmt.Sprintf("• l2 token: %s\n", message.L2TokenIds))
	buffer.WriteString(fmt.Sprintf("• l1 tx_hash: %s\n", message.L1TxHash))
	buffer.WriteString(fmt.Sprintf("• l2 tx_hash: %s\n", message.L2TxHash))
	buffer.WriteString(fmt.Sprintf("• msg_hash: %s\n", message.MessageHash))
	return buffer.String()
}

// MrkDwnETHCrossChainMessage make the markdown message of cross chain alert message
func MrkDwnETHCrossChainMessage(message orm.MessengerMessageMatch, checkResult types.MismatchType) string {
	crossChainETHEventNotMatchTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:bangbang: ")
	buffer.WriteString("*Cross chain messenger event check failed*\n")
	buffer.WriteString(fmt.Sprintf("• database id: %d\n", message.ID))
	buffer.WriteString(fmt.Sprintf("• l1 event type: %s\n", types.EventType(message.L1EventType).String()))
	buffer.WriteString(fmt.Sprintf("• l2 event type: %s\n", types.EventType(message.L2EventType).String()))
	buffer.WriteString(fmt.Sprintf("• mismatch type: %s\n", checkResult.String()))
	buffer.WriteString(fmt.Sprintf("• l1 block number: %d\n", message.L1BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l2 block number: %d\n", message.L2BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l1 tx_hash: %s\n", message.L1TxHash))
	buffer.WriteString(fmt.Sprintf("• l2 tx_hash: %s\n", message.L2TxHash))
	buffer.WriteString(fmt.Sprintf("• msg_hash: %s\n", message.MessageHash))
	return buffer.String()
}

// MrkDwnETHGatewayMessage make the markdown message of cross chain eth alert message
func MrkDwnETHGatewayMessage(message *orm.MessengerMessageMatch, expectedEndBalance, actualEndBalance *big.Int) string {
	crossChainETHEventBalanceNotMatchTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:bangbang: ")
	buffer.WriteString("*Cross chain ETH balance check failed*\n")
	buffer.WriteString(fmt.Sprintf("• database id: %d\n", message.ID))
	buffer.WriteString(fmt.Sprintf("• l1 event type: %s\n", types.EventType(message.L1EventType).String()))
	buffer.WriteString(fmt.Sprintf("• l2 event type: %s\n", types.EventType(message.L2EventType).String()))
	buffer.WriteString(fmt.Sprintf("• l1 block number: %d\n", message.L1BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l2 block number: %d\n", message.L2BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l1 tx_hash: %s\n", message.L1TxHash))
	buffer.WriteString(fmt.Sprintf("• l2 tx_hash: %s\n", message.L2TxHash))
	buffer.WriteString(fmt.Sprintf("• msg_hash: %s\n", message.MessageHash))
	buffer.WriteString(fmt.Sprintf("• expected end balance: %s\n", expectedEndBalance.String()))
	buffer.WriteString(fmt.Sprintf("• actual end balance: %s\n", actualEndBalance.String()))
	return buffer.String()
}

// MrkDwnGatewayMessageMatchDuplicated make the markdown message of duplicated gateway message
func MrkDwnGatewayMessageMatchDuplicated(layer types.LayerType, message orm.GatewayMessageMatch) string {
	gatewayEventDuplicatedTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:bangbang: ")
	buffer.WriteString("*Gateway event duplicated*\n")
	buffer.WriteString(fmt.Sprintf("• layer: %s\n", layer.String()))
	if layer == types.Layer1 {
		buffer.WriteString(fmt.Sprintf("• l1 event type: %s\n", types.EventType(message.L1EventType).String()))
		buffer.WriteString(fmt.Sprintf("• l1 block number: %d\n", message.L1BlockNumber))
		buffer.WriteString(fmt.Sprintf("• l1 tx_hash: %s\n", message.L1TxHash))
	} else {
		buffer.WriteString(fmt.Sprintf("• l2 event type: %s\n", types.EventType(message.L2EventType).String()))
		buffer.WriteString(fmt.Sprintf("• l2 block number: %d\n", message.L2BlockNumber))
		buffer.WriteString(fmt.Sprintf("• l2 tx_hash: %s\n", message.L2TxHash))
	}
	buffer.WriteString(fmt.Sprintf("• msg_hash: %s\n", message.MessageHash))
	return buffer.String()
}

// MrkDwnMessengerMessageMatchDuplicated make the markdown message of duplicated messenger message
func MrkDwnMessengerMessageMatchDuplicated(layer types.LayerType, message orm.MessengerMessageMatch) string {
	messengerEventDuplicatedTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:bangbang: ")
	buffer.WriteString("*Messenger event duplicated*\n")
	buffer.WriteString(fmt.Sprintf("• layer: %s\n", layer.String()))
	if layer == types.Layer1 {
		buffer.WriteString(fmt.Sprintf("• l1 event type: %s\n", types.EventType(message.L1EventType).String()))
		buffer.WriteString(fmt.Sprintf("• l1 block number: %d\n", message.L1BlockNumber))
		buffer.WriteString(fmt.Sprintf("• l1 tx_hash: %s\n", message.L1TxHash))
	} else {
		buffer.WriteString(fmt.Sprintf("• l2 event type: %s\n", types.EventType(message.L2EventType).String()))
		buffer.WriteString(fmt.Sprintf("• l2 block number: %d\n", message.L2BlockNumber))
		buffer.WriteString(fmt.Sprintf("• l2 tx_hash: %s\n", message.L2TxHash))
	}
	buffer.WriteString(fmt.Sprintf("• msg_hash: %s\n", message.MessageHash))
	return buffer.String()
}

// MrkDwnNodeSyncHeightDiffAlert make the markdown message of node sync height difference alert
func MrkDwnNodeSyncHeightDiffAlert(info NodeSyncHeightDiffInfo) string {
	nodeSyncHeightDiffAlertTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:warning: ")
	buffer.WriteString("*Node Height Difference Alert*\n")
	buffer.WriteString(fmt.Sprintf("• reth height: `%d`\n", info.RethHeight))
	buffer.WriteString(fmt.Sprintf("• geth height: `%d`\n", info.GethHeight))
	buffer.WriteString(fmt.Sprintf("• difference: `%d`\n", info.Difference))
	buffer.WriteString(fmt.Sprintf("• threshold: `%d`\n", info.Threshold))
	buffer.WriteString("• status: :x: *ALERTING*\n")
	return buffer.String()
}

// MrkDwnNodeSyncHeightDiffRecovered make the markdown message of node sync height difference recovery
func MrkDwnNodeSyncHeightDiffRecovered(info NodeSyncHeightDiffInfo) string {
	var buffer bytes.Buffer
	buffer.WriteString("\n:white_check_mark: ")
	buffer.WriteString("*Node Height Difference Recovered*\n")
	buffer.WriteString(fmt.Sprintf("• reth height: `%d`\n", info.RethHeight))
	buffer.WriteString(fmt.Sprintf("• geth height: `%d`\n", info.GethHeight))
	buffer.WriteString(fmt.Sprintf("• difference: `%d`\n", info.Difference))
	buffer.WriteString(fmt.Sprintf("• threshold: `%d`\n", info.Threshold))
	buffer.WriteString("• status: :white_check_mark: *NORMAL*\n")
	return buffer.String()
}

// MrkDwnNodeSyncHashMismatchAlert make the markdown message of node sync hash mismatch alert
func MrkDwnNodeSyncHashMismatchAlert(info NodeSyncHashMismatchInfo) string {
	nodeSyncHashMismatchAlertTotal.Inc()

	var buffer bytes.Buffer
	buffer.WriteString("\n:rotating_light: ")
	buffer.WriteString("*Block Hash Mismatch Alert*\n")
	buffer.WriteString(fmt.Sprintf("• block height: `%d`\n", info.Height))
	buffer.WriteString(fmt.Sprintf("• reth hash: `%s`\n", info.RethHash.Hex()))
	buffer.WriteString(fmt.Sprintf("• geth hash: `%s`\n", info.GethHash.Hex()))
	buffer.WriteString("• :warning: consensus status NOT updated\n")
	buffer.WriteString("• status: :x: *ALERTING*\n")
	return buffer.String()
}

// MrkDwnNodeSyncHashMismatchRecovered make the markdown message of node sync hash mismatch recovery
func MrkDwnNodeSyncHashMismatchRecovered(height uint64, blockHash common.Hash) string {
	var buffer bytes.Buffer
	buffer.WriteString("\n:white_check_mark: ")
	buffer.WriteString("*Block Hash Mismatch Recovered*\n")
	buffer.WriteString(fmt.Sprintf("• block height: `%d`\n", height))
	buffer.WriteString(fmt.Sprintf("• block hash: `%s`\n", blockHash.Hex()))
	buffer.WriteString("• status: :white_check_mark: *NORMAL*\n")
	return buffer.String()
}
