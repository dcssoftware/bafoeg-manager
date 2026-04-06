package applicants

import (
	"fmt"
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	applicantModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/http/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/http/models"
	"github.com/stretchr/testify/assert"
)

// Table-driven test for GET /api/v1/applications/applicants/by-school/{schoolID}?page=&isActive= when there are no applicants.
func TestGetApplicantsBySchool_Empty(t *testing.T) {
	cases := []struct {
		name     string
		isActive bool
	}{
		{name: "active applicants empty", isActive: true},
		{name: "inactive applicants empty", isActive: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			testSetup := integrationtestsetup.SetupTest(t)
			defer testSetup.Cleanup()

			// Fetch seeded school (should exist as verified by school tests)
			var schools models.SchoolShortResponseModel
			testSetup.Request(http.MethodGet, "/api/v1/schools", nil, &schools)
			if !assert.GreaterOrEqual(t, len(schools.Schools), 1, "expected at least one seeded school") {
				t.Skip("no seeded school to test")
			}
			schoolID := schools.Schools[0].ID.String()

			var resp applicantModels.ApplicantBySchoolModels
			testSetup.Request(
				http.MethodGet,
				fmt.Sprintf("/api/v1/applications/applicants/by-school/%s?page=1&isActive=%t", schoolID, c.isActive),
				nil,
				&resp,
			)

			assert.Equal(t, uint(0), resp.Count, "expected zero applicants count")
			assert.Equal(t, uint(0), resp.MaxCount, "expected zero maxCount")
			assert.Len(t, resp.Applicants, 0, "expected empty applicants slice")
		})
	}
}
