package service

import (
	"errors"
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/metadata"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/mimetype"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/webp"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (s *FileService) InsertFileProfilePicture(tx *sqlx.Tx, userID string, fileContent []byte) (uuid.UUID, *minio.UploadInfo, customerrors.ErrorInterface) {

	fileContent, convertErr := convertProfilePictureToWebP(fileContent)
	if convertErr != nil {
		return uuid.Nil, nil, convertErr
	}

	fileSize, mimeType, fileHash, fileMetadataErr := metadata.GetFileMetaData(fileContent)
	if fileMetadataErr != nil {
		return uuid.Nil, nil, fileMetadataErr
	}

	filename := "profilepicture.webp"
	fileID, fileErr := s.storage.InsertFile(tx, filename, mimeType, fileSize, fileHash)
	if fileErr != nil {
		return uuid.Nil, nil, fileErr
	}

	storagePath := fmt.Sprintf("/profile-pictures/%s/", userID)
	uploadInfo, uploadErr := s.storageS3.UploadFile(
		configuration.S3StoragePaths.EAktenDataBucket,
		storagePath,
		filename,
		fileContent,
	)

	return fileID, uploadInfo, uploadErr
}

func convertProfilePictureToWebP(fileContent []byte) ([]byte, customerrors.ErrorInterface) {
	var convertErr customerrors.ErrorInterface

	mimeType, mimeTypeErr := mimetype.DetectMineType(fileContent)
	if mimeTypeErr != nil {
		return nil, mimeTypeErr
	}

	switch mimeType {

	case "image/jpg", "image/jpeg":

		fileContent, convertErr = webp.ConvertJPGToWebP(fileContent)
		if convertErr != nil {
			return nil, convertErr
		}

	case "image/png":

		fileContent, convertErr = webp.ConvertPngToWebP(fileContent)
		if convertErr != nil {
			return nil, convertErr
		}

	default:
		return nil, customerrors.NewInternalServerError(errors.New(""), "unsupported image format", "")
	}

	return fileContent, nil
}

// func (s *FileService) InsertFile(tx *sqlx.Tx, bucketname, storagePath, fileName string, fileContent []byte) (uuid.UUID, customerrors.ErrorInterface) {

// 	fileSize, mimeType, fileHash, err := metadata.GetFileMetaData(fileContent)

// 	fileID, err := s.storage.InsertFile(tx, fileName, mimeType, fileSize, fileHash)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}

// 	_, err = s.storageS3.UploadFile(
// 		bucketname,
// 		storagePath,
// 		"profilepicture.webp",
// 		fileContent,
// 	)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}

// 	return fileID, nil
// }
