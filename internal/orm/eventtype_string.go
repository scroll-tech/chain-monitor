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
	_ = x[L2SentMessage-211]
	_ = x[L2RelayedMessage-212]
	_ = x[L2FailedRelayedMessage-213]
}

const _EventType_name = "L1DepositETHL1FinalizeWithdrawETHL1DepositWETHL1FinalizeWithdrawWETHL1DepositDAIL1FinalizeWithdrawDAIL1DepositStandardERC20L1FinalizeWithdrawStandardERC20L1DepositCustomERC20L1FinalizeWithdrawCustomERC20L1USDCDepositERC20L1USDCFinalizeWithdrawERC20L1LIDODepositERC20L1LIDOFinalizeWithdrawERC20L1RefundETHL1RefundWETHL1RefundDAIL1RefundStandardL1RefundCustomL1RefundUSDCL1RefundLIDOL1RefundERC721L1RefundERC1155L1DepositERC721L1FinalizeWithdrawERC721L1DepositERC1155L1FinalizeWithdrawERC1155L2FinalizeDepositETHL2WithdrawETHL2FinalizeDepositWETHL2WithdrawWETHL2FinalizeDepositDAIL2WithdrawDAIL2FinalizeDepositStandardERC20L2WithdrawStandardERC20L2FinalizeDepositCustomERC20L2WithdrawCustomERC20L2USDCWithdrawERC20L2USDCFinalizeDepositERC20L2LIDOWithdrawERC20L2LIDOFinalizeDepositERC20L2FinalizeDepositERC721L2WithdrawERC721L2FinalizeDepositERC1155L2WithdrawERC1155L1SentMessageL1RelayedMessageL1FailedRelayedMessageL2SentMessageL2RelayedMessageL2FailedRelayedMessage"

var _EventType_map = map[EventType]string{
	1:   _EventType_name[0:12],
	2:   _EventType_name[12:33],
	11:  _EventType_name[33:46],
	12:  _EventType_name[46:68],
	13:  _EventType_name[68:80],
	14:  _EventType_name[80:101],
	15:  _EventType_name[101:123],
	16:  _EventType_name[123:154],
	17:  _EventType_name[154:174],
	18:  _EventType_name[174:203],
	19:  _EventType_name[203:221],
	20:  _EventType_name[221:248],
	21:  _EventType_name[248:266],
	22:  _EventType_name[266:293],
	41:  _EventType_name[293:304],
	42:  _EventType_name[304:316],
	43:  _EventType_name[316:327],
	44:  _EventType_name[327:343],
	45:  _EventType_name[343:357],
	46:  _EventType_name[357:369],
	47:  _EventType_name[369:381],
	48:  _EventType_name[381:395],
	49:  _EventType_name[395:410],
	61:  _EventType_name[410:425],
	62:  _EventType_name[425:449],
	71:  _EventType_name[449:465],
	72:  _EventType_name[465:490],
	101: _EventType_name[490:510],
	102: _EventType_name[510:523],
	111: _EventType_name[523:544],
	112: _EventType_name[544:558],
	113: _EventType_name[558:578],
	114: _EventType_name[578:591],
	115: _EventType_name[591:621],
	116: _EventType_name[621:644],
	117: _EventType_name[644:672],
	118: _EventType_name[672:693],
	119: _EventType_name[693:712],
	120: _EventType_name[712:738],
	121: _EventType_name[738:757],
	122: _EventType_name[757:783],
	161: _EventType_name[783:806],
	162: _EventType_name[806:822],
	171: _EventType_name[822:846],
	172: _EventType_name[846:863],
	201: _EventType_name[863:876],
	202: _EventType_name[876:892],
	203: _EventType_name[892:914],
	211: _EventType_name[914:927],
	212: _EventType_name[927:943],
	213: _EventType_name[943:965],
}

func (i EventType) String() string {
	if str, ok := _EventType_map[i]; ok {
		return str
	}
	return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
}
