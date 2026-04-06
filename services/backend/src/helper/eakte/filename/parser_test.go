package filename

import (
	"testing"

	xdomeaconsts "github.com/dcssoftware/bafoeg-manager/src/helper/eakte/versions/3.0.0/const"
	"github.com/stretchr/testify/assert"
)

func TestParseXDomeaFileName(t *testing.T) {
	erg, err := ParseContainerDateiname("eff9c1ec-57f2-4195-ba70-034dab4868ab_0101_1.xdomea")
	assert.NoError(t, err)
	assert.Equal(t, erg.ProzessID, "eff9c1ec-57f2-4195-ba70-034dab4868ab", "process ID")
	assert.Equal(t, erg.Nachrichtengruppe, "", "Nachrichtengruppe")
	assert.Equal(t, erg.Nachrichtenname, "", "Nachrichtenname")
	assert.Equal(t, erg.MessageType, xdomeaconsts.MessageType_Information_Information.GetCode(), "MessageType")
	assert.Equal(t, erg.Nachrichtentyp, "", "Nachrichtentyp")
	assert.Equal(t, erg.Extension, xdomeaconsts.ContainerFileTypeXdomea.GetNameLowercase(), "Extension")

	erg2, err2 := ParseContainerDateiname("e25cf421-d7e2-4df0-ae8f-e8cd764cd2cc_Aussonderung.Anbieteverzeichnis.0501.zip")
	assert.NoError(t, err2)
	assert.Equal(t, erg2.ProzessID, "e25cf421-d7e2-4df0-ae8f-e8cd764cd2cc", "ProzessID")
	assert.Equal(t, erg2.Nachrichtengruppe, "Aussonderung", "Nachrichtengruppe")
	assert.Equal(t, erg2.Nachrichtenname, "Anbieteverzeichnis", "Nachrichtenname")
	assert.Equal(t, erg2.MessageType, xdomeaconsts.MessageType_Aussonderung_Anbieteverzeichnis.GetCode(), "MessageType")
	assert.Equal(t, erg2.Nachrichtentyp, "Aussonderung.Anbieteverzeichnis.0501", "Nachrichtentyp")
	assert.Equal(t, erg2.Extension, xdomeaconsts.ContainerFileTypeZip.GetNameLowercase(), "Extension")
}
