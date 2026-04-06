package pdf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPDFSite(t *testing.T) {
	size, err := GetPDFSize(bytes.NewReader(testFile))
	assert.NoError(t, err, "should get size without error")
	assert.Equal(t, int64(966501), size, "size should match")
}
