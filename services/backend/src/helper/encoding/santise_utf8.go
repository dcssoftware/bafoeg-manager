package encoding

import (
	"strings"
	"unicode"
)

// sanitizeUTF8 removes NUL bytes (\x00) and most control characters except for
// common whitespace (tab, newline, carriage return). It preserves valid runes.
func SanitizeUTF8(s string) string {
	if s == "" {
		return s
	}
	// Fast path: remove explicit NUL bytes first
	if strings.Contains(s, "\x00") {
		s = strings.ReplaceAll(s, "\x00", "")
	}
	// Remove other control characters except \n, \r, \t
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		// Keep standard whitespace
		if r == '\n' || r == '\r' || r == '\t' {
			b.WriteRune(r)
			continue
		}
		if unicode.IsControl(r) {
			// drop other control chars
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
