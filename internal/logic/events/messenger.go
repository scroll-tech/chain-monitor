package events

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
)

// SentMessageEvent represents a SentMessage event raised by the L1/L2 scroll messenger contract.
type SentMessageEvent struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	MessageHash  common.Hash
}
