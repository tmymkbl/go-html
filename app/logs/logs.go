package logs

import (
	"go-sample-todo/config"
	"io"
	"log/slog"
	"os"
)

var Logfile *os.File
var Log *slog.Logger

func SetLogfile() {
	f := io.Writer(Logfile)
	setLog(f)
}

func SetLogStd() {
	f := os.Stdout
	setLog(f)
}

func SetLogMulti() {
	f := io.MultiWriter(os.Stdout, Logfile)
	setLog(f)
}

func setLog(f io.Writer) {
	switch config.Config.LogLevel {
	case "debug":
		Log = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))
		Log.Debug("debug Logging start", "output", config.Config.LogOutput, "file", config.Config.LogFile)
	case "info":
		Log = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo}))
		Log.Info("info Logging start", "output", config.Config.LogOutput, "file", config.Config.LogFile)
	case "warn":
		Log = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelWarn}))
		Log.Warn("warn Logging start", "output", config.Config.LogOutput, "file", config.Config.LogFile)
	case "error":
		Log = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelError}))
		Log.Error("error Logging start", "output", config.Config.LogOutput, "file", config.Config.LogFile)
	default:
		Log = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
}
