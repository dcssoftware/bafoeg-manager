package filename

import (
	"fmt"
	"path/filepath"
	"regexp"
)

// ParsedName holds the structured parts of an xdomea file name.
type ParsedName struct {
	ProzessID         string // UUID
	Nachrichtengruppe string // e.g., Geschaeftsgang
	Nachrichtenname   string // e.g., Geschaeftsgang
	MessageType       string // 4 digits, e.g., 0201
	Nachrichtentyp    string // full "<Gruppe>.<Name>.<Nummer>"
	Extension         string // e.g., zip, xdomea, xml, txt, ...
}

// Regex per spec: "<UUID>_<Group>.<Name>.<NNNN>.<ext>"
var re = regexp.MustCompile(`^(?i)` +
	`([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})` + // ProzessID (UUID)
	`_` +
	`(([\p{L}]+)\.([\p{L}]+)\.(\d{4}))` + // Nachrichtentyp = Gruppe.Name.Nummer
	`\.(.+)$`) // extension (any)

// Regex for "<UUID>_<NNNN>_<lfdNr>.<ext>"
var reSimple = regexp.MustCompile(`^(?i)` +
	`([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})` + // ProzessID (UUID)
	`_` +
	`(\d{4})` + // Nachrichtennummer
	`_(\d+)` + // lfdNrNachrichtProTyp
	`\.(.+)$`) // extension

// ParseContainerDateiname parses container/package names (*.zip | *.xdomea).
func ParseContainerDateiname(name string) (*ParsedName, error) {
	p := filepath.Base(name)
	m := re.FindStringSubmatch(p)
	if m != nil {
		ext := m[6]
		if ext != "zip" && ext != "xdomea" {
			return nil, fmt.Errorf("unerwartete Container-Erweiterung %q (erwartet zip/xdomea)", ext)
		}
		return &ParsedName{
			ProzessID:         m[1],
			Nachrichtentyp:    m[2],
			Nachrichtengruppe: m[3],
			Nachrichtenname:   m[4],
			MessageType:       m[5],
			Extension:         ext,
		}, nil
	}
	// Try simple pattern
	m2 := reSimple.FindStringSubmatch(p)
	if m2 != nil {
		ext := m2[4]
		if ext != "zip" && ext != "xdomea" {
			return nil, fmt.Errorf("unerwartete Container-Erweiterung %q (erwartet zip/xdomea)", ext)
		}
		return &ParsedName{
			ProzessID:   m2[1],
			MessageType: m2[2],
			Extension:   ext,
			// Nachrichtentyp, Nachrichtengruppe, Nachrichtenname not available in this format
		}, nil
	}
	return nil, fmt.Errorf("kein xdomea-Containerdateiname: %q", p)
}

// ParseTransportDateiname parses the optional transport file name.
// The spec allows a transport file but does not fix its extension;
// we therefore accept any extension while enforcing the same stem.
func ParseTransportDateiname(name string) (*ParsedName, error) {
	p := filepath.Base(name)
	m := re.FindStringSubmatch(p)
	if m == nil {
		return nil, fmt.Errorf("kein xdomea-Transportdateiname: %q", p)
	}
	return &ParsedName{
		ProzessID:         m[1],
		Nachrichtentyp:    m[2],
		Nachrichtengruppe: m[3],
		Nachrichtenname:   m[4],
		MessageType:       m[5],
		Extension:         m[6],
	}, nil
}
