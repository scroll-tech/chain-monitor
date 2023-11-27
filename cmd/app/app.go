package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/urfave/cli/v2"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/controller"
	"github.com/scroll-tech/chain-monitor/internal/utils"
	"github.com/scroll-tech/chain-monitor/internal/utils/database"
)

var app *cli.App

func init() {
	// Set up event-watcher app info.
	app = cli.NewApp()
	app.Action = action
	app.Name = "chain-monitor"
	app.Usage = "The Scroll chain monitor"
	app.Version = utils.Version
	app.Flags = append(app.Flags, utils.CommonFlags...)
	app.Commands = []*cli.Command{}
	app.Before = func(ctx *cli.Context) error {
		return utils.LogSetup(ctx)
	}
}

func action(ctx *cli.Context) error {
	subCtx, cancel := context.WithCancel(ctx.Context)

	// Load config file.
	cfgFile := ctx.String(utils.ConfigFileFlag.Name)
	cfg, err := config.NewConfig(cfgFile)
	if err != nil {
		log.Crit("failed to load config file", "config file", cfgFile, "error", err)
	}

	// Create db instance.
	db, err := database.InitDB(cfg.DBConfig)
	if err != nil {
		log.Crit("failed to connect to db", "err", err)
	}

	l1Client, err := rpc.Dial(cfg.L1Config.L1URL)
	if err != nil {
		log.Crit("failed to connect to l1 geth", "l1 geth url", cfg.L1Config.L1URL, "err", err)
	}

	l2Client, err := rpc.Dial(cfg.L2Config.L2URL)
	if err != nil {
		log.Crit("failed to connect to l2 geth", "l2 geth url", cfg.L2Config.L2URL, "err", err)
	}

	slackAlert := controller.NewSlackAlertController(subCtx, cfg.AlertConfig)
	slackAlert.Start()

	contractCtl := controller.NewContractController(cfg, db, l1Client, l2Client)
	contractCtl.Watch(subCtx)

	crossChainCtl := controller.NewCrossChainController(cfg, db, ethclient.NewClient(l1Client), ethclient.NewClient(l2Client))
	crossChainCtl.Watch(subCtx)

	defer func() {
		contractCtl.Stop()
		crossChainCtl.Stop()
		slackAlert.Stop()
		cancel()
		if err = database.CloseDB(db); err != nil {
			log.Error("failed to close database", "err", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	return nil
}

// Run event watcher cmd instance.
func Run() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
