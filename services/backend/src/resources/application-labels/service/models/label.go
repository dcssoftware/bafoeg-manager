package models

import "github.com/google/uuid"

type ApplicationLabel struct {
	ID    uuid.UUID
	Name  string
	Style *ApplicationLabelStyle
}
type ApplicationLabelStyle struct {
	ID                   uuid.UUID
	Name                 string
	ColorDark            string
	BackgroundColorDark  string
	ColorLight           string
	BackgroundColorLight string
}
