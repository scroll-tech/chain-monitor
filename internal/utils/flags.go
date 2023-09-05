package utils

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	// CommonFlags is used for app common flags in different modules
	CommonFlags = []cli.Flag{
		&ConfigFileFlag,

		&httpEnabledFlag,
		&httpListenAddrFlag,
		&httpPortFlag,

		&verbosityFlag,
		&logFileFlag,
		&logJSONFormat,
		&logDebugFlag,

		&metricsEnabled,
		&metricsAddr,
		&metricsPort,

		&pprofFlag,
		&pprofPortFlag,
		&pprofAddrFlag,
	}
	// ConfigFileFlag load json type config file.
	ConfigFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "JSON configuration file",
		Value: "./config.json",
	}

	// httpEnabledFlag enable rpc server.
	httpEnabledFlag = cli.BoolFlag{
		Name:  "http",
		Usage: "Enable the HTTP-RPC server",
		Value: false,
	}
	// httpListenAddrFlag set the http address.
	httpListenAddrFlag = cli.StringFlag{
		Name:  "http.addr",
		Usage: "HTTP-RPC server listening interface",
		Value: "localhost",
	}
	// httpPortFlag set http.port.
	httpPortFlag = cli.IntFlag{
		Name:  "http.port",
		Usage: "HTTP-RPC server listening port",
		Value: 8750,
	}

	// verbosityFlag log level.
	verbosityFlag = cli.IntFlag{
		Name:  "verbosity",
		Usage: "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail",
		Value: 3,
	}
	// logFileFlag decides where the logger output is sent. If this flag is left
	// empty, it will log to stdout.
	logFileFlag = cli.StringFlag{
		Name:  "log.file",
		Usage: "Tells the module where to write log entries",
	}
	// logJSONFormat decides the log format is json or not
	logJSONFormat = cli.BoolFlag{
		Name:  "log.json",
		Usage: "Tells the module whether log format is json or not",
		Value: true,
	}
	// logDebugFlag make log messages with call-site location
	logDebugFlag = cli.BoolFlag{
		Name:  "log.debug",
		Usage: "Prepends log messages with call-site location (file and line number)",
	}

	// metricsEnabled enable metrics collection and reporting
	metricsEnabled = cli.BoolFlag{
		Name:     "metrics",
		Usage:    "Enable metrics collection and reporting",
		Category: "METRICS",
		Value:    false,
	}
	// metricsAddr is listening address of Metrics reporting server
	metricsAddr = cli.StringFlag{
		Name:     "metrics.addr",
		Usage:    "Metrics reporting server listening address",
		Category: "METRICS",
		Value:    "127.0.0.1",
	}
	// metricsPort is listening port of Metrics reporting server
	metricsPort = cli.IntFlag{
		Name:     "metrics.port",
		Usage:    "Metrics reporting server listening port",
		Category: "METRICS",
		Value:    6060,
	}

	pprofFlag = cli.BoolFlag{
		Name:  "pprof",
		Usage: "Enable the pprof HTTP server",
	}
	pprofPortFlag = cli.IntFlag{
		Name:  "pprof.port",
		Usage: "pprof HTTP server listening port",
		Value: 6060,
	}
	pprofAddrFlag = cli.StringFlag{
		Name:  "pprof.addr",
		Usage: "pprof HTTP server listening interface",
		Value: "127.0.0.1",
	}
)

// BeforeAction is an initialization function to be executed before the main action.
// It sets up logging based on the provided context flags and, if enabled, starts the pprof server.
func BeforeAction(ctx *cli.Context) error {
	// Init logger.
	err := LogSetup(ctx)
	if err != nil {
		return err
	}

	// Start pprof
	if ctx.Bool(pprofFlag.Name) {
		address := fmt.Sprintf("%s:%d", ctx.String(pprofAddrFlag.Name), ctx.Int(pprofPortFlag.Name))
		StartPProf(address, false)
	}

	return nil
}
