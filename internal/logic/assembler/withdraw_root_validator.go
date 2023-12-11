package assembler

import (
	"context"
	"fmt"
	"sort"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"

	"github.com/scroll-tech/chain-monitor/internal/logic/slack"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
	"github.com/scroll-tech/chain-monitor/internal/utils/msgproof"
)

func (c *MessageMatchAssembler) checkL2WithdrawRoots(ctx context.Context, startBlockNumber, endBlockNumber uint64, client *rpc.Client, messageQueueAddr common.Address) (*orm.MessengerMessageMatch, error) {
	log.Info("checking l2 withdraw roots", "start", startBlockNumber, "end", endBlockNumber)

	if startBlockNumber > endBlockNumber {
		return nil, nil
	}
	// recover latest withdraw trie.
	withdrawTrie := msgproof.NewWithdrawTrie()
	msg, err := c.messengerMessageMatchOrm.GetLatestValidL2SentMessageMatch(ctx)
	if err != nil {
		return nil, fmt.Errorf("get largest message nonce l2 message match failed, err: %w", err)
	}
	if msg != nil {
		withdrawTrie.Initialize(msg.NextMessageNonce-1, common.HexToHash(msg.MessageHash), msg.MessageProof)
	}

	l2SentMessages, err := c.messengerMessageMatchOrm.GetL2SentMessagesInBlockRange(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		return nil, fmt.Errorf("get l2 sent messages in block range failed, err: %w", err)
	}
	if len(l2SentMessages) == 0 {
		return nil, nil
	}

	sentMessageEventHashesMap := make(map[uint64][]common.Hash)
	for _, message := range l2SentMessages {
		sentMessageEventHashesMap[message.L2BlockNumber] = append(sentMessageEventHashesMap[message.L2BlockNumber], common.HexToHash(message.MessageHash))
	}

	var blockNums []uint64
	for blockNumber := range sentMessageEventHashesMap {
		blockNums = append(blockNums, blockNumber)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })

	withdrawRoots, err := utils.GetL2WithdrawRootsForBlocks(ctx, client, messageQueueAddr, blockNums)
	if err != nil {
		return nil, fmt.Errorf("get l2 withdraw roots failed, message queue addr: %v, blocks: %v, err: %w", messageQueueAddr, blockNums, err)
	}

	var lastMessage *orm.MessengerMessageMatch
	for _, blockNum := range blockNums {
		eventHashes := sentMessageEventHashesMap[blockNum]
		proofs := withdrawTrie.AppendMessages(eventHashes)
		lastWithdrawRoot := withdrawTrie.MessageRoot()
		if lastWithdrawRoot != withdrawRoots[blockNum] {
			info := slack.WithdrawRootInfo{
				BlockNumber:          blockNum,
				LastWithdrawRoot:     lastWithdrawRoot,
				ExpectedWithdrawRoot: withdrawRoots[blockNum],
			}
			log.Error("withdraw root mismatch",
				"block number", blockNum,
				"got", lastWithdrawRoot,
				"except", withdrawRoots[blockNum],
			)
			slack.Notify(slack.MrkDwnWithdrawRootMessage(info))
			return nil, fmt.Errorf("withdraw root mismatch in %v, got: %v, expected %v", blockNum, lastWithdrawRoot, withdrawRoots[blockNum])
		}
		// current block has SentMessage events.
		numEvents := len(eventHashes)
		if numEvents > 0 {
			// only update the last message of this check.
			lastMessage = &orm.MessengerMessageMatch{
				MessageHash:           eventHashes[numEvents-1].Hex(),
				MessageProof:          proofs[numEvents-1],
				WithdrawRootStatus:    int(types.WithdrawRootStatusTypeValid),
				MessageProofUpdatedAt: utils.NowUTC(),
				NextMessageNonce:      withdrawTrie.NextMessageNonce,
				L2BlockNumber:         blockNum,
			}
		}
	}
	return lastMessage, nil
}
