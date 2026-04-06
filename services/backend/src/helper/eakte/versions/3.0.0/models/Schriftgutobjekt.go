package models

// Schriftgutobjekt is Ein Dokument.
type Schriftgutobjekt struct {
	Akte     *AkteType     `xml:"Akte"`
	Vorgang  *VorgangType  `xml:"Vorgang"`
	Dokument *DokumentType `xml:"Dokument"`
}
