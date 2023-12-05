package orm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils/testcontainer"
)

func TestMessengerMessageMatch_InsertOrUpdateEventInfo(t *testing.T) {
	ctx := context.Background()
	db := testcontainer.SetupDB(ctx, t)
	messengerOrm := NewMessengerMessageMatch(db)

	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			"l1SentMessage", func(t *testing.T) {
				l1SentEventMsg1 := MessengerMessageMatch{
					MessageHash:   "0x1",
					L1EventType:   int(types.L1SentMessage),
					L1BlockNumber: 120,
					L1TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
					ETHAmount:     "1000",
				}
				affectRows, err := messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer1, l1SentEventMsg1)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(1))

				l1SentEventMsg1Duplicated := MessengerMessageMatch{
					MessageHash:   "0x1",
					L1EventType:   int(types.L1SentMessage),
					L1BlockNumber: 120,
					L1TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
					ETHAmount:     "1000",
				}
				affectRows, err = messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer1, l1SentEventMsg1Duplicated)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(0))
			},
		},
		{
			"l1RelayedMessage", func(t *testing.T) {
				l1RelayedEventMsg1 := MessengerMessageMatch{
					MessageHash:   "0x2",
					L1EventType:   int(types.L1RelayedMessage),
					L1BlockNumber: 1120,
					L1TxHash:      "0x1c7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
				}
				affectRows, err := messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer1, l1RelayedEventMsg1)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(1))

				l1RelayedEventMsg1Duplicated := MessengerMessageMatch{
					MessageHash:   "0x2",
					L1EventType:   int(types.L1RelayedMessage),
					L1BlockNumber: 120,
					L1TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
				}
				affectRows, err = messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer1, l1RelayedEventMsg1Duplicated)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(0))
			},
		},
		{
			"l2RelayedMessage", func(t *testing.T) {
				l2RelayedEventMsg1 := MessengerMessageMatch{
					MessageHash:   "0x1",
					L2EventType:   int(types.L2RelayedMessage),
					L2BlockNumber: 1200,
					L2TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb791",
				}
				affectRows, err := messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer2, l2RelayedEventMsg1)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(1))

				l2RelayedEventMsg1Duplicated := MessengerMessageMatch{
					MessageHash:   "0x1",
					L2EventType:   int(types.L2RelayedMessage),
					L2BlockNumber: 1200,
					L2TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb791",
				}
				affectRows, err = messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer2, l2RelayedEventMsg1Duplicated)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(0))
			},
		},
		{
			"l2SentMessage", func(t *testing.T) {
				l2RelayedEventMsg1 := MessengerMessageMatch{
					MessageHash:   "0x2",
					L2EventType:   int(types.L2SentMessage),
					L2BlockNumber: 12000,
					L2TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb792",
					ETHAmount:     "10000",
				}
				affectRows, err := messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer2, l2RelayedEventMsg1)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(1))

				l2RelayedEventMsg1Duplicated := MessengerMessageMatch{
					MessageHash:   "0x2",
					L2EventType:   int(types.L2SentMessage),
					L2BlockNumber: 1200,
					L2TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb792",
					ETHAmount:     "1",
				}
				affectRows, err = messengerOrm.InsertOrUpdateEventInfo(ctx, types.Layer2, l2RelayedEventMsg1Duplicated)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(0))
			},
		},
		{
			"sentRelayMatched0x1", func(t *testing.T) {
				msgMatch, err := messengerOrm.GetMessageMatchByMessageHash(ctx, "0x1")
				assert.NoError(t, err)
				assert.NotNil(t, msgMatch)

				assert.Equal(t, msgMatch.L1EventType, int(types.L1SentMessage))
				assert.Equal(t, msgMatch.L1BlockNumber, uint64(120))
				assert.Equal(t, msgMatch.L1TxHash, "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a")
				assert.Equal(t, msgMatch.ETHAmount, "1000")

				assert.Equal(t, msgMatch.L2EventType, int(types.L2RelayedMessage))
				assert.Equal(t, msgMatch.L2BlockNumber, uint64(1200))
				assert.Equal(t, msgMatch.L2TxHash, "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb791")
			},
		},
		{
			"sentRelayMatched0x2", func(t *testing.T) {
				msgMatch, err := messengerOrm.GetMessageMatchByMessageHash(ctx, "0x2")
				assert.NoError(t, err)
				assert.NotNil(t, msgMatch)

				assert.Equal(t, msgMatch.L1EventType, int(types.L1RelayedMessage))
				assert.Equal(t, msgMatch.L1BlockNumber, uint64(1120))
				assert.Equal(t, msgMatch.L1TxHash, "0x1c7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a")

				assert.Equal(t, msgMatch.L2EventType, int(types.L2SentMessage))
				assert.Equal(t, msgMatch.L2BlockNumber, uint64(12000))
				assert.Equal(t, msgMatch.L2TxHash, "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb792")
				assert.Equal(t, msgMatch.ETHAmount, "10000")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
