package models

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationShortModel struct {
	ID           uuid.UUID
	ClassLevel   string
	Applicant    *ApplicationShortApplicant
	AssignedUser *ApplicationAssignedUser
	School       *ApplicationSchoolWithDegree
	Status       *ApplicationStatus

	ProcessingTime *ApplicationProcessingTime
	Created        time.Time
	Updated        time.Time
}

type ApplicationShortApplicant struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
}
