package models

// ErfolgOderMisserfolgAbgabeType is Die Information in Form eines Codes, aus welchem Grund der Import zum abzugebenden Schriftgutobjekt nicht erfolgreich war.
type ErfolgOderMisserfolgAbgabeType struct {
	IDSGO             string                           `xml:"IDSGO"`
	Erfolgreich       bool                             `xml:"Erfolgreich"`
	Fehlermeldung     []string                         `xml:"Fehlermeldung,omitempty"`
	FehlermeldungCode []*SonstigeFehlermeldungCodeType `xml:"FehlermeldungCode,omitempty"`
}

// NkAbgabeType is Die Angabe, ob vom Empfänger der Nachricht eine Empfangsbestätigung erwünscht wird. Eine erwünschte Empfangsbestätigung wird mit 1 gekennzeichnet. Wird keine Empfangsbestätigung gewünscht, so wird dies mit 0 gekennzeichnet.
type NkAbgabeType struct {
	Importbestaetigung   bool `xml:"Importbestaetigung"`
	Empfangsbestaetigung bool `xml:"Empfangsbestaetigung"`
	*NkNichtFVDatenWeitereEmpfaengerType
}
