package app

import (
	"chain-monitor/internal/config"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/controller/monitor"
	"chain-monitor/internal/route"
	"chain-monitor/internal/utils"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"chain-monitor/orm"
)

var (
	app      *cli.App
	initFlag = &cli.BoolFlag{
		Name:  "init",
		Usage: "Clean and rebuild chain-monitor database",
		Value: false,
	}
)

func init() {
	// Set up event-watcher app info.
	app = cli.NewApp()
	app.Action = action
	app.Name = "event-watcher"
	app.Usage = "The Scroll Event Watcher"
	app.Version = utils.Version
	app.Flags = append(app.Flags, utils.CommonFlags...)
	app.Flags = append(app.Flags, initFlag)
	app.Commands = []*cli.Command{}
	app.Before = utils.BeforeAction
}

func action(ctx *cli.Context) error {
	subCtx := ctx.Context
	// Load config file.
	cfgFile := ctx.String(utils.ConfigFileFlag.Name)
	cfg, err := config.NewConfig(cfgFile)
	if err != nil {
		log.Crit("failed to load config file", "config file", cfgFile, "error", err)
	}
	// Create db instance.
	db, err := utils.InitDB(cfg.DBConfig)
	if err != nil {
		log.Error("failed to init db", "err", err)
		return err
	}
	defer utils.CloseDB(db)

	// Clean and rebuild db tables.
	if ctx.Bool(initFlag.Name) {
		// Clean tables.
		if err = orm.DropTables(db); err != nil {
			log.Error("failed to drop tables", "err", err)
			return err
		}
		// Create db tables.
		if err = orm.CreateTables(db); err != nil {
			log.Error("failed to migrate tables", "err", err)
			return err
		}
	}

	// Start onchain_metrics server.
	utils.StartServer(subCtx, ctx, route.Route(db.WithContext(subCtx)))

	l1Watcher, err := controller.NewL1Watcher(cfg.L1Config, db.WithContext(subCtx))
	if err != nil {
		log.Error("failed to create l1 watcher instance", "err", err)
		return err
	}
	_ = l1Watcher

	l2Watcher, err := controller.NewL2Watcher(cfg.L2Config, db.WithContext(subCtx))
	if err != nil {
		log.Error("failed to create l2 watcher instance", "err", err)
		return err
	}
	_ = l2Watcher

	chainMonitor := monitor.NewChainMonitor(cfg.ChainMonitor, db.WithContext(subCtx), l2Watcher)
	_ = chainMonitor

	go utils.LoopWithContext(subCtx, time.Second*2, l1Watcher.ScanL1Chain)
	go utils.LoopWithContext(subCtx, time.Second*2, l2Watcher.ScanL2Chain)
	go utils.LoopWithContext(subCtx, time.Second*3, chainMonitor.ChainMonitor)

	// Catch CTRL-C to ensure a graceful shutdown.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Wait until the interrupt signal is received from an OS signal.
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
