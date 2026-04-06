package applications

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	applicationStates "github.com/dcssoftware/bafoeg-manager/src/resources/applications/states"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestUpdateApplicationStatus covers happy path and validation error cases
func TestUpdateApplicationStatus(t *testing.T) {

	type testCase struct {
		name        string
		pathBuilder func(t *testing.T, ts *integrationtestsetup.TestInstance) string
	}

	cases := []testCase{
		{
			name: "happy path",
			pathBuilder: func(t *testing.T, ts *integrationtestsetup.TestInstance) string {
				appID, _, _ := insertBasicApplication(t, ts)
				return "/api/v1/applications/" + appID.String() + "/change-status/" + applicationStates.StatusResponseAwaited.ToString()
			},
		},
		{
			name: "invalid uuid",
			pathBuilder: func(t *testing.T, ts *integrationtestsetup.TestInstance) string {
				return "/api/v1/applications/not-a-uuid/change-status/" + applicationStates.StatusDenied.ToString()
			},
		},
		{
			name: "invalid status",
			pathBuilder: func(t *testing.T, ts *integrationtestsetup.TestInstance) string {
				appID, _, _ := insertBasicApplication(t, ts)
				return "/api/v1/applications/" + appID.String() + "/change-status/DOES_NOT_EXIST"
			},
		},
		{
			name: "not found unknown id valid status",
			pathBuilder: func(t *testing.T, ts *integrationtestsetup.TestInstance) string {
				return "/api/v1/applications/" + uuid.New().String() + "/change-status/" + applicationStates.StatusResponseAwaited.ToString()
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testSetup := integrationtestsetup.SetupTest(t)
			defer testSetup.Cleanup()

			var resp struct{}
			testSetup.Request(http.MethodPatch, tc.pathBuilder(t, testSetup), nil, &resp)
			assert.Equal(t, struct{}{}, resp)
		})
	}
}
