package types

//go:generate stringer -type CheckStatus

// CheckStatus the status of cross chain whether it checked
type CheckStatus int

const (
	// CheckStatusUnknown represents an unknown or undefined balance status.
	CheckStatusUnknown CheckStatus = iota
	// CheckStatusUnchecked the cross chain status hasn't checked
	CheckStatusUnchecked
	// CheckStatusChecked the cross  chain status have checked
	CheckStatusChecked
)
