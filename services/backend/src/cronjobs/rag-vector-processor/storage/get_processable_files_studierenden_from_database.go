package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor/storage/models"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *CronjobRagVectorProcessorStorage) GetProcessableFilesStudierendenFromDatabase(tx *sqlx.Tx) ([]serviceModel.DocumentModel, customerrors.ErrorInterface) {
	limit := 50

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("pgvector_rag_studierendenbafoeg_files_overview").
		Where(squirrel.Eq{"processed_timestamp": nil}).
		PlaceholderFormat(squirrel.Dollar).
		Limit(uint64(limit))

	sqlquerybuilder.ToSql()

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var documents []serviceModel.DocumentModel
	var err error

	if tx != nil {
		rows, err = tx.Queryx(sqlquery, sqlArgs...)
	} else {
		rows, err = s.db.Queryx(sqlquery, sqlArgs...)
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
		var file models.DocumentModel
		if err := rows.StructScan(&file); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		documents = append(documents, *file.ToServiceFileModel())
	}

	return documents, nil
}
