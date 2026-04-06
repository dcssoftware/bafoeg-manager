package schools

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/http/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSchools(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	var schoolStruct models.SchoolShortResponseModel
	testSetup.Request(http.MethodGet, "/api/v1/schools", nil, &schoolStruct)
	assert.Len(t, schoolStruct.Schools, 1, "should be one school seeded")
}
