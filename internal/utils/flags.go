package utils

import (
	"github.com/urfave/cli/v2"
)

var (
	// CommonFlags is used for app common flags in different modules
	CommonFlags = []cli.Flag{
		&ConfigFileFlag,

		&HTTPEnabledFlag,
		&HTTPListenAddrFlag,
		&HTTPPortFlag,

		&verbosityFlag,
		&logFileFlag,
		&logJSONFormat,
		&logDebugFlag,

		&MetricsEnabled,
		&MetricsAddr,
		&MetricsPort,
	}
	// ConfigFileFlag load json type config file.
	ConfigFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "JSON configuration file.",
		Value: "./config.json",
	}

	// HTTPEnabledFlag enable rpc server.
	HTTPEnabledFlag = cli.BoolFlag{
		Name:  "http",
		Usage: "Enable the HTTP-RPC server.",
		Value: false,
	}
	// HTTPListenAddrFlag set the http address.
	HTTPListenAddrFlag = cli.StringFlag{
		Name:  "http.addr",
		Usage: "HTTP-RPC server listening interface.",
		Value: "localhost",
	}
	// HTTPPortFlag set http.port.
	HTTPPortFlag = cli.IntFlag{
		Name:  "http.port",
		Usage: "HTTP-RPC server listening port.",
		Value: 8750,
	}

	// verbosityFlag log level.
	verbosityFlag = cli.IntFlag{
		Name:  "verbosity",
		Usage: "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail.",
		Value: 3,
	}
	// logFileFlag decides where the logger output is sent. If this flag is left
	// empty, it will log to stdout.
	logFileFlag = cli.StringFlag{
		Name:  "log.file",
		Usage: "Tells the module where to write log entries.",
	}
	// logJSONFormat decides the log format is json or not
	logJSONFormat = cli.BoolFlag{
		Name:  "log.json",
		Usage: "Tells the module whether log format is json or not.",
		Value: true,
	}
	// logDebugFlag make log messages with call-site location
	logDebugFlag = cli.BoolFlag{
		Name:  "log.debug",
		Usage: "Prepends log messages with call-site location (file and line number).",
	}

	// MetricsEnabled enable metrics collection and reporting
	MetricsEnabled = cli.BoolFlag{
		Name:     "metrics",
		Usage:    "Enable metrics collection and reporting.",
		Category: "METRICS",
		Value:    false,
	}
	// MetricsAddr is listening address of Metrics reporting server
	MetricsAddr = cli.StringFlag{
		Name:     "metrics.addr",
		Usage:    "Metrics reporting server listening address.",
		Category: "METRICS",
		Value:    "127.0.0.1",
	}
	// MetricsPort is listening port of Metrics reporting server
	MetricsPort = cli.IntFlag{
		Name:     "metrics.port",
		Usage:    "Metrics reporting server listening port.",
		Category: "METRICS",
		Value:    6060,
	}
)
