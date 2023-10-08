// Code generated by "stringer -type EventType"; DO NOT EDIT.

package orm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[L1DepositETH-1]
	_ = x[L1FinalizeWithdrawETH-2]
	_ = x[L2FinalizeDepositETH-101]
	_ = x[L2WithdrawETH-102]
	_ = x[L1DepositWETH-11]
	_ = x[L1FinalizeWithdrawWETH-12]
	_ = x[L1DepositDAI-13]
	_ = x[L1FinalizeWithdrawDAI-14]
	_ = x[L1DepositStandardERC20-15]
	_ = x[L1FinalizeWithdrawStandardERC20-16]
	_ = x[L1DepositCustomERC20-17]
	_ = x[L1FinalizeWithdrawCustomERC20-18]
	_ = x[L1USDCDepositERC20-19]
	_ = x[L1USDCFinalizeWithdrawERC20-20]
	_ = x[L1LIDODepositERC20-21]
	_ = x[L1LIDOFinalizeWithdrawERC20-22]
	_ = x[L1RefundETH-41]
	_ = x[L1RefundWETH-42]
	_ = x[L1RefundDAI-43]
	_ = x[L1RefundStandard-44]
	_ = x[L1RefundCustom-45]
	_ = x[L1RefundUSDC-46]
	_ = x[L1RefundLIDO-47]
	_ = x[L1RefundERC721-48]
	_ = x[L1RefundERC1155-49]
	_ = x[L2FinalizeDepositWETH-111]
	_ = x[L2WithdrawWETH-112]
	_ = x[L2FinalizeDepositDAI-113]
	_ = x[L2WithdrawDAI-114]
	_ = x[L2FinalizeDepositStandardERC20-115]
	_ = x[L2WithdrawStandardERC20-116]
	_ = x[L2FinalizeDepositCustomERC20-117]
	_ = x[L2WithdrawCustomERC20-118]
	_ = x[L2USDCWithdrawERC20-119]
	_ = x[L2USDCFinalizeDepositERC20-120]
	_ = x[L2LIDOWithdrawERC20-121]
	_ = x[L2LIDOFinalizeDepositERC20-122]
	_ = x[L1DepositERC721-61]
	_ = x[L1FinalizeWithdrawERC721-62]
	_ = x[L2FinalizeDepositERC721-161]
	_ = x[L2WithdrawERC721-162]
	_ = x[L1DepositERC1155-71]
	_ = x[L1FinalizeWithdrawERC1155-72]
	_ = x[L2FinalizeDepositERC1155-171]
	_ = x[L2WithdrawERC1155-172]
	_ = x[L1SentMessage-201]
	_ = x[L1RelayedMessage-202]
	_ = x[L1FailedRelayedMessage-203]
	_ = x[L2SentMessage-201]
	_ = x[L2RelayedMessage-202]
	_ = x[L2FailedRelayedMessage-203]
}

const (
	_EventType_name_0 = "L1DepositETHL1FinalizeWithdrawETH"
	_EventType_name_1 = "L1DepositWETHL1FinalizeWithdrawWETHL1DepositDAIL1FinalizeWithdrawDAIL1DepositStandardERC20L1FinalizeWithdrawStandardERC20L1DepositCustomERC20L1FinalizeWithdrawCustomERC20L1USDCDepositERC20L1USDCFinalizeWithdrawERC20L1LIDODepositERC20L1LIDOFinalizeWithdrawERC20"
	_EventType_name_2 = "L1RefundETHL1RefundWETHL1RefundDAIL1RefundStandardL1RefundCustomL1RefundUSDCL1RefundLIDOL1RefundERC721L1RefundERC1155"
	_EventType_name_3 = "L1DepositERC721L1FinalizeWithdrawERC721"
	_EventType_name_4 = "L1DepositERC1155L1FinalizeWithdrawERC1155"
	_EventType_name_5 = "L2FinalizeDepositETHL2WithdrawETH"
	_EventType_name_6 = "L2FinalizeDepositWETHL2WithdrawWETHL2FinalizeDepositDAIL2WithdrawDAIL2FinalizeDepositStandardERC20L2WithdrawStandardERC20L2FinalizeDepositCustomERC20L2WithdrawCustomERC20L2USDCWithdrawERC20L2USDCFinalizeDepositERC20L2LIDOWithdrawERC20L2LIDOFinalizeDepositERC20"
	_EventType_name_7 = "L2FinalizeDepositERC721L2WithdrawERC721"
	_EventType_name_8 = "L2FinalizeDepositERC1155L2WithdrawERC1155"
	_EventType_name_9 = "L1SentMessageL1RelayedMessageL1FailedRelayedMessage"
)

var (
	_EventType_index_0 = [...]uint8{0, 12, 33}
	_EventType_index_1 = [...]uint16{0, 13, 35, 47, 68, 90, 121, 141, 170, 188, 215, 233, 260}
	_EventType_index_2 = [...]uint8{0, 11, 23, 34, 50, 64, 76, 88, 102, 117}
	_EventType_index_3 = [...]uint8{0, 15, 39}
	_EventType_index_4 = [...]uint8{0, 16, 41}
	_EventType_index_5 = [...]uint8{0, 20, 33}
	_EventType_index_6 = [...]uint16{0, 21, 35, 55, 68, 98, 121, 149, 170, 189, 215, 234, 260}
	_EventType_index_7 = [...]uint8{0, 23, 39}
	_EventType_index_8 = [...]uint8{0, 24, 41}
	_EventType_index_9 = [...]uint8{0, 13, 29, 51}
)

func (i EventType) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _EventType_name_0[_EventType_index_0[i]:_EventType_index_0[i+1]]
	case 11 <= i && i <= 22:
		i -= 11
		return _EventType_name_1[_EventType_index_1[i]:_EventType_index_1[i+1]]
	case 41 <= i && i <= 49:
		i -= 41
		return _EventType_name_2[_EventType_index_2[i]:_EventType_index_2[i+1]]
	case 61 <= i && i <= 62:
		i -= 61
		return _EventType_name_3[_EventType_index_3[i]:_EventType_index_3[i+1]]
	case 71 <= i && i <= 72:
		i -= 71
		return _EventType_name_4[_EventType_index_4[i]:_EventType_index_4[i+1]]
	case 101 <= i && i <= 102:
		i -= 101
		return _EventType_name_5[_EventType_index_5[i]:_EventType_index_5[i+1]]
	case 111 <= i && i <= 122:
		i -= 111
		return _EventType_name_6[_EventType_index_6[i]:_EventType_index_6[i+1]]
	case 161 <= i && i <= 162:
		i -= 161
		return _EventType_name_7[_EventType_index_7[i]:_EventType_index_7[i+1]]
	case 171 <= i && i <= 172:
		i -= 171
		return _EventType_name_8[_EventType_index_8[i]:_EventType_index_8[i+1]]
	case 201 <= i && i <= 203:
		i -= 201
		return _EventType_name_9[_EventType_index_9[i]:_EventType_index_9[i+1]]
	default:
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
