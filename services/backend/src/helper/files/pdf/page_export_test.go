package pdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPDFPageExtraction(t *testing.T) {
	_, pagesBytesErr := ExportPDFPages([]byte{})
	assert.Error(t, pagesBytesErr, "should error on empty input")
}

func TestPDFPageExtractionValidFile(t *testing.T) {
	pagesBytes, pagesBytesErr := ExportPDFPages(testFile)
	assert.NoError(t, pagesBytesErr, "should not error on valid input")
	assert.Equal(t, 8, len(pagesBytes), "should extract eight pages from test file")
}
