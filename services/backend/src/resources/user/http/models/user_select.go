package models

import serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"

type UserSelectModel struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

func ToHttpUserSelectModel(model *serviceModel.UserModel) *UserSelectModel {
	if model == nil {
		return nil
	}

	return &UserSelectModel{
		ID:          model.ID,
		Username:    model.Username,
		DisplayName: model.DisplayName,
	}
}

func ToHttpUserSelectModels(models []serviceModel.UserModel) []*UserSelectModel {
	var httpModels []*UserSelectModel

	for _, m := range models {
		httpModels = append(httpModels, ToHttpUserSelectModel(&m))
	}

	return httpModels
}
