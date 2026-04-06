package branding

import (
	"image"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/dcssoftware/bafoeg-manager/src/static/assets"
	"github.com/qeesung/image2ascii/convert"
)

func getASCIILogo() (string, error) {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 80
	convertOptions.FixedHeight = 35

	// Create the image converter
	converter := convert.NewImageConverter()

	// Open the image file
	file, err := assets.TerminalFS.Open("terminal-assets/logo.png")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Decode the image file
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Convert the image to an ASCII matrix and print it
	asciiImage := converter.Image2ASCIIString(img, &convertOptions)

	lipglossWidth := 100
	lipglossMargin := 5

	var styleImage = lipgloss.NewStyle().
		Width(lipglossWidth).
		Align(lipgloss.Center).
		MarginLeft(lipglossMargin).
		MarginRight(lipglossMargin)

	return styleImage.Render(asciiImage), nil
}
