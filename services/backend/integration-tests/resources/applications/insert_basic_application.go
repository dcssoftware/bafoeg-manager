package applications

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"testing"
	"time"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	applicantServiceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	httpModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func insertBasicApplication(t *testing.T, testSetup *integrationtestsetup.TestInstance) (
	applicationID uuid.UUID,
	newApplicant *applicantServiceModel.ApplicantModel,
	appStatus []serviceModel.ApplicationStatus,
) {
	var applicationStructBeforeInsert httpModel.ApplicationShortModels
	testSetup.Request(http.MethodGet, "/api/v1/applications", nil, &applicationStructBeforeInsert)
	assert.Empty(t, applicationStructBeforeInsert.Application, "should be empty: applications")

	// ---- Contact Information

	newApplicantContactInformationModel := applicantServiceModel.ApplicantContactModel{
		Phone: faker.Phonenumber(),
		Email: faker.Email(),
	}
	newApplicantContactInformationID, newApplicantContactInformationErr := testSetup.AppServices.ApplicantSvc.InsertApplicantContactInformation(nil, newApplicantContactInformationModel)
	assert.NoError(t, newApplicantContactInformationErr, "could not create new applicant contact information")

	// ---- Bank Account

	newBankAccountModel := applicantServiceModel.ApplicantBankAccountModel{
		Iban:          "DE02120300000000202051",
		Bic:           "BYLADEM1001",
		BankName:      faker.Word(),
		AccountHolder: faker.Name(),
	}
	newApplicantBankAccountID, newApplicantBankAccountErr := testSetup.AppServices.ApplicantSvc.InsertApplicantBankAccount(nil, newBankAccountModel)
	assert.NoError(t, newApplicantBankAccountErr, "could not create new applicant bank account")

	// ---- Applicant Address

	newAddressModel := applicantServiceModel.ApplicantAddressModel{
		Street:      faker.GetRealAddress().Address,
		HouseNumber: fmt.Sprintf("%d", rand.Uint32()),
		ZipCode:     faker.GetRealAddress().PostalCode,
		City:        faker.GetRealAddress().City,
		Country:     faker.GetRealAddress().State,
	}
	newApplicantAddressID, newApplicantBankAccountErr := testSetup.AppServices.ApplicantSvc.InsertApplicantAddress(nil, newAddressModel)
	assert.NoError(t, newApplicantBankAccountErr, "could not create new applicant bank account")

	// ---- Applicant

	newApplicantModel := applicantServiceModel.ApplicantInsertModel{
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),

		AddressID:     newApplicantAddressID,
		ContactID:     newApplicantContactInformationID,
		BankAccountID: newApplicantBankAccountID,
	}
	newApplicant, newApplicantErr := testSetup.AppServices.ApplicantSvc.InsertApplicantUnittest(nil, newApplicantModel)
	assert.NoError(t, newApplicantErr, "could not create new applicant")

	// ---- School

	schoolDegrees, schoolDegreesErr := testSetup.AppServices.SchoolService.GetSchoolDegree(nil, 1)
	assert.NoError(t, schoolDegreesErr, "could not retrieve school degrees")
	assert.NotEmpty(t, schoolDegrees, "should not be empty: school degrees")

	// ---- Assigned User

	user, userErr := testSetup.AppServices.UserSvc.GetUsers(nil)
	assert.NoError(t, userErr, "could not retrieve assignable users")
	assert.NotEmpty(t, user, "should not be empty: assignable users")

	// ---- Application Status

	appStatus, appStatusErr := testSetup.AppServices.ApplicationsService.GetApplicationStatus(nil, 1)
	assert.NoError(t, appStatusErr, "could not retrieve application status")
	assert.NotEmpty(t, appStatus, "should not be empty: application status")

	// ---- Application

	educationStart, educationStartErr := time.Parse("2006-01-02 15:04:05", "2024-09-01 00:00:00")
	assert.NoError(t, educationStartErr, "could not parse application education start date")
	educationEnd, educationEndErr := time.Parse("2006-01-02 15:04:05", "2025-08-31 00:00:00")
	assert.NoError(t, educationEndErr, "could not parse application education end date")

	newApplication := serviceModel.ApplicationInsertModel{
		ClassLevel:         "10a 😅",
		LabelIDs:           []string{},
		ApplicantID:        newApplicant.ID.String(),
		SchoolWithDegreeID: schoolDegrees[0].ID.String(),
		AssignedUserID:     user[0].ID,
		StatusID:           appStatus[0].ID.String(),

		EducationStart: educationStart,
		EducationEnd:   educationEnd,
	}
	applicationID, cErr := testSetup.AppServices.ApplicationsService.InsertApplication(nil, newApplication)
	assert.Nil(t, cErr, "error on application insertion")

	return
}
