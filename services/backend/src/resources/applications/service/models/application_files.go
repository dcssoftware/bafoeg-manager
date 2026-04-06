package models

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationFile struct {
	Source        string
	ID            uuid.UUID
	ApplicationID uuid.UUID
	Eakte         *ApplicationFileEakteModel
	File          *ApplicationFileFileModel
}

type ApplicationFileEakteModel struct {
	AkteID          uuid.UUID
	Vertraulichkeit string
	Created         time.Time
}

type ApplicationFileFileModel struct {
	FileID  uuid.UUID
	Name    string
	Size    float64
	Type    string
	Created time.Time
}
