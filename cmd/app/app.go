package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/controller"
	"github.com/scroll-tech/chain-monitor/internal/orm/migrate"
	"github.com/scroll-tech/chain-monitor/internal/route"
	"github.com/scroll-tech/chain-monitor/internal/utils"
	"github.com/scroll-tech/chain-monitor/internal/utils/database"
	"github.com/scroll-tech/chain-monitor/internal/utils/observability"
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
	defer cancel()

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

	// db operation.
	if ctx.Bool(utils.DBFlag.Name) {
		if ctx.Bool(utils.DBMigrateFlag.Name) {
			return migrate.Migrate(db)
		}
		if ctx.Bool(utils.DBResetFlag.Name) {
			return migrate.ResetDB(db)
		}
		if ctx.IsSet(utils.DBRollBackFlag.Name) {
			return migrate.Rollback(db, ctx.Int64(utils.DBRollBackFlag.Name))
		}
	}

	l1Client, err := rpc.Dial(cfg.L1Config.L1URL)
	if err != nil {
		log.Crit("failed to connect to l1 geth", "l1 geth url", cfg.L1Config.L1URL, "err", err)
	}

	l2Client, err := rpc.Dial(cfg.L2Config.L2URL)
	if err != nil {
		log.Crit("failed to connect to l2 geth", "l2 geth url", cfg.L2Config.L2URL, "err", err)
	}

	observability.Server(ctx, db)

	slackAlert := controller.NewSlackAlertController(subCtx, cfg.AlertConfig)
	slackAlert.Start()

	contractCtl := controller.NewContractController(cfg, db, l1Client, l2Client)
	contractCtl.Watch(subCtx)

	crossChainCtl := controller.NewCrossChainController(cfg, db, ethclient.NewClient(l1Client), ethclient.NewClient(l2Client))
	crossChainCtl.Watch(subCtx)

	apiSrv := apiServer(ctx, cfg, db)

	log.Info("Start chain-monitor successfully.")

	defer func() {
		contractCtl.Stop()
		crossChainCtl.Stop()
		slackAlert.Stop()
		if err = database.CloseDB(db); err != nil {
			log.Error("failed to close database", "err", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Info("start shutdown chain-monitor server ...")

	closeCtx, cancelExit := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelExit()
	if err = apiSrv.Shutdown(closeCtx); err != nil {
		log.Warn("shutdown chain-monitor server failure", "error", err)
		return nil
	}

	<-closeCtx.Done()
	log.Info("chain-monitor server exiting success")
	return nil
}

func apiServer(ctx *cli.Context, cfg *config.Config, db *gorm.DB) *http.Server {
	log.Info("api controller start successful")

	router := gin.New()
	controller.InitApi(cfg, db)
	route.Route(router)
	port := ctx.String(utils.HTTPPortFlag.Name)
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
	}

	go func() {
		if runServerErr := srv.ListenAndServe(); runServerErr != nil && !errors.Is(runServerErr, http.ErrServerClosed) {
			log.Crit("run coordinator http server failure", "error", runServerErr)
		}
	}()
	return srv
}

// Run event watcher cmd instance.
func Run() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
