package qrcode

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed assets/bafög-1.png
var scanBaföGAntrag01Version01 []byte

func TestReadDocumentQRCode(t *testing.T) {

	result, err := ReadDocumentQRCode(bytes.NewReader(scanBaföGAntrag01Version01))
	assert.NoError(t, err)

	assert.Equal(t, "Formblatt 01", result.String())
}
