package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsService) InsertApplication(tx *sqlx.Tx, model models.ApplicationInsertModel) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertApplication(tx, model)
}
