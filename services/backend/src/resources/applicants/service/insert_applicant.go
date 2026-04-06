package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicantService) InsertApplicant(tx *sqlx.Tx, model models.CreateApplicantModel) (*models.ApplicantModel, customerrors.ErrorInterface) {

	insertedAddressID, insertedAddressErr := s.InsertApplicantAddress(tx, models.ApplicantAddressModel{
		Street:      model.Street,
		HouseNumber: model.HouseNumber,
		ZipCode:     model.ZipCode,
		City:        model.City,
		Country:     model.Country,
	})

	if insertedAddressErr != nil {
		return nil, insertedAddressErr
	}

	insertedContactInfoID, insertedContactInfoErr := s.InsertApplicantContactInformation(tx, models.ApplicantContactModel{
		Phone: "",
		Email: "",
	})
	if insertedContactInfoErr != nil {
		return nil, insertedContactInfoErr
	}

	insertedApplicantID, insertedApplicantErr := s.storage.InsertApplicant(tx, models.ApplicantInsertModel{
		Firstname: model.Firstname,
		Lastname:  model.Lastname,

		AddressID:     insertedAddressID,
		ContactID:     insertedContactInfoID,
		BankAccountID: uuid.Nil,
	})
	if insertedApplicantErr != nil {
		return nil, insertedApplicantErr
	}

	return s.GetApplicantByID(tx, insertedApplicantID.String())
}
