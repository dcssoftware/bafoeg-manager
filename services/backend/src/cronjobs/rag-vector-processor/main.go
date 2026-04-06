package ragvectorprocessor

import (
	"context"
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor/service"
	"github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor/storage"
	storages3 "github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor/storageS3"
	"github.com/go-sqlx/sqlx"
	"github.com/minio/minio-go/v7"
)

type RagVectorProcessor struct {
	db *sqlx.DB
	s3 *minio.Client
}

func NewCronjob(db *sqlx.DB, s3 *minio.Client) *RagVectorProcessor {
	return &RagVectorProcessor{
		db: db,
		s3: s3,
	}
}

func (c *RagVectorProcessor) RunCronjob() {
	ctx, ctxCancelFunc := context.WithTimeout(
		context.Background(),
		configuration.Conjobs.RagVectorProcessor.Timeout,
	)

	defer ctxCancelFunc()

	fmt.Println("---- CRONJOB -----")

	storeS3 := storages3.NewCronjobRagVectorProcessorStorageS3(c.s3)
	store := storage.NewCronjobRagVectorProcessorStorage(c.db)
	svc := service.NewCronjobRagVectorProcessorService(store, storeS3)

	err := svc.ProcessFiles(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---- CRONJOB Ende -----")
}
