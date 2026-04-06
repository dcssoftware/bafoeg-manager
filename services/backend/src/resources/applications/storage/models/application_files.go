package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

type ApplicationFile struct {
	Source        string                     `db:"source"`
	ID            uuid.UUID                  `db:"id"`
	ApplicationID uuid.UUID                  `db:"application_id"`
	Eakte         *ApplicationFileEakteModel `db:"eakte,omitempty"`
	File          *ApplicationFileFileModel  `db:"file"`
}

type ApplicationFileEakteModel struct {
	AkteID          uuid.UUID `json:"akte_id"`
	Vertraulichkeit string    `json:"vertraulichkeit"`
	Created         time.Time `json:"created"`
}

type ApplicationFileFileModel struct {
	FileID  uuid.UUID `json:"file_id"`
	Name    string    `json:"name"`
	Size    float64   `json:"size"`
	Type    string    `json:"type"`
	Created time.Time `json:"created"`
}

func (model *ApplicationFile) ToServiceFileModel() *serviceModel.ApplicationFile {
	var eakte *serviceModel.ApplicationFileEakteModel
	if model.Eakte != nil {
		eakte = model.Eakte.ToServiceModel()
	}

	var file *serviceModel.ApplicationFileFileModel
	if model.File != nil {
		file = model.File.ToServiceModel()
	}

	return &serviceModel.ApplicationFile{
		Source:        model.Source,
		ID:            model.ID,
		ApplicationID: model.ApplicationID,
		Eakte:         eakte,
		File:          file,
	}
}

func (model *ApplicationFileEakteModel) ToServiceModel() *serviceModel.ApplicationFileEakteModel {
	return &serviceModel.ApplicationFileEakteModel{
		AkteID:          model.AkteID,
		Vertraulichkeit: model.Vertraulichkeit,
		Created:         model.Created,
	}
}
func (model *ApplicationFileFileModel) ToServiceModel() *serviceModel.ApplicationFileFileModel {
	return &serviceModel.ApplicationFileFileModel{
		FileID:  model.FileID,
		Name:    model.Name,
		Size:    model.Size,
		Type:    model.Type,
		Created: model.Created,
	}
}
