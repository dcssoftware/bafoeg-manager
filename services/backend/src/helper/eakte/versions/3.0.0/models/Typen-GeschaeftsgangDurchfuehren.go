package models

// Aenderungsart is Der Beteiligungsschritt wird in dem Laufweg als übersprungen markiert, d.h. der Schritt bleibt im Geschäftsgang erhalten, wird jedoch nicht mehr aktiv beteiligt.
type Aenderungsart struct {
	Neu           bool `xml:"Neu"`
	Uebersprungen bool `xml:"Uebersprungen"`
}

// GeaenderterBeteiligungsschrittType is Die Nummer des Beteiligungsschrittes innerhalb der fortlaufenden Nummerierung im Geschäftsgang. Die Nummerierung ist innerhalb eines xdomea-Geschäftsgangsobjekts eindeutig. Werden in den xdomea-Geschäftsgang neue Bearbeitungsschritte eingefügt, werden dementsprechend die Nummern der nachfolgenden Beteiligungsschritte verändert. Beteiligungsschritte, die abgeschlossen sind, sind unveränderlich.
type GeaenderterBeteiligungsschrittType struct {
	Nummer        string         `xml:"Nummer"`
	Aenderungsart *Aenderungsart `xml:"Aenderungsart"`
}

// HauptobjektType is Ein in der Geschäftsgangs-Nachricht enthaltenes Dokument, auf das sich die Verfügungen im externen Geschäftsgang beziehen.
type HauptobjektType struct {
	Dokument []*DokumentType `xml:"Dokument"`
}

// NkGeschaeftsgangType is Die Angabe, ob vom Empfänger der Nachricht eine Empfangsbestätigung für den Absender der Nachricht erwünscht wird. Eine erwünschte Empfangsbestätigung wird mit 1 gekennzeichnet. Wird keine Empfangsbestätigung gewünscht, so wird dies mit 0 gekennzeichnet.
type NkGeschaeftsgangType struct {
	EmpfangsbestaetigungAnInitiator  bool `xml:"EmpfangsbestaetigungAnInitiator"`
	EmpfangsbestaetigungAnVorgaenger bool `xml:"EmpfangsbestaetigungAnVorgaenger"`
	*NkBasisType
}

// NkRueckmeldungGeschaeftsgangType is Der weitere Empfänger der Nachricht.
type NkRueckmeldungGeschaeftsgangType struct {
	WeitererEmpfaenger *KontaktType `xml:"WeitererEmpfaenger,omitempty"`
	*NkBasisType
}
