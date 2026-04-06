package logger

import (
	"log/slog"
	"os"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/lmittmann/tint"
)

var (
	loggerInstances = loggerInstancesStruct{
		terminal: nil,
		grafana:  nil,
	}
)

type loggerInstancesStruct struct {
	terminal *slog.Logger
	grafana  *slog.Logger
}

/*

	GRAFANA LOKI PROJECT IS CURRENTLY BROKEN AND THEREFORE DISABLED
	to be able to build the project, all loki references are commented out

*/

func NewLogger() {
	stdErr := os.Stderr
	// var grafanaLogger *slog.Logger

	stdLogger := slog.New(tint.NewHandler(stdErr, &tint.Options{
		Level: slog.LevelDebug,
	}))

	// if configuration.Logger.Loki.Enabled {
	// 	config, _ := loki.NewDefaultConfig(configuration.Logger.Loki.URL)
	// 	config.TenantID = "xyz"
	// 	client, _ := loki.New(config)

	// 	grafanaLogger = slog.New(slogloki.Option{
	// 		Level:  slog.LevelDebug,
	// 		Client: client,
	// 	}.NewLokiHandler(),
	// 	)
	// }

	// grafanaLogger = grafanaLogger.With(slog.String("application", configuration.CONST_APPLICATION_SERVICE_NAME))
	stdLogger = stdLogger.With(slog.String("application", configuration.CONST_APPLICATION_SERVICE_NAME))

	// if configuration.Logger.Loki.Enabled {
	// 	loggerInstances.grafana = grafanaLogger
	// }

	loggerInstances.terminal = stdLogger
}
