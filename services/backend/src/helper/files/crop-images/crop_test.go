package cropimages

import (
	"bytes"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// loadTestImage loads a test image from the assets directory
func loadTestImage() ([]byte, error) {
	file, err := os.Open("assets/bafög-1.png")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	// Encode the image to bytes
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// TestCropImage tests the CropImage function with valid parameters
func TestCropImage(t *testing.T) {
	// Load a test image
	imgData, err := loadTestImage()
	assert.NoError(t, err)
	assert.NotNil(t, imgData)

	// Test cropping with valid parameters
	croppedImg, cropErr := CropImage(imgData, 10, 10, 100, 100, "png")
	assert.NoError(t, cropErr)
	assert.NotNil(t, croppedImg)

	// Export cropped image to filesystem
	err = os.WriteFile(filepath.Join("testresults", "cropped_image.png"), croppedImg, 0644)
	assert.NoError(t, err)

	// Verify the cropped image can be decoded
	decodedImg, _, decodeErr := image.DecodeConfig(bytes.NewReader(croppedImg))
	assert.NoError(t, decodeErr)
	assert.Equal(t, 100, decodedImg.Width)
	assert.Equal(t, 100, decodedImg.Height)
}

// TestCropImageJpeg tests the CropImage function with JPEG format
func TestCropImageJpeg(t *testing.T) {
	// Load a test image
	imgData, err := loadTestImage()
	assert.NoError(t, err)
	assert.NotNil(t, imgData)

	// Test cropping with JPEG format
	croppedImg, cropErr := CropImage(imgData, 5, 5, 50, 50, "jpeg")
	assert.NoError(t, cropErr)
	assert.NotNil(t, croppedImg)

	// Export cropped image to filesystem
	err = os.WriteFile(filepath.Join("testresults", "cropped_image_jpeg.jpg"), croppedImg, 0644)
	assert.NoError(t, err)

	// Verify the cropped image can be decoded
	decodedImg, format, decodeErr := image.DecodeConfig(bytes.NewReader(croppedImg))
	assert.NoError(t, decodeErr)
	assert.Equal(t, "jpeg", format)
	assert.Equal(t, 50, decodedImg.Width)
	assert.Equal(t, 50, decodedImg.Height)
}

// TestCropImageInvalidParams tests the CropImage function with invalid parameters
func TestCropImageInvalidParams(t *testing.T) {
	// Load a test image
	imgData, err := loadTestImage()
	assert.NoError(t, err)
	assert.NotNil(t, imgData)

	// Test with negative coordinates
	croppedImg, cropErr := CropImage(imgData, -10, -10, 100, 100, "png")
	assert.Error(t, cropErr)
	assert.Nil(t, croppedImg)
	assert.Contains(t, cropErr.Error(), "crop coordinates and dimensions must be positive")

	// Test with negative dimensions
	croppedImg, cropErr = CropImage(imgData, 10, 10, -100, -100, "png")
	assert.Error(t, cropErr)
	assert.Nil(t, croppedImg)
	assert.Contains(t, cropErr.Error(), "crop coordinates and dimensions must be positive")

	// Test with zero dimensions
	croppedImg, cropErr = CropImage(imgData, 10, 10, 0, 0, "png")
	assert.Error(t, cropErr)
	assert.Nil(t, croppedImg)
	assert.Contains(t, cropErr.Error(), "crop coordinates and dimensions must be positive")

	// Test with crop area outside image bounds
	// First, get the image dimensions
	imgConfig, _, configErr := image.DecodeConfig(bytes.NewReader(imgData))
	assert.NoError(t, configErr)

	// Try to crop outside the image bounds
	croppedImg, cropErr = CropImage(imgData, imgConfig.Width-10, imgConfig.Height-10, 100, 100, "png")
	assert.Error(t, cropErr)
	assert.Nil(t, croppedImg)
	assert.Contains(t, cropErr.Error(), "crop area must be within image dimensions")
}

// TestCropAndResizeImage tests the CropAndResizeImage function
func TestCropAndResizeImage(t *testing.T) {
	// Load a test image
	imgData, err := loadTestImage()
	assert.NoError(t, err)
	assert.NotNil(t, imgData)

	// Test cropping and resizing with valid parameters
	croppedResizedImg, cropResizeErr := CropAndResizeImage(imgData, 10, 10, 100, 100, 50, 50, "png")
	assert.NoError(t, cropResizeErr)
	assert.NotNil(t, croppedResizedImg)

	// Export cropped and resized image to filesystem
	err = os.WriteFile(filepath.Join("testresults", "cropped_resized_image.png"), croppedResizedImg, 0644)
	assert.NoError(t, err)

	// Verify the cropped and resized image has the expected dimensions
	decodedImg, _, decodeErr := image.DecodeConfig(bytes.NewReader(croppedResizedImg))
	assert.NoError(t, decodeErr)
	assert.Equal(t, 50, decodedImg.Width)
	assert.Equal(t, 50, decodedImg.Height)
}

// TestCropAndResizeImageJpeg tests the CropAndResizeImage function with JPEG format
func TestCropAndResizeImageJpeg(t *testing.T) {
	// Load a test image
	imgData, err := loadTestImage()
	assert.NoError(t, err)
	assert.NotNil(t, imgData)

	// Test cropping and resizing with JPEG format
	croppedResizedImg, cropResizeErr := CropAndResizeImage(imgData, 5, 5, 80, 80, 40, 40, "jpeg")
	assert.NoError(t, cropResizeErr)
	assert.NotNil(t, croppedResizedImg)

	// Export cropped and resized image to filesystem
	err = os.WriteFile(filepath.Join("testresults", "cropped_resized_image_jpeg.jpg"), croppedResizedImg, 0644)
	assert.NoError(t, err)

	// Verify the cropped and resized image has the expected dimensions
	decodedImg, format, decodeErr := image.DecodeConfig(bytes.NewReader(croppedResizedImg))
	assert.NoError(t, decodeErr)
	assert.Equal(t, "jpeg", format)
	assert.Equal(t, 40, decodedImg.Width)
	assert.Equal(t, 40, decodedImg.Height)
}
