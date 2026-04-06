package models

import (
	"encoding/xml"
)

// Anlage is Ein Dokument.
type Anlage struct {
	Akte     *AkteType     `xml:"Akte"`
	Vorgang  *VorgangType  `xml:"Vorgang"`
	Dokument *DokumentType `xml:"Dokument"`
}

// GeschaeftsgangGeschaeftsgang0201 is Im externen Geschäftsgang werden die Verfügungen an externe Bearbeitungsstationen festgelegt.
type GeschaeftsgangGeschaeftsgang0201 struct {
	XMLName                xml.Name              `xml:"Geschaeftsgang.Geschaeftsgang.0201"`
	Kopf                   *NkGeschaeftsgangType `xml:"Kopf"`
	Anschreiben            *DokumentType         `xml:"Anschreiben,omitempty"`
	Hauptobjekt            *HauptobjektType      `xml:"Hauptobjekt"`
	ExternerGeschaeftsgang *GeschaeftsgangType   `xml:"ExternerGeschaeftsgang"`
	Anlage                 []*Anlage             `xml:"Anlage,omitempty"`
}

// GeschaeftsgangEmpfangBestaetigen0202 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type GeschaeftsgangEmpfangBestaetigen0202 struct {
	XMLName xml.Name                          `xml:"Geschaeftsgang.EmpfangBestaetigen.0202"`
	Kopf    *NkRueckmeldungGeschaeftsgangType `xml:"Kopf"`
}

// GeschaeftsgangGeaendertenLaufwegMitteilen0203 is Informationen zu geänderten Beteiligungsschritten eines Geschäftsgangs.
type GeschaeftsgangGeaendertenLaufwegMitteilen0203 struct {
	XMLName                        xml.Name                              `xml:"Geschaeftsgang.GeaendertenLaufwegMitteilen.0203"`
	Kopf                           *NkBasisType                          `xml:"Kopf"`
	GeaenderterGeschaeftsgang      *GeschaeftsgangType                   `xml:"GeaenderterGeschaeftsgang"`
	GeaenderterBeteiligungsschritt []*GeaenderterBeteiligungsschrittType `xml:"GeaenderterBeteiligungsschritt"`
}
