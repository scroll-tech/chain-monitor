package controller

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// BlockNumber Stores the Prometheus gauge for Layer 1 or Layer 2 block number.
	BlockNumber *prometheus.GaugeVec

	// ETHEventTotal Prometheus counter vector for tracking total Ethereum events.
	ETHEventTotal *prometheus.CounterVec
	// DAIEventTotal Prometheus counter vector for tracking total DAI events.
	DAIEventTotal *prometheus.CounterVec
	// WETHEventTotal Prometheus counter vector for tracking total Wrapped ETH events.
	WETHEventTotal *prometheus.CounterVec
	// StandardERC20EventsTotal Prometheus counter vector for tracking total standard ERC-20 events.
	StandardERC20EventsTotal *prometheus.CounterVec
	// CustomERC20EventsTotal Prometheus counter vector for tracking total custom ERC-20 events.
	CustomERC20EventsTotal *prometheus.CounterVec
	// ERC721EventsTotal Prometheus counter vector for tracking total ERC-721 events.
	ERC721EventsTotal *prometheus.CounterVec
	// ERC1155EventsTotal Prometheus counter vector for tracking total ERC-1155 events.
	ERC1155EventsTotal *prometheus.CounterVec

	// ParseLogsFailureTotal The count of failed parse logs.
	ParseLogsFailureTotal *prometheus.CounterVec
	// ReorgTotal The count of reorg appears.
	ReorgTotal *prometheus.CounterVec

	// DepositFailedTotal The count of failed deposit events.
	DepositFailedTotal *prometheus.CounterVec
	// WithdrawFailedTotal The count of failed withdraw events.
	WithdrawFailedTotal *prometheus.CounterVec
	// WithdrawRootMisMatchTotal The count of mismatch withdraw root.
	WithdrawRootMisMatchTotal prometheus.Counter

	// WorkerStartedTotal The count of started goroutine workers.
	WorkerStartedTotal *prometheus.CounterVec
)

// InitMonitorMetrics returns a new monitor metrics instance.

func init() {
	factory := promauto.With(prometheus.DefaultRegisterer)
	chainName := "chain_name"
	eventName := "event_name"
	BlockNumber = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "l1_block_number",
		Help: "L1watcher scanned latest block number.",
	}, []string{chainName})

	ETHEventTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "eth_gateway_event_total",
		Help: "Four types of eth events on l1chain or l2chain.",
	}, []string{chainName, eventName})
	DAIEventTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "dai_gateway_event_total",
		Help: "Four types of dai events on l1chain or l2chain.",
	}, []string{chainName, eventName})
	WETHEventTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "weth_gateway_event_total",
		Help: "Four types of weth events on l1chain or l2chain.",
	}, []string{chainName, eventName})
	StandardERC20EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "standard_erc20_gateway_event_total",
		Help: "Four types of standard_erc20 events on l1chain or l2chain.",
	}, []string{chainName, eventName})
	CustomERC20EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "custom_erc20_gateway_event_total",
		Help: "Four types of custom_erc20 events on l1chain or l2chain.",
	}, []string{chainName, eventName})
	ERC721EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "erc721_gateway_event_total",
		Help: "Four types of erc721 events on l1chain or l2chain.",
	}, []string{chainName, eventName})
	ERC1155EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "erc1155_gateway_event_total",
		Help: "Four types of erc1155 events on l1chain or l2chain.",
	}, []string{chainName, eventName})

	ParseLogsFailureTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "failed_times_parse_logs_total",
		Help: "The count of failed parse logs",
	}, []string{chainName})
	ReorgTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "reorg_total",
		Help: "The count of reorg appears",
	}, []string{chainName})

	DepositFailedTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "failed_deposit_events_total",
		Help: "The count of failed deposit events.",
	}, []string{eventName})
	WithdrawFailedTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "failed_withdraw_events_total",
		Help: "The count of failed withdraw events.",
	}, []string{eventName})
	WithdrawRootMisMatchTotal = factory.NewCounter(prometheus.CounterOpts{
		Name: "mismatch_withdraw_root_total",
		Help: "The count of mismatch withdraw root.",
	})
	WorkerStartedTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "started_worker_total",
		Help: "The count of started goroutine workers.",
	}, []string{"worker_name"})
}
