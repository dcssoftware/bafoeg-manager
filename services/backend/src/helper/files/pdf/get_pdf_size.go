package pdf

import (
	"io"
	"os"
)

// GetPDFSize returns the size of the provided PDF reader in bytes.
// It supports common types like *os.File, bytes.Reader, strings.Reader,
// and falls back to scanning via ReadAt if necessary.
func GetPDFSize(pdfFile io.ReaderAt) (int64, error) {
	// Fast path: readers that expose Size()
	type sizer interface{ Size() int64 }
	if s, ok := pdfFile.(sizer); ok {
		return s.Size(), nil
	}

	// Files that expose Stat()
	type statter interface{ Stat() (os.FileInfo, error) }
	if st, ok := pdfFile.(statter); ok {
		fi, err := st.Stat()
		if err != nil {
			return 0, err
		}
		return fi.Size(), nil
	}

	// Readers that can seek: seek to end and back to restore position
	if rs, ok := pdfFile.(io.ReadSeeker); ok {
		cur, err := rs.Seek(0, io.SeekCurrent)
		if err != nil {
			return 0, err
		}
		end, err := rs.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, err
		}
		_, err = rs.Seek(cur, io.SeekStart)
		if err != nil {
			return 0, err
		}
		return end, nil
	}

	// Fallback: iterate using ReadAt in chunks until EOF
	const chunk = 1 << 20 // 1 MiB
	buf := make([]byte, chunk)
	var off int64
	for {
		n, err := pdfFile.ReadAt(buf, off)
		off += int64(n)
		if err == io.EOF {
			return off, nil
		}
		if err != nil {
			return 0, err
		}
		if n == 0 { // defensive: avoid infinite loop
			return off, nil
		}
	}
}
