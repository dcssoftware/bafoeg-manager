package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *RAGStorage) DeleteRagFileStudierenden(tx *sqlx.Tx, documentID uuid.UUID) customerrors.ErrorInterface {
	sqlquery := `DELETE FROM pgvector_rag_studierendenbafoeg_files WHERE id::text = $1`

	var err error
	if tx != nil {
		_, err = tx.Exec(sqlquery, documentID.String())
	} else {
		_, err = s.db.Exec(sqlquery, documentID)
	}

	sqlErrorData := customerrors.SQLData{
		"cmetadata->>'originDocumentID'": documentID.String(),
	}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return customerrors.NewDatabaseError(err, "", "Cannot delete rag vectors for schueler by document id", sqlquery, sqlErrorData)
		}
	}

	return nil
}
