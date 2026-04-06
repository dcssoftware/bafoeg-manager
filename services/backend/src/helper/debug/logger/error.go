package logger

import (
	"log/slog"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

func Error(
	uuidIdentifier *uuid.UUID,
	devMessage string,
	goError error,
	extraContext string,
) {

	loggerConfig := []*slog.Logger{loggerInstances.terminal}
	if loggerInstances.grafana != nil {
		loggerConfig = append(loggerConfig, loggerInstances.grafana)
	}

	if uuidIdentifier == nil {
		newUuid := uuid.New()
		uuidIdentifier = &newUuid
	}

	callerFile, callerLine, _ := runtime.GetCallerFunction(1)

	logError(
		loggerConfig,
		devMessage,
		slog.String("id", uuidIdentifier.String()),
		slog.String("file", callerFile),
		slog.Int("lineNumber", callerLine),
		slog.String("go-error", goError.Error()),
		slog.String("data", extraContext),
	)
}

func ErrorWithCustomLocation(
	uuidIdentifier *uuid.UUID,
	devMessage string,
	occurredFile string,
	occurredLineNumber int,
	goError error,
	extraContext string,
) {

	loggerConfig := []*slog.Logger{loggerInstances.terminal}
	if loggerInstances.grafana != nil {
		loggerConfig = append(loggerConfig, loggerInstances.grafana)
	}

	if uuidIdentifier == nil {
		newUuid := uuid.New()
		uuidIdentifier = &newUuid
	}

	logError(
		loggerConfig,
		devMessage,
		slog.String("id", uuidIdentifier.String()),
		slog.String("file", occurredFile),
		slog.Int("file_lineNumber", occurredLineNumber),
		slog.String("go-error", goError.Error()),
		slog.String("data", extraContext),
	)
}

func logError(loggers []*slog.Logger, message string, args ...any) {

	for _, logger := range loggers {
		if logger == nil {
			panic("could not load logger instance for logging")
		}

		logger.Error(
			message,
			args...,
		)

	}
}
