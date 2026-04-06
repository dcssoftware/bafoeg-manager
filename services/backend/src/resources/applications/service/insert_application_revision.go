package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsService) InsertApplicationRevision(tx *sqlx.Tx, applicationID, header, description string, newData *models.ApplicationRevisionDataModel) (uuid.UUID, customerrors.ErrorInterface) {
	var baseRevision *models.ApplicationRevisionModel

	existingRevisions, existingRevisionsErr := s.storage.GetApplicationRevisionsByApplicationIDCount(tx, applicationID)
	if existingRevisionsErr != nil {
		return uuid.Nil, existingRevisionsErr
	}

	if existingRevisions > 0 {
		var err customerrors.ErrorInterface
		baseRevision, err = s.storage.GetApplicationRevisionLatestRevisionByApplicationID(tx, applicationID)
		if err != nil {
			return uuid.Nil, err
		}
	}

	s.storage.InsertApplicationRevision(
		tx,
		applicationID,
		header,
		description,
		baseRevision.ToDataModel(),
		newData,
	)

	return uuid.Nil, nil
}
