package models

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/google/uuid"
)

type ApplicationRevisionModel struct {
	ID            uuid.UUID `db:"id"`
	ApplicationID uuid.UUID `db:"application_id"`

	Header      string `db:"message_header"`
	Description string `db:"message_description"`

	// revision data
	TrainingsAddressID uuid.UUID `db:"trainings_address_id"`

	Created time.Time `db:"created"`
}

func (model *ApplicationRevisionModel) ToServiceShortModel() *serviceModel.ApplicationRevisionShortModel {
	return &serviceModel.ApplicationRevisionShortModel{
		ID:            model.ID,
		Header:        model.Header,
		Description:   model.Description,
		ApplicationID: model.ApplicationID,
		Created:       model.Created,
	}
}
