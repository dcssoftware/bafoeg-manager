package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/google/uuid"
)

type DokumenteResponseModels struct {
	Count    uint `json:"count"`
	MaxCount uint `json:"maxCount"`

	Documents []DokumentModel `json:"documents"`
}

type DokumentModel struct {
	ID               uuid.UUID           `json:"id"`
	SourceXdomeaFile bool                `json:"source_xdomea_file"`
	Source           DokumentModelSource `json:"source"`
	File             DokumentModelFile   `json:"files"`
	Created          time.Time           `json:"created"`
}

type DokumentModelSource struct {
	ID         uuid.UUID `json:"id"`
	Identidier string    `json:"identifier"`
	Name       string    `json:"name"`
}

type DokumentModelFile struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	Size    float64   `json:"size"`
	Created time.Time `json:"created"`
}

type DokumentModelVorgang struct {
	ID              uuid.UUID `json:"id"`
	Vorgangszeichen string    `json:"vorgangszeichen"`
}

type DokumentModelEakte struct {
	ID              uuid.UUID `json:"id"`
	Aktenzeichen    string    `json:"aktenzeichen"`
	Vertraulichkeit string    `json:"vertraulichkeit"`
}

func ToDokumenteResponseModels(models []serviceModels.DokumentModel, maxCount uint) *DokumenteResponseModels {
	return &DokumenteResponseModels{
		MaxCount:  maxCount,
		Count:     uint(len(models)),
		Documents: ToDokumentModels(models),
	}
}

func ToDokumentModels(models []serviceModels.DokumentModel) []DokumentModel {
	var httpModels []DokumentModel

	for _, model := range models {
		httpModels = append(httpModels, *ToDokumentModel(model))
	}

	return httpModels
}

func ToDokumentModel(model serviceModels.DokumentModel) *DokumentModel {
	return &DokumentModel{
		ID:               model.ID,
		SourceXdomeaFile: model.SourceXdomeaFile,
		Source:           *ToDokumentModelSource(model.Source),
		File:             *ToDokumentModelFile(model.File),
		Created:          model.Created,
	}
}

func ToDokumentModelSource(model serviceModels.DokumentModelSource) *DokumentModelSource {
	return &DokumentModelSource{
		ID:         model.ID,
		Identidier: model.Identidier,
		Name:       model.Name,
	}
}

func ToDokumentModelFile(model serviceModels.DokumentModelFile) *DokumentModelFile {
	return &DokumentModelFile{
		ID:      model.ID,
		Name:    model.Name,
		Type:    model.Type,
		Size:    model.Size,
		Created: model.Created,
	}
}
