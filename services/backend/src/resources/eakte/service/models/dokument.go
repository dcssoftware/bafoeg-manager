package models

import (
	"time"

	"github.com/google/uuid"
)

type DokumentModel struct {
	ID               uuid.UUID
	SourceXdomeaFile bool
	Source           DokumentModelSource
	File             DokumentModelFile
	Vorgang          DokumentModelVorgang
	Eakte            DokumentModelEakte
	Created          time.Time
}

type DokumentModelSource struct {
	ID         uuid.UUID
	Identidier string
	Name       string
}

type DokumentModelFile struct {
	ID      uuid.UUID
	Name    string
	Size    float64
	Type    string
	Created time.Time
}

type DokumentModelVorgang struct {
	ID              uuid.UUID
	Vorgangszeichen string
}

type DokumentModelEakte struct {
	ID           uuid.UUID
	Aktenzeichen string
}
