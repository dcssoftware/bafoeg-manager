package models

import (
	"encoding/xml"
)

// AbgabeAbgabe0401 is Ein Dokument, das dem Empfänger zur Erläuterung der Abgabe mitgegeben wird.
type AbgabeAbgabe0401 struct {
	XMLName          xml.Name            `xml:"Abgabe.Abgabe.0401"`
	Kopf             *NkAbgabeType       `xml:"Kopf"`
	Anschreiben      *DokumentType       `xml:"Anschreiben,omitempty"`
	Schriftgutobjekt []*Schriftgutobjekt `xml:"Schriftgutobjekt"`
	Aktenplan        *AktenplanType      `xml:"Aktenplan,omitempty"`
}

// AbgabeImportBestaetigen0402 is Die Information zum erfolgreichen oder nicht erfolgreichen Import eines Schriftgutobjekts zur Abgabe.
type AbgabeImportBestaetigen0402 struct {
	XMLName           xml.Name                          `xml:"Abgabe.ImportBestaetigen.0402"`
	Kopf              *NkBasisType                      `xml:"Kopf"`
	AbgegebenesObjekt []*ErfolgOderMisserfolgAbgabeType `xml:"AbgegebenesObjekt"`
}

// AbgabeEmpfangBestaetigen0403 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type AbgabeEmpfangBestaetigen0403 struct {
	XMLName xml.Name     `xml:"Abgabe.EmpfangBestaetigen.0403"`
	Kopf    *NkBasisType `xml:"Kopf"`
}
