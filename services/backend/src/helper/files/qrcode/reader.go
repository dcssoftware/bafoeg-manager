package qrcode

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func ReadDocumentQRCode(input io.Reader) (*gozxing.Result, error) {
	img, _, imgErr := image.Decode(input)
	if imgErr != nil {
		return nil, imgErr
	}

	bmp, bmpErr := gozxing.NewBinaryBitmapFromImage(img)
	if bmpErr != nil {
		return nil, bmpErr
	}

	qrReader := qrcode.NewQRCodeReader()
	result, resultErr := qrReader.Decode(bmp, nil)
	if resultErr != nil {
		return nil, resultErr
	}

	return result, nil
}
