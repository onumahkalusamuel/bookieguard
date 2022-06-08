package internal

import (
	"github.com/itrepablik/itrlog"
)

func init() {
	// Custom settings to initialize the itrlog.
	itrlog.SetLogInit(50, 90, "logs", "bg_log_")
}

func Logger() *itrlog.ITRLogger {
	return &itrlog.ITRLogger{}
}
