package logger

import (
	"log/slog"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
)

func Info(
	uuidIdentifier string,
	devMessage string,
	extraContext any,
) {

	callerFile, callerLine, _ := runtime.GetCallerFunction(1)

	loggerConfig := []*slog.Logger{loggerInstances.terminal}
	if loggerInstances.grafana != nil {
		loggerConfig = append(loggerConfig, loggerInstances.grafana)
	}

	logInfo(
		loggerConfig,
		devMessage,
		slog.String("id", uuidIdentifier),
		slog.String("infologtype", "general-type"),
		slog.String("file", callerFile),
		slog.Int("file_lineNumber", callerLine),
		slog.Group("data", extraContext),
	)
}

func InfoWithCustomLocation(
	uuidIdentifier string,
	devMessage string,
	occurredFile string,
	occurredLineNumber int,
	extraContext any,
) {

	loggerConfig := []*slog.Logger{loggerInstances.terminal}
	if loggerInstances.grafana != nil {
		loggerConfig = append(loggerConfig, loggerInstances.grafana)
	}

	logInfo(
		loggerConfig,
		devMessage,
		slog.String("id", uuidIdentifier),
		slog.String("infologtype", "general-type"),
		slog.String("file", occurredFile),
		slog.Int("lineNumber", occurredLineNumber),
		slog.Group("data", extraContext),
	)
}

func logInfo(loggers []*slog.Logger, message string, args ...any) {

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
