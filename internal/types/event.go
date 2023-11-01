package types

// EventType represents the type of blockchain event.
type EventType uint8

const (
	EventTypeUnknown EventType = iota

	// L1DepositETH represents the event for depositing ETH on Layer 1.
	L1DepositETH
	// L1FinalizeWithdrawETH represents the event for finalizing ETH withdrawal on Layer 1.
	L1FinalizeWithdrawETH
	// L1RefundETH the event for refund eth
	L1RefundETH

	// L1DepositWETH represents the event for depositing Wrapped ETH on Layer 1.
	L1DepositWETH
	// L1FinalizeWithdrawWETH represents the event for finalizing Wrapped ETH withdrawal on Layer 1.
	L1FinalizeWithdrawWETH
	// L1RefundWETH the event for refund weth
	L1RefundWETH

	// L2FinalizeDepositETH represents the event for finalizing ETH deposit on Layer 2.
	L2FinalizeDepositETH
	// L2WithdrawETH represents the event for withdrawing ETH on Layer 2.
	L2WithdrawETH

	// L1DepositDAI represents the event for depositing DAI on Layer 1.
	L1DepositDAI
	// L1FinalizeWithdrawDAI represents the event for finalizing DAI withdrawal on Layer 1.
	L1FinalizeWithdrawDAI
	// L1DepositStandardERC20 represents the event for depositing standard ERC20 tokens on Layer 1.
	L1DepositStandardERC20
	// L1FinalizeWithdrawStandardERC20 represents the event for finalizing standard ERC20 token withdrawal on Layer 1.
	L1FinalizeWithdrawStandardERC20
	// L1DepositCustomERC20 represents the event for depositing custom ERC20 tokens on Layer 1.
	L1DepositCustomERC20
	// L1FinalizeWithdrawCustomERC20 represents the event for finalizing custom ERC20 token withdrawal on Layer 1.
	L1FinalizeWithdrawCustomERC20
	// L1USDCDepositERC20 represents the event for finalizing USDC ERC20 token withdrawal on Layer 1.
	L1USDCDepositERC20
	// L1USDCFinalizeWithdrawERC20 represents the event for finalizing USDC ERC20 token withdrawal on Layer 1.
	L1USDCFinalizeWithdrawERC20
	// L1LIDODepositERC20 represents the event for finalizing LIDO ERC20 token withdrawal on Layer 1.
	L1LIDODepositERC20
	// L1LIDOFinalizeWithdrawERC20 represents the event for finalizing LIDO ERC20 token withdrawal on Layer 1.
	L1LIDOFinalizeWithdrawERC20
	// L2FinalizeDepositWETH represents the event for finalizing Wrapped ETH deposit on Layer 2.
	L2FinalizeDepositWETH
	// L2WithdrawWETH represents the event for withdrawing Wrapped ETH on Layer 2.
	L2WithdrawWETH
	// L2FinalizeDepositDAI represents the event for finalizing DAI deposit on Layer 2.
	L2FinalizeDepositDAI
	// L2WithdrawDAI represents the event for withdrawing DAI on Layer 2.
	L2WithdrawDAI
	// L2FinalizeDepositStandardERC20 represents the event for finalizing standard ERC20 token deposit on Layer 2.
	L2FinalizeDepositStandardERC20
	// L2WithdrawStandardERC20 represents the event for withdrawing standard ERC20 tokens on Layer 2.
	L2WithdrawStandardERC20
	// L2FinalizeDepositCustomERC20 represents the event for finalizing custom ERC20 token deposit on Layer 2.
	L2FinalizeDepositCustomERC20
	// L2WithdrawCustomERC20 represents the event for withdrawing custom ERC20 tokens on Layer 2.
	L2WithdrawCustomERC20
	// L2USDCWithdrawERC20 represents the event for withdrawing USDC ERC20 tokens on Layer 2.
	L2USDCWithdrawERC20
	// L2USDCFinalizeDepositERC20 represents the event for withdrawing USDC ERC20 tokens on Layer 2.
	L2USDCFinalizeDepositERC20
	// L2LIDOWithdrawERC20 represents the event for withdrawing LIDO ERC20 tokens on Layer 2.
	L2LIDOWithdrawERC20
	// L2LIDOFinalizeDepositERC20 represents the event for withdrawing LIDO ERC20 tokens on Layer 2.
	L2LIDOFinalizeDepositERC20

	// L1DepositERC721 represents the event for depositing ERC721 tokens on Layer 1.
	L1DepositERC721
	// L1FinalizeWithdrawERC721 represents the event for finalizing ERC721 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC721
	// L2FinalizeDepositERC721 represents the event for finalizing ERC721 token deposit on Layer 2.
	L2FinalizeDepositERC721
	// L2WithdrawERC721 represents the event for withdrawing ERC721 tokens on Layer 2.
	L2WithdrawERC721

	// L1DepositERC1155 represents the event for depositing ERC1155 tokens on Layer 1.
	L1DepositERC1155
	// L1FinalizeWithdrawERC1155 represents the event for finalizing ERC1155 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC1155
	// L2FinalizeDepositERC1155 represents the event for finalizing ERC1155 token deposit on Layer 2.
	L2FinalizeDepositERC1155
	// L2WithdrawERC1155 represents the event for withdrawing ERC1155 tokens on Layer 2.
	L2WithdrawERC1155
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
	DAI
	LIDO
	StandardERC20
	CustomERC20
	USDCERC20
)
