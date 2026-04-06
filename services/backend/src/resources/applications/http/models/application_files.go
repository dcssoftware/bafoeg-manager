package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"

	"github.com/google/uuid"
)

type ApplicationFileModels struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Files []ApplicationFile `json:"files"`
}

type ApplicationFile struct {
	Source        string                     `json:"source"`
	ID            uuid.UUID                  `json:"id"`
	ApplicationID uuid.UUID                  `json:"applicationID"`
	Eakte         *ApplicationFileEakteModel `json:"eakte"`
	File          *ApplicationFileFileModel  `json:"file"`
}

type ApplicationFileEakteModel struct {
	AkteID          uuid.UUID `json:"akteID"`
	Vertraulichkeit string    `json:"vertraulichkeit"`
	Created         time.Time `json:"created"`
}

type ApplicationFileFileModel struct {
	FileID  uuid.UUID `json:"fileID"`
	Name    string    `json:"name"`
	Size    float64   `json:"size"`
	Type    string    `json:"type"`
	Created time.Time `json:"created"`
}

func ToFilesHttpModels(files []serviceModel.ApplicationFile, maxCount uint) *ApplicationFileModels {
	return &ApplicationFileModels{
		Count:    uint(len(files)),
		MaxCount: maxCount,

		Files: toFilesHttpModels(files),
	}
}

func toFilesHttpModels(files []serviceModel.ApplicationFile) []ApplicationFile {

	var httpFiles []ApplicationFile

	for _, applicationFile := range files {

		var eakte *ApplicationFileEakteModel
		if applicationFile.Eakte != nil {
			eakte = &ApplicationFileEakteModel{
				AkteID:          applicationFile.Eakte.AkteID,
				Vertraulichkeit: applicationFile.Eakte.Vertraulichkeit,
				Created:         applicationFile.Eakte.Created,
			}
		}

		var file *ApplicationFileFileModel
		if applicationFile.File != nil {
			file = &ApplicationFileFileModel{
				FileID:  applicationFile.File.FileID,
				Name:    applicationFile.File.Name,
				Size:    applicationFile.File.Size,
				Type:    applicationFile.File.Type,
				Created: applicationFile.File.Created,
			}
		}

		httpFiles = append(httpFiles, ApplicationFile{
			Source:        applicationFile.Source,
			ID:            applicationFile.ID,
			ApplicationID: applicationFile.ApplicationID,
			Eakte:         eakte,
			File:          file,
		})
	}

	return httpFiles
}
