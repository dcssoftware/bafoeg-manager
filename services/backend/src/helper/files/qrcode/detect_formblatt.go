package qrcode

import "bytes"

type FormBlattVersion string

const (
	Formblatt01Old   FormBlattVersion = "FORMBLATT_01_OLD"
	FormblattUnknown FormBlattVersion = "FORMBLATT_UNKNOWN"
)

func (v FormBlattVersion) String() string {
	return string(v)
}

func DetectFormBlatt(input []byte) (*FormBlattVersion, error) {

	result, resultErr := ReadDocumentQRCode(
		bytes.NewReader(input),
	)
	if resultErr != nil {
		return nil, resultErr
	}

	identifier := result.GetText()
	var value FormBlattVersion = FormblattUnknown

	switch identifier {
	case "Formblatt 01":
		value = Formblatt01Old
	}

	return &value, nil
}
