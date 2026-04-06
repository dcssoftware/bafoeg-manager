package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/ccoveille/go-safecast/v2"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *RAGService) UploadRAGrelevantDocumentsPDStudierende(tx *sqlx.Tx, filename, filetype string, filesize int64, filecontent []byte, createdFromUserID string) error {

	var commitTransaction bool = false
	var txErr error

	if tx == nil {

		tx, txErr = s.storage.StartTx()
		if txErr != nil {
			return txErr
		}

		commitTransaction = true

	}

	fileSizeUint, fileSizeErr := safecast.Convert[uint](filesize)
	if fileSizeErr != nil {
		return customerrors.NewInternalServerError(fileSizeErr, "", "")
	}

	hash := md5.New()
	if _, err := io.Copy(hash, bytes.NewBuffer(filecontent)); err != nil {
		return customerrors.NewInternalServerError(err, "", "could not generate file hash")
	}
	fileHash := hex.EncodeToString(hash.Sum(nil))

	// insert in postgres as reference
	// mark file in database as processable for ai
	documentID, documentIDErr := s.storage.InsertDocumentStudierende(tx, filename, filetype, fileSizeUint, fileHash, createdFromUserID)
	if documentIDErr != nil {
		return documentIDErr
	}
	// store file in s3
	s.storageS3.UploadDocumentStudierenden(documentID.String(), filecontent)

	// commit transacion
	if commitTransaction {
		commitErr := tx.Commit()
		if commitErr != nil {
			return commitErr
		}
	}

	return nil
}
