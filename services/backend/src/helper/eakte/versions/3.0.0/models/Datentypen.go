package models

// AllgemeineRueckmeldungCodeType is Der Datentyp zur Werteliste von Rückmeldungen für Fachverfahrensnachrichten unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type AllgemeineRueckmeldungCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              any    `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// AnschriftstypCodeType is Der Datentyp zur Werteliste von Anschriftentypen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type AnschriftstypCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// AussonderungsartCodeType is Der Datentyp zur Werteliste von Aussonderungsarten unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type AussonderungsartCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// BeteiligungsstatusCodeType is Der Datentyp zur Werteliste von Beteiligungsstatus unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes. Für die Umsetzung in den Systemen wird empfohlen, sofern kein anderer Wert hinterlegt wurde, als Standardwert den Wert 001 für "Der Schritt liegt in der Zukunft." anzugeben.
type BeteiligungsstatusCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// BewertungsvorschlagCodeType is Der Datentyp zur Werteliste von Bewertungsvorschlägen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type BewertungsvorschlagCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// DateiformatCodeType is Der Datentyp zur Werteliste von Dateiformaten unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type DateiformatCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr"`
	Code              string `xml:"code"`
}

// DatenaustauschartCodeType is Der Datentyp zur Werteliste von Datenaustauscharten unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type DatenaustauschartCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// DatentypCodeType is Der Datentyp zur Werteliste von Datentypen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type DatentypCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// KommunikationsartCodeType is Der Datentyp zur Werteliste von Kommunikationsarten unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type KommunikationsartCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr"`
	Code              string `xml:"code"`
}

// KompressionsverfahrenCodeType is Der Datentyp zur Werteliste von Kompressionsverfahren unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type KompressionsverfahrenCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// KonfigurationsparameterCodeType is Der Datentyp zur Werteliste von Konfigurationsparametern unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type KonfigurationsparameterCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// KonfigurationsparameterGruppeCodeType is Der Datentyp zur Werteliste von Konfigurationsparametergruppen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type KonfigurationsparameterGruppeCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// MediumCodeType is Der Datentyp zur Werteliste von Medientypen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type MediumCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// NachrichtentypCodeType is Der Datentyp zur Werteliste von Nachrichtentypen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type NachrichtentypCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SchriftgutobjekttypCodeType is Der Datentyp zur Werteliste von Schriftgutobjekttypen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SchriftgutobjekttypCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SonstigeFehlermeldungCodeType is Der Datentyp zu sonstigen Fehlermeldungen aus einer Codeliste für die Importbestätigungen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SonstigeFehlermeldungCodeType struct {
	ListURIAttr       string `xml:"listURI,attr"`
	ListVersionIDAttr string `xml:"listVersionID,attr"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0601CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0601 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0601CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0602CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0602 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0602CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0603CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0603 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0603CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0604CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0604 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0604CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0605CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0605 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0605CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0606CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0606 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0606CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0607CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0607 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0607CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0608CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0608 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0608CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0609CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0609 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0609CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0610CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0610 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0610CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0611CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0611 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0611CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0612CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0612 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0612CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0613CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0613 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0613CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0614CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0614 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0614CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0615CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0615 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0615CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0616CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0616 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0616CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0617CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0617 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0617CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0618CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0618 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0618CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0619CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0619 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0619CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0620CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0620 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0620CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0621CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0621 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0621CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0622CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0622 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0622CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0623CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0623 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0623CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0624CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0624 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0624CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0625CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0625 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0625CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0626CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0626 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0626CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0627CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0627 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0627CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0628CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0628 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0628CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0631CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0631 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0631CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// SpezifischeFehlermeldung0632CodeType is Der Datentyp zur Werteliste von spezifischen Fehlermeldungen für die Nachricht 0632 unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type SpezifischeFehlermeldung0632CodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// StaatCodeType is Der Datentyp zur Werteliste von Staaten unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes. Genutzt wird hier die ISO-3166-1-Liste. Für den Schlüssel (code) wird die Spalte "ALPHA2" und für den beschreibenden Namen (name) die Spalte "Name des Landes" verwendet.
type StaatCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// TransportwegCodeType is Der Datentyp zur Werteliste von Transportwegen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type TransportwegCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// VertraulichkeitsstufeCodeType is Der Datentyp zur Werteliste von Vertraulichkeitsstufen unter Angabe weiterer Informationen zur konkreten Angabe eines Wertes.
type VertraulichkeitsstufeCodeType struct {
	ListURIAttr       string `xml:"listURI,attr,omitempty"`
	ListVersionIDAttr string `xml:"listVersionID,attr,omitempty"`
	Code              string `xml:"code"`
	Name              string `xml:"name,omitempty"`
}

// StringDateinameType is Der Datentyp zur Angabe eines Dateinamens.
type StringDateinameType string

// StringUUIDType is Der Datentyp zur Angabe einer UUID.
type StringUUIDType string
