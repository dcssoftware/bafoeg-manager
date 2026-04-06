package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/google/uuid"
)

type EaktenFileModel struct {
	ID               string    `db:"id"`
	Source           uuid.UUID `db:"source"`
	VorgangID        uuid.UUID `db:"vorgang_id"`
	SourceXdomeaFile bool      `db:"source_xdomea_file"`
	FileID           uuid.UUID `db:"file_id"`
	Created          time.Time `db:"created"`
}

func (model *EaktenFileModel) ToServiceModel() *serviceModels.EaktenFileModel {
	return &serviceModels.EaktenFileModel{
		ID:               model.ID,
		Source:           model.Source,
		VorgangID:        model.VorgangID,
		SourceXdomeaFile: model.SourceXdomeaFile,
		FileID:           model.FileID,
		Created:          model.Created,
	}
}
