package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"

	filesService "github.com/dcssoftware/bafoeg-manager/src/resources/files/service"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
)

type UserService struct {
	storage     UserStorageInterface
	storageS3   UserStorageS3Interface
	fileService *filesService.FileService
}

func NewUserService(storage UserStorageInterface, storageS3 UserStorageS3Interface, fileService *filesService.FileService) *UserService {
	return &UserService{
		storage:     storage,
		storageS3:   storageS3,
		fileService: fileService,
	}
}

type UserStorageInterface interface {
	StartTransaction() (tx *sqlx.Tx, err customerrors.ErrorInterface)
	Create(
		tx *sqlx.Tx,
		displayName string,
		username string,
		email string,
		emailVerified bool,
	) (string, customerrors.ErrorInterface)
	CreateLoginIdentity(
		tx *sqlx.Tx,
		createdUserID string,
		authProvider string,
		authProviderID string,
	) customerrors.ErrorInterface
	GetPermissionsByID(tx *sqlx.Tx, userID string) ([]string, customerrors.ErrorInterface)
	GetByID(tx *sqlx.Tx, userID string) (*serviceModel.UserModel, customerrors.ErrorInterface)
	GetByUsername(tx *sqlx.Tx, username string) (*serviceModel.UserModel, customerrors.ErrorInterface)
	GetAuthSessionByID(tx *sqlx.Tx, sessionID string) (*serviceModel.SessionModel, customerrors.ErrorInterface)
	GetIDByLogin(provider string, providerID string) (string, customerrors.ErrorInterface)
	GetUsers(tx *sqlx.Tx) ([]serviceModel.UserModel, customerrors.ErrorInterface)
	GetUsersCount(tx *sqlx.Tx) (uint, customerrors.ErrorInterface)
	GetUserAuthSessionByUserID(tx *sqlx.Tx, userID string) (*serviceModel.SessionModel, customerrors.ErrorInterface)
}

type UserStorageS3Interface interface {
	InsertProfilePictureByID(userID string, pictureJPG []byte) customerrors.ErrorInterface
	GetProfilePictureByID(id string) ([]byte, customerrors.ErrorInterface)
}

func (s *UserService) StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface) {
	return s.storage.StartTransaction()
}
