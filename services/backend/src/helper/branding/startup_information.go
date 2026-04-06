package branding

import (
	lipgloss "charm.land/lipgloss/v2"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
)

func getASCIIStartupInformation() (string, lipgloss.Style) {
	lipgloassWidth := 100
	lipglossMargin := 5

	var styleInformation = lipgloss.NewStyle().
		Width(lipgloassWidth).
		Bold(true).
		Align(lipgloss.Center).
		Underline(true).
		MarginLeft(lipglossMargin).
		MarginRight(lipglossMargin)

	return styleInformation.Render(configuration.CONST_APPLICATION_BRANDING_HEADER), styleInformation
}
