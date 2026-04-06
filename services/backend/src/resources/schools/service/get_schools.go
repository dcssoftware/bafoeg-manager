package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolService) GetSchools(tx *sqlx.Tx, page uint, filter string) ([]model.SchoolShortModel, customerrors.ErrorInterface) {
	return s.storage.GetSchools(tx, page, filter)
}
