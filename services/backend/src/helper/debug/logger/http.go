package logger

import (
	"log/slog"
	"time"
)

func HTTP(
	method string,
	endpoint string,
	statusCode int,
	executionTime time.Duration,
) {
	loggerConfig := []*slog.Logger{loggerInstances.terminal}
	if loggerInstances.grafana != nil {
		loggerConfig = append(loggerConfig, loggerInstances.grafana)
	}

	logHTTP(
		loggerConfig,
		"http request",
		slog.String("infologtype", "http-request"),
		slog.String("method", method),
		slog.String("endpoint", endpoint),
		slog.Int("statusCode", statusCode),
		slog.Int("executionTime", int(executionTime.Milliseconds())),
	)
}

func logHTTP(loggers []*slog.Logger, message string, args ...any) {

	for _, logger := range loggers {

		if logger == nil {
			panic("could not load logger instance for logging")
		}

		logger.Info(
			message,
			args...,
		)

	}
}
