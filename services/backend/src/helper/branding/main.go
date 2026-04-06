package branding

import (
	"fmt"
	"strconv"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/branding/models"
)

func PrintBranding() {
	logo, logoErr := getASCIILogo()
	if logoErr != nil {
		panic(logoErr)
	}

	startupInformation, _ := getASCIIStartupInformation()

	fmt.Println(logo)
	fmt.Println(startupInformation)
	fmt.Println()
}

func PrintBrandingWithConfig() {
	logo, logoErr := getASCIILogo()
	if logoErr != nil {
		panic(logoErr)
	}

	startupInformation, defaultConfConfiguraton := getASCIIStartupInformation()

	startupPresetTable := getASCIIConfigurationTable(models.ConfigurationTable{
		DisplayHost: configuration.Webserver.Displayname,
		Host:        configuration.Webserver.Host,
		Port:        strconv.Itoa(configuration.Webserver.Port),
		ShowSwagger: configuration.Webserver.Display.Swagger,
	}, defaultConfConfiguraton)

	fmt.Println("\n" + logo)
	fmt.Println("\n" + startupInformation)
	fmt.Println("\n" + startupPresetTable)
	fmt.Println("")
}
