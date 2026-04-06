package schools

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/http/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSchoolByID(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	var schoolStruct models.SchoolShortResponseModel
	testSetup.Request(http.MethodGet, "/api/v1/schools", nil, &schoolStruct)
	assert.Len(t, schoolStruct.Schools, 1, "should be empty schools")

	schoolID := schoolStruct.Schools[0].ID
	var schoolByIDStruct models.SchoolModel
	testSetup.Request(http.MethodGet, "/api/v1/schools/"+schoolID.String(), nil, &schoolByIDStruct)

	assert.Equal(t, "other_school", schoolByIDStruct.Type.Identifier, "")
	assert.Equal(t, "Sonstige", schoolByIDStruct.Type.Name, "")
	assert.Equal(t, "Teststraße", schoolByIDStruct.Street, "")
	assert.Equal(t, "1a", schoolByIDStruct.HouseNumber, "")
	assert.Equal(t, "12345", schoolByIDStruct.ZipCode, "")
	assert.Equal(t, "Teststadt", schoolByIDStruct.City, "")
	assert.Equal(t, "Testland", schoolByIDStruct.Country, "")

	var degreenames []string
	for _, degree := range schoolByIDStruct.Degree {
		degreenames = append(degreenames, degree.Name)
	}

	assert.Contains(t, degreenames, "Test Abschluss (example degree)", "")

	// Assert degree flags of the first (and only) degree
	if assert.GreaterOrEqual(t, len(schoolByIDStruct.Degree), 1, "expected at least one degree") {
		degree := schoolByIDStruct.Degree[0]
		assert.Equal(t, false, degree.FosBerufsabschlussRequired, "unexpected FosBerufsabschlussRequired")
		assert.Equal(t, true, degree.BosBerufsqualifizierenderAbschluss, "unexpected BosBerufsqualifizierenderAbschluss")
		assert.Equal(t, false, degree.FachschuleBerufsschuleRequired, "unexpected FachschuleBerufsschuleRequired")
	}
}
