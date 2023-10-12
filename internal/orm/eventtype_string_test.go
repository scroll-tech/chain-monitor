package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventType_String(t *testing.T) {
	assert.Equal(t, "L1DepositETH", L1DepositETH.String())
	assert.Equal(t, "L1DepositWETH", L1DepositWETH.String())
	assert.Equal(t, "L1SentMessage", L1SentMessage.String())
	assert.Equal(t, "L1FinalizeWithdrawWETH", L1FinalizeWithdrawWETH.String())
}
