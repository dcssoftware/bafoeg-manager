package service

import (
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/states"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) UpdateApplicationStatus(tx *sqlx.Tx, applicationID string, newStatus states.ApplicationState) customerrors.ErrorInterface {

	oldApplication, oldApplicationErr := s.GetApplicationByID(tx, applicationID)
	if oldApplicationErr != nil {
		fmt.Println("Error: 1")
		return oldApplicationErr
	}

	oldStatus, oldStatusErr := states.ConvertStrToApplicationState(oldApplication.Status.Identifier)
	if oldStatusErr != nil {
		fmt.Println("Error: 2")
		return customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_ApplicationStatusInvalid)
	}

	statusChangeAllowedErr := states.CheckStatusChange(oldStatus, newStatus)
	if statusChangeAllowedErr != nil {
		fmt.Println("Error: 3")
		return statusChangeAllowedErr
	}

	statusChangeErr := s.storage.UpdateApplicationStatus(tx, oldApplication.ID.String(), newStatus)
	if statusChangeErr != nil {
		fmt.Println("Error: 4", statusChangeErr.Error())
		return statusChangeErr
	}

	return nil
}
