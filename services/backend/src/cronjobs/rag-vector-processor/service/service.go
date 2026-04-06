package service

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

type CronjobRagVectorProcessorService struct {
	storage   CronjobRagVectorProcessorStore
	storageS3 CronjobRagVectorProcessorStoreS3
}

func NewCronjobRagVectorProcessorService(
	storage CronjobRagVectorProcessorStore,
	storageS3 CronjobRagVectorProcessorStoreS3,
) *CronjobRagVectorProcessorService {
	return &CronjobRagVectorProcessorService{
		storage:   storage,
		storageS3: storageS3,
	}
}

type CronjobRagVectorProcessorStore interface {
	GetProcessableFilesSchülerFromDatabase(tx *sqlx.Tx) ([]serviceModel.DocumentModel, customerrors.ErrorInterface)
	GetProcessableFilesStudierendenFromDatabase(tx *sqlx.Tx) ([]serviceModel.DocumentModel, customerrors.ErrorInterface)
	SetProcessedFileStatusSchüler(tx *sqlx.Tx, id uuid.UUID, errorlog string) customerrors.ErrorInterface
	SetProcessedFileStatusStudierenden(tx *sqlx.Tx, id uuid.UUID, errorlog string) customerrors.ErrorInterface
}

type CronjobRagVectorProcessorStoreS3 interface {
	DownloadDocumentSchülerByID(fileID string) ([]byte, customerrors.ErrorInterface)
	DownloadDocumentStudierendenByID(fileID string) ([]byte, customerrors.ErrorInterface)
}
