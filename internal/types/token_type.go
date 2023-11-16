package types

// TokenType represents the type of token in Ethereum.
type TokenType uint8

const (
	// TokenTypeDefined represents a default value for TokenType.
	TokenTypeDefined TokenType = iota
	// TokenTypeETH represents the native currency of Ethereum.
	TokenTypeETH
	// TokenTypeERC20 represents fungible tokens complying with ERC20 standard.
	TokenTypeERC20
	// TokenTypeERC721 represents non-fungible tokens complying with ERC721 standard.
	TokenTypeERC721
	// TokenTypeERC1155 represents semi-fungible tokens complying with ERC1155 standard.
	TokenTypeERC1155
)
