package types

//go:generate stringer -type ERC20

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
