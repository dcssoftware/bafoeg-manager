package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor/service/models"
	"github.com/google/uuid"
)

type DocumentModel struct {
	ID             uuid.UUID                `db:"id"`
	Filename       string                   `db:"file_name"`
	FileType       string                   `db:"file_type"`
	FileSize       float64                  `db:"file_size"`
	FileHash       string                   `db:"file_hash"`
	Status         DocumentStatusModel      `db:"status"`
	ProcessedTime  *time.Time               `db:"processed_timestamp"`
	ProcessedError string                   `db:"processed_error"`
	Created        time.Time                `db:"created"`
	CreatedFrom    DocumentCreatedFromModel `db:"created_from"`
}

type DocumentStatusModel struct {
	ID         uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	Name       string    `json:"name"`
}

type DocumentCreatedFromModel struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"display_name"`
	Username    string    `json:"username"`
}

func (file *DocumentModel) ToServiceFileModel() *serviceModel.DocumentModel {
	return &serviceModel.DocumentModel{
		ID: file.ID,

		Filename: file.Filename,
		FileType: file.FileType,
		FileSize: file.FileSize,
		Status: serviceModel.DocumentStatusModel{
			ID:         file.Status.ID,
			Identifier: file.Status.Identifier,
			Name:       file.Status.Name,
		},
		Created: file.Created,
		CreatedFrom: serviceModel.DocumentCreatedFromModel{
			ID:          file.CreatedFrom.ID,
			DisplayName: file.CreatedFrom.DisplayName,
			Username:    file.CreatedFrom.Username,
		},
	}
}
