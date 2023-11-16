package types

// EventType represents the type of blockchain event.
type EventType uint8

const (
	// EventTypeUnknown represents an unknown or undefined event.
	EventTypeUnknown EventType = iota

	// L1SentMessage represents a message sent on Layer 1.
	L1SentMessage
	// L1RelayedMessage represents a message relayed on Layer 1.
	L1RelayedMessage

	// L2SentMessage represents a message sent on Layer 2.
	L2SentMessage
	// L2RelayedMessage represents a message relayed on Layer 2.
	L2RelayedMessage

	// L1DepositETH represents the event for depositing ETH on Layer 1.
	L1DepositETH
	// L1FinalizeWithdrawETH represents the event for finalizing ETH withdrawal on Layer 1.
	L1FinalizeWithdrawETH
	// L1RefundETH represents the event for refunding ETH on Layer 1.
	L1RefundETH

	// L2FinalizeDepositETH represents the event for finalizing ETH deposit on Layer 2.
	L2FinalizeDepositETH
	// L2WithdrawETH represents the event for withdrawing ETH on Layer 2.
	L2WithdrawETH

	// L1DepositERC20 represents the event for depositing ERC20 tokens (e.g., WETH, standard erc20, custom erc20, usdc, dai, lido, etc) on Layer 1.
	L1DepositERC20
	// L1FinalizeWithdrawERC20 represents the event for finalizing ERC20 tokens withdrawal on Layer 1.
	L1FinalizeWithdrawERC20
	// L1RefundERC20 represents the event for refunding ERC20 tokens on Layer 1.
	L1RefundERC20

	// L2FinalizeDepositERC20 represents the event for finalizing ERC20 tokens (e.g., WETH, standard erc20, custom erc20, usdc, dai, lido, etc) deposit on Layer 2.
	L2FinalizeDepositERC20
	// L2WithdrawERC20 represents the event for withdrawing ERC20 tokens on Layer 2.
	L2WithdrawERC20

	// L1DepositERC721 represents the event for depositing ERC721 tokens on Layer 1.
	L1DepositERC721
	// L1FinalizeWithdrawERC721 represents the event for finalizing ERC721 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC721
	// L1RefundERC721 represents the event for refunding ERC721 tokens on Layer 1.
	L1RefundERC721

	// L2FinalizeDepositERC721 represents the event for finalizing ERC721 token deposit on Layer 2.
	L2FinalizeDepositERC721
	// L2WithdrawERC721 represents the event for withdrawing ERC721 tokens on Layer 2.
	L2WithdrawERC721

	// L1DepositERC1155 represents the event for depositing ERC1155 tokens on Layer 1.
	L1DepositERC1155
	// L1FinalizeWithdrawERC1155 represents the event for finalizing ERC1155 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC1155
	// L1RefundERC1155 represents the event for refunding ERC1155 tokens on Layer 1.
	L1RefundERC1155

	// L2FinalizeDepositERC1155 represents the event for finalizing ERC1155 token deposit on Layer 2.
	L2FinalizeDepositERC1155
	// L2WithdrawERC1155 represents the event for withdrawing ERC1155 tokens on Layer 2.
	L2WithdrawERC1155

	// L1BatchDepositERC721 represents the event for batch depositing ERC721 tokens on Layer 1.
	L1BatchDepositERC721
	// L1FinalizeBatchWithdrawERC721 represents the event for batch finalizing ERC721 token withdrawals on Layer 1.
	L1FinalizeBatchWithdrawERC721
	// L1BatchRefundERC721 represents the event for batch refunding ERC721 tokens on Layer 1.
	L1BatchRefundERC721

	// L2FinalizeBatchDepositERC721 represents the event for batch finalizing ERC721 token deposits on Layer 2.
	L2FinalizeBatchDepositERC721
	// L2BatchWithdrawERC721 represents the event for batch withdrawing ERC721 tokens on Layer 2.
	L2BatchWithdrawERC721

	// L1BatchDepositERC1155 represents the event for batch depositing ERC1155 tokens on Layer 1.
	L1BatchDepositERC1155
	// L1FinalizeBatchWithdrawERC1155 represents the event for batch finalizing ERC1155 token withdrawals on Layer 1.
	L1FinalizeBatchWithdrawERC1155
	// L1BatchRefundERC1155 represents the event for batch refunding ERC1155 tokens on Layer 1.
	L1BatchRefundERC1155

	// L2FinalizeBatchDepositERC1155 represents the event for batch finalizing ERC1155 token deposits on Layer 2.
	L2FinalizeBatchDepositERC1155
	// L2BatchWithdrawERC1155 represents the event for batch withdrawing ERC1155 tokens on Layer 2.
	L2BatchWithdrawERC1155
)

// EventCategory represents the category of events.
type EventCategory int

const (
	// EventCategoryUnknown represents an unknown or undefined event category.
	EventCategoryUnknown EventCategory = iota
	// ERC20EventCategory represents the ERC20 gateway events.
	ERC20EventCategory
	// ERC721EventCategory represents the ERC721 gateway events.
	ERC721EventCategory
	// ERC1155EventCategory represents the ERC1155 gateway events.
	ERC1155EventCategory
	// MessengerEventCategory represents the messenger events.
	MessengerEventCategory
)

// ERC20 represents the type of ERC20 token.
type ERC20 int

const (
	// ERC20Unknown represents an unknown or undefined ERC20 token.
	ERC20Unknown ERC20 = iota
	// WETH represents Wrapped Ethereum (WETH) token.
	WETH
	// StandardERC20 represents a standard ERC20 token.
	StandardERC20
	// CustomERC20 represents a custom ERC20 token.
	CustomERC20
	// USDC represents USD Coin (USDC) token.
	USDC
	// DAI represents DAI stablecoin token.
	DAI
	// LIDO represents Lido DAO token.
	LIDO
)
