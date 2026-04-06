package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/google/uuid"
)

type SchoolDegreeModel struct {
	ID                                 uuid.UUID `db:"id" json:"id"`
	Name                               string    `db:"name" json:"name"`
	SchoolID                           uuid.UUID `db:"school_id" json:"school_id"`
	FosBerufsabschlussRequired         bool      `db:"fos_berufsabschluss_required" json:"fos_berufsabschluss_required"`
	BosBerufsqualifizierenderAbschluss bool      `db:"bos_berufsqualifizierender_abschluss" json:"bos_berufsqualifizierender_abschluss"`
	FachschuleBerufsschuleRequired     bool      `db:"fachschule_berufsabschluss_required" json:"fachschule_berufsabschluss_required"`
}

func (model *SchoolDegreeModel) ToServiceModel() serviceModel.SchoolDegreeModel {

	return serviceModel.SchoolDegreeModel{
		ID:                                 model.ID,
		Name:                               model.Name,
		FosBerufsabschlussRequired:         model.FosBerufsabschlussRequired,
		BosBerufsqualifizierenderAbschluss: model.BosBerufsqualifizierenderAbschluss,
		FachschuleBerufsschuleRequired:     model.FachschuleBerufsschuleRequired,
	}
}

func ToServiceModels(models []SchoolDegreeModel) []serviceModel.SchoolDegreeModel {

	var degrees []serviceModel.SchoolDegreeModel
	for _, degree := range models {
		degrees = append(degrees, degree.ToServiceModel())
	}

	return degrees
}

func (model *SchoolDegreeModels) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &model)
	case string:
		return json.Unmarshal([]byte(v), &model)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (model *SchoolDegreeModels) Value() (driver.Value, error) {
	return json.Marshal(model)
}
