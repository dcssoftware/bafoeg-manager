package states

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
)

func CheckStatusChange(old, new ApplicationState) customerrors.ErrorInterface {

	if old == StatusInProgress {
		return checkStatusChangeFromInProcess(new)
	}
	return customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_ApplicationStatusChangeInvalid)
}

func checkStatusChangeFromInProcess(new ApplicationState) customerrors.ErrorInterface {

	if new == StatusResponseAwaited ||
		new == StatusApproved ||
		new == StatusDenied {
		return nil
	}

	return customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_ApplicationStatusChangeInvalid)
}
