package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

type SchoolService struct {
	storage SchoolStorage
}

func NewSchoolService(storage SchoolStorage) *SchoolService {
	return &SchoolService{
		storage,
	}
}

type SchoolStorage interface {
	GetSchools(tx *sqlx.Tx, page uint, filter string) ([]model.SchoolShortModel, customerrors.ErrorInterface)
	GetSchoolsCount(tx *sqlx.Tx, filter string) (uint, customerrors.ErrorInterface)

	GetSchoolByID(tx *sqlx.Tx, schoolID string) (*model.SchoolModel, customerrors.ErrorInterface)

	GetSchoolDegreeByID(tx *sqlx.Tx, degreeID uuid.UUID) (*model.SchoolDegreeModel, customerrors.ErrorInterface)
	GetSchoolDegree(tx *sqlx.Tx, page uint) ([]model.SchoolDegreeModel, customerrors.ErrorInterface)

	GetSchoolDegreeBySchoolID(tx *sqlx.Tx, page uint, schoolID string) ([]model.SchoolDegreeModel, customerrors.ErrorInterface)
	GetSchoolDegreeBySchoolIDCount(tx *sqlx.Tx, schoolID string) (uint, customerrors.ErrorInterface)
}
