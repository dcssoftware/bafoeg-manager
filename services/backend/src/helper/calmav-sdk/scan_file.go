package calmavsdk

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	clamavrestsdk "github.com/dcssoftware/clamav-rest-go-sdk"
)

func ScanFile(fileContent []byte) (bool, customerrors.ErrorInterface) {

	// if no virus check is desired, disable it
	if !configuration.ClamAVWrapper.Enabled {
		return false, nil
	}

	instance := clamavrestsdk.NewclamAVRestInstance(
		configuration.ClamAVWrapper.Address,
		int(configuration.ClamAVWrapper.Port),
		configuration.ClamAVWrapper.IsSecure,
	)

	result, resultErr := instance.ScanFile(fileContent)
	if resultErr != nil {
		return false, customerrors.NewVirusScannerError(resultErr)
	}

	return result, nil
}
