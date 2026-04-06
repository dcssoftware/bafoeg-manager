package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

type RAGService struct {
	storage   RAGStorage
	storageS3 RAGStorageS3
}

func NewRAGService(storage RAGStorage, storageS3 RAGStorageS3) *RAGService {
	return &RAGService{
		storage:   storage,
		storageS3: storageS3,
	}
}

type RAGStorage interface {
	StartTx() (*sqlx.Tx, error)

	GetDocumentsSchüler(tx *sqlx.Tx, page uint, filterResult string) ([]models.DocumentModel, customerrors.ErrorInterface)
	GetDocumentsStudierenden(tx *sqlx.Tx, page uint, filterResult string) ([]models.DocumentModel, customerrors.ErrorInterface)

	GetDocumentSchülerByID(tx *sqlx.Tx, id string) (*models.DocumentModel, customerrors.ErrorInterface)
	GetDocumentsSchülerCount(tx *sqlx.Tx, filterResult string) (uint, customerrors.ErrorInterface)
	GetDocumentStudierendenByID(tx *sqlx.Tx, id string) (*models.DocumentModel, customerrors.ErrorInterface)
	GetDocumentsStudierendenCount(tx *sqlx.Tx, filterResult string) (uint, customerrors.ErrorInterface)

	InsertDocumentSchüler(tx *sqlx.Tx, fileName, fileType string, fileSize uint, fileHash string, createdFromUserID string) (uuid.UUID, customerrors.ErrorInterface)
	InsertDocumentStudierende(tx *sqlx.Tx, fileName, fileType string, fileSize uint, fileHash string, createdFromUserID string) (uuid.UUID, customerrors.ErrorInterface)

	InsertRagConversation(tx *sqlx.Tx, userID uuid.UUID, bafoegType string) (uuid.UUID, customerrors.ErrorInterface)
	InsertRagConversationMessage(tx *sqlx.Tx, conversationID uuid.UUID, messageContent string, isUserInputMessage bool) (uuid.UUID, customerrors.ErrorInterface)
	GetRagConversationMessagesByConversationID(tx *sqlx.Tx, page uint, conversationID uuid.UUID) ([]models.ConversationMessage, customerrors.ErrorInterface)

	GetRagConversationByID(tx *sqlx.Tx, conversationID uuid.UUID) (*models.Conversation, customerrors.ErrorInterface)

	DeleteRagVectorSchueler(tx *sqlx.Tx, documentID uuid.UUID) customerrors.ErrorInterface
	DeleteRagVectorStudierenden(tx *sqlx.Tx, documentID uuid.UUID) customerrors.ErrorInterface

	DeleteRagFileSchueler(tx *sqlx.Tx, documentID uuid.UUID) customerrors.ErrorInterface
	DeleteRagFileStudierenden(tx *sqlx.Tx, documentID uuid.UUID) customerrors.ErrorInterface
}

type RAGStorageS3 interface {
	UploadDocumentSchüler(fileID string, file []byte) customerrors.ErrorInterface
	UploadDocumentStudierenden(fileID string, file []byte) customerrors.ErrorInterface
	DownloadDocumentSchülerByID(id string) ([]byte, customerrors.ErrorInterface)
	DownloadDocumentStudierendenByID(id string) ([]byte, customerrors.ErrorInterface)
}
