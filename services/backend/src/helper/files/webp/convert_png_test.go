package webp

import (
	"bytes"
	"embed"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/webp"
)

//go:embed assets/png/*
var pngAssets embed.FS

func TestConvertPNGToWebP(t *testing.T) {

	files, filesErr := jpgAssets.ReadDir(".")
	assert.NoError(t, filesErr)

	var testFiles []TestFiles
	for _, file := range files {

		extention := strings.ToLower(filepath.Ext(file.Name()))
		if !file.IsDir() && extention == ".png" {

			fileContent, fileContentErr := jpgAssets.ReadFile(file.Name())
			assert.NoError(t, fileContentErr)

			testFiles = append(testFiles, TestFiles{
				Name:    file.Name(),
				Content: fileContent,
			})

		}
	}

	for _, tt := range testFiles {
		t.Run("Test Convert "+tt.Name+" from jpg to webp", func(t *testing.T) {

			require.NotEmpty(t, tt.Content, "JPG file should not be empty: %s", tt.Name)

			webpData, err := ConvertPngToWebP(tt.Content)

			assert.NoError(t, err, "ConvertPngToWebP should not return an error for: %s", tt.Name)
			assert.NotEmpty(t, webpData, "WebP data should not be empty for: %s", tt.Name)

			_, webpImageErr := webp.Decode(bytes.NewReader(webpData))
			assert.NoError(t, webpImageErr)

			assert.True(t, len(webpData) > 12, "WebP data should be at least 13 bytes for: %s", tt.Name)
			assert.Equal(t, []byte("RIFF"), webpData[0:4], "WebP file should start with RIFF header for: %s", tt.Name)
			assert.Equal(t, []byte("WEBP"), webpData[8:12], "WebP file should contain WEBP identifier for: %s", tt.Name)
		})
	}
}
