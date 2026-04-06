package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	servicemodels "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *RAGStorage) GetDocumentStudierendenByID(tx *sqlx.Tx, fileID string) (*servicemodels.DocumentModel, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.
		Select(`*`).
		From("pgvector_rag_studierendenbafoeg_files_overview").
		Where(squirrel.Eq{"id": fileID}).
		PlaceholderFormat(squirrel.Dollar).
		Limit(1)

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, sqlArgs...)
	} else {
		row = s.db.QueryRowx(sqlquery, sqlArgs...)
	}

	var fileModel models.DocumentModel
	err = row.StructScan(&fileModel)

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get application file by file id", sqlquery, sqlErrorData)
		}
	}

	return fileModel.ToServiceFileModel(), nil
}
