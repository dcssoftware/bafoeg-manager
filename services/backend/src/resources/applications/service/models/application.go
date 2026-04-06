package models

import (
	"time"

	applicantServiceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/google/uuid"
)

type ApplicationModel struct {
	ID         uuid.UUID
	ClassLevel string

	Labels           *[]ApplicationLabel
	Applicant        *ApplicationApplicant
	SchoolWithDegree *ApplicationSchoolWithDegree
	AssignedUser     *ApplicationAssignedUser
	Status           *ApplicationStatus
	ProcessingTime   *ApplicationProcessingTime

	Created time.Time
	Updated time.Time
}

type ApplicationFromEakteInsertModel struct {
	EakteAkteID uuid.UUID

	ClassLevel string

	LabelIDs           []uuid.UUID
	ApplicantID        uuid.UUID
	SchoolWithDegreeID uuid.UUID

	EducationStart time.Time
	EducationEnd   time.Time
}

type ApplicationInsertModel struct {
	ClassLevel string

	LabelIDs           []string
	ApplicantID        string
	SchoolWithDegreeID string
	AssignedUserID     string
	StatusID           string

	EducationStart time.Time
	EducationEnd   time.Time
}

type ApplicationLabel struct {
	ID    uuid.UUID
	Name  string
	Color *ApplicationLabelColor
}
type ApplicationLabelColor struct {
	ID   uuid.UUID
	Name string

	ColorDark       string
	BgColorDark     string
	BorderColorDark string

	ColorLight       string
	BgColorLight     string
	BorderColorLight string
}

type ApplicationApplicant struct {
	ID               uuid.UUID
	Firstname        string
	Lastname         string
	Address          *ApplicationApplicantAddress
	TrainingsAddress *applicantServiceModel.ApplicantTrainingsAddressModel
	ContactData      *ApplicationApplicantContactData
}

type ApplicationApplicantAddress struct {
	Street      string
	HouseNumber string
	ZipCode     string
	City        string
	Country     string
}

type ApplicationApplicantContactData struct {
	Email string
	Phone string
}

type ApplicationSchoolWithDegree struct {
	ID      uuid.UUID
	Name    string
	Address *ApplicationSchoolWithDegreeAddress
	Type    *ApplicationSchoolWithDegreeType
	Degree  *ApplicationSchoolWithDegreeDegree
}

type ApplicationSchoolWithDegreeAddress struct {
	Street      string
	HouseNumber string
	ZipCode     string
	City        string
	Country     string
}

type ApplicationSchoolWithDegreeType struct {
	ID         uuid.UUID
	Name       string
	Identifier string
}

type ApplicationSchoolWithDegreeDegree struct {
	ID                                 uuid.UUID
	Name                               string
	FosBerufsabschlussRequired         bool
	BosBerufsqualifizierenderAbschluss bool
	FachschuleBerufsabschlussRequired  bool
}

type ApplicationAssignedUser struct {
	ID          uuid.UUID
	Username    string
	DisplayName string
}

type ApplicationStatus struct {
	ID         uuid.UUID
	Name       string
	Identifier string
}

type ApplicationProcessingTime struct {
	MaxValidity            uint
	RemainingTimeInDays    int
	RemainingTimeInPercent float64
	IsStillLegal           bool
}

func (s ApplicationStatus) IsUpdatable() bool {
	switch s.Identifier {
	case "IN_PROGRESS", "RESPONSE_AWAITED":
		return true
	}
	return false
}
