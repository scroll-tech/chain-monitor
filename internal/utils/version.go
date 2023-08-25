package utils

import (
	"fmt"
	"runtime/debug"
)

var tag = "v4.1.61"

var commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				value := setting.Value
				if len(value) >= 7 {
					return value[:7]
				}
				return value
			}
		}
	}
	// Set default value for integration test.
	return "000000"
}()

// Version denote the version of scroll protocol, including the l2geth, relayer, coordinator, prover, contracts and etc.
var Version = fmt.Sprintf("%s-%s", tag, commit)
