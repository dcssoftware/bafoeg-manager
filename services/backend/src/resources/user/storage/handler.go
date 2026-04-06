package storage

import (
	"encoding/json"
	"errors"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/minio/minio-go/v7"
)

type UserStore struct {
	db *sqlx.DB
	s3 *minio.Client
}

func NewUserStore(db *sqlx.DB, s3 *minio.Client) *UserStore {
	return &UserStore{
		db,
		s3,
	}
}

func (s *UserStore) StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, customerrors.NewDatabaseTransactionBeginError(err, "Failed to start transaction")
	}
	return tx, nil
}

func (permission *Permission) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &permission)
}

func (s *UserPermissions) ToServiceModel() *serviceModel.UserPermissionsModel {
	return &serviceModel.UserPermissionsModel{
		Permissions: PermissionsToStringArray(s.Permissions),
	}
}

func PermissionsToStringArray(input []Permission) []string {
	permissionArray := make([]string, 0)

	for _, permission := range input {
		permissionArray = append(permissionArray, string(permission))
	}

	return permissionArray
}
