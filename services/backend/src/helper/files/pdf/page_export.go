package pdf

import (
	"bytes"
	"io"
	"sort"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	pdfcpuModel "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ExportPDFPages(file []byte) (map[int][]byte, error) {
	f := bytes.NewReader(file)
	var result map[int][]byte = make(map[int][]byte)

	conf := pdfcpuModel.NewDefaultConfiguration()
	conf.Cmd = pdfcpuModel.EXTRACTPAGES

	ctx, readAllErr := api.ReadValidateAndOptimize(f, conf)
	if readAllErr != nil {
		return nil, readAllErr
	}

	pages, readAllErr := api.PagesForPageSelection(ctx.PageCount, []string{}, true, true)
	if readAllErr != nil {
		return nil, readAllErr
	}

	if len(pages) == 0 {
		return nil, nil
	}

	for _, pageNumber := range sortedPages(pages) {
		reader, readerErr := api.ExtractPage(ctx, pageNumber)
		if readerErr != nil {
			return nil, readerErr
		}

		result[pageNumber], readAllErr = io.ReadAll(reader)
		if readAllErr != nil {
			return nil, readAllErr
		}

		// os.WriteFile(fmt.Sprintf("./tmp/page_%d.pdf", pageNumber), result[pageNumber], 0644)
	}

	return result, nil
}

func sortedPages(selectedPages types.IntSet) []int {
	p := []int(nil)
	for i, v := range selectedPages {
		if v {
			p = append(p, i)
		}
	}
	sort.Ints(p)
	return p
}
