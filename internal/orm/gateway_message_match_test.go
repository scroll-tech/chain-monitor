package orm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils/testcontainer"
)

func TestGatewayMessageMatch_InsertOrUpdateEventInfo(t *testing.T) {
	ctx := context.Background()
	db := testcontainer.SetupDB(ctx, t)
	gatewayMessageMatchOrm := NewGatewayMessageMatch(db)

	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "l1ERC20DepositDuplicated",
			test: func(t *testing.T) {
				l1EventMsg1 := GatewayMessageMatch{
					MessageHash:   "0x1",
					TokenType:     int(types.TokenTypeERC20),
					L1EventType:   int(types.L1DepositERC20),
					L1BlockNumber: 120,
					L1TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
					L1TokenIds:    "",
					L1Amounts:     "200000000",
				}
				affectRows, err := gatewayMessageMatchOrm.InsertOrUpdateEventInfo(ctx, types.Layer1, l1EventMsg1)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(1))

				l1EventMsg1InsertDuplicated := GatewayMessageMatch{
					MessageHash:   "0x1",
					TokenType:     int(types.TokenTypeERC20),
					L1EventType:   int(types.L1DepositERC20),
					L1BlockNumber: 120,
					L1TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
					L1TokenIds:    "",
					L1Amounts:     "200000000",
				}
				affectRows, err = gatewayMessageMatchOrm.InsertOrUpdateEventInfo(ctx, types.Layer1, l1EventMsg1InsertDuplicated)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(0))
			},
		},
		{
			name: "L2FinalizeDepositERC20Duplicated",
			test: func(t *testing.T) {
				l2EventMsg1 := GatewayMessageMatch{
					MessageHash:   "0x1",
					TokenType:     int(types.TokenTypeERC20),
					L2EventType:   int(types.L2FinalizeDepositERC20),
					L2BlockNumber: 1200,
					L2TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
					L2TokenIds:    "",
					L2Amounts:     "200000000",
				}
				affectRows, err := gatewayMessageMatchOrm.InsertOrUpdateEventInfo(ctx, types.Layer2, l2EventMsg1)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(1))

				l2EventMsgInsertDuplicated := GatewayMessageMatch{
					MessageHash:   "0x1",
					TokenType:     int(types.TokenTypeERC20),
					L2EventType:   int(types.L2FinalizeDepositERC20),
					L2BlockNumber: 1200,
					L2TxHash:      "0xfc7d3ea5ec8dc9b664a5a886c3b33d21e665355057601033481a439498efb79a",
					L2TokenIds:    "",
					L2Amounts:     "200000000",
				}
				affectRows, err = gatewayMessageMatchOrm.InsertOrUpdateEventInfo(ctx, types.Layer2, l2EventMsgInsertDuplicated)
				assert.NoError(t, err)
				assert.Equal(t, affectRows, int64(0))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
