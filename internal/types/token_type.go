package types

//go:generate stringer -type TokenType

// TokenType represents the type of token
type TokenType uint8

const (
	TokenTypeUnknown TokenType = iota
	TokenTypeETH
	TokenTypeERC20
	TokenTypeERC721
	TokenTypeERC1155
)
