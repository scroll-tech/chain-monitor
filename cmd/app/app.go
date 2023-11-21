package app

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/orm/migrate"
	"github.com/scroll-tech/chain-monitor/internal/route"
	"github.com/scroll-tech/chain-monitor/internal/utils"
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
	app.Before = func(ctx *cli.Context) error {
		return utils.LogSetup(ctx)
	}
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

	// Start chain-monitor api server.
	if ctx.Bool(utils.HTTPEnabledFlag.Name) {
		endpoint := fmt.Sprintf(
			"%s:%d",
			ctx.String(utils.HTTPListenAddrFlag.Name),
			ctx.Int(utils.HTTPPortFlag.Name),
		)
		utils.StartServer(subCtx, endpoint, route.APIHandler(db.WithContext(subCtx)))
	}

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
