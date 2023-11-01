package contracts

import (
	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/go-ethereum/ethclient"
)

type l2Contracts struct {
	client *ethclient.Client
}

func newL2Contracts(c *ethclient.Client) *l2Contracts {
	return &l2Contracts{
		client: c,
	}
}

func (l *l2Contracts) register(conf config.Config) error {
	// TODO add others
	return nil
}
