package types

//go:generate stringer -type CheckStatus

// CheckStatus the status of cross chain whether it checked
type CheckStatus int

const (
	// CheckStatusUnchecked the cross chain status hasn't checked
	CheckStatusUnchecked CheckStatus = iota
	// CheckStatusChecked the cross  chain status have checked
	CheckStatusChecked
)
