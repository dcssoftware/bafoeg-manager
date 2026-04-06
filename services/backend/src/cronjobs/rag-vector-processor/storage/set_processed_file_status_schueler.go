package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *CronjobRagVectorProcessorStorage) SetProcessedFileStatusSchüler(tx *sqlx.Tx, id uuid.UUID, errorlog string) customerrors.ErrorInterface {

	var err error
	sqlquery := `UPDATE pgvector_rag_schuelerbafoeg_files SET processed_timestamp = now(), processed_error = $2 WHERE id = $1 RETURNING id`

	if tx != nil {
		_, err = tx.Exec(sqlquery, id, errorlog)
	} else {
		_, err = s.db.Exec(sqlquery, id, errorlog)
	}

	if err != nil {
		sqlErrorData := customerrors.SQLData{
			"id":       id,
			"errorlog": errorlog,
		}

		switch err {
		case sql.ErrNoRows:
			return customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)
		default:
			return customerrors.NewDatabaseError(err, "", "Cannot update processed file status by file id", sqlquery, sqlErrorData)
		}
	}

	return nil
}
