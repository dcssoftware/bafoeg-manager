package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

type ApplicationRevisionShortModel struct {
	ID            uuid.UUID `db:"id"`
	ApplicationID uuid.UUID `db:"application_id"`

	Header      string `db:"message_header"`
	Description string `db:"message_description"`

	Created time.Time `db:"created"`
}

func (model *ApplicationRevisionShortModel) ToServiceShortModel() *serviceModel.ApplicationRevisionShortModel {
	return &serviceModel.ApplicationRevisionShortModel{
		ID:            model.ID,
		Header:        model.Header,
		Description:   model.Description,
		ApplicationID: model.ApplicationID,
		Created:       model.Created,
	}
}
