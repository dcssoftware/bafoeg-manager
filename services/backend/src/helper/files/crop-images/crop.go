package cropimages

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	"golang.org/x/image/draw"
)

// CropImage takes an image byte slice, crops it to the specified coordinates, and returns the cropped image
// Parameters:
//   - imageData: the original image data as byte slice
//   - x, y: the top-left corner coordinates of the crop area
//   - width, height: the dimensions of the crop area
//   - format: the format of the output image ("jpeg", "png", etc.)
//
// Returns:
//   - []byte: the cropped image data
//   - customerrors.ErrorInterface: error if any
func CropImage(imageData []byte, x, y, width, height int, format string) ([]byte, customerrors.ErrorInterface) {
	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return nil, customerrors.NewInternalServerError(err, "failed to decode image", "")
	}

	// Validate crop coordinates and dimensions
	imgBounds := img.Bounds()
	if x < 0 || y < 0 || width <= 0 || height <= 0 {
		constraint := &custombadrequestconstraints.BadRequestContraint{
			KeyName:           "crop_parameters",
			ConstraintMessage: "crop coordinates and dimensions must be positive",
		}
		return nil, customerrors.NewBadRequestError(constraint)
	}

	if x+width > imgBounds.Max.X || y+height > imgBounds.Max.Y {
		constraint := &custombadrequestconstraints.BadRequestContraint{
			KeyName:           "crop_area",
			ConstraintMessage: "crop area must be within image dimensions",
		}
		return nil, customerrors.NewBadRequestError(constraint)
	}

	// Create a rectangle for the crop area
	cropRect := image.Rect(x, y, x+width, y+height)

	// Crop the image using the SubImage method
	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(cropRect)

	// Encode the cropped image to the specified format
	var buf bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, croppedImg, nil)
		if err != nil {
			return nil, customerrors.NewInternalServerError(err, "failed to encode cropped image as JPEG", "")
		}
	case "png":
		err = png.Encode(&buf, croppedImg)
		if err != nil {
			return nil, customerrors.NewInternalServerError(err, "failed to encode cropped image as PNG", "")
		}
	default:
		// If format is not specified or unsupported, try to preserve original format
		// For simplicity, we'll default to PNG
		err = png.Encode(&buf, croppedImg)
		if err != nil {
			return nil, customerrors.NewInternalServerError(err, "failed to encode cropped image", "")
		}
	}

	return buf.Bytes(), nil
}

// CropAndResizeImage takes an image byte slice, crops it to the specified coordinates,
// and then resizes it to the target dimensions
// Parameters:
//   - imageData: the original image data as byte slice
//   - cropX, cropY: the top-left corner coordinates of the crop area
//   - cropWidth, cropHeight: the dimensions of the crop area
//   - targetWidth, targetHeight: the target dimensions for the resized image
//   - format: the format of the output image ("jpeg", "png", etc.)
//
// Returns:
//   - []byte: the cropped and resized image data
//   - customerrors.ErrorInterface: error if any
func CropAndResizeImage(imageData []byte, cropX, cropY, cropWidth, cropHeight, targetWidth, targetHeight int, format string) ([]byte, customerrors.ErrorInterface) {
	// First crop the image
	croppedImg, err := CropImage(imageData, cropX, cropY, cropWidth, cropHeight, format)
	if err != nil {
		return nil, err
	}

	// Decode the cropped image for resizing
	img, _, decodeErr := image.Decode(bytes.NewReader(croppedImg))
	if decodeErr != nil {
		return nil, customerrors.NewInternalServerError(decodeErr, "failed to decode cropped image", "")
	}

	// Create a new RGBA image with target dimensions
	resizedImg := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	// Resize the image using draw.CatmullRom for high-quality resampling
	draw.CatmullRom.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	// Encode the resized image to the specified format
	var buf bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		encodeErr := jpeg.Encode(&buf, resizedImg, nil)
		if encodeErr != nil {
			return nil, customerrors.NewInternalServerError(encodeErr, "failed to encode resized image as JPEG", "")
		}
	case "png":
		encodeErr := png.Encode(&buf, resizedImg)
		if encodeErr != nil {
			return nil, customerrors.NewInternalServerError(encodeErr, "failed to encode resized image as PNG", "")
		}
	default:
		// Default to PNG if format is not specified or unsupported
		encodeErr := png.Encode(&buf, resizedImg)
		if encodeErr != nil {
			return nil, customerrors.NewInternalServerError(encodeErr, "failed to encode resized image", "")
		}
	}

	return buf.Bytes(), nil
}
