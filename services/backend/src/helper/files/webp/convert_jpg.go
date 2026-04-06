package webp

import (
	"bytes"
	"image/jpeg"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func ConvertJPGToWebP(jpgData []byte) ([]byte, customerrors.ErrorInterface) {
	// Decode the JPG image
	img, err := jpeg.Decode(bytes.NewReader(jpgData))
	if err != nil {
		return nil, customerrors.NewInternalServerError(err, "failed to decode JPG image", "")
	}

	// Configure WebP encoding options
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 80)
	if err != nil {
		return nil, customerrors.NewInternalServerError(err, "failed to create WebP encoder options", "")
	}

	// Encode the image to WebP format
	var webpBuffer bytes.Buffer
	err = webp.Encode(&webpBuffer, img, options)
	if err != nil {
		return nil, customerrors.NewInternalServerError(err, "failed to encode image to WebP", "")
	}

	return webpBuffer.Bytes(), nil
}
