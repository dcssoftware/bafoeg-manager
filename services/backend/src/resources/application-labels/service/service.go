package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/storage"
	"github.com/go-sqlx/sqlx"
)

type ApplicationLabelsService struct {
	storage ApplicationLabelssStorageInterface
}

func NewApplicationLabelsService(storage *storage.ApplicationLabelsStorage) *ApplicationLabelsService {
	return &ApplicationLabelsService{
		storage,
	}
}

type ApplicationLabelssStorageInterface interface {
	GetLabels(tx *sqlx.Tx, page uint) ([]models.ApplicationLabel, customerrors.ErrorInterface)
}
