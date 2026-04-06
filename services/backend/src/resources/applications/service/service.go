package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	applicantsService "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/states"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage"
	filesService "github.com/dcssoftware/bafoeg-manager/src/resources/files/service"
	schoolService "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

type ApplicationsService struct {
	storage           ApplicationsStorageInterface
	applicantsService *applicantsService.ApplicantService
	schoolService     *schoolService.SchoolService
	filesService      *filesService.FileService
}

func NewApplicationsService(storage *storage.ApplicationsStorage, applicantsService *applicantsService.ApplicantService, schoolService *schoolService.SchoolService, filesService *filesService.FileService) *ApplicationsService {
	return &ApplicationsService{
		storage,
		applicantsService,
		schoolService,
		filesService,
	}
}

type ApplicationsStorageInterface interface {
	StartTx() (*sqlx.Tx, error)

	GetApplicationByID(tx *sqlx.Tx, applicationID string) (*models.ApplicationModel, customerrors.ErrorInterface)
	GetApplicationFilesByApplicationID(tx *sqlx.Tx, page uint, applicationID string) ([]models.ApplicationFile, customerrors.ErrorInterface)
	GetApplicationFilesByApplicationIDCount(tx *sqlx.Tx, applicationID string) (uint, customerrors.ErrorInterface)
	GetApplicationFileByFileID(tx *sqlx.Tx, fileID string) (*models.ApplicationFile, customerrors.ErrorInterface)

	GetFileIDByApplicationFileID(tx *sqlx.Tx, applicationFileID string) (uuid.UUID, customerrors.ErrorInterface)

	GetApplicationsCount(tx *sqlx.Tx, userID string, applicantID string) (uint, customerrors.ErrorInterface)
	GetApplications(tx *sqlx.Tx, page uint, userID string, filter string, filterOnlyInProgress bool) ([]models.ApplicationShortModel, customerrors.ErrorInterface)

	GetApplicationsByApplicantID(tx *sqlx.Tx, page uint, applicantID string) ([]models.ApplicationShortModel, customerrors.ErrorInterface)
	GetApplicationsByApplicantIDCount(tx *sqlx.Tx, applicantID string) (uint, customerrors.ErrorInterface)

	GetApplicationRevisionsByApplicationID(tx *sqlx.Tx, page uint, applicantID string) ([]models.ApplicationRevisionShortModel, customerrors.ErrorInterface)
	GetApplicationRevisionsByApplicationIDCount(tx *sqlx.Tx, applicantID string) (uint, customerrors.ErrorInterface)

	GetApplicationsMetricsTotal(tx *sqlx.Tx) (uint, customerrors.ErrorInterface)
	GetApplicationsMetricsInProgress(tx *sqlx.Tx) (uint, customerrors.ErrorInterface)
	GetApplicationsMetricsNewToday(tx *sqlx.Tx) (uint, customerrors.ErrorInterface)
	GetApplicationsMetricsUserAssigned(tx *sqlx.Tx, userID string) (uint, customerrors.ErrorInterface)

	InsertApplication(tx *sqlx.Tx, model models.ApplicationInsertModel) (uuid.UUID, customerrors.ErrorInterface)
	InsertApplicationFromEakte(tx *sqlx.Tx, model models.ApplicationFromEakteInsertModel) (uuid.UUID, customerrors.ErrorInterface)

	InsertApplicationEakteMapping(tx *sqlx.Tx, applicationID, eakteAkteID uuid.UUID) (uuid.UUID, customerrors.ErrorInterface)

	GetApplicationStatus(tx *sqlx.Tx, page uint) ([]models.ApplicationStatus, customerrors.ErrorInterface)
	GetApplicationStatusByID(tx *sqlx.Tx, identifier string) (*models.ApplicationStatus, customerrors.ErrorInterface)
	GetApplicationStatusByIdentifier(tx *sqlx.Tx, identifier string) (*models.ApplicationStatus, customerrors.ErrorInterface)

	InsertApplicationRevision(tx *sqlx.Tx, applicationID string, header, description string, baseRevision *models.ApplicationRevisionDataModel, newRevisionData *models.ApplicationRevisionDataModel) (uuid.UUID, customerrors.ErrorInterface)
	UploadApplicationFile(tx *sqlx.Tx, applicationID string, fileID uuid.UUID) (uuid.UUID, customerrors.ErrorInterface)
	ExistsUploadedApplicationFileHash(tx *sqlx.Tx, applicationID string, filetype string, fileSize uint, fileHash string) (bool, customerrors.ErrorInterface)

	GetApplicationRevisionLatestRevisionByApplicationID(tx *sqlx.Tx, applicationID string) (*models.ApplicationRevisionModel, customerrors.ErrorInterface)

	DeleteApplicationFileByFileID(tx *sqlx.Tx, fileID string) (uuid.UUID, customerrors.ErrorInterface)

	UpdateApplicationStatus(tx *sqlx.Tx, applicationID string, newStatus states.ApplicationState) customerrors.ErrorInterface
	UpdateApplicationAssignedUser(tx *sqlx.Tx, applicationID string, newAssignedUser string) customerrors.ErrorInterface
	UpdateApplicationAssignedSchoolDegree(tx *sqlx.Tx, applicationID uuid.UUID, newSchoolID uuid.UUID) customerrors.ErrorInterface
	UpdateApplicationAssignedApplicant(tx *sqlx.Tx, applicationID uuid.UUID, newApplicantID uuid.UUID) customerrors.ErrorInterface
}
