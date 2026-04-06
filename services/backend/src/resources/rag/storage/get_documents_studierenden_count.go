package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *RAGStorage) GetDocumentsStudierendenCount(tx *sqlx.Tx, filterResult string) (uint, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.
		Select(`COUNT(id) AS count`).
		From("pgvector_rag_studierendenbafoeg_files_overview").
		Where(squirrel.Like{"file_name": "%" + filterResult + "%"}).
		PlaceholderFormat(squirrel.Dollar)

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return 0, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var model storageModel.CountModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, sqlArgs...).StructScan(&model)
	} else {
		err = s.db.QueryRowx(sqlquery, sqlArgs...).StructScan(&model)
	}

	if err != nil {
		data := customerrors.SQLData{}

		switch err {

		case sql.ErrNoRows:
			return 0, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get pgvector_rag_schuelerbafoeg_files_overview files count", sqlquery, data)
		}
	}

	return model.Count, nil
}
