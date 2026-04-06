package pdf

import (
	"io"

	"github.com/ledongthuc/pdf"
)

type schemaDocument struct {
	PageContent string
	Metadata    struct {
		Page       int
		TotalPages int
	}
}

func ReadPlainText(fileReader io.ReaderAt) ([]schemaDocument, error) {
	var reader *pdf.Reader
	var err error

	fileSize, fileSizeErr := GetPDFSize(fileReader)
	if fileSizeErr != nil {
		return nil, fileSizeErr
	}

	// if p.password != "" {
	// 	reader, err = pdf.NewReaderEncrypted(p.r, p.s, p.getPassword)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// } else {

	reader, err = pdf.NewReader(fileReader, fileSize)
	if err != nil {
		return nil, err
	}

	// }

	numPages := reader.NumPage()

	docs := []schemaDocument{}

	// fonts to be used when getting plain text from pages
	fonts := make(map[string]*pdf.Font)
	for i := 1; i < numPages+1; i++ {
		p := reader.Page(i)
		// add fonts to map
		for _, name := range p.Fonts() {
			// only add the font if we don't already have it
			if _, ok := fonts[name]; !ok {
				f := p.Font(name)
				fonts[name] = &f
			}
		}
		text, err := p.GetPlainText(fonts)
		if err != nil {
			return nil, err
		}

		// add the document to the doc list
		docs = append(docs, schemaDocument{
			PageContent: text,
			Metadata: struct {
				Page       int
				TotalPages int
			}{
				Page:       i,
				TotalPages: numPages,
			},
		})
	}

	return docs, nil
}
