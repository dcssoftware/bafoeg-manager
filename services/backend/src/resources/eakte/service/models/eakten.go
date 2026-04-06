package models

import (
	"time"

	stateModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/states/models"
	"github.com/google/uuid"
)

type EakteModel struct {
	ID              uuid.UUID
	Aktenzeichen    string
	Typ             EakteModelType
	Vertraulichkeit stateModels.VertraulichkeitEnum
	Created         *time.Time
}

type EakteModelType struct {
	ID         uuid.UUID
	Name       string
	Identifier string
}
