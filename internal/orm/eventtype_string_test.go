package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventType_String(t *testing.T) {
	assert.Equal(t, "L1DepositETH", L1DepositETH.String())
	assert.Equal(t, "L1FinalizeWithdrawETH", L1FinalizeWithdrawETH.String())
}
