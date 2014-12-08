package logger

import (
	"github.com/kdar/factorlog"
	"os"
)

var Log *factorlog.FactorLog

func init() {
	modded := `%{Color "red" "ERROR"}%{Color "yellow" "WARN"}%{Color "green" "INFO"}%{Color "cyan" "DEBUG"}%{Color "magenta" "TRACE"}[%{Date} %{Time}] [%{SEVERITY}:%{File}:%{Line}] %{Message}%{Color "reset"}`
	Log = factorlog.New(os.Stdout, factorlog.NewStdFormatter(modded))

	Log.Info("Initializing logger system")
}
