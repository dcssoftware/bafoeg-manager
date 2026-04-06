package models

import (
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
)

type SelfInformationModel struct {
	ID             string   `json:"id"`
	Username       string   `json:"username"`
	DisplayName    string   `json:"displayName"`
	Email          string   `json:"email"`
	ProfilePicture string   `json:"profilePicture"`
	Permissions    []string `json:"permissions"`
}

func NewSelfInformationModel(model *serviceModel.UserSelfInformationModel) *SelfInformationModel {
	return &SelfInformationModel{
		ID:             model.User.ID,
		Username:       model.User.Username,
		DisplayName:    model.User.DisplayName,
		Email:          model.User.Email,
		ProfilePicture: "/api/v1/self/picture",
		Permissions:    model.Permissions,
	}
}
