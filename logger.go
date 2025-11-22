package logger

import "log/slog"

func WarnOnFunctionError(f func() error, message ...string) {
	if err := f(); err != nil {
		slog.Warn("logger.WarnIfErr: error encountered", "error", err, "message", message)
	}
}
