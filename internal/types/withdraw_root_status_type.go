package types

//go:generate stringer -type WithdrawRootStatus

// WithdrawRootStatus represents the status of withdraw root check in layer 2.
type WithdrawRootStatus int

const (
	// WithdrawRootStatusTypeUnknown represents an unknown l2 withdraw root status.
	WithdrawRootStatusTypeUnknown WithdrawRootStatus = iota
	// WithdrawRootStatusTypeValid represents a valid l2 withdraw root status.
	WithdrawRootStatusTypeValid
)
