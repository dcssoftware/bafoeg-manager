package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	exportapplication "github.com/dcssoftware/bafoeg-manager/src/helper/files/export-application"
)

func (h *ApplicationsService) DownloadDigitalApplicationPDF() ([]byte, customerrors.ErrorInterface) {
	return exportapplication.ExportPdfApplication()
}
