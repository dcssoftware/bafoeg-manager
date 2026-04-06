package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsService) InsertApplicationFromEakte(tx *sqlx.Tx, model models.ApplicationFromEakteInsertModel) (applicationID uuid.UUID, revisionID uuid.UUID, mappingID uuid.UUID, err customerrors.ErrorInterface) {

	var txStarted bool
	if tx == nil {
		var txErr error
		tx, txErr = s.storage.StartTx()
		if txErr != nil {
			err = customerrors.NewDatabaseTransactionBeginError(txErr, "")
			return
		}
		txStarted = true
	}

	applicationID, err = s.storage.InsertApplicationFromEakte(tx, model)
	if err != nil {
		tx.Rollback()
		return
	}

	revisionID, err = s.InsertApplicationRevision(
		tx,
		applicationID.String(),
		"Antrag aus eAkte erstellt",
		"Der Antrag wurde automatisch aus der eAkte erstellt.",
		nil,
	)
	if err != nil {
		tx.Rollback()
		return
	}

	mappingID, err = s.InsertApplicationEakteMapping(tx, applicationID, model.EakteAkteID)
	if err != nil {
		tx.Rollback()
		return
	}

	if txStarted {
		commitErr := tx.Commit()
		if commitErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = customerrors.NewInternalServerError(rollbackErr, "", "")
				return
			}
			err = customerrors.NewDatabaseTransactionCommitError(commitErr, "")
			return
		}
	}

	return
}
