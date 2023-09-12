package monitor

type msgEvents struct {
	L1Number uint64 `json:"l1_number,omitempty" gorm:"l1_number"`
	L2Number uint64 `json:"l2_number,omitempty" gorm:"l2_number"`

	L1TxHash string `gorm:"l1_tx_hash"`
	L2TxHash string `gorm:"l2_tx_hash"`

	// asset fields
	L1Amount  string `gorm:"l1_amount"`
	L2Amount  string `gorm:"l2_amount"`
	L1TokenID string `json:"l1_token_id,omitempty" gorm:"l1_token_id"`
	L2TokenID string `json:"l2_token_id,omitempty" gorm:"l2_token_id"`
}
