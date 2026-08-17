package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/states"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) UpdateApplicationStatus(tx *sqlx.Tx, applicationID string, newStatus states.ApplicationState) customerrors.ErrorInterface {

	oldApplication, oldApplicationErr := s.GetApplicationByID(tx, applicationID)
	if oldApplicationErr != nil {
		return oldApplicationErr
	}

	oldStatus, oldStatusErr := states.ConvertStrToApplicationState(oldApplication.Status.Identifier)
	if oldStatusErr != nil {
		return customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_ApplicationStatusInvalid)
	}

	statusChangeAllowedErr := states.CheckStatusChange(oldStatus, newStatus)
	if statusChangeAllowedErr != nil {
		return statusChangeAllowedErr
	}

	statusChangeErr := s.storage.UpdateApplicationStatus(tx, oldApplication.ID.String(), newStatus)
	if statusChangeErr != nil {
		return statusChangeErr
	}

	return nil
}
