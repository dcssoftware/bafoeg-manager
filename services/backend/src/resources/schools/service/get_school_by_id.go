package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
)

func (s *SchoolService) GetSchoolByID(schoolID string) (*model.SchoolModel, customerrors.ErrorInterface) {
	return s.storage.GetSchoolByID(nil, schoolID)
}
