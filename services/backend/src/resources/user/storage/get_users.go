package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *UserStore) GetUsers(tx *sqlx.Tx) ([]serviceModel.UserModel, customerrors.ErrorInterface) {
	sqlquery := `SELECT * FROM users`

	var rows *sqlx.Rows
	var users []serviceModel.UserModel
	var err error

	if tx != nil {
		rows, err = tx.Queryx(sqlquery)
	} else {
		rows, err = s.db.Queryx(sqlquery)
	}

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get user", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var user models.UserModel
		if err := rows.StructScan(&user); err != nil {
			customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
			panic(err)
		}
		users = append(users, *user.ToServiceModel())
	}

	return users, nil
}
