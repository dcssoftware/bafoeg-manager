/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/branding"
	webApp "github.com/dcssoftware/bafoeg-manager/src/web-app"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "starts the webserver to provide the application",
		Long:  `starts the webserver to provide the application`,

		Run: func(cmd *cobra.Command, args []string) {
			branding.PrintBrandingWithConfig()

			webApp.NewApp().RunApp()
		},
	}
}
