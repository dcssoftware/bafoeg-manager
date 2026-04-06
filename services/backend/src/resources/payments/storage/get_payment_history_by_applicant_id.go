package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModell "github.com/dcssoftware/bafoeg-manager/src/resources/payments/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/payments/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *PaymentStorage) GetPaymentHistoryByApplicantID(tx *sqlx.Tx, page uint, applicantID string) ([]serviceModell.PaymentsModell, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquery := `
	SELECT payments.*, payment_status.identifier AS status_identifier FROM payments 
	INNER JOIN payment_status ON payment_status.id = payments.status_id
	WHERE applicant_id = $1
	LIMIT $2
	OFFSET $3
	`

	var rows *sqlx.Rows
	var payments []models.PaymentsModell
	var err error

	if tx != nil {
		rows, err = tx.Queryx(sqlquery, applicantID, limit, offset)
	} else {
		rows, err = s.db.Queryx(sqlquery, applicantID, limit, offset)
	}

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get payment history", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var payment models.PaymentsModell
		if err := rows.StructScan(&payment); err != nil {
			customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
			panic(err)
		}
		payments = append(payments, payment)
	}

	return models.ToPaymentsServiceModels(payments), nil
}
