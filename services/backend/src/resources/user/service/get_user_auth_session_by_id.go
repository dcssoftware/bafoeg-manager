package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
)

func (s *UserService) GetUserAuthSessionByID(
	tx *sqlx.Tx,
	sessionID string,
) (
	*serviceModel.SessionModel,
	customerrors.ErrorInterface,
) {
	return s.storage.GetAuthSessionByID(tx, sessionID)
}
