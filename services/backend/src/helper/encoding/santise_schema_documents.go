package encoding

import (
	"strings"

	"github.com/tmc/langchaingo/schema"
)

// sanitizeDocuments cleans each document's PageContent and string metadata to ensure
// Postgres-compatible UTF-8 (no NUL bytes) and removes other disallowed control characters.
// Empty documents (after cleaning) are dropped.
func SanitizeDocuments(docs []schema.Document) []schema.Document {
	out := make([]schema.Document, 0, len(docs))
	for _, d := range docs {
		// Sanitize content
		content := SanitizeUTF8(d.PageContent)
		content = strings.TrimSpace(content)
		if content == "" {
			continue
		}

		// Sanitize string metadata at top level
		if d.Metadata != nil {
			for k, v := range d.Metadata {
				switch vv := v.(type) {
				case string:
					d.Metadata[k] = SanitizeUTF8(vv)
				}
			}
		}

		d.PageContent = content
		out = append(out, d)
	}
	return out
}
