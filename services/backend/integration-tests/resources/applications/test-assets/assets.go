package assets

import _ "embed"

//go:embed example-documents/infected/example01.pdf
var ExampleDocumentPDF01 []byte

//go:embed example-documents/infected/example02.pdf
var ExampleDocumentPDF02 []byte

//go:embed example-documents/infected/example-infected01.txt
var ExampleDocumentTxtInfected []byte

//go:embed example-documents/infected/example-infected01.zip
var ExampleDocumentZipInfected01 []byte

//go:embed example-documents/infected/example-infected02.zip
var ExampleDocumentTxtInfected02 []byte

//go:embed example-documents/xdomea-messages/xdomea-message-stupid.zip
var XDomeaMessageStupid []byte

//go:embed example-documents/files/example-file.pdf
var ExampleFilePDF []byte
