package applicants

import (
	"fmt"
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	applicantModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/http/models"
	"github.com/stretchr/testify/assert"
)

// Table-driven test for GET /api/v1/applications/applicants pagination when database has no applicants.
func TestGetApplicants_EmptyPagination(t *testing.T) {
	tests := []struct {
		name          string
		page          int
		expectedCount uint
		expectedMax   uint
	}{
		{name: "page 1 empty", page: 1, expectedCount: 0, expectedMax: 0},
		{name: "page 2 empty", page: 2, expectedCount: 0, expectedMax: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testSetup := integrationtestsetup.SetupTest(t)
			defer testSetup.Cleanup()

			var resp applicantModels.ApplicantResponseModels
			testSetup.Request(http.MethodGet,
				"/api/v1/applications/applicants/?page="+intToString(tc.page),
				nil,
				&resp,
			)

			assert.Equal(t, tc.expectedCount, resp.Count, "unexpected count for page")
			assert.Equal(t, tc.expectedMax, resp.MaxCount, "unexpected maxCount for page")
			assert.Len(t, resp.Applicants, int(tc.expectedCount), "unexpected applicants length")
		})
	}
}

func intToString(i int) string { return fmt.Sprintf("%d", i) }
