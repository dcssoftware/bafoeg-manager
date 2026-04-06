package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicantStorage) InsertApplicant(tx *sqlx.Tx, model serviceModels.ApplicantInsertModel) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `
		INSERT INTO applicants (
			firstname,
			lastname,
			contact_id,
			address_id,
			bank_account_id
		) VALUES
		 ($1,$2,$3,$4,$5)
		 RETURNING id
	`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			model.Firstname,
			model.Lastname,
			models.UUIDToStringPtr(model.ContactID),
			models.UUIDToStringPtr(model.AddressID),
			models.UUIDToStringPtr(model.BankAccountID),
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			model.Firstname,
			model.Lastname,
			models.UUIDToStringPtr(model.ContactID),
			models.UUIDToStringPtr(model.AddressID),
			models.UUIDToStringPtr(model.BankAccountID),
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
