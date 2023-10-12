package l1watcher

import (
	"chain-monitor/bytecode/scroll/L2/gateway"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMessage(t *testing.T) {
	//cfg, err := config.NewConfig("../../../config.json")
	//assert.NoError(t, err)
	//client, err := ethclient.Dial(cfg.L1Config.L1URL)
	//assert.NoError(t, err)
	//ctx := context.Background()

	data := common.FromHex("8431f5c1000000000000000000000000fff9976782d46cc05630d1f6ebab18b2324d6b1400000000000000000000000053000000000000000000000000000000000000040000000000000000000000006fe0dd81a09e23430391d76320731a3cb6f1d8cf0000000000000000000000006fe0dd81a09e23430391d76320731a3cb6f1d8cf000000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000000")
	method, err := gateway.L2ERC20GatewayABI.MethodById(data)
	assert.NoError(t, err)
	t.Log(method.Name, common.BytesToHash(method.ID))
	event := new(gateway.L2ERC20GatewayFinalizeDepositERC20Event)
	//err = gateway.L2ETHGatewayABI.UnpackIntoInterface(event, "FinalizeDepositETH", data[4:])
	//outputs := method.Outputs
	//unpacked, err := outputs.Unpack(data)
	//assert.NoError(t, err)
	//_ = unpacked

	params, err := method.Inputs.Unpack(data[4:])
	assert.NoError(t, err)
	method.Inputs.Copy(event, params)
	t.Log(event)
}
