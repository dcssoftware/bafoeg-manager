package http

import (
	"time"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// CreateApplicationFromEakteRequest represents the request body for creating an application from eAkte
type CreateApplicationFromEakteRequest struct {
	ApplicantID        uuid.UUID   `json:"applicantID"`
	EakteID            uuid.UUID   `json:"eakteID"`
	SchoolDegreeID     uuid.UUID   `json:"schoolDegreeID"`
	EducationStartDate time.Time   `json:"educationStartDate"`
	EducationEndDate   time.Time   `json:"educationEndDate"`
	ClassLevel         string      `json:"classLevel"`
	LabelIDs           []uuid.UUID `json:"labelIDs"`
}

func (h *ApplicationsHandler) CreateApplicationFromEakte(c fiber.Ctx) error {
	// Parse request body
	var req CreateApplicationFromEakteRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	applicationID, revisionID, mappingID, applicationErr := h.service.InsertApplicationFromEakte(nil, serviceModel.ApplicationFromEakteInsertModel{
		ClassLevel:         req.ClassLevel,
		EakteAkteID:        req.EakteID,
		LabelIDs:           req.LabelIDs,
		ApplicantID:        req.ApplicantID,
		SchoolWithDegreeID: req.SchoolDegreeID,
		EducationStart:     req.EducationStartDate,
		EducationEnd:       req.EducationEndDate,
	})

	if applicationErr != nil {
		status, message := applicationErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.JSON(map[string]any{
		"success":          true,
		"application_id":   applicationID,
		"base_revision_id": revisionID,
		"mapping_id":       mappingID,
	})
}
