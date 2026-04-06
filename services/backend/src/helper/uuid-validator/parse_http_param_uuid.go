package uuidvalidator

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	"github.com/google/uuid"
)

func ParseHttpParamUUID(param string) (uuid.UUID, customerrors.ErrorInterface) {
	if !ValidateUUID(param) {
		return uuid.Nil, customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_IdNotUUID)
	}
	parsedUUID, err := uuid.Parse(param)
	if err != nil {
		return uuid.Nil, customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_IdNotUUID)
	}
	return parsedUUID, nil
}
