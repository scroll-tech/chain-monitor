package monitor

import "chain-monitor/internal/orm"

type msgEvents struct {
	L1Type orm.EventType `json:"l1_type,omitempty" gorm:"l1_type"`
	L2Type orm.EventType `json:"l2_type,omitempty" gorm:"l2_type"`

	L1Number uint64 `json:"l1_number,omitempty" gorm:"l1_number"`
	L2Number uint64 `json:"l2_number,omitempty" gorm:"l2_number"`

	L1TxHash string `gorm:"l1_tx_hash"`
	L2TxHash string `gorm:"l2_tx_hash"`

	// asset fields
	L1Amount string `gorm:"l1_amount"`
	L2Amount string `gorm:"l2_amount"`

	L1TokenIDList string `json:"l1_token_id_list,omitempty" gorm:"l1_token_id_list"`
	L2TokenIDList string `json:"l2_token_id_list,omitempty" gorm:"l2_token_id_list"`
	L1AmountList  string `json:"l1_amount_list,omitempty" gorm:"l1_amount_list"`
	L2AmountList  string `json:"l2_amount_list,omitempty" gorm:"l2_amount_list"`
}
