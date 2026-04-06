package applications

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	httpModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestGetApplicationByID ensures retrieving an application by ID returns the expected payload
func TestGetApplicationByID_Scenarios(t *testing.T) {
	type testCase struct {
		name           string
		pathBuilder    func(t *testing.T, ts *integrationtestsetup.TestInstance) string
		expectIDExists bool
	}

	cases := []testCase{
		{
			name: "happy path",
			pathBuilder: func(t *testing.T, ts *integrationtestsetup.TestInstance) string {
				appID, _, _ := insertBasicApplication(t, ts)
				return "/api/v1/applications/" + appID.String()
			},
			expectIDExists: true,
		},
		{
			name: "bad request invalid uuid",
			pathBuilder: func(t *testing.T, ts *integrationtestsetup.TestInstance) string {
				return "/api/v1/applications/not-a-uuid"
			},
			expectIDExists: false,
		},
		{
			name: "not found unknown id",
			pathBuilder: func(t *testing.T, ts *integrationtestsetup.TestInstance) string {
				return "/api/v1/applications/" + uuid.New().String()
			},
			expectIDExists: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testSetup := integrationtestsetup.SetupTest(t)
			defer testSetup.Cleanup()

			var resp httpModel.ApplicationModel
			testSetup.Request(http.MethodGet, tc.pathBuilder(t, testSetup), nil, &resp)

			if tc.expectIDExists {
				assert.NotEqual(t, uuid.Nil, resp.ID, "expected a valid application id")
			} else {
				assert.Equal(t, uuid.Nil, resp.ID, "expected empty application id")
			}
		})
	}
}
