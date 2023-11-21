package slack

import (
	"bytes"
	"fmt"

	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

func MrkDwnCrossChainMessage(message *orm.MessageMatch, checkResult types.MismatchType) string {
	var buffer bytes.Buffer
	buffer.WriteString(":warning: ")
	buffer.WriteString("*Cross chain check failed*")
	buffer.WriteString(" :warning:\n")
	buffer.WriteString(fmt.Sprintf("• database id: %d\n", message.ID))
	buffer.WriteString(fmt.Sprintf("• token type: %s\n", types.TokenType(message.TokenType).String()))
	buffer.WriteString(fmt.Sprintf("• l1 event type: %s\n", types.EventType(message.L1EventType).String()))
	buffer.WriteString(fmt.Sprintf("• l2 event type: %s\n", types.EventType(message.L2EventType).String()))
	buffer.WriteString(fmt.Sprintf("• mismatch type: %s\n", checkResult.String()))
	buffer.WriteString(fmt.Sprintf("• l1 block number: %d\n", message.L1BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l2 block number: %d\n", message.L2BlockNumber))
	buffer.WriteString(fmt.Sprintf("• l1 mount: %s\n", message.L1Amounts))
	buffer.WriteString(fmt.Sprintf("• l2 mount: %s\n", message.L2Amounts))
	buffer.WriteString(fmt.Sprintf("• l1 token: %s\n", message.L1TokenIds))
	buffer.WriteString(fmt.Sprintf("• l2 token: %s\n", message.L2TokenIds))
	buffer.WriteString(fmt.Sprintf("• l1 tx_hash: %s\n", message.L1TxHash))
	buffer.WriteString(fmt.Sprintf("• l2 tx_hash: %s\n", message.L2TxHash))
	buffer.WriteString(fmt.Sprintf("• msg_hash: %s\n", message.MessageHash))
	return buffer.String()
}

func MrkDwnETHMessage() {

}
