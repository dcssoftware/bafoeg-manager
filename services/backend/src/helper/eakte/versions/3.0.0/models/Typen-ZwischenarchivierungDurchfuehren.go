package models

// ErfolgOderMisserfolgZwischenarchivierungType is Die Information in Form eines Codes, aus welchem Grund der Import zum auszulagernden Schriftgutobjekt nicht erfolgreich war.
type ErfolgOderMisserfolgZwischenarchivierungType struct {
	IDSGO             string                           `xml:"IDSGO"`
	Erfolgreich       bool                             `xml:"Erfolgreich"`
	Fehlermeldung     []string                         `xml:"Fehlermeldung,omitempty"`
	FehlermeldungCode []*SonstigeFehlermeldungCodeType `xml:"FehlermeldungCode,omitempty"`
}

// NkZwischenarchivierungType is Die Angabe, ob vom Empfänger der Nachricht eine Empfangsbestätigung erwünscht wird. Eine erwünschte Empfangsbestätigung wird mit 1 gekennzeichnet. Wird keine Empfangsbestätigung gewünscht, so wird dies mit 0 gekennzeichnet.
type NkZwischenarchivierungType struct {
	Importbestaetigung   bool `xml:"Importbestaetigung"`
	Empfangsbestaetigung bool `xml:"Empfangsbestaetigung"`
	*NkNichtFVDatenWeitereEmpfaengerType
}
