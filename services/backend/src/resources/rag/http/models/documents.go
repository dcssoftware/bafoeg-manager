package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/models"
	"github.com/google/uuid"
)

type DocumentModelResponse struct {
	Count     uint             `json:"count"`
	MaxCount  uint             `json:"maxCount"`
	Documents []*DocumentModel `json:"documents"`
}

type DocumentModel struct {
	ID          uuid.UUID                `json:"id"`
	Filename    string                   `json:"fileName"`
	FileType    string                   `json:"fileType"`
	FileSize    float64                  `json:"fileSize"`
	Status      DocumentStatusModel      `json:"status"`
	Created     time.Time                `json:"created"`
	CreatedFrom DocumentCreatedFromModel `json:"createdFrom"`
}

type DocumentStatusModel struct {
	ID         uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	Name       string    `json:"name"`
}

type DocumentCreatedFromModel struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"displayName"`
	Username    string    `json:"username"`
}

func ToDocumentsHttpModels(documents []serviceModel.DocumentModel, maxCount uint) *DocumentModelResponse {
	return &DocumentModelResponse{
		Count:    uint(len(documents)),
		MaxCount: maxCount,

		Documents: toDocumentsHttpModels(documents),
	}
}

func toDocumentsHttpModels(documents []serviceModel.DocumentModel) []*DocumentModel {

	var responseDocuments []*DocumentModel

	for _, document := range documents {
		responseDocuments = append(responseDocuments, &DocumentModel{
			ID:       document.ID,
			Filename: document.Filename,
			FileType: document.FileType,
			FileSize: document.FileSize,
			Status: DocumentStatusModel{
				ID:         document.Status.ID,
				Identifier: document.Status.Identifier,
				Name:       document.Status.Name,
			},
			Created: document.Created,
			CreatedFrom: DocumentCreatedFromModel{
				ID:          document.CreatedFrom.ID,
				DisplayName: document.CreatedFrom.DisplayName,
				Username:    document.CreatedFrom.Username,
			},
		})
	}
	return responseDocuments
}
