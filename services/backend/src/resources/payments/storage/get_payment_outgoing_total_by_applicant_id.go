package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/payments/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *PaymentStorage) GetPaymentOutgoingTotalByApplicantID(tx *sqlx.Tx, applicantID string) (float64, customerrors.ErrorInterface) {
	sqlquery := `
			SELECT COALESCE(SUM(amount), 0) AS amount FROM payments 
			WHERE applicant_id = $1
	`

	var result models.Amount
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, applicantID).StructScan(&result)
	} else {
		err = s.db.QueryRowx(sqlquery, applicantID).StructScan(&result)
	}

	if err != nil {
		data := customerrors.SQLData{
			"es_user_id": applicantID,
		}

		switch err {

		case sql.ErrNoRows:
			return 0, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get payment outgoing total amount", sqlquery, data)
		}
	}

	return result.Amount, nil
}
