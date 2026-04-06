package models

import (
	"encoding/xml"
)

// InformationInformation0101 is Ein Dokument, das dem Empfänger zur Erläuterung der Information mitgegeben wird.
type InformationInformation0101 struct {
	XMLName          xml.Name                                                    `xml:"Information.Information.0101"`
	Kopf             *NkNichtFVDatenWeitereEmpfaengerMitEmpfangsbestaetigungType `xml:"Kopf"`
	Anschreiben      *DokumentType                                               `xml:"Anschreiben,omitempty"`
	Schriftgutobjekt []*Schriftgutobjekt                                         `xml:"Schriftgutobjekt"`
}

// InformationEmpfangBestaetigen0102 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type InformationEmpfangBestaetigen0102 struct {
	XMLName xml.Name     `xml:"Information.EmpfangBestaetigen.0102"`
	Kopf    *NkBasisType `xml:"Kopf"`
}
