package models

import serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"

type UserModel struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	DisplayName   string `json:"displayName"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
}

func ToHttpUserModel(model *serviceModel.UserModel) *UserModel {
	if model == nil {
		return nil
	}

	return &UserModel{
		ID:            model.ID,
		Username:      model.Username,
		DisplayName:   model.DisplayName,
		Email:         model.Email,
		EmailVerified: model.EmailVerified,
	}
}

func ToHttpUserModels(models []serviceModel.UserModel) []*UserModel {
	var httpModels []*UserModel

	for _, m := range models {
		httpModels = append(httpModels, ToHttpUserModel(&m))
	}

	return httpModels
}
