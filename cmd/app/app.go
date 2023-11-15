package app

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"chain-monitor/internal/config"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/controller/l1watcher"
	"chain-monitor/internal/controller/l2watcher"
	"chain-monitor/internal/controller/monitor"
	"chain-monitor/internal/orm/migrate"
	"chain-monitor/internal/route"
	"chain-monitor/internal/utils"
)

var (
	app *cli.App
)

func init() {
	// Set up event-watcher app info.
	app = cli.NewApp()
	app.Action = action
	app.Name = "chain-monitor"
	app.Usage = "The Scroll chain monitor"
	app.Version = utils.Version
	app.Flags = append(app.Flags, utils.CommonFlags...)
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
		log.Error("failed to connect to db", "err", err)
		return err
	}
	defer func() {
		if err = utils.CloseDB(db); err != nil {
			log.Error("failed to close database", "err", err)
		}
	}()

	// db operation.
	if ctx.Bool(utils.DBFlag.Name) {
		if ctx.Bool(utils.DBMigrateFlag.Name) {
			return migrate.Migrate(db)
		}
		if ctx.Bool(utils.DBResetFlag.Name) {
			return migrate.Rollback(db, 0)
		}
		if ctx.IsSet(utils.DBRollBackFlag.Name) {
			return migrate.Rollback(db, ctx.Int64(utils.DBRollBackFlag.Name))
		}
	}

	// Init metrics.
	controller.InitChainMonitorMetrics()

	// Start chain-monitor api server.
	if ctx.Bool(utils.HTTPEnabledFlag.Name) {
		endpoint := fmt.Sprintf(
			"%s:%d",
			ctx.String(utils.HTTPListenAddrFlag.Name),
			ctx.Int(utils.HTTPPortFlag.Name),
		)
		utils.StartServer(subCtx, endpoint, route.APIHandler(db.WithContext(subCtx)))
	}

	// Start metrics server.
	if ctx.Bool(utils.MetricsEnabled.Name) {
		endpoint := fmt.Sprintf("%s:%d",
			ctx.String(utils.MetricsAddr.Name),
			ctx.Int(utils.MetricsPort.Name),
		)
		utils.StartServer(subCtx, endpoint, route.MetricsHandler(db))
	}

	// Init webhook alert.
	controller.InitWebhookAlert(cfg.AlertConfig)

	l1Watcher, err := l1watcher.NewL1Watcher(cfg.L1Config, db.WithContext(subCtx))
	if err != nil {
		log.Error("failed to create l1 watcher instance", "err", err)
		return err
	}
	_ = l1Watcher

	l2Watcher, err := l2watcher.NewL2Watcher(cfg.L2Config, db.WithContext(subCtx))
	if err != nil {
		log.Error("failed to create l2 watcher instance", "err", err)
		return err
	}
	_ = l2Watcher

	chainMonitor, err := monitor.NewChainMonitor(db.WithContext(subCtx), l1Watcher, l2Watcher)
	if err != nil {
		log.Error("failed to create chain chainMonitor instance", "err", err)
		return err
	}
	_ = chainMonitor

	go utils.LoopWithContext(subCtx, time.Millisecond*1500, l1Watcher.ScanL1Chain)
	go utils.LoopWithContext(subCtx, time.Millisecond*1500, l2Watcher.ScanL2Chain)
	go utils.LoopWithContext(subCtx, time.Millisecond*200, chainMonitor.DepositConfirm)
	go utils.LoopWithContext(subCtx, time.Millisecond*500, chainMonitor.WithdrawConfirm)

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
