package assembler

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

func (c *MessageMatchAssembler) messengerMessageMatchAssembler(messengerEvents []events.EventUnmarshaler) ([]orm.MessengerMessageMatch, error) {
	if len(messengerEvents) == 0 {
		return nil, nil
	}

	var messageMatches []orm.MessengerMessageMatch
	for _, eventData := range messengerEvents {
		var tmpMessageMatch orm.MessengerMessageMatch
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return nil, fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		switch messengerEventUnmarshaler.Type {
		case types.L1SentMessage:
			tmpMessageMatch = orm.MessengerMessageMatch{
				MessageHash:     messengerEventUnmarshaler.MessageHash.Hex(),
				L1EventType:     int(messengerEventUnmarshaler.Type),
				L1BlockNumber:   messengerEventUnmarshaler.Number,
				L1TxHash:        messengerEventUnmarshaler.TxHash.Hex(),
				ETHAmount:       decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0).String(),
				ETHAmountStatus: int(types.ETHAmountStatusTypeSet),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L1RelayedMessage:
			tmpMessageMatch = orm.MessengerMessageMatch{
				MessageHash:   messengerEventUnmarshaler.MessageHash.Hex(),
				L1EventType:   int(messengerEventUnmarshaler.Type),
				L1BlockNumber: messengerEventUnmarshaler.Number,
				L1TxHash:      messengerEventUnmarshaler.TxHash.Hex(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L2SentMessage:
			tmpMessageMatch = orm.MessengerMessageMatch{
				MessageHash:      messengerEventUnmarshaler.MessageHash.Hex(),
				L2EventType:      int(messengerEventUnmarshaler.Type),
				L2BlockNumber:    messengerEventUnmarshaler.Number,
				L2TxHash:         messengerEventUnmarshaler.TxHash.Hex(),
				ETHAmount:        decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0).String(),
				ETHAmountStatus:  int(types.ETHAmountStatusTypeSet),
				NextMessageNonce: messengerEventUnmarshaler.MessageNonce.Uint64() + 1,
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L2RelayedMessage:
			tmpMessageMatch = orm.MessengerMessageMatch{
				MessageHash:   messengerEventUnmarshaler.MessageHash.Hex(),
				L2EventType:   int(messengerEventUnmarshaler.Type),
				L2BlockNumber: messengerEventUnmarshaler.Number,
				L2TxHash:      messengerEventUnmarshaler.TxHash.Hex(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		}
	}
	return messageMatches, nil
}
