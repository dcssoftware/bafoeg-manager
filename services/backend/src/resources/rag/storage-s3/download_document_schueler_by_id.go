package storages3

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/minio/minio-go/v7"
)

func (s *RAGStorageS3) DownloadDocumentSchülerByID(fileID string) ([]byte, customerrors.ErrorInterface) {
	obj, objErr := s.s3.GetObject(
		context.Background(),
		configuration.S3StoragePaths.RAGDataBucket,
		fmt.Sprintf("/rag-data/schueler-data/uploaded-documents/%s", fileID),
		minio.GetObjectOptions{},
	)
	if objErr != nil {
		log.Fatalln(objErr)
	}

	content, contentErr := io.ReadAll(obj)
	if contentErr != nil {
		log.Fatalln(contentErr)
	}

	return content, nil
}
