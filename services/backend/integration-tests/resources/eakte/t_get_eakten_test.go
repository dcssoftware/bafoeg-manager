package eakte

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/http/models"
	"github.com/stretchr/testify/assert"
)

func TestGetEakten(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	var eaktenStruct models.EaktenResponseModels
	testSetup.Request(http.MethodGet, "/api/v1/eakten?page=1", nil, &eaktenStruct)
	assert.Empty(t, eaktenStruct.Eakten, "should be empty applications")
}
