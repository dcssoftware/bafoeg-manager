package models

import (
	"time"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/google/uuid"
)

type ApplicantsBySchoolModel struct {
	SchholID                    uuid.UUID `db:"school_id"`
	ApplicationClassLevel       string    `db:"application_class_level"`
	ApplicationStatusIdentifier string    `db:"application_status_identifier"`
	ApplicationValidityStarts   time.Time `db:"application_validity_starts"`
	ApplicationValidityEnds     time.Time `db:"application_validity_ends"`

	ApplicantID        uuid.UUID `db:"applicants_id"`
	ApplicantFirstname string    `db:"applicants_firstname"`
	ApplicantLastname  string    `db:"applicants_lastname"`
	ApplicantZipCode   string    `db:"applicants_address_zip_code"`
	ApplicantCity      string    `db:"applicants_address_city"`
	ApplicantCountry   string    `db:"applicants_address_country"`

	SchoolDegreeID   uuid.UUID `db:"school_degree_id"`
	SchoolDegreeName string    `db:"school_degree_name"`
}

func (u *ApplicantsBySchoolModel) ToApplicantsBySchoolServiceModel() *serviceModels.ApplicantBySchoolModel {
	if u == nil {
		return nil
	}

	return &serviceModels.ApplicantBySchoolModel{
		ID:               u.ApplicantID,
		Firstname:        u.ApplicantFirstname,
		Lastname:         u.ApplicantLastname,
		ClassLevel:       u.ApplicationClassLevel,
		StatusIdentifier: u.ApplicationStatusIdentifier,

		Degree:  u.toApplicantDegreeServiceModel(),
		Address: u.toApplicantAddressServiceModel(),
	}
}

func ToApplicantsBySchoolServiceModels(user []ApplicantsBySchoolModel) []serviceModels.ApplicantBySchoolModel {
	var serviceModels []serviceModels.ApplicantBySchoolModel

	for _, u := range user {
		serviceModels = append(serviceModels, *u.ToApplicantsBySchoolServiceModel())
	}

	return serviceModels
}

func (u *ApplicantsBySchoolModel) toApplicantAddressServiceModel() *serviceModels.ApplicantBySchoolAddressModel {
	if u == nil {
		return nil
	}

	return &serviceModels.ApplicantBySchoolAddressModel{
		ZipCode: u.ApplicantZipCode,
		City:    u.ApplicantCity,
		Country: u.ApplicantCountry,
	}
}

func (u *ApplicantsBySchoolModel) toApplicantDegreeServiceModel() *serviceModels.ApplicantBySchoolDegreeModel {
	if u == nil {
		return nil
	}

	return &serviceModels.ApplicantBySchoolDegreeModel{
		ID:   u.SchoolDegreeID,
		Name: u.SchoolDegreeName,
	}
}
