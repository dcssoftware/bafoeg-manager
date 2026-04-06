package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	stateModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/states/models"
	fileService "github.com/dcssoftware/bafoeg-manager/src/resources/files/service"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

type EakteService struct {
	storage     EakteStorage
	fileService *fileService.FileService
}

func NewEakteService(storage EakteStorage, fileService *fileService.FileService) *EakteService {
	return &EakteService{
		storage:     storage,
		fileService: fileService,
	}
}

type EakteStorage interface {
	StartTx() (*sqlx.Tx, error)

	GetEakten(tx *sqlx.Tx, page uint) ([]models.EakteModel, customerrors.ErrorInterface)
	GetEaktenCount(tx *sqlx.Tx) (uint, customerrors.ErrorInterface)

	GetVorgängeByEaktenID(tx *sqlx.Tx, id string) ([]models.VorgangModel, customerrors.ErrorInterface)

	GetFilesByVorgangsID(tx *sqlx.Tx, vorgangID string) ([]models.DokumentModel, customerrors.ErrorInterface)
	GetFilesByVorgangsIDCount(tx *sqlx.Tx, vorgangID string) (uint, customerrors.ErrorInterface)
	GetFilesByAkteID(tx *sqlx.Tx, akteID string) ([]models.DokumentModel, customerrors.ErrorInterface)
	GetFilesByAkteIDCount(tx *sqlx.Tx, akteID string) (uint, customerrors.ErrorInterface)
	GetEakteApplicationMapping(tx *sqlx.Tx, eakteAkteID string) (*models.EaktenApplicationMappingModel, customerrors.ErrorInterface)
	GetFileByFileID(tx *sqlx.Tx, fileID uuid.UUID) (*models.EaktenFileModel, customerrors.ErrorInterface)
	GetEakteByID(tx *sqlx.Tx, id uuid.UUID) (*models.EakteModel, customerrors.ErrorInterface)

	ExistsUploadedEakteFileHash(tx *sqlx.Tx, filetype string, fileSize uint, fileHash string) (bool, customerrors.ErrorInterface)

	// InsertEakteFile(tx *sqlx.Tx, vorgangID, source, fileID uuid.UUID) (uuid.UUID, customerrors.ErrorInterface)
	InsertEakteAkte(tx *sqlx.Tx, aktenzeichen string, typIdentifier string, vertraulichkeit stateModels.VertraulichkeitEnum) (uuid.UUID, customerrors.ErrorInterface)
	InsertEakteVorgang(tx *sqlx.Tx, akteID uuid.UUID, vorgangszeichen string) (uuid.UUID, customerrors.ErrorInterface)
	InsertEakteDocument(tx *sqlx.Tx, vorgangID, fileID uuid.UUID, sourceIdentidier string, isXdomeaZipFile bool) (uuid.UUID, customerrors.ErrorInterface)
}
