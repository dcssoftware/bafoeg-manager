package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) GetProfilePictureByID(tx *sqlx.Tx, id string) ([]byte, customerrors.ErrorInterface) {
	user, userErr := s.storage.GetByID(tx, id)
	if userErr != nil {
		return nil, userErr
	}

	return s.fileService.DownloadFileProfilePicture(user.ID)
}
