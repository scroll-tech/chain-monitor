package controller

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	once sync.Once

	// L1BlockNumber Stores the Prometheus gauge for Layer 1 block number.
	L1BlockNumber prometheus.Gauge
	// L2BlockNumber Stores the Prometheus gauge for Layer 2 block number.
	L2BlockNumber prometheus.Gauge

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

	// BatchStatusTimeSecUsed Prometheus gauge for measuring time used for batch status.
	BatchStatusTimeSecUsed prometheus.Gauge
	// BlocksStatusTimeSecUsed Prometheus gauge for measuring time used for blocks status.
	BlocksStatusTimeSecUsed prometheus.Gauge
)

// InitMonitorMetrics returns a new monitor metrics instance.
func InitMonitorMetrics(reg prometheus.Registerer) {
	once.Do(func() {
		initMonitorMetrics(reg)
	})
}

func initMonitorMetrics(reg prometheus.Registerer) {
	factory := promauto.With(reg)
	L1BlockNumber = factory.NewGauge(prometheus.GaugeOpts{
		Name: "l1_block_number",
		Help: "L1watcher scanned latest block number.",
	})
	L2BlockNumber = factory.NewGauge(prometheus.GaugeOpts{
		Name: "l2_block_number",
		Help: "L2watcher scanned latest block number.",
	})
	ETHEventTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "eth_gateway_event_total",
		Help: "Four types of eth events on l1chain or l2chain.",
	}, []string{"name"})
	DAIEventTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "dai_gateway_event_total",
		Help: "Four types of dai events on l1chain or l2chain.",
	}, []string{"name"})
	WETHEventTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "weth_gateway_event_total",
		Help: "Four types of weth events on l1chain or l2chain.",
	}, []string{"name"})
	StandardERC20EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "standard_erc20_gateway_event_total",
		Help: "Four types of standard_erc20 events on l1chain or l2chain.",
	}, []string{"name"})
	CustomERC20EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "custom_erc20_gateway_event_total",
		Help: "Four types of custom_erc20 events on l1chain or l2chain.",
	}, []string{"name"})
	ERC721EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "erc721_gateway_event_total",
		Help: "Four types of erc721 events on l1chain or l2chain.",
	}, []string{"name"})
	ERC1155EventsTotal = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "erc1155_gateway_event_total",
		Help: "Four types of erc1155 events on l1chain or l2chain.",
	}, []string{"name"})
	BatchStatusTimeSecUsed = factory.NewGauge(prometheus.GaugeOpts{
		Name: "batch_status_api_time_sec_used",
		Help: "batch_status api time(sec) used",
	})
	BlocksStatusTimeSecUsed = factory.NewGauge(prometheus.GaugeOpts{
		Name: "blocks_status_api_time_sec_used",
		Help: "blocks api time(sec) used",
	})
}
