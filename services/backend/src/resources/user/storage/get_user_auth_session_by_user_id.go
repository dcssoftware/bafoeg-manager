package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *UserStore) GetUserAuthSessionByUserID(tx *sqlx.Tx, userID string) (*serviceModel.SessionModel, customerrors.ErrorInterface) {
	sqlquery := "SELECT user_sessions.* FROM user_sessions JOIN users ON user_sessions.user_id = users.id WHERE users.username = $1 LIMIT 1;"
	var session storageModel.SessionModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, userID).StructScan(&session)
	} else {
		err = s.db.QueryRowx(sqlquery, userID).StructScan(&session)
	}

	if err != nil {
		data := customerrors.SQLData{
			"userID": userID,
		}

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get session", sqlquery, data)
		}
	}

	return session.ToServiceModel(), nil
}
