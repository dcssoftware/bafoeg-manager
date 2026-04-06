package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/google/uuid"
)

type SchoolDegreeResponseModel struct {
	Count         uint                `json:"count"`
	MaxCount      uint                `json:"maxCount"`
	SchoolDegrees []SchoolDegreeModel `json:"degrees"`
}

type SchoolDegreeModel struct {
	ID                                 uuid.UUID `json:"id"`
	Name                               string    `json:"name"`
	SchoolID                           uuid.UUID `json:"schoolID"`
	FosBerufsabschlussRequired         bool      `json:"fosBerufsabschlussRequired"`
	BosBerufsqualifizierenderAbschluss bool      `json:"bosBerufsqualifizierenderAbschluss"`
	FachschuleBerufsschuleRequired     bool      `json:"fachschuleBerufsschuleRequired"`
}

func ToHttpSchoolDegreeModel(degree serviceModel.SchoolDegreeModel) SchoolDegreeModel {
	return SchoolDegreeModel{
		ID:                                 degree.ID,
		Name:                               degree.Name,
		SchoolID:                           degree.SchoolID,
		FosBerufsabschlussRequired:         degree.FosBerufsabschlussRequired,
		BosBerufsqualifizierenderAbschluss: degree.BosBerufsqualifizierenderAbschluss,
		FachschuleBerufsschuleRequired:     degree.FachschuleBerufsschuleRequired,
	}
}

func ToHttpSchoolDegreeResponseModel(models []serviceModel.SchoolDegreeModel, maxCount uint) SchoolDegreeResponseModel {

	var httpModels []SchoolDegreeModel
	for _, m := range models {
		httpModels = append(httpModels, ToHttpSchoolDegreeModel(m))
	}

	return SchoolDegreeResponseModel{
		Count:         uint(len(httpModels)),
		MaxCount:      uint(maxCount),
		SchoolDegrees: httpModels,
	}
}
