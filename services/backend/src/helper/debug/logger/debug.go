package logger

import (
	"log/slog"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

func Debug(
	uuidIdentifier string,
	devMessage string,
	extraContext any,
) {

	loggerConfig := []*slog.Logger{loggerInstances.terminal}
	if loggerInstances.grafana != nil {
		loggerConfig = append(loggerConfig, loggerInstances.grafana)
	}

	callerFile, callerLine, _ := runtime.GetCallerFunction(1)

	logDebug(
		loggerConfig,
		devMessage,
		slog.String("id", uuidIdentifier),
		slog.String("file", callerFile),
		slog.Int("lineNumber", callerLine),
		slog.Group("data", extraContext),
	)
}

func DebugWithCustomLocation(
	uuidIdentifier *uuid.UUID,
	devMessage string,
	occurredFile string,
	occurredLineNumber int,
	extraContext any,
) {

	loggerConfig := []*slog.Logger{loggerInstances.terminal}
	if loggerInstances.grafana != nil {
		loggerConfig = append(loggerConfig, loggerInstances.grafana)
	}

	if uuidIdentifier == nil {
		newUuid := uuid.New()
		uuidIdentifier = &newUuid
	}

	logDebug(
		loggerConfig,
		devMessage,
		slog.String("id", uuidIdentifier.String()),
		slog.String("file", occurredFile),
		slog.Int("file_lineNumber", occurredLineNumber),
		slog.Group("data", extraContext),
	)
}

func logDebug(loggers []*slog.Logger, message string, args ...any) {

	for _, logger := range loggers {
		if logger == nil {
			panic("could not load logger instance for logging")
		}

		logger.Debug(
			message,
			args...,
		)

	}
}
