package webp

import (
	"bytes"
	"image/png"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func ConvertPngToWebP(pngData []byte) ([]byte, customerrors.ErrorInterface) {

	img, err := png.Decode(bytes.NewReader(pngData))
	if err != nil {
		return []byte(""), customerrors.NewInternalServerError(err, "failed to decode PNG image", "")
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 80)
	if err != nil {
		return []byte(""), customerrors.NewInternalServerError(err, "failed to create WebP encoder options", "")
	}

	var webpBuffer bytes.Buffer
	err = webp.Encode(&webpBuffer, img, options)
	if err != nil {
		return []byte(""), customerrors.NewInternalServerError(err, "failed to encode PNG to WebP", "")
	}

	return webpBuffer.Bytes(), nil
}
