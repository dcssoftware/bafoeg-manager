package cmd

import (
	"log"

	// "time"

	"github.com/spf13/cobra"

	// "github.com/spf13/cobra/doc"
	migrate "github.com/dcssoftware/bafoeg-manager/src/cmd/migrate"
	runApp "github.com/dcssoftware/bafoeg-manager/src/cmd/run"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/branding"
	webApp "github.com/dcssoftware/bafoeg-manager/src/web-app"
)

// runCmd represents the run command
func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "application",
		Short: configuration.CONST_APPLICATION_BRANDING_HEADER,
		Long:  configuration.CONST_APPLICATION_BRANDING_HEADER,
		Run: func(cmd *cobra.Command, args []string) {
			// branding.PrintBranding()
			// _ = cmd.Help()
			// os.Exit(0)

			branding.PrintBrandingWithConfig()
			webApp.NewApp().RunApp()
		},
	}
}
func Execute() {
	runCmd := GetCommand()
	runCmd.AddCommand(runApp.GetCommand())
	runCmd.AddCommand(migrate.GetCommand())
	// runCmd.AddCommand(version.VersionCmd)

	err := runCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
