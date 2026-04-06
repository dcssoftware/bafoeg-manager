package service

import "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"

func (s *ApplicationsService) GetApplicationsCount(userID string, applicantID string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetApplicationsCount(nil, userID, applicantID)
}
