package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicantStorage) InsertApplicantBankAccount(tx *sqlx.Tx, model models.ApplicantBankAccountModel) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `
		INSERT INTO applicants_bank_data (
			iban,
			bic,
			bank_name,
			account_holder
		) VALUES
		 ($1,$2,$3,$4)
		 RETURNING id
	`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			model.Iban,
			model.Bic,
			model.BankName,
			model.AccountHolder,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			model.Iban,
			model.Bic,
			model.BankName,
			model.AccountHolder,
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application bank information", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
