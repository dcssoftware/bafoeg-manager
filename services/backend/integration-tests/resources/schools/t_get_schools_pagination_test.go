package schools

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/http/models"
	"github.com/stretchr/testify/assert"
)

// Tests pagination behaviour of GET /api/v1/schools endpoint.
// Page param defaults to 1 if not provided. With provided seed data we
// expect exactly 1 school on page 1 and 0 schools on a higher (empty) page.
func TestGetSchoolsPagination(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	// Page 1 (explicit)
	var schoolsPage1 models.SchoolShortResponseModel
	testSetup.Request(http.MethodGet, "/api/v1/schools?page=1", nil, &schoolsPage1)
	assert.Len(t, schoolsPage1.Schools, 1, "expected exactly one school on page 1")

	// Page 2 should be empty because only one seeded school exists
	var schoolsPage2 models.SchoolShortResponseModel
	testSetup.Request(http.MethodGet, "/api/v1/schools?page=2", nil, &schoolsPage2)
	assert.Len(t, schoolsPage2.Schools, 0, "expected no schools on page 2")
}
