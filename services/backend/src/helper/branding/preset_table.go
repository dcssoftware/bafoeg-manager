package branding

import (
	"strconv"

	lipgloss "charm.land/lipgloss/v2"
	table "charm.land/lipgloss/v2/table"
	"github.com/dcssoftware/bafoeg-manager/src/helper/branding/models"
)

func getASCIIConfigurationTable(configuration models.ConfigurationTable, presetStle lipgloss.Style) string {
	configurationRows := [][]string{
		{"URL:", configuration.DisplayHost},
		{"", ""},
		{"Host (Container) :", configuration.Host},
		{"Port (Container) :", configuration.Port},
		{"", ""},
		{"ShowSwagger:", strconv.FormatBool(configuration.ShowSwagger)},
	}

	lipglossTableWidth := 90
	informationTable := table.New().
		Border(lipgloss.NormalBorder()).
		Headers("Key", "Value").
		Width(lipglossTableWidth).
		Rows(configurationRows...)

	lipglossTablePadding := 3
	styleInformation := presetStle.Align(lipgloss.Left).
		PaddingLeft(lipglossTablePadding).
		PaddingRight(lipglossTablePadding).
		Underline(false).
		Bold(false)

	return styleInformation.Render(informationTable.Render())
}
