package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
)

// StartServer initializes and starts an HTTP server based on provided CLI context settings.
// If the HTTP-RPC server is enabled via CLI context, it sets up the server using the provided
// handler and other context flags such as listen address and port.
// The server will listen and serve until the context is done, after which it will gracefully shut down.
func StartServer(ctx context.Context, endpoint string, handler http.Handler) {
	srv := &http.Server{
		Handler:      handler,
		Addr:         endpoint,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		IdleTimeout:  time.Second * 12,
	}

	go func() {
		<-ctx.Done()
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
