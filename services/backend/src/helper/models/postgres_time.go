package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type PostgresJsonTime struct {
	time.Time
}

var timeFormats = []string{
	time.RFC3339Nano,                   // "2006-01-02T15:04:05.999999999Z07:00"
	time.RFC3339,                       // "2006-01-02T15:04:05Z07:00"
	"2006-01-02T15:04:05.999999999+00", // postgres without colon (no tz offset)
	"2006-01-02T15:04:05+00",           // postgres without colon (no tz offset)
}

func (ft *PostgresJsonTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "null" {
		ft.Time = time.Time{}
		return nil
	}

	s = normalizeTimezone(s)

	for _, layout := range timeFormats {
		if t, err := time.Parse(layout, s); err == nil {
			ft.Time = t
			return nil
		}
	}

	return fmt.Errorf("FlexTime: cannot parse %q", s)
}

// normalizeTimezone turns a bare "+HH" or "-HH" suffix into "+HH:00"
func normalizeTimezone(s string) string {
	// Matches +HH or -HH at end of string (no colon, no minutes)
	re := regexp.MustCompile(`([+-]\d{2})$`)
	return re.ReplaceAllString(s, "$1:00")
}
