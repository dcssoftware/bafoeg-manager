package storage

// func (s *EakteStorage) InsertEakteFile(tx *sqlx.Tx, vorgangID, source, fileID uuid.UUID, isXdomeaSourceFile bool) (uuid.UUID, customerrors.ErrorInterface) {
// 	sqlquery := `
// 		INSERT INTO eakte_import_dokument (source, vorgang_id, source_xdomea_file, file_id) VALUES ($1, $2, $3, $4) RETURNING id
// 	`

// 	var result storageModels.IDModel
// 	var row *sqlx.Row
// 	if tx != nil {
// 		row = tx.QueryRowx(
// 			sqlquery,
// 			vorgangID,
// 			source,
// 			isXdomeaSourceFile,
// 			fileID,
// 		)
// 	} else {
// 		row = s.db.QueryRowx(
// 			sqlquery,
// 			vorgangID,
// 			source,
// 			isXdomeaSourceFile,
// 			fileID,
// 		)
// 	}

// 	err := row.StructScan(&result)

// 	sqlErrorData := customerrors.SQLData{}
// 	if err != nil {
// 		switch err {

// 		case sql.ErrNoRows:
// 			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

// 		default:
// 			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert eakte file", sqlquery, sqlErrorData)
// 		}
// 	}

// 	return result.ID, nil
// }
