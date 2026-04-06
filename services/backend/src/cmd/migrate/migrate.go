/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/dcssoftware/bafoeg-manager/src/helper/branding"
	"github.com/dcssoftware/bafoeg-manager/src/migrator"
	gomigrator "github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate-db",
		Short: "🚀 Migrate your database to a newer version",
		Long:  `🚀 Migrate your database to a newer version`,
		Run: func(cmd *cobra.Command, args []string) {

			branding.PrintBranding()
			migratorInstance := migrator.NewMigrator()
			migrateErr := migratorInstance.MigrateUp()

			messageStyle := lipgloss.NewStyle().Bold(true)
			successMessageStyle := messageStyle.Foreground(lipgloss.Color("#1eb523"))
			errorMessageStyle := messageStyle.Foreground(lipgloss.Color("#c42f2d"))

			if migrateErr != nil {

				var isExitErrCode bool = true
				var errMsg string

				switch migrateErr {
				case gomigrator.ErrNoChange:
					isExitErrCode = false
					errMsg = "Database is on the newest version"
				case gomigrator.ErrNilVersion:
					errMsg = "No Migration found."
				case gomigrator.ErrInvalidVersion:
					errMsg = "Database is on a newer version than your software. Please use the newer version"
				case gomigrator.ErrLocked:
					errMsg = "Database is currently locked by another application"
				case gomigrator.ErrLockTimeout:
					errMsg = "Database Migration timed out"
				default:
					errMsg = migrateErr.Error()
				}

				log.Println(errorMessageStyle.Render(errMsg))

				if isExitErrCode {
					os.Exit(1)
				}
			}

			log.Println(successMessageStyle.Render("Database Migration was successful"))
			os.Exit(0)
		},
	}
}
