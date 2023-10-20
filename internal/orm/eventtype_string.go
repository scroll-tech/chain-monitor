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
	_ = x[L1DepositERC721-51]
	_ = x[L1FinalizeWithdrawERC721-52]
	_ = x[L1BatchDepositERC721-53]
	_ = x[L1BatchFinalizeWithdrawERC721-54]
	_ = x[L2FinalizeDepositERC721-151]
	_ = x[L2WithdrawERC721-152]
	_ = x[L2BatchFinalizeDepositERC721-153]
	_ = x[L2BatchWithdrawERC721-154]
	_ = x[L1DepositERC1155-71]
	_ = x[L1FinalizeWithdrawERC1155-72]
	_ = x[L1BatchDepositERC1155-73]
	_ = x[L1BatchFinalizeWithdrawERC1155-74]
	_ = x[L2FinalizeDepositERC1155-171]
	_ = x[L2WithdrawERC1155-172]
	_ = x[L2BatchFinalizeDepositERC1155-173]
	_ = x[L2BatchWithdrawERC1155-174]
	_ = x[L1SentMessage-201]
	_ = x[L1RelayedMessage-202]
	_ = x[L1FailedRelayedMessage-203]
	_ = x[L2SentMessage-211]
	_ = x[L2RelayedMessage-212]
	_ = x[L2FailedRelayedMessage-213]
}

const (
	_EventType_name_0 = "L1DepositETHL1FinalizeWithdrawETH"
	_EventType_name_1 = "L1DepositWETHL1FinalizeWithdrawWETHL1DepositDAIL1FinalizeWithdrawDAIL1DepositStandardERC20L1FinalizeWithdrawStandardERC20L1DepositCustomERC20L1FinalizeWithdrawCustomERC20L1USDCDepositERC20L1USDCFinalizeWithdrawERC20L1LIDODepositERC20L1LIDOFinalizeWithdrawERC20"
	_EventType_name_2 = "L1DepositERC721L1FinalizeWithdrawERC721L1BatchDepositERC721L1BatchFinalizeWithdrawERC721"
	_EventType_name_3 = "L1DepositERC1155L1FinalizeWithdrawERC1155L1BatchDepositERC1155L1BatchFinalizeWithdrawERC1155"
	_EventType_name_4 = "L2FinalizeDepositETHL2WithdrawETH"
	_EventType_name_5 = "L2FinalizeDepositWETHL2WithdrawWETHL2FinalizeDepositDAIL2WithdrawDAIL2FinalizeDepositStandardERC20L2WithdrawStandardERC20L2FinalizeDepositCustomERC20L2WithdrawCustomERC20L2USDCWithdrawERC20L2USDCFinalizeDepositERC20L2LIDOWithdrawERC20L2LIDOFinalizeDepositERC20"
	_EventType_name_6 = "L2FinalizeDepositERC721L2WithdrawERC721L2BatchFinalizeDepositERC721L2BatchWithdrawERC721"
	_EventType_name_7 = "L2FinalizeDepositERC1155L2WithdrawERC1155L2BatchFinalizeDepositERC1155L2BatchWithdrawERC1155"
	_EventType_name_8 = "L1SentMessageL1RelayedMessageL1FailedRelayedMessage"
	_EventType_name_9 = "L2SentMessageL2RelayedMessageL2FailedRelayedMessage"
)

var (
	_EventType_index_0 = [...]uint8{0, 12, 33}
	_EventType_index_1 = [...]uint16{0, 13, 35, 47, 68, 90, 121, 141, 170, 188, 215, 233, 260}
	_EventType_index_2 = [...]uint8{0, 15, 39, 59, 88}
	_EventType_index_3 = [...]uint8{0, 16, 41, 62, 92}
	_EventType_index_4 = [...]uint8{0, 20, 33}
	_EventType_index_5 = [...]uint16{0, 21, 35, 55, 68, 98, 121, 149, 170, 189, 215, 234, 260}
	_EventType_index_6 = [...]uint8{0, 23, 39, 67, 88}
	_EventType_index_7 = [...]uint8{0, 24, 41, 70, 92}
	_EventType_index_8 = [...]uint8{0, 13, 29, 51}
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
	case 51 <= i && i <= 54:
		i -= 51
		return _EventType_name_2[_EventType_index_2[i]:_EventType_index_2[i+1]]
	case 71 <= i && i <= 74:
		i -= 71
		return _EventType_name_3[_EventType_index_3[i]:_EventType_index_3[i+1]]
	case 101 <= i && i <= 102:
		i -= 101
		return _EventType_name_4[_EventType_index_4[i]:_EventType_index_4[i+1]]
	case 111 <= i && i <= 122:
		i -= 111
		return _EventType_name_5[_EventType_index_5[i]:_EventType_index_5[i+1]]
	case 151 <= i && i <= 154:
		i -= 151
		return _EventType_name_6[_EventType_index_6[i]:_EventType_index_6[i+1]]
	case 171 <= i && i <= 174:
		i -= 171
		return _EventType_name_7[_EventType_index_7[i]:_EventType_index_7[i+1]]
	case 201 <= i && i <= 203:
		i -= 201
		return _EventType_name_8[_EventType_index_8[i]:_EventType_index_8[i+1]]
	case 211 <= i && i <= 213:
		i -= 211
		return _EventType_name_9[_EventType_index_9[i]:_EventType_index_9[i+1]]
	default:
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
