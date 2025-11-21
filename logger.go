package logger

import "log/slog"

func WarnOnError(err error, message ...string) {
	if err != nil {
		slog.Warn("logger.WarnIfErr: error encountered", "error", err, "message", message)
	}
}
