package storages3

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/minio/minio-go/v7"
)

func (s *RAGStorageS3) UploadDocumentSchüler(fileID string, fileContent []byte) customerrors.ErrorInterface {
	reader := bytes.NewReader(fileContent)

	_, objErr := s.s3.PutObject(
		context.Background(),
		configuration.S3StoragePaths.RAGDataBucket,
		fmt.Sprintf("/rag-data/schueler-data/uploaded-documents/%s", fileID),
		reader,
		int64(len(fileContent)),
		minio.PutObjectOptions{},
	)
	if objErr != nil {
		log.Fatalln(objErr)
	}

	return nil
}
