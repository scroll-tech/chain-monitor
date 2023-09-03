package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

func StartServer(ctx *cli.Context, handler http.Handler) {
	if !ctx.Bool(httpEnabledFlag.Name) {
		return
	}
	endpoint := fmt.Sprintf("%s:%d", ctx.String(httpListenAddrFlag.Name), ctx.Int(httpPortFlag.Name))

	srv := &http.Server{
		Handler:      handler,
		Addr:         endpoint,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		IdleTimeout:  time.Second * 12,
	}

	go func() {
		<-ctx.Context.Done()
		if err := srv.Close(); err != nil {
			log.Crit("failed to close chain_monitor serer", "err", err)
		}
	}()

	log.Info("starting chain_monitor server", "endpoint", fmt.Sprintf("http://%s/v1", endpoint))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("start chain_monitor server error", "error", err)
		}
	}()
}
