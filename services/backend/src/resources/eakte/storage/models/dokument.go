package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/google/uuid"
)

type DokumentModel struct {
	ID               uuid.UUID            `db:"id"`
	SourceXdomeaFile bool                 `db:"source_xdomea_file"`
	Source           DokumentModelSource  `db:"source"`
	File             DokumentModelFile    `db:"files"`
	Vorgang          DokumentModelVorgang `db:"vorgang"`
	Eakte            DokumentModelEakte   `db:"akte"`
	Created          *time.Time           `db:"created"`
}

type DokumentModelSource struct {
	ID         uuid.UUID `json:"id"`
	Identidier string    `json:"identifier"`
	Name       string    `json:"name"`
}

type DokumentModelFile struct {
	ID      uuid.UUID  `json:"id"`
	Name    string     `json:"name"`
	Type    string     `json:"file_type"`
	Size    float64    `json:"file_size"`
	Created *time.Time `json:"created"`
}

type DokumentModelVorgang struct {
	ID              uuid.UUID `json:"id"`
	Vorgangszeichen string    `json:"vorgangszeichen"`
}

type DokumentModelEakte struct {
	ID           uuid.UUID `json:"id"`
	Aktenzeichen string    `json:"aktenzeichen"`
}

func (m DokumentModel) ToServiceModel() *serviceModels.DokumentModel {
	return &serviceModels.DokumentModel{
		ID:               m.ID,
		SourceXdomeaFile: m.SourceXdomeaFile,
		Source:           *m.Source.ToServiceModel(),
		File:             *m.File.ToServiceModel(),
		Vorgang:          *m.Vorgang.ToServiceModel(),
		Eakte:            *m.Eakte.ToServiceModel(),
		Created:          *m.Created,
	}
}

func (m DokumentModelSource) ToServiceModel() *serviceModels.DokumentModelSource {
	return &serviceModels.DokumentModelSource{
		ID:         m.ID,
		Identidier: m.Identidier,
		Name:       m.Name,
	}
}

func (m DokumentModelFile) ToServiceModel() *serviceModels.DokumentModelFile {
	return &serviceModels.DokumentModelFile{
		ID:      m.ID,
		Name:    m.Name,
		Size:    m.Size,
		Type:    m.Type,
		Created: *m.Created,
	}
}

func (m DokumentModelVorgang) ToServiceModel() *serviceModels.DokumentModelVorgang {
	return &serviceModels.DokumentModelVorgang{
		ID:              m.ID,
		Vorgangszeichen: m.Vorgangszeichen,
	}
}

func (m DokumentModelEakte) ToServiceModel() *serviceModels.DokumentModelEakte {
	return &serviceModels.DokumentModelEakte{
		ID:           m.ID,
		Aktenzeichen: m.Aktenzeichen,
	}
}

func ToDokumentenServiceModel(models []DokumentModel) []serviceModels.DokumentModel {
	var svcModels []serviceModels.DokumentModel

	for _, model := range models {
		svcModels = append(svcModels, *model.ToServiceModel())
	}

	return svcModels
}
