package types

//go:generate stringer -type EventType

// EventType represents the type of blockchain event.
type EventType uint8

const (
	EventTypeUnknown EventType = iota

	L1SentMessage
	L1RelayedMessage

	L2SentMessage
	L2RelayedMessage

	// L1DepositETH represents the event for depositing ETH on Layer 1.
	L1DepositETH
	// L1FinalizeWithdrawETH represents the event for finalizing ETH withdrawal on Layer 1.
	L1FinalizeWithdrawETH
	// L1RefundETH the event for refund eth on Layer 1.
	L1RefundETH

	// L2FinalizeDepositETH represents the event for finalizing ETH deposit on Layer 2.
	L2FinalizeDepositETH
	// L2WithdrawETH represents the event for withdrawing ETH on Layer 2.
	L2WithdrawETH

	// L1DepositERC20 represents the event for depositing ERC20 tokens (e.g., WETH, standard erc20, custom erc20, usdc, dai, lido, etc) on Layer 1.
	L1DepositERC20
	// L1FinalizeWithdrawERC20 represents the event for finalizing ERC20 tokens withdrawal on Layer 1.
	L1FinalizeWithdrawERC20
	// L1RefundERC20 the event for refund erc20 on Layer 1.
	L1RefundERC20

	// L2FinalizeDepositERC20 represents the event for finalizing ERC20 tokens (e.g., WETH, standard erc20, custom erc20, usdc, dai, lido, etc) deposit on Layer 2.
	L2FinalizeDepositERC20
	// L2WithdrawERC20 represents the event for withdrawing ERC20 tokens on Layer 2.
	L2WithdrawERC20

	// L1DepositERC721 represents the event for depositing ERC721 tokens on Layer 1.
	L1DepositERC721
	// L1FinalizeWithdrawERC721 represents the event for finalizing ERC721 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC721
	// L1RefundERC721 the event for refund erc721 on Layer 1.
	L1RefundERC721

	// L2FinalizeDepositERC721 represents the event for finalizing ERC721 token deposit on Layer 2.
	L2FinalizeDepositERC721
	// L2WithdrawERC721 represents the event for withdrawing ERC721 tokens on Layer 2.
	L2WithdrawERC721

	// L1DepositERC1155 represents the event for depositing ERC1155 tokens on Layer 1.
	L1DepositERC1155
	// L1FinalizeWithdrawERC1155 represents the event for finalizing ERC1155 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC1155
	// L1RefundERC1155 the event for refund erc1155 on Layer 1.
	L1RefundERC1155

	// L2FinalizeDepositERC1155 represents the event for finalizing ERC1155 token deposit on Layer 2.
	L2FinalizeDepositERC1155
	// L2WithdrawERC1155 represents the event for withdrawing ERC1155 tokens on Layer 2.
	L2WithdrawERC1155

	// L1BatchDepositERC721 represents the event for batch depositing ERC721 tokens on Layer 1.
	L1BatchDepositERC721
	// L1FinalizeBatchWithdrawERC721 represents the event for batch finalizing ERC721 token withdrawals on Layer 1.
	L1FinalizeBatchWithdrawERC721
	// L1BatchRefundERC721 the event for batch refund erc721 on Layer 1.
	L1BatchRefundERC721

	// L2FinalizeBatchDepositERC721 represents the event for batch finalizing ERC721 token deposits on Layer 2.
	L2FinalizeBatchDepositERC721
	// L2BatchWithdrawERC721 represents the event for batch withdrawing ERC721 tokens on Layer 2.
	L2BatchWithdrawERC721

	// L1BatchDepositERC1155 represents the event for batch depositing ERC1155 tokens on Layer 1.
	L1BatchDepositERC1155
	// L1FinalizeBatchWithdrawERC1155 represents the event for batch finalizing ERC1155 token withdrawals on Layer 1.
	L1FinalizeBatchWithdrawERC1155
	// L1BatchRefundERC1155 the event for batch refund erc1155 on Layer 1.
	L1BatchRefundERC1155

	// L2FinalizeBatchDepositERC1155 represents the event for batch finalizing ERC1155 token deposits on Layer 2.
	L2FinalizeBatchDepositERC1155
	// L2BatchWithdrawERC1155 represents the event for batch withdrawing ERC1155 tokens on Layer 2.
	L2BatchWithdrawERC1155
)

type TxEventCategory int

const (
	ERC20EventCategory TxEventCategory = iota + 1
	ERC721EventCategory
	ERC1155EventCategory
	ETHEventCategory
	MessengerEventCategory
)

type ERC20 int

const (
	WETH ERC20 = iota + 1
	StandardERC20
	CustomERC20
	USDC
	DAI
	LIDO
)
