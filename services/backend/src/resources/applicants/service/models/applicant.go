package models

import "github.com/google/uuid"

type ApplicantModel struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string

	BalancePayback  *float64
	BalanceOutgoing *float64

	Address          *ApplicantAddressModel
	TrainingsAddress *ApplicantTrainingsAddressModel
	Contact          *ApplicantContactModel
}

type ApplicantInsertModel struct {
	Firstname string
	Lastname  string

	AddressID     uuid.UUID
	ContactID     uuid.UUID
	BankAccountID uuid.UUID
}

type ApplicantBankAccountModel struct {
	Iban          string
	Bic           string
	BankName      string
	AccountHolder string
}

type ApplicantAddressModel struct {
	Street      string
	HouseNumber string
	ZipCode     string
	City        string
	Country     string
}

type ApplicantTrainingsAddressModel struct {
	Street      string
	HouseNumber string
	ZipCode     string
	City        string
	Country     string
}

type ApplicantContactModel struct {
	Phone string
	Email string
}
