package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/urfave/cli/v2"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/controller"
	"github.com/scroll-tech/chain-monitor/internal/utils"
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
	db, err := utils.InitDB(cfg.DBConfig)
	if err != nil {
		log.Crit("failed to connect to db", "err", err)
	}

	l1Client, err := rpc.Dial(cfg.L1Config.L1URL)
	if err != nil {
		log.Crit("failed to connect to l1 eth", "l1 eth url", cfg.L1Config.L1URL, "err", err)
	}

	l2Client, err := rpc.Dial(cfg.L2Config.L2URL)
	if err != nil {
		log.Crit("failed to connect to l2 eth", "l2 eth url", cfg.L2Config.L2URL, "err", err)
	}

	slackAlert := controller.NewSlackAlertController(subCtx, cfg.AlertConfig)
	slackAlert.Start()

	contractCtl := controller.NewContractController(*cfg, db, l1Client, l2Client)
	go contractCtl.Watch(subCtx)

	crossChainCtl := controller.NewCrossChainController(cfg, db, l1Client, l2Client)
	go crossChainCtl.Proposer(subCtx)

	defer func() {
		contractCtl.Stop()
		crossChainCtl.Stop()
		slackAlert.Stop()
		cancel()
		if err = utils.CloseDB(db); err != nil {
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
