package pdf

import (
	_ "embed"
)

//go:embed testdata/01.pdf
var testFile []byte
